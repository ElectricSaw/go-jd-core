package util

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	iterator := NewIteratorWithSlice(data)

	// 반복문으로 요소 출력
	for iterator.HasNext() {
		element := iterator.Next()
		fmt.Println("Next element:", element)

		// 특정 조건에서 요소 제거
		if element == 3 {
			err := iterator.Remove()
			if err != nil {
				fmt.Println("Remove error:", err)
			}
		}
	}

	// ForEachRemaining로 남은 요소 처리
	fmt.Println("Remaining elements:")
	iterator = NewIteratorWithSlice(data) // 새로운 Iterator 생성
	iterator.ForEachRemaining(func(value int) {
		fmt.Println("Processing:", value)
	})
}
