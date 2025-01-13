package javafragment

func NewStartBodyFragment(minimalLineCount, lineCount, maximalLineCount, weight int,
	label string) StartBodyFragment {
	return StartBodyFragment{
		StartFlexibleBlockFragment: NewStartFlexibleBlockFragment(minimalLineCount,
			lineCount, maximalLineCount, weight, label),
	}
}

type StartBodyFragment struct {
	StartFlexibleBlockFragment

	End *EndBodyFragment
}

func (f *StartBodyFragment) IncLineCount(force bool) bool {
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

func (f *StartBodyFragment) DecLineCount(force bool) bool {
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

func (f *StartBodyFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitStartBodyFragment(f)
}
