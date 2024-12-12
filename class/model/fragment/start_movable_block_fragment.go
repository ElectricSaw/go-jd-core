package fragment

import intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"

func NewStartMovableBlockFragment(typ int) intmod.IStartMovableBlockFragment {
	return &StartMovableBlockFragment{
		FlexibleFragment: *NewFlexibleFragment(0, 0,
			0, 0, "Start movable block").(*FlexibleFragment),
		typ: typ,
	}
}

type StartMovableBlockFragment struct {
	FlexibleFragment

	typ int
}

func (f *StartMovableBlockFragment) Type() int {
	return f.typ
}

func (f *StartMovableBlockFragment) SetType(typ int) {
	f.typ = typ
}

func (f *StartMovableBlockFragment) AcceptFragmentVisitor(visitor intmod.IFragmentVisitor) {
	visitor.VisitStartMovableBlockFragment(f)
}

func (f *StartMovableBlockFragment) String() string {
	return "{start-movable-block}"
}
