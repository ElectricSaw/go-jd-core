package javafragment

import "bitbucket.org/coontec/javaClass/java/model/fragment"

func NewEndSingleStatementBlockFragment(minimalLineCount int, lineCount int, maximalLineCount int, weight int, label string, start *StartSingleStatementBlockFragment) EndSingleStatementBlockFragment {
	f := EndSingleStatementBlockFragment{
		EndFlexibleBlockFragment: fragment.NewEndFlexibleBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight, label),
		start:                    start,
	}

	f.start.SetEndSingleStatementBlockFragment(&f)

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
	if f.LineCount() < f.MaximalLineCount {
		f.SetLineCount(f.LineCount() - 1)

		if !force {
			if f.start.LineCount() == 0 {
				f.start.SetLineCount(1)
			}
		}

		return true
	}
	return false
}

func (f *EndSingleStatementBlockFragment) DecLineCount(force bool) bool {
	if f.LineCount() > f.MinimalLineCount {
		f.SetLineCount(f.LineCount() - 1)

		if !force {
			if f.LineCount() == 0 {
				f.start.SetLineCount(0)
			}
		}

		return true
	}
	return false
}

func (f *EndSingleStatementBlockFragment) Accept(visitor JavaFragmentVisitor) {
	visitor.VisitEndSingleStatementBlockFragment(f)
}
