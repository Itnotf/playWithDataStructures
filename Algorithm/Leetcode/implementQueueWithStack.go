package Leetcode

type Stack struct {
	data []int
}

func (s *Stack) Push(value int) {
	s.data = append(s.data, value)
}

func (s *Stack) Pop() int {
	length := len(s.data)
	head := s.data[length-1]
	s.data = s.data[:length-1]
	return head
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

type CQueue struct {
	stackA *Stack
	stackB *Stack
}

func QueueConstructor() CQueue {
	s1 := &Stack{data: []int{}}
	s2 := &Stack{data: []int{}}
	return CQueue{
		stackA: s1,
		stackB: s2,
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stackA.Push(value)
}

func (this *CQueue) DeleteHead() int {
	if !this.stackB.IsEmpty() {
		return this.stackB.Pop()
	} else {
		for !this.stackA.IsEmpty() {
			this.stackB.Push(this.stackA.Pop())
		}
		if this.stackB.IsEmpty() {
			return -1
		}
		return this.stackB.Pop()
	}
}
