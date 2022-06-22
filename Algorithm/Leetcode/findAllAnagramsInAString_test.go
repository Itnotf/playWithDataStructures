package Leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findAnagrams(t *testing.T) {
	assert.Equal(t, []int{0, 6}, findAnagrams("cbaebabacd", "abc"))
	assert.Equal(t, []int{0, 1, 2}, findAnagrams("abab", "ab"))
}
