package util

import (
	"fmt"
	"testing"
)

func TestDefaultBase(t *testing.T) {
	// 기본 IBase 객체 생성
	base := &DefaultBase[int]{value: 42}

	// IBase 인터페이스의 메서드 사용 예제
	fmt.Println("IsList:", base.IsList())
	fmt.Println("First:", base.First())
	fmt.Println("Last:", base.Last())
	fmt.Println("Size:", base.Size())

	// Iterator 사용 예제
	iterator := base.Iterator()
	for iterator.HasNext() {
		val := iterator.Next()
		fmt.Println("Iterator value:", val)
	}
}
