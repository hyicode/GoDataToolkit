package sync

import "sync"

type Map[K comparable, V any] struct {
	_map sync.Map
}

// Load returns the value stored in the map for a key, or nil if no
// value is present.
// The ok result indicates whether value was found in the map.
func (m *Map[K, V]) Load(key K) (value V, ok bool) {
	v, exist := m._map.Load(key)
	if !exist {
		var val V
		return val, false
	}
	return v.(V), true
}

// Store sets the value for a key.
func (m *Map[K, V]) Store(key K, value V) {
	m._map.Store(key, value)
}

// Swap swaps the value for a key and returns the previous value if any.
// The loaded result reports whether the key was present.
func (m *Map[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	v, ok := m._map.Swap(key, value)
	if !ok {
		var val V
		return val, false
	}
	return v.(V), true
}

// CompareAndSwap swaps the old and new values for key
// if the value stored in the map is equal to old.
// The old value must be of a comparable type.
func (m *Map[K, V]) CompareAndSwap(key K, old, new V) bool {
	return m._map.CompareAndSwap(key, old, new)
}

func (m *Map[K, V]) Range(f func(key K, value V) bool) {
	m._map.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

// Delete deletes the value for a key.
func (m *Map[K, V]) Delete(key K) {
	m._map.Delete(key)
}

// CompareAndDelete deletes the entry for key if its value is equal to old.
// The old value must be of a comparable type.
//
// If there is no current value for key in the map, CompareAndDelete
// returns false (even if the old value is the nil interface value).
func (m *Map[K, V]) CompareAndDelete(key K, old V) bool {
	return m._map.CompareAndDelete(key, old)
}

// LoadAndDelete deletes the value for a key, returning the previous value if any.
// The loaded result reports whether the key was present.
func (m *Map[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	v, exist := m._map.LoadAndDelete(key)
	if !exist {
		var val V
		return val, false
	}
	return v.(V), true
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (m *Map[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	v, exist := m._map.LoadOrStore(key, value)
	if !exist {
		return value, false
	}
	return v.(V), true
}
