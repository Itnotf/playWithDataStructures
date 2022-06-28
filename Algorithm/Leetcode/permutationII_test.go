package Leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_permutation(t *testing.T) {
	assert.Equal(t, []string{"abc", "acb", "bac", "bca", "cab", "cba"}, permutation("abc"))
}
