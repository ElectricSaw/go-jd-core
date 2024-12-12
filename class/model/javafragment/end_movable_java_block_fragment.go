package javafragment

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/class/model/fragment"
)

var EndMovableBlock = NewEndMovableJavaBlockFragment()

func NewEndMovableJavaBlockFragment() intmod.IEndMovableJavaBlockFragment {
	return &EndMovableJavaBlockFragment{
		EndMovableBlockFragment: *fragment.NewEndMovableBlockFragment().(*fragment.EndMovableBlockFragment),
	}
}

type EndMovableJavaBlockFragment struct {
	fragment.EndMovableBlockFragment
}

func (f *EndMovableJavaBlockFragment) Accept(visitor intmod.IJavaFragmentVisitor) {
	visitor.VisitEndMovableJavaBlockFragment(f)
}
