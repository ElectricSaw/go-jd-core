package javafragment

func NewEndSingleStatementBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight int,
	label string, start *StartSingleStatementBlockFragment) EndSingleStatementBlockFragment {
	f := EndSingleStatementBlockFragment{
		EndFlexibleBlockFragment: NewEndFlexibleBlockFragment(minimalLineCount,
			lineCount, maximalLineCount, weight, label),
		Start: start,
	}

	f.Start.End = &f

	return f
}

type EndSingleStatementBlockFragment struct {
	EndFlexibleBlockFragment

	Start *StartSingleStatementBlockFragment
}

func (f *EndSingleStatementBlockFragment) IncLineCount(force bool) bool {
	if f.LineCount < f.MaximalLineCount {
		f.LineCount = f.LineCount + 1

		if !force {
			if f.Start.LineCount == 0 {
				f.Start.LineCount = 1
			}
		}

		return true
	}
	return false
}

func (f *EndSingleStatementBlockFragment) DecLineCount(force bool) bool {
	if f.LineCount > f.MinimalLineCount {
		f.LineCount = f.LineCount - 1

		if !force {
			if f.LineCount == 0 {
				f.Start.LineCount = 0
			}
		}

		return true
	}
	return false
}

func (f *EndSingleStatementBlockFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitEndSingleStatementBlockFragment(f)
}
