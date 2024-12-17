package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax/statement"
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

// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$                           $$$
// $$$ AbstractJavaSyntaxVisitor $$$
// $$$                           $$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func (v *DeclaredSyntheticLocalVariableVisitor) VisitCompilationUnit(compilationUnit intmod.ICompilationUnit) {
	compilationUnit.TypeDeclarations().AcceptDeclaration(v)
}

// --- DeclarationVisitor ---

func (v *DeclaredSyntheticLocalVariableVisitor) VisitAnnotationDeclaration(decl intmod.IAnnotationDeclaration) {
	v.SafeAcceptDeclaration(decl.AnnotationDeclarators())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitArrayVariableInitializer(decl intmod.IArrayVariableInitializer) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitBodyDeclaration(decl intmod.IBodyDeclaration) {
	v.SafeAcceptDeclaration(decl.MemberDeclarations())
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitClassDeclaration(decl intmod.IClassDeclaration) {
	superType := decl.SuperType()

	if superType != nil {
		superType.AcceptTypeVisitor(v)
	}

	v.SafeAcceptTypeParameter(decl.TypeParameters())
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitConstructorDeclaration(decl intmod.IConstructorDeclaration) {
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.FormalParameters())
	v.SafeAcceptType(decl.ExceptionTypes())
	v.SafeAcceptStatement(decl.Statements())
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitEnumDeclaration(decl intmod.IEnumDeclaration) {
	v.VisitTypeDeclaration(decl.(intmod.ITypeDeclaration))
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptListConstant(decl.Constants())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitEnumDeclarationConstant(decl intmod.IConstant) {
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptExpression(decl.Arguments())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitExpressionVariableInitializer(decl intmod.IExpressionVariableInitializer) {
	decl.Expression().Accept(v)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitFieldDeclarator(decl intmod.IFieldDeclarator) {
	v.SafeAcceptDeclaration(decl.VariableInitializer())
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitFieldDeclarators(decl intmod.IFieldDeclarators) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitFormalParameters(decl intmod.IFormalParameters) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitInstanceInitializerDeclaration(decl intmod.IInstanceInitializerDeclaration) {
	v.SafeAcceptStatement(decl.Statements())
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitInterfaceDeclaration(decl intmod.IInterfaceDeclaration) {
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitLocalVariableDeclarator(decl intmod.ILocalVariableDeclarator) {
	v.SafeAcceptDeclaration(decl.VariableInitializer())
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitLocalVariableDeclarators(decl intmod.ILocalVariableDeclarators) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitMemberDeclarations(decl intmod.IMemberDeclarations) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitModuleDeclaration(_ intmod.IModuleDeclaration) {
	// Empty
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitStaticInitializerDeclaration(_ intmod.IStaticInitializerDeclaration) {
	// Empty
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitTypeDeclarations(decl intmod.ITypeDeclarations) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

// --- IExpressionVisitor ---

func (v *DeclaredSyntheticLocalVariableVisitor) VisitBinaryOperatorExpression(expr intmod.IBinaryOperatorExpression) {
	expr.LeftExpression().Accept(v)
	expr.RightExpression().Accept(v)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitBooleanExpression(_ intmod.IBooleanExpression) {
	// Empty
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitCommentExpression(_ intmod.ICommentExpression) {
	// Empty
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitConstructorInvocationExpression(expression intmod.IConstructorInvocationExpression) {
	t := expression.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expression.Parameters())
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitEnumConstantReferenceExpression(expr intmod.IEnumConstantReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitExpressions(expr intmod.IExpressions) {
	list := make([]intmod.IExpression, 0, expr.Size())
	for _, element := range expr.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListExpression(list)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitLambdaFormalParametersExpression(expr intmod.ILambdaFormalParametersExpression) {
	v.SafeAcceptDeclaration(expr.FormalParameters())
	expr.Statements().AcceptStatement(v)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitLambdaIdentifiersExpression(expr intmod.ILambdaIdentifiersExpression) {
	v.SafeAcceptStatement(expr.Statements())
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitLengthExpression(expr intmod.ILengthExpression) {
	expr.Expression().Accept(v)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitMethodInvocationExpression(expr intmod.IMethodInvocationExpression) {
	expr.Expression().Accept(v)
	v.SafeAcceptTypeArgumentVisitable(expr.NonWildcardTypeArguments().(intmod.IWildcardSuperTypeArgument))
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitNoExpression(_ intmod.INoExpression) {
	// Empty
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitParenthesesExpression(expr intmod.IParenthesesExpression) {
	expr.Expression().Accept(v)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitPostOperatorExpression(expr intmod.IPostOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitPreOperatorExpression(expr intmod.IPreOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitStringConstantExpression(_ intmod.IStringConstantExpression) {
	// Empty
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitSuperConstructorInvocationExpression(expr intmod.ISuperConstructorInvocationExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitSuperExpression(expr intmod.ISuperExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitTernaryOperatorExpression(expr intmod.ITernaryOperatorExpression) {
	expr.Condition().Accept(v)
	expr.TrueExpression().Accept(v)
	expr.FalseExpression().Accept(v)
}

// --- IReferenceVisitor ---

func (v *DeclaredSyntheticLocalVariableVisitor) VisitAnnotationElementValue(ref intmod.IAnnotationElementValue) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitAnnotationReference(ref intmod.IAnnotationReference) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitAnnotationReferences(ref intmod.IAnnotationReferences) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitElementValueArrayInitializerElementValue(ref intmod.IElementValueArrayInitializerElementValue) {
	v.SafeAcceptReference(ref.ElementValueArrayInitializer())
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitElementValues(ref intmod.IElementValues) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitElementValuePair(ref intmod.IElementValuePair) {
	ref.ElementValue().Accept(v)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitElementValuePairs(ref intmod.IElementValuePairs) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitExpressionElementValue(ref intmod.IExpressionElementValue) {
	ref.Expression().Accept(v)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitInnerObjectReference(ref intmod.IInnerObjectReference) {
	v.VisitInnerObjectType(ref.(intmod.IInnerObjectType))
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitObjectReference(ref intmod.IObjectReference) {
	v.VisitObjectType(ref.(intmod.IObjectType))
}

// --- IStatementVisitor ---

func (v *DeclaredSyntheticLocalVariableVisitor) VisitAssertStatement(stat intmod.IAssertStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptExpression(stat.Message())
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitBreakStatement(_ intmod.IBreakStatement) {
	// Empty
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitByteCodeStatement(_ intmod.IByteCodeStatement) {
	// Empty
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitCommentStatement(_ intmod.ICommentStatement) {
	// Empty
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitContinueStatement(_ intmod.IContinueStatement) {
	// Empty
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitDoWhileStatement(stat intmod.IDoWhileStatement) {
	v.SafeAcceptExpression(stat.Condition())
	v.SafeAcceptStatement(stat.Statements())
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitExpressionStatement(stat intmod.IExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitForStatement(stat intmod.IForStatement) {
	v.SafeAcceptDeclaration(stat.Declaration())
	v.SafeAcceptExpression(stat.Init())
	v.SafeAcceptExpression(stat.Condition())
	v.SafeAcceptExpression(stat.Update())
	v.SafeAcceptStatement(stat.Statements())
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitIfStatement(stat intmod.IIfStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitIfElseStatement(stat intmod.IIfElseStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
	stat.ElseStatements().AcceptStatement(v)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitLabelStatement(stat intmod.ILabelStatement) {
	v.SafeAcceptStatement(stat.Statements())
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitLambdaExpressionStatement(stat intmod.ILambdaExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitLocalVariableDeclarationStatement(stat intmod.ILocalVariableDeclarationStatement) {
	v.VisitLocalVariableDeclaration(&stat.(*statement.LocalVariableDeclarationStatement).LocalVariableDeclaration)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitNoStatement(_ intmod.INoStatement) {
	// Empty
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitReturnExpressionStatement(stat intmod.IReturnExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitReturnStatement(_ intmod.IReturnStatement) {
	// Empty
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitStatements(stat intmod.IStatements) {
	list := make([]intmod.IStatement, 0, stat.Size())
	for _, element := range stat.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListStatement(list)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitSwitchStatement(stat intmod.ISwitchStatement) {
	stat.Condition().Accept(v)
	v.AcceptListStatement(stat.List())
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitSwitchStatementDefaultLabel(_ intmod.IDefaultLabel) {
	// Empty
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitSwitchStatementExpressionLabel(stat intmod.IExpressionLabel) {
	stat.Expression().Accept(v)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitSynchronizedStatement(stat intmod.ISynchronizedStatement) {
	stat.Monitor().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitThrowStatement(stat intmod.IThrowStatement) {
	stat.Expression().Accept(v)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitTryStatement(stat intmod.ITryStatement) {
	v.SafeAcceptListStatement(stat.ResourceList())
	stat.TryStatements().AcceptStatement(v)
	v.SafeAcceptListStatement(stat.CatchClauseList())
	v.SafeAcceptStatement(stat.FinallyStatements())
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitTypeDeclarationStatement(stat intmod.ITypeDeclarationStatement) {
	stat.TypeDeclaration().AcceptDeclaration(v)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitWhileStatement(stat intmod.IWhileStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

// --- ITypeVisitor ---

func (v *DeclaredSyntheticLocalVariableVisitor) VisitTypes(types intmod.ITypes) {
	for _, value := range types.ToSlice() {
		value.AcceptTypeVisitor(v)
	}
}

// --- ITypeParameterVisitor --- //

func (v *DeclaredSyntheticLocalVariableVisitor) VisitTypeParameter(_ intmod.ITypeParameter) {
	// Empty
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitTypeParameterWithTypeBounds(parameter intmod.ITypeParameterWithTypeBounds) {
	parameter.TypeBounds().AcceptTypeVisitor(v)
}

func (v *DeclaredSyntheticLocalVariableVisitor) VisitTypeParameters(parameters intmod.ITypeParameters) {
	for _, param := range parameters.ToSlice() {
		param.AcceptTypeParameterVisitor(v)
	}
}

// --- ITypeArgumentVisitor ---

func (v *DeclaredSyntheticLocalVariableVisitor) VisitTypeDeclaration(decl intmod.ITypeDeclaration) {
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *DeclaredSyntheticLocalVariableVisitor) AcceptListDeclaration(list []intmod.IDeclaration) {
	for _, value := range list {
		value.AcceptDeclaration(v)
	}
}

func (v *DeclaredSyntheticLocalVariableVisitor) AcceptListExpression(list []intmod.IExpression) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *DeclaredSyntheticLocalVariableVisitor) AcceptListReference(list []intmod.IReference) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *DeclaredSyntheticLocalVariableVisitor) AcceptListStatement(list []intmod.IStatement) {
	for _, value := range list {
		value.AcceptStatement(v)
	}
}

func (v *DeclaredSyntheticLocalVariableVisitor) SafeAcceptDeclaration(decl intmod.IDeclaration) {
	if decl != nil {
		decl.AcceptDeclaration(v)
	}
}

func (v *DeclaredSyntheticLocalVariableVisitor) SafeAcceptExpression(expr intmod.IExpression) {
	if expr != nil {
		expr.Accept(v)
	}
}

func (v *DeclaredSyntheticLocalVariableVisitor) SafeAcceptReference(ref intmod.IReference) {
	if ref != nil {
		ref.Accept(v)
	}
}

func (v *DeclaredSyntheticLocalVariableVisitor) SafeAcceptStatement(list intmod.IStatement) {
	if list != nil {
		list.AcceptStatement(v)
	}
}

func (v *DeclaredSyntheticLocalVariableVisitor) SafeAcceptType(list intmod.IType) {
	if list != nil {
		list.AcceptTypeVisitor(v)
	}
}

func (v *DeclaredSyntheticLocalVariableVisitor) SafeAcceptTypeParameter(list intmod.ITypeParameter) {
	if list != nil {
		list.AcceptTypeParameterVisitor(v)
	}
}

func (v *DeclaredSyntheticLocalVariableVisitor) SafeAcceptListDeclaration(list []intmod.IDeclaration) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *DeclaredSyntheticLocalVariableVisitor) SafeAcceptListConstant(list []intmod.IConstant) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *DeclaredSyntheticLocalVariableVisitor) SafeAcceptListStatement(list []intmod.IStatement) {
	if list != nil {
		for _, value := range list {
			value.AcceptStatement(v)
		}
	}
}
