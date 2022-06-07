package loopQueue

import (
	"fmt"
)

type LoopQueue[T any] struct {
	data []T
	head int
	tail int
	size int
}

func NewLoopQueue[T any](capacity int) *LoopQueue[T] {
	return &LoopQueue[T]{
		data: make([]T, capacity+1),
		head: 0,
		tail: 0,
		size: 0,
	}
}

func (l *LoopQueue[T]) Enqueue(e T) {
	if (l.tail+1)%cap(l.data) == l.head {
		l.resize(2 * l.GetCapacity())
	}
	l.data[l.tail] = e
	l.tail = (l.tail + 1) % cap(l.data)
	l.size++
}

func (l *LoopQueue[T]) Dequeue() T {
	if l.size == 0 {
		panic("queue is empty")
	}
	e := l.data[l.head]
	l.head = (l.head + 1) % cap(l.data)
	l.size--
	if (l.GetSize() == l.GetCapacity()/4) && (l.GetCapacity()/2 != 0) {
		l.resize(l.GetCapacity() / 2)
	}
	return e
}

func (l *LoopQueue[T]) GetFront() T {
	return l.data[l.head]
}

func (l *LoopQueue[T]) GetSize() int {
	return l.size
}

func (l *LoopQueue[T]) GetCapacity() int {
	return cap(l.data) - 1
}

func (l *LoopQueue[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *LoopQueue[T]) resize(newCapacity int) {
	newData := make([]T, newCapacity)
	for i := 0; i < l.size; i++ {
		newData[i] = l.data[(l.head+i)%cap(l.data)]
	}
	l.data = newData
	l.head = 0
	l.tail = l.size
}

func (l *LoopQueue[T]) ToString() string {
	str := fmt.Sprintf("Size:%d,Capacity:%d\n%v\n", l.GetSize(), l.GetCapacity(), l.data)
	str += "head ["
	for i := 0; i < l.GetSize(); i++ {
		str += fmt.Sprintf(" %v ", l.data[((l.head+i)%(cap(l.data)))])
	}
	str += "] tail"
	return str
}
