package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax"
)

func NewSearchLocalVariableVisitor() intsrv.ISearchLocalVariableVisitor {
	return &SearchLocalVariableVisitor{
		variables: make([]intsrv.ILocalVariable, 0),
	}
}

type SearchLocalVariableVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	variables []intsrv.ILocalVariable
}

func (v *SearchLocalVariableVisitor) Init() {
	if v.variables == nil {
		v.variables = make([]intsrv.ILocalVariable, 0)
	}

	v.variables = v.variables[:0]
}

func (v *SearchLocalVariableVisitor) Variables() []intsrv.ILocalVariable {
	return v.variables
}

func (v *SearchLocalVariableVisitor) VisitLocalVariableReferenceExpression(expression intmod.ILocalVariableReferenceExpression) {
	lv := expression.(intsrv.IClassFileLocalVariableReferenceExpression).LocalVariable().(intsrv.ILocalVariable)

	if !lv.IsDeclared() {
		v.variables = append(v.variables, lv)
	}
}

func (v *SearchLocalVariableVisitor) VisitIntegerConstantExpression(_ intmod.IIntegerConstantExpression) {
}

func (v *SearchLocalVariableVisitor) VisitTypeArguments(_ intmod.ITypeArguments)             {}
func (v *SearchLocalVariableVisitor) VisitDiamondTypeArgument(_ intmod.IDiamondTypeArgument) {}
func (v *SearchLocalVariableVisitor) VisitWildcardExtendsTypeArgument(_ intmod.IWildcardExtendsTypeArgument) {
}
func (v *SearchLocalVariableVisitor) VisitWildcardSuperTypeArgument(_ intmod.IWildcardSuperTypeArgument) {
}
func (v *SearchLocalVariableVisitor) VisitWildcardTypeArgument(_ intmod.IWildcardTypeArgument) {
}
func (v *SearchLocalVariableVisitor) VisitPrimitiveType(_ intmod.IPrimitiveType)     {}
func (v *SearchLocalVariableVisitor) VisitObjectType(_ intmod.IObjectType)           {}
func (v *SearchLocalVariableVisitor) VisitInnerObjectType(_ intmod.IInnerObjectType) {}
func (v *SearchLocalVariableVisitor) VisitGenericType(_ intmod.IGenericType)         {}
