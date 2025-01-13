package javafragment

import (
	"fmt"
	"strings"
)

func NewEndFlexibleBlockFragment(minimalLineCount, lineCount, maximalLineCount,
	weight int, label string) EndFlexibleBlockFragment {
	return EndFlexibleBlockFragment{
		FlexibleFragment: NewFlexibleFragment(minimalLineCount, lineCount,
			maximalLineCount, weight, label),
	}
}

func NewEndMovableBlockFragment() EndMovableBlockFragment {
	return EndMovableBlockFragment{
		FlexibleFragment: NewFlexibleFragment(0, 0,
			0, 0, "End movable block"),
	}
}

func NewFixedFragment(firstLineNumber, lastLineNumber int) FixedFragment {
	return FixedFragment{
		FirstLineNumber: firstLineNumber,
		LastLineNumber:  lastLineNumber,
	}
}

func NewFlexibleFragment(minimalLineCount, lineCount, maximalLineCount, weight int, label string) FlexibleFragment {
	return FlexibleFragment{
		MinimalLineCount: minimalLineCount,
		MaximalLineCount: maximalLineCount,
		InitialLineCount: lineCount,
		LineCount:        lineCount,
		Weight:           weight,
		Label:            label,
	}
}

func NewSpacerBetweenMovableBlocksFragment(minimalLineCount, lineCount, maximalLineCount, weight int, label string) SpacerBetweenMovableBlocksFragment {
	return SpacerBetweenMovableBlocksFragment{
		FlexibleFragment: NewFlexibleFragment(minimalLineCount, lineCount,
			maximalLineCount, weight, label),
	}
}

func NewStartFlexibleBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight int, label string) StartFlexibleBlockFragment {
	return StartFlexibleBlockFragment{
		FlexibleFragment: NewFlexibleFragment(minimalLineCount, lineCount,
			maximalLineCount, weight, label),
	}
}

func NewStartMovableBlockFragment(typ int) StartMovableBlockFragment {
	return StartMovableBlockFragment{
		FlexibleFragment: NewFlexibleFragment(0, 0,
			0, 0, "Start movable block"),
		Type: typ,
	}
}

type IFragment interface {
	AcceptFragmentVisitor(visitor IFragmentVisitor)
	String() string
}

type IFragmentVisitor interface {
	VisitFlexibleFragment(fragment IFragment)
	VisitEndFlexibleBlockFragment(fragment IFragment)
	VisitEndMovableBlockFragment(fragment IFragment)
	VisitSpacerBetweenMovableBlocksFragment(fragment IFragment)
	VisitStartFlexibleBlockFragment(fragment IFragment)
	VisitStartMovableBlockFragment(fragment IFragment)
	VisitFixedFragment(fragment IFragment)
}

type AbstractNopFlexibleFragmentVisitor struct {
}

func (f AbstractNopFlexibleFragmentVisitor) VisitFlexibleFragment(_ IFragment)                   {}
func (f AbstractNopFlexibleFragmentVisitor) VisitEndFlexibleBlockFragment(_ IFragment)           {}
func (f AbstractNopFlexibleFragmentVisitor) VisitEndMovableBlockFragment(_ IFragment)            {}
func (f AbstractNopFlexibleFragmentVisitor) VisitSpacerBetweenMovableBlocksFragment(_ IFragment) {}
func (f AbstractNopFlexibleFragmentVisitor) VisitStartFlexibleBlockFragment(_ IFragment)         {}
func (f AbstractNopFlexibleFragmentVisitor) VisitStartMovableBlockFragment(_ IFragment)          {}
func (f AbstractNopFlexibleFragmentVisitor) VisitFixedFragment(_ IFragment)                      {}

type EndFlexibleBlockFragment struct {
	FlexibleFragment
}

func (f *EndFlexibleBlockFragment) AcceptFragmentVisitor(visitor IFragmentVisitor) {
	visitor.VisitEndFlexibleBlockFragment(f)
}

type EndMovableBlockFragment struct {
	FlexibleFragment
}

func (f *EndMovableBlockFragment) AcceptFragmentVisitor(visitor IFragmentVisitor) {
	visitor.VisitEndMovableBlockFragment(f)
}

func (f *EndMovableBlockFragment) String() string {
	return "{end-movable-block}"
}

type FixedFragment struct {
	FirstLineNumber int
	LastLineNumber  int
}

func (f *FixedFragment) AcceptFragmentVisitor(visitor IFragmentVisitor) {
	visitor.VisitFixedFragment(f)
}

func (f *FixedFragment) String() string {
	return fmt.Sprintf("{first-line-number=%d, last-line-number=%d}", f.FirstLineNumber, f.LastLineNumber)
}

type IFlexibleFragment interface {
	GetLineCount() int
	ResetLineCount()
	IncLineCount(force bool) bool
	DecLineCount(force bool) bool
	AcceptFragmentVisitor(visitor IFragmentVisitor)
	String() string
}

type FlexibleFragment struct {
	MinimalLineCount int
	MaximalLineCount int
	InitialLineCount int
	LineCount        int
	Weight           int
	Label            string
}

func (f *FlexibleFragment) GetLineCount() int {
	return f.LineCount
}

func (f *FlexibleFragment) ResetLineCount() {
	f.LineCount = f.InitialLineCount
}

func (f *FlexibleFragment) IncLineCount(force bool) bool {
	if f.LineCount < f.MaximalLineCount {
		f.LineCount++
		return true
	}
	return false
}

func (f *FlexibleFragment) DecLineCount(force bool) bool {
	if f.LineCount > f.MinimalLineCount {
		f.LineCount--
		return true
	}
	return false
}

func (f *FlexibleFragment) AcceptFragmentVisitor(visitor IFragmentVisitor) {
	visitor.VisitFlexibleFragment(f)
}

func (f *FlexibleFragment) String() string {
	var msg strings.Builder

	msg.WriteString(fmt.Sprintf("FlexibleFragment { minimal-line-count=%d", f.MinimalLineCount))
	msg.WriteString(fmt.Sprintf(", maximal-line-count=%d", f.MaximalLineCount))
	msg.WriteString(fmt.Sprintf(", initial-line-count=%d", f.InitialLineCount))
	msg.WriteString(fmt.Sprintf(", line-count=%d", f.LineCount))
	msg.WriteString(fmt.Sprintf(", weight=%d", f.Weight))

	if f.Label != "" {
		msg.WriteString(fmt.Sprintf(", label=%s }", f.Label))
	} else {
		msg.WriteString("}")
	}

	return msg.String()
}

type SpacerBetweenMovableBlocksFragment struct {
	FlexibleFragment
}

func (f *SpacerBetweenMovableBlocksFragment) SetInitialLineCount(initialLineCount int) {
	f.InitialLineCount = initialLineCount
	f.LineCount = initialLineCount
}

func (f *SpacerBetweenMovableBlocksFragment) AcceptFragmentVisitor(visitor IFragmentVisitor) {
	visitor.VisitSpacerBetweenMovableBlocksFragment(f)
}

type StartFlexibleBlockFragment struct {
	FlexibleFragment
}

func (f *StartFlexibleBlockFragment) AcceptFragmentVisitor(visitor IFragmentVisitor) {
	visitor.VisitStartFlexibleBlockFragment(f)
}

type StartMovableBlockFragment struct {
	FlexibleFragment

	Type int
}

func (f *StartMovableBlockFragment) AcceptFragmentVisitor(visitor IFragmentVisitor) {
	visitor.VisitStartMovableBlockFragment(f)
}

func (f *StartMovableBlockFragment) String() string {
	return "{start-movable-block}"
}
