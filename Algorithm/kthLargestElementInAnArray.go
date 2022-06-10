package algorithm

type MaxHeap struct {
	data []int
	size int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		data: []int{},
		size: 0,
	}
}
func (m *MaxHeap) GetSize() int {
	return m.size
}
func (m *MaxHeap) parent(index int) int {
	return (index - 1) / 2
}
func (m *MaxHeap) leftChild(index int) int {
	return 2*index + 1
}
func (m *MaxHeap) rightChild(index int) int {
	return 2*index + 2
}
func (m *MaxHeap) FindMax() int {
	return m.data[0]
}
func (m *MaxHeap) Add(e int) {
	m.data = append(m.data, e)
	m.size++
	m.siftUp()
}
func (m *MaxHeap) siftUp() {
	cur := m.size - 1
	for cur > 0 {
		parent := m.parent(cur)
		if m.data[cur] > m.data[m.parent(cur)] {
			m.data[cur], m.data[parent] = m.data[parent], m.data[cur]
			cur = parent
		} else {
			break
		}
	}
}
func (m *MaxHeap) ExtractMax() int {
	max := m.data[0]
	m.data[0] = m.data[m.size-1]
	m.data[m.size-1] = 0
	m.size--
	m.siftDown(0)
	return max
}
func (m *MaxHeap) siftDown(index int) {
	cur := index
	for m.leftChild(cur) < m.size {
		left := m.leftChild(cur)
		right := m.rightChild(cur)
		if right < m.size && m.data[right] > m.data[left] && m.data[right] > m.data[cur] {
			m.data[right], m.data[cur] = m.data[cur], m.data[right]
			cur = right
			continue
		}
		//left
		if m.data[left] > m.data[cur] {
			m.data[left], m.data[cur] = m.data[cur], m.data[left]
			cur = left
		} else {
			break
		}
	}
}

func (m *MaxHeap) Replace(e int) {
	m.data[0] = e
	m.siftDown(0)
}

func findKthLargest(nums []int, k int) int {
	maxHeap := NewMaxHeap()
	for _, num := range nums {
		maxHeap.Add(num)
	}
	for i := 0; i < k-1; i++ {
		maxHeap.ExtractMax()
	}
	return maxHeap.FindMax()
}
