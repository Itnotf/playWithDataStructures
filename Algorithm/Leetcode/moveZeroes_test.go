package Leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	a := []int{0, 1, 0, 3, 12}
	moveZeroes(a)
	assert.Equal(t, []int{1, 3, 12, 0, 0}, a)
}
