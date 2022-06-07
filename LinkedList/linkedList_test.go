package linkedList

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList(t *testing.T) {
	l := NewLinkedList[int]()
	l.AddLast(33)
	l.AddLast(22)
	l.AddLast(11)
	l.AddFirst(9)
	l.AddFirst(8)
	l.AddFirst(7)
	fmt.Println(l.ToString())
	l.Add(4, 8)
	l.Set(1, 2)
	assert.Equal(t, 33, l.Get(3))
	assert.Equal(t, false, l.IsEmpty())
	assert.Equal(t, 7, l.GetSize())
	assert.Panics(t, func() { l.Add(999, 9) })
	assert.Panics(t, func() { l.Get(999) })
	assert.Panics(t, func() { l.Set(999, 9) })
	assert.Equal(t, true, l.Contains(11))
	assert.Equal(t, false, l.Contains(112))
	fmt.Println(l.ToString())
	assert.Equal(t, 7, l.RemoveFirst())
	assert.Equal(t, 11, l.RemoveLast())
	assert.Equal(t, 9, l.Remove(1))
	assert.Panics(t, func() { l.Remove(100) })

}
