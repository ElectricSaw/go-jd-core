package javafragment

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/class/model/fragment"
)

func NewEndBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight int, label string, start intmod.IStartBlockFragment) intmod.IEndBlockFragment {
	f := &EndBlockFragment{
		EndFlexibleBlockFragment: *fragment.NewEndFlexibleBlockFragment(minimalLineCount,
			lineCount, maximalLineCount, weight, label).(*fragment.EndFlexibleBlockFragment),
		start: start,
	}

	f.start.SetEnd(f)

	return f
}

type EndBlockFragment struct {
	fragment.EndFlexibleBlockFragment

	start intmod.IStartBlockFragment
}

func (f *EndBlockFragment) Start() intmod.IStartBlockFragment {
	return f.start
}

func (f *EndBlockFragment) SetStart(start intmod.IStartBlockFragment) {
	f.start = start
}

func (f *EndBlockFragment) IncLineCount(force bool) bool {
	if f.LineCount() < f.MaximalLineCount() {
		f.SetLineCount(f.LineCount() + 1)
		return true
	}
	return false
}

func (f *EndBlockFragment) DecLineCount(force bool) bool {
	if f.LineCount() > f.MinimalLineCount() {
		f.SetLineCount(f.LineCount() - 1)
		return true
	}
	return false
}

func (f *EndBlockFragment) Accept(visitor intmod.IJavaFragmentVisitor) {
	visitor.VisitEndBlockFragment(f)
}
