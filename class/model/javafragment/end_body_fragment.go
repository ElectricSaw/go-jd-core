package javafragment

import "bitbucket.org/coontec/go-jd-core/class/model/fragment"

func NewEndBodyFragment(minimalLineCount int, lineCount int, maximalLineCount int, weight int, label string, start *StartBodyFragment) *EndBodyFragment {
	f := &EndBodyFragment{
		EndFlexibleBlockFragment: *fragment.NewEndFlexibleBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight, label),
		start:                    start,
	}

	f.start.SetEndBodyFragment(f)

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
	if f.LineCount < f.MaximalLineCount {
		f.LineCount++

		if !force {
			if f.LineCount == 1 && f.start.LineCount == 0 {
				f.start.LineCount = f.LineCount
			}
		}

		return true
	}
	return false
}

func (f *EndBodyFragment) DecLineCount(force bool) bool {
	if f.LineCount > f.MinimalLineCount {
		f.LineCount--

		if !force {
			if f.LineCount == 0 {
				f.start.LineCount = f.LineCount
			}
		}

		return true
	}
	return false
}

func (f *EndBodyFragment) Accept(visitor JavaFragmentVisitor) {
	visitor.VisitEndBodyFragment(f)
}
