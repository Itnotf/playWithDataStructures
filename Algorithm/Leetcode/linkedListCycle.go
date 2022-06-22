package Leetcode

func hasCycle(head *ListNode) bool {
	first := head
	second := head

	for first != nil && first.Next != nil {
		first = first.Next.Next
		second = second.Next
		if first == second {
			return true
		}
	}

	return false
}
