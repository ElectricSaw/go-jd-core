package visitor

import (
	"reflect"

	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax"
	modsts "github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/statement"
)

func NewRemoveLastContinueStatementVisitor() intsrv.IRemoveLastContinueStatementVisitor {
	return &RemoveLastContinueStatementVisitor{}
}

type RemoveLastContinueStatementVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor
}

func (v *RemoveLastContinueStatementVisitor) VisitStatements(list intmod.IStatements) {
	if !list.IsEmpty() {
		last := list.Last()

		if reflect.TypeOf(last) == reflect.TypeOf(modsts.ContinueStatement{}) {
			list.RemoveLast()
			v.VisitStatements(list)
		} else {
			last.AcceptStatement(v)
		}
	}
}

func (v *RemoveLastContinueStatementVisitor) VisitIfElseStatement(state intmod.IIfElseStatement) {
	v.SafeAcceptStatement(state.Statements())
	state.ElseStatements().AcceptStatement(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitTryStatement(state intmod.ITryStatement) {
	state.TryStatements().AcceptStatement(v)
	tmp := make([]intmod.IStatement, 0)
	for _, block := range state.CatchClauses() {
		tmp = append(tmp, block)
	}
	v.SafeAcceptListStatement(tmp)
	v.SafeAcceptStatement(state.FinallyStatements())
}

func (v *RemoveLastContinueStatementVisitor) VisitSwitchStatement(state intmod.ISwitchStatement) {
	tmp := make([]intmod.IStatement, 0)
	for _, block := range state.Blocks() {
		tmp = append(tmp, block)
	}
	v.AcceptListStatement(tmp)
}

func (v *RemoveLastContinueStatementVisitor) VisitSwitchStatementLabelBlock(state intmod.ILabelBlock) {
	state.Statements().AcceptStatement(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitSwitchStatementMultiLabelsBlock(state intmod.IMultiLabelsBlock) {
	state.Statements().AcceptStatement(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitIfStatement(state intmod.IIfStatement) {
	v.SafeAcceptStatement(state.Statements())
}

func (v *RemoveLastContinueStatementVisitor) VisitSynchronizedStatement(state intmod.ISynchronizedStatement) {
	v.SafeAcceptStatement(state.Statements())
}

func (v *RemoveLastContinueStatementVisitor) VisitTryStatementCatchClause(state intmod.ICatchClause) {
	v.SafeAcceptStatement(state.Statements())
}

func (v *RemoveLastContinueStatementVisitor) VisitDoWhileStatement(_ intmod.IDoWhileStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitForEachStatement(_ intmod.IForEachStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitForStatement(_ intmod.IForStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitWhileStatement(_ intmod.IWhileStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitAssertStatement(_ intmod.IAssertStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitBreakStatement(_ intmod.IBreakStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitByteCodeStatement(_ intmod.IByteCodeStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitCommentStatement(_ intmod.ICommentStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitContinueStatement(_ intmod.IContinueStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitExpressionStatement(_ intmod.IExpressionStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitLabelStatement(_ intmod.ILabelStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitLambdaExpressionStatement(_ intmod.ILambdaExpressionStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitLocalVariableDeclarationStatement(_ intmod.ILocalVariableDeclarationStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitReturnExpressionStatement(_ intmod.IReturnExpressionStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitReturnStatement(_ intmod.IReturnStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitSwitchStatementExpressionLabel(_ intmod.IExpressionLabel) {
}

func (v *RemoveLastContinueStatementVisitor) VisitThrowStatement(_ intmod.IThrowStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitTypeDeclarationStatement(_ intmod.ITypeDeclarationStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitTryStatementResource(_ intmod.IResource) {
}

// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$                           $$$
// $$$ AbstractJavaSyntaxVisitor $$$
// $$$                           $$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func (v *RemoveLastContinueStatementVisitor) VisitCompilationUnit(compilationUnit intmod.ICompilationUnit) {
	compilationUnit.TypeDeclarations().AcceptDeclaration(v)
}

// --- DeclarationVisitor ---

func (v *RemoveLastContinueStatementVisitor) VisitAnnotationDeclaration(decl intmod.IAnnotationDeclaration) {
	v.SafeAcceptDeclaration(decl.AnnotationDeclarators())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *RemoveLastContinueStatementVisitor) VisitArrayVariableInitializer(decl intmod.IArrayVariableInitializer) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *RemoveLastContinueStatementVisitor) VisitBodyDeclaration(decl intmod.IBodyDeclaration) {
	v.SafeAcceptDeclaration(decl.MemberDeclarations())
}

func (v *RemoveLastContinueStatementVisitor) VisitClassDeclaration(decl intmod.IClassDeclaration) {
	superType := decl.SuperType()

	if superType != nil {
		superType.AcceptTypeVisitor(v)
	}

	v.SafeAcceptTypeParameter(decl.TypeParameters())
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *RemoveLastContinueStatementVisitor) VisitConstructorDeclaration(decl intmod.IConstructorDeclaration) {
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.FormalParameters())
	v.SafeAcceptType(decl.ExceptionTypes())
	v.SafeAcceptStatement(decl.Statements())
}

func (v *RemoveLastContinueStatementVisitor) VisitEnumDeclaration(decl intmod.IEnumDeclaration) {
	v.VisitTypeDeclaration(decl.(intmod.ITypeDeclaration))
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptListConstant(decl.Constants())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *RemoveLastContinueStatementVisitor) VisitEnumDeclarationConstant(decl intmod.IConstant) {
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptExpression(decl.Arguments())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *RemoveLastContinueStatementVisitor) VisitExpressionVariableInitializer(decl intmod.IExpressionVariableInitializer) {
	decl.Expression().Accept(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitFieldDeclaration(decl intmod.IFieldDeclaration) {
	t := decl.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptReference(decl.AnnotationReferences())
	decl.FieldDeclarators().AcceptDeclaration(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitFieldDeclarator(decl intmod.IFieldDeclarator) {
	v.SafeAcceptDeclaration(decl.VariableInitializer())
}

func (v *RemoveLastContinueStatementVisitor) VisitFieldDeclarators(decl intmod.IFieldDeclarators) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *RemoveLastContinueStatementVisitor) VisitFormalParameter(decl intmod.IFormalParameter) {
	t := decl.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *RemoveLastContinueStatementVisitor) VisitFormalParameters(decl intmod.IFormalParameters) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *RemoveLastContinueStatementVisitor) VisitInstanceInitializerDeclaration(decl intmod.IInstanceInitializerDeclaration) {
	v.SafeAcceptStatement(decl.Statements())
}

func (v *RemoveLastContinueStatementVisitor) VisitInterfaceDeclaration(decl intmod.IInterfaceDeclaration) {
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *RemoveLastContinueStatementVisitor) VisitLocalVariableDeclaration(decl intmod.ILocalVariableDeclaration) {
	v.SafeAcceptDeclaration(decl.LocalVariableDeclarators())
}

func (v *RemoveLastContinueStatementVisitor) VisitLocalVariableDeclarator(decl intmod.ILocalVariableDeclarator) {
	v.SafeAcceptDeclaration(decl.VariableInitializer())
}

func (v *RemoveLastContinueStatementVisitor) VisitLocalVariableDeclarators(decl intmod.ILocalVariableDeclarators) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *RemoveLastContinueStatementVisitor) VisitMethodDeclaration(decl intmod.IMethodDeclaration) {
	t := decl.ReturnedType()
	t.AcceptTypeVisitor(v)

	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.FormalParameters())
	v.SafeAcceptType(decl.ExceptionTypes())
	v.SafeAcceptStatement(decl.Statements())
}

func (v *RemoveLastContinueStatementVisitor) VisitMemberDeclarations(decl intmod.IMemberDeclarations) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *RemoveLastContinueStatementVisitor) VisitModuleDeclaration(_ intmod.IModuleDeclaration) {
	// Empty
}

func (v *RemoveLastContinueStatementVisitor) VisitStaticInitializerDeclaration(_ intmod.IStaticInitializerDeclaration) {
	// Empty
}

func (v *RemoveLastContinueStatementVisitor) VisitTypeDeclarations(decl intmod.ITypeDeclarations) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

// --- IExpressionVisitor ---
func (v *RemoveLastContinueStatementVisitor) VisitArrayExpression(expr intmod.IArrayExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)

	expr.Expression().Accept(v)
	expr.Index().Accept(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitBinaryOperatorExpression(expr intmod.IBinaryOperatorExpression) {
	expr.LeftExpression().Accept(v)
	expr.RightExpression().Accept(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitBooleanExpression(_ intmod.IBooleanExpression) {
	// Empty
}

func (v *RemoveLastContinueStatementVisitor) VisitCastExpression(expr intmod.ICastExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	expr.Expression().Accept(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitCommentExpression(_ intmod.ICommentExpression) {
	// Empty
}

func (v *RemoveLastContinueStatementVisitor) VisitConstructorInvocationExpression(expression intmod.IConstructorInvocationExpression) {
	t := expression.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expression.Parameters())
}

func (v *RemoveLastContinueStatementVisitor) VisitConstructorReferenceExpression(expr intmod.IConstructorReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitDoubleConstantExpression(expr intmod.IDoubleConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitEnumConstantReferenceExpression(expr intmod.IEnumConstantReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitExpressions(expr intmod.IExpressions) {
	list := make([]intmod.IExpression, 0, expr.Size())
	for _, element := range expr.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListExpression(list)
}

func (v *RemoveLastContinueStatementVisitor) VisitFieldReferenceExpression(expr intmod.IFieldReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)

	v.SafeAcceptExpression(expr.Expression())
}

func (v *RemoveLastContinueStatementVisitor) VisitFloatConstantExpression(expr intmod.IFloatConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitIntegerConstantExpression(expr intmod.IIntegerConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitInstanceOfExpression(expr intmod.IInstanceOfExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	expr.Expression().Accept(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitLambdaFormalParametersExpression(expr intmod.ILambdaFormalParametersExpression) {
	v.SafeAcceptDeclaration(expr.FormalParameters())
	expr.Statements().AcceptStatement(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitLambdaIdentifiersExpression(expr intmod.ILambdaIdentifiersExpression) {
	v.SafeAcceptStatement(expr.Statements())
}

func (v *RemoveLastContinueStatementVisitor) VisitLengthExpression(expr intmod.ILengthExpression) {
	expr.Expression().Accept(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitLocalVariableReferenceExpression(expr intmod.ILocalVariableReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitLongConstantExpression(expr intmod.ILongConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitMethodInvocationExpression(expr intmod.IMethodInvocationExpression) {
	expr.Expression().Accept(v)
	v.SafeAcceptTypeArgumentVisitable(expr.NonWildcardTypeArguments().(intmod.IWildcardSuperTypeArgument))
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *RemoveLastContinueStatementVisitor) VisitMethodReferenceExpression(expr intmod.IMethodReferenceExpression) {
	expr.Expression().Accept(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitNewArray(expr intmod.INewArray) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.DimensionExpressionList())
}

func (v *RemoveLastContinueStatementVisitor) VisitNewExpression(expr intmod.INewExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *RemoveLastContinueStatementVisitor) VisitNewInitializedArray(expr intmod.INewInitializedArray) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptDeclaration(expr.ArrayInitializer())
}

func (v *RemoveLastContinueStatementVisitor) VisitNoExpression(_ intmod.INoExpression) {
	// Empty
}

func (v *RemoveLastContinueStatementVisitor) VisitNullExpression(expr intmod.INullExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitObjectTypeReferenceExpression(expr intmod.IObjectTypeReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitParenthesesExpression(expr intmod.IParenthesesExpression) {
	expr.Expression().Accept(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitPostOperatorExpression(expr intmod.IPostOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitPreOperatorExpression(expr intmod.IPreOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitStringConstantExpression(_ intmod.IStringConstantExpression) {
	// Empty
}

func (v *RemoveLastContinueStatementVisitor) VisitSuperConstructorInvocationExpression(expr intmod.ISuperConstructorInvocationExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *RemoveLastContinueStatementVisitor) VisitSuperExpression(expr intmod.ISuperExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitTernaryOperatorExpression(expr intmod.ITernaryOperatorExpression) {
	expr.Condition().Accept(v)
	expr.TrueExpression().Accept(v)
	expr.FalseExpression().Accept(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitThisExpression(expr intmod.IThisExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitTypeReferenceDotClassExpression(expr intmod.ITypeReferenceDotClassExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

// --- IReferenceVisitor ---

func (v *RemoveLastContinueStatementVisitor) VisitAnnotationElementValue(ref intmod.IAnnotationElementValue) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *RemoveLastContinueStatementVisitor) VisitAnnotationReference(ref intmod.IAnnotationReference) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *RemoveLastContinueStatementVisitor) VisitAnnotationReferences(ref intmod.IAnnotationReferences) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *RemoveLastContinueStatementVisitor) VisitElementValueArrayInitializerElementValue(ref intmod.IElementValueArrayInitializerElementValue) {
	v.SafeAcceptReference(ref.ElementValueArrayInitializer())
}

func (v *RemoveLastContinueStatementVisitor) VisitElementValues(ref intmod.IElementValues) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *RemoveLastContinueStatementVisitor) VisitElementValuePair(ref intmod.IElementValuePair) {
	ref.ElementValue().Accept(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitElementValuePairs(ref intmod.IElementValuePairs) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *RemoveLastContinueStatementVisitor) VisitExpressionElementValue(ref intmod.IExpressionElementValue) {
	ref.Expression().Accept(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitInnerObjectReference(ref intmod.IInnerObjectReference) {
	v.VisitInnerObjectType(ref.(intmod.IInnerObjectType))
}

func (v *RemoveLastContinueStatementVisitor) VisitObjectReference(ref intmod.IObjectReference) {
	v.VisitObjectType(ref.(intmod.IObjectType))
}

// --- IStatementVisitor ---

func (v *RemoveLastContinueStatementVisitor) VisitNoStatement(_ intmod.INoStatement) {
	// Empty
}

func (v *RemoveLastContinueStatementVisitor) VisitSwitchStatementDefaultLabel(_ intmod.IDefaultLabel) {
	// Empty
}

// --- ITypeVisitor ---

func (v *RemoveLastContinueStatementVisitor) VisitTypes(types intmod.ITypes) {
	for _, value := range types.ToSlice() {
		value.AcceptTypeVisitor(v)
	}
}

// --- ITypeParameterVisitor --- //

func (v *RemoveLastContinueStatementVisitor) VisitTypeParameter(_ intmod.ITypeParameter) {
	// Empty
}

func (v *RemoveLastContinueStatementVisitor) VisitTypeParameterWithTypeBounds(parameter intmod.ITypeParameterWithTypeBounds) {
	parameter.TypeBounds().AcceptTypeVisitor(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitTypeParameters(parameters intmod.ITypeParameters) {
	for _, param := range parameters.ToSlice() {
		param.AcceptTypeParameterVisitor(v)
	}
}

// --- ITypeArgumentVisitor ---

func (v *RemoveLastContinueStatementVisitor) VisitTypeDeclaration(decl intmod.ITypeDeclaration) {
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *RemoveLastContinueStatementVisitor) AcceptListDeclaration(list []intmod.IDeclaration) {
	for _, value := range list {
		value.AcceptDeclaration(v)
	}
}

func (v *RemoveLastContinueStatementVisitor) AcceptListExpression(list []intmod.IExpression) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *RemoveLastContinueStatementVisitor) AcceptListReference(list []intmod.IReference) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *RemoveLastContinueStatementVisitor) AcceptListStatement(list []intmod.IStatement) {
	for _, value := range list {
		value.AcceptStatement(v)
	}
}

func (v *RemoveLastContinueStatementVisitor) SafeAcceptDeclaration(decl intmod.IDeclaration) {
	if decl != nil {
		decl.AcceptDeclaration(v)
	}
}

func (v *RemoveLastContinueStatementVisitor) SafeAcceptExpression(expr intmod.IExpression) {
	if expr != nil {
		expr.Accept(v)
	}
}

func (v *RemoveLastContinueStatementVisitor) SafeAcceptReference(ref intmod.IReference) {
	if ref != nil {
		ref.Accept(v)
	}
}

func (v *RemoveLastContinueStatementVisitor) SafeAcceptStatement(list intmod.IStatement) {
	if list != nil {
		list.AcceptStatement(v)
	}
}

func (v *RemoveLastContinueStatementVisitor) SafeAcceptType(list intmod.IType) {
	if list != nil {
		list.AcceptTypeVisitor(v)
	}
}

func (v *RemoveLastContinueStatementVisitor) SafeAcceptTypeParameter(list intmod.ITypeParameter) {
	if list != nil {
		list.AcceptTypeParameterVisitor(v)
	}
}

func (v *RemoveLastContinueStatementVisitor) SafeAcceptListDeclaration(list []intmod.IDeclaration) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *RemoveLastContinueStatementVisitor) SafeAcceptListConstant(list []intmod.IConstant) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *RemoveLastContinueStatementVisitor) SafeAcceptListStatement(list []intmod.IStatement) {
	if list != nil {
		for _, value := range list {
			value.AcceptStatement(v)
		}
	}
}
