package javafragment

import intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"

func NewEndBlockInParameterFragment(minimalLineCount, lineCount, maximalLineCount,
	weight int, label string, start intmod.IStartBlockFragment) intmod.IEndBlockInParameterFragment {
	return &EndBlockInParameterFragment{
		EndBlockFragment: *NewEndBlockFragment(minimalLineCount, lineCount,
			maximalLineCount, weight, label, start).(*EndBlockFragment),
	}
}

type EndBlockInParameterFragment struct {
	EndBlockFragment
}

func (f *EndBlockInParameterFragment) Accept(visitor intmod.IJavaFragmentVisitor) {
	visitor.VisitEndBlockInParameterFragment(f)
}
