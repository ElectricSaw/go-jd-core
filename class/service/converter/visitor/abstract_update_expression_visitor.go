package visitor

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/declaration"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/expression"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/reference"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/statement"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
)

type AbstractUpdateExpressionVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor
}

func (v *AbstractUpdateExpressionVisitor) UpdateExpression(expression expression.Expression) expression.Expression {
	return nil
}

func (v *AbstractUpdateExpressionVisitor) UpdateBaseExpression(baseExpression expression.Expression) expression.Expression {
	if baseExpression == nil {
		return nil
	}

	if baseExpression.IsList() {
		iterator := baseExpression.Iterator()

		for iterator.HasNext() {
			value, _ := iterator.Next()
			_ = iterator.Set(v.UpdateExpression(value))
		}

		return baseExpression
	}

	return v.UpdateExpression(baseExpression.First())
}

func (v *AbstractUpdateExpressionVisitor) VisitAnnotationDeclaration(decl *declaration.AnnotationDeclaration) {
	v.SafeAcceptDeclaration(decl.AnnotationDeclarators())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *AbstractUpdateExpressionVisitor) VisitClassDeclaration(declaration *declaration.ClassDeclaration) {
	v.VisitInterfaceDeclaration(&declaration.InterfaceDeclaration)
}

func (v *AbstractUpdateExpressionVisitor) VisitConstructorInvocationExpression(expression *expression.ConstructorInvocationExpression) {
	if expression.Parameters() != nil {
		expression.SetParameters(v.UpdateBaseExpression(expression.Parameters()))
		expression.Parameters().Accept(v)
	}
}

func (v *AbstractUpdateExpressionVisitor) VisitConstructorDeclaration(declaration *declaration.ConstructorDeclaration) {
	v.SafeAcceptStatement(declaration.Statements())
}

func (v *AbstractUpdateExpressionVisitor) VisitEnumDeclaration(declaration *declaration.EnumDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *AbstractUpdateExpressionVisitor) VisitEnumDeclarationConstant(declaration *declaration.Constant) {
	if declaration.Arguments() != nil {
		declaration.SetArguments(v.UpdateBaseExpression(declaration.Arguments()))
		declaration.Arguments().Accept(v)
	}
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *AbstractUpdateExpressionVisitor) VisitExpressionVariableInitializer(declaration *declaration.ExpressionVariableInitializer) {
	if declaration.Expression() != nil {
		declaration.SetExpression(v.UpdateExpression(declaration.Expression()))
		declaration.Expression().Accept(v)
	}
}

func (v *AbstractUpdateExpressionVisitor) VisitFieldDeclaration(declaration *declaration.FieldDeclaration) {
	declaration.FieldDeclarators().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitFieldDeclarator(declaration *declaration.FieldDeclarator) {
	v.SafeAcceptDeclaration(declaration.VariableInitializer())
}

func (v *AbstractUpdateExpressionVisitor) VisitFormalParameter(declaration *declaration.FormalParameter) {
}

func (v *AbstractUpdateExpressionVisitor) VisitInterfaceDeclaration(declaration *declaration.InterfaceDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *AbstractUpdateExpressionVisitor) VisitLocalVariableDeclaration(declaration *declaration.LocalVariableDeclaration) {
	declaration.LocalVariableDeclarators().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitLocalVariableDeclarator(declarator *declaration.LocalVariableDeclarator) {
	v.SafeAcceptDeclaration(declarator.VariableInitializer())
}

func (v *AbstractUpdateExpressionVisitor) VisitMethodDeclaration(declaration *declaration.MethodDeclaration) {
	v.SafeAcceptReference(declaration.AnnotationReferences())
	v.SafeAcceptStatement(declaration.Statements())
}

func (v *AbstractUpdateExpressionVisitor) VisitArrayExpression(expression *expression.ArrayExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.SetIndex(v.UpdateExpression(expression.Index()))
	expression.Expression().Accept(v)
	expression.Index().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitBinaryOperatorExpression(expression *expression.BinaryOperatorExpression) {
	expression.SetLeftExpression(v.UpdateExpression(expression.LeftExpression()))
	expression.SetRightExpression(v.UpdateExpression(expression.RightExpression()))
	expression.LeftExpression().Accept(v)
	expression.RightExpression().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitCastExpression(expression *expression.CastExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.Expression().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitFieldReferenceExpression(expression *expression.FieldReferenceExpression) {
	if expression.Expression() != nil {
		expression.SetExpression(v.UpdateExpression(expression.Expression()))
		expression.Expression().Accept(v)
	}
}

func (v *AbstractUpdateExpressionVisitor) VisitInstanceOfExpression(expression *expression.InstanceOfExpression) {
	if expression.Expression() != nil {
		expression.SetExpression(v.UpdateExpression(expression.Expression()))
		expression.Expression().Accept(v)
	}
}

func (v *AbstractUpdateExpressionVisitor) VisitLambdaFormalParametersExpression(expression *expression.LambdaFormalParametersExpression) {
	expression.Statements().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitLengthExpression(expression *expression.LengthExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.Expression().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitMethodInvocationExpression(expression *expression.MethodInvocationExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	if expression.Parameters() != nil {
		expression.SetParameters(v.UpdateBaseExpression(expression.Parameters()))
		expression.Parameters().Accept(v)
	}
	expression.Expression().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitMethodReferenceExpression(expression *expression.MethodReferenceExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.Expression().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitNewArray(expression *expression.NewArray) {
	if expression.DimensionExpressionList() != nil {
		expression.SetDimensionExpressionList(v.UpdateBaseExpression(expression.DimensionExpressionList()))
		expression.DimensionExpressionList().Accept(v)
	}
}

func (v *AbstractUpdateExpressionVisitor) VisitNewExpression(expression *expression.NewExpression) {
	if expression.Parameters() != nil {
		expression.SetParameters(v.UpdateBaseExpression(expression.Parameters()))
		expression.Parameters().Accept(v)
	}
	// v.SafeAccept(expression.BodyDeclaration());
}

func (v *AbstractUpdateExpressionVisitor) VisitNewInitializedArray(expression *expression.NewInitializedArray) {
	v.SafeAcceptDeclaration(expression.ArrayInitializer())
}

func (v *AbstractUpdateExpressionVisitor) VisitParenthesesExpression(expression *expression.ParenthesesExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.Expression().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitPostOperatorExpression(expression *expression.PostOperatorExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.Expression().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitPreOperatorExpression(expression *expression.PreOperatorExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.Expression().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitSuperConstructorInvocationExpression(expression *expression.SuperConstructorInvocationExpression) {
	if expression.Parameters() != nil {
		expression.SetParameters(v.UpdateBaseExpression(expression.Parameters()))
		expression.Parameters().Accept(v)
	}
}

func (v *AbstractUpdateExpressionVisitor) VisitTernaryOperatorExpression(expression *expression.TernaryOperatorExpression) {
	expression.SetCondition(v.UpdateExpression(expression.Condition()))
	expression.SetTrueExpression(v.UpdateExpression(expression.TrueExpression()))
	expression.SetFalseExpression(v.UpdateExpression(expression.FalseExpression()))
	expression.Condition().Accept(v)
	expression.TrueExpression().Accept(v)
	expression.FalseExpression().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitExpressionElementValue(reference *reference.ExpressionElementValue) {
	reference.SetExpression(v.UpdateExpression(reference.Expression()))
	reference.Expression().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitAssertStatement(statement *statement.AssertStatement) {
	statement.SetCondition(v.UpdateExpression(statement.Condition()))
	statement.Condition().Accept(v)
	v.SafeAcceptExpression(statement.Message())
}

func (v *AbstractUpdateExpressionVisitor) VisitDoWhileStatement(statement *statement.DoWhileStatement) {
	statement.SetCondition(v.UpdateExpression(statement.Condition()))
	v.SafeAcceptExpression(statement.Condition())
	v.SafeAcceptStatement(statement.Statements())
}

func (v *AbstractUpdateExpressionVisitor) VisitExpressionStatement(statement *statement.ExpressionStatement) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitForEachStatement(statement *statement.ForEachStatement) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
	v.SafeAcceptStatement(statement.Statements())
}

func (v *AbstractUpdateExpressionVisitor) VisitForStatement(statement *statement.ForStatement) {
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

func (v *AbstractUpdateExpressionVisitor) VisitIfStatement(statement *statement.IfStatement) {
	statement.SetCondition(v.UpdateExpression(statement.Condition()))
	statement.Condition().Accept(v)
	v.SafeAcceptStatement(statement.Statements())
}

func (v *AbstractUpdateExpressionVisitor) VisitIfElseStatement(statement *statement.IfElseStatement) {
	statement.SetCondition(v.UpdateExpression(statement.Condition()))
	statement.Condition().Accept(v)
	v.SafeAcceptStatement(statement.Statements())
	statement.ElseStatements().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitLambdaExpressionStatement(statement *statement.LambdaExpressionStatement) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitReturnExpressionStatement(statement *statement.ReturnExpressionStatement) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitSwitchStatement(statement *statement.SwitchStatement) {
	statement.SetCondition(v.UpdateExpression(statement.Condition()))
	statement.Condition().Accept(v)
	v.AcceptListStatement(statement.List())
}

func (v *AbstractUpdateExpressionVisitor) VisitSwitchStatementExpressionLabel(statement *statement.ExpressionLabel) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitSynchronizedStatement(statement *statement.SynchronizedStatement) {
	statement.SetMonitor(v.UpdateExpression(statement.Monitor()))
	statement.Monitor().Accept(v)
	v.SafeAcceptStatement(statement.Statements())
}

func (v *AbstractUpdateExpressionVisitor) VisitThrowStatement(statement *statement.ThrowStatement) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitTryStatementCatchClause(statement *statement.CatchClause) {
	v.SafeAcceptStatement(statement.Statements())
}

func (v *AbstractUpdateExpressionVisitor) VisitTryStatementResource(statement *statement.Resource) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
}

func (v *AbstractUpdateExpressionVisitor) VisitWhileStatement(statement *statement.WhileStatement) {
	statement.SetCondition(v.UpdateExpression(statement.Condition()))
	statement.Condition().Accept(v)
	v.SafeAcceptStatement(statement.Statements())
}

func (v *AbstractUpdateExpressionVisitor) VisitConstructorReferenceExpression(expression *expression.ConstructorReferenceExpression) {
}
func (v *AbstractUpdateExpressionVisitor) VisitDoubleConstantExpression(expression *expression.DoubleConstantExpression) {
}
func (v *AbstractUpdateExpressionVisitor) VisitEnumConstantReferenceExpression(expression *expression.EnumConstantReferenceExpression) {
}
func (v *AbstractUpdateExpressionVisitor) VisitFloatConstantExpression(expression *expression.FloatConstantExpression) {
}
func (v *AbstractUpdateExpressionVisitor) VisitIntegerConstantExpression(expression *expression.IntegerConstantExpression) {
}
func (v *AbstractUpdateExpressionVisitor) VisitLocalVariableReferenceExpression(expression *expression.LocalVariableReferenceExpression) {
}
func (v *AbstractUpdateExpressionVisitor) VisitLongConstantExpression(expression *expression.LongConstantExpression) {
}
func (v *AbstractUpdateExpressionVisitor) VisitNullExpression(expression *expression.NullExpression) {
}
func (v *AbstractUpdateExpressionVisitor) VisitTypeReferenceDotClassExpression(expression *expression.TypeReferenceDotClassExpression) {
}
func (v *AbstractUpdateExpressionVisitor) VisitObjectTypeReferenceExpression(expression *expression.ObjectTypeReferenceExpression) {
}
func (v *AbstractUpdateExpressionVisitor) VisitStringConstantExpression(expression *expression.StringConstantExpression) {
}
func (v *AbstractUpdateExpressionVisitor) VisitSuperExpression(expression *expression.SuperExpression) {
}
func (v *AbstractUpdateExpressionVisitor) VisitThisExpression(expression *expression.ThisExpression) {
}
func (v *AbstractUpdateExpressionVisitor) VisitAnnotationReference(reference *reference.AnnotationReference) {
}
func (v *AbstractUpdateExpressionVisitor) VisitElementValueArrayInitializerElementValue(reference *reference.ElementValueArrayInitializerElementValue) {
}
func (v *AbstractUpdateExpressionVisitor) VisitAnnotationElementValue(reference *reference.AnnotationElementValue) {
}
func (v *AbstractUpdateExpressionVisitor) VisitObjectReference(reference *reference.ObjectReference) {
}
func (v *AbstractUpdateExpressionVisitor) VisitBreakStatement(statement *statement.BreakStatement) {}
func (v *AbstractUpdateExpressionVisitor) VisitByteCodeStatement(statement *statement.ByteCodeStatement) {
}
func (v *AbstractUpdateExpressionVisitor) VisitContinueStatement(statement *statement.ContinueStatement) {
}
func (v *AbstractUpdateExpressionVisitor) VisitReturnStatement(statement *statement.ReturnStatement) {
}
func (v *AbstractUpdateExpressionVisitor) VisitSwitchStatementDefaultLabel(statement *statement.DefaultLabe1) {
}

func (v *AbstractUpdateExpressionVisitor) VisitInnerObjectReference(reference *reference.InnerObjectReference) {
}
func (v *AbstractUpdateExpressionVisitor) VisitTypeArguments(typ *_type.TypeArguments) {}
func (v *AbstractUpdateExpressionVisitor) VisitWildcardExtendsTypeArgument(typ *_type.WildcardExtendsTypeArgument) {
}
func (v *AbstractUpdateExpressionVisitor) VisitObjectType(typ *_type.ObjectType)           {}
func (v *AbstractUpdateExpressionVisitor) VisitInnerObjectType(typ *_type.InnerObjectType) {}
func (v *AbstractUpdateExpressionVisitor) VisitWildcardSuperTypeArgument(typ *_type.WildcardSuperTypeArgument) {
}
func (v *AbstractUpdateExpressionVisitor) VisitTypes(list *_type.Types) {}
func (v *AbstractUpdateExpressionVisitor) VisitTypeParameterWithTypeBounds(typ *_type.TypeParameterWithTypeBounds) {
}
