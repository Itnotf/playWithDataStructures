package Leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solveNQueens(t *testing.T) {
	assert.Equal(t, 92, len(solveNQueens(8)))
	assert.Equal(t, 1, len(solveNQueens(1)))
	assert.Equal(t, 2, len(solveNQueens(4)))
}
