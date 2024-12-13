package util

// IIterable 인터페이스: Java의 Iterable<T>에 해당
type IIterable[T comparable] interface {
	Iterator() IIterator[T] // Iterator 반환
	ForEach(action func(T)) // 각 요소에 대해 액션 수행
}

// Iterable 슬라이스를 Iterable로 구현
type Iterable[T comparable] struct {
	data []T
}

// NewIterable 슬라이스를 기반으로 한 Iterator 생성자
func NewIterable[T comparable]() IIterable[T] {
	return NewIterableWithSlice(make([]T, 0))
}

// NewIterableWithSlice 슬라이스를 기반으로 한 Iterator 생성자
func NewIterableWithSlice[T comparable](data []T) IIterable[T] {
	return &Iterable[T]{data: data}
}

// Iterator Iterable에 대한 Iterator 반환
func (s *Iterable[T]) Iterator() IIterator[T] {
	return &Iterator[T]{data: s.data, index: 0}
}

// ForEach 각 요소에 대해 주어진 액션 수행
func (s *Iterable[T]) ForEach(action func(T)) {
	for _, v := range s.data {
		action(v)
	}
}
