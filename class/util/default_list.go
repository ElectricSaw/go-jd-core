package util

import (
	"fmt"
)

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

func main() {
	list := NewDefaultListWithElements(1, 2, 3, 4)

	// 첫 번째 요소 확인
	first := list.First()
	fmt.Println("First element:", first)

	// 마지막 요소 확인
	last := list.Last()
	fmt.Println("Last element:", last)

	// 요소 추가
	_ = list.Add(5)
	last = list.Last()
	fmt.Println("New last element:", last)

	// 첫 번째 요소 제거
	removedFirst := list.RemoveFirst()
	fmt.Println("Removed first element:", removedFirst)

	// 마지막 요소 제거
	removedLast := list.RemoveLast()
	fmt.Println("Removed last element:", removedLast)
}
