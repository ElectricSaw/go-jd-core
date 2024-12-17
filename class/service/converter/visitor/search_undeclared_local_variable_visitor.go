package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax/statement"
	"github.com/ElectricSaw/go-jd-core/class/service/converter/visitor/utils"
)

func NewSearchUndeclaredLocalVariableVisitor() intsrv.ISearchUndeclaredLocalVariableVisitor {
	return &SearchUndeclaredLocalVariableVisitor{
		variables: make([]intsrv.ILocalVariable, 0),
	}
}

type SearchUndeclaredLocalVariableVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	variables []intsrv.ILocalVariable
}

func (v *SearchUndeclaredLocalVariableVisitor) Init() {
	if v.variables == nil {
		v.variables = make([]intsrv.ILocalVariable, 0)
	}

	v.variables = v.variables[:0]
}

func (v *SearchUndeclaredLocalVariableVisitor) Variables() []intsrv.ILocalVariable {
	return v.variables
}

func (v *SearchUndeclaredLocalVariableVisitor) RemoveAll(removal []intsrv.ILocalVariable) {
	variables := make([]intsrv.ILocalVariable, len(v.variables))
	copy(variables, v.variables)

	for _, item := range removal {
		v.variables = utils.SliceInItemRemove(variables, item)
	}
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitBinaryOperatorExpression(expression intmod.IBinaryOperatorExpression) {
	if expression.LeftExpression().IsLocalVariableReferenceExpression() && expression.Operator() == "=" {
		lv := expression.LeftExpression().(intsrv.IClassFileLocalVariableReferenceExpression).
			LocalVariable().(intsrv.ILocalVariable)

		if !lv.IsDeclared() {
			v.variables = append(v.variables, lv)
		}
	}

	expression.LeftExpression().Accept(v)
	expression.RightExpression().Accept(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitDoWhileStatement(statement intmod.IDoWhileStatement) {
	v.SafeAcceptExpression(statement.Condition())
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitForEachStatement(statement intmod.IForEachStatement) {
	statement.Expression().Accept(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitForStatement(statement intmod.IForStatement) {
	v.SafeAcceptDeclaration(statement.Declaration())
	v.SafeAcceptExpression(statement.Init())
	v.SafeAcceptExpression(statement.Condition())
	v.SafeAcceptExpression(statement.Update())
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitIfStatement(statement intmod.IIfStatement) {
	statement.Condition().Accept(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitIfElseStatement(statement intmod.IIfElseStatement) {
	statement.Condition().Accept(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitLambdaExpressionStatement(statement intmod.ILambdaExpressionStatement) {
	statement.Expression().Accept(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitSwitchStatement(statement intmod.ISwitchStatement) {
	statement.Condition().Accept(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitSynchronizedStatement(statement intmod.ISynchronizedStatement) {
	statement.Monitor().Accept(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitTryStatement(statement intmod.ITryStatement) {}

func (v *SearchUndeclaredLocalVariableVisitor) VisitWhileStatement(statement intmod.IWhileStatement) {
	statement.Condition().Accept(v)
}

// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$                           $$$
// $$$ AbstractJavaSyntaxVisitor $$$
// $$$                           $$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func (v *SearchUndeclaredLocalVariableVisitor) VisitCompilationUnit(compilationUnit intmod.ICompilationUnit) {
	compilationUnit.TypeDeclarations().AcceptDeclaration(v)
}

// --- DeclarationVisitor ---

func (v *SearchUndeclaredLocalVariableVisitor) VisitAnnotationDeclaration(decl intmod.IAnnotationDeclaration) {
	v.SafeAcceptDeclaration(decl.AnnotationDeclarators())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitArrayVariableInitializer(decl intmod.IArrayVariableInitializer) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitBodyDeclaration(decl intmod.IBodyDeclaration) {
	v.SafeAcceptDeclaration(decl.MemberDeclarations())
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitClassDeclaration(decl intmod.IClassDeclaration) {
	superType := decl.SuperType()

	if superType != nil {
		superType.AcceptTypeVisitor(v)
	}

	v.SafeAcceptTypeParameter(decl.TypeParameters())
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitConstructorDeclaration(decl intmod.IConstructorDeclaration) {
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.FormalParameters())
	v.SafeAcceptType(decl.ExceptionTypes())
	v.SafeAcceptStatement(decl.Statements())
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitEnumDeclaration(decl intmod.IEnumDeclaration) {
	v.VisitTypeDeclaration(decl.(intmod.ITypeDeclaration))
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptListConstant(decl.Constants())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitEnumDeclarationConstant(decl intmod.IConstant) {
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptExpression(decl.Arguments())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitExpressionVariableInitializer(decl intmod.IExpressionVariableInitializer) {
	decl.Expression().Accept(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitFieldDeclaration(decl intmod.IFieldDeclaration) {
	t := decl.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptReference(decl.AnnotationReferences())
	decl.FieldDeclarators().AcceptDeclaration(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitFieldDeclarator(decl intmod.IFieldDeclarator) {
	v.SafeAcceptDeclaration(decl.VariableInitializer())
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitFieldDeclarators(decl intmod.IFieldDeclarators) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitFormalParameter(decl intmod.IFormalParameter) {
	t := decl.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitFormalParameters(decl intmod.IFormalParameters) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitInstanceInitializerDeclaration(decl intmod.IInstanceInitializerDeclaration) {
	v.SafeAcceptStatement(decl.Statements())
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitInterfaceDeclaration(decl intmod.IInterfaceDeclaration) {
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitLocalVariableDeclaration(decl intmod.ILocalVariableDeclaration) {
	v.SafeAcceptDeclaration(decl.LocalVariableDeclarators())
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitLocalVariableDeclarator(decl intmod.ILocalVariableDeclarator) {
	v.SafeAcceptDeclaration(decl.VariableInitializer())
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitLocalVariableDeclarators(decl intmod.ILocalVariableDeclarators) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitMethodDeclaration(decl intmod.IMethodDeclaration) {
	t := decl.ReturnedType()
	t.AcceptTypeVisitor(v)

	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.FormalParameters())
	v.SafeAcceptType(decl.ExceptionTypes())
	v.SafeAcceptStatement(decl.Statements())
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitMemberDeclarations(decl intmod.IMemberDeclarations) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitModuleDeclaration(_ intmod.IModuleDeclaration) {
	// Empty
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitStaticInitializerDeclaration(_ intmod.IStaticInitializerDeclaration) {
	// Empty
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitTypeDeclarations(decl intmod.ITypeDeclarations) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

// --- IExpressionVisitor ---
func (v *SearchUndeclaredLocalVariableVisitor) VisitArrayExpression(expr intmod.IArrayExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)

	expr.Expression().Accept(v)
	expr.Index().Accept(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitBooleanExpression(_ intmod.IBooleanExpression) {
	// Empty
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitCastExpression(expr intmod.ICastExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	expr.Expression().Accept(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitCommentExpression(_ intmod.ICommentExpression) {
	// Empty
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitConstructorInvocationExpression(expression intmod.IConstructorInvocationExpression) {
	t := expression.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expression.Parameters())
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitConstructorReferenceExpression(expr intmod.IConstructorReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitDoubleConstantExpression(expr intmod.IDoubleConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitEnumConstantReferenceExpression(expr intmod.IEnumConstantReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitExpressions(expr intmod.IExpressions) {
	list := make([]intmod.IExpression, 0, expr.Size())
	for _, element := range expr.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListExpression(list)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitFieldReferenceExpression(expr intmod.IFieldReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)

	v.SafeAcceptExpression(expr.Expression())
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitFloatConstantExpression(expr intmod.IFloatConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitIntegerConstantExpression(expr intmod.IIntegerConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitInstanceOfExpression(expr intmod.IInstanceOfExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	expr.Expression().Accept(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitLambdaFormalParametersExpression(expr intmod.ILambdaFormalParametersExpression) {
	v.SafeAcceptDeclaration(expr.FormalParameters())
	expr.Statements().AcceptStatement(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitLambdaIdentifiersExpression(expr intmod.ILambdaIdentifiersExpression) {
	v.SafeAcceptStatement(expr.Statements())
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitLengthExpression(expr intmod.ILengthExpression) {
	expr.Expression().Accept(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitLocalVariableReferenceExpression(expr intmod.ILocalVariableReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitLongConstantExpression(expr intmod.ILongConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitMethodInvocationExpression(expr intmod.IMethodInvocationExpression) {
	expr.Expression().Accept(v)
	v.SafeAcceptTypeArgumentVisitable(expr.NonWildcardTypeArguments().(intmod.IWildcardSuperTypeArgument))
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitMethodReferenceExpression(expr intmod.IMethodReferenceExpression) {
	expr.Expression().Accept(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitNewArray(expr intmod.INewArray) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.DimensionExpressionList())
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitNewExpression(expr intmod.INewExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitNewInitializedArray(expr intmod.INewInitializedArray) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptDeclaration(expr.ArrayInitializer())
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitNoExpression(_ intmod.INoExpression) {
	// Empty
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitNullExpression(expr intmod.INullExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitObjectTypeReferenceExpression(expr intmod.IObjectTypeReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitParenthesesExpression(expr intmod.IParenthesesExpression) {
	expr.Expression().Accept(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitPostOperatorExpression(expr intmod.IPostOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitPreOperatorExpression(expr intmod.IPreOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitStringConstantExpression(_ intmod.IStringConstantExpression) {
	// Empty
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitSuperConstructorInvocationExpression(expr intmod.ISuperConstructorInvocationExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitSuperExpression(expr intmod.ISuperExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitTernaryOperatorExpression(expr intmod.ITernaryOperatorExpression) {
	expr.Condition().Accept(v)
	expr.TrueExpression().Accept(v)
	expr.FalseExpression().Accept(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitThisExpression(expr intmod.IThisExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitTypeReferenceDotClassExpression(expr intmod.ITypeReferenceDotClassExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

// --- IReferenceVisitor ---

func (v *SearchUndeclaredLocalVariableVisitor) VisitAnnotationElementValue(ref intmod.IAnnotationElementValue) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitAnnotationReference(ref intmod.IAnnotationReference) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitAnnotationReferences(ref intmod.IAnnotationReferences) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitElementValueArrayInitializerElementValue(ref intmod.IElementValueArrayInitializerElementValue) {
	v.SafeAcceptReference(ref.ElementValueArrayInitializer())
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitElementValues(ref intmod.IElementValues) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitElementValuePair(ref intmod.IElementValuePair) {
	ref.ElementValue().Accept(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitElementValuePairs(ref intmod.IElementValuePairs) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitExpressionElementValue(ref intmod.IExpressionElementValue) {
	ref.Expression().Accept(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitInnerObjectReference(ref intmod.IInnerObjectReference) {
	v.VisitInnerObjectType(ref.(intmod.IInnerObjectType))
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitObjectReference(ref intmod.IObjectReference) {
	v.VisitObjectType(ref.(intmod.IObjectType))
}

// --- IStatementVisitor ---

func (v *SearchUndeclaredLocalVariableVisitor) VisitAssertStatement(stat intmod.IAssertStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptExpression(stat.Message())
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitBreakStatement(_ intmod.IBreakStatement) {
	// Empty
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitByteCodeStatement(_ intmod.IByteCodeStatement) {
	// Empty
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitCommentStatement(_ intmod.ICommentStatement) {
	// Empty
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitContinueStatement(_ intmod.IContinueStatement) {
	// Empty
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitExpressionStatement(stat intmod.IExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitLabelStatement(stat intmod.ILabelStatement) {
	v.SafeAcceptStatement(stat.Statements())
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitLocalVariableDeclarationStatement(stat intmod.ILocalVariableDeclarationStatement) {
	v.VisitLocalVariableDeclaration(&stat.(*statement.LocalVariableDeclarationStatement).LocalVariableDeclaration)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitNoStatement(_ intmod.INoStatement) {
	// Empty
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitReturnExpressionStatement(stat intmod.IReturnExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitReturnStatement(_ intmod.IReturnStatement) {
	// Empty
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitStatements(stat intmod.IStatements) {
	list := make([]intmod.IStatement, 0, stat.Size())
	for _, element := range stat.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListStatement(list)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitSwitchStatementDefaultLabel(_ intmod.IDefaultLabel) {
	// Empty
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitSwitchStatementExpressionLabel(stat intmod.IExpressionLabel) {
	stat.Expression().Accept(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitSwitchStatementLabelBlock(stat intmod.ILabelBlock) {
	stat.Label().AcceptStatement(v)
	stat.Statements().AcceptStatement(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitSwitchStatementMultiLabelsBlock(stat intmod.IMultiLabelsBlock) {
	v.SafeAcceptListStatement(stat.ToSlice())
	stat.Statements().AcceptStatement(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitThrowStatement(stat intmod.IThrowStatement) {
	stat.Expression().Accept(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitTryStatementResource(stat intmod.IResource) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	stat.Expression().Accept(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitTryStatementCatchClause(stat intmod.ICatchClause) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitTypeDeclarationStatement(stat intmod.ITypeDeclarationStatement) {
	stat.TypeDeclaration().AcceptDeclaration(v)
}

// --- ITypeVisitor ---

func (v *SearchUndeclaredLocalVariableVisitor) VisitTypes(types intmod.ITypes) {
	for _, value := range types.ToSlice() {
		value.AcceptTypeVisitor(v)
	}
}

// --- ITypeParameterVisitor --- //

func (v *SearchUndeclaredLocalVariableVisitor) VisitTypeParameter(_ intmod.ITypeParameter) {
	// Empty
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitTypeParameterWithTypeBounds(parameter intmod.ITypeParameterWithTypeBounds) {
	parameter.TypeBounds().AcceptTypeVisitor(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitTypeParameters(parameters intmod.ITypeParameters) {
	for _, param := range parameters.ToSlice() {
		param.AcceptTypeParameterVisitor(v)
	}
}

// --- ITypeArgumentVisitor ---

func (v *SearchUndeclaredLocalVariableVisitor) VisitTypeDeclaration(decl intmod.ITypeDeclaration) {
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *SearchUndeclaredLocalVariableVisitor) AcceptListDeclaration(list []intmod.IDeclaration) {
	for _, value := range list {
		value.AcceptDeclaration(v)
	}
}

func (v *SearchUndeclaredLocalVariableVisitor) AcceptListExpression(list []intmod.IExpression) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *SearchUndeclaredLocalVariableVisitor) AcceptListReference(list []intmod.IReference) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *SearchUndeclaredLocalVariableVisitor) AcceptListStatement(list []intmod.IStatement) {
	for _, value := range list {
		value.AcceptStatement(v)
	}
}

func (v *SearchUndeclaredLocalVariableVisitor) SafeAcceptDeclaration(decl intmod.IDeclaration) {
	if decl != nil {
		decl.AcceptDeclaration(v)
	}
}

func (v *SearchUndeclaredLocalVariableVisitor) SafeAcceptExpression(expr intmod.IExpression) {
	if expr != nil {
		expr.Accept(v)
	}
}

func (v *SearchUndeclaredLocalVariableVisitor) SafeAcceptReference(ref intmod.IReference) {
	if ref != nil {
		ref.Accept(v)
	}
}

func (v *SearchUndeclaredLocalVariableVisitor) SafeAcceptStatement(list intmod.IStatement) {
	if list != nil {
		list.AcceptStatement(v)
	}
}

func (v *SearchUndeclaredLocalVariableVisitor) SafeAcceptType(list intmod.IType) {
	if list != nil {
		list.AcceptTypeVisitor(v)
	}
}

func (v *SearchUndeclaredLocalVariableVisitor) SafeAcceptTypeParameter(list intmod.ITypeParameter) {
	if list != nil {
		list.AcceptTypeParameterVisitor(v)
	}
}

func (v *SearchUndeclaredLocalVariableVisitor) SafeAcceptListDeclaration(list []intmod.IDeclaration) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *SearchUndeclaredLocalVariableVisitor) SafeAcceptListConstant(list []intmod.IConstant) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *SearchUndeclaredLocalVariableVisitor) SafeAcceptListStatement(list []intmod.IStatement) {
	if list != nil {
		for _, value := range list {
			value.AcceptStatement(v)
		}
	}
}
