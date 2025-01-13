package javafragment

func NewEndBodyInParameterFragment(minimalLineCount, lineCount, maximalLineCount,
	weight int, label string, start *StartBodyFragment) EndBodyInParameterFragment {
	return EndBodyInParameterFragment{
		EndBodyFragment: NewEndBodyFragment(minimalLineCount, lineCount,
			maximalLineCount, weight, label, start),
	}
}

type EndBodyInParameterFragment struct {
	EndBodyFragment
}

func (f *EndBodyInParameterFragment) IncLineCount(force bool) bool {
	if f.LineCount < f.MaximalLineCount {
		f.LineCount = f.LineCount + 1
		return true
	}
	return false
}

func (f *EndBodyInParameterFragment) DecLineCount(force bool) bool {
	if f.LineCount > f.MinimalLineCount {
		f.LineCount = f.LineCount - 1
		return true
	}
	return false
}

func (f *EndBodyInParameterFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitEndBodyInParameterFragment(f)
}
