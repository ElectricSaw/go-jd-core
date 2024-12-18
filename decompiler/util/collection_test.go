package util

import (
	"fmt"
	"testing"
)

func TestCollection(t *testing.T) {
	collection := NewCollection[int]()
	collection.Add(1)
	collection.Add(2)
	collection.Add(3)

	fmt.Println("Size:", collection.Size())            // Size: 3
	fmt.Println("Contains 2:", collection.Contains(2)) // Contains 2: true

	collection.Remove(2)
	fmt.Println("Contains 2 after removal:", collection.Contains(2)) // Contains 2 after removal: false

	collection.AddAll([]int{4, 5, 6})
	fmt.Println("All elements:", collection.ToSlice()) // All elements: [1 3 4 5 6]

	collection.RetainAll([]int{1, 4})
	fmt.Println("After retain only [1, 4]:", collection.ToSlice()) // After retain only [1, 4]: [1 4]

	collection.Clear()
	fmt.Println("Size after clear:", collection.Size()) // Size after clear: 0
}
