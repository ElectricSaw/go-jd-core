package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax/statement"
)

func NewSearchFirstLineNumberVisitor() intsrv.ISearchFirstLineNumberVisitor {
	return &SearchFirstLineNumberVisitor{
		lineNumber: -1,
	}
}

type SearchFirstLineNumberVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	lineNumber int
}

func (v *SearchFirstLineNumberVisitor) Init() {
	v.lineNumber = -1
}

func (v *SearchFirstLineNumberVisitor) LineNumber() int {
	return v.lineNumber
}

func (v *SearchFirstLineNumberVisitor) VisitStatements(statements intmod.IStatements) {
	if v.lineNumber == -1 {

		for _, value := range statements.ToSlice() {
			value.AcceptStatement(v)
			if v.lineNumber != -1 {
				break
			}
		}
	}
}

func (v *SearchFirstLineNumberVisitor) VisitArrayExpression(expression intmod.IArrayExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitBinaryOperatorExpression(expression intmod.IBinaryOperatorExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitBooleanExpression(expression intmod.IBooleanExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitCastExpression(expression intmod.ICastExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitCommentExpression(expression intmod.ICommentExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitConstructorInvocationExpression(expression intmod.IConstructorInvocationExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitConstructorReferenceExpression(expression intmod.IConstructorReferenceExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitDoubleConstantExpression(expression intmod.IDoubleConstantExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitEnumConstantReferenceExpression(expression intmod.IEnumConstantReferenceExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitFieldReferenceExpression(expression intmod.IFieldReferenceExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitFloatConstantExpression(expression intmod.IFloatConstantExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitIntegerConstantExpression(expression intmod.IIntegerConstantExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitInstanceOfExpression(expression intmod.IInstanceOfExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitLambdaFormalParametersExpression(expression intmod.ILambdaFormalParametersExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitLambdaIdentifiersExpression(expression intmod.ILambdaIdentifiersExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitLengthExpression(expression intmod.ILengthExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitLocalVariableReferenceExpression(expression intmod.ILocalVariableReferenceExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitLongConstantExpression(expression intmod.ILongConstantExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitMethodReferenceExpression(expression intmod.IMethodReferenceExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitNewArray(expression intmod.INewArray) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitNewExpression(expression intmod.INewExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitNewInitializedArray(expression intmod.INewInitializedArray) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitNullExpression(expression intmod.INullExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitObjectTypeReferenceExpression(expression intmod.IObjectTypeReferenceExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitParenthesesExpression(expression intmod.IParenthesesExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitPostOperatorExpression(expression intmod.IPostOperatorExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitPreOperatorExpression(expression intmod.IPreOperatorExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitStringConstantExpression(expression intmod.IStringConstantExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitSuperExpression(expression intmod.ISuperExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitTernaryOperatorExpression(expression intmod.ITernaryOperatorExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitThisExpression(expression intmod.IThisExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitTypeReferenceDotClassExpression(expression intmod.ITypeReferenceDotClassExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitAssertStatement(statement intmod.IAssertStatement) {
	statement.Condition().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitExpressionStatement(statement intmod.IExpressionStatement) {
	statement.Expression().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitForEachStatement(statement intmod.IForEachStatement) {
	statement.Expression().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitIfStatement(statement intmod.IIfStatement) {
	statement.Condition().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitIfElseStatement(statement intmod.IIfElseStatement) {
	statement.Condition().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitLambdaExpressionStatement(statement intmod.ILambdaExpressionStatement) {
	statement.Expression().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitMethodInvocationExpression(expression intmod.IMethodInvocationExpression) {
	expression.Expression().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitSwitchStatement(statement intmod.ISwitchStatement) {
	statement.Condition().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitSynchronizedStatement(statement intmod.ISynchronizedStatement) {
	statement.Monitor().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitThrowStatement(statement intmod.IThrowStatement) {
	statement.Expression().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitWhileStatement(statement intmod.IWhileStatement) {
	if statement.Condition() != nil {
		statement.Condition().Accept(v)
	} else if statement.Statements() != nil {
		statement.Statements().AcceptStatement(v)
	}
}

func (v *SearchFirstLineNumberVisitor) VisitDoWhileStatement(statement intmod.IDoWhileStatement) {
	if statement.Statements() != nil {
		statement.Statements().AcceptStatement(v)
	} else if statement.Condition() != nil {
		statement.Condition().Accept(v)
	}
}

func (v *SearchFirstLineNumberVisitor) VisitForStatement(statement intmod.IForStatement) {
	if statement.Init() != nil {
		statement.Init().Accept(v)
	} else if statement.Condition() != nil {
		statement.Condition().Accept(v)
	} else if statement.Update() != nil {
		statement.Update().Accept(v)
	} else if statement.Statements() != nil {
		statement.Statements().AcceptStatement(v)
	}
}

func (v *SearchFirstLineNumberVisitor) VisitReturnExpressionStatement(statement intmod.IReturnExpressionStatement) {
	statement.Expression().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitTryStatement(statement intmod.ITryStatement) {
	if statement.Resources() != nil {
		v.AcceptListStatement(statement.ResourceList())
	} else {
		statement.TryStatements().AcceptStatement(v)
	}
}

// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$                           $$$
// $$$ AbstractJavaSyntaxVisitor $$$
// $$$                           $$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func (v *SearchFirstLineNumberVisitor) VisitCompilationUnit(compilationUnit intmod.ICompilationUnit) {
	compilationUnit.TypeDeclarations().AcceptDeclaration(v)
}

// --- DeclarationVisitor ---

func (v *SearchFirstLineNumberVisitor) VisitAnnotationDeclaration(decl intmod.IAnnotationDeclaration) {
	v.SafeAcceptDeclaration(decl.AnnotationDeclarators())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *SearchFirstLineNumberVisitor) VisitArrayVariableInitializer(decl intmod.IArrayVariableInitializer) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *SearchFirstLineNumberVisitor) VisitBodyDeclaration(decl intmod.IBodyDeclaration) {
	v.SafeAcceptDeclaration(decl.MemberDeclarations())
}

func (v *SearchFirstLineNumberVisitor) VisitClassDeclaration(decl intmod.IClassDeclaration) {
	superType := decl.SuperType()

	if superType != nil {
		superType.AcceptTypeVisitor(v)
	}

	v.SafeAcceptTypeParameter(decl.TypeParameters())
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *SearchFirstLineNumberVisitor) VisitConstructorDeclaration(decl intmod.IConstructorDeclaration) {
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.FormalParameters())
	v.SafeAcceptType(decl.ExceptionTypes())
	v.SafeAcceptStatement(decl.Statements())
}

func (v *SearchFirstLineNumberVisitor) VisitEnumDeclaration(decl intmod.IEnumDeclaration) {
	v.VisitTypeDeclaration(decl.(intmod.ITypeDeclaration))
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptListConstant(decl.Constants())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *SearchFirstLineNumberVisitor) VisitEnumDeclarationConstant(decl intmod.IConstant) {
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptExpression(decl.Arguments())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *SearchFirstLineNumberVisitor) VisitExpressionVariableInitializer(decl intmod.IExpressionVariableInitializer) {
	decl.Expression().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitFieldDeclaration(decl intmod.IFieldDeclaration) {
	t := decl.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptReference(decl.AnnotationReferences())
	decl.FieldDeclarators().AcceptDeclaration(v)
}

func (v *SearchFirstLineNumberVisitor) VisitFieldDeclarator(decl intmod.IFieldDeclarator) {
	v.SafeAcceptDeclaration(decl.VariableInitializer())
}

func (v *SearchFirstLineNumberVisitor) VisitFieldDeclarators(decl intmod.IFieldDeclarators) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *SearchFirstLineNumberVisitor) VisitFormalParameter(decl intmod.IFormalParameter) {
	t := decl.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *SearchFirstLineNumberVisitor) VisitFormalParameters(decl intmod.IFormalParameters) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *SearchFirstLineNumberVisitor) VisitInstanceInitializerDeclaration(decl intmod.IInstanceInitializerDeclaration) {
	v.SafeAcceptStatement(decl.Statements())
}

func (v *SearchFirstLineNumberVisitor) VisitInterfaceDeclaration(decl intmod.IInterfaceDeclaration) {
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *SearchFirstLineNumberVisitor) VisitLocalVariableDeclaration(decl intmod.ILocalVariableDeclaration) {
	v.SafeAcceptDeclaration(decl.LocalVariableDeclarators())
}

func (v *SearchFirstLineNumberVisitor) VisitLocalVariableDeclarator(decl intmod.ILocalVariableDeclarator) {
	v.SafeAcceptDeclaration(decl.VariableInitializer())
}

func (v *SearchFirstLineNumberVisitor) VisitLocalVariableDeclarators(decl intmod.ILocalVariableDeclarators) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *SearchFirstLineNumberVisitor) VisitMethodDeclaration(decl intmod.IMethodDeclaration) {
	t := decl.ReturnedType()
	t.AcceptTypeVisitor(v)

	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.FormalParameters())
	v.SafeAcceptType(decl.ExceptionTypes())
	v.SafeAcceptStatement(decl.Statements())
}

func (v *SearchFirstLineNumberVisitor) VisitMemberDeclarations(decl intmod.IMemberDeclarations) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *SearchFirstLineNumberVisitor) VisitModuleDeclaration(_ intmod.IModuleDeclaration) {
	// Empty
}

func (v *SearchFirstLineNumberVisitor) VisitStaticInitializerDeclaration(_ intmod.IStaticInitializerDeclaration) {
	// Empty
}

func (v *SearchFirstLineNumberVisitor) VisitTypeDeclarations(decl intmod.ITypeDeclarations) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

// --- IExpressionVisitor ---

func (v *SearchFirstLineNumberVisitor) VisitExpressions(expr intmod.IExpressions) {
	list := make([]intmod.IExpression, 0, expr.Size())
	for _, element := range expr.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListExpression(list)
}

func (v *SearchFirstLineNumberVisitor) VisitNoExpression(_ intmod.INoExpression) {
	// Empty
}

func (v *SearchFirstLineNumberVisitor) VisitSuperConstructorInvocationExpression(expr intmod.ISuperConstructorInvocationExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.Parameters())
}

// --- IReferenceVisitor ---

func (v *SearchFirstLineNumberVisitor) VisitAnnotationElementValue(ref intmod.IAnnotationElementValue) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *SearchFirstLineNumberVisitor) VisitAnnotationReference(ref intmod.IAnnotationReference) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *SearchFirstLineNumberVisitor) VisitAnnotationReferences(ref intmod.IAnnotationReferences) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *SearchFirstLineNumberVisitor) VisitElementValueArrayInitializerElementValue(ref intmod.IElementValueArrayInitializerElementValue) {
	v.SafeAcceptReference(ref.ElementValueArrayInitializer())
}

func (v *SearchFirstLineNumberVisitor) VisitElementValues(ref intmod.IElementValues) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *SearchFirstLineNumberVisitor) VisitElementValuePair(ref intmod.IElementValuePair) {
	ref.ElementValue().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitElementValuePairs(ref intmod.IElementValuePairs) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *SearchFirstLineNumberVisitor) VisitExpressionElementValue(ref intmod.IExpressionElementValue) {
	ref.Expression().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitInnerObjectReference(ref intmod.IInnerObjectReference) {
	v.VisitInnerObjectType(ref.(intmod.IInnerObjectType))
}

func (v *SearchFirstLineNumberVisitor) VisitObjectReference(ref intmod.IObjectReference) {
	v.VisitObjectType(ref.(intmod.IObjectType))
}

// --- IStatementVisitor ---

func (v *SearchFirstLineNumberVisitor) VisitBreakStatement(_ intmod.IBreakStatement) {
	// Empty
}

func (v *SearchFirstLineNumberVisitor) VisitByteCodeStatement(_ intmod.IByteCodeStatement) {
	// Empty
}

func (v *SearchFirstLineNumberVisitor) VisitCommentStatement(_ intmod.ICommentStatement) {
	// Empty
}

func (v *SearchFirstLineNumberVisitor) VisitContinueStatement(_ intmod.IContinueStatement) {
	// Empty
}

func (v *SearchFirstLineNumberVisitor) VisitLabelStatement(stat intmod.ILabelStatement) {
	v.SafeAcceptStatement(stat.Statements())
}

func (v *SearchFirstLineNumberVisitor) VisitLocalVariableDeclarationStatement(stat intmod.ILocalVariableDeclarationStatement) {
	v.VisitLocalVariableDeclaration(&stat.(*statement.LocalVariableDeclarationStatement).LocalVariableDeclaration)
}

func (v *SearchFirstLineNumberVisitor) VisitNoStatement(_ intmod.INoStatement) {
	// Empty
}

func (v *SearchFirstLineNumberVisitor) VisitReturnStatement(_ intmod.IReturnStatement) {
	// Empty
}

func (v *SearchFirstLineNumberVisitor) VisitSwitchStatementDefaultLabel(_ intmod.IDefaultLabel) {
	// Empty
}

func (v *SearchFirstLineNumberVisitor) VisitSwitchStatementExpressionLabel(stat intmod.IExpressionLabel) {
	stat.Expression().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitSwitchStatementLabelBlock(stat intmod.ILabelBlock) {
	stat.Label().AcceptStatement(v)
	stat.Statements().AcceptStatement(v)
}

func (v *SearchFirstLineNumberVisitor) VisitSwitchStatementMultiLabelsBlock(stat intmod.IMultiLabelsBlock) {
	v.SafeAcceptListStatement(stat.ToSlice())
	stat.Statements().AcceptStatement(v)
}

func (v *SearchFirstLineNumberVisitor) VisitTryStatementResource(stat intmod.IResource) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	stat.Expression().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitTryStatementCatchClause(stat intmod.ICatchClause) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *SearchFirstLineNumberVisitor) VisitTypeDeclarationStatement(stat intmod.ITypeDeclarationStatement) {
	stat.TypeDeclaration().AcceptDeclaration(v)
}

// --- ITypeVisitor ---

func (v *SearchFirstLineNumberVisitor) VisitTypes(types intmod.ITypes) {
	for _, value := range types.ToSlice() {
		value.AcceptTypeVisitor(v)
	}
}

// --- ITypeParameterVisitor --- //

func (v *SearchFirstLineNumberVisitor) VisitTypeParameter(_ intmod.ITypeParameter) {
	// Empty
}

func (v *SearchFirstLineNumberVisitor) VisitTypeParameterWithTypeBounds(parameter intmod.ITypeParameterWithTypeBounds) {
	parameter.TypeBounds().AcceptTypeVisitor(v)
}

func (v *SearchFirstLineNumberVisitor) VisitTypeParameters(parameters intmod.ITypeParameters) {
	for _, param := range parameters.ToSlice() {
		param.AcceptTypeParameterVisitor(v)
	}
}

// --- ITypeArgumentVisitor ---

func (v *SearchFirstLineNumberVisitor) VisitTypeDeclaration(decl intmod.ITypeDeclaration) {
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *SearchFirstLineNumberVisitor) AcceptListDeclaration(list []intmod.IDeclaration) {
	for _, value := range list {
		value.AcceptDeclaration(v)
	}
}

func (v *SearchFirstLineNumberVisitor) AcceptListExpression(list []intmod.IExpression) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *SearchFirstLineNumberVisitor) AcceptListReference(list []intmod.IReference) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *SearchFirstLineNumberVisitor) AcceptListStatement(list []intmod.IStatement) {
	for _, value := range list {
		value.AcceptStatement(v)
	}
}

func (v *SearchFirstLineNumberVisitor) SafeAcceptDeclaration(decl intmod.IDeclaration) {
	if decl != nil {
		decl.AcceptDeclaration(v)
	}
}

func (v *SearchFirstLineNumberVisitor) SafeAcceptExpression(expr intmod.IExpression) {
	if expr != nil {
		expr.Accept(v)
	}
}

func (v *SearchFirstLineNumberVisitor) SafeAcceptReference(ref intmod.IReference) {
	if ref != nil {
		ref.Accept(v)
	}
}

func (v *SearchFirstLineNumberVisitor) SafeAcceptStatement(list intmod.IStatement) {
	if list != nil {
		list.AcceptStatement(v)
	}
}

func (v *SearchFirstLineNumberVisitor) SafeAcceptType(list intmod.IType) {
	if list != nil {
		list.AcceptTypeVisitor(v)
	}
}

func (v *SearchFirstLineNumberVisitor) SafeAcceptTypeParameter(list intmod.ITypeParameter) {
	if list != nil {
		list.AcceptTypeParameterVisitor(v)
	}
}

func (v *SearchFirstLineNumberVisitor) SafeAcceptListDeclaration(list []intmod.IDeclaration) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *SearchFirstLineNumberVisitor) SafeAcceptListConstant(list []intmod.IConstant) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *SearchFirstLineNumberVisitor) SafeAcceptListStatement(list []intmod.IStatement) {
	if list != nil {
		for _, value := range list {
			value.AcceptStatement(v)
		}
	}
}
