package javafragment

import (
	"bitbucket.org/coontec/go-jd-core/class/model/fragment"
	"math"
)

func NewStartStatementsBlockFragment(minimalLineCount int, lineCount int, maximalLineCount int, weight int, label string) *StartStatementsBlockFragment {
	f := &StartStatementsBlockFragment{
		StartFlexibleBlockFragment: *fragment.NewStartFlexibleBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight, label),
	}
	f.Group = NewStartStatementsBlockFragmentGroup(f)
	return f
}

func NewStartStatementsBlockFragmentWithGroup(minimalLineCount int, lineCount int, maximalLineCount int, weight int, label string, group *StartStatementsBlockFragmentGroup) *StartStatementsBlockFragment {
	f := &StartStatementsBlockFragment{
		StartFlexibleBlockFragment: *fragment.NewStartFlexibleBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight, label),
		Group:                      group,
	}
	f.Group.add(f)
	return f
}

type StartStatementsBlockFragment struct {
	fragment.StartFlexibleBlockFragment

	Group *StartStatementsBlockFragmentGroup
}

func (f *StartStatementsBlockFragment) Accept(visitor JavaFragmentVisitor) {
	visitor.VisitStartStatementsBlockFragment(f)
}

func NewStartStatementsBlockFragmentGroup(frag fragment.IFlexibleFragment) *StartStatementsBlockFragmentGroup {
	f := &StartStatementsBlockFragmentGroup{minimalLineCount: math.MaxInt}
	f.fragments = append(f.fragments, frag)
	return f
}

type StartStatementsBlockFragmentGroup struct {
	fragments        []fragment.IFlexibleFragment
	minimalLineCount int
}

func (g *StartStatementsBlockFragmentGroup) add(frag fragment.IFlexibleFragment) {
	g.fragments = append(g.fragments, frag)
}

func (g *StartStatementsBlockFragmentGroup) MinimalLineCount() int {
	if g.minimalLineCount == math.MaxInt {
		for _, frag := range g.fragments {
			if g.minimalLineCount > frag.GetLineCount() {
				g.minimalLineCount = frag.GetLineCount()
			}
		}
	}
	return g.minimalLineCount
}
