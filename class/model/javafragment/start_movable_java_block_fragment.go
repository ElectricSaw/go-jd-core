package javafragment

import "bitbucket.org/coontec/go-jd-core/class/model/fragment"

var StartMovableTypeBlock = NewStartMovableJavaBlockFragment(1)
var StartMovableFieldBlock = NewStartMovableJavaBlockFragment(2)
var StartMovableMethodBlock = NewStartMovableJavaBlockFragment(3)

func NewStartMovableJavaBlockFragment(typ int) *StartMovableJavaBlockFragment {
	return &StartMovableJavaBlockFragment{
		StartMovableBlockFragment: *fragment.NewStartMovableBlockFragment(typ),
	}
}

type StartMovableJavaBlockFragment struct {
	fragment.StartMovableBlockFragment
}

func (f *StartMovableJavaBlockFragment) Accept(visitor JavaFragmentVisitor) {
	visitor.VisitStartMovableJavaBlockFragment(f)
}
