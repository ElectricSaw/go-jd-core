package javasyntax

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax/statement"
	_type "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/type"
)

type AbstractJavaSyntaxVisitor struct {
	_type.AbstractTypeArgumentVisitor
}

func (v *AbstractJavaSyntaxVisitor) VisitCompilationUnit(compilationUnit *CompilationUnit) {
	compilationUnit.TypeDeclarations().AcceptDeclaration(v)
}

// --- DeclarationVisitor ---

func (v *AbstractJavaSyntaxVisitor) VisitAnnotationDeclaration(decl intmod.IAnnotationDeclaration) {
	v.SafeAcceptDeclaration(decl.AnnotationDeclarators())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *AbstractJavaSyntaxVisitor) VisitArrayVariableInitializer(decl intmod.IArrayVariableInitializer) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *AbstractJavaSyntaxVisitor) VisitBodyDeclaration(decl intmod.IBodyDeclaration) {
	v.SafeAcceptDeclaration(decl.MemberDeclarations())
}

func (v *AbstractJavaSyntaxVisitor) VisitClassDeclaration(decl intmod.IClassDeclaration) {
	superType := decl.SuperType()

	if superType != nil {
		superType.AcceptTypeVisitor(v)
	}

	v.SafeAcceptTypeParameter(decl.TypeParameters())
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *AbstractJavaSyntaxVisitor) VisitConstructorDeclaration(decl intmod.IConstructorDeclaration) {
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.FormalParameters())
	v.SafeAcceptType(decl.ExceptionTypes())
	v.SafeAcceptStatement(decl.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitEnumDeclaration(decl intmod.IEnumDeclaration) {
	v.VisitTypeDeclaration(decl.(intmod.ITypeDeclaration))
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptListConstant(decl.Constants())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *AbstractJavaSyntaxVisitor) VisitEnumDeclarationConstant(decl intmod.IConstant) {
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptExpression(decl.Arguments())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *AbstractJavaSyntaxVisitor) VisitExpressionVariableInitializer(decl intmod.IExpressionVariableInitializer) {
	decl.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitFieldDeclaration(decl intmod.IFieldDeclaration) {
	t := decl.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptReference(decl.AnnotationReferences())
	decl.FieldDeclarators().AcceptDeclaration(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitFieldDeclarator(decl intmod.IFieldDeclarator) {
	v.SafeAcceptDeclaration(decl.VariableInitializer())
}

func (v *AbstractJavaSyntaxVisitor) VisitFieldDeclarators(decl intmod.IFieldDeclarators) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *AbstractJavaSyntaxVisitor) VisitFormalParameter(decl intmod.IFormalParameter) {
	t := decl.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *AbstractJavaSyntaxVisitor) VisitFormalParameters(decl intmod.IFormalParameters) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *AbstractJavaSyntaxVisitor) VisitInstanceInitializerDeclaration(decl intmod.IInstanceInitializerDeclaration) {
	v.SafeAcceptStatement(decl.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitInterfaceDeclaration(decl intmod.IInterfaceDeclaration) {
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *AbstractJavaSyntaxVisitor) VisitLocalVariableDeclaration(decl intmod.ILocalVariableDeclaration) {
	v.SafeAcceptDeclaration(decl.LocalVariableDeclarators())
}

func (v *AbstractJavaSyntaxVisitor) VisitLocalVariableDeclarator(decl intmod.ILocalVariableDeclarator) {
	v.SafeAcceptDeclaration(decl.VariableInitializer())
}

func (v *AbstractJavaSyntaxVisitor) VisitLocalVariableDeclarators(decl intmod.ILocalVariableDeclarators) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *AbstractJavaSyntaxVisitor) VisitMethodDeclaration(decl intmod.IMethodDeclaration) {
	t := decl.ReturnedType()
	t.AcceptTypeVisitor(v)

	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.FormalParameters())
	v.SafeAcceptType(decl.ExceptionTypes())
	v.SafeAcceptStatement(decl.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitMemberDeclarations(decl intmod.IMemberDeclarations) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *AbstractJavaSyntaxVisitor) VisitModuleDeclaration(decl intmod.IModuleDeclaration) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitStaticInitializerDeclaration(decl intmod.IStaticInitializerDeclaration) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitTypeDeclarations(decl intmod.ITypeDeclarations) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

// --- IExpressionVisitor ---
func (v *AbstractJavaSyntaxVisitor) VisitArrayExpression(expr intmod.IArrayExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)

	expr.Expression().Accept(v)
	expr.Index().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitBinaryOperatorExpression(expr intmod.IBinaryOperatorExpression) {
	expr.LeftExpression().Accept(v)
	expr.RightExpression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitBooleanExpression(expr intmod.IBooleanExpression) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitCastExpression(expr intmod.ICastExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	expr.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitCommentExpression(expr intmod.ICommentExpression) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitConstructorInvocationExpression(expression intmod.IConstructorInvocationExpression) {
	t := expression.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expression.Parameters())
}

func (v *AbstractJavaSyntaxVisitor) VisitConstructorReferenceExpression(expr intmod.IConstructorReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitDoubleConstantExpression(expr intmod.IDoubleConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitEnumConstantReferenceExpression(expr intmod.IEnumConstantReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitExpressions(expr intmod.IExpressions) {
	list := make([]intmod.IExpression, 0, expr.Size())
	for _, element := range expr.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListExpression(list)
}

func (v *AbstractJavaSyntaxVisitor) VisitFieldReferenceExpression(expr intmod.IFieldReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)

	v.SafeAcceptExpression(expr.Expression())
}

func (v *AbstractJavaSyntaxVisitor) VisitFloatConstantExpression(expr intmod.IFloatConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitIntegerConstantExpression(expr intmod.IIntegerConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitInstanceOfExpression(expr intmod.IInstanceOfExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	expr.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitLambdaFormalParametersExpression(expr intmod.ILambdaFormalParametersExpression) {
	v.SafeAcceptDeclaration(expr.FormalParameters())
	expr.Statements().AcceptStatement(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitLambdaIdentifiersExpression(expr intmod.ILambdaIdentifiersExpression) {
	v.SafeAcceptStatement(expr.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitLengthExpression(expr intmod.ILengthExpression) {
	expr.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitLocalVariableReferenceExpression(expr intmod.ILocalVariableReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitLongConstantExpression(expr intmod.ILongConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitMethodInvocationExpression(expr intmod.IMethodInvocationExpression) {
	expr.Expression().Accept(v)
	v.SafeAcceptTypeArgumentVisitable(expr.NonWildcardTypeArguments().(intmod.IWildcardSuperTypeArgument))
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *AbstractJavaSyntaxVisitor) VisitMethodReferenceExpression(expr intmod.IMethodReferenceExpression) {
	expr.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitNewArray(expr intmod.INewArray) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.DimensionExpressionList())
}

func (v *AbstractJavaSyntaxVisitor) VisitNewExpression(expr intmod.INewExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *AbstractJavaSyntaxVisitor) VisitNewInitializedArray(expr intmod.INewInitializedArray) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptDeclaration(expr.ArrayInitializer())
}

func (v *AbstractJavaSyntaxVisitor) VisitNoExpression(expr intmod.INoExpression) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitNullExpression(expr intmod.INullExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitObjectTypeReferenceExpression(expr intmod.IObjectTypeReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitParenthesesExpression(expr intmod.IParenthesesExpression) {
	expr.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitPostOperatorExpression(expr intmod.IPostOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitPreOperatorExpression(expr intmod.IPreOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitStringConstantExpression(expr intmod.IStringConstantExpression) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitSuperConstructorInvocationExpression(expr intmod.ISuperConstructorInvocationExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *AbstractJavaSyntaxVisitor) VisitSuperExpression(expr intmod.ISuperExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitTernaryOperatorExpression(expr intmod.ITernaryOperatorExpression) {
	expr.Condition().Accept(v)
	expr.TrueExpression().Accept(v)
	expr.FalseExpression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitThisExpression(expr intmod.IThisExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitTypeReferenceDotClassExpression(expr intmod.ITypeReferenceDotClassExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

// --- IReferenceVisitor ---

func (v *AbstractJavaSyntaxVisitor) VisitAnnotationElementValue(ref intmod.IAnnotationElementValue) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *AbstractJavaSyntaxVisitor) VisitAnnotationReference(ref intmod.IAnnotationReference) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *AbstractJavaSyntaxVisitor) VisitAnnotationReferences(ref intmod.IAnnotationReferences) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *AbstractJavaSyntaxVisitor) VisitElementValueArrayInitializerElementValue(ref intmod.IElementValueArrayInitializerElementValue) {
	v.SafeAcceptReference(ref.ElementValueArrayInitializer())
}

func (v *AbstractJavaSyntaxVisitor) VisitElementValues(ref intmod.IElementValues) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *AbstractJavaSyntaxVisitor) VisitElementValuePair(ref intmod.IElementValuePair) {
	ref.ElementValue().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitElementValuePairs(ref intmod.IElementValuePairs) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *AbstractJavaSyntaxVisitor) VisitExpressionElementValue(ref intmod.IExpressionElementValue) {
	ref.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitInnerObjectReference(ref intmod.IInnerObjectReference) {
	v.VisitInnerObjectType(ref.(intmod.IInnerObjectType))
}

func (v *AbstractJavaSyntaxVisitor) VisitObjectReference(ref intmod.IObjectReference) {
	v.VisitObjectType(ref.(intmod.IObjectType))
}

// --- IStatementVisitor ---

func (v *AbstractJavaSyntaxVisitor) VisitAssertStatement(stat intmod.IAssertStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptExpression(stat.Message())
}

func (v *AbstractJavaSyntaxVisitor) VisitBreakStatement(stat intmod.IBreakStatement) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitByteCodeStatement(stat intmod.IByteCodeStatement) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitCommentStatement(stat intmod.ICommentStatement) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitContinueStatement(stat intmod.IContinueStatement) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitDoWhileStatement(stat intmod.IDoWhileStatement) {
	v.SafeAcceptExpression(stat.Condition())
	v.SafeAcceptStatement(stat.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitExpressionStatement(stat intmod.IExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitForEachStatement(stat intmod.IForEachStatement) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	stat.Expression().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitForStatement(stat intmod.IForStatement) {
	v.SafeAcceptDeclaration(stat.Declaration())
	v.SafeAcceptExpression(stat.Init())
	v.SafeAcceptExpression(stat.Condition())
	v.SafeAcceptExpression(stat.Update())
	v.SafeAcceptStatement(stat.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitIfStatement(stat intmod.IIfStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitIfElseStatement(stat intmod.IIfElseStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
	stat.ElseStatements().AcceptStatement(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitLabelStatement(stat intmod.ILabelStatement) {
	v.SafeAcceptStatement(stat.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitLambdaExpressionStatement(stat intmod.ILambdaExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitLocalVariableDeclarationStatement(stat intmod.ILocalVariableDeclarationStatement) {
	v.VisitLocalVariableDeclaration(&stat.(*statement.LocalVariableDeclarationStatement).LocalVariableDeclaration)
}

func (v *AbstractJavaSyntaxVisitor) VisitNoStatement(stat intmod.INoStatement) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitReturnExpressionStatement(stat intmod.IReturnExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitReturnStatement(stat intmod.IReturnStatement) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitStatements(stat intmod.IStatements) {
	list := make([]intmod.IStatement, 0, stat.Size())
	for _, element := range stat.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListStatement(list)
}

func (v *AbstractJavaSyntaxVisitor) VisitSwitchStatement(stat intmod.ISwitchStatement) {
	stat.Condition().Accept(v)
	v.AcceptListStatement(stat.List())
}

func (v *AbstractJavaSyntaxVisitor) VisitSwitchStatementDefaultLabel(stat intmod.IDefaultLabel) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitSwitchStatementExpressionLabel(stat intmod.IExpressionLabel) {
	stat.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitSwitchStatementLabelBlock(stat intmod.ILabelBlock) {
	stat.Label().AcceptStatement(v)
	stat.Statements().AcceptStatement(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitSwitchStatementMultiLabelsBlock(stat intmod.IMultiLabelsBlock) {
	v.SafeAcceptListStatement(stat.ToSlice())
	stat.Statements().AcceptStatement(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitSynchronizedStatement(stat intmod.ISynchronizedStatement) {
	stat.Monitor().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitThrowStatement(stat intmod.IThrowStatement) {
	stat.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitTryStatement(stat intmod.ITryStatement) {
	v.SafeAcceptListStatement(stat.ResourceList())
	stat.TryStatements().AcceptStatement(v)
	v.SafeAcceptListStatement(stat.CatchClauseList())
	v.SafeAcceptStatement(stat.FinallyStatements())
}

func (v *AbstractJavaSyntaxVisitor) VisitTryStatementResource(stat intmod.IResource) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	stat.Expression().Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitTryStatementCatchClause(stat intmod.ICatchClause) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *AbstractJavaSyntaxVisitor) VisitTypeDeclarationStatement(stat intmod.ITypeDeclarationStatement) {
	stat.TypeDeclaration().AcceptDeclaration(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitWhileStatement(stat intmod.IWhileStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

// --- ITypeVisitor ---
// func (v *AbstractJavaSyntaxVisitor) VisitPrimitiveType(y *PrimitiveType)     {}
// func (v *AbstractJavaSyntaxVisitor) VisitObjectType(y *ObjectType)           {}
// func (v *AbstractJavaSyntaxVisitor) VisitInnerObjectType(y *InnerObjectType) {}

func (v *AbstractJavaSyntaxVisitor) VisitTypes(types intmod.ITypes) {
	for _, value := range types.ToSlice() {
		value.AcceptTypeVisitor(v)
	}
}

//func (v *AbstractJavaSyntaxVisitor) VisitGenericType(y *GenericType)         {}

// --- ITypeParameterVisitor --- //

func (v *AbstractJavaSyntaxVisitor) VisitTypeParameter(parameter intmod.ITypeParameter) {
	// TODO: Empty
}

func (v *AbstractJavaSyntaxVisitor) VisitTypeParameterWithTypeBounds(parameter intmod.ITypeParameterWithTypeBounds) {
	parameter.TypeBounds().AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitTypeParameters(parameters intmod.ITypeParameters) {
	for _, param := range parameters.ToSlice() {
		param.AcceptTypeParameterVisitor(v)
	}
}

// --- ITypeArgumentVisitor ---

//func (v *AbstractJavaSyntaxVisitor) VisitTypeArguments(arguments intmod.ITypeArguments) {}
//
//func (v *AbstractJavaSyntaxVisitor) VisitDiamondTypeArgument(argument intmod.IDiamondTypeArgument) {}
//
//func (v *AbstractJavaSyntaxVisitor) VisitWildcardExtendsTypeArgument(argument intmod.IWildcardExtendsTypeArgument) {
//}
//
//func (v *AbstractJavaSyntaxVisitor) VisitWildcardSuperTypeArgument(argument intmod.IWildcardSuperTypeArgument) {
//}
//
//func (v *AbstractJavaSyntaxVisitor) VisitWildcardTypeArgument(argument intmod.IWildcardTypeArgument) {}
//
//func (v *AbstractJavaSyntaxVisitor) VisitPrimitiveType(t intmod.IPrimitiveType) {}
//
//func (v *AbstractJavaSyntaxVisitor) VisitObjectType(t intmod.IObjectType) {}
//
//func (v *AbstractJavaSyntaxVisitor) VisitInnerObjectType(t intmod.IInnerObjectType) {}
//
//func (v *AbstractJavaSyntaxVisitor) VisitGenericType(t intmod.IGenericType) {}

func (v *AbstractJavaSyntaxVisitor) VisitTypeDeclaration(decl intmod.ITypeDeclaration) {
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *AbstractJavaSyntaxVisitor) AcceptListDeclaration(list []intmod.IDeclaration) {
	for _, value := range list {
		value.AcceptDeclaration(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) AcceptListExpression(list []intmod.IExpression) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) AcceptListReference(list []intmod.IReference) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) AcceptListStatement(list []intmod.IStatement) {
	for _, value := range list {
		value.AcceptStatement(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptDeclaration(decl intmod.IDeclaration) {
	if decl != nil {
		decl.AcceptDeclaration(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptExpression(expr intmod.IExpression) {
	if expr != nil {
		expr.Accept(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptReference(ref intmod.IReference) {
	if ref != nil {
		ref.Accept(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptStatement(list intmod.IStatement) {
	if list != nil {
		list.AcceptStatement(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptType(list intmod.IType) {
	if list != nil {
		list.AcceptTypeVisitor(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptTypeParameter(list intmod.ITypeParameter) {
	if list != nil {
		list.AcceptTypeParameterVisitor(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptListDeclaration(list []intmod.IDeclaration) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptListConstant(list []intmod.IConstant) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptListStatement(list []intmod.IStatement) {
	if list != nil {
		for _, value := range list {
			value.AcceptStatement(v)
		}
	}
}
