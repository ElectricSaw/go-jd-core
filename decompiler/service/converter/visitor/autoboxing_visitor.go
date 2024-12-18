package visitor

import (
	"strings"

	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
)

func NewAutoboxingVisitor() intmod.IJavaSyntaxVisitor {
	return &AutoboxingVisitor{}
}

var ValueOfDescriptorMap = map[string]string{
	"java/lang/Byte":      "(B)Ljava/lang/Byte;",
	"java/lang/Character": "(C)Ljava/lang/Character;",
	"java/lang/Float":     "(F)Ljava/lang/Float;",
	"java/lang/Integer":   "(I)Ljava/lang/Integer;",
	"java/lang/Long":      "(J)Ljava/lang/Long;",
	"java/lang/Short":     "(S)Ljava/lang/Short;",
	"java/lang/Double":    "(D)Ljava/lang/Double;",
	"java/lang/Boolean":   "(Z)Ljava/lang/Boolean;",
}

var ValueDescriptorMap = map[string]string{
	"java/lang/Byte":      "()B",
	"java/lang/Character": "()C",
	"java/lang/Float":     "()F",
	"java/lang/Integer":   "()I",
	"java/lang/Long":      "()J",
	"java/lang/Short":     "()S",
	"java/lang/Double":    "()D",
	"java/lang/Boolean":   "()Z",
}

var ValueMethodNameMap = map[string]string{
	"java/lang/Byte":      "byteValue",
	"java/lang/Character": "charValue",
	"java/lang/Float":     "floatValue",
	"java/lang/Integer":   "intValue",
	"java/lang/Long":      "longValue",
	"java/lang/Short":     "shortValue",
	"java/lang/Double":    "doubleValue",
	"java/lang/Boolean":   "booleanValue",
}

type AutoboxingVisitor struct {
	AbstractUpdateExpressionVisitor
}

func (v *AutoboxingVisitor) VisitBodyDeclaration(declaration intmod.IBodyDeclaration) {
	cfbd := declaration.(intsrv.IClassFileBodyDeclaration)
	autoBoxingSupported := cfbd.ClassFile().MajorVersion() >= 49 // (majorVersion >= Java 5)

	if autoBoxingSupported {
		v.SafeAcceptDeclaration(declaration.MemberDeclarations())
	}
}

func (v *AutoboxingVisitor) updateExpression(expression intmod.IExpression) intmod.IExpression {
	if expression.IsMethodInvocationExpression() && strings.HasPrefix(expression.InternalTypeName(), "java/lang/") {
		var parameterSize int

		if expression.Parameters() != nil {
			parameterSize = expression.Parameters().Size()
		}

		if expression.Expression().IsObjectTypeReferenceExpression() {
			// static method invocation
			if parameterSize == 1 &&
				expression.Name() == "valueOf" &&
				expression.Descriptor() == ValueOfDescriptorMap[expression.InternalTypeName()] {
				return expression.Parameters().First()
			}
		} else {
			// non-static method invocation
			if (parameterSize == 0) &&
				expression.Name() == ValueMethodNameMap[expression.InternalTypeName()] &&
				expression.Descriptor() == ValueDescriptorMap[expression.InternalTypeName()] {
				return expression.Expression()
			}
		}
	}

	return expression
}

// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$                                 $$$
// $$$ AbstractUpdateExpressionVisitor $$$
// $$$                                 $$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func (v *AutoboxingVisitor) UpdateExpression(_ intmod.IExpression) intmod.IExpression {
	return nil
}

func (v *AutoboxingVisitor) UpdateBaseExpression(baseExpression intmod.IExpression) intmod.IExpression {
	if baseExpression == nil {
		return nil
	}

	if baseExpression.IsList() {
		iterator := baseExpression.ToList().ListIterator()

		for iterator.HasNext() {
			value := iterator.Next()
			_ = iterator.Set(v.UpdateExpression(value))
		}

		return baseExpression
	}

	return v.UpdateExpression(baseExpression.First())
}

func (v *AutoboxingVisitor) VisitAnnotationDeclaration(decl intmod.IAnnotationDeclaration) {
	v.SafeAcceptDeclaration(decl.AnnotationDeclarators())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *AutoboxingVisitor) VisitClassDeclaration(declaration intmod.IClassDeclaration) {
	v.VisitInterfaceDeclaration(declaration)
}

func (v *AutoboxingVisitor) VisitConstructorInvocationExpression(expression intmod.IConstructorInvocationExpression) {
	if expression.Parameters() != nil {
		expression.SetParameters(v.UpdateBaseExpression(expression.Parameters()))
		expression.Parameters().Accept(v)
	}
}

func (v *AutoboxingVisitor) VisitConstructorDeclaration(declaration intmod.IConstructorDeclaration) {
	v.SafeAcceptStatement(declaration.Statements())
}

func (v *AutoboxingVisitor) VisitEnumDeclaration(declaration intmod.IEnumDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *AutoboxingVisitor) VisitEnumDeclarationConstant(declaration intmod.IConstant) {
	if declaration.Arguments() != nil {
		declaration.SetArguments(v.UpdateBaseExpression(declaration.Arguments()))
		declaration.Arguments().Accept(v)
	}
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *AutoboxingVisitor) VisitExpressionVariableInitializer(declaration intmod.IExpressionVariableInitializer) {
	if declaration.Expression() != nil {
		declaration.SetExpression(v.UpdateExpression(declaration.Expression()))
		declaration.Expression().Accept(v)
	}
}

func (v *AutoboxingVisitor) VisitFieldDeclaration(declaration intmod.IFieldDeclaration) {
	declaration.FieldDeclarators().AcceptDeclaration(v)
}

func (v *AutoboxingVisitor) VisitFieldDeclarator(declaration intmod.IFieldDeclarator) {
	v.SafeAcceptDeclaration(declaration.VariableInitializer())
}

func (v *AutoboxingVisitor) VisitFormalParameter(_ intmod.IFormalParameter) {
}

func (v *AutoboxingVisitor) VisitInterfaceDeclaration(declaration intmod.IInterfaceDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *AutoboxingVisitor) VisitLocalVariableDeclaration(declaration intmod.ILocalVariableDeclaration) {
	declaration.LocalVariableDeclarators().AcceptDeclaration(v)
}

func (v *AutoboxingVisitor) VisitLocalVariableDeclarator(declarator intmod.ILocalVariableDeclarator) {
	v.SafeAcceptDeclaration(declarator.VariableInitializer())
}

func (v *AutoboxingVisitor) VisitMethodDeclaration(declaration intmod.IMethodDeclaration) {
	v.SafeAcceptReference(declaration.AnnotationReferences())
	v.SafeAcceptStatement(declaration.Statements())
}

func (v *AutoboxingVisitor) VisitArrayExpression(expression intmod.IArrayExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.SetIndex(v.UpdateExpression(expression.Index()))
	expression.Expression().Accept(v)
	expression.Index().Accept(v)
}

func (v *AutoboxingVisitor) VisitBinaryOperatorExpression(expression intmod.IBinaryOperatorExpression) {
	expression.SetLeftExpression(v.UpdateExpression(expression.LeftExpression()))
	expression.SetRightExpression(v.UpdateExpression(expression.RightExpression()))
	expression.LeftExpression().Accept(v)
	expression.RightExpression().Accept(v)
}

func (v *AutoboxingVisitor) VisitCastExpression(expression intmod.ICastExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.Expression().Accept(v)
}

func (v *AutoboxingVisitor) VisitFieldReferenceExpression(expression intmod.IFieldReferenceExpression) {
	if expression.Expression() != nil {
		expression.SetExpression(v.UpdateExpression(expression.Expression()))
		expression.Expression().Accept(v)
	}
}

func (v *AutoboxingVisitor) VisitInstanceOfExpression(expression intmod.IInstanceOfExpression) {
	if expression.Expression() != nil {
		expression.SetExpression(v.UpdateExpression(expression.Expression()))
		expression.Expression().Accept(v)
	}
}

func (v *AutoboxingVisitor) VisitLambdaFormalParametersExpression(expression intmod.ILambdaFormalParametersExpression) {
	expression.Statements().AcceptStatement(v)
}

func (v *AutoboxingVisitor) VisitLengthExpression(expression intmod.ILengthExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.Expression().Accept(v)
}

func (v *AutoboxingVisitor) VisitMethodInvocationExpression(expression intmod.IMethodInvocationExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	if expression.Parameters() != nil {
		expression.SetParameters(v.UpdateBaseExpression(expression.Parameters()))
		expression.Parameters().Accept(v)
	}
	expression.Expression().Accept(v)
}

func (v *AutoboxingVisitor) VisitMethodReferenceExpression(expression intmod.IMethodReferenceExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.Expression().Accept(v)
}

func (v *AutoboxingVisitor) VisitNewArray(expression intmod.INewArray) {
	if expression.DimensionExpressionList() != nil {
		expression.SetDimensionExpressionList(v.UpdateBaseExpression(expression.DimensionExpressionList()))
		expression.DimensionExpressionList().Accept(v)
	}
}

func (v *AutoboxingVisitor) VisitNewExpression(expression intmod.INewExpression) {
	if expression.Parameters() != nil {
		expression.SetParameters(v.UpdateBaseExpression(expression.Parameters()))
		expression.Parameters().Accept(v)
	}
	// v.SafeAccept(expression.BodyDeclaration());
}

func (v *AutoboxingVisitor) VisitNewInitializedArray(expression intmod.INewInitializedArray) {
	v.SafeAcceptDeclaration(expression.ArrayInitializer())
}

func (v *AutoboxingVisitor) VisitParenthesesExpression(expression intmod.IParenthesesExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.Expression().Accept(v)
}

func (v *AutoboxingVisitor) VisitPostOperatorExpression(expression intmod.IPostOperatorExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.Expression().Accept(v)
}

func (v *AutoboxingVisitor) VisitPreOperatorExpression(expression intmod.IPreOperatorExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.Expression().Accept(v)
}

func (v *AutoboxingVisitor) VisitSuperConstructorInvocationExpression(expression intmod.ISuperConstructorInvocationExpression) {
	if expression.Parameters() != nil {
		expression.SetParameters(v.UpdateBaseExpression(expression.Parameters()))
		expression.Parameters().Accept(v)
	}
}

func (v *AutoboxingVisitor) VisitTernaryOperatorExpression(expression intmod.ITernaryOperatorExpression) {
	expression.SetCondition(v.UpdateExpression(expression.Condition()))
	expression.SetTrueExpression(v.UpdateExpression(expression.TrueExpression()))
	expression.SetFalseExpression(v.UpdateExpression(expression.FalseExpression()))
	expression.Condition().Accept(v)
	expression.TrueExpression().Accept(v)
	expression.FalseExpression().Accept(v)
}

func (v *AutoboxingVisitor) VisitExpressionElementValue(reference intmod.IExpressionElementValue) {
	reference.SetExpression(v.UpdateExpression(reference.Expression()))
	reference.Expression().Accept(v)
}

func (v *AutoboxingVisitor) VisitAssertStatement(statement intmod.IAssertStatement) {
	statement.SetCondition(v.UpdateExpression(statement.Condition()))
	statement.Condition().Accept(v)
	v.SafeAcceptExpression(statement.Message())
}

func (v *AutoboxingVisitor) VisitDoWhileStatement(statement intmod.IDoWhileStatement) {
	statement.SetCondition(v.UpdateExpression(statement.Condition()))
	v.SafeAcceptExpression(statement.Condition())
	v.SafeAcceptStatement(statement.Statements())
}

func (v *AutoboxingVisitor) VisitExpressionStatement(statement intmod.IExpressionStatement) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
}

func (v *AutoboxingVisitor) VisitForEachStatement(statement intmod.IForEachStatement) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
	v.SafeAcceptStatement(statement.Statements())
}

func (v *AutoboxingVisitor) VisitForStatement(statement intmod.IForStatement) {
	v.SafeAcceptDeclaration(statement.Declaration())
	if statement.Init() != nil {
		statement.SetInit(v.UpdateBaseExpression(statement.Init()))
		statement.Init().Accept(v)
	}
	if statement.Condition() != nil {
		statement.SetCondition(v.UpdateExpression(statement.Condition()))
		statement.Condition().Accept(v)
	}
	if statement.Update() != nil {
		statement.SetUpdate(v.UpdateBaseExpression(statement.Update()))
		statement.Update().Accept(v)
	}
	v.SafeAcceptStatement(statement.Statements())
}

func (v *AutoboxingVisitor) VisitIfStatement(statement intmod.IIfStatement) {
	statement.SetCondition(v.UpdateExpression(statement.Condition()))
	statement.Condition().Accept(v)
	v.SafeAcceptStatement(statement.Statements())
}

func (v *AutoboxingVisitor) VisitIfElseStatement(statement intmod.IIfElseStatement) {
	statement.SetCondition(v.UpdateExpression(statement.Condition()))
	statement.Condition().Accept(v)
	v.SafeAcceptStatement(statement.Statements())
	statement.ElseStatements().AcceptStatement(v)
}

func (v *AutoboxingVisitor) VisitLambdaExpressionStatement(statement intmod.ILambdaExpressionStatement) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
}

func (v *AutoboxingVisitor) VisitReturnExpressionStatement(statement intmod.IReturnExpressionStatement) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
}

func (v *AutoboxingVisitor) VisitSwitchStatement(statement intmod.ISwitchStatement) {
	statement.SetCondition(v.UpdateExpression(statement.Condition()))
	statement.Condition().Accept(v)
	v.AcceptListStatement(statement.List())
}

func (v *AutoboxingVisitor) VisitSwitchStatementExpressionLabel(statement intmod.IExpressionLabel) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
}

func (v *AutoboxingVisitor) VisitSynchronizedStatement(statement intmod.ISynchronizedStatement) {
	statement.SetMonitor(v.UpdateExpression(statement.Monitor()))
	statement.Monitor().Accept(v)
	v.SafeAcceptStatement(statement.Statements())
}

func (v *AutoboxingVisitor) VisitThrowStatement(statement intmod.IThrowStatement) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
}

func (v *AutoboxingVisitor) VisitTryStatementCatchClause(statement intmod.ICatchClause) {
	v.SafeAcceptStatement(statement.Statements())
}

func (v *AutoboxingVisitor) VisitTryStatementResource(statement intmod.IResource) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
}

func (v *AutoboxingVisitor) VisitWhileStatement(statement intmod.IWhileStatement) {
	statement.SetCondition(v.UpdateExpression(statement.Condition()))
	statement.Condition().Accept(v)
	v.SafeAcceptStatement(statement.Statements())
}

func (v *AutoboxingVisitor) VisitConstructorReferenceExpression(_ intmod.IConstructorReferenceExpression) {
}
func (v *AutoboxingVisitor) VisitDoubleConstantExpression(_ intmod.IDoubleConstantExpression) {
}
func (v *AutoboxingVisitor) VisitEnumConstantReferenceExpression(_ intmod.IEnumConstantReferenceExpression) {
}
func (v *AutoboxingVisitor) VisitFloatConstantExpression(_ intmod.IFloatConstantExpression) {
}
func (v *AutoboxingVisitor) VisitIntegerConstantExpression(_ intmod.IIntegerConstantExpression) {
}
func (v *AutoboxingVisitor) VisitLocalVariableReferenceExpression(_ intmod.ILocalVariableReferenceExpression) {
}
func (v *AutoboxingVisitor) VisitLongConstantExpression(_ intmod.ILongConstantExpression) {
}
func (v *AutoboxingVisitor) VisitNullExpression(_ intmod.INullExpression) {
}
func (v *AutoboxingVisitor) VisitTypeReferenceDotClassExpression(_ intmod.ITypeReferenceDotClassExpression) {
}
func (v *AutoboxingVisitor) VisitObjectTypeReferenceExpression(_ intmod.IObjectTypeReferenceExpression) {
}
func (v *AutoboxingVisitor) VisitStringConstantExpression(_ intmod.IStringConstantExpression) {
}
func (v *AutoboxingVisitor) VisitSuperExpression(_ intmod.ISuperExpression) {
}
func (v *AutoboxingVisitor) VisitThisExpression(_ intmod.IThisExpression) {
}
func (v *AutoboxingVisitor) VisitAnnotationReference(_ intmod.IAnnotationReference) {
}
func (v *AutoboxingVisitor) VisitElementValueArrayInitializerElementValue(_ intmod.IElementValueArrayInitializerElementValue) {
}
func (v *AutoboxingVisitor) VisitAnnotationElementValue(_ intmod.IAnnotationElementValue) {
}
func (v *AutoboxingVisitor) VisitObjectReference(_ intmod.IObjectReference) {
}
func (v *AutoboxingVisitor) VisitBreakStatement(_ intmod.IBreakStatement) {}
func (v *AutoboxingVisitor) VisitByteCodeStatement(_ intmod.IByteCodeStatement) {
}
func (v *AutoboxingVisitor) VisitContinueStatement(_ intmod.IContinueStatement) {
}
func (v *AutoboxingVisitor) VisitReturnStatement(_ intmod.IReturnStatement) {
}
func (v *AutoboxingVisitor) VisitSwitchStatementDefaultLabel(_ intmod.IDefaultLabel) {
}

func (v *AutoboxingVisitor) VisitInnerObjectReference(_ intmod.IInnerObjectReference) {
}
func (v *AutoboxingVisitor) VisitTypeArguments(_ intmod.ITypeArguments) {}
func (v *AutoboxingVisitor) VisitWildcardExtendsTypeArgument(_ intmod.IWildcardExtendsTypeArgument) {
}
func (v *AutoboxingVisitor) VisitObjectType(_ intmod.IObjectType)           {}
func (v *AutoboxingVisitor) VisitInnerObjectType(_ intmod.IInnerObjectType) {}
func (v *AutoboxingVisitor) VisitWildcardSuperTypeArgument(_ intmod.IWildcardSuperTypeArgument) {
}
func (v *AutoboxingVisitor) VisitTypes(_ intmod.ITypes) {}
func (v *AutoboxingVisitor) VisitTypeParameterWithTypeBounds(_ intmod.ITypeParameterWithTypeBounds) {
}

func (v *AutoboxingVisitor) AcceptListDeclaration(list []intmod.IDeclaration) {
	for _, value := range list {
		value.AcceptDeclaration(v)
	}
}

func (v *AutoboxingVisitor) AcceptListExpression(list []intmod.IExpression) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *AutoboxingVisitor) AcceptListReference(list []intmod.IReference) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *AutoboxingVisitor) AcceptListStatement(list []intmod.IStatement) {
	for _, value := range list {
		value.AcceptStatement(v)
	}
}

func (v *AutoboxingVisitor) SafeAcceptDeclaration(decl intmod.IDeclaration) {
	if decl != nil {
		decl.AcceptDeclaration(v)
	}
}

func (v *AutoboxingVisitor) SafeAcceptExpression(expr intmod.IExpression) {
	if expr != nil {
		expr.Accept(v)
	}
}

func (v *AutoboxingVisitor) SafeAcceptReference(ref intmod.IReference) {
	if ref != nil {
		ref.Accept(v)
	}
}

func (v *AutoboxingVisitor) SafeAcceptStatement(list intmod.IStatement) {
	if list != nil {
		list.AcceptStatement(v)
	}
}

func (v *AutoboxingVisitor) SafeAcceptType(list intmod.IType) {
	if list != nil {
		list.AcceptTypeVisitor(v)
	}
}

func (v *AutoboxingVisitor) SafeAcceptTypeParameter(list intmod.ITypeParameter) {
	if list != nil {
		list.AcceptTypeParameterVisitor(v)
	}
}

func (v *AutoboxingVisitor) SafeAcceptListDeclaration(list []intmod.IDeclaration) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *AutoboxingVisitor) SafeAcceptListConstant(list []intmod.IConstant) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *AutoboxingVisitor) SafeAcceptListStatement(list []intmod.IStatement) {
	if list != nil {
		for _, value := range list {
			value.AcceptStatement(v)
		}
	}
}
