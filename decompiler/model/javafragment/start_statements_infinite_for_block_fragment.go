package javafragment

func NewStartStatementsInfiniteForBlockFragment(minimalLineCount, lineCount, maximalLineCount,
	weight int, label string) StartStatementsInfiniteForBlockFragment {
	return NewStartStatementsInfiniteForBlockFragmentWithGroup(minimalLineCount, lineCount,
		maximalLineCount, weight, label, NewStartStatementsBlockFragmentGroup())
}

func NewStartStatementsInfiniteForBlockFragmentWithGroup(minimalLineCount, lineCount, maximalLineCount, weight int,
	label string, group StartStatementsBlockFragmentGroup) StartStatementsInfiniteForBlockFragment {
	return StartStatementsInfiniteForBlockFragment{
		StartStatementsBlockFragment: NewStartStatementsBlockFragmentWithGroup(minimalLineCount,
			lineCount, maximalLineCount, weight, label, group),
	}
}

type StartStatementsInfiniteForBlockFragment struct {
	StartStatementsBlockFragment
}

func (f *StartStatementsInfiniteForBlockFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitStartStatementsInfiniteForBlockFragment(f)
}
