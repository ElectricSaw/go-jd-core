package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
)

func NewCreateTypeFromTypeArgumentVisitor() *CreateTypeFromTypeArgumentVisitor {
	return &CreateTypeFromTypeArgumentVisitor{}
}

type CreateTypeFromTypeArgumentVisitor struct {
	typ intmod.IType
}

func (v *CreateTypeFromTypeArgumentVisitor) Init() {
	v.typ = nil
}

func (v *CreateTypeFromTypeArgumentVisitor) Type() intmod.IType {
	return v.typ
}

func (v *CreateTypeFromTypeArgumentVisitor) VisitTypeArguments(_ intmod.ITypeArguments) {
	v.typ = nil
}
func (v *CreateTypeFromTypeArgumentVisitor) VisitDiamondTypeArgument(_ intmod.IDiamondTypeArgument) {
	v.typ = nil
}
func (v *CreateTypeFromTypeArgumentVisitor) VisitWildcardTypeArgument(_ intmod.IWildcardTypeArgument) {
	v.typ = nil
}
func (v *CreateTypeFromTypeArgumentVisitor) VisitWildcardExtendsTypeArgument(typ intmod.IWildcardExtendsTypeArgument) {
	v.typ = typ.Type()
}
func (v *CreateTypeFromTypeArgumentVisitor) VisitWildcardSuperTypeArgument(typ intmod.IWildcardSuperTypeArgument) {
	v.typ = typ.Type()
}
func (v *CreateTypeFromTypeArgumentVisitor) VisitPrimitiveType(typ intmod.IPrimitiveType) {
	v.typ = typ
}
func (v *CreateTypeFromTypeArgumentVisitor) VisitObjectType(typ intmod.IObjectType) {
	v.typ = typ
}
func (v *CreateTypeFromTypeArgumentVisitor) VisitInnerObjectType(typ intmod.IInnerObjectType) {
	v.typ = typ
}
func (v *CreateTypeFromTypeArgumentVisitor) VisitGenericType(typ intmod.IGenericType) {
	v.typ = typ
}
