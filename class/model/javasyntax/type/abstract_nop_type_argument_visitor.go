package _type

import intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"

type AbstractTypeArgumentVisitor struct {
}

func (v *AbstractTypeArgumentVisitor) VisitTypeArguments(arguments intsyn.ITypeArguments) {
	arguments.AcceptTypeArgumentVisitor(v)
}

func (v *AbstractTypeArgumentVisitor) VisitDiamondTypeArgument(argument intsyn.IDiamondTypeArgument) {
}

func (v *AbstractTypeArgumentVisitor) VisitWildcardExtendsTypeArgument(argument intsyn.IWildcardExtendsTypeArgument) {
	argument.AcceptTypeArgumentVisitor(v)
}

func (v *AbstractTypeArgumentVisitor) VisitWildcardSuperTypeArgument(argument intsyn.IWildcardSuperTypeArgument) {
}

func (v *AbstractTypeArgumentVisitor) VisitWildcardTypeArgument(argument intsyn.IWildcardTypeArgument) {
}

func (v *AbstractTypeArgumentVisitor) VisitPrimitiveType(t intsyn.IPrimitiveType) {
}

func (v *AbstractTypeArgumentVisitor) VisitObjectType(t intsyn.IObjectType) {
	if visitable, ok := t.TypeArguments().(intsyn.ITypeArgumentVisitable); ok {
		v.SafeAcceptTypeArgumentVisitable(visitable)
	}
}

func (v *AbstractTypeArgumentVisitor) VisitInnerObjectType(t intsyn.IInnerObjectType) {
	t.OuterType().AcceptTypeArgumentVisitor(v)
	if visitable, ok := t.TypeArguments().(intsyn.ITypeArgumentVisitable); ok {
		v.SafeAcceptTypeArgumentVisitable(visitable)
	}
}

func (v *AbstractTypeArgumentVisitor) VisitGenericType(t intsyn.IGenericType) {
}

func (v *AbstractTypeArgumentVisitor) SafeAcceptTypeArgumentVisitable(visitable intsyn.ITypeArgumentVisitable) {
	if visitable != nil {
		visitable.AcceptTypeArgumentVisitor(v)
	}
}

type AbstractNopTypeArgumentVisitor struct {
}

func (v *AbstractNopTypeArgumentVisitor) VisitTypeArguments(arguments intsyn.ITypeArguments) {}
func (v *AbstractNopTypeArgumentVisitor) VisitDiamondTypeArgument(argument intsyn.IDiamondTypeArgument) {
}
func (v *AbstractNopTypeArgumentVisitor) VisitWildcardExtendsTypeArgument(argument intsyn.IWildcardExtendsTypeArgument) {
}
func (v *AbstractNopTypeArgumentVisitor) VisitWildcardSuperTypeArgument(argument intsyn.IWildcardSuperTypeArgument) {
}
func (v *AbstractNopTypeArgumentVisitor) VisitWildcardTypeArgument(argument intsyn.IWildcardTypeArgument) {
}
func (v *AbstractNopTypeArgumentVisitor) VisitPrimitiveType(t intsyn.IPrimitiveType)     {}
func (v *AbstractNopTypeArgumentVisitor) VisitObjectType(t intsyn.IObjectType)           {}
func (v *AbstractNopTypeArgumentVisitor) VisitInnerObjectType(t intsyn.IInnerObjectType) {}
func (v *AbstractNopTypeArgumentVisitor) VisitGenericType(t intsyn.IGenericType)         {}
