package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := TrieConstructor()
	trie.Insert("apple")
	assert.Equal(t, true, trie.Search("apple"))
	assert.Equal(t, false, trie.Search("app"))
	assert.Equal(t, true, trie.StartsWith("app"))
	trie.Insert("app")
	assert.Equal(t, true, trie.Search("app"))
}
