package loopQueue

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoopQueue(t *testing.T) {
	q := NewLoopQueue[int](10)
	q.Enqueue(99)
	q.Enqueue(88)
	assert.Equal(t, 99, q.GetFront())
	assert.Equal(t, 2, q.GetSize())
	assert.Equal(t, false, q.IsEmpty())
	fmt.Println(q.ToString())

	assert.Equal(t, 99, q.Dequeue())
	assert.Equal(t, 88, q.Dequeue())
	assert.Equal(t, true, q.IsEmpty())
	assert.Panics(t, func() { q.Dequeue() })
	fmt.Println(q.ToString())
	q.Enqueue(88)
	q.Dequeue()
	for i := 3; i < 15; i++ {
		q.Enqueue(i)
	}
	fmt.Println(q.ToString())

	for i := 3; i < 12; i++ {
		q.Dequeue()
	}

	fmt.Println(q.ToString())
}
