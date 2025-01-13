package javafragment

func NewStartStatementsTryBlockFragment(minimalLineCount, lineCount, maximalLineCount,
	weight int, label string) StartStatementsTryBlockFragment {
	return NewStartStatementsTryBlockFragmentWithGroup(minimalLineCount, lineCount,
		maximalLineCount, weight, label, NewStartStatementsBlockFragmentGroup())
}

func NewStartStatementsTryBlockFragmentWithGroup(minimalLineCount, lineCount, maximalLineCount, weight int,
	label string, group StartStatementsBlockFragmentGroup) StartStatementsTryBlockFragment {
	return StartStatementsTryBlockFragment{
		StartStatementsBlockFragment: NewStartStatementsBlockFragmentWithGroup(minimalLineCount,
			lineCount, maximalLineCount, weight, label, group),
	}
}

type StartStatementsTryBlockFragment struct {
	StartStatementsBlockFragment
}

func (f *StartStatementsTryBlockFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitStartStatementsTryBlockFragment(f)
}
