package Leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast := head

	for i := 0; i < k; i++ {
		if fast == nil {
			return nil
		}

		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		head = head.Next
	}
	return head
}
