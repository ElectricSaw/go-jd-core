package javafragment

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/model/fragment"
)

func NewStartBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight int,
	label string) intmod.IStartBlockFragment {
	return &StartBlockFragment{
		StartFlexibleBlockFragment: *fragment.NewStartFlexibleBlockFragment(minimalLineCount,
			lineCount, maximalLineCount, weight, label).(*fragment.StartFlexibleBlockFragment),
	}
}

type StartBlockFragment struct {
	fragment.StartFlexibleBlockFragment

	end intmod.IEndBlockFragment
}

func (f *StartBlockFragment) End() intmod.IEndBlockFragment {
	return f.end
}

func (f *StartBlockFragment) SetEnd(end intmod.IEndBlockFragment) {
	f.end = end
}

func (f *StartBlockFragment) IncLineCount(force bool) bool {
	if f.LineCount() < f.MaximalLineCount() {
		f.SetLineCount(f.LineCount() + 1)

		if !force {
			if f.LineCount() == 1 && f.end.LineCount() == 0 {
				f.end.SetLineCount(f.LineCount())
			}
		}

		return true
	}
	return false
}

func (f *StartBlockFragment) DecLineCount(force bool) bool {
	if f.LineCount() > f.MinimalLineCount() {
		f.SetLineCount(f.LineCount() - 1)

		if !force {
			if f.LineCount() == 1 {
				f.end.SetLineCount(f.LineCount())
			}
		}

		return true
	}
	return false
}

func (f *StartBlockFragment) Accept(visitor intmod.IJavaFragmentVisitor) {
	visitor.VisitStartBlockFragment(f)
}
