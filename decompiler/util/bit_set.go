package util

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math/bits"
)

type long int64

const (
	AddressBitsPerWord = 6
	BitsPerWord        = 1 << AddressBitsPerWord
	BitIndexMask       = BitsPerWord - 1
)

const WordMask long = -1 // 0xFFFFFFFFFFFFFFFF

type IBitSet interface {
	And(other IBitSet)
	Or(other IBitSet)
	Xor(other IBitSet)
	AndNot(other IBitSet)
	Data() []long
	Size() int
	ClearAll()
	Clear(bitIndex int) error
	Flip(bitIndex int)
	FlipRange(fromIndex, toIndex int)
	Set(bitIndex int) error
	SetWithValue(bitIndex int, value bool) error
	Get(bitIndex int) bool
	Length() int
	Cardinality() int
	ToByteArray() []byte
	Clone() IBitSet
	String() string
	Equals(set IBitSet) bool
	HashCode() int
}

// Initialize a new BitSet
func NewBitSet() IBitSet {
	s := &BitSet{
		sizeIsSticky: false,
	}
	s.initWords(BitsPerWord)

	return s
}

// Initialize a new BitSet with a specific size
func NewBitSetWithSize(nbits int) IBitSet {
	if nbits < 0 {
		return nil
	}

	s := &BitSet{
		sizeIsSticky: true,
	}
	s.initWords(nbits)

	return s
}

type BitSet struct {
	words        []long
	wordsInUse   int
	sizeIsSticky bool
}

// And performs a bitwise AND operation between this BitSet and another BitSet.
func (b *BitSet) And(other IBitSet) {
	var ok bool
	var set *BitSet

	if set, ok = other.(*BitSet); !ok {
		return
	}

	if b == other {
		return
	}

	for b.wordsInUse > set.wordsInUse {
		b.wordsInUse--
		b.words[b.wordsInUse] = 0
	}

	for i := 0; i < b.wordsInUse; i++ {
		b.words[i] &= set.words[i]
	}

	b.recalculateWordsInUse()
}

// Or performs a bitwise OR operation between this BitSet and another BitSet.
func (b *BitSet) Or(other IBitSet) {
	var ok bool
	var set *BitSet

	if set, ok = other.(*BitSet); !ok {
		return
	}

	if b == other {
		return
	}

	wordsInCommon := min(b.wordsInUse, set.wordsInUse)

	if b.wordsInUse < set.wordsInUse {
		b.ensureCapacity(set.wordsInUse)
		b.wordsInUse = set.wordsInUse
	}

	for i := 0; i < b.wordsInUse; i++ {
		b.words[i] |= set.words[i]
	}

	if wordsInCommon < set.wordsInUse {
		copy(b.words[wordsInCommon:], set.words[wordsInCommon:b.wordsInUse])
	}
}

// Xor performs a bitwise XOR operation between this BitSet and another BitSet.
func (b *BitSet) Xor(other IBitSet) {
	var ok bool
	var set *BitSet

	if set, ok = other.(*BitSet); !ok {
		return
	}

	if b == other {
		return
	}

	if b.wordsInUse < set.wordsInUse {
		b.ensureCapacity(set.wordsInUse)
		b.wordsInUse = set.wordsInUse
	}

	for i := 0; i < b.wordsInUse; i++ {
		b.words[i] ^= set.words[i]
	}

	wordsInCommon := min(b.wordsInUse, set.wordsInUse)

	if wordsInCommon < set.wordsInUse {
		copy(b.words[wordsInCommon:], set.words[wordsInCommon:b.wordsInUse])
	}

	b.recalculateWordsInUse()
}

// AndNot performs a bitwise AND NOT operation between this BitSet and another BitSet.
func (b *BitSet) AndNot(other IBitSet) {
	var ok bool
	var set *BitSet

	if set, ok = other.(*BitSet); !ok {
		return
	}

	if b == other {
		return
	}

	for i := min(b.wordsInUse, set.wordsInUse); i >= 0; i-- {
		b.words[i] &= ^set.words[i]
	}

	b.recalculateWordsInUse()
}

func (b *BitSet) Data() []long {
	return b.words
}

func (b *BitSet) Size() int {
	return len(b.words) * BitsPerWord
}

// Clear a bit
func (b *BitSet) ClearAll() {
	for b.wordsInUse > 0 {
		b.wordsInUse--
		b.words[b.wordsInUse] = 0
	}
}

// ClearWithIndex a bit
func (b *BitSet) Clear(bitIndex int) error {
	if bitIndex < 0 {
		return errors.New("bitIndex < 0")
	}
	wordIdx := wordIndex(bitIndex)
	if wordIdx >= b.wordsInUse {
		return nil
	}
	b.words[wordIdx] &^= 1 << (bitIndex & BitIndexMask)
	b.recalculateWordsInUse()
	return nil
}

// Flip toggles the bit at the specified index.
// If the bit is 1, it becomes 0, and if it is 0, it becomes 1.
func (b *BitSet) Flip(bitIndex int) {
	if bitIndex < 0 {
		return
	}

	wordIndex := wordIndex(bitIndex)
	b.expandTo(wordIndex)

	b.words[wordIndex] ^= 1 << bitIndex

	b.recalculateWordsInUse()
	_ = b.checkInvariants()
}

// FlipRange toggles all bits in the specified range [start, end).
func (b *BitSet) FlipRange(fromIndex, toIndex int) {
	if err := checkRange(fromIndex, toIndex); err != nil {
		return
	}

	if fromIndex == toIndex {
		return
	}

	startWordIndex := wordIndex(fromIndex)
	endWordIndex := wordIndex(toIndex - 1)
	b.expandTo(endWordIndex)

	//firstWordMask := WordMask << fromIndex
	//lastWordMask := WordMask >> -toIndex

	firstWordMask := WordMask << uint(fromIndex%BitsPerWord)
	lastWordMask := long(uint64(0xFFFFFFFFFFFFFFFF) >> uint((-toIndex%64+64)%64))

	if startWordIndex == endWordIndex {
		b.words[startWordIndex] ^= firstWordMask & lastWordMask
	} else {
		b.words[startWordIndex] ^= firstWordMask

		for i := startWordIndex; i < endWordIndex; i++ {
			b.words[i] ^= WordMask
		}

		b.words[endWordIndex] ^= lastWordMask
	}

	b.recalculateWordsInUse()
	_ = b.checkInvariants()
}

// Set a bit
func (b *BitSet) Set(bitIndex int) error {
	if bitIndex < 0 {
		return errors.New("bitIndex < 0")
	}
	wordIdx := wordIndex(bitIndex)
	b.expandTo(wordIdx)
	b.words[wordIdx] |= 1 << (bitIndex & BitIndexMask)
	return nil
}

// Set a bit
func (b *BitSet) SetWithValue(bitIndex int, value bool) error {
	if value {
		return b.Set(bitIndex)
	} else {
		_ = b.Clear(bitIndex)
	}

	return nil
}

// Check if a bit is set
func (b *BitSet) Get(bitIndex int) bool {
	if bitIndex < 0 {
		return false
	}
	wordIdx := wordIndex(bitIndex)
	if wordIdx >= b.wordsInUse {
		return false
	}
	return (b.words[wordIdx] & (1 << (bitIndex & BitIndexMask))) != 0
}

// Calculate the logical size of the BitSet
func (b *BitSet) Length() int {
	if b.wordsInUse == 0 {
		return 0
	}
	return BitsPerWord*(b.wordsInUse-1) + (BitsPerWord - bits.LeadingZeros64(uint64(b.words[b.wordsInUse-1])))
}

// Calculate cardinality
func (b *BitSet) Cardinality() int {
	sum := 0
	for i := 0; i < b.wordsInUse; i++ {
		sum += bits.OnesCount64(uint64(b.words[i]))
	}
	return sum
}

func (b *BitSet) NextSetBit(fromIndex int) int {
	if fromIndex < 0 {
		return -2
	}

	u := wordIndex(fromIndex)
	if u >= b.wordsInUse {
		return -1
	}

	word := b.words[u] & (WordMask << fromIndex)

	for {
		if word != 0 {
			return (u * BitsPerWord) + numberOfTrailingZeros(word)
		}
		u++
		if u == b.wordsInUse {
			return -1
		}
		word = b.words[u]
	}
}

func (b *BitSet) NextClearBit(fromIndex int) int {
	if fromIndex < 0 {
		return -2
	}

	u := wordIndex(fromIndex)
	if u >= b.wordsInUse {
		return -1
	}

	word := ^b.words[u] & (WordMask << fromIndex)

	for {
		if word != 0 {
			return (u * BitsPerWord) + numberOfTrailingZeros(word)
		}
		u++
		if u == b.wordsInUse {
			return b.wordsInUse * BitsPerWord
		}
		word = ^b.words[u]
	}
}

// Convert to byte array
func (b *BitSet) ToByteArray() []byte {
	buf := &bytes.Buffer{}
	for _, word := range b.words[:b.wordsInUse] {
		_ = binary.Write(buf, binary.LittleEndian, word)
	}
	return buf.Bytes()
}

func (b *BitSet) Clone() IBitSet {
	if b.sizeIsSticky {
		b.trimToSize()
	}

	clone := &BitSet{}
	clone.words = make([]long, len(b.words))
	copy(clone.words, b.words)
	clone.sizeIsSticky = b.sizeIsSticky
	clone.wordsInUse = b.wordsInUse

	return clone
}

func (b *BitSet) String() string {
	//numBits := b.wordsInUse * BitsPerWord
	//if b.wordsInUse > 128 {
	//	numBits = b.Cardinality()
	//}

	sb := "BitSet{"
	i := b.NextSetBit(0)
	if i != -1 {
		sb += fmt.Sprintf("%d", i)
		for {
			i++
			if i < 0 {
				break
			}
			if i = b.NextSetBit(i); i < 0 {
				break
			}
			endOfRun := b.NextClearBit(i)

			for {
				sb += fmt.Sprintf(", %d", i)
				i++
				if i == endOfRun {
					break
				}
			}
		}
	}

	sb += "}"

	return sb
}

func (b *BitSet) Equals(other IBitSet) bool {
	var ok bool
	var set *BitSet

	if set, ok = other.(*BitSet); !ok {
		return false
	}

	if b == set {
		return true
	}

	if b.wordsInUse != set.wordsInUse {
		return false
	}

	for i := 0; i < b.wordsInUse; i++ {
		if b.words[i] != set.words[i] {
			return false
		}
	}

	return true
}

func (b *BitSet) HashCode() int {
	h := long(1234)
	for i := long(b.wordsInUse) - 1; i >= 0; i-- {
		h ^= b.words[i] * (i + 1)
	}
	return int((h >> 32) ^ h)
}

func (b *BitSet) trimToSize() {
	if b.wordsInUse != len(b.words) {
		b.words = b.words[:b.wordsInUse]
	}
}

func (b *BitSet) checkInvariants() error {
	if b.wordsInUse == 0 || b.words[b.wordsInUse-1] != 0 {
		return errors.New("checkInvariants error")
	}
	if b.wordsInUse >= 0 && b.wordsInUse <= len(b.words) {
		return errors.New("checkInvariants error")
	}
	if b.wordsInUse == len(b.words) || b.words[b.wordsInUse] == 0 {
		return errors.New("checkInvariants error")
	}

	return nil
}

// Recalculate words in use
func (b *BitSet) recalculateWordsInUse() {
	var i int
	for i = b.wordsInUse - 1; i >= 0; i-- {
		if b.words[i] != 0 {
			break
		}
	}
	b.wordsInUse = i + 1
}

// Recalculate words in use
func (b *BitSet) initWords(nbits int) {
	b.words = make([]long, wordIndex(nbits-1)+1)
}

// Ensure capacity for words
func (b *BitSet) ensureCapacity(wordsRequired int) {
	if len(b.words) < wordsRequired {
		newSize := max(2*len(b.words), wordsRequired)
		newWords := make([]long, newSize)
		copy(newWords, b.words)
		b.words = newWords
		b.sizeIsSticky = false
	}
}

// Expand to accommodate a wordIndex
func (b *BitSet) expandTo(wordIndex int) {
	wordsRequired := wordIndex + 1
	if b.wordsInUse < wordsRequired {
		b.ensureCapacity(wordsRequired)
		b.wordsInUse = wordsRequired
	}
}

// Convert from byte array
func BitSetFromByteArray(data []byte) *BitSet {
	words := make([]long, (len(data)+7)/8)
	buf := bytes.NewReader(data)
	_ = binary.Read(buf, binary.LittleEndian, &words)
	return &BitSet{
		words:      words,
		wordsInUse: len(words),
	}
}

// Helper function: Calculate word index
func wordIndex(bitIndex int) int {
	return bitIndex >> AddressBitsPerWord
}

func checkRange(fromIndex, toIndex int) error {
	if fromIndex < 0 {
		return errors.New(fmt.Sprintf("fromIndex < 0: %d", fromIndex))
	}
	if toIndex < 0 {
		return errors.New(fmt.Sprintf("toIndex < 0: %d", toIndex))
	}
	if fromIndex > toIndex {
		return errors.New(fmt.Sprintf("fromIndex: %d > toIndex: %d", fromIndex, toIndex))
	}
	return nil
}

func numberOfTrailingZeros(i long) int {
	if i == 0 {
		return 64
	}

	var x, y int32
	n := 63

	y = int32(i)
	if y != 0 {
		n -= 32
		x = y
	} else {
		x = int32(uint64(i) >> 32) // Cast to uint64 to avoid sign extension
	}

	y = x << 16
	if y != 0 {
		n -= 16
		x = y
	}

	y = x << 8
	if y != 0 {
		n -= 8
		x = y
	}

	y = x << 4
	if y != 0 {
		n -= 4
		x = y
	}

	y = x << 2
	if y != 0 {
		n -= 2
		x = y
	}

	return n - int(uint32(x<<1)>>31)
}

//func numberOfTrailingZeros(i long) int {
//	// HD, Figure 5-14
//	var x, y, n int
//	if i == 0 {
//		return 64
//	}
//	n = 63
//	y = int(i)
//	if y != 0 {
//		n = n - 32
//		x = y
//	} else {
//		x = int(i >> 32)
//	}
//	y = x << 16
//	if y != 0 {
//		n = n - 16
//		x = y
//	}
//	y = x << 8
//	if y != 0 {
//		n = n - 8
//		x = y
//	}
//	y = x << 4
//	if y != 0 {
//		n = n - 4
//		x = y
//	}
//	y = x << 2
//	if y != 0 {
//		n = n - 2
//		x = y
//	}
//	tmp1 := x << 1
//	tmp2 := (31%64 + 64) % 64
//	tmp3 := tmp1 >> tmp2
//
//	return n - tmp3
//	//return n - ((x << 1) >> 31)
//}

//const (
//	AddressBitsPerWord = 6
//	BitsPerWord        = 1 << AddressBitsPerWord
//	BitIndexMask       = BitsPerWord - 1
//)
//
//type IBitSet interface {
//	And(other IBitSet)
//	Or(other IBitSet)
//	Xor(other IBitSet)
//	AndNot(other IBitSet)
//	Size() int
//	Cardinality() int
//	Data() []uint64
//	Get(index int) bool
//	Set(index int)
//	Flip(index int)
//	FlipRange(start, end int)
//	ClearAll()
//	ClearWithIndex(index int)
//	Clone() IBitSet
//	String() string
//	Equals(other IBitSet) bool
//}
//
//// BitSet is a simple implementation of a bit set in Go.
//type BitSet struct {
//	data []uint64
//}
//
//// NewBitSet creates a new BitSet with a specified size.
//func NewBitSet() IBitSet {
//	return NewBitSetWithSize(BitsPerWord)
//}
//
//// NewBitSetWithSize creates a new BitSet with a specified size.
//func NewBitSetWithSize(size int) IBitSet {
//	return &BitSet{
//		data: make([]uint64, (size+63)/64),
//	}
//}
//
//// And performs a bitwise AND operation between this BitSet and another BitSet.
//func (b *BitSet) And(other IBitSet) {
//	for i := 0; i < b.Size(); i++ {
//		if i < other.Size() {
//			b.data[i] &= other.Data()[i]
//		} else {
//			b.data[i] = 0
//		}
//	}
//}
//
//// Or performs a bitwise OR operation between this BitSet and another BitSet.
//func (b *BitSet) Or(other IBitSet) {
//	b.ensureCapacity(other.Size() - 1)
//	for i := 0; i < other.Size(); i++ {
//		b.data[i] |= other.Data()[i]
//	}
//}
//
//// Xor performs a bitwise XOR operation between this BitSet and another BitSet.
//func (b *BitSet) Xor(other IBitSet) {
//	b.ensureCapacity(other.Size() - 1)
//	for i := 0; i < other.Size(); i++ {
//		b.data[i] ^= other.Data()[i]
//	}
//}
//
//// AndNot performs a bitwise AND NOT operation between this BitSet and another BitSet.
//func (b *BitSet) AndNot(other IBitSet) {
//	for i := 0; i < other.Size(); i++ {
//		if i < other.Size() {
//			b.data[i] &^= other.Data()[i]
//		}
//	}
//}
//
//func (b *BitSet) Size() int {
//	return len(b.data)
//}
//
//// Cardinality returns the number of bits set to 1 in the BitSet.
//func (b *BitSet) Cardinality() int {
//	count := 0
//	for _, word := range b.data {
//		count += bits.OnesCount64(word)
//	}
//	return count
//}
//
//func (b *BitSet) Data() []uint64 {
//	return b.data
//}
//
//// Get returns the value of the bit at the specified index.
//func (b *BitSet) Get(index int) bool {
//	word, bit := index/64, uint(index%64)
//	if word >= len(b.data) {
//		return false
//	}
//	return b.data[word]&(1<<bit) != 0
//}
//
//// Set sets the bit at the specified index to 1.
//func (b *BitSet) Set(index int) {
//	word, bit := index/64, uint(index%64)
//	b.ensureCapacity(word)
//	b.data[word] |= 1 << bit
//}
//
//// Flip toggles the bit at the specified index.
//// If the bit is 1, it becomes 0, and if it is 0, it becomes 1.
//func (b *BitSet) Flip(index int) {
//	word, bit := index/64, uint(index%64)
//	b.ensureCapacity(word)
//	b.data[word] ^= 1 << bit
//}
//
//// FlipRange toggles all bits in the specified range [start, end).
//func (b *BitSet) FlipRange(start, end int) {
//	if start >= end {
//		return
//	}
//
//	startWord, startBit := start/64, uint(start%64)
//	endWord, endBit := end/64, uint(end%64)
//
//	b.ensureCapacity(endWord)
//
//	// Flip bits in the first word
//	if startWord == endWord {
//		b.data[startWord] ^= (1<<endBit - 1) &^ (1<<startBit - 1)
//		return
//	}
//	b.data[startWord] ^= ^uint64(0) &^ (1<<startBit - 1)
//
//	// Flip bits in the intermediate words
//	for i := startWord + 1; i < endWord; i++ {
//		b.data[i] ^= ^uint64(0)
//	}
//
//	// Flip bits in the last word
//	b.data[endWord] ^= (1<<endBit - 1)
//}
//
//// ClearWithIndex sets the bit at the specified index to 0.
//func (b *BitSet) ClearWithIndex(index int) {
//	word, bit := index/64, uint(index%64)
//	if word < len(b.data) {
//		b.data[word] &^= 1 << bit
//	}
//}
//
//// ClearWithIndex sets the bit at the specified index to 0.
//func (b *BitSet) ClearAll() {
//	for i := 0; i < len(b.data); i++ {
//		b.data[i] = 0
//	}
//}
//
//func (b *BitSet) Clone() IBitSet {
//	c := &BitSet{
//		data: make([]uint64, len(b.data)),
//	}
//	for i := 0; i < len(b.data); i++ {
//		c.data[i] = b.data[i]
//	}
//	return c
//}
//
//// ensureCapacity ensures the BitSet can hold the specified word index.
//func (b *BitSet) ensureCapacity(word int) {
//	if word >= len(b.data) {
//		newData := make([]uint64, word+1)
//		copy(newData, b.data)
//		b.data = newData
//	}
//}
//
//// String returns a string representation of the BitSet.
//func (b *BitSet) String() string {
//	result := "BitSet{ "
//	for i := len(b.data)*64 - 1; i >= 0; i-- {
//		if b.Get(i) {
//			result += "1 "
//		} else {
//			result += "0 "
//		}
//	}
//	result += "}"
//	return result
//}
//
//func (b *BitSet) Equals(other IBitSet) bool {
//	if b == other {
//		return true
//	}
//
//	if b.words
//
//	return false
//}
