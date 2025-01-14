package model

func NewSpaceSpacerFragment(minimalLineCount, lineCount, maximalLineCount, weight int, label string) SpaceSpacerFragment {
	return SpaceSpacerFragment{
		SpacerFragment: NewSpacerFragment(minimalLineCount, lineCount, maximalLineCount, weight, label),
	}
}

func NewSpacerBetweenMembersFragment(minimalLineCount, lineCount, maximalLineCount, weight int, label string) SpacerBetweenMembersFragment {
	return SpacerBetweenMembersFragment{
		SpacerBetweenMovableBlocksFragment: NewSpacerBetweenMovableBlocksFragment(
			minimalLineCount, lineCount, maximalLineCount, weight, label),
	}
}

func NewSpacerFragment(minimalLineCount, lineCount, maximalLineCount, weight int, label string) SpacerFragment {
	return SpacerFragment{
		FlexibleFragment: NewFlexibleFragment(minimalLineCount, lineCount,
			maximalLineCount, weight, label),
	}
}

type SpaceSpacerFragment struct {
	SpacerFragment
}

func (f *SpaceSpacerFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitSpaceSpacerFragment(f)
}

type SpacerBetweenMembersFragment struct {
	SpacerBetweenMovableBlocksFragment
}

func (f *SpacerBetweenMembersFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitSpacerBetweenMembersFragment(f)
}

type SpacerFragment struct {
	FlexibleFragment
}

func (f *SpacerFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitSpacerFragment(f)
}
