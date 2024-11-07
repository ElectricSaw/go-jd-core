package fragment

func NewSpacerBetweenMovableBlocksFragment(minimalLineCount, lineCount, maximalLineCount, weight int, label string) *SpacerBetweenMovableBlocksFragment {
	return &SpacerBetweenMovableBlocksFragment{
		FlexibleFragment: *NewFlexibleFragment(minimalLineCount, lineCount, maximalLineCount, weight, label),
	}
}

type SpacerBetweenMovableBlocksFragment struct {
	FlexibleFragment
}

func (f *SpacerBetweenMovableBlocksFragment) SetInitialLineCount(initialLineCount int) {
	f.InitialLineCount = initialLineCount
	f.LineCount = initialLineCount
}

func (f *SpacerBetweenMovableBlocksFragment) Accept(visitor FragmentVisitor) {
	visitor.VisitSpacerBetweenMovableBlocksFragment(f)
}
