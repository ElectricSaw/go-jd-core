package fragment

func NewStartFlexibleBlockFragment(minimalLineCount int, lineCount int, maximalLineCount int, weight int, label string) *StartFlexibleBlockFragment {
	return &StartFlexibleBlockFragment{
		FlexibleFragment: *NewFlexibleFragment(minimalLineCount, lineCount, maximalLineCount, weight, label),
	}
}

type StartFlexibleBlockFragment struct {
	FlexibleFragment
}

func (f *StartFlexibleBlockFragment) Accept(visitor FragmentVisitor) {
	visitor.VisitStartFlexibleBlockFragment(f)
}
