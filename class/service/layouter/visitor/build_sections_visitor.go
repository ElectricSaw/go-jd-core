package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/class/util"
)

func NewBuildSectionsVisitor() *BuildSectionsVisitor {
	return &BuildSectionsVisitor{
		sections:          util.NewDefaultList[ISection](),
		flexibleFragments: util.NewDefaultList[intmod.IFlexibleFragment](),
		previousSection:   nil,
	}
}

type BuildSectionsVisitor struct {
	sections          util.IList[ISection]
	flexibleFragments util.IList[intmod.IFlexibleFragment]
	previousSection   ISection
}

func (v *BuildSectionsVisitor) Sections() util.IList[ISection] {
	v.sections.Add(NewSection(v.flexibleFragments, nil, v.previousSection))
	return v.sections
}

func (v *BuildSectionsVisitor) VisitFlexibleFragment(frag intmod.IFlexibleFragment) {
	v.flexibleFragments.Add(frag)
}

func (v *BuildSectionsVisitor) VisitEndFlexibleBlockFragment(frag intmod.IEndFlexibleBlockFragment) {
	v.flexibleFragments.Add(frag)
}

func (v *BuildSectionsVisitor) VisitEndMovableBlockFragment(frag intmod.IEndMovableBlockFragment) {
	v.flexibleFragments.Add(frag)
}

func (v *BuildSectionsVisitor) VisitSpacerBetweenMovableBlocksFragment(frag intmod.ISpacerBetweenMovableBlocksFragment) {
	v.flexibleFragments.Add(frag)
}

func (v *BuildSectionsVisitor) VisitStartFlexibleBlockFragment(frag intmod.IStartFlexibleBlockFragment) {
	v.flexibleFragments.Add(frag)
}

func (v *BuildSectionsVisitor) VisitStartMovableBlockFragment(frag intmod.IStartMovableBlockFragment) {
	v.flexibleFragments.Add(frag)
}

func (v *BuildSectionsVisitor) VisitFixedFragment(frag intmod.IFixedFragment) {
	v.sections.Add(NewSection(v.flexibleFragments, frag, v.previousSection))
	v.flexibleFragments = util.NewDefaultList[intmod.IFlexibleFragment]()
}
