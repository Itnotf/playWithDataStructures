package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArray(t *testing.T) {
	arr := NewArray[[]int](4)
	assert.Panics(t, func() { arr.Add(1, []int{1}) })

	arr.Add(0, []int{1})
	arr.Add(0, []int{2})
	arr.Add(1, []int{3})
	fmt.Println(arr.ToString())
	arr.AddFirst([]int{999})
	arr.AddFirst([]int{88})
	fmt.Println(arr.ToString())
	arr.AddLast([]int{9})
	arr.AddLast([]int{8})
	fmt.Println(arr.ToString())

	assert.Equal(t, true, arr.Contains([]int{9}))
	assert.Equal(t, false, arr.Contains([]int{6}))
	assert.Equal(t, 0, arr.Find([]int{88}))
	assert.Equal(t, -1, arr.Find([]int{99}))
	assert.Equal(t, arr.Get(1), []int{999})
	assert.Panics(t, func() { arr.Get(999) })
	assert.Equal(t, 8, arr.GetCapacity())
	assert.Equal(t, 7, arr.GetSize())
	assert.Equal(t, arr.Remove(0), []int{88})
	assert.Panics(t, func() { arr.Remove(999) })
	assert.Equal(t, arr.RemoveLast(), []int{8})
	assert.Equal(t, arr.RemoveFirst(), []int{999})
	assert.Equal(t, arr.RemoveElement([]int{2}), []int{2})
	assert.Panics(t, func() { arr.RemoveElement([]int{900}) })

	fmt.Println(arr.ToString())
	arr.Set(0, []int{9897})
	assert.Panics(t, func() { arr.Set(999, []int{999}) })
	assert.Equal(t, arr.Get(0), []int{9897})
	assert.Equal(t, false, arr.IsEmpty())
	arr.RemoveLast()
	arr.RemoveLast()
	arr.RemoveLast()
}
