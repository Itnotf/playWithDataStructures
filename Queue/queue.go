package queue

type queue[T any] interface {
	Enqueue(T)
	Dequeue() T
	GetFront() T
	GetSize() int
	IsEmpty() bool
}
