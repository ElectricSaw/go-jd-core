package fragment

import intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"

func NewStartFlexibleBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight int, label string) intmod.IStartFlexibleBlockFragment {
	return &StartFlexibleBlockFragment{
		FlexibleFragment: *NewFlexibleFragment(minimalLineCount, lineCount,
			maximalLineCount, weight, label).(*FlexibleFragment),
	}
}

type StartFlexibleBlockFragment struct {
	FlexibleFragment
}

func (f *StartFlexibleBlockFragment) AcceptFragmentVisitor(visitor intmod.IFragmentVisitor) {
	visitor.VisitStartFlexibleBlockFragment(f)
}
