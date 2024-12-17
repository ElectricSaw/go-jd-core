package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax/statement"
)

func NewSearchLocalVariableReferenceVisitor() intsrv.ISearchLocalVariableReferenceVisitor {
	return &SearchLocalVariableReferenceVisitor{}
}

type SearchLocalVariableReferenceVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	index int
	found bool
}

func (v *SearchLocalVariableReferenceVisitor) Init(index int) {
	v.index = index
	v.found = false
}

func (v *SearchLocalVariableReferenceVisitor) ContainsReference() bool {
	return v.found
}

func (v *SearchLocalVariableReferenceVisitor) VisitLocalVariableReferenceExpression(expr intmod.ILocalVariableReferenceExpression) {
	if v.index < 0 {
		v.found = true
	} else {
		refExpr := expr.(intsrv.IClassFileLocalVariableReferenceExpression)
		v.found = v.found || refExpr.LocalVariable().(intsrv.ILocalVariable).Index() == v.index
	}
}

// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$                           $$$
// $$$ AbstractJavaSyntaxVisitor $$$
// $$$                           $$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func (v *SearchLocalVariableReferenceVisitor) VisitCompilationUnit(compilationUnit intmod.ICompilationUnit) {
	compilationUnit.TypeDeclarations().AcceptDeclaration(v)
}

// --- DeclarationVisitor ---

func (v *SearchLocalVariableReferenceVisitor) VisitAnnotationDeclaration(decl intmod.IAnnotationDeclaration) {
	v.SafeAcceptDeclaration(decl.AnnotationDeclarators())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *SearchLocalVariableReferenceVisitor) VisitArrayVariableInitializer(decl intmod.IArrayVariableInitializer) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *SearchLocalVariableReferenceVisitor) VisitBodyDeclaration(decl intmod.IBodyDeclaration) {
	v.SafeAcceptDeclaration(decl.MemberDeclarations())
}

func (v *SearchLocalVariableReferenceVisitor) VisitClassDeclaration(decl intmod.IClassDeclaration) {
	superType := decl.SuperType()

	if superType != nil {
		superType.AcceptTypeVisitor(v)
	}

	v.SafeAcceptTypeParameter(decl.TypeParameters())
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *SearchLocalVariableReferenceVisitor) VisitConstructorDeclaration(decl intmod.IConstructorDeclaration) {
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.FormalParameters())
	v.SafeAcceptType(decl.ExceptionTypes())
	v.SafeAcceptStatement(decl.Statements())
}

func (v *SearchLocalVariableReferenceVisitor) VisitEnumDeclaration(decl intmod.IEnumDeclaration) {
	v.VisitTypeDeclaration(decl.(intmod.ITypeDeclaration))
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptListConstant(decl.Constants())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *SearchLocalVariableReferenceVisitor) VisitEnumDeclarationConstant(decl intmod.IConstant) {
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptExpression(decl.Arguments())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *SearchLocalVariableReferenceVisitor) VisitExpressionVariableInitializer(decl intmod.IExpressionVariableInitializer) {
	decl.Expression().Accept(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitFieldDeclaration(decl intmod.IFieldDeclaration) {
	t := decl.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptReference(decl.AnnotationReferences())
	decl.FieldDeclarators().AcceptDeclaration(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitFieldDeclarator(decl intmod.IFieldDeclarator) {
	v.SafeAcceptDeclaration(decl.VariableInitializer())
}

func (v *SearchLocalVariableReferenceVisitor) VisitFieldDeclarators(decl intmod.IFieldDeclarators) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *SearchLocalVariableReferenceVisitor) VisitFormalParameter(decl intmod.IFormalParameter) {
	t := decl.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *SearchLocalVariableReferenceVisitor) VisitFormalParameters(decl intmod.IFormalParameters) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *SearchLocalVariableReferenceVisitor) VisitInstanceInitializerDeclaration(decl intmod.IInstanceInitializerDeclaration) {
	v.SafeAcceptStatement(decl.Statements())
}

func (v *SearchLocalVariableReferenceVisitor) VisitInterfaceDeclaration(decl intmod.IInterfaceDeclaration) {
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *SearchLocalVariableReferenceVisitor) VisitLocalVariableDeclaration(decl intmod.ILocalVariableDeclaration) {
	v.SafeAcceptDeclaration(decl.LocalVariableDeclarators())
}

func (v *SearchLocalVariableReferenceVisitor) VisitLocalVariableDeclarator(decl intmod.ILocalVariableDeclarator) {
	v.SafeAcceptDeclaration(decl.VariableInitializer())
}

func (v *SearchLocalVariableReferenceVisitor) VisitLocalVariableDeclarators(decl intmod.ILocalVariableDeclarators) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *SearchLocalVariableReferenceVisitor) VisitMethodDeclaration(decl intmod.IMethodDeclaration) {
	t := decl.ReturnedType()
	t.AcceptTypeVisitor(v)

	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.FormalParameters())
	v.SafeAcceptType(decl.ExceptionTypes())
	v.SafeAcceptStatement(decl.Statements())
}

func (v *SearchLocalVariableReferenceVisitor) VisitMemberDeclarations(decl intmod.IMemberDeclarations) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *SearchLocalVariableReferenceVisitor) VisitModuleDeclaration(_ intmod.IModuleDeclaration) {
	// Empty
}

func (v *SearchLocalVariableReferenceVisitor) VisitStaticInitializerDeclaration(_ intmod.IStaticInitializerDeclaration) {
	// Empty
}

func (v *SearchLocalVariableReferenceVisitor) VisitTypeDeclarations(decl intmod.ITypeDeclarations) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

// --- IExpressionVisitor ---
func (v *SearchLocalVariableReferenceVisitor) VisitArrayExpression(expr intmod.IArrayExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)

	expr.Expression().Accept(v)
	expr.Index().Accept(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitBinaryOperatorExpression(expr intmod.IBinaryOperatorExpression) {
	expr.LeftExpression().Accept(v)
	expr.RightExpression().Accept(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitBooleanExpression(_ intmod.IBooleanExpression) {
	// Empty
}

func (v *SearchLocalVariableReferenceVisitor) VisitCastExpression(expr intmod.ICastExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	expr.Expression().Accept(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitCommentExpression(_ intmod.ICommentExpression) {
	// Empty
}

func (v *SearchLocalVariableReferenceVisitor) VisitConstructorInvocationExpression(expression intmod.IConstructorInvocationExpression) {
	t := expression.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expression.Parameters())
}

func (v *SearchLocalVariableReferenceVisitor) VisitConstructorReferenceExpression(expr intmod.IConstructorReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitDoubleConstantExpression(expr intmod.IDoubleConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitEnumConstantReferenceExpression(expr intmod.IEnumConstantReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitExpressions(expr intmod.IExpressions) {
	list := make([]intmod.IExpression, 0, expr.Size())
	for _, element := range expr.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListExpression(list)
}

func (v *SearchLocalVariableReferenceVisitor) VisitFieldReferenceExpression(expr intmod.IFieldReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)

	v.SafeAcceptExpression(expr.Expression())
}

func (v *SearchLocalVariableReferenceVisitor) VisitFloatConstantExpression(expr intmod.IFloatConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitIntegerConstantExpression(expr intmod.IIntegerConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitInstanceOfExpression(expr intmod.IInstanceOfExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	expr.Expression().Accept(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitLambdaFormalParametersExpression(expr intmod.ILambdaFormalParametersExpression) {
	v.SafeAcceptDeclaration(expr.FormalParameters())
	expr.Statements().AcceptStatement(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitLambdaIdentifiersExpression(expr intmod.ILambdaIdentifiersExpression) {
	v.SafeAcceptStatement(expr.Statements())
}

func (v *SearchLocalVariableReferenceVisitor) VisitLengthExpression(expr intmod.ILengthExpression) {
	expr.Expression().Accept(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitLongConstantExpression(expr intmod.ILongConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitMethodInvocationExpression(expr intmod.IMethodInvocationExpression) {
	expr.Expression().Accept(v)
	v.SafeAcceptTypeArgumentVisitable(expr.NonWildcardTypeArguments().(intmod.IWildcardSuperTypeArgument))
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *SearchLocalVariableReferenceVisitor) VisitMethodReferenceExpression(expr intmod.IMethodReferenceExpression) {
	expr.Expression().Accept(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitNewArray(expr intmod.INewArray) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.DimensionExpressionList())
}

func (v *SearchLocalVariableReferenceVisitor) VisitNewExpression(expr intmod.INewExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *SearchLocalVariableReferenceVisitor) VisitNewInitializedArray(expr intmod.INewInitializedArray) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptDeclaration(expr.ArrayInitializer())
}

func (v *SearchLocalVariableReferenceVisitor) VisitNoExpression(_ intmod.INoExpression) {
	// Empty
}

func (v *SearchLocalVariableReferenceVisitor) VisitNullExpression(expr intmod.INullExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitObjectTypeReferenceExpression(expr intmod.IObjectTypeReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitParenthesesExpression(expr intmod.IParenthesesExpression) {
	expr.Expression().Accept(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitPostOperatorExpression(expr intmod.IPostOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitPreOperatorExpression(expr intmod.IPreOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitStringConstantExpression(_ intmod.IStringConstantExpression) {
	// Empty
}

func (v *SearchLocalVariableReferenceVisitor) VisitSuperConstructorInvocationExpression(expr intmod.ISuperConstructorInvocationExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *SearchLocalVariableReferenceVisitor) VisitSuperExpression(expr intmod.ISuperExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitTernaryOperatorExpression(expr intmod.ITernaryOperatorExpression) {
	expr.Condition().Accept(v)
	expr.TrueExpression().Accept(v)
	expr.FalseExpression().Accept(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitThisExpression(expr intmod.IThisExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitTypeReferenceDotClassExpression(expr intmod.ITypeReferenceDotClassExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

// --- IReferenceVisitor ---

func (v *SearchLocalVariableReferenceVisitor) VisitAnnotationElementValue(ref intmod.IAnnotationElementValue) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *SearchLocalVariableReferenceVisitor) VisitAnnotationReference(ref intmod.IAnnotationReference) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *SearchLocalVariableReferenceVisitor) VisitAnnotationReferences(ref intmod.IAnnotationReferences) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *SearchLocalVariableReferenceVisitor) VisitElementValueArrayInitializerElementValue(ref intmod.IElementValueArrayInitializerElementValue) {
	v.SafeAcceptReference(ref.ElementValueArrayInitializer())
}

func (v *SearchLocalVariableReferenceVisitor) VisitElementValues(ref intmod.IElementValues) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *SearchLocalVariableReferenceVisitor) VisitElementValuePair(ref intmod.IElementValuePair) {
	ref.ElementValue().Accept(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitElementValuePairs(ref intmod.IElementValuePairs) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *SearchLocalVariableReferenceVisitor) VisitExpressionElementValue(ref intmod.IExpressionElementValue) {
	ref.Expression().Accept(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitInnerObjectReference(ref intmod.IInnerObjectReference) {
	v.VisitInnerObjectType(ref.(intmod.IInnerObjectType))
}

func (v *SearchLocalVariableReferenceVisitor) VisitObjectReference(ref intmod.IObjectReference) {
	v.VisitObjectType(ref.(intmod.IObjectType))
}

// --- IStatementVisitor ---

func (v *SearchLocalVariableReferenceVisitor) VisitAssertStatement(stat intmod.IAssertStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptExpression(stat.Message())
}

func (v *SearchLocalVariableReferenceVisitor) VisitBreakStatement(_ intmod.IBreakStatement) {
	// Empty
}

func (v *SearchLocalVariableReferenceVisitor) VisitByteCodeStatement(_ intmod.IByteCodeStatement) {
	// Empty
}

func (v *SearchLocalVariableReferenceVisitor) VisitCommentStatement(_ intmod.ICommentStatement) {
	// Empty
}

func (v *SearchLocalVariableReferenceVisitor) VisitContinueStatement(_ intmod.IContinueStatement) {
	// Empty
}

func (v *SearchLocalVariableReferenceVisitor) VisitDoWhileStatement(stat intmod.IDoWhileStatement) {
	v.SafeAcceptExpression(stat.Condition())
	v.SafeAcceptStatement(stat.Statements())
}

func (v *SearchLocalVariableReferenceVisitor) VisitExpressionStatement(stat intmod.IExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitForEachStatement(stat intmod.IForEachStatement) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	stat.Expression().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *SearchLocalVariableReferenceVisitor) VisitForStatement(stat intmod.IForStatement) {
	v.SafeAcceptDeclaration(stat.Declaration())
	v.SafeAcceptExpression(stat.Init())
	v.SafeAcceptExpression(stat.Condition())
	v.SafeAcceptExpression(stat.Update())
	v.SafeAcceptStatement(stat.Statements())
}

func (v *SearchLocalVariableReferenceVisitor) VisitIfStatement(stat intmod.IIfStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *SearchLocalVariableReferenceVisitor) VisitIfElseStatement(stat intmod.IIfElseStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
	stat.ElseStatements().AcceptStatement(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitLabelStatement(stat intmod.ILabelStatement) {
	v.SafeAcceptStatement(stat.Statements())
}

func (v *SearchLocalVariableReferenceVisitor) VisitLambdaExpressionStatement(stat intmod.ILambdaExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitLocalVariableDeclarationStatement(stat intmod.ILocalVariableDeclarationStatement) {
	v.VisitLocalVariableDeclaration(&stat.(*statement.LocalVariableDeclarationStatement).LocalVariableDeclaration)
}

func (v *SearchLocalVariableReferenceVisitor) VisitNoStatement(_ intmod.INoStatement) {
	// Empty
}

func (v *SearchLocalVariableReferenceVisitor) VisitReturnExpressionStatement(stat intmod.IReturnExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitReturnStatement(_ intmod.IReturnStatement) {
	// Empty
}

func (v *SearchLocalVariableReferenceVisitor) VisitStatements(stat intmod.IStatements) {
	list := make([]intmod.IStatement, 0, stat.Size())
	for _, element := range stat.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListStatement(list)
}

func (v *SearchLocalVariableReferenceVisitor) VisitSwitchStatement(stat intmod.ISwitchStatement) {
	stat.Condition().Accept(v)
	v.AcceptListStatement(stat.List())
}

func (v *SearchLocalVariableReferenceVisitor) VisitSwitchStatementDefaultLabel(_ intmod.IDefaultLabel) {
	// Empty
}

func (v *SearchLocalVariableReferenceVisitor) VisitSwitchStatementExpressionLabel(stat intmod.IExpressionLabel) {
	stat.Expression().Accept(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitSwitchStatementLabelBlock(stat intmod.ILabelBlock) {
	stat.Label().AcceptStatement(v)
	stat.Statements().AcceptStatement(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitSwitchStatementMultiLabelsBlock(stat intmod.IMultiLabelsBlock) {
	v.SafeAcceptListStatement(stat.ToSlice())
	stat.Statements().AcceptStatement(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitSynchronizedStatement(stat intmod.ISynchronizedStatement) {
	stat.Monitor().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *SearchLocalVariableReferenceVisitor) VisitThrowStatement(stat intmod.IThrowStatement) {
	stat.Expression().Accept(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitTryStatement(stat intmod.ITryStatement) {
	v.SafeAcceptListStatement(stat.ResourceList())
	stat.TryStatements().AcceptStatement(v)
	v.SafeAcceptListStatement(stat.CatchClauseList())
	v.SafeAcceptStatement(stat.FinallyStatements())
}

func (v *SearchLocalVariableReferenceVisitor) VisitTryStatementResource(stat intmod.IResource) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	stat.Expression().Accept(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitTryStatementCatchClause(stat intmod.ICatchClause) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *SearchLocalVariableReferenceVisitor) VisitTypeDeclarationStatement(stat intmod.ITypeDeclarationStatement) {
	stat.TypeDeclaration().AcceptDeclaration(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitWhileStatement(stat intmod.IWhileStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

// --- ITypeVisitor ---

func (v *SearchLocalVariableReferenceVisitor) VisitTypes(types intmod.ITypes) {
	for _, value := range types.ToSlice() {
		value.AcceptTypeVisitor(v)
	}
}

// --- ITypeParameterVisitor --- //

func (v *SearchLocalVariableReferenceVisitor) VisitTypeParameter(_ intmod.ITypeParameter) {
	// Empty
}

func (v *SearchLocalVariableReferenceVisitor) VisitTypeParameterWithTypeBounds(parameter intmod.ITypeParameterWithTypeBounds) {
	parameter.TypeBounds().AcceptTypeVisitor(v)
}

func (v *SearchLocalVariableReferenceVisitor) VisitTypeParameters(parameters intmod.ITypeParameters) {
	for _, param := range parameters.ToSlice() {
		param.AcceptTypeParameterVisitor(v)
	}
}

// --- ITypeArgumentVisitor ---

func (v *SearchLocalVariableReferenceVisitor) VisitTypeDeclaration(decl intmod.ITypeDeclaration) {
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *SearchLocalVariableReferenceVisitor) AcceptListDeclaration(list []intmod.IDeclaration) {
	for _, value := range list {
		value.AcceptDeclaration(v)
	}
}

func (v *SearchLocalVariableReferenceVisitor) AcceptListExpression(list []intmod.IExpression) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *SearchLocalVariableReferenceVisitor) AcceptListReference(list []intmod.IReference) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *SearchLocalVariableReferenceVisitor) AcceptListStatement(list []intmod.IStatement) {
	for _, value := range list {
		value.AcceptStatement(v)
	}
}

func (v *SearchLocalVariableReferenceVisitor) SafeAcceptDeclaration(decl intmod.IDeclaration) {
	if decl != nil {
		decl.AcceptDeclaration(v)
	}
}

func (v *SearchLocalVariableReferenceVisitor) SafeAcceptExpression(expr intmod.IExpression) {
	if expr != nil {
		expr.Accept(v)
	}
}

func (v *SearchLocalVariableReferenceVisitor) SafeAcceptReference(ref intmod.IReference) {
	if ref != nil {
		ref.Accept(v)
	}
}

func (v *SearchLocalVariableReferenceVisitor) SafeAcceptStatement(list intmod.IStatement) {
	if list != nil {
		list.AcceptStatement(v)
	}
}

func (v *SearchLocalVariableReferenceVisitor) SafeAcceptType(list intmod.IType) {
	if list != nil {
		list.AcceptTypeVisitor(v)
	}
}

func (v *SearchLocalVariableReferenceVisitor) SafeAcceptTypeParameter(list intmod.ITypeParameter) {
	if list != nil {
		list.AcceptTypeParameterVisitor(v)
	}
}

func (v *SearchLocalVariableReferenceVisitor) SafeAcceptListDeclaration(list []intmod.IDeclaration) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *SearchLocalVariableReferenceVisitor) SafeAcceptListConstant(list []intmod.IConstant) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *SearchLocalVariableReferenceVisitor) SafeAcceptListStatement(list []intmod.IStatement) {
	if list != nil {
		for _, value := range list {
			value.AcceptStatement(v)
		}
	}
}
