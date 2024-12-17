package util

import (
	"testing"
)

func TestDefaultBase(t *testing.T) {
	bitset := NewBitSet()
	_ = bitset.Set(10)
	_ = bitset.Set(64)
	bitset.Clear(10)
	isSet := bitset.Get(64)
	println("Bit 64 set:", isSet)
	println("Length:", bitset.Length())
	println("Cardinality:", bitset.Cardinality())
}
