package sync

type Hash interface {
	comparable
	Hash() uint64
}

type BucketMap[K Hash, V any] struct {
	shards  uint64
	buckets []*Map[K, V]
}

func (m *BucketMap[K, V]) bucket(k K) *Map[K, V] {
	return m.buckets[k.Hash()%m.shards]
}

// Load returns the value stored in the map for a key, or nil if no
// value is present.
// The ok result indicates whether value was found in the map.
func (m *BucketMap[K, V]) Load(key K) (value V, ok bool) {
	return m.bucket(key).Load(key)
}

// Store sets the value for a key.
func (m *BucketMap[K, V]) Store(key K, value V) {
	m.bucket(key).Store(key, value)
}

// Swap swaps the value for a key and returns the previous value if any.
// The loaded result reports whether the key was present.
func (m *BucketMap[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	return m.bucket(key).Swap(key, value)
}

// CompareAndSwap swaps the old and new values for key
// if the value stored in the map is equal to old.
// The old value must be of a comparable type.
func (m *BucketMap[K, V]) CompareAndSwap(key K, old, new V) bool {
	return m.bucket(key).CompareAndSwap(key, old, new)
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
//
// Range does not necessarily correspond to any consistent snapshot of the Map's
// contents: no key will be visited more than once, but if the value for any key
// is stored or deleted concurrently (including by f), Range may reflect any
// mapping for that key from any point during the Range call. Range does not
// block other methods on the receiver; even f itself may call any method on m.
//
// Range may be O(N) with the number of elements in the map even if f returns
// false after a constant number of calls.
func (m *BucketMap[K, V]) Range(f func(key K, value V) bool) {
	continueRange := true
	for _, bucket := range m.buckets {
		bucket._map.Range(func(key, value any) bool {
			continueRange = f(key.(K), value.(V))
			return continueRange
		})
		if !continueRange {
			return
		}
	}
}

// Delete deletes the value for a key.
func (m *BucketMap[K, V]) Delete(key K) {
	m.bucket(key).Delete(key)
}

// CompareAndDelete deletes the entry for key if its value is equal to old.
// The old value must be of a comparable type.
//
// If there is no current value for key in the map, CompareAndDelete
// returns false (even if the old value is the nil interface value).
func (m *BucketMap[K, V]) CompareAndDelete(key K, old V) bool {
	return m.bucket(key).CompareAndDelete(key, old)
}

// LoadAndDelete deletes the value for a key, returning the previous value if any.
// The loaded result reports whether the key was present.
func (m *BucketMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	return m.bucket(key).LoadAndDelete(key)
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (m *BucketMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	return m.bucket(key).LoadOrStore(key, value)
}
