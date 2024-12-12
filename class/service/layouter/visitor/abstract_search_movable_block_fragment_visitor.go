package visitor

import intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"

func NewAbstractSearchMovableBlockFragmentVisitor() *AbstractSearchMovableBlockFragmentVisitor {
	return &AbstractSearchMovableBlockFragmentVisitor{
		depth: 1,
		index: 0,
	}
}

type ISearchMovableBlockFragmentVisitor interface {
	intmod.IFragmentVisitor

	Depth() int
	Index() int
	Reset()
	ResetIndex()
}

type AbstractSearchMovableBlockFragmentVisitor struct {
	depth int
	index int
}

func (v *AbstractSearchMovableBlockFragmentVisitor) Depth() int {
	return v.depth
}

func (v *AbstractSearchMovableBlockFragmentVisitor) Index() int {
	return v.index
}

func (v *AbstractSearchMovableBlockFragmentVisitor) Reset() {
	v.depth = 1
	v.index = 0
}

func (v *AbstractSearchMovableBlockFragmentVisitor) ResetIndex() {
	v.index = 0
}

func (v *AbstractSearchMovableBlockFragmentVisitor) VisitFlexibleFragment(_ intmod.IFlexibleFragment) {
	v.index++
}

func (v *AbstractSearchMovableBlockFragmentVisitor) VisitEndFlexibleBlockFragment(_ intmod.IEndFlexibleBlockFragment) {
	v.index++
}

func (v *AbstractSearchMovableBlockFragmentVisitor) VisitEndMovableBlockFragment(_ intmod.IEndMovableBlockFragment) {
	v.index++
}

func (v *AbstractSearchMovableBlockFragmentVisitor) VisitSpacerBetweenMovableBlocksFragment(_ intmod.ISpacerBetweenMovableBlocksFragment) {
	v.index++
}

func (v *AbstractSearchMovableBlockFragmentVisitor) VisitStartFlexibleBlockFragment(_ intmod.IStartFlexibleBlockFragment) {
	v.index++
}

func (v *AbstractSearchMovableBlockFragmentVisitor) VisitStartMovableBlockFragment(_ intmod.IStartMovableBlockFragment) {
	v.index++
}

func (v *AbstractSearchMovableBlockFragmentVisitor) VisitFixedFragment(_ intmod.IFixedFragment) {
}
