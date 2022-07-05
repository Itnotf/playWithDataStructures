package Leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_runningSum(t *testing.T) {
	assert.Equal(t, []int{1, 3, 6, 10}, runningSum([]int{1, 2, 3, 4}))
}
