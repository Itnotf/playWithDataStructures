package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var q []*TreeNode
	step := 1
	q = append(q, root)

	for len(q) != 0 {
		for _, node := range q {
			q = q[1:]

			if node.Left == nil && node.Right == nil {
				return step
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		step++
	}

	return step
}
