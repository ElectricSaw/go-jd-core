package javafragment

import (
	"fmt"
)

func NewStartBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight int,
	label string) StartBlockFragment {
	return StartBlockFragment{
		StartFlexibleBlockFragment: NewStartFlexibleBlockFragment(minimalLineCount,
			lineCount, maximalLineCount, weight, label),
	}
}

type StartBlockFragment struct {
	StartFlexibleBlockFragment

	End *EndBlockFragment
}

func (f *StartBlockFragment) IncLineCount(force bool) bool {
	if f.LineCount < f.MaximalLineCount {
		f.LineCount = f.LineCount + 1

		if !force {
			if f.LineCount == 1 && f.End.LineCount == 0 {
				f.End.LineCount = f.LineCount
			}
		}

		return true
	}
	return false
}

func (f *StartBlockFragment) DecLineCount(force bool) bool {
	if f.LineCount > f.MinimalLineCount {
		f.LineCount = f.LineCount - 1

		if !force {
			if f.LineCount == 1 {
				f.End.LineCount = f.LineCount
			}
		}

		return true
	}
	return false
}

func (f *StartBlockFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitStartBlockFragment(f)
}

func (f *StartBlockFragment) String() string {
	return fmt.Sprintf("StartBlockFragment { start: %s, end: %s", f.StartFlexibleBlockFragment.String(), f.End.String())
}
