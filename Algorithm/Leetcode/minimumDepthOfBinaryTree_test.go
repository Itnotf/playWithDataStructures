package Leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minDepth(t *testing.T) {
	assert.Equal(t, 0, minDepth(nil))
	assert.Equal(t, 1, minDepth(&TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}))
	assert.Equal(t, 2, minDepth(&TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		},
	}))
	assert.Equal(t, 5, minDepth(&TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val:   0,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
				Right: nil,
			},
		},
	}))
}
