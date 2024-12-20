package javafragment

import intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"

func NewStartStatementsInfiniteWhileBlockFragment(minimalLineCount, lineCount, maximalLineCount,
	weight int, label string) intmod.IStartStatementsInfiniteWhileBlockFragment {
	return NewStartStatementsInfiniteWhileBlockFragmentWithGroup(minimalLineCount, lineCount,
		maximalLineCount, weight, label, NewStartStatementsBlockFragmentGroup())
}

func NewStartStatementsInfiniteWhileBlockFragmentWithGroup(minimalLineCount, lineCount, maximalLineCount, weight int,
	label string, group intmod.IStartStatementsBlockFragmentGroup) intmod.IStartStatementsInfiniteWhileBlockFragment {
	return &StartStatementsInfiniteWhileBlockFragment{
		StartStatementsBlockFragment: *NewStartStatementsBlockFragmentWithGroup(minimalLineCount,
			lineCount, maximalLineCount, weight, label, group).(*StartStatementsBlockFragment),
	}
}

type StartStatementsInfiniteWhileBlockFragment struct {
	StartStatementsBlockFragment
}

func (f *StartStatementsInfiniteWhileBlockFragment) Accept(visitor intmod.IJavaFragmentVisitor) {
	visitor.VisitStartStatementsInfiniteWhileBlockFragment(f)
}
