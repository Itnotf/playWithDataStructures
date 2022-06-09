package Map

type Map[K any, V any] interface {
	GetSize() int
	IsEmpty() bool
	Get(K) V
	Set(K, V)
	Contains(V)
	Remove(K)
}
