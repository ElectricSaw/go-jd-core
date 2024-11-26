package visitor

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
)

func NewSearchInTypeArgumentVisitor() *SearchInTypeArgumentVisitor {
	return &SearchInTypeArgumentVisitor{
		wildcardFound:                   false,
		wildcardSuperOrExtendsTypeFound: false,
		genericFound:                    false,
	}
}

type SearchInTypeArgumentVisitor struct {
	_type.AbstractTypeArgumentVisitor

	wildcardFound                   bool
	wildcardSuperOrExtendsTypeFound bool
	genericFound                    bool
}

func (v *SearchInTypeArgumentVisitor) Init() {
	v.wildcardFound = false
	v.wildcardSuperOrExtendsTypeFound = false
	v.genericFound = false
}

func (v *SearchInTypeArgumentVisitor) ContainsWildcard() bool {
	return v.wildcardFound
}

func (v *SearchInTypeArgumentVisitor) ContainsWildcardSuperOrExtendsType() bool {
	return v.wildcardSuperOrExtendsTypeFound
}

func (v *SearchInTypeArgumentVisitor) ContainsGeneric() bool {
	return v.genericFound
}

func (v *SearchInTypeArgumentVisitor) VisitWildcardTypeArgument(_ intmod.IWildcardTypeArgument) {
	v.wildcardFound = true
}

func (v *SearchInTypeArgumentVisitor) VisitWildcardExtendsTypeArgument(typ intmod.IWildcardExtendsTypeArgument) {
	v.wildcardSuperOrExtendsTypeFound = true
	typ.Type().AcceptTypeArgumentVisitor(v)
}

func (v *SearchInTypeArgumentVisitor) VisitWildcardSuperTypeArgument(typ intmod.IWildcardSuperTypeArgument) {
	v.wildcardSuperOrExtendsTypeFound = true
	typ.Type().AcceptTypeArgumentVisitor(v)
}

func (v *SearchInTypeArgumentVisitor) VisitGenericType(_ intmod.IGenericType) {
	v.genericFound = true
}
