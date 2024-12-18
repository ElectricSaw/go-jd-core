package _type

import intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"

type AbstractTypeArgumentVisitor struct {
}

func (v *AbstractTypeArgumentVisitor) VisitTypeArguments(arguments intmod.ITypeArguments) {
	arguments.AcceptTypeArgumentVisitor(v)
}

func (v *AbstractTypeArgumentVisitor) VisitDiamondTypeArgument(argument intmod.IDiamondTypeArgument) {
}

func (v *AbstractTypeArgumentVisitor) VisitWildcardExtendsTypeArgument(argument intmod.IWildcardExtendsTypeArgument) {
	argument.AcceptTypeArgumentVisitor(v)
}

func (v *AbstractTypeArgumentVisitor) VisitWildcardSuperTypeArgument(argument intmod.IWildcardSuperTypeArgument) {
}

func (v *AbstractTypeArgumentVisitor) VisitWildcardTypeArgument(argument intmod.IWildcardTypeArgument) {
}

func (v *AbstractTypeArgumentVisitor) VisitPrimitiveType(t intmod.IPrimitiveType) {
}

func (v *AbstractTypeArgumentVisitor) VisitObjectType(t intmod.IObjectType) {
	if visitable, ok := t.TypeArguments().(intmod.ITypeArgumentVisitable); ok {
		v.SafeAcceptTypeArgumentVisitable(visitable)
	}
}

func (v *AbstractTypeArgumentVisitor) VisitInnerObjectType(t intmod.IInnerObjectType) {
	t.OuterType().AcceptTypeArgumentVisitor(v)
	if visitable, ok := t.TypeArguments().(intmod.ITypeArgumentVisitable); ok {
		v.SafeAcceptTypeArgumentVisitable(visitable)
	}
}

func (v *AbstractTypeArgumentVisitor) VisitGenericType(t intmod.IGenericType) {
}

func (v *AbstractTypeArgumentVisitor) SafeAcceptTypeArgumentVisitable(visitable intmod.ITypeArgumentVisitable) {
	if visitable != nil {
		visitable.AcceptTypeArgumentVisitor(v)
	}
}

type AbstractNopTypeArgumentVisitor struct {
}

func (v *AbstractNopTypeArgumentVisitor) VisitTypeArguments(arguments intmod.ITypeArguments) {}
func (v *AbstractNopTypeArgumentVisitor) VisitDiamondTypeArgument(argument intmod.IDiamondTypeArgument) {
}
func (v *AbstractNopTypeArgumentVisitor) VisitWildcardExtendsTypeArgument(argument intmod.IWildcardExtendsTypeArgument) {
}
func (v *AbstractNopTypeArgumentVisitor) VisitWildcardSuperTypeArgument(argument intmod.IWildcardSuperTypeArgument) {
}
func (v *AbstractNopTypeArgumentVisitor) VisitWildcardTypeArgument(argument intmod.IWildcardTypeArgument) {
}
func (v *AbstractNopTypeArgumentVisitor) VisitPrimitiveType(t intmod.IPrimitiveType)     {}
func (v *AbstractNopTypeArgumentVisitor) VisitObjectType(t intmod.IObjectType)           {}
func (v *AbstractNopTypeArgumentVisitor) VisitInnerObjectType(t intmod.IInnerObjectType) {}
func (v *AbstractNopTypeArgumentVisitor) VisitGenericType(t intmod.IGenericType)         {}
