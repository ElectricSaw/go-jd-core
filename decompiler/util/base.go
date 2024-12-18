package util

// IBase 인터페이스 정의
type IBase[T comparable] interface {
	IIterable[T]

	IsList() bool
	First() T
	Last() T
	ToSlice() []T
	ToList() *DefaultList[T]
	Size() int
}

func NewDefaultBase[T comparable]() IBase[T] {
	return &DefaultBase[T]{
		Iterable: *NewIterable[T]().(*Iterable[T]),
	}
}

// DefaultBase 구조체는 IBase 인터페이스의 기본 동작을 제공합니다
type DefaultBase[T comparable] struct {
	Iterable[T]

	value T
}

// SetValue 메서드
func (b *DefaultBase[T]) SetValue(value T) {
	b.value = value
}

// IsList 메서드
func (b *DefaultBase[T]) IsList() bool {
	return false
}

// First 메서드
func (b *DefaultBase[T]) First() T {
	return b.value
}

// Last 메서드
func (b *DefaultBase[T]) Last() T {
	return b.value
}

// ToSlice 메서드 (기본적으로 지원하지 않음을 나타내는 에러 반환)
func (b *DefaultBase[T]) ToSlice() []T {
	return nil
}

func (b *DefaultBase[T]) ToList() *DefaultList[T] {
	list := make([]T, 0, b.Size())
	list = append(list, b.value)
	return NewDefaultListWithSlice[T](list).(*DefaultList[T])
}

// Size 메서드
func (b *DefaultBase[T]) Size() int {
	return 1
}

// Iterator 메서드
func (b *DefaultBase[T]) Iterator() IIterator[T] {
	list := make([]T, 0, b.Size())
	list = append(list, b.value)
	return NewListIterator(list)
}
