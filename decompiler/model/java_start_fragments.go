package model

import (
	"fmt"
)

var StartMovableTypeBlock = NewStartMovableJavaBlockFragment(1)
var StartMovableFieldBlock = NewStartMovableJavaBlockFragment(2)
var StartMovableMethodBlock = NewStartMovableJavaBlockFragment(3)

func NewStartBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight int,
	label string) StartBlockFragment {
	return StartBlockFragment{
		StartFlexibleBlockFragment: NewStartFlexibleBlockFragment(minimalLineCount,
			lineCount, maximalLineCount, weight, label),
	}
}

func NewStartBodyFragment(minimalLineCount, lineCount, maximalLineCount, weight int,
	label string) StartBodyFragment {
	return StartBodyFragment{
		StartFlexibleBlockFragment: NewStartFlexibleBlockFragment(minimalLineCount,
			lineCount, maximalLineCount, weight, label),
	}
}

func NewStartMovableJavaBlockFragment(typ int) StartMovableJavaBlockFragment {
	return StartMovableJavaBlockFragment{
		StartMovableBlockFragment: NewStartMovableBlockFragment(typ),
	}
}

func NewStartSingleStatementBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight int, label string) StartSingleStatementBlockFragment {
	return StartSingleStatementBlockFragment{
		StartFlexibleBlockFragment: NewStartFlexibleBlockFragment(minimalLineCount, lineCount,
			maximalLineCount, weight, label),
	}
}

func NewStartStatementsBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight int,
	label string) StartStatementsBlockFragment {
	return NewStartStatementsBlockFragmentWithGroup(minimalLineCount, lineCount,
		maximalLineCount, weight, label, NewStartStatementsBlockFragmentGroup())
}

func NewStartStatementsBlockFragmentWithGroup(minimalLineCount, lineCount, maximalLineCount, weight int,
	label string, group StartStatementsBlockFragmentGroup) StartStatementsBlockFragment {
	f := StartStatementsBlockFragment{
		StartFlexibleBlockFragment: NewStartFlexibleBlockFragment(minimalLineCount, lineCount,
			maximalLineCount, weight, label),
		Group: group,
	}
	f.Group.Add(&f)
	return f
}

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

func NewStartStatementsInfiniteForBlockFragment(minimalLineCount, lineCount, maximalLineCount,
	weight int, label string) StartStatementsInfiniteForBlockFragment {
	return NewStartStatementsInfiniteForBlockFragmentWithGroup(minimalLineCount, lineCount,
		maximalLineCount, weight, label, NewStartStatementsBlockFragmentGroup())
}

func NewStartStatementsInfiniteForBlockFragmentWithGroup(minimalLineCount, lineCount, maximalLineCount, weight int,
	label string, group StartStatementsBlockFragmentGroup) StartStatementsInfiniteForBlockFragment {
	return StartStatementsInfiniteForBlockFragment{
		StartStatementsBlockFragment: NewStartStatementsBlockFragmentWithGroup(minimalLineCount,
			lineCount, maximalLineCount, weight, label, group),
	}
}

func NewStartStatementsInfiniteWhileBlockFragment(minimalLineCount, lineCount, maximalLineCount,
	weight int, label string) StartStatementsInfiniteWhileBlockFragment {
	return NewStartStatementsInfiniteWhileBlockFragmentWithGroup(minimalLineCount, lineCount,
		maximalLineCount, weight, label, NewStartStatementsBlockFragmentGroup())
}

func NewStartStatementsInfiniteWhileBlockFragmentWithGroup(minimalLineCount, lineCount, maximalLineCount, weight int,
	label string, group StartStatementsBlockFragmentGroup) StartStatementsInfiniteWhileBlockFragment {
	return StartStatementsInfiniteWhileBlockFragment{
		StartStatementsBlockFragment: NewStartStatementsBlockFragmentWithGroup(minimalLineCount,
			lineCount, maximalLineCount, weight, label, group),
	}
}

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

type StartBlockFragment struct {
	StartFlexibleBlockFragment

	End *EndBlockFragment
}

func (f *StartBlockFragment) IncLineCount(force bool) bool {
	if f.LineCount < f.MaximalLineCount {
		f.LineCount = f.LineCount + 1

		if !force {
			if f.LineCount == 1 && f.End.LineCount == 0 {
				f.End.LineCount = f.LineCount
			}
		}

		return true
	}
	return false
}

func (f *StartBlockFragment) DecLineCount(force bool) bool {
	if f.LineCount > f.MinimalLineCount {
		f.LineCount = f.LineCount - 1

		if !force {
			if f.LineCount == 1 {
				f.End.LineCount = f.LineCount
			}
		}

		return true
	}
	return false
}

func (f *StartBlockFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitStartBlockFragment(f)
}

func (f *StartBlockFragment) String() string {
	return fmt.Sprintf("StartBlockFragment { start: %s, end: %s", f.StartFlexibleBlockFragment.String(), f.End.String())
}

type StartBodyFragment struct {
	StartFlexibleBlockFragment

	End *EndBodyFragment
}

func (f *StartBodyFragment) IncLineCount(force bool) bool {
	if f.LineCount < f.MaximalLineCount {
		f.LineCount = f.LineCount + 1

		if !force {
			if f.LineCount == 1 && f.End.LineCount == 0 {
				f.End.LineCount = f.LineCount
			}
		}

		return true
	}
	return false
}

func (f *StartBodyFragment) DecLineCount(force bool) bool {
	if f.LineCount > f.MinimalLineCount {
		f.LineCount = f.LineCount - 1

		if !force {
			if f.LineCount == 1 {
				f.End.LineCount = f.LineCount
			}
		}

		return true
	}
	return false
}

func (f *StartBodyFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitStartBodyFragment(f)
}

type StartMovableJavaBlockFragment struct {
	StartMovableBlockFragment
}

func (f *StartMovableJavaBlockFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitStartMovableJavaBlockFragment(f)
}

type StartSingleStatementBlockFragment struct {
	StartFlexibleBlockFragment

	End *EndSingleStatementBlockFragment
}

func (f *StartSingleStatementBlockFragment) IncLineCount(force bool) bool {
	if f.LineCount < f.MaximalLineCount {
		f.LineCount = f.LineCount + 1

		if !force {
			if f.End.LineCount == 0 {
				f.End.LineCount = 1
			}
		}

		return true
	}
	return false
}

func (f *StartSingleStatementBlockFragment) DecLineCount(force bool) bool {
	if f.LineCount > f.MinimalLineCount {
		f.LineCount = f.LineCount - 1

		if !force {
			if f.LineCount == 1 {
				f.End.LineCount = 1
			}
		}

		return true
	}
	return false
}

func (f *StartSingleStatementBlockFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitStartSingleStatementBlockFragment(f)
}

type StartStatementsBlockFragment struct {
	StartFlexibleBlockFragment

	Group StartStatementsBlockFragmentGroup
}

func (f *StartStatementsBlockFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitStartStatementsBlockFragment(f)
}

type StartStatementsDoWhileBlockFragment struct {
	StartStatementsBlockFragment
}

func (f *StartStatementsDoWhileBlockFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitStartStatementsDoWhileBlockFragment(f)
}

type StartStatementsInfiniteForBlockFragment struct {
	StartStatementsBlockFragment
}

func (f *StartStatementsInfiniteForBlockFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitStartStatementsInfiniteForBlockFragment(f)
}

type StartStatementsInfiniteWhileBlockFragment struct {
	StartStatementsBlockFragment
}

func (f *StartStatementsInfiniteWhileBlockFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitStartStatementsInfiniteWhileBlockFragment(f)
}

type StartStatementsTryBlockFragment struct {
	StartStatementsBlockFragment
}

func (f *StartStatementsTryBlockFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitStartStatementsTryBlockFragment(f)
}
