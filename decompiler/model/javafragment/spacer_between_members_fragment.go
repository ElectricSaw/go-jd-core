package javafragment

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/fragment"
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
