package fragment

import (
	"fmt"
	"strings"
)

func NewFlexibleFragment(minimalLineCount int, lineCount int, maximalLineCount int, weight int, label string) *FlexibleFragment {
	return &FlexibleFragment{
		MinimalLineCount: minimalLineCount,
		MaximalLineCount: maximalLineCount,
		InitialLineCount: lineCount,
		LineCount:        lineCount,
		Weight:           weight,
		Label:            label,
	}
}

type IFlexibleFragment interface {
	GetLineCount() int
	flexibleFragmentIgnore()
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

func (f *FlexibleFragment) flexibleFragmentIgnore() {}
