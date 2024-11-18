package util

import (
	"fmt"
)

type IList[T any] interface {
	Elements() []T
	Get(index int) T
	First() T
	Last() T
	RemoveFirst() T
	RemoveLast() T
	Add(element T)
	AddAll(elements []T)
	IsEmpty() bool
	Size() int
}

// DefaultList 구조체 정의
type DefaultList[T any] struct {
	DefaultBase[T]
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

func (d *DefaultList[T]) Elements() []T {
	return d.elements
}

func (d *DefaultList[T]) Get(index int) T {
	return d.elements[index]
}

// 첫 번째 요소 반환
func (d *DefaultList[T]) First() T {
	if len(d.elements) == 0 {
		var zero T
		return zero
	}
	return d.elements[0]
}

// 마지막 요소 반환
func (d *DefaultList[T]) Last() T {
	if len(d.elements) == 0 {
		var zero T
		return zero
	}
	return d.elements[len(d.elements)-1]
}

// 첫 번째 요소 제거
func (d *DefaultList[T]) RemoveFirst() T {
	if len(d.elements) == 0 {
		var zero T
		return zero
	}
	first := d.elements[0]
	d.elements = d.elements[1:]
	return first
}

// 마지막 요소 제거
func (d *DefaultList[T]) RemoveLast() T {
	if len(d.elements) == 0 {
		var zero T
		return zero
	}
	last := d.elements[len(d.elements)-1]
	d.elements = d.elements[:len(d.elements)-1]
	return last
}

// 요소 추가
func (d *DefaultList[T]) Add(element T) {
	d.elements = append(d.elements, element)
}

func (d *DefaultList[T]) AddAll(elements []T) {
	d.elements = append(d.elements, elements...)
}

// 리스트가 비어있는지 확인
func (d *DefaultList[T]) IsEmpty() bool {
	return len(d.elements) == 0
}

func (d *DefaultList[T]) Size() int {
	return len(d.elements)
}

// 빈 리스트를 반환하는 함수
func EmptyList[T any]() *DefaultList[T] {
	return NewDefaultList[T](0)
}

func main() {
	list := NewDefaultListWithElements(1, 2, 3, 4)

	// 첫 번째 요소 확인
	first := list.First()
	fmt.Println("First element:", first)

	// 마지막 요소 확인
	last := list.Last()
	fmt.Println("Last element:", last)

	// 요소 추가
	list.Add(5)
	last = list.Last()
	fmt.Println("New last element:", last)

	// 첫 번째 요소 제거
	removedFirst := list.RemoveFirst()
	fmt.Println("Removed first element:", removedFirst)

	// 마지막 요소 제거
	removedLast := list.RemoveLast()
	fmt.Println("Removed last element:", removedLast)
}
