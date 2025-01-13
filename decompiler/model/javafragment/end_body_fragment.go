package javafragment

func NewEndBodyFragment(minimalLineCount, lineCount, maximalLineCount, weight int,
	label string, start *StartBodyFragment) EndBodyFragment {
	f := EndBodyFragment{
		EndFlexibleBlockFragment: NewEndFlexibleBlockFragment(minimalLineCount,
			lineCount, maximalLineCount, weight, label),
		Start: start,
	}

	f.Start.End = &f

	return f
}

type EndBodyFragment struct {
	EndFlexibleBlockFragment

	Start *StartBodyFragment
}

func (f *EndBodyFragment) IncLineCount(force bool) bool {
	if f.LineCount < f.MaximalLineCount {
		f.LineCount = f.LineCount + 1

		if !force {
			if f.LineCount == 1 && f.Start.LineCount == 0 {
				f.Start.LineCount = f.LineCount
			}
		}

		return true
	}
	return false
}

func (f *EndBodyFragment) DecLineCount(force bool) bool {
	if f.LineCount > f.MinimalLineCount {
		f.LineCount = f.LineCount - 1

		if !force {
			if f.LineCount == 0 {
				f.Start.LineCount = f.LineCount
			}
		}

		return true
	}
	return false
}

func (f *EndBodyFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitEndBodyFragment(f)
}
