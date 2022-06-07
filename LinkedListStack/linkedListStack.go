package linkedListStack

import (
	"fmt"
	linkedList "playWithDataStructures/LinkedList"
)

type linkedListStack[T any] struct {
	data *linkedList.LinkedList[T]
}

func NewLinkedListStack[T any]() *linkedListStack[T] {
	return &linkedListStack[T]{
		data: linkedList.NewLinkedList[T](),
	}
}

func (l *linkedListStack[T]) ToString() string {
	return fmt.Sprintf("top:%s", l.data.ToString())
}

func (l *linkedListStack[T]) IsEmpty() bool {
	return l.GetSize() == 0
}

func (l *linkedListStack[T]) GetSize() int {
	return l.data.GetSize()
}

func (l *linkedListStack[T]) Push(e T) {
	l.data.AddFirst(e)
}

func (l *linkedListStack[T]) Pop() T {
	return l.data.RemoveFirst()
}

func (l *linkedListStack[T]) Peek() T {
	return l.data.Get(0)
}
