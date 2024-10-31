package javafragment

func NewStartStatementsInfiniteWhileBlockFragment(minimalLineCount int, lineCount int, maximalLineCount int, weight int, label string) *StartStatementsInfiniteWhileBlockFragment {
	return &StartStatementsInfiniteWhileBlockFragment{
		StartStatementsBlockFragment: *NewStartStatementsBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight, label),
	}
}

func NewStartStatementsInfiniteWhileBlockFragmentWithGroup(minimalLineCount int, lineCount int, maximalLineCount int, weight int, label string, group *StartStatementsBlockFragmentGroup) *StartStatementsInfiniteWhileBlockFragment {
	return &StartStatementsInfiniteWhileBlockFragment{
		StartStatementsBlockFragment: *NewStartStatementsBlockFragmentWithGroup(minimalLineCount, lineCount, maximalLineCount, weight, label, group),
	}
}

type StartStatementsInfiniteWhileBlockFragment struct {
	StartStatementsBlockFragment
}

func (f *StartStatementsInfiniteWhileBlockFragment) Accept(visitor JavaFragmentVisitor) {
	visitor.VisitStartStatementsInfiniteWhileBlockFragment(f)
}
