package arrayQueue

import (
	"fmt"
	"playWithDataStructures/array"
)

type ArrayQueue[T any] struct {
	data *array.Array[T]
}

func NewArrayQueue[T any](capacity int) *ArrayQueue[T] {
	return &ArrayQueue[T]{
		data: array.NewArray[T](capacity),
	}
}

func (q *ArrayQueue[T]) Enqueue(e T) {
	q.data.AddLast(e)
}
func (q *ArrayQueue[T]) Dequeue() T {
	if q.IsEmpty() {
		panic("queue is empty")
	}
	return q.data.RemoveFirst()
}
func (q *ArrayQueue[T]) GetFront() T {
	return q.data.Get(0)
}
func (q *ArrayQueue[T]) GetSize() int {
	return q.data.GetSize()
}
func (q *ArrayQueue[T]) IsEmpty() bool {
	return q.data.GetSize() == 0
}
func (q *ArrayQueue[T]) ToString() string {
	return fmt.Sprintf("Size:%d,Capacity:%d\nfront %v tail", q.data.GetSize(), q.data.GetCapacity(), q.data)
}
