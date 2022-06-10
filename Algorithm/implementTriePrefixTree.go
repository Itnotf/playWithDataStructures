package algorithm

type TrieNode struct {
	isWord bool
	next   map[byte]*TrieNode
}
type Trie struct {
	root *TrieNode
}

func TrieConstructor() Trie {
	return Trie{
		root: &TrieNode{
			isWord: false,
			next:   map[byte]*TrieNode{},
		},
	}
}

func (t *Trie) Insert(word string) {
	cur := t.root
	for _, char := range word {
		_, ok := cur.next[byte(char)]
		if !ok {
			cur.next[byte(char)] = &TrieNode{
				isWord: false,
				next:   map[byte]*TrieNode{},
			}
		}
		cur = cur.next[byte(char)]
	}
	cur.isWord = true
}

func (t *Trie) Search(word string) bool {
	cur := t.root
	for _, char := range word {
		_, ok := cur.next[byte(char)]
		if !ok {
			return false
		}
		cur = cur.next[byte(char)]
	}
	return cur.isWord
}

func (t *Trie) StartsWith(prefix string) bool {
	cur := t.root
	for _, char := range prefix {
		_, ok := cur.next[byte(char)]
		if !ok {
			return false
		}
		cur = cur.next[byte(char)]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
