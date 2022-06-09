package algorithm

type MinHeap struct {
	data [][2]int
	size int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		data: [][2]int{},
		size: 0,
	}
}
func (m *MinHeap) GetSize() int {
	return m.size
}
func (m *MinHeap) parent(index int) int {
	return (index - 1) / 2
}
func (m *MinHeap) leftChild(index int) int {
	return 2*index + 1
}
func (m *MinHeap) rightChild(index int) int {
	return 2*index + 2
}
func (m *MinHeap) FindMin() [2]int {
	return m.data[0]
}
func (m *MinHeap) Add(e [2]int) {
	m.data = append(m.data, e)
	m.size++
	m.siftUp()
}
func (m *MinHeap) siftUp() {
	cur := m.size - 1
	for cur > 0 {
		parent := m.parent(cur)
		if m.data[cur][1] < m.data[m.parent(cur)][1] {
			m.data[cur], m.data[parent] = m.data[parent], m.data[cur]
			cur = parent
		} else {
			break
		}
	}
}
func (m *MinHeap) ExtractMin() [2]int {
	min := m.data[0]
	m.data[0] = m.data[m.size-1]
	m.size--
	m.siftDown(0)
	return min
}
func (m *MinHeap) siftDown(index int) {
	cur := index
	for m.leftChild(cur) < m.size {
		left := m.leftChild(cur)
		right := m.rightChild(cur)
		if right < m.size && m.data[right][1] < m.data[left][1] && m.data[right][1] < m.data[cur][1] {
			m.data[right], m.data[cur] = m.data[cur], m.data[right]
			cur = right
			continue
		}
		//left
		if m.data[left][1] < m.data[cur][1] {
			m.data[left], m.data[cur] = m.data[cur], m.data[left]
			cur = left
		} else {
			break
		}
	}
}
func (m *MinHeap) Replace(e [2]int) {
	m.data[0] = e
	m.siftDown(0)
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num] += 1
	}
	minHeap := NewMinHeap()
	for num, freq := range m {
		if minHeap.GetSize() < k {
			minHeap.Add([2]int{num, freq})
		} else {
			if minHeap.FindMin()[1] < freq {
				minHeap.Replace([2]int{num, freq})
			}
		}
	}
	var ret []int
	for minHeap.GetSize() > 0 {
		ret = append(ret, minHeap.ExtractMin()[0])
	}
	return ret
}
