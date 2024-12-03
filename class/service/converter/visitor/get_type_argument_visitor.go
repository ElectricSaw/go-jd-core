package visitor

import intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"

func NewGetTypeArgumentVisitor() *GetTypeArgumentVisitor {
	return &GetTypeArgumentVisitor{}
}

type GetTypeArgumentVisitor struct {
	typeArguments intmod.ITypeArgument
}

func (v *GetTypeArgumentVisitor) Init() {
	v.typeArguments = nil
}

func (v *GetTypeArgumentVisitor) TypeArguments() intmod.ITypeArgument {
	return v.typeArguments
}

func (v *GetTypeArgumentVisitor) VisitObjectType(t intmod.IObjectType) {
	v.typeArguments = t.TypeArguments()
}

func (v *GetTypeArgumentVisitor) VisitInnerObjectType(t intmod.IInnerObjectType) {
	v.typeArguments = t.TypeArguments()
}

func (v *GetTypeArgumentVisitor) VisitPrimitiveType(_ intmod.IPrimitiveType) {
	v.typeArguments = nil
}

func (v *GetTypeArgumentVisitor) VisitGenericType(_ intmod.IGenericType) {
	v.typeArguments = nil
}

func (v *GetTypeArgumentVisitor) VisitTypes(_ intmod.ITypes) {
	v.typeArguments = nil
}
