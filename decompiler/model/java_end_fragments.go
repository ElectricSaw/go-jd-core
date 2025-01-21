package model

import (
	"fmt"
)

func NewEndBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight int, label string, start *StartBlockFragment) EndBlockFragment {
	f := EndBlockFragment{
		EndFlexibleBlockFragment: NewEndFlexibleBlockFragment(minimalLineCount,
			lineCount, maximalLineCount, weight, label),
		Start: start,
	}

	f.Start.End = &f

	return f
}

func NewEndBlockInParameterFragment(minimalLineCount, lineCount, maximalLineCount,
	weight int, label string, start *StartBlockFragment) EndBlockInParameterFragment {
	return EndBlockInParameterFragment{
		EndBlockFragment: NewEndBlockFragment(minimalLineCount, lineCount,
			maximalLineCount, weight, label, start),
	}
}

func NewEndBodyFragment(minimalLineCount, lineCount, maximalLineCount, weight int,
	label string, start *StartBodyFragment) EndBodyFragment {
	f := EndBodyFragment{
		EndFlexibleBlockFragment: NewEndFlexibleBlockFragment(minimalLineCount,
			lineCount, maximalLineCount, weight, label),
		Start: start,
	}

	f.Start.End = &f

	return f
}

func NewEndBodyInParameterFragment(minimalLineCount, lineCount, maximalLineCount,
	weight int, label string, start *StartBodyFragment) EndBodyInParameterFragment {
	return EndBodyInParameterFragment{
		EndBodyFragment: NewEndBodyFragment(minimalLineCount, lineCount,
			maximalLineCount, weight, label, start),
	}
}

func NewEndMovableJavaBlockFragment() EndMovableJavaBlockFragment {
	return EndMovableJavaBlockFragment{
		EndMovableBlockFragment: NewEndMovableBlockFragment(),
	}
}

func NewEndSingleStatementBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight int,
	label string, start *StartSingleStatementBlockFragment) EndSingleStatementBlockFragment {
	f := EndSingleStatementBlockFragment{
		EndFlexibleBlockFragment: NewEndFlexibleBlockFragment(minimalLineCount,
			lineCount, maximalLineCount, weight, label),
		Start: start,
	}

	f.Start.End = &f

	return f
}

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

type IJavaFragment interface {
	Accept(visitor IJavaFragmentVisitor)
	String() string
}

type IJavaFragmentVisitor interface {
	VisitEndBodyFragment(fragment *EndBodyFragment)
	VisitEndBlockInParameterFragment(fragment *EndBlockInParameterFragment)
	VisitEndBlockFragment(fragment *EndBlockFragment)
	VisitEndBodyInParameterFragment(fragment *EndBodyInParameterFragment)
	VisitEndMovableJavaBlockFragment(fragment *EndMovableJavaBlockFragment)
	VisitEndSingleStatementBlockFragment(fragment *EndSingleStatementBlockFragment)
	VisitEndStatementsBlockFragment(fragment *EndStatementsBlockFragment)
	VisitImportsFragment(fragment *ImportsFragment)
	VisitLineNumberTokensFragment(fragment *LineNumberTokensFragment)
	VisitSpacerBetweenMembersFragment(fragment *SpacerBetweenMembersFragment)
	VisitSpacerFragment(fragment *SpacerFragment)
	VisitSpaceSpacerFragment(fragment *SpaceSpacerFragment)
	VisitStartBlockFragment(fragment *StartBlockFragment)
	VisitStartBodyFragment(fragment *StartBodyFragment)
	VisitStartMovableJavaBlockFragment(fragment *StartMovableJavaBlockFragment)
	VisitStartSingleStatementBlockFragment(fragment *StartSingleStatementBlockFragment)
	VisitStartStatementsBlockFragment(fragment *StartStatementsBlockFragment)
	VisitStartStatementsDoWhileBlockFragment(fragment *StartStatementsDoWhileBlockFragment)
	VisitStartStatementsInfiniteForBlockFragment(fragment *StartStatementsInfiniteForBlockFragment)
	VisitStartStatementsInfiniteWhileBlockFragment(fragment *StartStatementsInfiniteWhileBlockFragment)
	VisitStartStatementsTryBlockFragment(fragment *StartStatementsTryBlockFragment)
	VisitTokensFragment(fragment *TokensFragment)
}

type EndBlockFragment struct {
	EndFlexibleBlockFragment

	Start *StartBlockFragment
}

func (f *EndBlockFragment) IncLineCount(_ bool) bool {
	if f.LineCount < f.MaximalLineCount {
		f.LineCount = f.LineCount + 1
		return true
	}
	return false
}

func (f *EndBlockFragment) DecLineCount(_ bool) bool {
	if f.LineCount > f.MinimalLineCount {
		f.LineCount = f.LineCount - 1
		return true
	}
	return false
}

func (f *EndBlockFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitEndBlockFragment(f)
}

func (f *EndBlockFragment) String() string {
	return fmt.Sprintf("EndBlockFragment { start: %s, end: %s }", f.Start.String(), f.EndFlexibleBlockFragment.String())
}

type EndBlockInParameterFragment struct {
	EndBlockFragment
}

func (f *EndBlockInParameterFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitEndBlockInParameterFragment(f)
}

func (f *EndBlockInParameterFragment) String() string {
	return fmt.Sprintf("EndBlockInParameterFragment { start: %s, end: %s }", f.Start.String(), f.EndFlexibleBlockFragment.String())
}

type EndBodyFragment struct {
	EndFlexibleBlockFragment

	Start *StartBodyFragment
}

func (f *EndBodyFragment) IncLineCount(force bool) bool {
	if f.LineCount < f.MaximalLineCount {
		f.LineCount = f.LineCount + 1

		if !force {
			if f.LineCount == 1 && f.Start.LineCount == 0 {
				f.Start.LineCount = f.LineCount
			}
		}

		return true
	}
	return false
}

func (f *EndBodyFragment) DecLineCount(force bool) bool {
	if f.LineCount > f.MinimalLineCount {
		f.LineCount = f.LineCount - 1

		if !force {
			if f.LineCount == 0 {
				f.Start.LineCount = f.LineCount
			}
		}

		return true
	}
	return false
}

func (f *EndBodyFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitEndBodyFragment(f)
}

func (f *EndBodyFragment) String() string {
	return fmt.Sprintf("EndBodyFragment { start: %s, end: %s }", f.Start.String(), f.EndFlexibleBlockFragment.String())
}

type EndBodyInParameterFragment struct {
	EndBodyFragment
}

func (f *EndBodyInParameterFragment) IncLineCount(force bool) bool {
	if f.LineCount < f.MaximalLineCount {
		f.LineCount = f.LineCount + 1
		return true
	}
	return false
}

func (f *EndBodyInParameterFragment) DecLineCount(force bool) bool {
	if f.LineCount > f.MinimalLineCount {
		f.LineCount = f.LineCount - 1
		return true
	}
	return false
}

func (f *EndBodyInParameterFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitEndBodyInParameterFragment(f)
}

func (f *EndBodyInParameterFragment) String() string {
	return fmt.Sprintf("EndBodyInParameterFragment { start: %s, end: %s }", f.Start.String(), f.EndFlexibleBlockFragment.String())
}

var EndMovableBlock = NewEndMovableJavaBlockFragment()

type EndMovableJavaBlockFragment struct {
	EndMovableBlockFragment
}

func (f *EndMovableJavaBlockFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitEndMovableJavaBlockFragment(f)
}

func (f *EndMovableJavaBlockFragment) String() string {
	return fmt.Sprintf("EndMovableJavaBlockFragment { %s }", f.EndMovableBlockFragment.String())
}

type EndSingleStatementBlockFragment struct {
	EndFlexibleBlockFragment

	Start *StartSingleStatementBlockFragment
}

func (f *EndSingleStatementBlockFragment) IncLineCount(force bool) bool {
	if f.LineCount < f.MaximalLineCount {
		f.LineCount = f.LineCount + 1

		if !force {
			if f.Start.LineCount == 0 {
				f.Start.LineCount = 1
			}
		}

		return true
	}
	return false
}

func (f *EndSingleStatementBlockFragment) DecLineCount(force bool) bool {
	if f.LineCount > f.MinimalLineCount {
		f.LineCount = f.LineCount - 1

		if !force {
			if f.LineCount == 0 {
				f.Start.LineCount = 0
			}
		}

		return true
	}
	return false
}

func (f *EndSingleStatementBlockFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitEndSingleStatementBlockFragment(f)
}

func (f *EndSingleStatementBlockFragment) String() string {
	return fmt.Sprintf("EndSingleStatementBlockFragment { start: %s, end: %s }", f.Start.String(), f.EndFlexibleBlockFragment.String())
}

type EndStatementsBlockFragment struct {
	EndFlexibleBlockFragment

	Group StartStatementsBlockFragmentGroup
}

func (f *EndStatementsBlockFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitEndStatementsBlockFragment(f)
}

func (f *EndStatementsBlockFragment) String() string {
	return fmt.Sprintf("EndSingleStatementBlockFragment { start: %s, end: %s }", f.EndFlexibleBlockFragment.String(), f.Group.String())
}
