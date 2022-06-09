package LinkedListSet

import linkedList "playWithDataStructures/LinkedList"

type linkedListSet[T comparable] struct {
	data *linkedList.LinkedList[T]
	size int
}

func NewLinkedListSet[T comparable]() *linkedListSet[T] {
	return &linkedListSet[T]{
		data: linkedList.NewLinkedList[T](),
		size: 0,
	}
}

func (l *linkedListSet[T]) GetSize() int {
	return l.data.GetSize()
}
func (l *linkedListSet[T]) IsEmpty() bool {
	return l.size == 0
}
func (l *linkedListSet[T]) Add(e T) {
	if !l.data.Contains(e) {
		l.data.AddFirst(e)
		l.size++
	}
}
func (l *linkedListSet[T]) Remove(e T) {
	if l.size == 0 {
		panic("empty set")
	}
	if !l.data.Contains(e) {
		panic("no such element to remove")
	}
	l.data.RemoveElement(e)
	l.size--
}
func (l *linkedListSet[T]) Contains(e T) bool {
	return l.data.Contains(e)
}
