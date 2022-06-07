package loopQueue

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoopQueue(t *testing.T) {
	q := NewLoopQueue[int](10)
	q.Enqueue(99)
	q.Enqueue(88)
	assert.Equal(t, q.GetFront(), 99)
	assert.Equal(t, q.GetSize(), 2)
	assert.Equal(t, q.IsEmpty(), false)
	fmt.Println(q.ToString())

	assert.Equal(t, q.Dequeue(), 99)
	assert.Equal(t, q.Dequeue(), 88)
	assert.Equal(t, q.IsEmpty(), true)
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
