package Trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Add("book")
	trie.Add("bond")
	assert.Equal(t, true, trie.Contains("book"))
	assert.Equal(t, false, trie.Contains("bon"))
	assert.Equal(t, false, trie.Contains("a"))
	assert.Equal(t, false, trie.IsPrefix("pip"))
	assert.Equal(t, true, trie.IsPrefix("b"))
	assert.Equal(t, true, trie.IsPrefix("boo"))
}
