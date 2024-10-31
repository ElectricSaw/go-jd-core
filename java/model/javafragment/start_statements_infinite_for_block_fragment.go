package javafragment

func NewStartStatementsInfiniteForBlockFragment(minimalLineCount int, lineCount int, maximalLineCount int, weight int, label string) *StartStatementsInfiniteForBlockFragment {
	return &StartStatementsInfiniteForBlockFragment{
		StartStatementsBlockFragment: *NewStartStatementsBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight, label),
	}
}

func NewStartStatementsInfiniteForBlockFragmentWithGroup(minimalLineCount int, lineCount int, maximalLineCount int, weight int, label string, group *StartStatementsBlockFragmentGroup) *StartStatementsInfiniteForBlockFragment {
	return &StartStatementsInfiniteForBlockFragment{
		StartStatementsBlockFragment: *NewStartStatementsBlockFragmentWithGroup(minimalLineCount, lineCount, maximalLineCount, weight, label, group),
	}
}

type StartStatementsInfiniteForBlockFragment struct {
	StartStatementsBlockFragment
}

func (f *StartStatementsInfiniteForBlockFragment) Accept(visitor JavaFragmentVisitor) {
	visitor.VisitStartStatementsInfiniteForBlockFragment(f)
}
