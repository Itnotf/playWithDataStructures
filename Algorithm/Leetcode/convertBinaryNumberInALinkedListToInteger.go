package Leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {

	sum := 0

	for head != nil {
		sum = 2*sum + head.Val
		head = head.Next
	}

	return sum
}
