package Map

type Map[T any] interface {
	GetSize() int
	IsEmpty() bool
	Add(T)
	Remove(T) T
	Contains(T)
	Get(key int)
	Set(key int)
}
