package javafragment

func NewEndBlockInParameterFragment(minimalLineCount int, lineCount int, maximalLineCount int, weight int, label string, start *StartBlockFragment) *EndBlockInParameterFragment {
	return &EndBlockInParameterFragment{
		EndBlockFragment: *NewEndBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight, label, start),
	}
}

type EndBlockInParameterFragment struct {
	EndBlockFragment
}

func (f *EndBlockInParameterFragment) Accept(visitor JavaFragmentVisitor) {
	visitor.VisitEndBlockInParameterFragment(f)
}
