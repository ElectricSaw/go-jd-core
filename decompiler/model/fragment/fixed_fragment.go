package fragment

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewFixedFragment(firstLineNumber, lastLineNumber int) intmod.IFixedFragment {
	return &FixedFragment{firstLineNumber, lastLineNumber}
}

type FixedFragment struct {
	firstLineNumber int
	lastLineNumber  int
}

func (f *FixedFragment) FirstLineNumber() int {
	return f.firstLineNumber
}

func (f *FixedFragment) SetFirstLineNumber(firstLineNumber int) {
	f.firstLineNumber = firstLineNumber
}

func (f *FixedFragment) LastLineNumber() int {
	return f.lastLineNumber
}

func (f *FixedFragment) SetLastLineNumber(lastLineNumber int) {
	f.lastLineNumber = lastLineNumber
}

func (f *FixedFragment) AcceptFragmentVisitor(visitor intmod.IFragmentVisitor) {
	visitor.VisitFixedFragment(f)
}

func (f *FixedFragment) String() string {
	return fmt.Sprintf("{first-line-number=%d, last-line-number=%d}", f.firstLineNumber, f.lastLineNumber)
}
