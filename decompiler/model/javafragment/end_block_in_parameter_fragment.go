package javafragment

func NewEndBlockInParameterFragment(minimalLineCount, lineCount, maximalLineCount,
	weight int, label string, start *StartBlockFragment) EndBlockInParameterFragment {
	return EndBlockInParameterFragment{
		EndBlockFragment: NewEndBlockFragment(minimalLineCount, lineCount,
			maximalLineCount, weight, label, start),
	}
}

type EndBlockInParameterFragment struct {
	EndBlockFragment
}

func (f *EndBlockInParameterFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitEndBlockInParameterFragment(f)
}
