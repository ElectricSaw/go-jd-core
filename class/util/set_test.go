package util

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	// Set 생성
	set := NewSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	// 요소 확인
	fmt.Println("Size:", set.Size())            // Size: 3
	fmt.Println("Contains 2:", set.Contains(2)) // Contains 2: true

	// 요소 제거
	set.Remove(2)
	fmt.Println("Contains 2 after removal:", set.Contains(2)) // Contains 2 after removal: false

	// Iterator 사용
	it := set.Iterator()
	for it.HasNext() {
		value := it.Next()
		fmt.Println("Iterating:", value)
	}
}
