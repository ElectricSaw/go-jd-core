package javafragment

import intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"

func NewEndBodyInParameterFragment(minimalLineCount, lineCount, maximalLineCount,
	weight int, label string, start intmod.IStartBodyFragment) intmod.IEndBodyInParameterFragment {
	return &EndBodyInParameterFragment{
		EndBodyFragment: *NewEndBodyFragment(minimalLineCount, lineCount,
			maximalLineCount, weight, label, start).(*EndBodyFragment),
	}
}

type EndBodyInParameterFragment struct {
	EndBodyFragment
}

func (f *EndBodyInParameterFragment) IncLineCount(force bool) bool {
	if f.LineCount() < f.MaximalLineCount() {
		f.SetLineCount(f.LineCount() + 1)
		return true
	}
	return false
}

func (f *EndBodyInParameterFragment) DecLineCount(force bool) bool {
	if f.LineCount() > f.MinimalLineCount() {
		f.SetLineCount(f.LineCount() - 1)
		return true
	}
	return false
}

func (f *EndBodyInParameterFragment) Accept(visitor intmod.IJavaFragmentVisitor) {
	visitor.VisitEndBodyInParameterFragment(f)
}
