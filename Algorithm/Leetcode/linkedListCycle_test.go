package Leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_hasCycle(t *testing.T) {
	assert.Equal(t, false, hasCycle(nil))
	assert.Equal(t, false, hasCycle(&ListNode{
		Val:  0,
		Next: nil,
	}))
	assert.Equal(t, false, hasCycle(&ListNode{
		Val: 0,
		Next: &ListNode{
			Val:  0,
			Next: nil,
		},
	}))
}
