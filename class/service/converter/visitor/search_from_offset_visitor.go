package visitor

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
	"math"
)

func NewSearchFromOffsetVisitor() *SearchFromOffsetVisitor {
	return &SearchFromOffsetVisitor{offset: math.MaxInt}
}

type SearchFromOffsetVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	offset int
}

func (v *SearchFromOffsetVisitor) Init() {
	v.offset = math.MaxInt
}

func (v *SearchFromOffsetVisitor) Offset() int {
	return v.offset
}

func (v *SearchFromOffsetVisitor) VisitLocalVariableReferenceExpression(expression intmod.ILocalVariableReferenceExpression) {
	offset := expression.(intsrv.IClassFileLocalVariableReferenceExpression).Offset()

	if v.offset > offset {
		v.offset = offset
	}
}

func (v *SearchFromOffsetVisitor) VisitIntegerConstantExpression(expression intmod.IIntegerConstantExpression) {
}

func (v *SearchFromOffsetVisitor) VisitTypeArguments(arguments intmod.ITypeArguments)            {}
func (v *SearchFromOffsetVisitor) VisitDiamondTypeArgument(argument intmod.IDiamondTypeArgument) {}
func (v *SearchFromOffsetVisitor) VisitWildcardExtendsTypeArgument(argument intmod.IWildcardExtendsTypeArgument) {
}
func (v *SearchFromOffsetVisitor) VisitWildcardSuperTypeArgument(argument intmod.IWildcardSuperTypeArgument) {
}
func (v *SearchFromOffsetVisitor) VisitWildcardTypeArgument(argument intmod.IWildcardTypeArgument) {}
func (v *SearchFromOffsetVisitor) VisitPrimitiveType(typ intmod.IPrimitiveType)                    {}
func (v *SearchFromOffsetVisitor) VisitObjectType(typ intmod.IObjectType)                          {}
func (v *SearchFromOffsetVisitor) VisitInnerObjectType(typ intmod.IInnerObjectType)                {}
func (v *SearchFromOffsetVisitor) VisitGenericType(typ intmod.IGenericType)                        {}
