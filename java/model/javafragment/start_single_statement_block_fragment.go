package javafragment

import "bitbucket.org/coontec/javaClass/java/model/fragment"

func NewStartSingleStatementBlockFragment(minimalLineCount int, lineCount int, maximalLineCount int, weight int, label string) StartSingleStatementBlockFragment {
	return StartSingleStatementBlockFragment{
		StartFlexibleBlockFragment: fragment.NewStartFlexibleBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight, label),
	}
}

type StartSingleStatementBlockFragment struct {
	fragment.StartFlexibleBlockFragment

	end *EndSingleStatementBlockFragment
}

func (f *StartSingleStatementBlockFragment) EndSingleStatementBlockFragment() *EndSingleStatementBlockFragment {
	return f.end
}

func (f *StartSingleStatementBlockFragment) SetEndSingleStatementBlockFragment(end *EndSingleStatementBlockFragment) {
	f.end = end
}

func (f *StartSingleStatementBlockFragment) IncLineCount(force bool) bool {
	if f.LineCount() < f.MaximalLineCount {
		f.SetLineCount(f.LineCount() + 1)

		if !force {
			if f.end.LineCount() == 0 {
				f.end.SetLineCount(1)
			}
		}

		return true
	}
	return false
}

func (f *StartSingleStatementBlockFragment) DecLineCount(force bool) bool {
	if f.LineCount() > f.MinimalLineCount {
		f.SetLineCount(f.LineCount() - 1)

		if !force {
			if f.LineCount() == 1 {
				f.end.SetLineCount(1)
			}
		}

		return true
	}
	return false
}

func (f *StartSingleStatementBlockFragment) Accept(visitor JavaFragmentVisitor) {
	visitor.VisitStartSingleStatementBlockFragment(f)
}
