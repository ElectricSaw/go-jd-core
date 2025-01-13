package javafragment

func NewStartMovableJavaBlockFragment(typ int) StartMovableJavaBlockFragment {
	return StartMovableJavaBlockFragment{
		StartMovableBlockFragment: NewStartMovableBlockFragment(typ),
	}
}

type StartMovableJavaBlockFragment struct {
	StartMovableBlockFragment
}

func (f *StartMovableJavaBlockFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitStartMovableJavaBlockFragment(f)
}
