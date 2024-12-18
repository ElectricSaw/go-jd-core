package util

import (
	"errors"
)

// IIterator Java Iterator와 유사한 구조
type IIterator[T comparable] interface {
	HasNext() bool            // 다음 요소가 있는지 확인
	Next() T                  // 다음 요소 반환
	Remove() error            // 현재 요소 제거 (옵션)
	ForEachRemaining(func(T)) // 남은 요소 처리
}

// Iterator 구현
type Iterator[T comparable] struct {
	data    []T
	index   int
	removed bool // Remove 호출 여부를 추적
}

// NewIterator 슬라이스를 기반으로 한 Iterator 생성자
func NewIterator[T comparable]() IIterator[T] {
	return NewIteratorWithSlice(make([]T, 0))
}

// NewIteratorWithSlice 슬라이스를 기반으로 한 Iterator 생성자
func NewIteratorWithSlice[T comparable](data []T) IIterator[T] {
	return &Iterator[T]{data: data, index: 0, removed: false}
}

// HasNext 다음 요소가 있는지 확인
func (it *Iterator[T]) HasNext() bool {
	return it.index < len(it.data)
}

// Next 다음 요소 반환
func (it *Iterator[T]) Next() T {
	if !it.HasNext() {
		var zero T
		return zero
	}
	element := it.data[it.index]
	it.index++
	it.removed = false // 다음 요소로 이동 시 Remove 호출 상태 초기화
	return element
}

// Remove 현재 요소 제거
func (it *Iterator[T]) Remove() error {
	if it.index == 0 || it.removed {
		return errors.New("illegal state for remove")
	}
	// 슬라이스에서 요소 제거
	it.data = append(it.data[:it.index-1], it.data[it.index:]...)
	it.index-- // 인덱스 조정
	it.removed = true
	return nil
}

// ForEachRemaining 남은 요소를 모두 처리
func (it *Iterator[T]) ForEachRemaining(action func(T)) {
	for it.HasNext() {
		element := it.Next()
		action(element)
	}
}
