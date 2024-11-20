package javafragment

import "bitbucket.org/coontec/go-jd-core/class/model/fragment"

func NewStartBodyFragment(minimalLineCount int, lineCount int, maximalLineCount int, weight int, label string) *StartBodyFragment {
	return &StartBodyFragment{
		StartFlexibleBlockFragment: *fragment.NewStartFlexibleBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight, label),
	}
}

type StartBodyFragment struct {
	fragment.StartFlexibleBlockFragment

	end *EndBodyFragment
}

func (f *StartBodyFragment) EndBodyFragment() *EndBodyFragment {
	return f.end
}

func (f *StartBodyFragment) SetEndBodyFragment(end *EndBodyFragment) {
	f.end = end
}

func (f *StartBodyFragment) IncLineCount(force bool) bool {
	if f.LineCount < f.MaximalLineCount {
		f.LineCount++

		if !force {
			if f.LineCount == 1 && f.end.LineCount == 0 {
				f.end.LineCount = f.LineCount
			}
		}

		return true
	}
	return false
}

func (f *StartBodyFragment) DecLineCount(force bool) bool {
	if f.LineCount > f.MinimalLineCount {
		f.LineCount--

		if !force {
			if f.LineCount == 1 {
				f.end.LineCount = f.LineCount
			}
		}

		return true
	}
	return false
}

func (f *StartBodyFragment) Accept(visitor JavaFragmentVisitor) {
	visitor.VisitStartBodyFragment(f)
}
