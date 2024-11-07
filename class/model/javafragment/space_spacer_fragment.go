package javafragment

func NewSpaceSpacerFragment(minimalLineCount, lineCount, maximalLineCount, weight int, label string) *SpaceSpacerFragment {
	return &SpaceSpacerFragment{
		SpacerFragment: *NewSpacerFragment(minimalLineCount, lineCount, maximalLineCount, weight, label),
	}
}

type SpaceSpacerFragment struct {
	SpacerFragment
}

func (f *SpaceSpacerFragment) Accept(visitor JavaFragmentVisitor) {
	visitor.VisitSpaceSpacerFragment(f)
}
