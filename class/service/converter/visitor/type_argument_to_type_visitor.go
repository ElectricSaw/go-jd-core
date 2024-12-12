package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	_type "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/type"
)

func NewTypeArgumentToTypeVisitor() intsrv.ITypeArgumentToTypeVisitor {
	return &TypeArgumentToTypeVisitor{}
}

type TypeArgumentToTypeVisitor struct {
	_type.AbstractTypeArgumentVisitor

	typ intmod.IType
}

func (v *TypeArgumentToTypeVisitor) Init() {
	v.typ = nil
}

func (v *TypeArgumentToTypeVisitor) Type() intmod.IType {
	return v.typ
}

func (v *TypeArgumentToTypeVisitor) VisitTypes(_ intmod.ITypes) {

}

func (v *TypeArgumentToTypeVisitor) VisitDiamondTypeArgument(_ intmod.IDiamondTypeArgument) {
	v.typ = _type.OtTypeObject
}

func (v *TypeArgumentToTypeVisitor) VisitWildcardTypeArgument(_ intmod.IWildcardTypeArgument) {
	v.typ = _type.OtTypeObject
}

func (v *TypeArgumentToTypeVisitor) VisitPrimitiveType(typ intmod.IPrimitiveType) {
	v.typ = typ
}

func (v *TypeArgumentToTypeVisitor) VisitObjectType(typ intmod.IObjectType) {
	v.typ = typ
}

func (v *TypeArgumentToTypeVisitor) VisitInnerObjectType(typ intmod.IInnerObjectType) {
	v.typ = typ
}

func (v *TypeArgumentToTypeVisitor) VisitGenericType(typ intmod.IGenericType) {
	v.typ = typ
}

func (v *TypeArgumentToTypeVisitor) VisitWildcardExtendsTypeArgument(argument intmod.IWildcardExtendsTypeArgument) {
	argument.Type().AcceptTypeVisitor(v)
}

func (v *TypeArgumentToTypeVisitor) VisitWildcardSuperTypeArgument(argument intmod.IWildcardSuperTypeArgument) {
	argument.Type().AcceptTypeVisitor(v)
}

func (v *TypeArgumentToTypeVisitor) VisitTypeArguments(arguments intmod.ITypeArguments) {
	if arguments.IsEmpty() {
		v.typ = _type.OtTypeUndefinedObject
	} else {
		arguments.First().AcceptTypeArgumentVisitor(v)
	}
}
