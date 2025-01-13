package javafragment

func NewSpacerBetweenMembersFragment(minimalLineCount, lineCount, maximalLineCount, weight int, label string) SpacerBetweenMembersFragment {
	return SpacerBetweenMembersFragment{
		SpacerBetweenMovableBlocksFragment: NewSpacerBetweenMovableBlocksFragment(
			minimalLineCount, lineCount, maximalLineCount, weight, label),
	}
}

type SpacerBetweenMembersFragment struct {
	SpacerBetweenMovableBlocksFragment
}

func (f *SpacerBetweenMembersFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitSpacerBetweenMembersFragment(f)
}
