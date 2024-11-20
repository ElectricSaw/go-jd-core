package javafragment

import "bitbucket.org/coontec/go-jd-core/class/model/fragment"

func NewEndSingleStatementBlockFragment(minimalLineCount int, lineCount int, maximalLineCount int, weight int, label string, start *StartSingleStatementBlockFragment) *EndSingleStatementBlockFragment {
	f := &EndSingleStatementBlockFragment{
		EndFlexibleBlockFragment: *fragment.NewEndFlexibleBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight, label),
		start:                    start,
	}

	f.start.SetEndSingleStatementBlockFragment(f)

	return f
}

type EndSingleStatementBlockFragment struct {
	fragment.EndFlexibleBlockFragment

	start *StartSingleStatementBlockFragment
}

func (f *EndSingleStatementBlockFragment) StartSingleStatementBlockFragment() *StartSingleStatementBlockFragment {
	return f.start
}

func (f *EndSingleStatementBlockFragment) IncLineCount(force bool) bool {
	if f.LineCount < f.MaximalLineCount {
		f.LineCount++

		if !force {
			if f.start.LineCount == 0 {
				f.start.LineCount = 1
			}
		}

		return true
	}
	return false
}

func (f *EndSingleStatementBlockFragment) DecLineCount(force bool) bool {
	if f.LineCount > f.MinimalLineCount {
		f.LineCount--

		if !force {
			if f.LineCount == 0 {
				f.start.LineCount = 0
			}
		}

		return true
	}
	return false
}

func (f *EndSingleStatementBlockFragment) Accept(visitor JavaFragmentVisitor) {
	visitor.VisitEndSingleStatementBlockFragment(f)
}
