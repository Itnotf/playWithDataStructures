package ArrayStack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewArrayStack(t *testing.T) {
	s := NewArrayStack[int](10)
	s.Push(10)
	assert.Equal(t, false, s.IsEmpty())
	assert.Equal(t, 10, s.Pop())
	assert.Equal(t, true, s.IsEmpty())
	assert.Panics(t, func() { s.Pop() })
	for i := 1; i < 10; i++ {
		s.Push(i)
	}
	assert.Equal(t, 9, s.GetSize())
	assert.Equal(t, 10, s.GetCapacity())
	assert.Equal(t, 9, s.Peek())
	fmt.Println(s.ToString())
}
