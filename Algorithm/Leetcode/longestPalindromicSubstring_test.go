package Leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	assert.Equal(t, "bab", longestPalindrome("babad"))
	assert.Equal(t, "dddd", longestPalindrome("ddddbabad"))
	assert.Equal(t, "bb", longestPalindrome("cbbd"))
}
