package Leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_search(t *testing.T) {
	assert.Equal(t, 4, binarySearch([]int{-1, 0, 3, 5, 9, 12}, 9))
	assert.Equal(t, -1, binarySearch([]int{-1, 0, 3, 5, 9, 12}, 2))
}
