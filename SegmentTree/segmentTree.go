package SegmentTree

type SegmentTree struct {
	data []int
	tree []int
}

func NewSegmentTree(data []int) *SegmentTree {
	s := &SegmentTree{
		data: nil,
		tree: nil,
	}
	s.data = data
	s.tree = make([]int, 4*len(data))
	s.buildSegmentTree(0, 0, len(data)-1)
	return s
}

func (s *SegmentTree) buildSegmentTree(index int, left int, right int) {
	if left == right {
		s.tree[index] = s.data[left]
		return
	}
	leftChild := s.leftChild(index)
	rightChild := s.rightChild(index)

	mid := left + (right-left)/2
	s.buildSegmentTree(leftChild, left, mid)
	s.buildSegmentTree(rightChild, mid+1, right)

	s.tree[index] = s.tree[leftChild] + s.tree[rightChild]
}

func (s *SegmentTree) leftChild(i int) int {
	return 2*i + 1
}
func (s *SegmentTree) rightChild(i int) int {
	return 2*i + 2
}

func (s *SegmentTree) Query(left int, right int) int {
	return s.query(0, 0, len(s.data)-1, left, right)
}

func (s *SegmentTree) query(index int, l int, r int, left int, right int) int {
	if l == left && r == right {
		return s.tree[index]
	}
	mid := l + (r-l)/2
	if right <= mid {
		return s.query(s.leftChild(index), l, mid, left, right)
	}
	if left >= mid+1 {
		return s.query(s.rightChild(index), mid+1, r, left, right)
	}
	return s.query(s.leftChild(index), l, mid, left, mid) + s.query(s.rightChild(index), mid+1, r, mid+1, right)
}

func (s *SegmentTree) Set(key int, value int) {
	s.data[key] = value
	s.set(0, 0, len(s.data)-1, key)
}

func (s *SegmentTree) set(index int, l int, r int, key int) {
	if l == r {
		s.tree[index] = s.data[l]
		return
	}
	mid := l + (r-l)/2
	leftChild := s.leftChild(index)
	rightChild := s.rightChild(index)

	if key <= mid {
		s.set(leftChild, l, mid, key)
	} else {
		s.set(rightChild, mid+1, r, key)
	}

	s.tree[index] = s.tree[leftChild] + s.tree[rightChild]
}
