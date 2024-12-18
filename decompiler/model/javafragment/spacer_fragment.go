package javafragment

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/fragment"
)

func NewSpacerFragment(minimalLineCount, lineCount, maximalLineCount, weight int, label string) intmod.ISpacerFragment {
	return &SpacerFragment{
		FlexibleFragment: *fragment.NewFlexibleFragment(minimalLineCount, lineCount,
			maximalLineCount, weight, label).(*fragment.FlexibleFragment),
	}
}

type SpacerFragment struct {
	fragment.FlexibleFragment
}

func (f *SpacerFragment) Accept(visitor intmod.IJavaFragmentVisitor) {
	visitor.VisitSpacerFragment(f)
}
