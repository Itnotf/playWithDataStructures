package Leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getDecimalValue(t *testing.T) {
	assert.Equal(t, 5, getDecimalValue(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}))
}
