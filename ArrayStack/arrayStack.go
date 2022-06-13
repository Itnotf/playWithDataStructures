package ArrayStack

import (
	"fmt"

	array "playWithDataStructures/Array"
)

type arrayStack[T any] struct {
	data *array.Array[T]
}

func NewArrayStack[T any](capacity int) *arrayStack[T] {
	return &arrayStack[T]{
		data: array.NewArray[T](capacity),
	}
}

//Push TODO:此处用空接口是因为用泛型就不算实现了stack接口，不知道为啥
func (s arrayStack[T]) Push(e T) {
	s.data.AddLast(e)
}

func (s *arrayStack[T]) Pop() interface{} {
	if s.IsEmpty() {
		panic("empty stack")
	}
	return s.data.RemoveLast()
}

func (s arrayStack[T]) Peek() interface{} {
	return s.data.Get(s.data.GetSize() - 1)
}

func (s arrayStack[T]) GetSize() int {
	return s.data.GetSize()
}

func (s arrayStack[T]) GetCapacity() int {
	return s.data.GetCapacity()
}

func (s arrayStack[T]) IsEmpty() bool {
	return s.GetSize() == 0
}

func (s *arrayStack[T]) ToString() string {
	return fmt.Sprintf("Size:%d,Capacity:%d\n %v top", s.data.GetSize(), s.data.GetCapacity(), s.data)
}
