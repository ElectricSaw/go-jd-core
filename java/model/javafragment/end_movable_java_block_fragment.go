package javafragment

import "bitbucket.org/coontec/javaClass/java/model/fragment"

var EndMovableBlock = NewEndMovableJavaBlockFragment()

func NewEndMovableJavaBlockFragment() EndMovableJavaBlockFragment {
	return EndMovableJavaBlockFragment{
		EndMovableBlockFragment: fragment.NewEndMovableBlockFragment(),
	}
}

type EndMovableJavaBlockFragment struct {
	fragment.EndMovableBlockFragment
}

func (f *EndMovableJavaBlockFragment) Accept(visitor JavaFragmentVisitor) {
	visitor.VisitEndMovableJavaBlockFragment(f)
}
