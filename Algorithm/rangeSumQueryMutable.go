package algorithm

type NumArray struct {
	data []int
	tree []int
}

func Constructor(nums []int) NumArray {
	numArray := &NumArray{}
	numArray.data = nums
	numArray.tree = make([]int, 4*len(nums))
	numArray.buildSegmentTree(0, 0, len(nums)-1)
	return *numArray
}

func (n *NumArray) buildSegmentTree(index, l int, r int) {
	if l == r {
		n.tree[index] = n.data[l]
		return
	}
	mid := l + (r-l)/2
	left := n.leftChild(index)
	right := n.rightChild(index)

	n.buildSegmentTree(left, l, mid)
	n.buildSegmentTree(right, mid+1, r)

	n.tree[index] = n.tree[left] + n.tree[right]
}

func (n *NumArray) leftChild(i int) int {
	return 2*i + 1
}

func (n *NumArray) rightChild(i int) int {
	return 2*i + 2
}

func (n *NumArray) SumRange(left int, right int) int {
	return n.sumRange(0, 0, len(n.data)-1, left, right)
}
func (n *NumArray) sumRange(index int, l int, r int, left int, right int) int {
	if l == left && r == right {
		return n.tree[index]
	}
	mid := l + (r-l)/2
	if right <= mid {
		return n.sumRange(n.leftChild(index), l, mid, left, right)
	}
	if left >= mid+1 {
		return n.sumRange(n.rightChild(index), mid+1, r, left, right)
	}
	return n.sumRange(n.leftChild(index), l, mid, left, mid) + n.sumRange(n.rightChild(index), mid+1, r, mid+1, right)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

func (n *NumArray) Update(index int, val int) {
	n.data[index] = val
	n.update(0, 0, len(n.data)-1, index)
}

func (n *NumArray) update(index int, l int, r int, key int) {
	if l == r {
		n.tree[index] = n.data[l]
		return
	}
	mid := l + (r-l)/2
	leftChild := n.leftChild(index)
	rightChild := n.rightChild(index)

	if key <= mid {
		n.update(leftChild, l, mid, key)
	} else {
		n.update(rightChild, mid+1, r, key)
	}

	n.tree[index] = n.tree[leftChild] + n.tree[rightChild]
}
