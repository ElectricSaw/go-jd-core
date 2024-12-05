package javafragment

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/model/fragment"
)

func NewEndStatementsBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight int,
	label string, group intmod.IStartStatementsBlockFragmentGroup) intmod.IEndStatementsBlockFragment {
	f := &EndStatementsBlockFragment{
		EndFlexibleBlockFragment: *fragment.NewEndFlexibleBlockFragment(minimalLineCount,
			lineCount, maximalLineCount, weight, label).(*fragment.EndFlexibleBlockFragment),
		group: group,
	}
	f.group.add(f)
	return f
}

type EndStatementsBlockFragment struct {
	fragment.EndFlexibleBlockFragment

	group intmod.IStartStatementsBlockFragmentGroup
}

func (f *EndStatementsBlockFragment) Group() intmod.IStartStatementsBlockFragmentGroup {
	return f.group
}

func (f *EndStatementsBlockFragment) SetGroup(group intmod.IStartStatementsBlockFragmentGroup) {
	f.group = group
}

func (f *EndStatementsBlockFragment) Accept(visitor intmod.IJavaFragmentVisitor) {
	visitor.VisitEndStatementsBlockFragment(f)
}
