package javafragment

func NewStartSingleStatementBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight int, label string) StartSingleStatementBlockFragment {
	return StartSingleStatementBlockFragment{
		StartFlexibleBlockFragment: NewStartFlexibleBlockFragment(minimalLineCount, lineCount,
			maximalLineCount, weight, label),
	}
}

type StartSingleStatementBlockFragment struct {
	StartFlexibleBlockFragment

	End *EndSingleStatementBlockFragment
}

func (f *StartSingleStatementBlockFragment) IncLineCount(force bool) bool {
	if f.LineCount < f.MaximalLineCount {
		f.LineCount = f.LineCount + 1

		if !force {
			if f.End.LineCount == 0 {
				f.End.LineCount = 1
			}
		}

		return true
	}
	return false
}

func (f *StartSingleStatementBlockFragment) DecLineCount(force bool) bool {
	if f.LineCount > f.MinimalLineCount {
		f.LineCount = f.LineCount - 1

		if !force {
			if f.LineCount == 1 {
				f.End.LineCount = 1
			}
		}

		return true
	}
	return false
}

func (f *StartSingleStatementBlockFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitStartSingleStatementBlockFragment(f)
}
