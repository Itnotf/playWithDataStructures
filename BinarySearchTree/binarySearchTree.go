package BinarySearchTree

import "fmt"

type Node[T int | float32 | ~string] struct {
	val   T
	left  *Node[T]
	right *Node[T]
}

type BST[T int | float32 | ~string] struct {
	root *Node[T]
	size int
}

func NewBST[T int | float32 | ~string]() *BST[T] {
	return &BST[T]{}
}

func (b *BST[T]) GetSize() int {
	return b.size
}

func (b *BST[T]) IsEmpty() bool {
	return b.size == 0
}

func (b *BST[T]) Add(e T) {
	b.add(b.root, e)
	b.size++
}
func (b *BST[T]) add(root *Node[T], e T) {
	node := &Node[T]{val: e}
	if root == nil {
		b.root = node
	} else if root.val < e && root.right == nil {
		root.right = node
	} else if root.val > e && root.left == nil {
		root.left = node
	} else if root.val < e {
		b.add(root.right, e)
	} else if root.val > e {
		b.add(root.left, e)
	}
}

func (b *BST[T]) Contains(e T) bool {
	return b.contains(b.root, e)
}
func (b *BST[T]) contains(root *Node[T], e T) bool {
	if root == nil {
		return false
	}
	if root.val > e {
		return b.contains(root.left, e)
	}
	if root.val < e {
		return b.contains(root.right, e)
	}
	return true
}

func (b *BST[T]) PreOrder() {
	fmt.Println("Pre order Begin:")
	b.preOrder(b.root)
	fmt.Println("\nEnd.")
}

func (b *BST[T]) preOrder(root *Node[T]) {
	if root == nil {
		return
	}
	fmt.Printf("%v ", root.val)
	b.preOrder(root.left)
	b.preOrder(root.right)
}

func (b *BST[T]) InOrder() {
	fmt.Println("In order Begin:")
	b.inOrder(b.root)
	fmt.Println("\nEnd.")
}

func (b *BST[T]) inOrder(root *Node[T]) {
	if root == nil {
		return
	}
	b.inOrder(root.left)
	fmt.Printf("%v ", root.val)
	b.inOrder(root.right)
}

func (b *BST[T]) PostOrder() {
	fmt.Println("Post order Begin:")
	b.postOrder(b.root)
	fmt.Println("\nEnd.")
}

func (b *BST[T]) postOrder(root *Node[T]) {
	if root == nil {
		return
	}
	b.postOrder(root.left)
	b.postOrder(root.right)
	fmt.Printf("%v ", root.val)
}

func (b *BST[T]) Minimum() T {
	return b.minimum(b.root).val
}
func (b *BST[T]) minimum(root *Node[T]) *Node[T] {
	if root == nil {
		return nil
	}
	if root.left == nil {
		return root
	}
	return b.minimum(root.left)
}

func (b *BST[T]) Maximum() T {
	return b.maximum(b.root).val
}
func (b *BST[T]) maximum(root *Node[T]) *Node[T] {
	if root == nil {
		return nil
	}
	if root.right == nil {
		return root
	}
	return b.maximum(root.right)
}
func (b *BST[T]) RemoveMin() T {
	if b.size == 0 {
		panic("empty BST")
	}
	ret := b.Minimum()
	b.root = b.removeMin(b.root)
	return ret
}
func (b *BST[T]) removeMin(root *Node[T]) *Node[T] {
	if root.left == nil {
		b.size--
		return root.right
	}
	root.left = b.removeMin(root.left)
	return root
}
func (b *BST[T]) RemoveMax() T {
	if b.size == 0 {
		panic("empty BST")
	}
	ret := b.Maximum()
	b.root = b.removeMax(b.root)
	return ret
}
func (b *BST[T]) removeMax(root *Node[T]) *Node[T] {
	if root.right == nil {
		b.size--
		return root.left
	}
	root.right = b.removeMax(root.right)
	return root
}

func (b *BST[T]) Remove(e T) {
	if b.size == 0 {
		panic("empty BST")
	}
	b.root = b.remove(b.root, e)
}
func (b *BST[T]) remove(root *Node[T], e T) *Node[T] {
	if root.val > e {
		root.left = b.remove(root.left, e)
		return root
	}
	if root.val < e {
		root.right = b.remove(root.right, e)
		return root
	}

	if root.left == nil {
		node := root.right
		root.right = nil
		b.size--
		return node
	}
	if root.right == nil {
		node := root.left
		root.left = nil
		b.size--
		return node
	}
	successor := b.minimum(root.right)
	successor.right = b.removeMin(root.right)
	successor.left = root.left
	root.left = nil
	root.right = nil
	return successor
}

func (b *BST[T]) Rank(e T) int {
	return b.rank(b.root, e)
}
func (b *BST[T]) rank(root *Node[T], e T) int {
	if root.val > e {
		return b.rank(root.left, e)
	}
	if root.val < e {
		return b.Count(root.left) + 1 + b.rank(root.right, e)
	}
	return b.Count(root.left) + 1
}

func (b *BST[T]) Count(root *Node[T]) int {
	if root == nil {
		return 0
	}
	return b.Count(root.left) + b.Count(root.right) + 1
}

func (b BST[T]) Select(rank int) T {
	if rank < 1 || rank > b.size {
		panic("rank must in the range of  [1,size]")
	}
	return b.selectRank(b.root, rank)
}
func (b *BST[T]) selectRank(root *Node[T], rank int) T {
	if b.Count(root.left)+1 == rank {
		return root.val
	} else if b.Count(root.left)+1 > rank {
		return b.selectRank(root.left, rank)
	} else {
		return b.selectRank(root.right, rank-b.Count(root.left)-1)
	}
}

func (b BST[T]) Floor(e T) T {
	return b.floor(b.root, e).val
}

func (b BST[T]) floor(root *Node[T], e T) *Node[T] {
	if root == nil {
		return nil
	}
	if root.val == e {
		return root
	} else if root.val > e {
		return b.floor(root.left, e)
	} else {
		if right := b.floor(root.right, e); right != nil {
			return right
		}
		return root
	}
}

func (b BST[T]) Ceil(e T) T {
	return b.ceil(b.root, e).val
}

func (b *BST[T]) ceil(root *Node[T], e T) *Node[T] {
	if root == nil {
		return nil
	}
	if root.val == e {
		return root
	} else if root.val < e {
		return b.ceil(root.right, e)
	} else {
		if left := b.ceil(root.left, e); left != nil {
			return left
		}
		return root
	}
}
