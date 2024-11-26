package util

import (
	"fmt"
)

// IBase 인터페이스 정의
type IBase[T comparable] interface {
	IIterable[T]

	IsList() bool
	First() T
	Last() T
	ToSlice() []T
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

func main() {
	// 기본 IBase 객체 생성
	base := &DefaultBase[int]{value: 42}

	// IBase 인터페이스의 메서드 사용 예제
	fmt.Println("IsList:", base.IsList())
	fmt.Println("First:", base.First())
	fmt.Println("Last:", base.Last())
	fmt.Println("Size:", base.Size())

	// Iterator 사용 예제
	iterator := base.Iterator()
	for iterator.HasNext() {
		val := iterator.Next()
		fmt.Println("Iterator value:", val)
	}
}

//// BaseIterator 구조체는 IBase 인터페이스의 Iterator를 구현
//type BaseIterator[T any] struct {
//	base    *DefaultBase[T]
//	hasNext bool
//}
//
//// HasNext 메서드
//func (it *BaseIterator[T]) HasNext() bool {
//	return it.hasNext
//}
//
//// Next 메서드
//func (it *BaseIterator[T]) Next() (T, error) {
//	if it.hasNext {
//		it.hasNext = false
//		return it.base.value, nil
//	}
//	var zero T
//	return zero, errors.New("no such element")
//}
//
//// Remove 메서드
//func (it *BaseIterator[T]) Remove() error {
//	return errors.New("unsupported operation")
//}
