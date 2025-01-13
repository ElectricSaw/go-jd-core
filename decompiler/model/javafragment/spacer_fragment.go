package javafragment

func NewSpacerFragment(minimalLineCount, lineCount, maximalLineCount, weight int, label string) SpacerFragment {
	return SpacerFragment{
		FlexibleFragment: NewFlexibleFragment(minimalLineCount, lineCount,
			maximalLineCount, weight, label),
	}
}

type SpacerFragment struct {
	FlexibleFragment
}

func (f *SpacerFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitSpacerFragment(f)
}
