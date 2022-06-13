package BSTSet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBSTSet(t *testing.T) {
	set := NewBSTSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(2)
	set.Add(5)
	assert.Equal(t, 3, set.GetSize())
	assert.Equal(t, true, set.Contains(2))
	assert.Equal(t, false, set.Contains(3))
	set.Remove(2)
	assert.Panics(t, func() { set.Remove(3) })
	set.Remove(1)
	set.Remove(5)
	assert.Panics(t, func() { set.Remove(1) })
	assert.Equal(t, true, set.IsEmpty())
}
