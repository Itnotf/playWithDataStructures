package PriorityQueue

import "playWithDataStructures/Heap"

type PriorityQueue[T int | float32] struct {
	data *Heap.MaxHeap[T]
	size int
}

func NewPriorityQueue[T int | float32]() *PriorityQueue[T] {
	return &PriorityQueue[T]{
		data: Heap.NewMaxHeap[T](),
		size: 0,
	}
}

func (p *PriorityQueue[T]) GetSize() int {
	return p.size
}
func (p *PriorityQueue[T]) IsEmpty() bool {
	return p.size == 0
}
func (p *PriorityQueue[T]) Enqueue(e T) {
	p.data.Add(e)
	p.size++
}

func (p *PriorityQueue[T]) Dequeue() T {
	p.size--
	return p.data.ExtractMax()
}
func (p *PriorityQueue[T]) GetFront() T {
	return p.data.FindMax()
}
