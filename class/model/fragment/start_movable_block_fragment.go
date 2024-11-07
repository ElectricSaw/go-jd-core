package fragment

func NewStartMovableBlockFragment(typ int) *StartMovableBlockFragment {
	return &StartMovableBlockFragment{
		FlexibleFragment: *NewFlexibleFragment(0, 0, 0, 0, "Start movable block"),
		Type:             typ,
	}
}

type StartMovableBlockFragment struct {
	FlexibleFragment

	Type int
}

func (f *StartMovableBlockFragment) Accept(visitor FragmentVisitor) {
	visitor.VisitStartMovableBlockFragment(f)
}

func (f *StartMovableBlockFragment) String() string {
	return "{start-movable-block}"
}
