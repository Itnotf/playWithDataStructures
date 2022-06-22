package Leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_deleteNode(t *testing.T) {
	node := &ListNode{
		Val: 4,
		Next: &ListNode{Val: 5,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}
	assert.Equal(t, &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  9,
				Next: nil,
			},
		},
	}, deleteNode(node, 5))
	assert.Equal(t, (*ListNode)(nil), deleteNode(nil, 1))
}
