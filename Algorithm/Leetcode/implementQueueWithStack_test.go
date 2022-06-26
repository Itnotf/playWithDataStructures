package Leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCQueue_AppendTail(t *testing.T) {
	obj := QueueConstructor()
	obj.AppendTail(1)
	obj.AppendTail(2)
	obj.AppendTail(3)
	assert.Equal(t, 1, obj.DeleteHead())
	assert.Equal(t, 2, obj.DeleteHead())
	assert.Equal(t, 3, obj.DeleteHead())
	assert.Equal(t, -1, obj.DeleteHead())
}
