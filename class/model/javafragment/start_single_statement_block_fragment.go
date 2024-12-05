package javafragment

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/model/fragment"
)

func NewStartSingleStatementBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight int, label string) intmod.IStartSingleStatementBlockFragment {
	return &StartSingleStatementBlockFragment{
		StartFlexibleBlockFragment: *fragment.NewStartFlexibleBlockFragment(minimalLineCount, lineCount,
			maximalLineCount, weight, label).(*fragment.StartFlexibleBlockFragment),
	}
}

type StartSingleStatementBlockFragment struct {
	fragment.StartFlexibleBlockFragment

	end intmod.IEndSingleStatementBlockFragment
}

func (f *StartSingleStatementBlockFragment) End() intmod.IEndSingleStatementBlockFragment {
	return f.end
}

func (f *StartSingleStatementBlockFragment) SetEnd(end intmod.IEndSingleStatementBlockFragment) {
	f.end = end
}

func (f *StartSingleStatementBlockFragment) IncLineCount(force bool) bool {
	if f.LineCount() < f.MaximalLineCount() {
		f.SetLineCount(f.LineCount() + 1)

		if !force {
			if f.end.LineCount() == 0 {
				f.end.SetLineCount(1)
			}
		}

		return true
	}
	return false
}

func (f *StartSingleStatementBlockFragment) DecLineCount(force bool) bool {
	if f.LineCount() > f.MinimalLineCount() {
		f.SetLineCount(f.LineCount() - 1)

		if !force {
			if f.LineCount() == 1 {
				f.end.SetLineCount(1)
			}
		}

		return true
	}
	return false
}

func (f *StartSingleStatementBlockFragment) Accept(visitor intmod.IJavaFragmentVisitor) {
	visitor.VisitStartSingleStatementBlockFragment(f)
}
