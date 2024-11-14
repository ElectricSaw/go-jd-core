package util

import (
	"errors"
	"fmt"
)

// DefaultList 구조체 정의
type DefaultList[T any] struct {
	elements []T
}

// 생성자 함수
func NewDefaultList[T any](initialCapacity int) *DefaultList[T] {
	return &DefaultList[T]{elements: make([]T, 0, initialCapacity)}
}

func NewDefaultListWithElements[T any](elements ...T) *DefaultList[T] {
	list := NewDefaultList[T](len(elements))
	list.elements = append(list.elements, elements...)
	return list
}

// 첫 번째 요소 반환
func (d *DefaultList[T]) GetFirst() (T, error) {
	if len(d.elements) == 0 {
		var zero T
		return zero, errors.New("list is empty")
	}
	return d.elements[0], nil
}

// 마지막 요소 반환
func (d *DefaultList[T]) GetLast() (T, error) {
	if len(d.elements) == 0 {
		var zero T
		return zero, errors.New("list is empty")
	}
	return d.elements[len(d.elements)-1], nil
}

// 첫 번째 요소 제거
func (d *DefaultList[T]) RemoveFirst() (T, error) {
	if len(d.elements) == 0 {
		var zero T
		return zero, errors.New("list is empty")
	}
	first := d.elements[0]
	d.elements = d.elements[1:]
	return first, nil
}

// 마지막 요소 제거
func (d *DefaultList[T]) RemoveLast() (T, error) {
	if len(d.elements) == 0 {
		var zero T
		return zero, errors.New("list is empty")
	}
	last := d.elements[len(d.elements)-1]
	d.elements = d.elements[:len(d.elements)-1]
	return last, nil
}

// 요소 추가
func (d *DefaultList[T]) Add(element T) {
	d.elements = append(d.elements, element)
}

// 리스트가 비어있는지 확인
func (d *DefaultList[T]) IsEmpty() bool {
	return len(d.elements) == 0
}

// 빈 리스트를 반환하는 함수
func EmptyList[T any]() *DefaultList[T] {
	return NewDefaultList[T](0)
}

func main() {
	list := NewDefaultListWithElements(1, 2, 3, 4)

	// 첫 번째 요소 확인
	first, _ := list.GetFirst()
	fmt.Println("First element:", first)

	// 마지막 요소 확인
	last, _ := list.GetLast()
	fmt.Println("Last element:", last)

	// 요소 추가
	list.Add(5)
	last, _ = list.GetLast()
	fmt.Println("New last element:", last)

	// 첫 번째 요소 제거
	removedFirst, _ := list.RemoveFirst()
	fmt.Println("Removed first element:", removedFirst)

	// 마지막 요소 제거
	removedLast, _ := list.RemoveLast()
	fmt.Println("Removed last element:", removedLast)
}
