package LinkedListMap

type Node[K comparable, V comparable] struct {
	key   K
	value V
	next  *Node[K, V]
}

type LinkedListMap[K comparable, V comparable] struct {
	dummyHead *Node[K, V]
	size      int
}

func NewLinkedListMap[K comparable, V comparable]() *LinkedListMap[K, V] {
	return &LinkedListMap[K, V]{
		dummyHead: &Node[K, V]{
			next: nil,
		},
		size: 0,
	}
}

func (l *LinkedListMap[K, V]) GetSize() int {
	return l.size
}
func (l *LinkedListMap[K, V]) IsEmpty() bool {
	return l.size == 0
}
func (l *LinkedListMap[K, V]) Contains(value V) bool {
	pre := l.dummyHead
	for pre.next != nil {
		if pre.next.value == value {
			return true
		}
		pre = pre.next
	}
	return false
}
func (l *LinkedListMap[K, V]) Get(key K) V {
	pre := l.dummyHead
	for pre.next != nil {
		if pre.next.key == key {
			return pre.next.value
		}
		pre = pre.next
	}
	panic("no such key")
}
func (l *LinkedListMap[K, V]) Set(key K, value V) {
	pre := l.dummyHead
	for pre.next != nil {
		if pre.next.key == key {
			pre.next.value = value
			return
		}
		pre = pre.next
	}
	node := &Node[K, V]{key: key, value: value}
	node.next = l.dummyHead.next
	l.dummyHead.next = node
	l.size++
}
func (l *LinkedListMap[K, V]) Remove(key K) V {
	pre := l.dummyHead
	for pre.next != nil {
		if pre.next.key == key {
			r := pre.next
			pre.next = pre.next.next
			l.size--
			return r.value
		}
		pre = pre.next
	}
	panic("no such key")
}
