package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax"
	"github.com/ElectricSaw/go-jd-core/class/util"
)

func NewDeclaredSyntheticLocalVariableVisitor() intsrv.IDeclaredSyntheticLocalVariableVisitor {
	return &DeclaredSyntheticLocalVariableVisitor{
		localVariableReferenceExpressions: util.NewDefaultList[intmod.ILocalVariableReferenceExpression](),
	}
}

type DeclaredSyntheticLocalVariableVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	localVariableReferenceExpressions util.IList[intmod.ILocalVariableReferenceExpression]
}

func (v *DeclaredSyntheticLocalVariableVisitor) Init() {
	v.localVariableReferenceExpressions.Clear()
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitFieldDeclaration(decl intmod.IFieldDeclaration) {
	v.SafeAcceptReference(decl.AnnotationReferences())
	decl.FieldDeclarators().AcceptDeclaration(v)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitFormalParameter(decl intmod.IFormalParameter) {
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitLocalVariableDeclaration(decl intmod.ILocalVariableDeclaration) {
	decl.LocalVariableDeclarators().AcceptDeclaration(v)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitMethodDeclaration(decl intmod.IMethodDeclaration) {
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.FormalParameters())
	v.SafeAcceptStatement(decl.Statements())
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitArrayExpression(expr intmod.IArrayExpression) {
	expr.Expression().Accept(v)
	expr.Index().Accept(v)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitCastExpression(expr intmod.ICastExpression) {
	expr.Expression().Accept(v)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitFieldReferenceExpression(expr intmod.IFieldReferenceExpression) {
	v.SafeAcceptExpression(expr.Expression())
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitInstanceOfExpression(expr intmod.IInstanceOfExpression) {
	expr.Expression().Accept(v)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitLocalVariableReferenceExpression(expr intmod.ILocalVariableReferenceExpression) {
	localVariable := expr.(intsrv.IClassFileLocalVariableReferenceExpression).LocalVariable().(intsrv.ILocalVariable)
	v.localVariableReferenceExpressions.Add(expr)

	tmp := make([]intmod.ILocalVariableReferenceExpression, 0)
	for _, item := range localVariable.References() {
		tmp = append(tmp, item.(intmod.ILocalVariableReferenceExpression))
	}

	if v.localVariableReferenceExpressions.ContainsAll(tmp) {
		localVariable.SetDeclared(true)
	}
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitMethodReferenceExpression(expr intmod.IMethodReferenceExpression) {
	expr.Expression().Accept(v)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitNewArray(expr intmod.INewArray) {
	v.SafeAcceptExpression(expr.DimensionExpressionList())
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitNewExpression(expr intmod.INewExpression) {
	v.SafeAcceptExpression(expr.Parameters())
	v.SafeAcceptDeclaration(expr.BodyDeclaration())
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitNewInitializedArray(expr intmod.INewInitializedArray) {
	tmp := make([]intmod.IDeclaration, 0)
	for _, item := range expr.ArrayInitializer().ToSlice() {
		tmp = append(tmp, item.(intmod.IDeclaration))
	}
	v.SafeAcceptListDeclaration(tmp)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitForEachStatement(expr intmod.IForEachStatement) {
	expr.Expression().Accept(v)
	v.SafeAcceptStatement(expr.Statements())
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitSwitchStatementLabelBlock(state intmod.ILabelBlock) {
	state.Statements().AcceptStatement(v)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitSwitchStatementMultiLabelsBlock(state intmod.IMultiLabelsBlock) {
	state.Statements().AcceptStatement(v)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitTryStatementCatchClause(state intmod.ICatchClause) {
	v.SafeAcceptStatement(state.Statements())
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitTryStatementResource(state intmod.IResource) {
	state.Expression().Accept(v)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitConstructorReferenceExpression(_ intmod.IConstructorReferenceExpression) {
}
func (v *DeclaredSyntheticLocalVariableVisitor) VisitDoubleConstantExpression(_ intmod.IDoubleConstantExpression) {
}
func (v *DeclaredSyntheticLocalVariableVisitor) VisitFloatConstantExpression(_ intmod.IFloatConstantExpression) {
}
func (v *DeclaredSyntheticLocalVariableVisitor) VisitIntegerConstantExpression(_ intmod.IIntegerConstantExpression) {
}
func (v *DeclaredSyntheticLocalVariableVisitor) VisitLongConstantExpression(_ intmod.ILongConstantExpression) {
}
func (v *DeclaredSyntheticLocalVariableVisitor) VisitNullExpression(_ intmod.INullExpression) {}
func (v *DeclaredSyntheticLocalVariableVisitor) VisitObjectTypeReferenceExpression(_ intmod.IObjectTypeReferenceExpression) {
}
func (v *DeclaredSyntheticLocalVariableVisitor) VisitThisExpression(_ intmod.IThisExpression) {}
func (v *DeclaredSyntheticLocalVariableVisitor) VisitTypeReferenceDotClassExpression(_ intmod.ITypeReferenceDotClassExpression) {
}
