package javasyntax

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/declaration"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/expression"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/reference"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/statement"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
)

type AbstractJavaSyntaxVisitor struct {
	_type.AbstractTypeArgumentVisitor
}

func (v *AbstractJavaSyntaxVisitor) VisitCompilationUnit(compilationUnit CompilationUnit) {
	compilationUnit.TypeDeclarations().Accept(v)
}

// --- DeclarationVisitor ---

func (v *AbstractJavaSyntaxVisitor) VisitAnnotationDeclaration(decl *declaration.AnnotationDeclaration) {
	v.safeAcceptDeclaration(decl.AnnotationDeclarators())
	v.safeAcceptDeclaration(decl.BodyDeclaration())
	v.safeAcceptReference(decl.AnnotationReferences())
}

func (v *AbstractJavaSyntaxVisitor) VisitArrayVariableInitializer(decl *declaration.ArrayVariableInitializer) {
	v.acceptListDeclaration(decl.List())
}

func (v *AbstractJavaSyntaxVisitor) VisitBodyDeclaration(decl *declaration.BodyDeclaration) {
	v.safeAcceptDeclaration(decl.MemberDeclaration())
}

func (v *AbstractJavaSyntaxVisitor) VisitClassDeclaration(decl *declaration.ClassDeclaration) {
	superType := decl.SuperType()

	if superType != nil {
		superType.AcceptTypeVisitor(v)
	}

	v.safeAcceptTypeParameter(decl.TypeParameters())
	v.safeAcceptType(decl.Interfaces())
	v.safeAcceptReference(decl.AnnotationReferences())
	v.safeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *AbstractJavaSyntaxVisitor) VisitConstructorDeclaration(decl *declaration.ConstructorDeclaration) {
	v.safeAcceptReference(decl.AnnotationReferences())
	v.safeAcceptDeclaration(decl.FormalParameters())
	v.safeAcceptType(decl.ExceptionTypes())
	v.safeAcceptStatement(decl.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitEnumDeclaration(decl *declaration.EnumDeclaration) {
	v.VisitTypeDeclaration(&decl.TypeDeclaration)
	v.safeAcceptType(decl.Interfaces())
	v.safeAcceptListConstant(decl.Constants())
	v.safeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *AbstractJavaSyntaxVisitor) VisitEnumDeclarationConstant(decl *declaration.Constant) {
	v.safeAcceptReference(decl.AnnotationReferences())
	v.safeAcceptExpression(decl.Arguments())
	v.safeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *AbstractJavaSyntaxVisitor) VisitExpressionVariableInitializer(decl *declaration.ExpressionVariableInitializer) {
	decl.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitFieldDeclaration(decl *declaration.FieldDeclaration) {
	t := decl.Type()
	t.AcceptTypeVisitor(v)
	v.safeAcceptReference(decl.AnnotationReferences())
	decl.FieldDeclaration().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitFieldDeclarator(decl *declaration.FieldDeclarator) {
	v.safeAcceptDeclaration(decl.VariableInitializer())
}

func (v *AbstractJavaSyntaxVisitor) VisitFieldDeclarators(decl *declaration.FieldDeclarators) {
	v.acceptListDeclaration(decl.List())
}

func (v *AbstractJavaSyntaxVisitor) VisitFormalParameter(decl *declaration.FormalParameter) {
	t := decl.Type()
	t.AcceptTypeVisitor(v)
	v.safeAcceptReference(decl.AnnotationReferences())
}

func (v *AbstractJavaSyntaxVisitor) VisitFormalParameters(decl *declaration.FormalParameters) {
	v.acceptListDeclaration(decl.List())
}

func (v *AbstractJavaSyntaxVisitor) VisitInstanceInitializerDeclaration(decl *declaration.InstanceInitializerDeclaration) {
	v.safeAcceptStatement(decl.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitInterfaceDeclaration(decl *declaration.InterfaceDeclaration) {
	v.safeAcceptType(decl.Interfaces())
	v.safeAcceptReference(decl.AnnotationReferences())
	v.safeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *AbstractJavaSyntaxVisitor) VisitLocalVariableDeclaration(decl *declaration.LocalVariableDeclaration) {
	v.safeAcceptDeclaration(decl.LocalVariableDeclarators())
}

func (v *AbstractJavaSyntaxVisitor) VisitLocalVariableDeclarator(decl *declaration.LocalVariableDeclarator) {
	v.safeAcceptDeclaration(decl.VariableInitializer())
}

func (v *AbstractJavaSyntaxVisitor) VisitLocalVariableDeclarators(decl *declaration.LocalVariableDeclarators) {
	v.acceptListDeclaration(decl.List())
}

func (v *AbstractJavaSyntaxVisitor) VisitMethodDeclaration(decl *declaration.MethodDeclaration) {
	t := decl.ReturnType()
	t.AcceptTypeVisitor(v)

	v.safeAcceptReference(decl.AnnotationReferences())
	v.safeAcceptDeclaration(decl.FormalParameter())
	v.safeAcceptType(decl.ExceptionTypes())
	v.safeAcceptStatement(decl.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitMemberDeclarations(decl *declaration.MemberDeclarations) {
	v.acceptListDeclaration(decl.List())
}

func (v *AbstractJavaSyntaxVisitor) VisitModuleDeclaration(decl *declaration.ModuleDeclaration) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitStaticInitializerDeclaration(decl *declaration.StaticInitializerDeclaration) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitTypeDeclarations(decl *declaration.TypeDeclarations) {
	v.acceptListDeclaration(decl.List())
}

// --- ExpressionVisitor ---
func (v *AbstractJavaSyntaxVisitor) VisitArrayExpression(expr *expression.ArrayExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)

	expr.Expression().Accept(v)
	expr.Index().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitBinaryOperatorExpression(expr *expression.BinaryOperatorExpression) {
	expr.LeftExpression().Accept(v)
	expr.RightExpression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitBooleanExpression(expr *expression.BooleanExpression) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitCastExpression(expr *expression.CastExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	expr.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitCommentExpression(expr *expression.CommentExpression) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitConstructorInvocationExpression(expression *expression.ConstructorInvocationExpression) {
	t := expression.Type()
	t.AcceptTypeVisitor(v)
	v.safeAcceptExpression(expression.Parameters())
}

func (v *AbstractJavaSyntaxVisitor) VisitConstructorReferenceExpression(expr *expression.ConstructorReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitDoubleConstantExpression(expr *expression.DoubleConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitEnumConstantReferenceExpression(expr *expression.EnumConstantReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitExpressions(expr *expression.Expressions) {
	v.acceptListExpression(expr.List())
}

func (v *AbstractJavaSyntaxVisitor) VisitFieldReferenceExpression(expr *expression.FieldReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)

	v.safeAcceptExpression(expr.Expression())
}

func (v *AbstractJavaSyntaxVisitor) VisitFloatConstantExpression(expr *expression.FloatConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitIntegerConstantExpression(expr *expression.IntegerConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitInstanceOfExpression(expr *expression.InstanceOfExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	expr.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitLambdaFormalParametersExpression(expr *expression.LambdaFormalParametersExpression) {
	v.safeAcceptDeclaration(expr.FormalParameters())
	expr.Statements().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitLambdaIdentifiersExpression(expr *expression.LambdaIdentifiersExpression) {
	v.safeAcceptStatement(expr.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitLengthExpression(expr *expression.LengthExpression) {
	expr.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitLocalVariableReferenceExpression(expr *expression.LocalVariableReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitLongConstantExpression(expr *expression.LongConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitMethodInvocationExpression(expr *expression.MethodInvocationExpression) {
	expr.Expression().Accept(v)
	v.SafeAcceptTypeArgumentVisitable(expr.NonWildcardTypeArguments().(*_type.WildcardSuperTypeArgument))
	v.safeAcceptExpression(expr.Parameters())
}

func (v *AbstractJavaSyntaxVisitor) VisitMethodReferenceExpression(expr *expression.MethodReferenceExpression) {
	expr.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitNewArray(expr *expression.NewArray) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.safeAcceptExpression(expr.DimensionExpressionList())
}

func (v *AbstractJavaSyntaxVisitor) VisitNewExpression(expr *expression.NewExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.safeAcceptExpression(expr.Parameters())
}

func (v *AbstractJavaSyntaxVisitor) VisitNewInitializedArray(expr *expression.NewInitializedArray) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.safeAcceptDeclaration(expr.ArrayInitializer())
}

func (v *AbstractJavaSyntaxVisitor) VisitNoExpression(expr *expression.NoExpression) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitNullExpression(expr *expression.NullExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitObjectTypeReferenceExpression(expr *expression.ObjectTypeReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitParenthesesExpression(expr *expression.ParenthesesExpression) {
	expr.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitPostOperatorExpression(expr *expression.PostOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitPreOperatorExpression(expr *expression.PreOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitStringConstantExpression(expr *expression.StringConstantExpression) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitSuperConstructorInvocationExpression(expr *expression.SuperConstructorInvocationExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.safeAcceptExpression(expr.Parameters())
}

func (v *AbstractJavaSyntaxVisitor) VisitSuperExpression(expr *expression.SuperExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitTernaryOperatorExpression(expr *expression.TernaryOperatorExpression) {
	expr.Condition().Accept(v)
	expr.TrueExpression().Accept(v)
	expr.FalseExpression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitThisExpression(expr *expression.ThisExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitTypeReferenceDotClassExpression(expr *expression.TypeReferenceDotClassExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

// --- ReferenceVisitor ---

func (v *AbstractJavaSyntaxVisitor) VisitAnnotationElementValue(ref *reference.AnnotationElementValue) {
	v.safeAcceptReference(ref.ElementValue())
	v.safeAcceptReference(ref.ElementValuePairs())
}

func (v *AbstractJavaSyntaxVisitor) VisitAnnotationReference(ref *reference.AnnotationReference) {
	v.safeAcceptReference(ref.ElementValue())
	v.safeAcceptReference(ref.ElementValuePairs())
}

func (v *AbstractJavaSyntaxVisitor) VisitAnnotationReferences(ref *reference.AnnotationReferences) {
	v.acceptListReference(ref.List())
}

func (v *AbstractJavaSyntaxVisitor) VisitElementValueArrayInitializerElementValue(ref *reference.ElementValueArrayInitializerElementValue) {
	v.safeAcceptReference(ref.ElementValueArrayInitializer())
}

func (v *AbstractJavaSyntaxVisitor) VisitElementValues(ref *reference.ElementValues) {
	v.acceptListReference(ref.List())
}

func (v *AbstractJavaSyntaxVisitor) VisitElementValuePair(ref *reference.ElementValuePair) {
	ref.ElementValue().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitElementValuePairs(ref *reference.ElementValuePairs) {
	v.acceptListReference(ref.List())
}

func (v *AbstractJavaSyntaxVisitor) VisitExpressionElementValue(ref *reference.ExpressionElementValue) {
	ref.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitInnerObjectReference(ref *reference.InnerObjectReference) {
	v.VisitInnerObjectType(&ref.InnerObjectType)
}

func (v *AbstractJavaSyntaxVisitor) VisitObjectReference(ref *reference.ObjectReference) {
	v.VisitObjectType(&ref.ObjectType)
}

// --- StatementVisitor ---

func (v *AbstractJavaSyntaxVisitor) VisitAssertStatement(stat *statement.AssertStatement) {
	stat.Condition().Accept(v)
	v.safeAcceptExpression(stat.Message())
}

func (v *AbstractJavaSyntaxVisitor) VisitBreakStatement(stat *statement.BreakStatement) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitByteCodeStatement(stat *statement.ByteCodeStatement) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitCommentStatement(stat *statement.CommentStatement) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitContinueStatement(stat *statement.ContinueStatement) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitDoWhileStatement(stat *statement.DoWhileStatement) {
	v.safeAcceptExpression(stat.Condition())
	v.safeAcceptStatement(stat.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitExpressionStatement(stat *statement.ExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitForEachStatement(stat *statement.ForEachStatement) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	stat.Expression().Accept(v)
	v.safeAcceptStatement(stat.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitForStatement(stat *statement.ForStatement) {
	v.safeAcceptDeclaration(stat.Declaration())
	v.safeAcceptExpression(stat.Init())
	v.safeAcceptExpression(stat.Condition())
	v.safeAcceptExpression(stat.Update())
	v.safeAcceptStatement(stat.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitIfStatement(stat *statement.IfStatement) {
	stat.Condition().Accept(v)
	v.safeAcceptStatement(stat.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitIfElseStatement(stat *statement.IfElseStatement) {
	stat.Condition().Accept(v)
	v.safeAcceptStatement(stat.Statements())
	stat.ElseStatements().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitLabelStatement(stat *statement.LabelStatement) {
	v.safeAcceptStatement(stat.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitLambdaExpressionStatement(stat *statement.LambdaExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitLocalVariableDeclarationStatement(stat *statement.LocalVariableDeclarationStatement) {
	v.VisitLocalVariableDeclaration(&stat.LocalVariableDeclaration)
}

func (v *AbstractJavaSyntaxVisitor) VisitNoStatement(stat *statement.NoStatement) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitReturnExpressionStatement(stat *statement.ReturnExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitReturnStatement(stat *statement.ReturnStatement) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitStatements(stat *statement.Statements) {
	v.acceptListStatement(stat.List())
}

func (v *AbstractJavaSyntaxVisitor) VisitSwitchStatement(stat *statement.SwitchStatement) {
	stat.Condition().Accept(v)
	v.acceptListStatement(stat.List())
}

func (v *AbstractJavaSyntaxVisitor) VisitSwitchStatementDefaultLabel(stat *statement.DefaultLabe1) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitSwitchStatementExpressionLabel(stat *statement.ExpressionLabel) {
	stat.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitSwitchStatementLabelBlock(stat *statement.LabelBlock) {
	stat.Label().Accept(v)
	stat.Statements().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitSwitchStatementMultiLabelsBlock(stat *statement.MultiLabelsBlock) {
	v.safeAcceptListStatement(stat.List())
	stat.Statements().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitSynchronizedStatement(stat *statement.SynchronizedStatement) {
	stat.Monitor().Accept(v)
	v.safeAcceptStatement(stat.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitThrowStatement(stat *statement.ThrowStatement) {
	stat.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitTryStatement(stat *statement.TryStatement) {
	v.safeAcceptListStatement(stat.ResourceList())
	stat.TryStatements().Accept(v)
	v.safeAcceptListStatement(stat.CatchClauseList())
	v.safeAcceptStatement(stat.FinallyStatements())
}

func (v *AbstractJavaSyntaxVisitor) VisitTryStatementResource(stat *statement.Resource) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	stat.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitTryStatementCatchClause(stat *statement.CatchClause) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	v.safeAcceptStatement(stat.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitTypeDeclarationStatement(stat *statement.TypeDeclarationStatement) {
	stat.TypeDeclaration().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitWhileStatement(stat *statement.WhileStatement) {
	stat.Condition().Accept(v)
	v.safeAcceptStatement(stat.Statements())
}

// --- TypeVisitor ---
// func (v *AbstractJavaSyntaxVisitor) VisitPrimitiveType(y *PrimitiveType)     {}
// func (v *AbstractJavaSyntaxVisitor) VisitObjectType(y *ObjectType)           {}
// func (v *AbstractJavaSyntaxVisitor) VisitInnerObjectType(y *InnerObjectType) {}

func (v *AbstractJavaSyntaxVisitor) VisitTypes(types *_type.Types) {
	for _, value := range types.Types {
		value.AcceptTypeVisitor(v)
	}
}

//func (v *AbstractJavaSyntaxVisitor) VisitGenericType(y *GenericType)         {}

// --- TypeParameterVisitor --- //

func (v *AbstractJavaSyntaxVisitor) VisitTypeParameter(parameter *_type.TypeParameter) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitTypeParameterWithTypeBounds(parameter *_type.TypeParameterWithTypeBounds) {
	parameter.TypeBounds().AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitTypeParameters(parameters *_type.TypeParameters) {
	for _, param := range parameters.TypeParameters {
		param.AcceptTypeParameterVisitor(v)
	}
}

// --- TypeArgumentVisitor ---

//func (v *AbstractJavaSyntaxVisitor) VisitTypeArguments(arguments *_type.TypeArguments) {}
//
//func (v *AbstractJavaSyntaxVisitor) VisitDiamondTypeArgument(argument *_type.DiamondTypeArgument) {}
//
//func (v *AbstractJavaSyntaxVisitor) VisitWildcardExtendsTypeArgument(argument *_type.WildcardExtendsTypeArgument) {
//}
//
//func (v *AbstractJavaSyntaxVisitor) VisitWildcardSuperTypeArgument(argument *_type.WildcardSuperTypeArgument) {
//}
//
//func (v *AbstractJavaSyntaxVisitor) VisitWildcardTypeArgument(argument *_type.WildcardTypeArgument) {}
//
//func (v *AbstractJavaSyntaxVisitor) VisitPrimitiveType(t *_type.PrimitiveType) {}
//
//func (v *AbstractJavaSyntaxVisitor) VisitObjectType(t *_type.ObjectType) {}
//
//func (v *AbstractJavaSyntaxVisitor) VisitInnerObjectType(t *_type.InnerObjectType) {}
//
//func (v *AbstractJavaSyntaxVisitor) VisitGenericType(t *_type.GenericType) {}

func (v *AbstractJavaSyntaxVisitor) VisitTypeDeclaration(decl *declaration.TypeDeclaration) {
	v.safeAcceptReference(decl.AnnotationReferences())
}

func (v *AbstractJavaSyntaxVisitor) acceptListDeclaration(list []declaration.Declaration) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) acceptListExpression(list []expression.Expression) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) acceptListReference(list []reference.IReference) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) acceptListStatement(list []statement.Statement) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) safeAcceptDeclaration(decl declaration.Declaration) {
	if decl != nil {
		decl.Accept(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) safeAcceptExpression(expr expression.Expression) {
	if expr != nil {
		expr.Accept(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) safeAcceptReference(ref reference.IReference) {
	if ref != nil {
		ref.Accept(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) safeAcceptStatement(list statement.Statement) {
	if list != nil {
		list.Accept(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) safeAcceptType(list _type.IType) {
	if list != nil {
		list.AcceptTypeVisitor(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) safeAcceptTypeParameter(list _type.ITypeParameter) {
	if list != nil {
		list.AcceptTypeParameterVisitor(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) safeAcceptListDeclaration(list []declaration.Declaration) {
	if list != nil {
		for _, value := range list {
			value.Accept(v)
		}
	}
}

func (v *AbstractJavaSyntaxVisitor) safeAcceptListConstant(list []declaration.Constant) {
	if list != nil {
		for _, value := range list {
			value.Accept(v)
		}
	}
}

func (v *AbstractJavaSyntaxVisitor) safeAcceptListStatement(list []statement.Statement) {
	if list != nil {
		for _, value := range list {
			value.Accept(v)
		}
	}
}
