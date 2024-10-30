package javafragment

import "bitbucket.org/coontec/javaClass/java/model/fragment"

func NewEndStatementsBlockFragment(minimalLineCount int, lineCount int, maximalLineCount int, weight int, label string, group *StartStatementsBlockFragmentGroup) EndStatementsBlockFragment {
	f := EndStatementsBlockFragment{
		EndFlexibleBlockFragment: fragment.NewEndFlexibleBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight, label),
		Group:                    group,
	}
	f.Group.add(&f)
	return f
}

type EndStatementsBlockFragment struct {
	fragment.EndFlexibleBlockFragment

	Group *StartStatementsBlockFragmentGroup
}

func (f *EndStatementsBlockFragment) Accept(visitor JavaFragmentVisitor) {
	visitor.VisitEndStatementsBlockFragment(f)
}
