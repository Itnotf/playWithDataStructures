package Leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{Val: 0}
	head := l3
	tmp := 0

	for l1 != nil || l2 != nil || tmp != 0 {
		sum := 0
		if l1 == nil && l2 == nil {
			sum += tmp
		} else if l1 == nil {
			sum += tmp + l2.Val
			l2 = l2.Next
		} else if l2 == nil {
			sum += tmp + l1.Val
			l1 = l1.Next
		} else {
			sum += tmp + l1.Val + l2.Val
			l1 = l1.Next
			l2 = l2.Next
		}

		tmp = sum / 10

		l3.Next = &ListNode{Val: sum % 10}
		l3 = l3.Next
	}

	return head.Next
}
