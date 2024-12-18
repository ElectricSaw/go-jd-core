package util

import (
	"fmt"
	"testing"
)

func TestBitSet(t *testing.T) {
	bs := NewBitSetWithSize(128)
	bs.Set(1)
	bs.Set(64)
	bs.Set(127)

	fmt.Println("BitSet:", bs.String())
	fmt.Println("Cardinality:", bs.Cardinality())

	bs.ClearAll(64)
	fmt.Println("After clearing bit 64:")
	fmt.Println("BitSet:", bs.String())
	fmt.Println("Cardinality:", bs.Cardinality())
}
