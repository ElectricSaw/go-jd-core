package util

import "math/bits"

const (
	AddressBitsPerWord = 6
	BitsPerWord        = 1 << AddressBitsPerWord
	BitIndexMask       = BitsPerWord - 1
)

type IBitSet interface {
	And(other IBitSet)
	Or(other IBitSet)
	Xor(other IBitSet)
	AndNot(other IBitSet)
	Size() int
	Cardinality() int
	Data() []uint64
	Get(index int) bool
	Set(index int)
	Flip(index int)
	FlipRange(start, end int)
	ClearAll()
	Clear(index int)
	Clone() IBitSet
	String() string
}

// BitSet is a simple implementation of a bit set in Go.
type BitSet struct {
	data []uint64
}

// NewBitSet creates a new BitSet with a specified size.
func NewBitSet() IBitSet {
	return NewBitSetWithSize(BitsPerWord)
}

// NewBitSetWithSize creates a new BitSet with a specified size.
func NewBitSetWithSize(size int) IBitSet {
	return &BitSet{
		data: make([]uint64, (size+63)/64),
	}
}

// And performs a bitwise AND operation between this BitSet and another BitSet.
func (b *BitSet) And(other IBitSet) {
	for i := 0; i < b.Size(); i++ {
		if i < other.Size() {
			b.data[i] &= other.Data()[i]
		} else {
			b.data[i] = 0
		}
	}
}

// Or performs a bitwise OR operation between this BitSet and another BitSet.
func (b *BitSet) Or(other IBitSet) {
	b.ensureCapacity(other.Size() - 1)
	for i := 0; i < other.Size(); i++ {
		b.data[i] |= other.Data()[i]
	}
}

// Xor performs a bitwise XOR operation between this BitSet and another BitSet.
func (b *BitSet) Xor(other IBitSet) {
	b.ensureCapacity(other.Size() - 1)
	for i := 0; i < other.Size(); i++ {
		b.data[i] ^= other.Data()[i]
	}
}

// AndNot performs a bitwise AND NOT operation between this BitSet and another BitSet.
func (b *BitSet) AndNot(other IBitSet) {
	for i := 0; i < other.Size(); i++ {
		if i < other.Size() {
			b.data[i] &^= other.Data()[i]
		}
	}
}

func (b *BitSet) Size() int {
	return len(b.data)
}

// Cardinality returns the number of bits set to 1 in the BitSet.
func (b *BitSet) Cardinality() int {
	count := 0
	for _, word := range b.data {
		count += bits.OnesCount64(word)
	}
	return count
}

func (b *BitSet) Data() []uint64 {
	return b.data
}

// Get returns the value of the bit at the specified index.
func (b *BitSet) Get(index int) bool {
	word, bit := index/64, uint(index%64)
	if word >= len(b.data) {
		return false
	}
	return b.data[word]&(1<<bit) != 0
}

// Set sets the bit at the specified index to 1.
func (b *BitSet) Set(index int) {
	word, bit := index/64, uint(index%64)
	b.ensureCapacity(word)
	b.data[word] |= 1 << bit
}

// Flip toggles the bit at the specified index.
// If the bit is 1, it becomes 0, and if it is 0, it becomes 1.
func (b *BitSet) Flip(index int) {
	word, bit := index/64, uint(index%64)
	b.ensureCapacity(word)
	b.data[word] ^= 1 << bit
}

// FlipRange toggles all bits in the specified range [start, end).
func (b *BitSet) FlipRange(start, end int) {
	if start >= end {
		return
	}

	startWord, startBit := start/64, uint(start%64)
	endWord, endBit := end/64, uint(end%64)

	b.ensureCapacity(endWord)

	// Flip bits in the first word
	if startWord == endWord {
		b.data[startWord] ^= (1<<endBit - 1) &^ (1<<startBit - 1)
		return
	}
	b.data[startWord] ^= ^uint64(0) &^ (1<<startBit - 1)

	// Flip bits in the intermediate words
	for i := startWord + 1; i < endWord; i++ {
		b.data[i] ^= ^uint64(0)
	}

	// Flip bits in the last word
	b.data[endWord] ^= (1<<endBit - 1)
}

// Clear sets the bit at the specified index to 0.
func (b *BitSet) Clear(index int) {
	word, bit := index/64, uint(index%64)
	if word < len(b.data) {
		b.data[word] &^= 1 << bit
	}
}

// Clear sets the bit at the specified index to 0.
func (b *BitSet) ClearAll() {
	for i := 0; i < len(b.data); i++ {
		b.data[i] = 0
	}
}

func (b *BitSet) Clone() IBitSet {
	c := &BitSet{
		data: make([]uint64, len(b.data)),
	}
	for i := 0; i < len(b.data); i++ {
		c.data[i] = b.data[i]
	}
	return c
}

// ensureCapacity ensures the BitSet can hold the specified word index.
func (b *BitSet) ensureCapacity(word int) {
	if word >= len(b.data) {
		newData := make([]uint64, word+1)
		copy(newData, b.data)
		b.data = newData
	}
}

// String returns a string representation of the BitSet.
func (b *BitSet) String() string {
	result := ""
	for i := len(b.data)*64 - 1; i >= 0; i-- {
		if b.Get(i) {
			result += "1"
		} else {
			result += "0"
		}
	}
	return result
}
