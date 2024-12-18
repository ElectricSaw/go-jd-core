package util

import (
	"fmt"
	"testing"
)

func TestArrayList(t *testing.T) {
	list := NewArrayList[int]()
	_ = list.Add(10)
	_ = list.Add(20)
	_ = list.Add(30)

	err := list.AddAt(1, 15)
	if err != nil {
		fmt.Println("Error:", err)
	}

	element := list.RemoveAt(2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Removed element:", element)
	}

	value := list.Get(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Element at index 1:", value)
	}

	_ = list.Set(0, 5)

	fmt.Println("Contains 20:", list.Contains(20))
	list.Clear()
}
