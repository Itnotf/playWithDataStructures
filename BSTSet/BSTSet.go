package BSTSet

import "playWithDataStructures/BST"

type BSTSet[T int | float32 | ~string] struct {
	data *BST.BST[T]
	size int
}

func NewBSTSet[T int | float32 | ~string]() *BSTSet[T] {
	return &BSTSet[T]{
		data: BST.NewBST[T](),
		size: 0,
	}
}

func (b *BSTSet[T]) GetSize() int {
	return b.size
}

func (b *BSTSet[T]) IsEmpty() bool {
	return b.size == 0
}

func (b *BSTSet[T]) Add(e T) {
	if !b.data.Contains(e) {
		b.data.Add(e)
		b.size++
	}
}

func (b *BSTSet[T]) Remove(e T) {
	if b.size == 0 {
		panic("empty set")
	}
	if !b.data.Contains(e) {
		panic("no such element to remove")
	}
	b.data.Remove(e)
	b.size--
}
func (b *BSTSet[T]) Contains(e T) bool {
	return b.data.Contains(e)
}
