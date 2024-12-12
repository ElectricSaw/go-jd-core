package javafragment

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/class/model/fragment"
)

func NewEndSingleStatementBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight int,
	label string, start intmod.IStartSingleStatementBlockFragment) intmod.IEndSingleStatementBlockFragment {
	f := &EndSingleStatementBlockFragment{
		EndFlexibleBlockFragment: *fragment.NewEndFlexibleBlockFragment(minimalLineCount,
			lineCount, maximalLineCount, weight, label).(*fragment.EndFlexibleBlockFragment),
		start: start,
	}

	f.start.SetEndSingleStatementBlockFragment(f)

	return f
}

type EndSingleStatementBlockFragment struct {
	fragment.EndFlexibleBlockFragment

	start intmod.IStartSingleStatementBlockFragment
}

func (f *EndSingleStatementBlockFragment) Start() intmod.IStartSingleStatementBlockFragment {
	return f.start
}

func (f *EndSingleStatementBlockFragment) SetStart(start intmod.IStartSingleStatementBlockFragment) {
	f.start = start
}

func (f *EndSingleStatementBlockFragment) IncLineCount(force bool) bool {
	if f.LineCount() < f.MaximalLineCount() {
		f.SetLineCount(f.LineCount() + 1)

		if !force {
			if f.start.LineCount() == 0 {
				f.start.SetLineCount(1)
			}
		}

		return true
	}
	return false
}

func (f *EndSingleStatementBlockFragment) DecLineCount(force bool) bool {
	if f.LineCount() > f.MinimalLineCount() {
		f.SetLineCount(f.LineCount() - 1)

		if !force {
			if f.LineCount() == 0 {
				f.start.SetLineCount(0)
			}
		}

		return true
	}
	return false
}

func (f *EndSingleStatementBlockFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitEndSingleStatementBlockFragment(f)
}
