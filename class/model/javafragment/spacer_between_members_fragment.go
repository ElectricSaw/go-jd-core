package javafragment

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/model/fragment"
)

func NewSpacerBetweenMembersFragment(minimalLineCount, lineCount, maximalLineCount, weight int, label string) intmod.ISpacerBetweenMembersFragment {
	return &SpacerBetweenMembersFragment{
		SpacerBetweenMovableBlocksFragment: *fragment.NewSpacerBetweenMovableBlocksFragment(
			minimalLineCount, lineCount, maximalLineCount, weight, label).(*fragment.SpacerBetweenMovableBlocksFragment),
	}
}

type SpacerBetweenMembersFragment struct {
	fragment.SpacerBetweenMovableBlocksFragment
}

func (f *SpacerBetweenMembersFragment) Accept(visitor intmod.IJavaFragmentVisitor) {
	visitor.VisitSpacerBetweenMembersFragment(f)
}
