package javafragment

func NewStartStatementsTryBlockFragment(minimalLineCount int, lineCount int, maximalLineCount int, weight int, label string) StartStatementsTryBlockFragment {
	return StartStatementsTryBlockFragment{
		StartStatementsBlockFragment: *NewStartStatementsBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight, label),
	}
}

func NewStartStatementsTryBlockFragmentWithGroup(minimalLineCount int, lineCount int, maximalLineCount int, weight int, label string, group *StartStatementsBlockFragmentGroup) StartStatementsTryBlockFragment {
	return StartStatementsTryBlockFragment{
		StartStatementsBlockFragment: *NewStartStatementsBlockFragmentWithGroup(minimalLineCount, lineCount, maximalLineCount, weight, label, group),
	}
}

type StartStatementsTryBlockFragment struct {
	StartStatementsBlockFragment
}

func (f *StartStatementsTryBlockFragment) Accept(visitor JavaFragmentVisitor) {
	visitor.VisitStartStatementsTryBlockFragment(f)
}
