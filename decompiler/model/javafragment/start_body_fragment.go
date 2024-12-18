package javafragment

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/fragment"
)

func NewStartBodyFragment(minimalLineCount, lineCount, maximalLineCount, weight int,
	label string) intmod.IStartBodyFragment {
	return &StartBodyFragment{
		StartFlexibleBlockFragment: *fragment.NewStartFlexibleBlockFragment(minimalLineCount,
			lineCount, maximalLineCount, weight, label).(*fragment.StartFlexibleBlockFragment),
	}
}

type StartBodyFragment struct {
	fragment.StartFlexibleBlockFragment

	end intmod.IEndBodyFragment
}

func (f *StartBodyFragment) End() intmod.IEndBodyFragment {
	return f.end
}

func (f *StartBodyFragment) SetEnd(end intmod.IEndBodyFragment) {
	f.end = end
}

func (f *StartBodyFragment) IncLineCount(force bool) bool {
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

func (f *StartBodyFragment) DecLineCount(force bool) bool {
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

func (f *StartBodyFragment) Accept(visitor intmod.IJavaFragmentVisitor) {
	visitor.VisitStartBodyFragment(f)
}
