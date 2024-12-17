package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax/statement"
)

func NewChangeFrameOfLocalVariablesVisitor(localVariableMaker intsrv.ILocalVariableMaker) intsrv.IChangeFrameOfLocalVariablesVisitor {
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

// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$                           $$$
// $$$ AbstractJavaSyntaxVisitor $$$
// $$$                           $$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func (v *ChangeFrameOfLocalVariablesVisitor) VisitCompilationUnit(compilationUnit intmod.ICompilationUnit) {
	compilationUnit.TypeDeclarations().AcceptDeclaration(v)
}

// --- DeclarationVisitor ---

func (v *ChangeFrameOfLocalVariablesVisitor) VisitAnnotationDeclaration(decl intmod.IAnnotationDeclaration) {
	v.SafeAcceptDeclaration(decl.AnnotationDeclarators())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitArrayVariableInitializer(decl intmod.IArrayVariableInitializer) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitBodyDeclaration(decl intmod.IBodyDeclaration) {
	v.SafeAcceptDeclaration(decl.MemberDeclarations())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitClassDeclaration(decl intmod.IClassDeclaration) {
	superType := decl.SuperType()

	if superType != nil {
		superType.AcceptTypeVisitor(v)
	}

	v.SafeAcceptTypeParameter(decl.TypeParameters())
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitConstructorDeclaration(decl intmod.IConstructorDeclaration) {
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.FormalParameters())
	v.SafeAcceptType(decl.ExceptionTypes())
	v.SafeAcceptStatement(decl.Statements())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitEnumDeclaration(decl intmod.IEnumDeclaration) {
	v.VisitTypeDeclaration(decl.(intmod.ITypeDeclaration))
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptListConstant(decl.Constants())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitEnumDeclarationConstant(decl intmod.IConstant) {
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptExpression(decl.Arguments())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitExpressionVariableInitializer(decl intmod.IExpressionVariableInitializer) {
	decl.Expression().Accept(v)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitFieldDeclaration(decl intmod.IFieldDeclaration) {
	t := decl.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptReference(decl.AnnotationReferences())
	decl.FieldDeclarators().AcceptDeclaration(v)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitFieldDeclarator(decl intmod.IFieldDeclarator) {
	v.SafeAcceptDeclaration(decl.VariableInitializer())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitFieldDeclarators(decl intmod.IFieldDeclarators) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitFormalParameter(decl intmod.IFormalParameter) {
	t := decl.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitFormalParameters(decl intmod.IFormalParameters) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitInstanceInitializerDeclaration(decl intmod.IInstanceInitializerDeclaration) {
	v.SafeAcceptStatement(decl.Statements())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitInterfaceDeclaration(decl intmod.IInterfaceDeclaration) {
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitLocalVariableDeclaration(decl intmod.ILocalVariableDeclaration) {
	v.SafeAcceptDeclaration(decl.LocalVariableDeclarators())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitLocalVariableDeclarator(decl intmod.ILocalVariableDeclarator) {
	v.SafeAcceptDeclaration(decl.VariableInitializer())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitLocalVariableDeclarators(decl intmod.ILocalVariableDeclarators) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitMethodDeclaration(decl intmod.IMethodDeclaration) {
	t := decl.ReturnedType()
	t.AcceptTypeVisitor(v)

	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.FormalParameters())
	v.SafeAcceptType(decl.ExceptionTypes())
	v.SafeAcceptStatement(decl.Statements())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitMemberDeclarations(decl intmod.IMemberDeclarations) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitModuleDeclaration(_ intmod.IModuleDeclaration) {
	// Empty
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitStaticInitializerDeclaration(_ intmod.IStaticInitializerDeclaration) {
	// Empty
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitTypeDeclarations(decl intmod.ITypeDeclarations) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

// --- IExpressionVisitor ---
func (v *ChangeFrameOfLocalVariablesVisitor) VisitArrayExpression(expr intmod.IArrayExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)

	expr.Expression().Accept(v)
	expr.Index().Accept(v)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitBinaryOperatorExpression(expr intmod.IBinaryOperatorExpression) {
	expr.LeftExpression().Accept(v)
	expr.RightExpression().Accept(v)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitBooleanExpression(_ intmod.IBooleanExpression) {
	// Empty
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitCastExpression(expr intmod.ICastExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	expr.Expression().Accept(v)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitCommentExpression(_ intmod.ICommentExpression) {
	// Empty
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitConstructorInvocationExpression(expression intmod.IConstructorInvocationExpression) {
	t := expression.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expression.Parameters())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitExpressions(expr intmod.IExpressions) {
	list := make([]intmod.IExpression, 0, expr.Size())
	for _, element := range expr.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListExpression(list)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitFieldReferenceExpression(expr intmod.IFieldReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)

	v.SafeAcceptExpression(expr.Expression())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitInstanceOfExpression(expr intmod.IInstanceOfExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	expr.Expression().Accept(v)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitLambdaFormalParametersExpression(expr intmod.ILambdaFormalParametersExpression) {
	v.SafeAcceptDeclaration(expr.FormalParameters())
	expr.Statements().AcceptStatement(v)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitLambdaIdentifiersExpression(expr intmod.ILambdaIdentifiersExpression) {
	v.SafeAcceptStatement(expr.Statements())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitLengthExpression(expr intmod.ILengthExpression) {
	expr.Expression().Accept(v)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitMethodInvocationExpression(expr intmod.IMethodInvocationExpression) {
	expr.Expression().Accept(v)
	v.SafeAcceptTypeArgumentVisitable(expr.NonWildcardTypeArguments().(intmod.IWildcardSuperTypeArgument))
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitMethodReferenceExpression(expr intmod.IMethodReferenceExpression) {
	expr.Expression().Accept(v)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitNewArray(expr intmod.INewArray) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.DimensionExpressionList())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitNewExpression(expr intmod.INewExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitNewInitializedArray(expr intmod.INewInitializedArray) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptDeclaration(expr.ArrayInitializer())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitNoExpression(_ intmod.INoExpression) {
	// Empty
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitParenthesesExpression(expr intmod.IParenthesesExpression) {
	expr.Expression().Accept(v)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitPostOperatorExpression(expr intmod.IPostOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitPreOperatorExpression(expr intmod.IPreOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitStringConstantExpression(_ intmod.IStringConstantExpression) {
	// Empty
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitSuperConstructorInvocationExpression(expr intmod.ISuperConstructorInvocationExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitTernaryOperatorExpression(expr intmod.ITernaryOperatorExpression) {
	expr.Condition().Accept(v)
	expr.TrueExpression().Accept(v)
	expr.FalseExpression().Accept(v)
}

// --- IReferenceVisitor ---

func (v *ChangeFrameOfLocalVariablesVisitor) VisitAnnotationElementValue(ref intmod.IAnnotationElementValue) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitAnnotationReference(ref intmod.IAnnotationReference) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitAnnotationReferences(ref intmod.IAnnotationReferences) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitElementValueArrayInitializerElementValue(ref intmod.IElementValueArrayInitializerElementValue) {
	v.SafeAcceptReference(ref.ElementValueArrayInitializer())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitElementValues(ref intmod.IElementValues) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitElementValuePair(ref intmod.IElementValuePair) {
	ref.ElementValue().Accept(v)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitElementValuePairs(ref intmod.IElementValuePairs) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitExpressionElementValue(ref intmod.IExpressionElementValue) {
	ref.Expression().Accept(v)
}

// --- IStatementVisitor ---

func (v *ChangeFrameOfLocalVariablesVisitor) VisitAssertStatement(stat intmod.IAssertStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptExpression(stat.Message())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitCommentStatement(_ intmod.ICommentStatement) {
	// Empty
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitDoWhileStatement(stat intmod.IDoWhileStatement) {
	v.SafeAcceptExpression(stat.Condition())
	v.SafeAcceptStatement(stat.Statements())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitExpressionStatement(stat intmod.IExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitForEachStatement(stat intmod.IForEachStatement) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	stat.Expression().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitForStatement(stat intmod.IForStatement) {
	v.SafeAcceptDeclaration(stat.Declaration())
	v.SafeAcceptExpression(stat.Init())
	v.SafeAcceptExpression(stat.Condition())
	v.SafeAcceptExpression(stat.Update())
	v.SafeAcceptStatement(stat.Statements())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitIfStatement(stat intmod.IIfStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitIfElseStatement(stat intmod.IIfElseStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
	stat.ElseStatements().AcceptStatement(v)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitLabelStatement(stat intmod.ILabelStatement) {
	v.SafeAcceptStatement(stat.Statements())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitLambdaExpressionStatement(stat intmod.ILambdaExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitLocalVariableDeclarationStatement(stat intmod.ILocalVariableDeclarationStatement) {
	v.VisitLocalVariableDeclaration(&stat.(*statement.LocalVariableDeclarationStatement).LocalVariableDeclaration)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitNoStatement(_ intmod.INoStatement) {
	// Empty
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitReturnExpressionStatement(stat intmod.IReturnExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitReturnStatement(_ intmod.IReturnStatement) {
	// Empty
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitStatements(stat intmod.IStatements) {
	list := make([]intmod.IStatement, 0, stat.Size())
	for _, element := range stat.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListStatement(list)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitSwitchStatement(stat intmod.ISwitchStatement) {
	stat.Condition().Accept(v)
	v.AcceptListStatement(stat.List())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitSwitchStatementDefaultLabel(_ intmod.IDefaultLabel) {
	// Empty
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitSwitchStatementExpressionLabel(stat intmod.IExpressionLabel) {
	stat.Expression().Accept(v)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitSwitchStatementLabelBlock(stat intmod.ILabelBlock) {
	stat.Label().AcceptStatement(v)
	stat.Statements().AcceptStatement(v)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitSwitchStatementMultiLabelsBlock(stat intmod.IMultiLabelsBlock) {
	v.SafeAcceptListStatement(stat.ToSlice())
	stat.Statements().AcceptStatement(v)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitSynchronizedStatement(stat intmod.ISynchronizedStatement) {
	stat.Monitor().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitThrowStatement(stat intmod.IThrowStatement) {
	stat.Expression().Accept(v)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitTryStatement(stat intmod.ITryStatement) {
	v.SafeAcceptListStatement(stat.ResourceList())
	stat.TryStatements().AcceptStatement(v)
	v.SafeAcceptListStatement(stat.CatchClauseList())
	v.SafeAcceptStatement(stat.FinallyStatements())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitTryStatementResource(stat intmod.IResource) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	stat.Expression().Accept(v)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitTryStatementCatchClause(stat intmod.ICatchClause) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitTypeDeclarationStatement(stat intmod.ITypeDeclarationStatement) {
	stat.TypeDeclaration().AcceptDeclaration(v)
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitWhileStatement(stat intmod.IWhileStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

// --- ITypeVisitor ---

// --- ITypeParameterVisitor --- //

func (v *ChangeFrameOfLocalVariablesVisitor) VisitTypeParameter(_ intmod.ITypeParameter) {
	// Empty
}

func (v *ChangeFrameOfLocalVariablesVisitor) VisitTypeParameters(parameters intmod.ITypeParameters) {
	for _, param := range parameters.ToSlice() {
		param.AcceptTypeParameterVisitor(v)
	}
}

// --- ITypeArgumentVisitor ---

func (v *ChangeFrameOfLocalVariablesVisitor) VisitTypeDeclaration(decl intmod.ITypeDeclaration) {
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *ChangeFrameOfLocalVariablesVisitor) AcceptListDeclaration(list []intmod.IDeclaration) {
	for _, value := range list {
		value.AcceptDeclaration(v)
	}
}

func (v *ChangeFrameOfLocalVariablesVisitor) AcceptListExpression(list []intmod.IExpression) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *ChangeFrameOfLocalVariablesVisitor) AcceptListReference(list []intmod.IReference) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *ChangeFrameOfLocalVariablesVisitor) AcceptListStatement(list []intmod.IStatement) {
	for _, value := range list {
		value.AcceptStatement(v)
	}
}

func (v *ChangeFrameOfLocalVariablesVisitor) SafeAcceptDeclaration(decl intmod.IDeclaration) {
	if decl != nil {
		decl.AcceptDeclaration(v)
	}
}

func (v *ChangeFrameOfLocalVariablesVisitor) SafeAcceptExpression(expr intmod.IExpression) {
	if expr != nil {
		expr.Accept(v)
	}
}

func (v *ChangeFrameOfLocalVariablesVisitor) SafeAcceptReference(ref intmod.IReference) {
	if ref != nil {
		ref.Accept(v)
	}
}

func (v *ChangeFrameOfLocalVariablesVisitor) SafeAcceptStatement(list intmod.IStatement) {
	if list != nil {
		list.AcceptStatement(v)
	}
}

func (v *ChangeFrameOfLocalVariablesVisitor) SafeAcceptType(list intmod.IType) {
	if list != nil {
		list.AcceptTypeVisitor(v)
	}
}

func (v *ChangeFrameOfLocalVariablesVisitor) SafeAcceptTypeParameter(list intmod.ITypeParameter) {
	if list != nil {
		list.AcceptTypeParameterVisitor(v)
	}
}

func (v *ChangeFrameOfLocalVariablesVisitor) SafeAcceptListDeclaration(list []intmod.IDeclaration) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *ChangeFrameOfLocalVariablesVisitor) SafeAcceptListConstant(list []intmod.IConstant) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *ChangeFrameOfLocalVariablesVisitor) SafeAcceptListStatement(list []intmod.IStatement) {
	if list != nil {
		for _, value := range list {
			value.AcceptStatement(v)
		}
	}
}
