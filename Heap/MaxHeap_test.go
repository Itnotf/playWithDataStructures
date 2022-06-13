package Heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
	array "playWithDataStructures/Array"
)

func TestMaxHeap(t *testing.T) {
	maxHeap := NewMaxHeap[int]()
	assert.Equal(t, true, maxHeap.IsEmpty())
	assert.Equal(t, 0, maxHeap.GetSize())
	assert.Panics(t, func() { maxHeap.ExtractMax() })
	maxHeap.Add(7)
	maxHeap.Add(5)
	maxHeap.Add(9)
	maxHeap.Add(1)
	maxHeap.Add(16)
	maxHeap.Add(1620)
	maxHeap.Add(3)
	maxHeap.Add(2)
	assert.Equal(t, 1620, maxHeap.ExtractMax())
	assert.Equal(t, 16, maxHeap.ExtractMax())
	assert.Equal(t, 9, maxHeap.ExtractMax())
	assert.Equal(t, 5, maxHeap.GetSize())
	assert.Equal(t, 7, maxHeap.Replace(4))
	assert.Equal(t, 5, maxHeap.ExtractMax())
	arr := array.NewArray[int](10)
	arr.AddLast(7)
	arr.AddLast(5)
	arr.AddLast(9)
	arr.AddLast(1)
	arr.AddLast(16)
	arr.AddLast(1620)
	arr.AddLast(3)
	maxHeap.Heapify(arr)
	assert.Equal(t, 1620, maxHeap.ExtractMax())
	assert.Equal(t, 16, maxHeap.ExtractMax())
	assert.Equal(t, 9, maxHeap.ExtractMax())
	assert.Equal(t, 4, maxHeap.GetSize())
	assert.Equal(t, 7, maxHeap.Replace(4))
	assert.Equal(t, 5, maxHeap.ExtractMax())
}
