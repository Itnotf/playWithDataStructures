package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	assert.Equal(t, (*ListNode)(nil), detectCycle(nil))
	assert.Equal(t, (*ListNode)(nil), detectCycle(&ListNode{
		Val:  0,
		Next: nil,
	}))
}
