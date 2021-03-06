package Leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_coinChange(t *testing.T) {
	assert.Equal(t, 3, coinChange([]int{1, 2, 5}, 11))
	assert.Equal(t, -1, coinChange([]int{2}, 3))
	assert.Equal(t, 0, coinChange([]int{1}, 0))
}
