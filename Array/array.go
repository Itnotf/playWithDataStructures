package array

import (
	"fmt"
	"reflect"
)

type Array[T any] struct {
	data []T
	size int
}

func NewArray[T any](capacity int) *Array[T] {
	return &Array[T]{
		data: make([]T, capacity),
		size: 0,
	}
}

func (arr *Array[T]) ToString() string {
	return fmt.Sprintf("Size:%d,Capacity:%d\n%v", arr.GetSize(), arr.GetCapacity(), arr.data)
}

func (arr *Array[T]) GetSize() int {
	return arr.size
}

func (arr *Array[T]) GetCapacity() int {
	return len(arr.data)
}

func (arr *Array[T]) IsEmpty() bool {
	return arr.size == 0
}

func (arr *Array[T]) AddLast(e T) {
	arr.Add(arr.size, e)
}

func (arr *Array[T]) AddFirst(e T) {
	arr.Add(0, e)
}

func (arr *Array[T]) Add(index int, e T) {
	if arr.size == arr.GetCapacity() {
		arr.resize(2 * arr.GetCapacity())
	}

	if index < 0 || index > arr.size {
		panic("add Failed. Require index<=0 and index<=arr.size")
	}

	for i := arr.size - 1; i >= index; i-- {
		arr.data[i+1] = arr.data[i]
	}
	arr.data[index] = e
	arr.size++
}

func (arr *Array[T]) Get(index int) T {
	if index < 0 || index > arr.size {
		panic("get Failed. Require index<=0 and index<=arr.size")
	}
	return arr.data[index]
}

func (arr *Array[T]) Set(index int, e T) {
	if index < 0 || index > arr.size {
		panic("set Failed. Require index<=0 and index<=arr.size")
	}
	arr.data[index] = e
}

func (arr *Array[T]) Contains(e T) bool {
	for i := 0; i < arr.size; i++ {
		if reflect.DeepEqual(arr.data[i], e) {
			return true
		}
	}
	return false
}

func (arr *Array[T]) Find(e T) int {
	for i := 0; i < arr.size; i++ {
		if reflect.DeepEqual(arr.data[i], e) {
			return i
		}
	}
	return -1
}

func (arr *Array[T]) Remove(index int) T {
	if index < 0 || index > arr.size {
		panic("get Failed. Require index<=0 and index<=arr.size")
	}
	t := arr.data[index]
	for i := index; i < arr.size-1; i++ {
		arr.data[i] = arr.data[i+1]
	}
	arr.size--

	//lazy
	if arr.size < arr.GetCapacity()/4 && arr.GetCapacity()/2 != 0 {
		arr.resize(arr.GetCapacity() / 2)
	}
	//Important
	arr.data[arr.size] = *new(T)
	return t
}

func (arr *Array[T]) RemoveFirst() T {
	return arr.Remove(0)
}

func (arr *Array[T]) RemoveLast() T {
	return arr.Remove(arr.size - 1)
}

func (arr *Array[T]) RemoveElement(e T) T {
	index := arr.Find(e)
	if index != -1 {
		return arr.Remove(index)
	}
	panic("remove failed:no such element")
}

func (arr *Array[T]) resize(newCapacity int) {
	newData := make([]T, newCapacity)
	for i := 0; i < arr.size; i++ {
		newData[i] = arr.data[i]
	}
	arr.data = newData
}
