package stack

type stack[T any] interface {
	Push(T)
	Pop() T
	Peek()
	GetSize() int
	IsEmpty() bool
}
