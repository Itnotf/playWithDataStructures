package PriorityQueue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	p := NewPriorityQueue[int]()
	p.Enqueue(89)
	p.Enqueue(9)
	p.Enqueue(4)
	p.Enqueue(88)
	p.Enqueue(22)
	p.Enqueue(1600)
	p.Enqueue(1)
	p.Enqueue(6)
	p.Enqueue(3)
	assert.Equal(t, 9, p.GetSize())
	assert.Equal(t, false, p.IsEmpty())
	assert.Equal(t, 1600, p.Dequeue())
	assert.Equal(t, 89, p.GetFront())
}
