package javafragment

func NewEndStatementsBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight int,
	label string, group StartStatementsBlockFragmentGroup) EndStatementsBlockFragment {
	f := EndStatementsBlockFragment{
		EndFlexibleBlockFragment: NewEndFlexibleBlockFragment(minimalLineCount,
			lineCount, maximalLineCount, weight, label),
		Group: group,
	}

	f.Group.Add(&f)

	return f
}

type EndStatementsBlockFragment struct {
	EndFlexibleBlockFragment

	Group StartStatementsBlockFragmentGroup
}

func (f *EndStatementsBlockFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitEndStatementsBlockFragment(f)
}
