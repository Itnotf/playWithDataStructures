package Leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	//遍历一遍获取长度
	//前一半入栈，后一半跟栈顶元素比较
	length := 0

	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}

	var stack []int
	cur = head
	i := 0
	for cur != nil {
		if i == length/2 && length%2 == 1 {
			cur = cur.Next
			i++
			continue
		}

		if i < length/2 {
			//入栈
			stack = append(stack, cur.Val)
		} else {
			//比较，不同返回false，相同则出栈
			if len(stack) == 0 || stack[len(stack)-1] != cur.Val {
				return false
			}
			stack = stack[:len(stack)-1]
		}

		cur = cur.Next
		i++
	}

	if len(stack) == 0 {
		return true
	}
	return false
}
