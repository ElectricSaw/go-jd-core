package fragment

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"strings"
)

func NewFlexibleFragment(minimalLineCount, lineCount, maximalLineCount, weight int, label string) intmod.IFlexibleFragment {
	return &FlexibleFragment{
		minimalLineCount: minimalLineCount,
		maximalLineCount: maximalLineCount,
		initialLineCount: lineCount,
		lineCount:        lineCount,
		weight:           weight,
		label:            label,
	}
}

type FlexibleFragment struct {
	minimalLineCount int
	maximalLineCount int
	initialLineCount int
	lineCount        int
	weight           int
	label            string
}

func (f *FlexibleFragment) MinimalLineCount() int {
	return f.minimalLineCount
}

func (f *FlexibleFragment) SetMinimalLineCount(minimalLineCount int) {
	f.minimalLineCount = minimalLineCount
}

func (f *FlexibleFragment) MaximalLineCount() int {
	return f.maximalLineCount
}

func (f *FlexibleFragment) SetMaximalLineCount(maximalLineCount int) {
	f.maximalLineCount = maximalLineCount
}

func (f *FlexibleFragment) InitialLineCount() int {
	return f.initialLineCount
}

func (f *FlexibleFragment) SetInitialLineCount(initialLineCount int) {
	f.initialLineCount = initialLineCount
}

func (f *FlexibleFragment) LineCount() int {
	return f.lineCount
}

func (f *FlexibleFragment) SetLineCount(lineCount int) {
	f.lineCount = lineCount
}

func (f *FlexibleFragment) Weight() int {
	return f.weight
}

func (f *FlexibleFragment) SetWeight(weight int) {
	f.weight = weight
}

func (f *FlexibleFragment) Label() string {
	return f.label
}

func (f *FlexibleFragment) SetLabel(label string) {
	f.label = label
}

func (f *FlexibleFragment) ResetLineCount() {
	f.lineCount = f.initialLineCount
}

func (f *FlexibleFragment) IncLineCount(force bool) bool {
	if f.lineCount < f.maximalLineCount {
		f.lineCount++
		return true
	}
	return false
}

func (f *FlexibleFragment) DecLineCount(force bool) bool {
	if f.lineCount > f.minimalLineCount {
		f.lineCount--
		return true
	}
	return false
}

func (f *FlexibleFragment) AcceptFragmentVisitor(visitor intmod.IFragmentVisitor) {
	visitor.VisitFlexibleFragment(f)
}

func (f *FlexibleFragment) String() string {
	var msg strings.Builder

	msg.WriteString(fmt.Sprintf("FlexibleFragment { minimal-line-count=%d", f.minimalLineCount))
	msg.WriteString(fmt.Sprintf(", maximal-line-count=%d", f.maximalLineCount))
	msg.WriteString(fmt.Sprintf(", initial-line-count=%d", f.initialLineCount))
	msg.WriteString(fmt.Sprintf(", line-count=%d", f.lineCount))
	msg.WriteString(fmt.Sprintf(", weight=%d", f.weight))

	if f.label != "" {
		msg.WriteString(fmt.Sprintf(", label=%s }", f.label))
	} else {
		msg.WriteString("}")
	}

	return msg.String()
}
