package visitor

import intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"

func NewVisitorsHolder() *VisitorsHolder {
	return &VisitorsHolder{}
}

type VisitorsHolder struct {
	visitor7  ISearchMovableBlockFragmentVisitor
	visitor8  ISearchMovableBlockFragmentVisitor
	visitor9  IStoreMovableBlockFragmentIndexVisitorAbstract
	visitor10 IStoreMovableBlockFragmentIndexVisitorAbstract
}

func (h *VisitorsHolder) ForwardSearchVisitor() ISearchMovableBlockFragmentVisitor {
	if h.visitor7 == nil {
		h.visitor7 = &Visitor7{}
	}
	return h.visitor7
}

func (h *VisitorsHolder) BackwardSearchVisitor() ISearchMovableBlockFragmentVisitor {
	if h.visitor8 == nil {
		h.visitor8 = &Visitor8{}
	}
	return h.visitor8
}

func (h *VisitorsHolder) BackwardSearchStartIndexesVisitor() IStoreMovableBlockFragmentIndexVisitorAbstract {
	if h.visitor9 == nil {
		h.visitor9 = &Visitor9{}
	}
	return h.visitor9
}

func (h *VisitorsHolder) ForwardSearchEndIndexesVisitor() IStoreMovableBlockFragmentIndexVisitorAbstract {
	if h.visitor10 == nil {
		h.visitor10 = &Visitor10{}
	}
	return h.visitor10
}

type Visitor7 struct {
	AbstractSearchMovableBlockFragmentVisitor
}

func (v *Visitor7) VisitEndMovableBlockFragment(_ intmod.IEndMovableBlockFragment) {
	v.depth--
	v.index++
}

func (v *Visitor7) VisitStartMovableBlockFragment(_ intmod.IStartMovableBlockFragment) {
	v.depth++
	v.index++
}

type Visitor8 struct {
	AbstractSearchMovableBlockFragmentVisitor
}

func (v *Visitor8) VisitEndMovableBlockFragment(_ intmod.IEndMovableBlockFragment) {
	v.depth++
	v.index++
}

func (v *Visitor8) VisitStartMovableBlockFragment(_ intmod.IStartMovableBlockFragment) {
	v.depth--
	v.index++
}

type Visitor9 struct {
	AbstractStoreMovableBlockFragmentIndexVisitorAbstract
}

func (v *Visitor9) VisitEndMovableBlockFragment(_ intmod.IEndMovableBlockFragment) {
	if v.enabled {
		v.depth++
		v.index++
	}
}

func (v *Visitor9) VisitStartMovableBlockFragment(_ intmod.IStartMovableBlockFragment) {
	if v.enabled {
		if v.depth == 0 {
			v.enabled = false
		} else {
			v.depth--
			if v.depth == 0 {
				v.storeIndex()
			}
			v.index++
		}
	}
}

type Visitor10 struct {
	AbstractStoreMovableBlockFragmentIndexVisitorAbstract
}

func (v *Visitor10) VisitEndMovableBlockFragment(_ intmod.IEndMovableBlockFragment) {
	if v.enabled {
		if v.depth == 0 {
			v.enabled = false
		} else {
			v.depth--
			if v.depth == 0 {
				v.storeIndex()
			}
			v.index++
		}
	}
}

func (v *Visitor10) VisitStartMovableBlockFragment(_ intmod.IStartMovableBlockFragment) {
	if v.enabled {
		v.depth++
		v.index++
	}
}
