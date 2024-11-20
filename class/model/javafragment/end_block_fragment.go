package javafragment

import "bitbucket.org/coontec/go-jd-core/class/model/fragment"

func NewEndBlockFragment(minimalLineCount int, lineCount int, maximalLineCount int, weight int, label string, start *StartBlockFragment) *EndBlockFragment {
	f := &EndBlockFragment{
		EndFlexibleBlockFragment: *fragment.NewEndFlexibleBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight, label),
		start:                    start,
	}

	f.start.SetEndArrayInitializerBlockFragment(f)

	return f
}

type EndBlockFragment struct {
	fragment.EndFlexibleBlockFragment

	start *StartBlockFragment
}

func (f *EndBlockFragment) StartArrayInitializerBlockFragment() *StartBlockFragment {
	return f.start
}

func (f *EndBlockFragment) IncLineCount(force bool) bool {
	if f.LineCount < f.MaximalLineCount {
		f.LineCount++
		return true
	}
	return false
}

func (f *EndBlockFragment) DecLineCount(force bool) bool {
	if f.LineCount > f.MinimalLineCount {
		f.LineCount--
		return true
	}
	return false
}

func (f *EndBlockFragment) Accept(visitor JavaFragmentVisitor) {
	visitor.VisitEndBlockFragment(f)
}
