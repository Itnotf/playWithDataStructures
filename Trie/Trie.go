package Trie

type Node struct {
	isWord bool
	next   map[byte]*Node
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{
		root: &Node{
			isWord: false,
			next:   map[byte]*Node{},
		},
	}
}

func (t *Trie) Add(word string) {
	cur := t.root
	for _, char := range word {
		_, ok := cur.next[byte(char)]
		if !ok {
			cur.next[byte(char)] = &Node{
				isWord: false,
				next:   map[byte]*Node{},
			}
		}
		cur = cur.next[byte(char)]
	}
	cur.isWord = true
}

func (t *Trie) Contains(word string) bool {
	cur := t.root
	for _, char := range word {
		node, ok := cur.next[byte(char)]
		if !ok {
			return false
		}
		cur = node
	}
	return cur.isWord
}

func (t *Trie) IsPrefix(prefix string) bool {
	cur := t.root
	for _, char := range prefix {
		node, ok := cur.next[byte(char)]
		if !ok {
			return false
		}
		cur = node
	}
	return true
}
