package Leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_invertTree(t *testing.T) {
	assert.Equal(t, &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
	}, invertTree(&TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}))
}
