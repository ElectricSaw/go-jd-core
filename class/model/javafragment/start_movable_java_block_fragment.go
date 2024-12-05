package javafragment

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/model/fragment"
)

var StartMovableTypeBlock = NewStartMovableJavaBlockFragment(1)
var StartMovableFieldBlock = NewStartMovableJavaBlockFragment(2)
var StartMovableMethodBlock = NewStartMovableJavaBlockFragment(3)

func NewStartMovableJavaBlockFragment(typ int) intmod.IStartMovableJavaBlockFragment {
	return &StartMovableJavaBlockFragment{
		StartMovableBlockFragment: *fragment.NewStartMovableBlockFragment(typ).(*fragment.StartMovableBlockFragment),
	}
}

type StartMovableJavaBlockFragment struct {
	fragment.StartMovableBlockFragment
}

func (f *StartMovableJavaBlockFragment) Accept(visitor intmod.IJavaFragmentVisitor) {
	visitor.VisitStartMovableJavaBlockFragment(f)
}
