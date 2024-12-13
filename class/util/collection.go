package util

// ICollection 인터페이스: Java Collection을 Go 스타일로 변환
type ICollection[T comparable] interface {
	Size() int                     // 요소 개수 반환
	IsEmpty() bool                 // 컬렉션이 비어 있는지 확인
	Contains(element T) bool       // 특정 요소가 포함되어 있는지 확인
	Add(element T) bool            // 요소 추가
	Remove(element T) bool         // 요소 제거
	ContainsAll(elements []T) bool // 여러 요소가 포함되어 있는지 확인
	AddAll(elements []T) bool      // 여러 요소 추가
	RemoveAll(elements []T) bool   // 여러 요소 제거
	RetainAll(elements []T) bool   // 특정 요소만 유지
	Clear()                        // 모든 요소 제거
	ToSlice() []T                  // 컬렉션을 슬라이스로 변환
	Iterator() IIterator[T]        // Iterator 반환
}

// Collection 구현
type Collection[T comparable] struct {
	data map[T]struct{}
}

// NewCollection 새로운 BasicCollection 생성
func NewCollection[T comparable]() ICollection[T] {
	return &Collection[T]{data: make(map[T]struct{})}
}

// Size 요소 개수 반환
func (c *Collection[T]) Size() int {
	return len(c.data)
}

// IsEmpty 컬렉션이 비어 있는지 확인
func (c *Collection[T]) IsEmpty() bool {
	return len(c.data) == 0
}

// Contains 특정 요소가 포함되어 있는지 확인
func (c *Collection[T]) Contains(element T) bool {
	_, exists := c.data[element]
	return exists
}

// Add 요소 추가
func (c *Collection[T]) Add(element T) bool {
	if _, exists := c.data[element]; exists {
		return false
	}
	c.data[element] = struct{}{}
	return true
}

// Remove 요소 제거
func (c *Collection[T]) Remove(element T) bool {
	if _, exists := c.data[element]; exists {
		delete(c.data, element)
		return true
	}
	return false
}

// ContainsAll 여러 요소가 포함되어 있는지 확인
func (c *Collection[T]) ContainsAll(elements []T) bool {
	for _, element := range elements {
		if !c.Contains(element) {
			return false
		}
	}
	return true
}

// AddAll 여러 요소 추가
func (c *Collection[T]) AddAll(elements []T) bool {
	changed := false
	for _, element := range elements {
		if c.Add(element) {
			changed = true
		}
	}
	return changed
}

// RemoveAll 여러 요소 제거
func (c *Collection[T]) RemoveAll(elements []T) bool {
	changed := false
	for _, element := range elements {
		if c.Remove(element) {
			changed = true
		}
	}
	return changed
}

// RetainAll 특정 요소만 유지
func (c *Collection[T]) RetainAll(elements []T) bool {
	changed := false
	elementsSet := make(map[T]struct{})
	for _, element := range elements {
		elementsSet[element] = struct{}{}
	}
	for element := range c.data {
		if _, exists := elementsSet[element]; !exists {
			c.Remove(element)
			changed = true
		}
	}
	return changed
}

// Clear 모든 요소 제거
func (c *Collection[T]) Clear() {
	c.data = make(map[T]struct{})
}

// ToSlice 컬렉션을 슬라이스로 변환
func (c *Collection[T]) ToSlice() []T {
	slice := make([]T, 0, len(c.data))
	for element := range c.data {
		slice = append(slice, element)
	}
	return slice
}

// Iterator 컬렉션의 Iterator 반환
func (c *Collection[T]) Iterator() IIterator[T] {
	elements := c.ToSlice()
	return &Iterator[T]{data: elements, index: 0}
}
