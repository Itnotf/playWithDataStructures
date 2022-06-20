package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_twoSum(t *testing.T) {
	assert.Equal(t, []int{0, 1}, twoSum([]int{2, 7, 11, 15}, 9))
	assert.Equal(t, []int{0, 1}, twoSum([]int{3, 3}, 6))
	assert.Equal(t, []int{1, 2}, twoSum([]int{3, 2, 4}, 6))
}
