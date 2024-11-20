package javafragment

import "bitbucket.org/coontec/go-jd-core/class/model/fragment"

func NewSpacerFragment(minimalLineCount, lineCount, maximalLineCount, weight int, label string) *SpacerFragment {
	return &SpacerFragment{
		FlexibleFragment: *fragment.NewFlexibleFragment(minimalLineCount, lineCount, maximalLineCount, weight, label),
	}
}

type SpacerFragment struct {
	fragment.FlexibleFragment
}

func (f *SpacerFragment) Accept(visitor JavaFragmentVisitor) {
	visitor.VisitSpacerFragment(f)
}
