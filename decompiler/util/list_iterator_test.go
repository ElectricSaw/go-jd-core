package util

import (
	"fmt"
	"testing"
)

func TestListIterator(t *testing.T) {
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
