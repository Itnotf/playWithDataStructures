package Leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

//deleteNode recursion
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	if head.Val == val {
		head = head.Next
	} else {
		head.Next = deleteNode(head.Next, val)
	}
	return head
}
