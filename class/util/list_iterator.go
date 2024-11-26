package util

import (
	"errors"
	"fmt"
)

// IListIterator Java의 ListIterator에 해당
type IListIterator[T comparable] interface {
	IIterator[T]

	HasPrevious() bool   // 이전 요소가 있는지 확인
	Previous() T         // 이전 요소 반환
	NextIndex() int      // 다음 요소의 인덱스 반환
	PreviousIndex() int  // 이전 요소의 인덱스 반환
	Remove() error       // 마지막 반환된 요소 제거
	Set(element T) error // 마지막 반환된 요소를 대체
	Add(element T) error // 현재 위치에 요소 추가
	ToSlice() []T
}

// ListIterator 슬라이스 기반 ListIterator 구현
type ListIterator[T comparable] struct {
	data      []T
	cursor    int
	lastIndex int // 마지막 반환된 요소의 인덱스
}

// NewListIterator 슬라이스 기반 ListIterator 생성자
func NewListIterator[T comparable](data []T) IListIterator[T] {
	return &ListIterator[T]{data: data, cursor: 0, lastIndex: -1}
}

// HasNext 다음 요소가 있는지 확인
func (it *ListIterator[T]) HasNext() bool {
	return it.cursor < len(it.data)
}

// Next 다음 요소 반환
func (it *ListIterator[T]) Next() T {
	if !it.HasNext() {
		var zero T
		return zero
	}
	element := it.data[it.cursor]
	it.lastIndex = it.cursor
	it.cursor++
	return element
}

// HasPrevious 이전 요소가 있는지 확인
func (it *ListIterator[T]) HasPrevious() bool {
	return it.cursor > 0
}

// Previous 이전 요소 반환
func (it *ListIterator[T]) Previous() T {
	if !it.HasPrevious() {
		var zero T
		return zero
	}
	it.cursor--
	it.lastIndex = it.cursor
	return it.data[it.cursor]
}

// NextIndex 다음 요소의 인덱스 반환
func (it *ListIterator[T]) NextIndex() int {
	return it.cursor
}

// PreviousIndex 이전 요소의 인덱스 반환
func (it *ListIterator[T]) PreviousIndex() int {
	return it.cursor - 1
}

// Remove 마지막 반환된 요소 제거
func (it *ListIterator[T]) Remove() error {
	if it.lastIndex == -1 {
		return errors.New("illegal state: no element to remove")
	}
	it.data = append(it.data[:it.lastIndex], it.data[it.lastIndex+1:]...)
	if it.lastIndex < it.cursor {
		it.cursor--
	}
	it.lastIndex = -1
	return nil
}

// Set 마지막 반환된 요소 대체
func (it *ListIterator[T]) Set(element T) error {
	if it.lastIndex == -1 {
		return errors.New("illegal state: no element to set")
	}
	it.data[it.lastIndex] = element
	return nil
}

// Add 현재 위치에 요소 추가
func (it *ListIterator[T]) Add(element T) error {
	it.data = append(it.data[:it.cursor], append([]T{element}, it.data[it.cursor:]...)...)
	it.cursor++
	it.lastIndex = -1
	return nil
}

func (it *ListIterator[T]) ToSlice() []T {
	return it.data
}

// ForEachRemaining 남은 요소를 모두 처리
func (it *ListIterator[T]) ForEachRemaining(action func(T)) {
	for it.HasNext() {
		element := it.Next()
		action(element)
	}
}

// 사용 예제
func main() {
	data := []int{10, 20, 30, 40}
	iterator := NewListIterator(data)

	fmt.Println("Forward iteration:")
	for iterator.HasNext() {
		element := iterator.Next()
		fmt.Println("Next element:", element)

		if element == 20 {
			iterator.Set(25) // Replace 20 with 25
		}
	}

	fmt.Println("\nBackward iteration:")
	for iterator.HasPrevious() {
		element := iterator.Previous()
		fmt.Println("Previous element:", element)

		if element == 25 {
			_ = iterator.Remove() // Remove 25
		}
	}

	iterator.Add(15) // Add 15 at the current position
	fmt.Println("\nFinal data:", iterator.ToSlice())
}
