package javafragment

import "bitbucket.org/coontec/javaClass/java/model/fragment"

func NewEndBodyFragment(minimalLineCount int, lineCount int, maximalLineCount int, weight int, label string, start *StartBodyFragment) EndBodyFragment {
	f := EndBodyFragment{
		EndFlexibleBlockFragment: fragment.NewEndFlexibleBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight, label),
		start:                    start,
	}

	f.start.SetEndBodyFragment(&f)

	return f
}

type EndBodyFragment struct {
	fragment.EndFlexibleBlockFragment

	start *StartBodyFragment
}

func (f *EndBodyFragment) StartBodyFragment() *StartBodyFragment {
	return f.start
}

func (f *EndBodyFragment) IncLineCount(force bool) bool {
	if f.LineCount() < f.MaximalLineCount {
		f.SetLineCount(f.LineCount() + 1)

		if !force {
			if f.LineCount() == 1 && f.start.LineCount() == 0 {
				f.start.SetLineCount(f.LineCount())
			}
		}

		return true
	}
	return false
}

func (f *EndBodyFragment) DecLineCount(force bool) bool {
	if f.LineCount() > f.MinimalLineCount {
		f.SetLineCount(f.LineCount() - 1)

		if !force {
			if f.LineCount() == 0 {
				f.start.SetLineCount(f.LineCount())
			}
		}

		return true
	}
	return false
}

func (f *EndBodyFragment) Accept(visitor JavaFragmentVisitor) {
	visitor.VisitEndBodyFragment(f)
}
