package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax/statement"
)

func NewUpdateJavaSyntaxTreeStep0Visitor(typeMaker intsrv.ITypeMaker) intsrv.IUpdateJavaSyntaxTreeStep0Visitor {
	return &UpdateJavaSyntaxTreeStep0Visitor{
		updateOuterFieldTypeVisitor:   NewUpdateOuterFieldTypeVisitor(typeMaker),
		updateBridgeMethodTypeVisitor: NewUpdateBridgeMethodTypeVisitor(typeMaker),
	}
}

type UpdateJavaSyntaxTreeStep0Visitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	updateOuterFieldTypeVisitor   intsrv.IUpdateOuterFieldTypeVisitor
	updateBridgeMethodTypeVisitor intsrv.IUpdateBridgeMethodTypeVisitor
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitBodyDeclaration(decl intmod.IBodyDeclaration) {
	bodyDeclaration := decl.(intsrv.IClassFileBodyDeclaration)
	genericTypesSupported := bodyDeclaration.ClassFile().MajorVersion() >= 49 // (majorVersion >= Java 5)

	if genericTypesSupported {
		v.updateOuterFieldTypeVisitor.SafeAcceptListDeclaration(
			ConvertTypeDeclarations(bodyDeclaration.InnerTypeDeclarations()))
		v.updateBridgeMethodTypeVisitor.VisitBodyDeclaration(decl)
	}
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitClassDeclaration(decl intmod.IClassDeclaration) {
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitInterfaceDeclaration(decl intmod.IInterfaceDeclaration) {
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitAnnotationDeclaration(_ intmod.IAnnotationDeclaration) {
}
func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitEnumDeclaration(_ intmod.IEnumDeclaration) {}

// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$                           $$$
// $$$ AbstractJavaSyntaxVisitor $$$
// $$$                           $$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitCompilationUnit(compilationUnit intmod.ICompilationUnit) {
	compilationUnit.TypeDeclarations().AcceptDeclaration(v)
}

// --- DeclarationVisitor ---

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitArrayVariableInitializer(decl intmod.IArrayVariableInitializer) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitConstructorDeclaration(decl intmod.IConstructorDeclaration) {
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.FormalParameters())
	v.SafeAcceptType(decl.ExceptionTypes())
	v.SafeAcceptStatement(decl.Statements())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitEnumDeclarationConstant(decl intmod.IConstant) {
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptExpression(decl.Arguments())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitExpressionVariableInitializer(decl intmod.IExpressionVariableInitializer) {
	decl.Expression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitFieldDeclaration(decl intmod.IFieldDeclaration) {
	t := decl.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptReference(decl.AnnotationReferences())
	decl.FieldDeclarators().AcceptDeclaration(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitFieldDeclarator(decl intmod.IFieldDeclarator) {
	v.SafeAcceptDeclaration(decl.VariableInitializer())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitFieldDeclarators(decl intmod.IFieldDeclarators) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitFormalParameter(decl intmod.IFormalParameter) {
	t := decl.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitFormalParameters(decl intmod.IFormalParameters) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitInstanceInitializerDeclaration(decl intmod.IInstanceInitializerDeclaration) {
	v.SafeAcceptStatement(decl.Statements())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitLocalVariableDeclaration(decl intmod.ILocalVariableDeclaration) {
	v.SafeAcceptDeclaration(decl.LocalVariableDeclarators())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitLocalVariableDeclarator(decl intmod.ILocalVariableDeclarator) {
	v.SafeAcceptDeclaration(decl.VariableInitializer())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitLocalVariableDeclarators(decl intmod.ILocalVariableDeclarators) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitMethodDeclaration(decl intmod.IMethodDeclaration) {
	t := decl.ReturnedType()
	t.AcceptTypeVisitor(v)

	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.FormalParameters())
	v.SafeAcceptType(decl.ExceptionTypes())
	v.SafeAcceptStatement(decl.Statements())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitMemberDeclarations(decl intmod.IMemberDeclarations) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitModuleDeclaration(_ intmod.IModuleDeclaration) {
	// Empty
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitStaticInitializerDeclaration(_ intmod.IStaticInitializerDeclaration) {
	// Empty
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitTypeDeclarations(decl intmod.ITypeDeclarations) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

// --- IExpressionVisitor ---
func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitArrayExpression(expr intmod.IArrayExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)

	expr.Expression().Accept(v)
	expr.Index().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitBinaryOperatorExpression(expr intmod.IBinaryOperatorExpression) {
	expr.LeftExpression().Accept(v)
	expr.RightExpression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitBooleanExpression(_ intmod.IBooleanExpression) {
	// Empty
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitCastExpression(expr intmod.ICastExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	expr.Expression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitCommentExpression(_ intmod.ICommentExpression) {
	// Empty
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitConstructorInvocationExpression(expression intmod.IConstructorInvocationExpression) {
	t := expression.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expression.Parameters())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitConstructorReferenceExpression(expr intmod.IConstructorReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitDoubleConstantExpression(expr intmod.IDoubleConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitEnumConstantReferenceExpression(expr intmod.IEnumConstantReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitExpressions(expr intmod.IExpressions) {
	list := make([]intmod.IExpression, 0, expr.Size())
	for _, element := range expr.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListExpression(list)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitFieldReferenceExpression(expr intmod.IFieldReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)

	v.SafeAcceptExpression(expr.Expression())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitFloatConstantExpression(expr intmod.IFloatConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitIntegerConstantExpression(expr intmod.IIntegerConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitInstanceOfExpression(expr intmod.IInstanceOfExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	expr.Expression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitLambdaFormalParametersExpression(expr intmod.ILambdaFormalParametersExpression) {
	v.SafeAcceptDeclaration(expr.FormalParameters())
	expr.Statements().AcceptStatement(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitLambdaIdentifiersExpression(expr intmod.ILambdaIdentifiersExpression) {
	v.SafeAcceptStatement(expr.Statements())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitLengthExpression(expr intmod.ILengthExpression) {
	expr.Expression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitLocalVariableReferenceExpression(expr intmod.ILocalVariableReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitLongConstantExpression(expr intmod.ILongConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitMethodInvocationExpression(expr intmod.IMethodInvocationExpression) {
	expr.Expression().Accept(v)
	v.SafeAcceptTypeArgumentVisitable(expr.NonWildcardTypeArguments().(intmod.IWildcardSuperTypeArgument))
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitMethodReferenceExpression(expr intmod.IMethodReferenceExpression) {
	expr.Expression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitNewArray(expr intmod.INewArray) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.DimensionExpressionList())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitNewExpression(expr intmod.INewExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitNewInitializedArray(expr intmod.INewInitializedArray) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptDeclaration(expr.ArrayInitializer())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitNoExpression(_ intmod.INoExpression) {
	// Empty
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitNullExpression(expr intmod.INullExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitObjectTypeReferenceExpression(expr intmod.IObjectTypeReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitParenthesesExpression(expr intmod.IParenthesesExpression) {
	expr.Expression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitPostOperatorExpression(expr intmod.IPostOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitPreOperatorExpression(expr intmod.IPreOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitStringConstantExpression(_ intmod.IStringConstantExpression) {
	// Empty
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitSuperConstructorInvocationExpression(expr intmod.ISuperConstructorInvocationExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitSuperExpression(expr intmod.ISuperExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitTernaryOperatorExpression(expr intmod.ITernaryOperatorExpression) {
	expr.Condition().Accept(v)
	expr.TrueExpression().Accept(v)
	expr.FalseExpression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitThisExpression(expr intmod.IThisExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitTypeReferenceDotClassExpression(expr intmod.ITypeReferenceDotClassExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

// --- IReferenceVisitor ---

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitAnnotationElementValue(ref intmod.IAnnotationElementValue) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitAnnotationReference(ref intmod.IAnnotationReference) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitAnnotationReferences(ref intmod.IAnnotationReferences) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitElementValueArrayInitializerElementValue(ref intmod.IElementValueArrayInitializerElementValue) {
	v.SafeAcceptReference(ref.ElementValueArrayInitializer())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitElementValues(ref intmod.IElementValues) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitElementValuePair(ref intmod.IElementValuePair) {
	ref.ElementValue().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitElementValuePairs(ref intmod.IElementValuePairs) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitExpressionElementValue(ref intmod.IExpressionElementValue) {
	ref.Expression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitInnerObjectReference(ref intmod.IInnerObjectReference) {
	v.VisitInnerObjectType(ref.(intmod.IInnerObjectType))
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitObjectReference(ref intmod.IObjectReference) {
	v.VisitObjectType(ref.(intmod.IObjectType))
}

// --- IStatementVisitor ---

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitAssertStatement(stat intmod.IAssertStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptExpression(stat.Message())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitBreakStatement(_ intmod.IBreakStatement) {
	// Empty
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitByteCodeStatement(_ intmod.IByteCodeStatement) {
	// Empty
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitCommentStatement(_ intmod.ICommentStatement) {
	// Empty
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitContinueStatement(_ intmod.IContinueStatement) {
	// Empty
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitDoWhileStatement(stat intmod.IDoWhileStatement) {
	v.SafeAcceptExpression(stat.Condition())
	v.SafeAcceptStatement(stat.Statements())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitExpressionStatement(stat intmod.IExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitForEachStatement(stat intmod.IForEachStatement) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	stat.Expression().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitForStatement(stat intmod.IForStatement) {
	v.SafeAcceptDeclaration(stat.Declaration())
	v.SafeAcceptExpression(stat.Init())
	v.SafeAcceptExpression(stat.Condition())
	v.SafeAcceptExpression(stat.Update())
	v.SafeAcceptStatement(stat.Statements())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitIfStatement(stat intmod.IIfStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitIfElseStatement(stat intmod.IIfElseStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
	stat.ElseStatements().AcceptStatement(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitLabelStatement(stat intmod.ILabelStatement) {
	v.SafeAcceptStatement(stat.Statements())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitLambdaExpressionStatement(stat intmod.ILambdaExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitLocalVariableDeclarationStatement(stat intmod.ILocalVariableDeclarationStatement) {
	v.VisitLocalVariableDeclaration(&stat.(*statement.LocalVariableDeclarationStatement).LocalVariableDeclaration)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitNoStatement(_ intmod.INoStatement) {
	// Empty
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitReturnExpressionStatement(stat intmod.IReturnExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitReturnStatement(_ intmod.IReturnStatement) {
	// Empty
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitStatements(stat intmod.IStatements) {
	list := make([]intmod.IStatement, 0, stat.Size())
	for _, element := range stat.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListStatement(list)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitSwitchStatement(stat intmod.ISwitchStatement) {
	stat.Condition().Accept(v)
	v.AcceptListStatement(stat.List())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitSwitchStatementDefaultLabel(_ intmod.IDefaultLabel) {
	// Empty
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitSwitchStatementExpressionLabel(stat intmod.IExpressionLabel) {
	stat.Expression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitSwitchStatementLabelBlock(stat intmod.ILabelBlock) {
	stat.Label().AcceptStatement(v)
	stat.Statements().AcceptStatement(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitSwitchStatementMultiLabelsBlock(stat intmod.IMultiLabelsBlock) {
	v.SafeAcceptListStatement(stat.ToSlice())
	stat.Statements().AcceptStatement(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitSynchronizedStatement(stat intmod.ISynchronizedStatement) {
	stat.Monitor().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitThrowStatement(stat intmod.IThrowStatement) {
	stat.Expression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitTryStatement(stat intmod.ITryStatement) {
	v.SafeAcceptListStatement(stat.ResourceList())
	stat.TryStatements().AcceptStatement(v)
	v.SafeAcceptListStatement(stat.CatchClauseList())
	v.SafeAcceptStatement(stat.FinallyStatements())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitTryStatementResource(stat intmod.IResource) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	stat.Expression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitTryStatementCatchClause(stat intmod.ICatchClause) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitTypeDeclarationStatement(stat intmod.ITypeDeclarationStatement) {
	stat.TypeDeclaration().AcceptDeclaration(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitWhileStatement(stat intmod.IWhileStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

// --- ITypeVisitor ---

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitTypes(types intmod.ITypes) {
	for _, value := range types.ToSlice() {
		value.AcceptTypeVisitor(v)
	}
}

// --- ITypeParameterVisitor --- //

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitTypeParameter(_ intmod.ITypeParameter) {
	// Empty
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitTypeParameterWithTypeBounds(parameter intmod.ITypeParameterWithTypeBounds) {
	parameter.TypeBounds().AcceptTypeVisitor(v)
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitTypeParameters(parameters intmod.ITypeParameters) {
	for _, param := range parameters.ToSlice() {
		param.AcceptTypeParameterVisitor(v)
	}
}

// --- ITypeArgumentVisitor ---

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitTypeDeclaration(decl intmod.ITypeDeclaration) {
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) AcceptListDeclaration(list []intmod.IDeclaration) {
	for _, value := range list {
		value.AcceptDeclaration(v)
	}
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) AcceptListExpression(list []intmod.IExpression) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) AcceptListReference(list []intmod.IReference) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) AcceptListStatement(list []intmod.IStatement) {
	for _, value := range list {
		value.AcceptStatement(v)
	}
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) SafeAcceptDeclaration(decl intmod.IDeclaration) {
	if decl != nil {
		decl.AcceptDeclaration(v)
	}
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) SafeAcceptExpression(expr intmod.IExpression) {
	if expr != nil {
		expr.Accept(v)
	}
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) SafeAcceptReference(ref intmod.IReference) {
	if ref != nil {
		ref.Accept(v)
	}
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) SafeAcceptStatement(list intmod.IStatement) {
	if list != nil {
		list.AcceptStatement(v)
	}
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) SafeAcceptType(list intmod.IType) {
	if list != nil {
		list.AcceptTypeVisitor(v)
	}
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) SafeAcceptTypeParameter(list intmod.ITypeParameter) {
	if list != nil {
		list.AcceptTypeParameterVisitor(v)
	}
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) SafeAcceptListDeclaration(list []intmod.IDeclaration) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) SafeAcceptListConstant(list []intmod.IConstant) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) SafeAcceptListStatement(list []intmod.IStatement) {
	if list != nil {
		for _, value := range list {
			value.AcceptStatement(v)
		}
	}
}
