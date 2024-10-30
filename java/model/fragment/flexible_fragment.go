package fragment

import (
	"fmt"
	"strings"
)

func NewFlexibleFragment(minimalLineCount int, lineCount int, maximalLineCount int, weight int, label string) FlexibleFragment {
	return FlexibleFragment{minimalLineCount, maximalLineCount, lineCount, lineCount, weight, label}
}

type IFlexibleFragment interface {
	LineCount() int
	SetLineCount(lineCount int)
	IncLineCount(force bool) bool
	DecLineCount(force bool) bool
}

type FlexibleFragment struct {
	MinimalLineCount int
	MaximalLineCount int
	InitialLineCount int
	lineCount        int
	Weight           int
	Label            string
}

func (f *FlexibleFragment) LineCount() int {
	return f.lineCount
}

func (f *FlexibleFragment) SetLineCount(lineCount int) {
	f.lineCount = lineCount
}

func (f *FlexibleFragment) IncLineCount(force bool) bool {
	if f.lineCount < f.MaximalLineCount {
		f.lineCount++
		return true
	}
	return false
}

func (f *FlexibleFragment) DecLineCount(force bool) bool {
	if f.lineCount > f.MinimalLineCount {
		f.lineCount--
		return true
	}
	return false
}

func (f *FlexibleFragment) Accept(visitor FragmentVisitor) {
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
