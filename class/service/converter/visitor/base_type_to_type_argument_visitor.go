package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	_type "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/type"
)

func NewBaseTypeToTypeArgumentVisitor() *BaseTypeToTypeArgumentVisitor {
	return &BaseTypeToTypeArgumentVisitor{}
}

type BaseTypeToTypeArgumentVisitor struct {
	typeArgument intmod.ITypeArgument
}

func (v *BaseTypeToTypeArgumentVisitor) Init() {
	v.typeArgument = nil
}

func (v *BaseTypeToTypeArgumentVisitor) TypeArgument() intmod.ITypeArgument {
	return v.typeArgument
}

func (v *BaseTypeToTypeArgumentVisitor) VisitPrimitiveType(y intmod.IPrimitiveType) {
	v.typeArgument = y
}

func (v *BaseTypeToTypeArgumentVisitor) VisitObjectType(y intmod.IObjectType) {
	v.typeArgument = y
}

func (v *BaseTypeToTypeArgumentVisitor) VisitInnerObjectType(y intmod.IInnerObjectType) {
	v.typeArgument = y
}

func (v *BaseTypeToTypeArgumentVisitor) VisitGenericType(y intmod.IGenericType) {
	v.typeArgument = y
}

func (v *BaseTypeToTypeArgumentVisitor) VisitTypes(types intmod.ITypes) {
	if types.IsEmpty() {
		v.typeArgument = _type.WildcardTypeArgumentEmpty
	} else {
		types.First().AcceptTypeVisitor(v)
	}
}
