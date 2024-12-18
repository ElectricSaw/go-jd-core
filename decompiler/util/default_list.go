package util

var EmptyList = NewDefaultList[interface{}]()

// DefaultList 구조체 정의
type DefaultList[T comparable] struct {
	ArrayList[T]
}

func NewDefaultList[T comparable]() IList[T] {
	return &DefaultList[T]{
		ArrayList: *NewArrayList[T]().(*ArrayList[T]),
	}
}

func NewDefaultListWithCapacity[T comparable](capacity int) IList[T] {
	return &DefaultList[T]{
		ArrayList: *NewArrayListWithCapacity[T](capacity).(*ArrayList[T]),
	}
}

func NewDefaultListWithSlice[T comparable](elements []T) IList[T] {
	return NewDefaultListWithElements[T](elements...)
}

func NewDefaultListWithElements[T comparable](elements ...T) IList[T] {
	list := NewDefaultListWithCapacity[T](len(elements))
	_ = list.AddAll(elements)
	return list
}
