package javafragment

var EndMovableBlock = NewEndMovableJavaBlockFragment()

func NewEndMovableJavaBlockFragment() EndMovableJavaBlockFragment {
	return EndMovableJavaBlockFragment{
		EndMovableBlockFragment: NewEndMovableBlockFragment(),
	}
}

type EndMovableJavaBlockFragment struct {
	EndMovableBlockFragment
}

func (f *EndMovableJavaBlockFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitEndMovableJavaBlockFragment(f)
}
