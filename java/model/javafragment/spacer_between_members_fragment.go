package javafragment

import "bitbucket.org/coontec/javaClass/java/model/fragment"

func NewSpacerBetweenMembersFragment(minimalLineCount, lineCount, maximalLineCount, weight int, label string) *SpacerBetweenMembersFragment {
	return &SpacerBetweenMembersFragment{
		SpacerBetweenMovableBlocksFragment: *fragment.NewSpacerBetweenMovableBlocksFragment(minimalLineCount, lineCount, maximalLineCount, weight, label),
	}
}

type SpacerBetweenMembersFragment struct {
	fragment.SpacerBetweenMovableBlocksFragment
}

func (f *SpacerBetweenMembersFragment) Accept(visitor JavaFragmentVisitor) {
	visitor.VisitSpacerBetweenMembersFragment(f)
}
