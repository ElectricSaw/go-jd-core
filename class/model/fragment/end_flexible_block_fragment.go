package fragment

import intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"

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
