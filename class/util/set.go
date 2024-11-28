package util

import (
	"errors"
	"fmt"
)

// ISet 인터페이스: Java Set 인터페이스에 해당
type ISet[T comparable] interface {
	Size() int               // Set의 크기 반환
	IsEmpty() bool           // Set이 비어 있는지 확인
	Contains(element T) bool // 요소 포함 여부 확인
	Get(index int) T
	Add(element T) bool            // 요소 추가
	Remove(element T) bool         // 요소 제거
	Clear()                        // 모든 요소 제거
	ContainsAll(elements []T) bool // 여러 요소 포함 여부 확인
	RemoveAll(elements []T) bool   // 여러 요소 제거
	ToSlice() []T                  // Set을 슬라이스로 변환
	RetainAll(elements []T) bool   //
	Iterator() IIterator[T]        // Iterator 반환
}

// Set 인터페이스의 기본 구현체
type Set[T comparable] struct {
	data map[T]Store[T]
}

type Store[T comparable] struct {
	index int
	data  T
}

// NewSet AbstractSet 생성자
func NewSet[T comparable]() ISet[T] {
	return &Set[T]{data: make(map[T]Store[T])}
}

func NewSetWithSlice[T comparable](data []T) ISet[T] {
	s := &Set[T]{data: make(map[T]Store[T])}
	for _, v := range data {
		s.Add(v)
	}
	return s
}

// Size Set의 크기 반환
func (s *Set[T]) Size() int {
	return len(s.data)
}

// IsEmpty Set이 비어 있는지 확인
func (s *Set[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Contains 특정 요소 포함 여부 확인
func (s *Set[T]) Contains(element T) bool {
	_, exists := s.data[element]
	return exists
}

func (s *Set[T]) Get(index int) T {
	for k, v := range s.data {
		if v.index == index {
			return k
		}
	}
	var zero T
	return zero
}

// Add Set에 요소 추가
func (s *Set[T]) Add(element T) bool {
	if _, exists := s.data[element]; exists {
		return false
	}
	s.data[element] = Store[T]{index: len(s.data), data: element}
	return true
}

// Remove Set에서 요소 제거
func (s *Set[T]) Remove(element T) bool {
	if _, exists := s.data[element]; exists {
		delete(s.data, element)
		return true
	}
	return false
}

// Clear 모든 요소 제거
func (s *Set[T]) Clear() {
	s.data = make(map[T]Store[T])
}

// ContainsAll 여러 요소 포함 여부 확인
func (s *Set[T]) ContainsAll(elements []T) bool {
	for _, element := range elements {
		if !s.Contains(element) {
			return false
		}
	}
	return true
}

// RemoveAll 여러 요소 제거
func (s *Set[T]) RemoveAll(elements []T) bool {
	modified := false
	for _, element := range elements {
		if s.Remove(element) {
			modified = true
		}
	}
	return modified
}

// ToSlice Set을 슬라이스로 변환
func (s *Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(s.data))
	for element := range s.data {
		slice = append(slice, element)
	}
	return slice
}

// Iterator Set의 Iterator 반환
func (s *Set[T]) Iterator() IIterator[T] {
	elements := s.ToSlice()
	return &SetIterator[T]{data: elements, index: 0, set: s}
}

// RetainAll 특정 요소만 유지
func (s *Set[T]) RetainAll(elements []T) bool {
	changed := false
	elementsSet := make(map[T]struct{})
	for _, element := range elements {
		elementsSet[element] = struct{}{}
	}
	for element := range s.data {
		if _, exists := elementsSet[element]; !exists {
			s.Remove(element)
			changed = true
		}
	}
	return changed
}

// SetIterator Set을 순회하는 Iterator 구현
type SetIterator[T comparable] struct {
	data  []T
	index int
	set   ISet[T]
}

// HasNext 다음 요소가 있는지 확인
func (it *SetIterator[T]) HasNext() bool {
	return it.index < len(it.data)
}

// Next 다음 요소 반환
func (it *SetIterator[T]) Next() T {
	if !it.HasNext() {
		var zero T
		return zero
	}
	element := it.data[it.index]
	it.index++
	return element
}

// Remove 현재 요소 제거
func (it *SetIterator[T]) Remove() error {
	if it.index == 0 || it.index > len(it.data) {
		return errors.New("invalid state for remove")
	}
	element := it.data[it.index-1]
	if !it.set.Remove(element) {
		return errors.New("failed to remove element")
	}
	return nil
}

// ForEachRemaining 남은 요소를 모두 처리
func (it *SetIterator[T]) ForEachRemaining(action func(T)) {
	for it.HasNext() {
		element := it.Next()
		action(element)
	}
}

// 사용 예제
func main() {
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
