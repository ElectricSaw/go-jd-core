package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/statement"
)

var aggregateFieldsVisitor = NewAggregateFieldsVisitor()
var sortMembersVisitor = NewSortMembersVisitor()
var autoboxingVisitor = NewAutoboxingVisitor()

func NewUpdateJavaSyntaxTreeStep2Visitor(typeMaker intsrv.ITypeMaker) intsrv.IUpdateJavaSyntaxTreeStep2Visitor {
	return &UpdateJavaSyntaxTreeStep2Visitor{
		initStaticFieldVisitor:          NewInitStaticFieldVisitor(),
		initInstanceFieldVisitor:        NewInitInstanceFieldVisitor(),
		initEnumVisitor:                 NewInitEnumVisitor(),
		removeDefaultConstructorVisitor: NewRemoveDefaultConstructorVisitor(),
		replaceBridgeMethodVisitor:      NewUpdateBridgeMethodVisitor(typeMaker),
		initInnerClassStep2Visitor:      NewUpdateNewExpressionVisitor(typeMaker),
		addCastExpressionVisitor:        NewAddCastExpressionVisitor(typeMaker),
	}
}

type UpdateJavaSyntaxTreeStep2Visitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	initStaticFieldVisitor          intsrv.IInitStaticFieldVisitor
	initInstanceFieldVisitor        intsrv.IInitInstanceFieldVisitor
	initEnumVisitor                 intsrv.IInitEnumVisitor
	removeDefaultConstructorVisitor intsrv.IRemoveDefaultConstructorVisitor
	replaceBridgeMethodVisitor      intsrv.IUpdateBridgeMethodVisitor
	initInnerClassStep2Visitor      intsrv.IUpdateNewExpressionVisitor
	addCastExpressionVisitor        intsrv.IAddCastExpressionVisitor
	typeDeclaration                 intmod.ITypeDeclaration
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitBodyDeclaration(declaration intmod.IBodyDeclaration) {
	bodyDeclaration := declaration.(intsrv.IClassFileBodyDeclaration)

	// Visit inner types
	if bodyDeclaration.InnerTypeDeclarations() != nil {
		td := v.typeDeclaration
		tmp := make([]intmod.IDeclaration, 0)
		for _, item := range bodyDeclaration.InnerTypeDeclarations() {
			tmp = append(tmp, item)
		}
		v.AcceptListDeclaration(tmp)
		v.typeDeclaration = td
	}

	// Init bindTypeArgumentVisitor
	v.initStaticFieldVisitor.SetInternalTypeName(v.typeDeclaration.InternalTypeName())

	// Visit declaration
	v.initInnerClassStep2Visitor.VisitBodyDeclaration(declaration)
	v.initStaticFieldVisitor.VisitBodyDeclaration(declaration)
	v.initInstanceFieldVisitor.VisitBodyDeclaration(declaration)
	v.removeDefaultConstructorVisitor.VisitBodyDeclaration(declaration)
	aggregateFieldsVisitor.VisitBodyDeclaration(declaration)
	sortMembersVisitor.VisitBodyDeclaration(declaration)

	if bodyDeclaration.OuterBodyDeclaration() == nil {
		// Main body declaration

		if (bodyDeclaration.InnerTypeDeclarations() != nil) && v.replaceBridgeMethodVisitor.Init(bodyDeclaration) {
			// Replace bridge method invocation
			v.replaceBridgeMethodVisitor.VisitBodyDeclaration(bodyDeclaration)
		}

		// Add cast expressions
		v.addCastExpressionVisitor.VisitBodyDeclaration(declaration)

		// Autoboxing
		autoboxingVisitor.VisitBodyDeclaration(declaration)
	}
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitAnnotationDeclaration(declaration intmod.IAnnotationDeclaration) {
	v.typeDeclaration = declaration
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitClassDeclaration(declaration intmod.IClassDeclaration) {
	v.typeDeclaration = declaration
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitInterfaceDeclaration(declaration intmod.IInterfaceDeclaration) {
	v.typeDeclaration = declaration
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitEnumDeclaration(declaration intmod.IEnumDeclaration) {
	v.typeDeclaration = declaration

	// Remove 'static', 'final' and 'abstract' flags
	cfed := declaration.(intsrv.IClassFileEnumDeclaration)

	cfed.SetFlags(cfed.Flags() & ^(intmod.FlagStatic | intmod.FlagFinal | intmod.FlagAbstract))
	cfed.BodyDeclaration().AcceptDeclaration(v)
	v.initEnumVisitor.VisitBodyDeclaration(cfed.BodyDeclaration())
	cfed.SetConstants(v.initEnumVisitor.Constants().ToSlice())
}

// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$                           $$$
// $$$ AbstractJavaSyntaxVisitor $$$
// $$$                           $$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitCompilationUnit(compilationUnit intmod.ICompilationUnit) {
	compilationUnit.TypeDeclarations().AcceptDeclaration(v)
}

// --- DeclarationVisitor ---

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitArrayVariableInitializer(decl intmod.IArrayVariableInitializer) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitConstructorDeclaration(decl intmod.IConstructorDeclaration) {
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.FormalParameters())
	v.SafeAcceptType(decl.ExceptionTypes())
	v.SafeAcceptStatement(decl.Statements())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitEnumDeclarationConstant(decl intmod.IConstant) {
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptExpression(decl.Arguments())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitExpressionVariableInitializer(decl intmod.IExpressionVariableInitializer) {
	decl.Expression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitFieldDeclaration(decl intmod.IFieldDeclaration) {
	t := decl.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptReference(decl.AnnotationReferences())
	decl.FieldDeclarators().AcceptDeclaration(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitFieldDeclarator(decl intmod.IFieldDeclarator) {
	v.SafeAcceptDeclaration(decl.VariableInitializer())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitFieldDeclarators(decl intmod.IFieldDeclarators) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitFormalParameter(decl intmod.IFormalParameter) {
	t := decl.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitFormalParameters(decl intmod.IFormalParameters) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitInstanceInitializerDeclaration(decl intmod.IInstanceInitializerDeclaration) {
	v.SafeAcceptStatement(decl.Statements())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitLocalVariableDeclaration(decl intmod.ILocalVariableDeclaration) {
	v.SafeAcceptDeclaration(decl.LocalVariableDeclarators())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitLocalVariableDeclarator(decl intmod.ILocalVariableDeclarator) {
	v.SafeAcceptDeclaration(decl.VariableInitializer())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitLocalVariableDeclarators(decl intmod.ILocalVariableDeclarators) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitMethodDeclaration(decl intmod.IMethodDeclaration) {
	t := decl.ReturnedType()
	t.AcceptTypeVisitor(v)

	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.FormalParameters())
	v.SafeAcceptType(decl.ExceptionTypes())
	v.SafeAcceptStatement(decl.Statements())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitMemberDeclarations(decl intmod.IMemberDeclarations) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitModuleDeclaration(_ intmod.IModuleDeclaration) {
	// Empty
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitStaticInitializerDeclaration(_ intmod.IStaticInitializerDeclaration) {
	// Empty
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitTypeDeclarations(decl intmod.ITypeDeclarations) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

// --- IExpressionVisitor ---
func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitArrayExpression(expr intmod.IArrayExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)

	expr.Expression().Accept(v)
	expr.Index().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitBinaryOperatorExpression(expr intmod.IBinaryOperatorExpression) {
	expr.LeftExpression().Accept(v)
	expr.RightExpression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitBooleanExpression(_ intmod.IBooleanExpression) {
	// Empty
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitCastExpression(expr intmod.ICastExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	expr.Expression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitCommentExpression(_ intmod.ICommentExpression) {
	// Empty
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitConstructorInvocationExpression(expression intmod.IConstructorInvocationExpression) {
	t := expression.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expression.Parameters())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitConstructorReferenceExpression(expr intmod.IConstructorReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitDoubleConstantExpression(expr intmod.IDoubleConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitEnumConstantReferenceExpression(expr intmod.IEnumConstantReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitExpressions(expr intmod.IExpressions) {
	list := make([]intmod.IExpression, 0, expr.Size())
	for _, element := range expr.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListExpression(list)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitFieldReferenceExpression(expr intmod.IFieldReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)

	v.SafeAcceptExpression(expr.Expression())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitFloatConstantExpression(expr intmod.IFloatConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitIntegerConstantExpression(expr intmod.IIntegerConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitInstanceOfExpression(expr intmod.IInstanceOfExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	expr.Expression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitLambdaFormalParametersExpression(expr intmod.ILambdaFormalParametersExpression) {
	v.SafeAcceptDeclaration(expr.FormalParameters())
	expr.Statements().AcceptStatement(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitLambdaIdentifiersExpression(expr intmod.ILambdaIdentifiersExpression) {
	v.SafeAcceptStatement(expr.Statements())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitLengthExpression(expr intmod.ILengthExpression) {
	expr.Expression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitLocalVariableReferenceExpression(expr intmod.ILocalVariableReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitLongConstantExpression(expr intmod.ILongConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitMethodInvocationExpression(expr intmod.IMethodInvocationExpression) {
	expr.Expression().Accept(v)
	v.SafeAcceptTypeArgumentVisitable(expr.NonWildcardTypeArguments().(intmod.IWildcardSuperTypeArgument))
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitMethodReferenceExpression(expr intmod.IMethodReferenceExpression) {
	expr.Expression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitNewArray(expr intmod.INewArray) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.DimensionExpressionList())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitNewExpression(expr intmod.INewExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitNewInitializedArray(expr intmod.INewInitializedArray) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptDeclaration(expr.ArrayInitializer())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitNoExpression(_ intmod.INoExpression) {
	// Empty
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitNullExpression(expr intmod.INullExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitObjectTypeReferenceExpression(expr intmod.IObjectTypeReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitParenthesesExpression(expr intmod.IParenthesesExpression) {
	expr.Expression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitPostOperatorExpression(expr intmod.IPostOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitPreOperatorExpression(expr intmod.IPreOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitStringConstantExpression(_ intmod.IStringConstantExpression) {
	// Empty
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitSuperConstructorInvocationExpression(expr intmod.ISuperConstructorInvocationExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitSuperExpression(expr intmod.ISuperExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitTernaryOperatorExpression(expr intmod.ITernaryOperatorExpression) {
	expr.Condition().Accept(v)
	expr.TrueExpression().Accept(v)
	expr.FalseExpression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitThisExpression(expr intmod.IThisExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitTypeReferenceDotClassExpression(expr intmod.ITypeReferenceDotClassExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

// --- IReferenceVisitor ---

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitAnnotationElementValue(ref intmod.IAnnotationElementValue) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitAnnotationReference(ref intmod.IAnnotationReference) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitAnnotationReferences(ref intmod.IAnnotationReferences) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitElementValueArrayInitializerElementValue(ref intmod.IElementValueArrayInitializerElementValue) {
	v.SafeAcceptReference(ref.ElementValueArrayInitializer())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitElementValues(ref intmod.IElementValues) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitElementValuePair(ref intmod.IElementValuePair) {
	ref.ElementValue().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitElementValuePairs(ref intmod.IElementValuePairs) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitExpressionElementValue(ref intmod.IExpressionElementValue) {
	ref.Expression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitInnerObjectReference(ref intmod.IInnerObjectReference) {
	v.VisitInnerObjectType(ref.(intmod.IInnerObjectType))
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitObjectReference(ref intmod.IObjectReference) {
	v.VisitObjectType(ref.(intmod.IObjectType))
}

// --- IStatementVisitor ---

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitAssertStatement(stat intmod.IAssertStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptExpression(stat.Message())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitBreakStatement(_ intmod.IBreakStatement) {
	// Empty
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitByteCodeStatement(_ intmod.IByteCodeStatement) {
	// Empty
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitCommentStatement(_ intmod.ICommentStatement) {
	// Empty
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitContinueStatement(_ intmod.IContinueStatement) {
	// Empty
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitDoWhileStatement(stat intmod.IDoWhileStatement) {
	v.SafeAcceptExpression(stat.Condition())
	v.SafeAcceptStatement(stat.Statements())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitExpressionStatement(stat intmod.IExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitForEachStatement(stat intmod.IForEachStatement) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	stat.Expression().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitForStatement(stat intmod.IForStatement) {
	v.SafeAcceptDeclaration(stat.Declaration())
	v.SafeAcceptExpression(stat.Init())
	v.SafeAcceptExpression(stat.Condition())
	v.SafeAcceptExpression(stat.Update())
	v.SafeAcceptStatement(stat.Statements())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitIfStatement(stat intmod.IIfStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitIfElseStatement(stat intmod.IIfElseStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
	stat.ElseStatements().AcceptStatement(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitLabelStatement(stat intmod.ILabelStatement) {
	v.SafeAcceptStatement(stat.Statements())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitLambdaExpressionStatement(stat intmod.ILambdaExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitLocalVariableDeclarationStatement(stat intmod.ILocalVariableDeclarationStatement) {
	v.VisitLocalVariableDeclaration(&stat.(*statement.LocalVariableDeclarationStatement).LocalVariableDeclaration)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitNoStatement(_ intmod.INoStatement) {
	// Empty
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitReturnExpressionStatement(stat intmod.IReturnExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitReturnStatement(_ intmod.IReturnStatement) {
	// Empty
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitStatements(stat intmod.IStatements) {
	list := make([]intmod.IStatement, 0, stat.Size())
	for _, element := range stat.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListStatement(list)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitSwitchStatement(stat intmod.ISwitchStatement) {
	stat.Condition().Accept(v)
	v.AcceptListStatement(stat.List())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitSwitchStatementDefaultLabel(_ intmod.IDefaultLabel) {
	// Empty
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitSwitchStatementExpressionLabel(stat intmod.IExpressionLabel) {
	stat.Expression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitSwitchStatementLabelBlock(stat intmod.ILabelBlock) {
	stat.Label().AcceptStatement(v)
	stat.Statements().AcceptStatement(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitSwitchStatementMultiLabelsBlock(stat intmod.IMultiLabelsBlock) {
	v.SafeAcceptListStatement(stat.ToSlice())
	stat.Statements().AcceptStatement(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitSynchronizedStatement(stat intmod.ISynchronizedStatement) {
	stat.Monitor().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitThrowStatement(stat intmod.IThrowStatement) {
	stat.Expression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitTryStatement(stat intmod.ITryStatement) {
	v.SafeAcceptListStatement(stat.ResourceList())
	stat.TryStatements().AcceptStatement(v)
	v.SafeAcceptListStatement(stat.CatchClauseList())
	v.SafeAcceptStatement(stat.FinallyStatements())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitTryStatementResource(stat intmod.IResource) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	stat.Expression().Accept(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitTryStatementCatchClause(stat intmod.ICatchClause) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitTypeDeclarationStatement(stat intmod.ITypeDeclarationStatement) {
	stat.TypeDeclaration().AcceptDeclaration(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitWhileStatement(stat intmod.IWhileStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

// --- ITypeVisitor ---

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitTypes(types intmod.ITypes) {
	for _, value := range types.ToSlice() {
		value.AcceptTypeVisitor(v)
	}
}

// --- ITypeParameterVisitor --- //

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitTypeParameter(_ intmod.ITypeParameter) {
	// Empty
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitTypeParameterWithTypeBounds(parameter intmod.ITypeParameterWithTypeBounds) {
	parameter.TypeBounds().AcceptTypeVisitor(v)
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitTypeParameters(parameters intmod.ITypeParameters) {
	for _, param := range parameters.ToSlice() {
		param.AcceptTypeParameterVisitor(v)
	}
}

// --- ITypeArgumentVisitor ---

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitTypeDeclaration(decl intmod.ITypeDeclaration) {
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) AcceptListDeclaration(list []intmod.IDeclaration) {
	for _, value := range list {
		value.AcceptDeclaration(v)
	}
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) AcceptListExpression(list []intmod.IExpression) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) AcceptListReference(list []intmod.IReference) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) AcceptListStatement(list []intmod.IStatement) {
	for _, value := range list {
		value.AcceptStatement(v)
	}
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) SafeAcceptDeclaration(decl intmod.IDeclaration) {
	if decl != nil {
		decl.AcceptDeclaration(v)
	}
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) SafeAcceptExpression(expr intmod.IExpression) {
	if expr != nil {
		expr.Accept(v)
	}
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) SafeAcceptReference(ref intmod.IReference) {
	if ref != nil {
		ref.Accept(v)
	}
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) SafeAcceptStatement(list intmod.IStatement) {
	if list != nil {
		list.AcceptStatement(v)
	}
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) SafeAcceptType(list intmod.IType) {
	if list != nil {
		list.AcceptTypeVisitor(v)
	}
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) SafeAcceptTypeParameter(list intmod.ITypeParameter) {
	if list != nil {
		list.AcceptTypeParameterVisitor(v)
	}
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) SafeAcceptListDeclaration(list []intmod.IDeclaration) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) SafeAcceptListConstant(list []intmod.IConstant) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) SafeAcceptListStatement(list []intmod.IStatement) {
	if list != nil {
		for _, value := range list {
			value.AcceptStatement(v)
		}
	}
}
