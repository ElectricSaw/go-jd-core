package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax"
)

type AbstractUpdateExpressionVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor
}

func (v *AbstractUpdateExpressionVisitor) UpdateExpression(_ intmod.IExpression) intmod.IExpression {
	return nil
}

func (v *AbstractUpdateExpressionVisitor) UpdateBaseExpression(baseExpression intmod.IExpression) intmod.IExpression {
	if baseExpression == nil {
		return nil
	}

	if baseExpression.IsList() {
		iterator := baseExpression.ToList().ListIterator()

		for iterator.HasNext() {
			value := iterator.Next()
			_ = iterator.Set(v.UpdateExpression(value))
		}

		return baseExpression
	}

	return v.UpdateExpression(baseExpression.First())
}

func (v *AbstractUpdateExpressionVisitor) VisitAnnotationDeclaration(decl intmod.IAnnotationDeclaration) {
	v.SafeAcceptDeclaration(decl.AnnotationDeclarators())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *AbstractUpdateExpressionVisitor) VisitClassDeclaration(declaration intmod.IClassDeclaration) {
	v.VisitInterfaceDeclaration(declaration)
}

func (v *AbstractUpdateExpressionVisitor) VisitConstructorInvocationExpression(expression intmod.IConstructorInvocationExpression) {
	if expression.Parameters() != nil {
		expression.SetParameters(v.UpdateBaseExpression(expression.Parameters()))
		expression.Parameters().Accept(v)
	}
}

func (v *AbstractUpdateExpressionVisitor) VisitConstructorDeclaration(declaration intmod.IConstructorDeclaration) {
	v.SafeAcceptStatement(declaration.Statements())
}

func (v *AbstractUpdateExpressionVisitor) VisitEnumDeclaration(declaration intmod.IEnumDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *AbstractUpdateExpressionVisitor) VisitEnumDeclarationConstant(declaration intmod.IConstant) {
	if declaration.Arguments() != nil {
		declaration.SetArguments(v.UpdateBaseExpression(declaration.Arguments()))
		declaration.Arguments().Accept(v)
	}
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *AbstractUpdateExpressionVisitor) VisitExpressionVariableInitializer(declaration intmod.IExpressionVariableInitializer) {
	if declaration.Expression() != nil {
		declaration.SetExpression(v.UpdateExpression(declaration.Expression()))
		declaration.Expression().Accept(v)
	}
}

func (v *AbstractUpdateExpressionVisitor) VisitFieldDeclaration(declaration intmod.IFieldDeclaration) {
	declaration.FieldDeclarators().AcceptDeclaration(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitFieldDeclarator(declaration intmod.IFieldDeclarator) {
	v.SafeAcceptDeclaration(declaration.VariableInitializer())
}

func (v *AbstractUpdateExpressionVisitor) VisitFormalParameter(_ intmod.IFormalParameter) {
}

func (v *AbstractUpdateExpressionVisitor) VisitInterfaceDeclaration(declaration intmod.IInterfaceDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *AbstractUpdateExpressionVisitor) VisitLocalVariableDeclaration(declaration intmod.ILocalVariableDeclaration) {
	declaration.LocalVariableDeclarators().AcceptDeclaration(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitLocalVariableDeclarator(declarator intmod.ILocalVariableDeclarator) {
	v.SafeAcceptDeclaration(declarator.VariableInitializer())
}

func (v *AbstractUpdateExpressionVisitor) VisitMethodDeclaration(declaration intmod.IMethodDeclaration) {
	v.SafeAcceptReference(declaration.AnnotationReferences())
	v.SafeAcceptStatement(declaration.Statements())
}

func (v *AbstractUpdateExpressionVisitor) VisitArrayExpression(expression intmod.IArrayExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.SetIndex(v.UpdateExpression(expression.Index()))
	expression.Expression().Accept(v)
	expression.Index().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitBinaryOperatorExpression(expression intmod.IBinaryOperatorExpression) {
	expression.SetLeftExpression(v.UpdateExpression(expression.LeftExpression()))
	expression.SetRightExpression(v.UpdateExpression(expression.RightExpression()))
	expression.LeftExpression().Accept(v)
	expression.RightExpression().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitCastExpression(expression intmod.ICastExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.Expression().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitFieldReferenceExpression(expression intmod.IFieldReferenceExpression) {
	if expression.Expression() != nil {
		expression.SetExpression(v.UpdateExpression(expression.Expression()))
		expression.Expression().Accept(v)
	}
}

func (v *AbstractUpdateExpressionVisitor) VisitInstanceOfExpression(expression intmod.IInstanceOfExpression) {
	if expression.Expression() != nil {
		expression.SetExpression(v.UpdateExpression(expression.Expression()))
		expression.Expression().Accept(v)
	}
}

func (v *AbstractUpdateExpressionVisitor) VisitLambdaFormalParametersExpression(expression intmod.ILambdaFormalParametersExpression) {
	expression.Statements().AcceptStatement(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitLengthExpression(expression intmod.ILengthExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.Expression().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitMethodInvocationExpression(expression intmod.IMethodInvocationExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	if expression.Parameters() != nil {
		expression.SetParameters(v.UpdateBaseExpression(expression.Parameters()))
		expression.Parameters().Accept(v)
	}
	expression.Expression().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitMethodReferenceExpression(expression intmod.IMethodReferenceExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.Expression().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitNewArray(expression intmod.INewArray) {
	if expression.DimensionExpressionList() != nil {
		expression.SetDimensionExpressionList(v.UpdateBaseExpression(expression.DimensionExpressionList()))
		expression.DimensionExpressionList().Accept(v)
	}
}

func (v *AbstractUpdateExpressionVisitor) VisitNewExpression(expression intmod.INewExpression) {
	if expression.Parameters() != nil {
		expression.SetParameters(v.UpdateBaseExpression(expression.Parameters()))
		expression.Parameters().Accept(v)
	}
	// v.SafeAccept(expression.BodyDeclaration());
}

func (v *AbstractUpdateExpressionVisitor) VisitNewInitializedArray(expression intmod.INewInitializedArray) {
	v.SafeAcceptDeclaration(expression.ArrayInitializer())
}

func (v *AbstractUpdateExpressionVisitor) VisitParenthesesExpression(expression intmod.IParenthesesExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.Expression().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitPostOperatorExpression(expression intmod.IPostOperatorExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.Expression().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitPreOperatorExpression(expression intmod.IPreOperatorExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.Expression().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitSuperConstructorInvocationExpression(expression intmod.ISuperConstructorInvocationExpression) {
	if expression.Parameters() != nil {
		expression.SetParameters(v.UpdateBaseExpression(expression.Parameters()))
		expression.Parameters().Accept(v)
	}
}

func (v *AbstractUpdateExpressionVisitor) VisitTernaryOperatorExpression(expression intmod.ITernaryOperatorExpression) {
	expression.SetCondition(v.UpdateExpression(expression.Condition()))
	expression.SetTrueExpression(v.UpdateExpression(expression.TrueExpression()))
	expression.SetFalseExpression(v.UpdateExpression(expression.FalseExpression()))
	expression.Condition().Accept(v)
	expression.TrueExpression().Accept(v)
	expression.FalseExpression().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitExpressionElementValue(reference intmod.IExpressionElementValue) {
	reference.SetExpression(v.UpdateExpression(reference.Expression()))
	reference.Expression().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitAssertStatement(statement intmod.IAssertStatement) {
	statement.SetCondition(v.UpdateExpression(statement.Condition()))
	statement.Condition().Accept(v)
	v.SafeAcceptExpression(statement.Message())
}

func (v *AbstractUpdateExpressionVisitor) VisitDoWhileStatement(statement intmod.IDoWhileStatement) {
	statement.SetCondition(v.UpdateExpression(statement.Condition()))
	v.SafeAcceptExpression(statement.Condition())
	v.SafeAcceptStatement(statement.Statements())
}

func (v *AbstractUpdateExpressionVisitor) VisitExpressionStatement(statement intmod.IExpressionStatement) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitForEachStatement(statement intmod.IForEachStatement) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
	v.SafeAcceptStatement(statement.Statements())
}

func (v *AbstractUpdateExpressionVisitor) VisitForStatement(statement intmod.IForStatement) {
	v.SafeAcceptDeclaration(statement.Declaration())
	if statement.Init() != nil {
		statement.SetInit(v.UpdateBaseExpression(statement.Init()))
		statement.Init().Accept(v)
	}
	if statement.Condition() != nil {
		statement.SetCondition(v.UpdateExpression(statement.Condition()))
		statement.Condition().Accept(v)
	}
	if statement.Update() != nil {
		statement.SetUpdate(v.UpdateBaseExpression(statement.Update()))
		statement.Update().Accept(v)
	}
	v.SafeAcceptStatement(statement.Statements())
}

func (v *AbstractUpdateExpressionVisitor) VisitIfStatement(statement intmod.IIfStatement) {
	statement.SetCondition(v.UpdateExpression(statement.Condition()))
	statement.Condition().Accept(v)
	v.SafeAcceptStatement(statement.Statements())
}

func (v *AbstractUpdateExpressionVisitor) VisitIfElseStatement(statement intmod.IIfElseStatement) {
	statement.SetCondition(v.UpdateExpression(statement.Condition()))
	statement.Condition().Accept(v)
	v.SafeAcceptStatement(statement.Statements())
	statement.ElseStatements().AcceptStatement(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitLambdaExpressionStatement(statement intmod.ILambdaExpressionStatement) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitReturnExpressionStatement(statement intmod.IReturnExpressionStatement) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitSwitchStatement(statement intmod.ISwitchStatement) {
	statement.SetCondition(v.UpdateExpression(statement.Condition()))
	statement.Condition().Accept(v)
	v.AcceptListStatement(statement.List())
}

func (v *AbstractUpdateExpressionVisitor) VisitSwitchStatementExpressionLabel(statement intmod.IExpressionLabel) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitSynchronizedStatement(statement intmod.ISynchronizedStatement) {
	statement.SetMonitor(v.UpdateExpression(statement.Monitor()))
	statement.Monitor().Accept(v)
	v.SafeAcceptStatement(statement.Statements())
}

func (v *AbstractUpdateExpressionVisitor) VisitThrowStatement(statement intmod.IThrowStatement) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitTryStatementCatchClause(statement intmod.ICatchClause) {
	v.SafeAcceptStatement(statement.Statements())
}

func (v *AbstractUpdateExpressionVisitor) VisitTryStatementResource(statement intmod.IResource) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitWhileStatement(statement intmod.IWhileStatement) {
	statement.SetCondition(v.UpdateExpression(statement.Condition()))
	statement.Condition().Accept(v)
	v.SafeAcceptStatement(statement.Statements())
}

func (v *AbstractUpdateExpressionVisitor) VisitConstructorReferenceExpression(_ intmod.IConstructorReferenceExpression) {
}
func (v *AbstractUpdateExpressionVisitor) VisitDoubleConstantExpression(_ intmod.IDoubleConstantExpression) {
}
func (v *AbstractUpdateExpressionVisitor) VisitEnumConstantReferenceExpression(_ intmod.IEnumConstantReferenceExpression) {
}
func (v *AbstractUpdateExpressionVisitor) VisitFloatConstantExpression(_ intmod.IFloatConstantExpression) {
}
func (v *AbstractUpdateExpressionVisitor) VisitIntegerConstantExpression(_ intmod.IIntegerConstantExpression) {
}
func (v *AbstractUpdateExpressionVisitor) VisitLocalVariableReferenceExpression(_ intmod.ILocalVariableReferenceExpression) {
}
func (v *AbstractUpdateExpressionVisitor) VisitLongConstantExpression(_ intmod.ILongConstantExpression) {
}
func (v *AbstractUpdateExpressionVisitor) VisitNullExpression(_ intmod.INullExpression) {
}
func (v *AbstractUpdateExpressionVisitor) VisitTypeReferenceDotClassExpression(_ intmod.ITypeReferenceDotClassExpression) {
}
func (v *AbstractUpdateExpressionVisitor) VisitObjectTypeReferenceExpression(_ intmod.IObjectTypeReferenceExpression) {
}
func (v *AbstractUpdateExpressionVisitor) VisitStringConstantExpression(_ intmod.IStringConstantExpression) {
}
func (v *AbstractUpdateExpressionVisitor) VisitSuperExpression(_ intmod.ISuperExpression) {
}
func (v *AbstractUpdateExpressionVisitor) VisitThisExpression(_ intmod.IThisExpression) {
}
func (v *AbstractUpdateExpressionVisitor) VisitAnnotationReference(_ intmod.IAnnotationReference) {
}
func (v *AbstractUpdateExpressionVisitor) VisitElementValueArrayInitializerElementValue(_ intmod.IElementValueArrayInitializerElementValue) {
}
func (v *AbstractUpdateExpressionVisitor) VisitAnnotationElementValue(_ intmod.IAnnotationElementValue) {
}
func (v *AbstractUpdateExpressionVisitor) VisitObjectReference(_ intmod.IObjectReference) {
}
func (v *AbstractUpdateExpressionVisitor) VisitBreakStatement(_ intmod.IBreakStatement) {}
func (v *AbstractUpdateExpressionVisitor) VisitByteCodeStatement(_ intmod.IByteCodeStatement) {
}
func (v *AbstractUpdateExpressionVisitor) VisitContinueStatement(_ intmod.IContinueStatement) {
}
func (v *AbstractUpdateExpressionVisitor) VisitReturnStatement(_ intmod.IReturnStatement) {
}
func (v *AbstractUpdateExpressionVisitor) VisitSwitchStatementDefaultLabel(_ intmod.IDefaultLabel) {
}

func (v *AbstractUpdateExpressionVisitor) VisitInnerObjectReference(_ intmod.IInnerObjectReference) {
}
func (v *AbstractUpdateExpressionVisitor) VisitTypeArguments(_ intmod.ITypeArguments) {}
func (v *AbstractUpdateExpressionVisitor) VisitWildcardExtendsTypeArgument(_ intmod.IWildcardExtendsTypeArgument) {
}
func (v *AbstractUpdateExpressionVisitor) VisitObjectType(_ intmod.IObjectType)           {}
func (v *AbstractUpdateExpressionVisitor) VisitInnerObjectType(_ intmod.IInnerObjectType) {}
func (v *AbstractUpdateExpressionVisitor) VisitWildcardSuperTypeArgument(_ intmod.IWildcardSuperTypeArgument) {
}
func (v *AbstractUpdateExpressionVisitor) VisitTypes(_ intmod.ITypes) {}
func (v *AbstractUpdateExpressionVisitor) VisitTypeParameterWithTypeBounds(_ intmod.ITypeParameterWithTypeBounds) {
}
