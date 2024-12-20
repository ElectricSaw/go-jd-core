package util

import (
	"errors"
)

// ISet 인터페이스: Java Set 인터페이스에 해당
type ISet[T comparable] interface {
	Size() int               // Set의 크기 반환
	IsEmpty() bool           // Set이 비어 있는지 확인
	Contains(element T) bool // 요소 포함 여부 확인
	Get(index int) T
	Add(element T) bool            // 요소 추가
	AddAll(element []T) bool       // 모든 요소 추가
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
	validate map[T]bool
	list     []T
}

// NewSet AbstractSet 생성자
func NewSet[T comparable]() ISet[T] {
	return &Set[T]{
		validate: make(map[T]bool),
		list:     make([]T, 0),
	}
}

func NewSetWithSlice[T comparable](data []T) ISet[T] {
	s := &Set[T]{
		validate: make(map[T]bool),
		list:     make([]T, 0),
	}
	for _, v := range data {
		s.Add(v)
	}
	return s
}

// Size Set의 크기 반환
func (s *Set[T]) Size() int {
	return len(s.list)
}

// IsEmpty Set이 비어 있는지 확인
func (s *Set[T]) IsEmpty() bool {
	return len(s.list) == 0
}

// Contains 특정 요소 포함 여부 확인
func (s *Set[T]) Contains(element T) bool {
	_, exists := s.validate[element]
	return exists
}

func (s *Set[T]) Get(index int) T {
	if index > s.Size() {
		var zero T
		return zero
	}
	return s.list[index]
}

// AddAll 모든 요소 추가
func (s *Set[T]) AddAll(element []T) bool {
	ret := true
	for _, v := range element {
		ret = ret && s.Add(v)
	}
	return ret
}

// Add Set에 요소 추가
func (s *Set[T]) Add(element T) bool {
	if _, exists := s.validate[element]; exists {
		return false
	}

	s.validate[element] = true
	s.list = append(s.list, element)

	return true
}

// Remove Set에서 요소 제거
func (s *Set[T]) Remove(element T) bool {
	if _, exists := s.validate[element]; exists {
		delete(s.validate, element)
		return true
	}
	return false
}

// Clear 모든 요소 제거
func (s *Set[T]) Clear() {
	s.validate = make(map[T]bool)
	s.list = make([]T, 0)
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
	slice := make([]T, 0, len(s.validate))
	for element := range s.validate {
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
	for element := range s.validate {
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
