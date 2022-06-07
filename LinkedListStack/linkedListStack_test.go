package linkedListStack

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedListStack(t *testing.T) {
	s := NewLinkedListStack[int]()
	s.Push(10)
	assert.Equal(t, s.IsEmpty(), false)
	assert.Equal(t, s.Pop(), 10)
	assert.Equal(t, s.IsEmpty(), true)
	assert.Panics(t, func() { s.Pop() })
	for i := 1; i < 10; i++ {
		s.Push(i)
	}
	assert.Equal(t, s.GetSize(), 9)
	assert.Equal(t, s.Peek(), 9)
	fmt.Println(s.ToString())
}
