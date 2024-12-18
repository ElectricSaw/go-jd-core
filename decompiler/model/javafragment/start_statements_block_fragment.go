package javafragment

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/fragment"
	"math"
)

func NewStartStatementsBlockFragment(minimalLineCount, lineCount, maximalLineCount, weight int,
	label string) intmod.IStartStatementsBlockFragment {
	return NewStartStatementsBlockFragmentWithGroup(minimalLineCount, lineCount,
		maximalLineCount, weight, label, NewStartStatementsBlockFragmentGroup())
}

func NewStartStatementsBlockFragmentWithGroup(minimalLineCount, lineCount, maximalLineCount, weight int,
	label string, group intmod.IStartStatementsBlockFragmentGroup) intmod.IStartStatementsBlockFragment {
	f := &StartStatementsBlockFragment{
		StartFlexibleBlockFragment: *fragment.NewStartFlexibleBlockFragment(minimalLineCount, lineCount,
			maximalLineCount, weight, label).(*fragment.StartFlexibleBlockFragment),
		group: group,
	}
	f.group.Add(f)
	return f
}

type StartStatementsBlockFragment struct {
	fragment.StartFlexibleBlockFragment

	group intmod.IStartStatementsBlockFragmentGroup
}

func (f *StartStatementsBlockFragment) Group() intmod.IStartStatementsBlockFragmentGroup {
	return f.group
}

func (f *StartStatementsBlockFragment) SetGroup(group intmod.IStartStatementsBlockFragmentGroup) {
	f.group = group
}

func (f *StartStatementsBlockFragment) Accept(visitor intmod.IJavaFragmentVisitor) {
	visitor.VisitStartStatementsBlockFragment(f)
}

func NewStartStatementsBlockFragmentGroup() intmod.IStartStatementsBlockFragmentGroup {
	return &StartStatementsBlockFragmentGroup{minimalLineCount: math.MaxInt}
}

type StartStatementsBlockFragmentGroup struct {
	fragments        []intmod.IFlexibleFragment
	minimalLineCount int
}

func (g *StartStatementsBlockFragmentGroup) MinimalLineCount() int {
	if g.minimalLineCount == math.MaxInt {
		for _, frag := range g.fragments {
			if g.minimalLineCount > frag.LineCount() {
				g.minimalLineCount = frag.LineCount()
			}
		}
	}
	return g.minimalLineCount
}

func (g *StartStatementsBlockFragmentGroup) SetMinimalLineCount(minimalLineCount int) {
	g.minimalLineCount = minimalLineCount
}

func (g *StartStatementsBlockFragmentGroup) Get(index int) intmod.IFlexibleFragment {
	return g.fragments[index]
}

func (g *StartStatementsBlockFragmentGroup) Add(frag intmod.IFlexibleFragment) {
	g.fragments = append(g.fragments, frag)
}

func (g *StartStatementsBlockFragmentGroup) Remove(index int) intmod.IFlexibleFragment {
	ret := g.fragments[index]
	g.fragments = append(g.fragments[:index], g.fragments[index+1:]...)
	return ret
}

func (g *StartStatementsBlockFragmentGroup) ToSlice() []intmod.IFlexibleFragment {
	return g.fragments
}
