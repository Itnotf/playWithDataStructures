package Set

type set[T any] interface {
	GetSize() int
	IsEmpty() bool
	Add(T)
	Remove(T)
	Contains(T) bool
}
