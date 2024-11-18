package javasyntax

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/model"
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
	list := make([]intsyn.IDeclaration, 0, decl.Size())
	for _, element := range decl.Elements() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *AbstractJavaSyntaxVisitor) VisitBodyDeclaration(decl intsyn.IBodyDeclaration) {
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
	v.VisitTypeDeclaration(decl.(intsyn.ITypeDeclaration))
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
	list := make([]intsyn.IDeclaration, 0, decl.Size())
	for _, element := range decl.Elements() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *AbstractJavaSyntaxVisitor) VisitFormalParameter(decl intsyn.IFormalParameter) {
	t := decl.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *AbstractJavaSyntaxVisitor) VisitFormalParameters(decl intsyn.IFormalParameters) {
	list := make([]intsyn.IDeclaration, 0, decl.Size())
	for _, element := range decl.Elements() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
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
	list := make([]intsyn.IDeclaration, 0, decl.Size())
	for _, element := range decl.Elements() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
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
	list := make([]intsyn.IDeclaration, 0, decl.Size())
	for _, element := range decl.Elements() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *AbstractJavaSyntaxVisitor) VisitModuleDeclaration(decl intsyn.IModuleDeclaration) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitStaticInitializerDeclaration(decl intsyn.IStaticInitializerDeclaration) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitTypeDeclarations(decl intsyn.ITypeDeclarations) {
	list := make([]intsyn.IDeclaration, 0, decl.Size())
	for _, element := range decl.Elements() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

// --- IExpressionVisitor ---
func (v *AbstractJavaSyntaxVisitor) VisitArrayExpression(expr intsyn.IArrayExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)

	expr.Expression().Accept(v)
	expr.Index().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitBinaryOperatorExpression(expr intsyn.IBinaryOperatorExpression) {
	expr.LeftExpression().Accept(v)
	expr.RightExpression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitBooleanExpression(expr intsyn.IBooleanExpression) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitCastExpression(expr intsyn.ICastExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	expr.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitCommentExpression(expr intsyn.ICommentExpression) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitConstructorInvocationExpression(expression intsyn.IConstructorInvocationExpression) {
	t := expression.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expression.Parameters())
}

func (v *AbstractJavaSyntaxVisitor) VisitConstructorReferenceExpression(expr intsyn.IConstructorReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitDoubleConstantExpression(expr intsyn.IDoubleConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitEnumConstantReferenceExpression(expr intsyn.IEnumConstantReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitExpressions(expr intsyn.IExpressions) {
	list := make([]intsyn.IExpression, 0, expr.Size())
	for _, element := range expr.Elements() {
		list = append(list, element)
	}
	v.AcceptListExpression(list)
}

func (v *AbstractJavaSyntaxVisitor) VisitFieldReferenceExpression(expr intsyn.IFieldReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)

	v.SafeAcceptExpression(expr.Expression())
}

func (v *AbstractJavaSyntaxVisitor) VisitFloatConstantExpression(expr intsyn.IFloatConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitIntegerConstantExpression(expr intsyn.IIntegerConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitInstanceOfExpression(expr intsyn.IInstanceOfExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	expr.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitLambdaFormalParametersExpression(expr intsyn.ILambdaFormalParametersExpression) {
	v.SafeAcceptDeclaration(expr.FormalParameters())
	expr.Statements().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitLambdaIdentifiersExpression(expr intsyn.ILambdaIdentifiersExpression) {
	v.SafeAcceptStatement(expr.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitLengthExpression(expr intsyn.ILengthExpression) {
	expr.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitLocalVariableReferenceExpression(expr intsyn.ILocalVariableReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitLongConstantExpression(expr intsyn.ILongConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitMethodInvocationExpression(expr intsyn.IMethodInvocationExpression) {
	expr.Expression().Accept(v)
	v.SafeAcceptTypeArgumentVisitable(expr.NonWildcardTypeArguments().(*_type.WildcardSuperTypeArgument))
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *AbstractJavaSyntaxVisitor) VisitMethodReferenceExpression(expr intsyn.IMethodReferenceExpression) {
	expr.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitNewArray(expr intsyn.INewArray) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.DimensionExpressionList())
}

func (v *AbstractJavaSyntaxVisitor) VisitNewExpression(expr intsyn.INewExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *AbstractJavaSyntaxVisitor) VisitNewInitializedArray(expr intsyn.INewInitializedArray) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptDeclaration(expr.ArrayInitializer())
}

func (v *AbstractJavaSyntaxVisitor) VisitNoExpression(expr intsyn.INoExpression) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitNullExpression(expr intsyn.INullExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitObjectTypeReferenceExpression(expr intsyn.IObjectTypeReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitParenthesesExpression(expr intsyn.IParenthesesExpression) {
	expr.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitPostOperatorExpression(expr intsyn.IPostOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitPreOperatorExpression(expr intsyn.IPreOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitStringConstantExpression(expr intsyn.IStringConstantExpression) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitSuperConstructorInvocationExpression(expr intsyn.ISuperConstructorInvocationExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *AbstractJavaSyntaxVisitor) VisitSuperExpression(expr intsyn.ISuperExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitTernaryOperatorExpression(expr intsyn.ITernaryOperatorExpression) {
	expr.Condition().Accept(v)
	expr.TrueExpression().Accept(v)
	expr.FalseExpression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitThisExpression(expr intsyn.IThisExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitTypeReferenceDotClassExpression(expr intsyn.ITypeReferenceDotClassExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

// --- IReferenceVisitor ---

func (v *AbstractJavaSyntaxVisitor) VisitAnnotationElementValue(ref intsyn.IAnnotationElementValue) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *AbstractJavaSyntaxVisitor) VisitAnnotationReference(ref intsyn.IAnnotationReference) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *AbstractJavaSyntaxVisitor) VisitAnnotationReferences(ref intsyn.IAnnotationReferences) {
	list := make([]intsyn.IReference, 0, ref.Size())
	for _, element := range ref.Elements() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *AbstractJavaSyntaxVisitor) VisitElementValueArrayInitializerElementValue(ref intsyn.IElementValueArrayInitializerElementValue) {
	v.SafeAcceptReference(ref.ElementValueArrayInitializer())
}

func (v *AbstractJavaSyntaxVisitor) VisitElementValues(ref intsyn.IElementValues) {
	list := make([]intsyn.IReference, 0, ref.Size())
	for _, element := range ref.Elements() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *AbstractJavaSyntaxVisitor) VisitElementValuePair(ref intsyn.IElementValuePair) {
	ref.ElementValue().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitElementValuePairs(ref intsyn.IElementValuePairs) {
	list := make([]intsyn.IReference, 0, ref.Size())
	for _, element := range ref.Elements() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *AbstractJavaSyntaxVisitor) VisitExpressionElementValue(ref intsyn.IExpressionElementValue) {
	ref.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitInnerObjectReference(ref intsyn.IInnerObjectReference) {
	v.VisitInnerObjectType(ref.(intsyn.IInnerObjectType))
}

func (v *AbstractJavaSyntaxVisitor) VisitObjectReference(ref intsyn.IObjectReference) {
	v.VisitObjectType(ref.(intsyn.IObjectType))
}

// --- IStatementVisitor ---

func (v *AbstractJavaSyntaxVisitor) VisitAssertStatement(stat intsyn.IAssertStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptExpression(stat.Message())
}

func (v *AbstractJavaSyntaxVisitor) VisitBreakStatement(stat intsyn.IBreakStatement) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitByteCodeStatement(stat intsyn.IByteCodeStatement) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitCommentStatement(stat intsyn.ICommentStatement) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitContinueStatement(stat intsyn.IContinueStatement) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitDoWhileStatement(stat intsyn.IDoWhileStatement) {
	v.SafeAcceptExpression(stat.Condition())
	v.SafeAcceptStatement(stat.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitExpressionStatement(stat intsyn.IExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitForEachStatement(stat intsyn.IForEachStatement) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	stat.Expression().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitForStatement(stat intsyn.IForStatement) {
	v.SafeAcceptDeclaration(stat.Declaration())
	v.SafeAcceptExpression(stat.Init())
	v.SafeAcceptExpression(stat.Condition())
	v.SafeAcceptExpression(stat.Update())
	v.SafeAcceptStatement(stat.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitIfStatement(stat intsyn.IIfStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitIfElseStatement(stat intsyn.IIfElseStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
	stat.ElseStatements().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitLabelStatement(stat intsyn.ILabelStatement) {
	v.SafeAcceptStatement(stat.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitLambdaExpressionStatement(stat intsyn.ILambdaExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitLocalVariableDeclarationStatement(stat intsyn.ILocalVariableDeclarationStatement) {
	v.VisitLocalVariableDeclaration(&stat.(*statement.LocalVariableDeclarationStatement).LocalVariableDeclaration)
}

func (v *AbstractJavaSyntaxVisitor) VisitNoStatement(stat intsyn.INoStatement) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitReturnExpressionStatement(stat intsyn.IReturnExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitReturnStatement(stat intsyn.IReturnStatement) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitStatements(stat intsyn.IStatements) {
	list := make([]intsyn.IStatement, 0, stat.Size())
	for _, element := range stat.Elements() {
		list = append(list, element)
	}
	v.AcceptListStatement(list)
}

func (v *AbstractJavaSyntaxVisitor) VisitSwitchStatement(stat intsyn.ISwitchStatement) {
	stat.Condition().Accept(v)
	v.AcceptListStatement(stat.List())
}

func (v *AbstractJavaSyntaxVisitor) VisitSwitchStatementDefaultLabel(stat intsyn.IDefaultLabel) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitSwitchStatementExpressionLabel(stat intsyn.IExpressionLabel) {
	stat.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitSwitchStatementLabelBlock(stat intsyn.ILabelBlock) {
	stat.Label().Accept(v)
	stat.Statements().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitSwitchStatementMultiLabelsBlock(stat intsyn.IMultiLabelsBlock) {
	v.SafeAcceptListStatement(stat.List())
	stat.Statements().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitSynchronizedStatement(stat intsyn.ISynchronizedStatement) {
	stat.Monitor().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitThrowStatement(stat intsyn.IThrowStatement) {
	stat.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitTryStatement(stat intsyn.ITryStatement) {
	v.SafeAcceptListStatement(stat.ResourceList())
	stat.TryStatements().Accept(v)
	v.SafeAcceptListStatement(stat.CatchClauseList())
	v.SafeAcceptStatement(stat.FinallyStatements())
}

func (v *AbstractJavaSyntaxVisitor) VisitTryStatementResource(stat intsyn.IResource) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	stat.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitTryStatementCatchClause(stat intsyn.ICatchClause) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitTypeDeclarationStatement(stat intsyn.ITypeDeclarationStatement) {
	stat.TypeDeclaration().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitWhileStatement(stat intsyn.IWhileStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

// --- ITypeVisitor ---
// func (v *AbstractJavaSyntaxVisitor) VisitPrimitiveType(y *PrimitiveType)     {}
// func (v *AbstractJavaSyntaxVisitor) VisitObjectType(y *ObjectType)           {}
// func (v *AbstractJavaSyntaxVisitor) VisitInnerObjectType(y *InnerObjectType) {}

func (v *AbstractJavaSyntaxVisitor) VisitTypes(types intsyn.ITypes) {
	for _, value := range types.Elements() {
		value.AcceptTypeVisitor(v)
	}
}

//func (v *AbstractJavaSyntaxVisitor) VisitGenericType(y *GenericType)         {}

// --- ITypeParameterVisitor --- //

func (v *AbstractJavaSyntaxVisitor) VisitTypeParameter(parameter intsyn.ITypeParameter) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitTypeParameterWithTypeBounds(parameter intsyn.ITypeParameterWithTypeBounds) {
	parameter.TypeBounds().AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitTypeParameters(parameters intsyn.ITypeParameters) {
	for _, param := range parameters.Elements() {
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

func (v *AbstractJavaSyntaxVisitor) AcceptListDeclaration(list []intsyn.IDeclaration) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) AcceptListExpression(list []intsyn.IExpression) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) AcceptListReference(list []intsyn.IReference) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) AcceptListStatement(list []intsyn.IStatement) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptDeclaration(decl intsyn.IDeclaration) {
	if decl != nil {
		decl.Accept(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptExpression(expr intsyn.IExpression) {
	if expr != nil {
		expr.Accept(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptReference(ref intsyn.IReference) {
	if ref != nil {
		ref.Accept(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptStatement(list intsyn.IStatement) {
	if list != nil {
		list.Accept(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptType(list intsyn.IType) {
	if list != nil {
		list.AcceptTypeVisitor(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptTypeParameter(list intsyn.ITypeParameter) {
	if list != nil {
		list.AcceptTypeParameterVisitor(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptListDeclaration(list []intsyn.IDeclaration) {
	if list != nil {
		for _, value := range list {
			value.Accept(v)
		}
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptListConstant(list []intsyn.IConstant) {
	if list != nil {
		for _, value := range list {
			value.Accept(v)
		}
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptListStatement(list []intsyn.IStatement) {
	if list != nil {
		for _, value := range list {
			value.Accept(v)
		}
	}
}
