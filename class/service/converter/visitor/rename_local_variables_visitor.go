package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax"
)

func NewRenameLocalVariablesVisitor() intsrv.IRenameLocalVariablesVisitor {
	return &RenameLocalVariablesVisitor{}
}

type RenameLocalVariablesVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	nameMapping map[string]string
}

func (v *RenameLocalVariablesVisitor) Init(nameMapping map[string]string) {
	v.nameMapping = nameMapping
}

func (v *RenameLocalVariablesVisitor) VisitLocalVariableReferenceExpression(expr intmod.ILocalVariableReferenceExpression) {
	lvre := expr.(intsrv.IClassFileLocalVariableReferenceExpression)
	newName := v.nameMapping[lvre.Name()]

	if newName != "" {
		lvre.LocalVariable().(intsrv.ILocalVariable).SetName(newName)
	}
}

func (v *RenameLocalVariablesVisitor) VisitFloatConstantExpression(_ intmod.IFloatConstantExpression) {
}

func (v *RenameLocalVariablesVisitor) VisitIntegerConstantExpression(_ intmod.IIntegerConstantExpression) {
}

func (v *RenameLocalVariablesVisitor) VisitConstructorReferenceExpression(_ intmod.IConstructorReferenceExpression) {
}

func (v *RenameLocalVariablesVisitor) VisitDoubleConstantExpression(_ intmod.IDoubleConstantExpression) {
}

func (v *RenameLocalVariablesVisitor) VisitEnumConstantReferenceExpression(_ intmod.IEnumConstantReferenceExpression) {
}

func (v *RenameLocalVariablesVisitor) VisitLongConstantExpression(_ intmod.ILongConstantExpression) {
}

func (v *RenameLocalVariablesVisitor) VisitBreakStatement(_ intmod.IBreakStatement) {
}

func (v *RenameLocalVariablesVisitor) VisitByteCodeStatement(_ intmod.IByteCodeStatement) {
}

func (v *RenameLocalVariablesVisitor) VisitContinueStatement(_ intmod.IContinueStatement) {
}

func (v *RenameLocalVariablesVisitor) VisitNullExpression(_ intmod.INullExpression) {
}

func (v *RenameLocalVariablesVisitor) VisitObjectTypeReferenceExpression(_ intmod.IObjectTypeReferenceExpression) {
}

func (v *RenameLocalVariablesVisitor) VisitSuperExpression(_ intmod.ISuperExpression) {
}

func (v *RenameLocalVariablesVisitor) VisitThisExpression(_ intmod.IThisExpression) {
}

func (v *RenameLocalVariablesVisitor) VisitTypeReferenceDotClassExpression(_ intmod.ITypeReferenceDotClassExpression) {
}

func (v *RenameLocalVariablesVisitor) VisitObjectReference(_ intmod.IObjectReference) {
}

func (v *RenameLocalVariablesVisitor) VisitInnerObjectReference(_ intmod.IInnerObjectReference) {
}

func (v *RenameLocalVariablesVisitor) VisitTypeArguments(_ intmod.ITypeArguments) {
}

func (v *RenameLocalVariablesVisitor) VisitWildcardExtendsTypeArgument(_ intmod.IWildcardExtendsTypeArgument) {
}

func (v *RenameLocalVariablesVisitor) VisitObjectType(_ intmod.IObjectType) {
}

func (v *RenameLocalVariablesVisitor) VisitInnerObjectType(_ intmod.IInnerObjectType) {
}

func (v *RenameLocalVariablesVisitor) VisitWildcardSuperTypeArgument(_ intmod.IWildcardSuperTypeArgument) {
}

func (v *RenameLocalVariablesVisitor) VisitTypes(_ intmod.ITypes) {
}

func (v *RenameLocalVariablesVisitor) VisitTypeParameterWithTypeBounds(_ intmod.ITypeParameterWithTypeBounds) {
}
