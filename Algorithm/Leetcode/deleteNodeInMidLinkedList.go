package Leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMidNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
