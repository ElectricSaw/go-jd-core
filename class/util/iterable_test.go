package util

import (
	"fmt"
	"testing"
)

func TestIterable(t *testing.T) {
	// 슬라이스를 Iterable로 감싸기
	data := []int{1, 2, 3, 4, 5}
	iterable := &Iterable[int]{data: data}

	// Iterator를 사용해 수동 반복
	iterator := iterable.Iterator()
	for iterator.HasNext() {
		value := iterator.Next()
		fmt.Println(value)
	}

	// ForEach로 각 요소에 대해 작업 수행
	iterable.ForEach(func(value int) {
		fmt.Printf("Value: %d\n", value)
	})
}
