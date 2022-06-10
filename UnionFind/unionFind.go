package UnionFind

type UnionFind struct {
	data []int
}

func NewUnionFind() *UnionFind {
	var data []int
	for i := 0; i < 10; i++ {
		data = append(data, i)
	}
	return &UnionFind{data}
}

func (u *UnionFind) UnionElement(p int, q int) {
	if p < 0 || p > len(u.data) || q < 0 || q > len(u.data) {
		panic("out of range")
	}

	u.data[p] = u.find(q)
}

func (u *UnionFind) find(p int) int {
	for u.data[p] != p {
		p = u.data[p]
	}

	return p
}

func (u *UnionFind) IsConnected(p int, q int) bool {
	if p < 0 || p > len(u.data) || q < 0 || q > len(u.data) {
		panic("out of range")
	}
	pRoot := u.find(p)
	qRoot := u.find(q)
	return pRoot == qRoot
}
