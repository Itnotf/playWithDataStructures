package LinkedListMap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedListMap(t *testing.T) {
	m := NewLinkedListMap[int, string]()
	assert.Equal(t, 0, m.GetSize())
	assert.Equal(t, true, m.IsEmpty())
	assert.Equal(t, false, m.Contains("a"))
	assert.Panics(t, func() { m.Get(1) })
	assert.Panics(t, func() { m.Remove(1) })
	m.Set(2, "a")
	m.Set(2, "B")
	m.Set(8, "9")
	m.Set(7, "7")
	m.Set(6, "6")
	assert.Equal(t, "B", m.Get(2))
	assert.Equal(t, true, m.Contains("9"))
	assert.Equal(t, "7", m.Remove(7))
	m.Set(20, "20")
	m.Set(15, "15")
	m.Set(22, "22")
	assert.Equal(t, 6, m.size)
	assert.Equal(t, "9", m.Remove(8))
	assert.Equal(t, "B", m.Remove(2))
	assert.Equal(t, "6", m.Remove(6))
}
