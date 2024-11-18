package util

import (
	"errors"
	"fmt"
)

// Base 인터페이스 정의
type Base[T any] interface {
	IsList() bool
	First() T
	Last() T
	List() []T
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

// First 메서드
func (b *DefaultBase[T]) First() T {
	return b.value
}

// Last 메서드
func (b *DefaultBase[T]) Last() T {
	return b.value
}

// GetList 메서드 (기본적으로 지원하지 않음을 나타내는 에러 반환)
func (b *DefaultBase[T]) List() []T {
	return nil
}

// Size 메서드
func (b *DefaultBase[T]) Size() int {
	return 1
}

// Iterator 메서드
func (b *DefaultBase[T]) Iterator() Iterator[T] {
	list := make([]T, 0, b.Size())
	list = append(list, b.value)
	return NewListIterator(list)
}

// Iterator 인터페이스 정의
type Iterator[T any] interface {
	HasNext() bool
	Next() (T, error)
	HasPrevious() bool
	Previous() (T, error)
	Add(element T)
	Set(element T) error
	Remove() error
	NextIndex() int
	PreviousIndex() int
	List() []T
}

// ListIterator 구조체 정의
type ListIterator[T any] struct {
	list     []T
	position int
}

// ListIterator 생성자 함수
func NewListIterator[T any](list []T) *ListIterator[T] {
	return &ListIterator[T]{list: list, position: -1}
}

// hasNext 메서드: 다음 요소가 있는지 확인
func (it *ListIterator[T]) HasNext() bool {
	return it.position < len(it.list)-1
}

// next 메서드: 다음 요소 반환 및 커서 이동
func (it *ListIterator[T]) Next() (T, error) {
	if it.HasNext() {
		it.position++
		return it.list[it.position], nil
	}
	var zero T
	return zero, errors.New("no next element")
}

// hasPrevious 메서드: 이전 요소가 있는지 확인
func (it *ListIterator[T]) HasPrevious() bool {
	return it.position > 0
}

// previous 메서드: 이전 요소 반환 및 커서 이동
func (it *ListIterator[T]) Previous() (T, error) {
	if it.HasPrevious() {
		it.position--
		return it.list[it.position], nil
	}
	var zero T
	return zero, errors.New("no previous element")
}

// add 메서드: 현재 커서 위치에 요소 추가
func (it *ListIterator[T]) Add(element T) {
	it.list = append(it.list[:it.position+1], append([]T{element}, it.list[it.position+1:]...)...)
	it.position++
}

// set 메서드: 현재 위치의 요소를 새로운 요소로 수정
func (it *ListIterator[T]) Set(element T) error {
	if it.position < 0 || it.position >= len(it.list) {
		return errors.New("no element to set")
	}
	it.list[it.position] = element
	return nil
}

// remove 메서드: 현재 위치의 요소를 삭제
func (it *ListIterator[T]) Remove() error {
	if it.position < 0 || it.position >= len(it.list) {
		return errors.New("no element to remove")
	}
	it.list = append(it.list[:it.position], it.list[it.position+1:]...)
	it.position--
	return nil
}

// 다음 요소의 인덱스 반환
func (it *ListIterator[T]) NextIndex() int {
	return it.position + 1
}

// 이전 요소의 인덱스 반환
func (it *ListIterator[T]) PreviousIndex() int {
	return it.position - 1
}

// 리스트 반환
func (it *ListIterator[T]) List() []T {
	return it.list
}

func main() {
	// 기본 Base 객체 생성
	base := &DefaultBase[int]{value: 42}

	// Base 인터페이스의 메서드 사용 예제
	fmt.Println("IsList:", base.IsList())
	fmt.Println("First:", base.First())
	fmt.Println("Last:", base.Last())
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

	list := []string{"Apple", "Banana", "Cherry"}
	iterator2 := NewListIterator(list)

	fmt.Println("Forward traversal:")
	for iterator2.HasNext() {
		val, _ := iterator2.Next()
		fmt.Println(val)
		if val == "Banana" {
			iterator2.Set("Blueberry") // Banana를 Blueberry로 수정
		}
	}

	fmt.Println("\nModified list:")
	fmt.Println(iterator2.List())

	fmt.Println("\nBackward traversal:")
	for iterator2.HasPrevious() {
		val, _ := iterator2.Previous()
		fmt.Println(val)
		if val == "Apple" {
			iterator2.Add("Avocado") // Apple 앞에 Avocado 추가
		}
	}

	fmt.Println("\nFinal list:")
	fmt.Println(iterator2.List())
}

//// BaseIterator 구조체는 Base 인터페이스의 Iterator를 구현
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
