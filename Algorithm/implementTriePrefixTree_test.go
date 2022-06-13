package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
