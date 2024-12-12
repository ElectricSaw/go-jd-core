package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax"
	"math"
)

func NewSearchFromOffsetVisitor() intsrv.ISearchFromOffsetVisitor {
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

func (v *SearchFromOffsetVisitor) VisitIntegerConstantExpression(_ intmod.IIntegerConstantExpression) {
}

func (v *SearchFromOffsetVisitor) VisitTypeArguments(_ intmod.ITypeArguments)             {}
func (v *SearchFromOffsetVisitor) VisitDiamondTypeArgument(_ intmod.IDiamondTypeArgument) {}
func (v *SearchFromOffsetVisitor) VisitWildcardExtendsTypeArgument(_ intmod.IWildcardExtendsTypeArgument) {
}
func (v *SearchFromOffsetVisitor) VisitWildcardSuperTypeArgument(_ intmod.IWildcardSuperTypeArgument) {
}
func (v *SearchFromOffsetVisitor) VisitWildcardTypeArgument(_ intmod.IWildcardTypeArgument) {}
func (v *SearchFromOffsetVisitor) VisitPrimitiveType(_ intmod.IPrimitiveType)               {}
func (v *SearchFromOffsetVisitor) VisitObjectType(_ intmod.IObjectType)                     {}
func (v *SearchFromOffsetVisitor) VisitInnerObjectType(_ intmod.IInnerObjectType)           {}
func (v *SearchFromOffsetVisitor) VisitGenericType(_ intmod.IGenericType)                   {}
