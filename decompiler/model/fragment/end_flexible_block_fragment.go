package fragment

import intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"

func NewEndFlexibleBlockFragment(minimalLineCount, lineCount, maximalLineCount,
	weight int, label string) intmod.IEndFlexibleBlockFragment {
	return &EndFlexibleBlockFragment{
		FlexibleFragment: *NewFlexibleFragment(minimalLineCount, lineCount,
			maximalLineCount, weight, label).(*FlexibleFragment),
	}
}

type EndFlexibleBlockFragment struct {
	FlexibleFragment
}

func (f *EndFlexibleBlockFragment) AcceptFragmentVisitor(visitor intmod.IFragmentVisitor) {
	visitor.VisitEndFlexibleBlockFragment(f)
}
