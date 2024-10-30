package fragment

func NewSpacerBetweenMovableBlocksFragment(minimalLineCount int, lineCount int, maximalLineCount int, weight int, label string) SpacerBetweenMovableBlocksFragment {
	return SpacerBetweenMovableBlocksFragment{
		FlexibleFragment: NewFlexibleFragment(minimalLineCount, lineCount, maximalLineCount, weight, label),
	}
}

type SpacerBetweenMovableBlocksFragment struct {
	FlexibleFragment
}

func (f *SpacerBetweenMovableBlocksFragment) SetInitialLineCount(initialLineCount int) {
	f.InitialLineCount = initialLineCount
	f.lineCount = initialLineCount
}

func (f *SpacerBetweenMovableBlocksFragment) Accept(visitor FragmentVisitor) {
	visitor.VisitSpacerBetweenMovableBlocksFragment(f)
}
