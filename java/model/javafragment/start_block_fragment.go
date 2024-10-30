package javafragment

import "bitbucket.org/coontec/javaClass/java/model/fragment"

func NewStartBlockFragment(minimalLineCount int, lineCount int, maximalLineCount int, weight int, label string) StartBlockFragment {
	return StartBlockFragment{
		StartFlexibleBlockFragment: fragment.NewStartFlexibleBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight, label),
	}
}

type StartBlockFragment struct {
	fragment.StartFlexibleBlockFragment

	end *EndBlockFragment
}

func (f *StartBlockFragment) EndArrayInitializerBlockFragment() *EndBlockFragment {
	return f.end
}

func (f *StartBlockFragment) SetEndArrayInitializerBlockFragment(end *EndBlockFragment) {
	f.end = end
}

func (f *StartBlockFragment) IncLineCount(force bool) bool {
	if f.LineCount() < f.MaximalLineCount {
		f.SetLineCount(f.LineCount() + 1)

		if !force {
			if f.LineCount() == 1 && f.end.LineCount() == 0 {
				f.end.SetLineCount(f.LineCount())
			}
		}

		return true
	}
	return false
}

func (f *StartBlockFragment) DecLineCount(force bool) bool {
	if f.LineCount() > f.MinimalLineCount {
		f.SetLineCount(f.LineCount() - 1)

		if !force {
			if f.LineCount() == 1 {
				f.end.SetLineCount(f.LineCount())
			}
		}

		return true
	}
	return false
}

func (f *StartBlockFragment) Accept(visitor JavaFragmentVisitor) {
	visitor.VisitStartBlockFragment(f)
}
