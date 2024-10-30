package fragment

import "fmt"

func NewFixedFragment(firstLineNumber int, lastLineNumber int) FixedFragment {
	return FixedFragment{firstLineNumber, lastLineNumber}
}

type FixedFragment struct {
	FirstLineNumber int
	LastLineNumber  int
}

func (f *FixedFragment) Accept(visitor FragmentVisitor) {
	visitor.VisitFixedFragment(f)
}

func (f *FixedFragment) String() string {
	return fmt.Sprintf("{first-line-number=%d, last-line-number=%d}", f.FirstLineNumber, f.LastLineNumber)
}
