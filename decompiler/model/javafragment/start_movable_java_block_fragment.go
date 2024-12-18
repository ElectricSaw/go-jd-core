package javafragment

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/fragment"
)

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
