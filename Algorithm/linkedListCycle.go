package algorithm

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}

	first := head
	second := head

	for first.Next != nil && second.Next != nil && second.Next.Next != nil {
		first = first.Next
		second = second.Next.Next
		if first == second {
			return true
		}
	}

	return false
}
