package Leetcode

type LRUCache struct {
	size                 int
	cap                  int
	cache                map[int]*DoubleLinkedList
	dummyHead, dummyTail *DoubleLinkedList
}

type DoubleLinkedList struct {
	key, value int
	prev, next *DoubleLinkedList
}

func NewLRUCache(capacity int) LRUCache {
	l := &LRUCache{
		size:      0,
		cap:       capacity,
		cache:     map[int]*DoubleLinkedList{},
		dummyHead: &DoubleLinkedList{},
		dummyTail: &DoubleLinkedList{},
	}
	l.dummyHead.next = l.dummyTail
	l.dummyTail.prev = l.dummyHead

	return *l
}

func (l *LRUCache) Get(key int) int {
	//找到对应节点，返回
	if node, ok := l.cache[key]; ok {
		//由于访问过，将节点放到链表头部，删除并插入头部

		l.deleteNode(node)
		l.addFront(node)

		return node.value
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	//如果key存在更新value，并将节点放入链表头部
	if node, ok := l.cache[key]; ok {
		node.value = value
		l.deleteNode(node)
		l.addFront(node)
		return
	}
	//如果key不存在
	//如果容量已满，删除链表尾部节点，将节点放入链表头部
	node := &DoubleLinkedList{key: key, value: value}

	if l.size == l.cap {
		tail := l.deleteTail()
		delete(l.cache, tail.key)
		l.size--
	}

	l.addFront(node)
	l.cache[key] = node
	l.size++
}

func (l *LRUCache) addFront(node *DoubleLinkedList) {
	node.next = l.dummyHead.next
	l.dummyHead.next = node
	node.prev = node.next.prev
	node.next.prev = node
}

func (l *LRUCache) deleteNode(node *DoubleLinkedList) *DoubleLinkedList {
	node.prev.next = node.next
	node.next.prev = node.prev

	return node
}

func (l *LRUCache) deleteTail() *DoubleLinkedList {
	node := l.dummyTail.prev
	node.prev.next = l.dummyTail
	l.dummyTail.prev = node.prev

	return node
}
