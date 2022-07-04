package Leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	node := reverseKGroup(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: &ListNode{},
					},
				},
			},
		},
	}, 2)
	assert.Equal(t, 2, node.Val)
	assert.Equal(t, 1, node.Next.Val)
	assert.Equal(t, 4, node.Next.Next.Val)
	assert.Equal(t, 3, node.Next.Next.Next.Val)
	assert.Equal(t, 5, node.Next.Next.Next.Next.Val)
}
