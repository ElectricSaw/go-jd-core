package util

import (
	"errors"
	"fmt"
	"sort"
)

// NewArrayList 생성자
func NewArrayList[T comparable]() IList[T] {
	return NewArrayListWithCapacity[T](0)
}

// NewArrayListWithCapacity 생성자
func NewArrayListWithCapacity[T comparable](capacity int) IList[T] {
	return &ArrayList[T]{
		data: make([]T, 0, capacity),
		size: 0,
	}
}

// NewArrayListWithData 생성자
func NewArrayListWithData[T comparable](data []T) IList[T] {
	l := NewArrayList[T]()
	_ = l.AddAll(data)
	return l
}

// IList 인터페이스 정의
type IList[T comparable] interface {
	ICollection[T]

	Get(index int) T
	Set(index int, element T) T
	AddAt(index int, element T) error
	RemoveAt(index int) T
	IndexOf(element T) int
	ToList() IList[T]
	ListIterator() IListIterator[T]

	// First 첫 번째 요소 반환
	First() T
	Last() T
	RemoveFirst() T
	RemoveLast() T

	// SubList 마지막 요소 제거
	SubList(start, end int) IList[T]
	Sort(action func(i, j int) bool)
}

// ArrayList 구조체 정의
type ArrayList[T comparable] struct {
	data []T
	size int
}

// Size 메서드
func (list *ArrayList[T]) Size() int {
	return list.size
}

// IsEmpty 메서드
func (list *ArrayList[T]) IsEmpty() bool {
	return list.size == 0
}

// Contains 메서드
func (list *ArrayList[T]) Contains(element T) bool {
	for _, e := range list.data {
		if e == element {
			return true
		}
	}
	return false
}

// Add 메서드
func (list *ArrayList[T]) Add(element T) bool {
	list.data = append(list.data, element)
	list.size++
	return true
}

func (list *ArrayList[T]) AddAll(elements []T) bool {
	for _, element := range elements {
		_ = list.Add(element)
	}
	return true
}

// Remove 메서드
func (list *ArrayList[T]) Remove(element T) bool {
	for i, e := range list.data {
		if e == element {
			list.data = append(list.data[:i], list.data[i+1:]...)
			list.size--
			return true
		}
	}
	return false
}

// Get 메서드
func (list *ArrayList[T]) Get(index int) T {
	if index < 0 || index >= list.size {
		var zero T
		return zero
	}
	return list.data[index]
}

// Set 메서드
func (list *ArrayList[T]) Set(index int, element T) T {
	if index < 0 || index >= list.size {
		var zero T
		return zero
	}
	oldValue := list.data[index]
	list.data[index] = element
	return oldValue
}

// AddAt 메서드
func (list *ArrayList[T]) AddAt(index int, element T) error {
	if index < 0 || index > list.size {
		return errors.New("index out of bounds")
	}
	list.data = append(list.data[:index], append([]T{element}, list.data[index:]...)...)
	list.size++
	return nil
}

// RemoveAt 메서드
func (list *ArrayList[T]) RemoveAt(index int) T {
	if index < 0 || index >= list.size {
		var zero T
		return zero
	}
	element := list.data[index]
	list.data = append(list.data[:index], list.data[index+1:]...)
	list.size--
	return element
}

// IndexOf 메서드
func (list *ArrayList[T]) IndexOf(element T) int {
	for i, e := range list.data {
		if e == element {
			return i
		}
	}
	return -1
}

// Clear 메서드
func (list *ArrayList[T]) Clear() {
	list.data = make([]T, 0)
	list.size = 0
}

func (list *ArrayList[T]) ToSlice() []T {
	return list.data
}

// First 첫 번째 요소 반환
func (list *ArrayList[T]) First() T {
	if list.IsEmpty() {
		var zero T
		return zero
	}
	v := list.Get(0)
	return v
}

// Last 마지막 요소 반환
func (list *ArrayList[T]) Last() T {
	if list.IsEmpty() {
		var zero T
		return zero
	}
	v := list.Get(list.Size() - 1)
	return v
}

// RemoveFirst 첫 번째 요소 제거
func (list *ArrayList[T]) RemoveFirst() T {
	if list.IsEmpty() {
		var zero T
		return zero
	}
	first := list.RemoveAt(0)
	return first
}

// RemoveLast 마지막 요소 제거
func (list *ArrayList[T]) RemoveLast() T {
	if list.IsEmpty() {
		var zero T
		return zero
	}
	last := list.RemoveAt(list.Size() - 1)
	return last
}

func (list *ArrayList[T]) ToList() IList[T] {
	return list
}

// SubList 마지막 요소 제거
func (list *ArrayList[T]) SubList(start, end int) IList[T] {
	return NewArrayListWithData[T](list.data[start:end])
}

func (list *ArrayList[T]) Sort(action func(i, j int) bool) {
	sort.SliceIsSorted(list.data, action)
}

// ContainsAll 여러 요소가 포함되어 있는지 확인
func (list *ArrayList[T]) ContainsAll(elements []T) bool {
	for _, element := range elements {
		if !list.Contains(element) {
			return false
		}
	}
	return true
}

// RemoveAll 여러 요소 제거
func (list *ArrayList[T]) RemoveAll(elements []T) bool {
	changed := false
	for _, element := range elements {
		if ok := list.Remove(element); ok {
			changed = true
		}
	}
	return changed
}

// RetainAll 특정 요소만 유지
func (list *ArrayList[T]) RetainAll(elements []T) bool {
	changed := false
	elementsSet := make(map[T]struct{})
	for _, element := range elements {
		elementsSet[element] = struct{}{}
	}
	for _, element := range list.data {
		if _, exists := elementsSet[element]; !exists {
			list.Remove(element)
			changed = true
		}
	}
	return changed
}

// Iterator 컬렉션의 Iterator 반환
func (list *ArrayList[T]) Iterator() IIterator[T] {
	elements := list.ToSlice()
	return &Iterator[T]{data: elements, index: 0}
}

// ListIterator 컬렉션의 ListIterator 반환
func (list *ArrayList[T]) ListIterator() IListIterator[T] {
	elements := list.ToSlice()
	return &ListIterator[T]{data: elements, cursor: 0, lastIndex: -1}
}

// Example usage
func main() {
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
