package javasyntax

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
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

func (v *AbstractJavaSyntaxVisitor) VisitAnnotationDeclaration(decl intsyn.IAnnotationDeclaration) {
	v.SafeAcceptDeclaration(decl.AnnotationDeclarators())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *AbstractJavaSyntaxVisitor) VisitArrayVariableInitializer(decl intsyn.IArrayVariableInitializer) {
	v.AcceptListDeclaration(decl.List())
}

func (v *AbstractJavaSyntaxVisitor) VisitBodyDeclaration(decl declaration.IBodyDeclaration) {
	v.SafeAcceptDeclaration(decl.MemberDeclarations())
}

func (v *AbstractJavaSyntaxVisitor) VisitClassDeclaration(decl intsyn.IClassDeclaration) {
	superType := decl.SuperType()

	if superType != nil {
		superType.AcceptTypeVisitor(v)
	}

	v.SafeAcceptTypeParameter(decl.TypeParameters())
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *AbstractJavaSyntaxVisitor) VisitConstructorDeclaration(decl intsyn.IConstructorDeclaration) {
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.FormalParameters())
	v.SafeAcceptType(decl.ExceptionTypes())
	v.SafeAcceptStatement(decl.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitEnumDeclaration(decl intsyn.IEnumDeclaration) {
	v.VisitTypeDeclaration(&decl.TypeDeclaration)
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptListConstant(decl.Constants())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *AbstractJavaSyntaxVisitor) VisitEnumDeclarationConstant(decl intsyn.IConstant) {
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptExpression(decl.Arguments())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *AbstractJavaSyntaxVisitor) VisitExpressionVariableInitializer(decl intsyn.IExpressionVariableInitializer) {
	decl.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitFieldDeclaration(decl intsyn.IFieldDeclaration) {
	t := decl.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptReference(decl.AnnotationReferences())
	decl.FieldDeclarators().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitFieldDeclarator(decl intsyn.IFieldDeclarator) {
	v.SafeAcceptDeclaration(decl.VariableInitializer())
}

func (v *AbstractJavaSyntaxVisitor) VisitFieldDeclarators(decl intsyn.IFieldDeclarators) {
	list := decl.List()
	decls := make([]declaration.Declaration, 0, len(list))
	for _, item := range list {
		decls = append(decls, item)
	}
	v.AcceptListDeclaration(decls)
}

func (v *AbstractJavaSyntaxVisitor) VisitFormalParameter(decl intsyn.IFormalParameter) {
	t := decl.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *AbstractJavaSyntaxVisitor) VisitFormalParameters(decl intsyn.IFormalParameters) {
	v.AcceptListDeclaration(decl.List())
}

func (v *AbstractJavaSyntaxVisitor) VisitInstanceInitializerDeclaration(decl intsyn.IInstanceInitializerDeclaration) {
	v.SafeAcceptStatement(decl.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitInterfaceDeclaration(decl intsyn.IInterfaceDeclaration) {
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *AbstractJavaSyntaxVisitor) VisitLocalVariableDeclaration(decl intsyn.ILocalVariableDeclaration) {
	v.SafeAcceptDeclaration(decl.LocalVariableDeclarators())
}

func (v *AbstractJavaSyntaxVisitor) VisitLocalVariableDeclarator(decl intsyn.ILocalVariableDeclarator) {
	v.SafeAcceptDeclaration(decl.VariableInitializer())
}

func (v *AbstractJavaSyntaxVisitor) VisitLocalVariableDeclarators(decl intsyn.ILocalVariableDeclarators) {
	v.AcceptListDeclaration(decl.List())
}

func (v *AbstractJavaSyntaxVisitor) VisitMethodDeclaration(decl intsyn.IMethodDeclaration) {
	t := decl.ReturnedType()
	t.AcceptTypeVisitor(v)

	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.FormalParameter())
	v.SafeAcceptType(decl.ExceptionTypes())
	v.SafeAcceptStatement(decl.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitMemberDeclarations(decl intsyn.IMemberDeclarations) {
	v.AcceptListDeclaration(decl.List())
}

func (v *AbstractJavaSyntaxVisitor) VisitModuleDeclaration(decl intsyn.IModuleDeclaration) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitStaticInitializerDeclaration(decl intsyn.IStaticInitializerDeclaration) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitTypeDeclarations(decl intsyn.ITypeDeclarations) {
	v.AcceptListDeclaration(decl.List())
}

// --- IExpressionVisitor ---
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
	v.SafeAcceptExpression(expression.Parameters())
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
	v.AcceptListExpression(expr.List())
}

func (v *AbstractJavaSyntaxVisitor) VisitFieldReferenceExpression(expr *expression.FieldReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)

	v.SafeAcceptExpression(expr.Expression())
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
	v.SafeAcceptDeclaration(expr.FormalParameters())
	expr.Statements().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitLambdaIdentifiersExpression(expr *expression.LambdaIdentifiersExpression) {
	v.SafeAcceptStatement(expr.Statements())
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
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *AbstractJavaSyntaxVisitor) VisitMethodReferenceExpression(expr *expression.MethodReferenceExpression) {
	expr.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitNewArray(expr *expression.NewArray) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.DimensionExpressionList())
}

func (v *AbstractJavaSyntaxVisitor) VisitNewExpression(expr *expression.NewExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *AbstractJavaSyntaxVisitor) VisitNewInitializedArray(expr *expression.NewInitializedArray) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptDeclaration(expr.ArrayInitializer())
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
	v.SafeAcceptExpression(expr.Parameters())
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

// --- IReferenceVisitor ---

func (v *AbstractJavaSyntaxVisitor) VisitAnnotationElementValue(ref *reference.AnnotationElementValue) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *AbstractJavaSyntaxVisitor) VisitAnnotationReference(ref *reference.AnnotationReference) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *AbstractJavaSyntaxVisitor) VisitAnnotationReferences(ref *reference.AnnotationReferences) {
	v.AcceptListReference(ref.List())
}

func (v *AbstractJavaSyntaxVisitor) VisitElementValueArrayInitializerElementValue(ref *reference.ElementValueArrayInitializerElementValue) {
	v.SafeAcceptReference(ref.ElementValueArrayInitializer())
}

func (v *AbstractJavaSyntaxVisitor) VisitElementValues(ref *reference.ElementValues) {
	v.AcceptListReference(ref.List())
}

func (v *AbstractJavaSyntaxVisitor) VisitElementValuePair(ref *reference.ElementValuePair) {
	ref.ElementValue().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitElementValuePairs(ref *reference.ElementValuePairs) {
	v.AcceptListReference(ref.List())
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

// --- IStatementVisitor ---

func (v *AbstractJavaSyntaxVisitor) VisitAssertStatement(stat *statement.AssertStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptExpression(stat.Message())
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
	v.SafeAcceptExpression(stat.Condition())
	v.SafeAcceptStatement(stat.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitExpressionStatement(stat *statement.ExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitForEachStatement(stat *statement.ForEachStatement) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	stat.Expression().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitForStatement(stat *statement.ForStatement) {
	v.SafeAcceptDeclaration(stat.Declaration())
	v.SafeAcceptExpression(stat.Init())
	v.SafeAcceptExpression(stat.Condition())
	v.SafeAcceptExpression(stat.Update())
	v.SafeAcceptStatement(stat.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitIfStatement(stat *statement.IfStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitIfElseStatement(stat *statement.IfElseStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
	stat.ElseStatements().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitLabelStatement(stat *statement.LabelStatement) {
	v.SafeAcceptStatement(stat.Statements())
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
	v.AcceptListStatement(stat.List())
}

func (v *AbstractJavaSyntaxVisitor) VisitSwitchStatement(stat *statement.SwitchStatement) {
	stat.Condition().Accept(v)
	v.AcceptListStatement(stat.List())
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
	v.SafeAcceptListStatement(stat.List())
	stat.Statements().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitSynchronizedStatement(stat *statement.SynchronizedStatement) {
	stat.Monitor().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitThrowStatement(stat *statement.ThrowStatement) {
	stat.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitTryStatement(stat *statement.TryStatement) {
	v.SafeAcceptListStatement(stat.ResourceList())
	stat.TryStatements().Accept(v)
	v.SafeAcceptListStatement(stat.CatchClauseList())
	v.SafeAcceptStatement(stat.FinallyStatements())
}

func (v *AbstractJavaSyntaxVisitor) VisitTryStatementResource(stat *statement.Resource) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	stat.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitTryStatementCatchClause(stat *statement.CatchClause) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitTypeDeclarationStatement(stat *statement.TypeDeclarationStatement) {
	stat.TypeDeclaration().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitWhileStatement(stat *statement.WhileStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

// --- ITypeVisitor ---
// func (v *AbstractJavaSyntaxVisitor) VisitPrimitiveType(y *PrimitiveType)     {}
// func (v *AbstractJavaSyntaxVisitor) VisitObjectType(y *ObjectType)           {}
// func (v *AbstractJavaSyntaxVisitor) VisitInnerObjectType(y *InnerObjectType) {}

func (v *AbstractJavaSyntaxVisitor) VisitTypes(types *_type.Types) {
	for _, value := range types.Types {
		value.AcceptTypeVisitor(v)
	}
}

//func (v *AbstractJavaSyntaxVisitor) VisitGenericType(y *GenericType)         {}

// --- ITypeParameterVisitor --- //

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

// --- ITypeArgumentVisitor ---

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

func (v *AbstractJavaSyntaxVisitor) VisitTypeDeclaration(decl intsyn.ITypeDeclaration) {
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *AbstractJavaSyntaxVisitor) AcceptListDeclaration(list []declaration.Declaration) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) AcceptListExpression(list []expression.IExpression) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) AcceptListReference(list []reference.IReference) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) AcceptListStatement(list []statement.IStatement) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptDeclaration(decl declaration.Declaration) {
	if decl != nil {
		decl.Accept(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptExpression(expr expression.IExpression) {
	if expr != nil {
		expr.Accept(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptReference(ref reference.IReference) {
	if ref != nil {
		ref.Accept(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptStatement(list statement.IStatement) {
	if list != nil {
		list.Accept(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptType(list _type.IType) {
	if list != nil {
		list.AcceptTypeVisitor(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptTypeParameter(list _type.ITypeParameter) {
	if list != nil {
		list.AcceptTypeParameterVisitor(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptListDeclaration(list []declaration.Declaration) {
	if list != nil {
		for _, value := range list {
			value.Accept(v)
		}
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptListConstant(list []declaration.Constant) {
	if list != nil {
		for _, value := range list {
			value.Accept(v)
		}
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptListStatement(list []statement.IStatement) {
	if list != nil {
		for _, value := range list {
			value.Accept(v)
		}
	}
}
