package fragment

func NewEndFlexibleBlockFragment(minimalLineCount int, lineCount int, maximalLineCount int, weight int, label string) EndFlexibleBlockFragment {
	return EndFlexibleBlockFragment{
		FlexibleFragment: NewFlexibleFragment(minimalLineCount, lineCount, maximalLineCount, weight, label),
	}
}

type EndFlexibleBlockFragment struct {
	FlexibleFragment
}

func (f *EndFlexibleBlockFragment) Accept(visitor FragmentVisitor) {
	visitor.VisitEndFlexibleBlockFragment(f)
}
