package javafragment

func NewStartStatementsDoWhileBlockFragment(minimalLineCount, lineCount, maximalLineCount,
	weight int, label string) StartStatementsDoWhileBlockFragment {
	return NewStartStatementsDoWhileBlockFragmentWithGroup(minimalLineCount, lineCount,
		maximalLineCount, weight, label, NewStartStatementsBlockFragmentGroup())
}

func NewStartStatementsDoWhileBlockFragmentWithGroup(minimalLineCount, lineCount, maximalLineCount, weight int,
	label string, group StartStatementsBlockFragmentGroup) StartStatementsDoWhileBlockFragment {
	return StartStatementsDoWhileBlockFragment{
		StartStatementsBlockFragment: NewStartStatementsBlockFragmentWithGroup(minimalLineCount,
			lineCount, maximalLineCount, weight, label, group),
	}
}

type StartStatementsDoWhileBlockFragment struct {
	StartStatementsBlockFragment
}

func (f *StartStatementsDoWhileBlockFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitStartStatementsDoWhileBlockFragment(f)
}
