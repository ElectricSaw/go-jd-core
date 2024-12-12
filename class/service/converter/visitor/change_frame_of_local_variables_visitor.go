package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax"
)

func NewChangeFrameOfLocalVariablesVisitor(localVariableMaker intsrv.ILocalVariableMaker) *ChangeFrameOfLocalVariablesVisitor {
	return &ChangeFrameOfLocalVariablesVisitor{
		localVariableMaker: localVariableMaker,
	}
}

type ChangeFrameOfLocalVariablesVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor
	
	localVariableMaker intsrv.ILocalVariableMaker
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitLocalVariableReferenceExpression(expr intmod.ILocalVariableReferenceExpression) {
	localVariable := expr.(intsrv.IClassFileLocalVariableReferenceExpression).
		LocalVariable().(intsrv.ILocalVariable)
	v.localVariableMaker.ChangeFrame(localVariable)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitFloatConstantExpression(_ intmod.IFloatConstantExpression) {
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitIntegerConstantExpression(_ intmod.IIntegerConstantExpression) {
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitConstructorReferenceExpression(_ intmod.IConstructorReferenceExpression) {
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitDoubleConstantExpression(_ intmod.IDoubleConstantExpression) {
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitEnumConstantReferenceExpression(_ intmod.IEnumConstantReferenceExpression) {
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitLongConstantExpression(_ intmod.ILongConstantExpression) {
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitBreakStatement(_ intmod.IBreakStatement) {
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitByteCodeStatement(_ intmod.IByteCodeStatement) {
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitContinueStatement(_ intmod.IContinueStatement) {
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitNullExpression(_ intmod.INullExpression) {
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitObjectTypeReferenceExpression(_ intmod.IObjectTypeReferenceExpression) {
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitSuperExpression(_ intmod.ISuperExpression) {
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitThisExpression(_ intmod.IThisExpression) {
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitTypeReferenceDotClassExpression(_ intmod.ITypeReferenceDotClassExpression) {
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitObjectReference(_ intmod.IObjectReference) {
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitInnerObjectReference(_ intmod.IInnerObjectReference) {
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitTypeArguments(_ intmod.ITypeArguments) {
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitWildcardExtendsTypeArgument(_ intmod.IWildcardExtendsTypeArgument) {
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitObjectType(_ intmod.IObjectType) {
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitInnerObjectType(_ intmod.IInnerObjectType) {
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitWildcardSuperTypeArgument(_ intmod.IWildcardSuperTypeArgument) {
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitTypes(_ intmod.ITypes) {
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitTypeParameterWithTypeBounds(_ intmod.ITypeParameterWithTypeBounds) {
}
