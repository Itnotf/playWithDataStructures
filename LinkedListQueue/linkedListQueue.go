package linkedListQueue

import (
	"fmt"
	linkedList "playWithDataStructures/LinkedList"
)

type linkedListQueue[T comparable] struct {
	data *linkedList.LinkedList[T]
}

func NewLinkedListQueue[T comparable]() *linkedListQueue[T] {
	return &linkedListQueue[T]{
		data: linkedList.NewLinkedList[T](),
	}
}

func (q *linkedListQueue[T]) GetSize() int {
	return q.data.GetSize()
}

func (q *linkedListQueue[T]) IsEmpty() bool {
	return q.data.GetSize() == 0
}
func (q *linkedListQueue[T]) ToString() string {
	return fmt.Sprintf("Front:%s:End", q.data.ToString())
}

func (q *linkedListQueue[T]) Enqueue(e T) {
	q.data.AddLast(e)
}

func (q *linkedListQueue[T]) Dequeue() T {
	return q.data.RemoveFirst()
}

func (q *linkedListQueue[T]) GetFront() T {
	return q.data.Get(0)
}
