package linkedListQueue

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedListQueue(t *testing.T) {
	q := NewLinkedListQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	assert.Equal(t, 1, q.GetFront())
	assert.Equal(t, 2, q.GetSize())
	assert.Equal(t, false, q.IsEmpty())
	fmt.Println(q.ToString())

	assert.Equal(t, 1, q.Dequeue())
	assert.Equal(t, 2, q.Dequeue())
	assert.Equal(t, true, q.IsEmpty())
	assert.Panics(t, func() { q.Dequeue() })
	fmt.Println(q.ToString())
	for i := 3; i < 15; i++ {
		q.Enqueue(i)
	}
	fmt.Println(q.ToString())
}
