package visitor

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

func NewUpdateSpacerBetweenMovableBlocksVisitor() *UpdateSpacerBetweenMovableBlocksVisitor {
	return &UpdateSpacerBetweenMovableBlocksVisitor{
		blocks:  util.NewDefaultList[intmod.IStartMovableBlockFragment](),
		spacers: util.NewDefaultList[intmod.ISpacerBetweenMovableBlocksFragment](),
	}
}

type UpdateSpacerBetweenMovableBlocksVisitor struct {
	blocks  util.IList[intmod.IStartMovableBlockFragment]
	spacers util.IList[intmod.ISpacerBetweenMovableBlocksFragment]

	lastStartMovableBlockFragmentType int
	lastSpacer                        intmod.ISpacerBetweenMovableBlocksFragment
	depth                             int
}

func (v *UpdateSpacerBetweenMovableBlocksVisitor) Reset() {
	v.lastStartMovableBlockFragmentType = 0
	v.lastSpacer = nil
	v.depth = 0
}

func (v *UpdateSpacerBetweenMovableBlocksVisitor) VisitStartMovableBlockFragment(frag intmod.IStartMovableBlockFragment) {
	if v.lastSpacer != nil {
		if v.lastStartMovableBlockFragmentType == 2 && frag.Type() == 2 {
			v.lastSpacer.SetInitialLineCount(1)
		} else {
			v.lastSpacer.SetInitialLineCount(2)
		}
	}

	if v.depth != 0 {
		v.blocks.Add(frag)
		v.spacers.Add(v.lastSpacer)
		v.lastSpacer = nil
	}

	v.lastStartMovableBlockFragmentType = frag.Type()
	v.depth = 1
}

func (v *UpdateSpacerBetweenMovableBlocksVisitor) VisitEndMovableBlockFragment(_ intmod.IEndMovableBlockFragment) {
	if v.depth != 1 {
		v.lastSpacer = v.spacers.RemoveAt(v.spacers.Size() - 1)
		v.lastStartMovableBlockFragmentType = v.blocks.RemoveLast().Type()
	}
	v.depth = 0
}

func (v *UpdateSpacerBetweenMovableBlocksVisitor) VisitSpacerBetweenMovableBlocksFragment(frag intmod.ISpacerBetweenMovableBlocksFragment) {
	v.lastSpacer = frag
}

func (v *UpdateSpacerBetweenMovableBlocksVisitor) VisitFlexibleFragment(_ intmod.IFlexibleFragment) {
}

func (v *UpdateSpacerBetweenMovableBlocksVisitor) VisitEndFlexibleBlockFragment(_ intmod.IEndFlexibleBlockFragment) {
}

func (v *UpdateSpacerBetweenMovableBlocksVisitor) VisitStartFlexibleBlockFragment(_ intmod.IStartFlexibleBlockFragment) {
}

func (v *UpdateSpacerBetweenMovableBlocksVisitor) VisitFixedFragment(_ intmod.IFixedFragment) {
}
