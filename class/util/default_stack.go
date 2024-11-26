package util

import (
	"fmt"
)

// DefaultStack 구조체 정의
type DefaultStack[T comparable] struct {
	elements []T
	head     int
}

// NewDefaultStack 기본 생성자 함수
func NewDefaultStack[T comparable]() *DefaultStack[T] {
	return &DefaultStack[T]{
		elements: make([]T, 0, 16),
		head:     0,
	}
}

// NewDefaultStackFrom 다른 스택으로부터 복사하여 생성
func NewDefaultStackFrom[T comparable](other *DefaultStack[T]) *DefaultStack[T] {
	newElements := make([]T, len(other.elements))
	copy(newElements, other.elements)
	return &DefaultStack[T]{
		elements: newElements,
		head:     other.head,
	}
}

// Clear 스택 초기화
func (s *DefaultStack[T]) Clear() {
	s.head = 0
}

// Size 스택 크기 반환
func (s *DefaultStack[T]) Size() int {
	return s.head
}

// IsEmpty 스택이 비어있는지 확인
func (s *DefaultStack[T]) IsEmpty() bool {
	return s.head == 0
}

// Copy 스택을 다른 스택으로부터 복사
func (s *DefaultStack[T]) Copy(other *DefaultStack[T]) {
	if len(s.elements) < other.head {
		s.elements = make([]T, other.head)
	}
	copy(s.elements, other.elements[:other.head])
	s.head = other.head
}

// Push 스택에 요소 추가
func (s *DefaultStack[T]) Push(expression T) {
	if s.head == len(s.elements) {
		newElements := make([]T, len(s.elements)*2)
		copy(newElements, s.elements)
		s.elements = newElements
	}
	s.elements[s.head] = expression
	s.head++
}

// Pop 스택에서 요소 제거 및 반환
func (s *DefaultStack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("stack is empty")
	}
	s.head--
	element := s.elements[s.head]
	// 메모리 관리 (nil로 설정)
	var zero T
	s.elements[s.head] = zero
	return element, nil
}

// Peek 스택의 마지막 요소 확인
func (s *DefaultStack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("stack is empty")
	}
	return s.elements[s.head-1], nil
}

// Replace 스택 내 요소 치환
func (s *DefaultStack[T]) Replace(old, new T) {
	for i := s.head - 1; i >= 0; i-- {
		if s.elements[i] == old {
			s.elements[i] = new
		}
	}
}

// String 문자열 표현 반환
func (s *DefaultStack[T]) String() string {
	result := "Stack{head=" + fmt.Sprintf("%d", s.head) + ", elements=["
	for i := 0; i < s.head; i++ {
		if i > 0 {
			result += ", "
		}
		result += fmt.Sprintf("%v", s.elements[i])
	}
	result += "]}"
	return result
}

func main() {
	stack := NewDefaultStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println("Stack after pushes:", stack)

	stack.Pop()
	fmt.Println("Stack after pop:", stack)

	stack.Replace(2, 5)
	fmt.Println("Stack after replace:", stack)

	value, _ := stack.Peek()
	fmt.Println("Peek element:", value)
}
