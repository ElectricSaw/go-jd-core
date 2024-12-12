package javafragment

import intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"

func NewStartStatementsTryBlockFragment(minimalLineCount, lineCount, maximalLineCount,
	weight int, label string) intmod.IStartStatementsTryBlockFragment {
	return NewStartStatementsTryBlockFragmentWithGroup(minimalLineCount, lineCount,
		maximalLineCount, weight, label, NewStartStatementsBlockFragmentGroup())
}

func NewStartStatementsTryBlockFragmentWithGroup(minimalLineCount, lineCount, maximalLineCount, weight int,
	label string, group intmod.IStartStatementsBlockFragmentGroup) intmod.IStartStatementsTryBlockFragment {
	return &StartStatementsTryBlockFragment{
		StartStatementsBlockFragment: *NewStartStatementsBlockFragmentWithGroup(minimalLineCount,
			lineCount, maximalLineCount, weight, label, group).(*StartStatementsBlockFragment),
	}
}

type StartStatementsTryBlockFragment struct {
	StartStatementsBlockFragment
}

func (f *StartStatementsTryBlockFragment) Accept(visitor intmod.IJavaFragmentVisitor) {
	visitor.VisitStartStatementsTryBlockFragment(f)
}
