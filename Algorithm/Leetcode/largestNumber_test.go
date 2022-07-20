package Leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_largestNumber(t *testing.T) {
	assert.Equal(t, "210", largestNumber([]int{10, 2}))
	assert.Equal(t, "9534330", largestNumber([]int{3, 30, 34, 5, 9}))
}
