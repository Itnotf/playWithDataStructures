package algorithm

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			for head != fast {
				head = head.Next
				fast = fast.Next
			}
			return head
		}
	}
	return nil
}
