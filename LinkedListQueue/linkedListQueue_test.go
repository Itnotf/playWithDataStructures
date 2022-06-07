package linkedListQueue

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedListQueue(t *testing.T) {
	q := NewLinkedListQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	assert.Equal(t, q.GetFront(), 1)
	assert.Equal(t, q.GetSize(), 2)
	assert.Equal(t, q.IsEmpty(), false)
	fmt.Println(q.ToString())

	assert.Equal(t, q.Dequeue(), 1)
	assert.Equal(t, q.Dequeue(), 2)
	assert.Equal(t, q.IsEmpty(), true)
	assert.Panics(t, func() { q.Dequeue() })
	fmt.Println(q.ToString())
	for i := 3; i < 15; i++ {
		q.Enqueue(i)
	}
	fmt.Println(q.ToString())
}
