package fragment

func NewEndMovableBlockFragment() *EndMovableBlockFragment {
	return &EndMovableBlockFragment{
		FlexibleFragment: *NewFlexibleFragment(0, 0, 0, 0, "End movable block"),
	}
}

type EndMovableBlockFragment struct {
	FlexibleFragment
}

func (f *EndMovableBlockFragment) Accept(visitor FragmentVisitor) {
	visitor.VisitEndMovableBlockFragment(f)
}

func (f *EndMovableBlockFragment) String() string {
	return "{end-movable-block}"
}
