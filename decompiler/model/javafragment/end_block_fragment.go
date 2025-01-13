package javafragment

import "fmt"

func NewEndBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight int, label string, start *StartBlockFragment) EndBlockFragment {
	f := EndBlockFragment{
		EndFlexibleBlockFragment: NewEndFlexibleBlockFragment(minimalLineCount,
			lineCount, maximalLineCount, weight, label),
		Start: start,
	}

	f.Start.End = &f

	return f
}

type EndBlockFragment struct {
	EndFlexibleBlockFragment

	Start *StartBlockFragment
}

func (f *EndBlockFragment) IncLineCount(force bool) bool {
	if f.LineCount < f.MaximalLineCount {
		f.LineCount = f.LineCount + 1
		return true
	}
	return false
}

func (f *EndBlockFragment) DecLineCount(force bool) bool {
	if f.LineCount > f.MinimalLineCount {
		f.LineCount = f.LineCount - 1
		return true
	}
	return false
}

func (f *EndBlockFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitEndBlockFragment(f)
}

func (f *EndBlockFragment) String() string {
	return fmt.Sprintf("EndBlockFragment { start: %s, end: %s", f.Start.String(), f.EndFlexibleBlockFragment.String())
}
