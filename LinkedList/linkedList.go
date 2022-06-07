package linkedList

import (
	"fmt"
	"reflect"
)

type Node[T any] struct {
	val  T
	next *Node[T]
}

type LinkedList[T any] struct {
	dummyHead *Node[T]
	size      int
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		dummyHead: &Node[T]{},
		size:      0,
	}
}

func (l *LinkedList[T]) GetSize() int {
	return l.size
}
func (l LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}
func (l *LinkedList[T]) ToString() string {
	cur := l.dummyHead
	str := ""
	for cur.next != nil {
		cur = cur.next
		str += fmt.Sprintf("%v->", cur.val)
	}
	str += "nil"
	return str
}

func (l *LinkedList[T]) Add(index int, e T) {
	if index < 0 || index > l.size {
		panic("Require index<=0 and index<=linkedList.size")
	}

	pre := l.dummyHead
	for i := 0; i < index; i++ {
		pre = pre.next
	}

	pre.next = &Node[T]{val: e, next: pre.next}

	l.size++
}
func (l *LinkedList[T]) AddFirst(e T) {
	l.Add(0, e)
}
func (l *LinkedList[T]) AddLast(e T) {
	l.Add(l.size, e)
}
func (l *LinkedList[T]) Get(index int) T {
	if index < 0 || index >= l.size {
		panic("Require index<=0 and index<=linkedList.size")
	}
	pre := l.dummyHead
	for i := 0; i < index; i++ {
		pre = pre.next
	}
	return pre.next.val
}

func (l *LinkedList[T]) Remove(index int) T {
	if index < 0 || index >= l.size {
		panic("Require index<=0 and index<=linkedList.size")
	}

	pre := l.dummyHead
	for i := 0; i < index; i++ {
		pre = pre.next
	}
	e := pre.next
	pre.next = pre.next.next
	l.size--
	return e.val
}
func (l *LinkedList[T]) RemoveFirst() T {
	return l.Remove(0)
}
func (l *LinkedList[T]) RemoveLast() T {
	return l.Remove(l.size - 1)
}
func (l *LinkedList[T]) Set(index int, e T) {
	if index < 0 || index >= l.size {
		panic("Require index<=0 and index<=linkedList.size")
	}
	pre := l.dummyHead
	for i := 0; i < index; i++ {
		pre = pre.next
	}
	pre.next.val = e
}
func (l *LinkedList[T]) Contains(e T) bool {
	pre := l.dummyHead
	for pre.next != nil {
		if reflect.DeepEqual(pre.next.val, e) {
			return true
		}
		pre = pre.next
	}
	return false
}
