package BSTMap

type Node[K int | float32, V comparable] struct {
	key   K
	value V
	left  *Node[K, V]
	right *Node[K, V]
}

type BSTMap[K int | float32, V comparable] struct {
	root *Node[K, V]
	size int
}

func NewBSTMap[K int | float32, V comparable]() *BSTMap[K, V] {
	return &BSTMap[K, V]{
		root: nil,
		size: 0,
	}
}

func (b *BSTMap[K, V]) GetSize() int {
	return b.size
}

func (b *BSTMap[K, V]) IsEmpty() bool {
	return b.size == 0
}
func (b *BSTMap[K, V]) Contains(value V) bool {
	return b.contains(b.root, value)
}
func (b *BSTMap[K, V]) contains(root *Node[K, V], value V) bool {
	if root == nil {
		return false
	}
	if root.value == value {
		return true
	}
	return b.contains(root.left, value) || b.contains(root.right, value)
}
func (b *BSTMap[K, V]) Get(key K) V {
	ret := b.get(b.root, key)
	return ret
}

func (b *BSTMap[K, V]) get(root *Node[K, V], key K) V {
	if root == nil {
		panic("no such key")
	}
	if root.key < key {
		return b.get(root.right, key)
	}
	if root.key > key {
		return b.get(root.left, key)
	}
	return root.value
}
func (b *BSTMap[K, V]) Set(key K, value V) {
	b.set(b.root, key, value)
}
func (b *BSTMap[K, V]) set(root *Node[K, V], key K, value V) {
	node := &Node[K, V]{key: key, value: value}
	if root == nil {
		b.root = node
		b.size++
	} else if root.key > key && root.left == nil {
		root.left = node
		b.size++
	} else if root.key < key && root.right == nil {
		root.right = node
		b.size++
	} else if root.key > key {
		b.set(root.left, key, value)
	} else if root.key < key {
		b.set(root.right, key, value)
	} else {
		root.value = value
	}
}
func (b *BSTMap[K, V]) Remove(key K) V {
	if b.size == 0 {
		panic("emptyMap")
	}
	node := b.Get(key)
	b.root = b.remove(b.root, key)
	return node
}
func (b *BSTMap[K, V]) remove(root *Node[K, V], key K) *Node[K, V] {
	if root.key > key {
		root.left = b.remove(root.left, key)
		return root
	}
	if root.key < key {
		root.right = b.remove(root.right, key)
		return root
	}
	//root.key=key
	if root.left == nil {
		b.size--
		return root.right
	}
	if root.right == nil {
		b.size--
		return root.left
	}
	//root.left != nil && root.right != nil
	successor := b.minimum(root.right)
	successor.right = b.removeMin(root.right)
	successor.left = root.left
	root.left = nil
	root.right = nil
	return successor
}

func (b *BSTMap[K, V]) minimum(root *Node[K, V]) *Node[K, V] {
	if root.left == nil {
		return root
	}
	return b.minimum(root.left)
}
func (b *BSTMap[K, V]) removeMin(root *Node[K, V]) *Node[K, V] {
	if root.left == nil {
		b.size--
		return root.right
	}
	return b.removeMin(root.left)
}
