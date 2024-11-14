package util

import (
	"errors"
	"fmt"
)

// Base 인터페이스 정의
type Base[T any] interface {
	IsList() bool
	GetFirst() T
	GetLast() T
	GetList() ([]T, error) // Go에는 예외가 없기 때문에 에러를 반환하도록 설계
	Size() int
	Iterator() Iterator[T]
}

// DefaultBase 구조체는 Base 인터페이스의 기본 동작을 제공합니다
type DefaultBase[T any] struct {
	value T
}

// IsList 메서드
func (b *DefaultBase[T]) IsList() bool {
	return false
}

// GetFirst 메서드
func (b *DefaultBase[T]) GetFirst() T {
	return b.value
}

// GetLast 메서드
func (b *DefaultBase[T]) GetLast() T {
	return b.value
}

// GetList 메서드 (기본적으로 지원하지 않음을 나타내는 에러 반환)
func (b *DefaultBase[T]) GetList() ([]T, error) {
	return nil, errors.New("unsupported operation")
}

// Size 메서드
func (b *DefaultBase[T]) Size() int {
	return 1
}

// Iterator 메서드
func (b *DefaultBase[T]) Iterator() Iterator[T] {
	return &BaseIterator[T]{base: b, hasNext: true}
}

// Iterator 인터페이스 정의
type Iterator[T any] interface {
	HasNext() bool
	Next() (T, error)
	Remove() error
}

// BaseIterator 구조체는 Base 인터페이스의 Iterator를 구현
type BaseIterator[T any] struct {
	base    *DefaultBase[T]
	hasNext bool
}

// HasNext 메서드
func (it *BaseIterator[T]) HasNext() bool {
	return it.hasNext
}

// Next 메서드
func (it *BaseIterator[T]) Next() (T, error) {
	if it.hasNext {
		it.hasNext = false
		return it.base.value, nil
	}
	var zero T
	return zero, errors.New("no such element")
}

// Remove 메서드
func (it *BaseIterator[T]) Remove() error {
	return errors.New("unsupported operation")
}

func main() {
	// 기본 Base 객체 생성
	base := &DefaultBase[int]{value: 42}

	// Base 인터페이스의 메서드 사용 예제
	fmt.Println("IsList:", base.IsList())
	fmt.Println("GetFirst:", base.GetFirst())
	fmt.Println("GetLast:", base.GetLast())
	fmt.Println("Size:", base.Size())

	// Iterator 사용 예제
	iterator := base.Iterator()
	for iterator.HasNext() {
		val, err := iterator.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}
		fmt.Println("Iterator value:", val)
	}
}
