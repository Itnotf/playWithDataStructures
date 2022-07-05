package Leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	//简单用map判断
	mapA := map[*ListNode]int{}

	for headA != nil {
		mapA[headA] = 1
		headA = headA.Next
	}
	for headB != nil {
		if _, ok := mapA[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

//TODO:双指针解法
