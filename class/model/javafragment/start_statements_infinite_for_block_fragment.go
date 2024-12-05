package javafragment

import intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"

func NewStartStatementsInfiniteForBlockFragment(minimalLineCount, lineCount, maximalLineCount,
	weight int, label string) intmod.IStartStatementsInfiniteForBlockFragment {
	return NewStartStatementsInfiniteForBlockFragmentWithGroup(minimalLineCount, lineCount,
		maximalLineCount, weight, label, NewStartStatementsBlockFragmentGroup())
}

func NewStartStatementsInfiniteForBlockFragmentWithGroup(minimalLineCount, lineCount, maximalLineCount, weight int,
	label string, group intmod.IStartStatementsBlockFragmentGroup) intmod.IStartStatementsInfiniteForBlockFragment {
	return &StartStatementsInfiniteForBlockFragment{
		StartStatementsBlockFragment: *NewStartStatementsBlockFragmentWithGroup(minimalLineCount,
			lineCount, maximalLineCount, weight, label, group).(*StartStatementsBlockFragment),
	}
}

type StartStatementsInfiniteForBlockFragment struct {
	StartStatementsBlockFragment
}

func (f *StartStatementsInfiniteForBlockFragment) Accept(visitor intmod.IJavaFragmentVisitor) {
	visitor.VisitStartStatementsInfiniteForBlockFragment(f)
}
