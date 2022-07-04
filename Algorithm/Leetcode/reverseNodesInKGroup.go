package Leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 传入头节点，k，返回翻转后的头节点
func reverseKGroup(head *ListNode, k int) *ListNode {
	start := head
	end := head
	for i := 0; i < k; i++ {
		if end == nil {
			return head
		}
		end = end.Next
	}

	head, tail := reverse(start, end)
	tail.Next = reverseKGroup(end, k)
	return head
}

//翻转从start到end的链表，并返回
func reverse(start, end *ListNode) (*ListNode, *ListNode) {
	head := start

	var pre *ListNode
	cur := start
	for cur != end {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre, head
}
