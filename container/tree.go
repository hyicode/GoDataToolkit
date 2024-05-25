package container

import "slices"

type Tree[T any] struct {
	value    T
	parent   *Tree[T]
	children []*Tree[T]
}

func NewTree[T any](value T) *Tree[T] {
	return &Tree[T]{
		value: value,
	}
}

func (tree *Tree[T]) SetParent(parent *Tree[T]) {
	tree.parent = parent
}

func (tree *Tree[T]) Parent() *Tree[T] {
	return tree.parent
}

func (tree *Tree[T]) Children() []*Tree[T] {
	return tree.children
}

func (tree *Tree[T]) AddChild(node *Tree[T]) {
	tree.children = append(tree.children, node)
	node.parent = tree
}

func (tree *Tree[T]) RemoveChild(f func(node *Tree[T]) bool) {
	tree.children = slices.DeleteFunc(tree.children, f)
}

func (tree *Tree[T]) SetValue(value T) {
	tree.value = value
}

func (tree *Tree[T]) Value() T {
	return tree.value
}

func (tree *Tree[T]) IsRoot() bool {
	return tree.parent == nil
}

func (tree *Tree[T]) IsLeaf() bool {
	return len(tree.children) == 0
}

func (tree *Tree[T]) TraversalPreorder(f func(node *Tree[T]) (stop bool)) {
	if f(tree) {
		return
	}
	for _, child := range tree.children {
		child.TraversalPreorder(f)
	}
}

func (tree *Tree[T]) TraversalPostorder(f func(node *Tree[T]) (stop bool)) {
	for _, child := range tree.children {
		child.TraversalPostorder(f)
	}
	if f(tree) {
		return
	}
}

func (tree *Tree[T]) TraversalBreadthFirst(f func(node *Tree[T]) (stop bool)) {
	if f(tree) {
		return
	}
	list := []*Tree[T]{tree}
	next := make([]*Tree[T], 0, len(tree.children))
	for len(list) > 0 {
		size := 0
		for _, node := range list {
			if f(node) {
				return
			}
			for _, child := range node.children {
				next = append(next, child)
				size += len(child.children)
			}
		}
		list = next
		next = make([]*Tree[T], 0, size)
	}
}
