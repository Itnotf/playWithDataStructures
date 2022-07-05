package Leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_majorityElement(t *testing.T) {
	assert.Equal(t, 3, majorityElement([]int{3, 2, 3}))
	assert.Equal(t, 2, majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
