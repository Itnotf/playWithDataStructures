package algorithm

type Node struct {
	value int
	next  map[int32]*Node
}

type MapSum struct {
	root *Node
}

func MapSumConstructor() *MapSum {
	return &MapSum{root: &Node{
		value: 0,
		next:  map[int32]*Node{},
	}}
}

func (t *MapSum) Insert(key string, val int) {
	cur := t.root
	for _, char := range key {
		_, ok := cur.next[char]
		if !ok {
			cur.next[char] = &Node{
				value: 0,
				next:  map[int32]*Node{},
			}
		}
		cur = cur.next[char]
	}
	cur.value = val
}

func (t *MapSum) Sum(prefix string) int {
	cur := t.root
	for _, char := range prefix {
		_, ok := cur.next[char]
		if !ok {
			return 0
		}
		cur = cur.next[char]
	}
	return t.sum(cur)
}

func (t *MapSum) sum(root *Node) int {
	sum := root.value
	for _, node := range root.next {
		sum += t.sum(node)
	}

	return sum
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
