package visitor

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
)

func NewPopulateBlackListNamesVisitor(names []string) *PopulateBlackListNamesVisitor {
	return &PopulateBlackListNamesVisitor{
		blackListNames: names,
	}
}

type PopulateBlackListNamesVisitor struct {
	_type.AbstractNopTypeArgumentVisitor

	blackListNames []string
}

func (v *PopulateBlackListNamesVisitor) VisitObjectType(t intmod.IObjectType) {
	v.blackListNames = append(v.blackListNames, t.Name())
}

func (v *PopulateBlackListNamesVisitor) VisitInnerObjectType(t intmod.IInnerObjectType) {
	v.blackListNames = append(v.blackListNames, t.Name())
}

func (v *PopulateBlackListNamesVisitor) VisitGenericType(t intmod.IGenericType) {
	v.blackListNames = append(v.blackListNames, t.Name())
}
