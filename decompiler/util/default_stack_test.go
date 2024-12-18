package util

import (
	"fmt"
	"testing"
)

func TestDefaultStack(t *testing.T) {
	stack := NewDefaultStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println("Stack after pushes:", stack)

	stack.Pop()
	fmt.Println("Stack after pop:", stack)

	stack.Replace(2, 5)
	fmt.Println("Stack after replace:", stack)

	value := stack.Peek()
	fmt.Println("Peek element:", value)
}
