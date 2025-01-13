package javafragment

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
	"math"
)

func NewStartStatementsBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight int,
	label string) StartStatementsBlockFragment {
	return NewStartStatementsBlockFragmentWithGroup(minimalLineCount, lineCount,
		maximalLineCount, weight, label, NewStartStatementsBlockFragmentGroup())
}

func NewStartStatementsBlockFragmentWithGroup(minimalLineCount, lineCount, maximalLineCount, weight int,
	label string, group StartStatementsBlockFragmentGroup) StartStatementsBlockFragment {
	f := StartStatementsBlockFragment{
		StartFlexibleBlockFragment: NewStartFlexibleBlockFragment(minimalLineCount, lineCount,
			maximalLineCount, weight, label),
		Group: group,
	}
	f.Group.Add(&f)
	return f
}

type StartStatementsBlockFragment struct {
	StartFlexibleBlockFragment

	Group StartStatementsBlockFragmentGroup
}

func (f *StartStatementsBlockFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitStartStatementsBlockFragment(f)
}

func NewStartStatementsBlockFragmentGroup() StartStatementsBlockFragmentGroup {
	return StartStatementsBlockFragmentGroup{minimalLineCount: math.MaxInt}
}

type StartStatementsBlockFragmentGroup struct {
	util.DefaultList[IFlexibleFragment]
	minimalLineCount int
}

func (f *StartStatementsBlockFragmentGroup) MinimalLineCount() int {
	if f.minimalLineCount == math.MaxInt {
		for _, frag := range f.ToSlice() {
			if f.minimalLineCount > frag.GetLineCount() {
				f.minimalLineCount = frag.GetLineCount()
			}
		}
	}

	return f.minimalLineCount
}
