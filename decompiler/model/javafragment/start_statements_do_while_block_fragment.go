package javafragment

import intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"

func NewStartStatementsDoWhileBlockFragment(minimalLineCount, lineCount, maximalLineCount,
	weight int, label string) intmod.IStartStatementsDoWhileBlockFragment {
	return NewStartStatementsDoWhileBlockFragmentWithGroup(minimalLineCount, lineCount,
		maximalLineCount, weight, label, NewStartStatementsBlockFragmentGroup())
}

func NewStartStatementsDoWhileBlockFragmentWithGroup(minimalLineCount, lineCount, maximalLineCount, weight int,
	label string, group intmod.IStartStatementsBlockFragmentGroup) intmod.IStartStatementsDoWhileBlockFragment {
	return &StartStatementsDoWhileBlockFragment{
		StartStatementsBlockFragment: *NewStartStatementsBlockFragmentWithGroup(minimalLineCount,
			lineCount, maximalLineCount, weight, label, group).(*StartStatementsBlockFragment),
	}
}

type StartStatementsDoWhileBlockFragment struct {
	StartStatementsBlockFragment
}

func (f *StartStatementsDoWhileBlockFragment) Accept(visitor intmod.IJavaFragmentVisitor) {
	visitor.VisitStartStatementsDoWhileBlockFragment(f)
}
