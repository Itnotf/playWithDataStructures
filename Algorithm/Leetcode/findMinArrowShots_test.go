package Leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMinArrowShots(t *testing.T) {
	assert.Equal(t, 2, findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}))
	assert.Equal(t, 4, findMinArrowShots([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}))
	assert.Equal(t, 2, findMinArrowShots([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}))
}
