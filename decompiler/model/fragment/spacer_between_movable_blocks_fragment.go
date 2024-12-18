package fragment

import intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"

func NewSpacerBetweenMovableBlocksFragment(minimalLineCount, lineCount, maximalLineCount, weight int, label string) intmod.ISpacerBetweenMovableBlocksFragment {
	return &SpacerBetweenMovableBlocksFragment{
		FlexibleFragment: *NewFlexibleFragment(minimalLineCount, lineCount,
			maximalLineCount, weight, label).(*FlexibleFragment),
	}
}

type SpacerBetweenMovableBlocksFragment struct {
	FlexibleFragment
}

func (f *SpacerBetweenMovableBlocksFragment) SetInitialLineCount(initialLineCount int) {
	f.initialLineCount = initialLineCount
	f.lineCount = initialLineCount
}

func (f *SpacerBetweenMovableBlocksFragment) AcceptFragmentVisitor(visitor intmod.IFragmentVisitor) {
	visitor.VisitSpacerBetweenMovableBlocksFragment(f)
}
