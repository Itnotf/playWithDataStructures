package UnionFind

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnionFind(t *testing.T) {
	unionFind := NewUnionFind()
	unionFind.UnionElement(2, 3)
	unionFind.UnionElement(3, 4)
	unionFind.UnionElement(5, 4)
	unionFind.UnionElement(7, 2)
	assert.Equal(t, true, unionFind.IsConnected(2, 5))
	assert.Equal(t, true, unionFind.IsConnected(4, 7))
	assert.Equal(t, false, unionFind.IsConnected(2, 6))
	assert.Panics(t, func() { unionFind.UnionElement(10, 20) })
	assert.Panics(t, func() { unionFind.IsConnected(10, 20) })
}
