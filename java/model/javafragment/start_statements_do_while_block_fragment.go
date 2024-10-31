package javafragment

func NewStartStatementsDoWhileBlockFragment(minimalLineCount int, lineCount int, maximalLineCount int, weight int, label string) *StartStatementsDoWhileBlockFragment {
	return &StartStatementsDoWhileBlockFragment{
		StartStatementsBlockFragment: *NewStartStatementsBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight, label),
	}
}

func NewStartStatementsDoWhileBlockFragmentWithGroup(minimalLineCount int, lineCount int, maximalLineCount int, weight int, label string, group *StartStatementsBlockFragmentGroup) *StartStatementsDoWhileBlockFragment {
	return &StartStatementsDoWhileBlockFragment{
		StartStatementsBlockFragment: *NewStartStatementsBlockFragmentWithGroup(minimalLineCount, lineCount, maximalLineCount, weight, label, group),
	}
}

type StartStatementsDoWhileBlockFragment struct {
	StartStatementsBlockFragment
}

func (f *StartStatementsDoWhileBlockFragment) Accept(visitor JavaFragmentVisitor) {
	visitor.VisitStartStatementsDoWhileBlockFragment(f)
}
