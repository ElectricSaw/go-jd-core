package fragment

import intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"

func NewEndMovableBlockFragment() intmod.IEndMovableBlockFragment {
	return &EndMovableBlockFragment{
		FlexibleFragment: *NewFlexibleFragment(0, 0,
			0, 0, "End movable block").(*FlexibleFragment),
	}
}

type EndMovableBlockFragment struct {
	FlexibleFragment
}

func (f *EndMovableBlockFragment) AcceptFragmentVisitor(visitor intmod.IFragmentVisitor) {
	visitor.VisitEndMovableBlockFragment(f)
}

func (f *EndMovableBlockFragment) String() string {
	return "{end-movable-block}"
}
