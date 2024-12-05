package javafragment

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/model/fragment"
)

func NewEndBodyFragment(minimalLineCount, lineCount, maximalLineCount, weight int,
	label string, start intmod.IStartBodyFragment) intmod.IEndBodyFragment {
	f := &EndBodyFragment{
		EndFlexibleBlockFragment: *fragment.NewEndFlexibleBlockFragment(minimalLineCount,
			lineCount, maximalLineCount, weight, label).(*fragment.EndFlexibleBlockFragment),
		start: start,
	}

	f.start.SetEndBodyFragment(f)

	return f
}

type EndBodyFragment struct {
	fragment.EndFlexibleBlockFragment

	start intmod.IStartBodyFragment
}

func (f *EndBodyFragment) Start() intmod.IStartBlockFragment {
	return f.start
}

func (f *EndBodyFragment) SetStart(start intmod.IStartBlockFragment) {
	f.start = start
}

func (f *EndBodyFragment) IncLineCount(force bool) bool {
	if f.LineCount() < f.MaximalLineCount() {
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
	if f.LineCount() > f.MinimalLineCount() {
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

func (f *EndBodyFragment) Accept(visitor intmod.IJavaFragmentVisitor) {
	visitor.VisitEndBodyFragment(f)
}
