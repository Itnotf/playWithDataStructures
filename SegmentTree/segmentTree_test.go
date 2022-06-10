package SegmentTree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSegmentTree(t *testing.T) {
	data := []int{-2, 0, 3, -5, 2, -1}
	segmentTree := NewSegmentTree(data)

	assert.Equal(t, 1, segmentTree.Query(0, 2))
	assert.Equal(t, -1, segmentTree.Query(2, 5))
	assert.Equal(t, -3, segmentTree.Query(0, 5))
	segmentTree.Set(1, 2)
	assert.Equal(t, -1, segmentTree.Query(0, 5))
}
