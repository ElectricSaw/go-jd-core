package util

import (
	"fmt"
	"testing"
)

func TestDefaultList(t *testing.T) {
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
