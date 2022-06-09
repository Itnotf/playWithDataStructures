package Heap

import (
	array "playWithDataStructures/Array"
)

type MaxHeap[K int | float32] struct {
	data *array.Array[K]
	size int
}

func NewMaxHeap[K int | float32]() *MaxHeap[K] {
	return &MaxHeap[K]{
		data: array.NewArray[K](10),
		size: 0,
	}
}

func (m *MaxHeap[K]) GetSize() int {
	return m.size
}

func (m *MaxHeap[K]) IsEmpty() bool {
	return m.size == 0
}

func (m MaxHeap[K]) parent(index int) int {
	return (index - 1) / 2
}
func (m *MaxHeap[K]) leftChild(index int) int {
	return 2*index + 1
}
func (m *MaxHeap[K]) rightChild(index int) int {
	return 2*index + 2
}
func (m *MaxHeap[K]) Add(e K) {
	m.data.AddLast(e)
	m.siftUp()
	m.size++
}
func (m *MaxHeap[K]) siftUp() {
	cur := m.size - 1
	for cur > 0 {
		if m.data.Get(cur) > m.data.Get(m.parent(cur)) {
			m.swap(cur, m.parent(cur))
			cur = m.parent(cur)
		} else {
			break
		}
	}
}

func (m *MaxHeap[K]) findMax() K {
	if m.size == 0 {
		panic("empty heap")
	}
	return m.data.Get(0)
}

func (m *MaxHeap[K]) ExtractMax() K {
	max := m.findMax()

	m.swap(0, m.size-1)
	m.data.RemoveLast()
	m.siftDown(0)
	m.size--

	return max
}

func (m *MaxHeap[K]) siftDown(cur int) {
	for m.leftChild(cur) < m.size {
		left := m.leftChild(cur)
		right := m.rightChild(cur)
		if right < m.size {
			if m.data.Get(right) > m.data.Get(left) && m.data.Get(right) > m.data.Get(cur) {
				m.swap(cur, right)
				cur = right
				break
			}
		}
		if m.data.Get(left) > m.data.Get(cur) {
			m.swap(cur, left)
			cur = left
		} else {
			break
		}
	}
}

func (m *MaxHeap[K]) swap(i int, j int) {
	tmp := m.data.Get(j)
	m.data.Set(j, m.data.Get(i))
	m.data.Set(i, tmp)
}

func (m *MaxHeap[K]) Replace(e K) K {
	max := m.findMax()

	m.data.Set(0, e)
	m.siftDown(0)

	return max
}

func (m *MaxHeap[K]) Heapify(arr *array.Array[K]) {
	m.data = arr
	m.size = arr.GetSize()
	lastParent := (m.size - 2) / 2
	for i := lastParent; i >= 0; i-- {
		m.siftDown(i)
	}
}
