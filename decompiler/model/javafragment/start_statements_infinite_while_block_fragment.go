package javafragment

func NewStartStatementsInfiniteWhileBlockFragment(minimalLineCount, lineCount, maximalLineCount,
	weight int, label string) StartStatementsInfiniteWhileBlockFragment {
	return NewStartStatementsInfiniteWhileBlockFragmentWithGroup(minimalLineCount, lineCount,
		maximalLineCount, weight, label, NewStartStatementsBlockFragmentGroup())
}

func NewStartStatementsInfiniteWhileBlockFragmentWithGroup(minimalLineCount, lineCount, maximalLineCount, weight int,
	label string, group StartStatementsBlockFragmentGroup) StartStatementsInfiniteWhileBlockFragment {
	return StartStatementsInfiniteWhileBlockFragment{
		StartStatementsBlockFragment: NewStartStatementsBlockFragmentWithGroup(minimalLineCount,
			lineCount, maximalLineCount, weight, label, group),
	}
}

type StartStatementsInfiniteWhileBlockFragment struct {
	StartStatementsBlockFragment
}

func (f *StartStatementsInfiniteWhileBlockFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitStartStatementsInfiniteWhileBlockFragment(f)
}
