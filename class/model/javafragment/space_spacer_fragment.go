package javafragment

import intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"

func NewSpaceSpacerFragment(minimalLineCount, lineCount, maximalLineCount, weight int, label string) intmod.ISpaceSpacerFragment {
	return &SpaceSpacerFragment{
		SpacerFragment: *NewSpacerFragment(minimalLineCount, lineCount, maximalLineCount, weight, label).(*SpacerFragment),
	}
}

type SpaceSpacerFragment struct {
	SpacerFragment
}

func (f *SpaceSpacerFragment) Accept(visitor intmod.IJavaFragmentVisitor) {
	visitor.VisitSpaceSpacerFragment(f)
}
