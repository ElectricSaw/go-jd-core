package _type

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

func NewTypes() intmod.ITypes {
	return &Types{}
}

type Types struct {
	AbstractType
	util.DefaultList[intmod.IType]
}

func (t *Types) ForEach(action func(intmod.IType)) {
	for _, item := range t.ToSlice() {
		action(item.(intmod.IType))
	}
}

func (t *Types) IsTypes() bool {
	return true
}

func (t *Types) AcceptTypeVisitor(visitor intmod.ITypeVisitor) {
	visitor.VisitTypes(t)
}
