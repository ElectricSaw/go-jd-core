package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	_type "github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/type"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewPopulateBlackListNamesVisitor(names util.ISet[string]) intsrv.IPopulateBlackListNamesVisitor {
	return &PopulateBlackListNamesVisitor{
		blackListNames: names,
	}
}

type PopulateBlackListNamesVisitor struct {
	_type.AbstractNopTypeArgumentVisitor

	blackListNames util.ISet[string]
}

func (v *PopulateBlackListNamesVisitor) VisitObjectType(t intmod.IObjectType) {
	v.blackListNames.Add(t.Name())
}

func (v *PopulateBlackListNamesVisitor) VisitInnerObjectType(t intmod.IInnerObjectType) {
	v.blackListNames.Add(t.Name())
}

func (v *PopulateBlackListNamesVisitor) VisitGenericType(t intmod.IGenericType) {
	v.blackListNames.Add(t.Name())
}
