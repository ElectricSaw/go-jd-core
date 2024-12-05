package fragment

import intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"

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
