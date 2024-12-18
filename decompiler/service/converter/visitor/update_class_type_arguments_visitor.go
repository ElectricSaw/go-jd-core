package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	_type "github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/type"
)

func NewUpdateClassTypeArgumentsVisitor() intsrv.IUpdateClassTypeArgumentsVisitor {
	return &UpdateClassTypeArgumentsVisitor{}
}

type UpdateClassTypeArgumentsVisitor struct {
	AbstractUpdateExpressionVisitor

	result intmod.ITypeArgument
}

func (v *UpdateClassTypeArgumentsVisitor) Init() {
	v.result = nil
}

func (v *UpdateClassTypeArgumentsVisitor) TypeArgument() intmod.ITypeArgument {
	return v.result
}

func (v *UpdateClassTypeArgumentsVisitor) VisitWildcardExtendsTypeArgument(argument intmod.IWildcardExtendsTypeArgument) {
	t := argument.Type()
	t.AcceptTypeArgumentVisitor(v)

	if v.result == t {
		v.result = argument
	} else {
		v.result = _type.NewWildcardExtendsTypeArgument(v.result.(intmod.IType))
	}
}

func (v *UpdateClassTypeArgumentsVisitor) VisitWildcardSuperTypeArgument(argument intmod.IWildcardSuperTypeArgument) {
	t := argument.Type()
	t.AcceptTypeArgumentVisitor(v)

	if v.result == t {
		v.result = argument
	} else {
		v.result = _type.NewWildcardSuperTypeArgument(v.result.(intmod.IType))
	}
}

func (v *UpdateClassTypeArgumentsVisitor) VisitDiamondTypeArgument(argument intmod.IDiamondTypeArgument) {
	v.result = argument
}
func (v *UpdateClassTypeArgumentsVisitor) VisitWildcardTypeArgument(argument intmod.IWildcardTypeArgument) {
	v.result = argument
}

func (v *UpdateClassTypeArgumentsVisitor) VisitPrimitiveType(t intmod.IPrimitiveType) { v.result = t }
func (v *UpdateClassTypeArgumentsVisitor) VisitGenericType(t intmod.IGenericType)     { v.result = t }

func (v *UpdateClassTypeArgumentsVisitor) VisitObjectType(t intmod.IObjectType) {
	typeArguments := t.TypeArguments()

	if typeArguments == nil {
		if t.InternalName() == _type.OtTypeClass.InternalName() {
			v.result = _type.OtTypeClassWildcard
		} else {
			v.result = t
		}
	} else {
		typeArguments.AcceptTypeArgumentVisitor(v)

		if v.result == typeArguments {
			v.result = t
		} else {
			v.result = t.CreateTypeWithArgs(typeArguments)
		}
	}
}

func (v *UpdateClassTypeArgumentsVisitor) VisitInnerObjectType(t intmod.IInnerObjectType) {
	t.OuterType().AcceptTypeArgumentVisitor(v)

	typeArguments := t.TypeArguments()

	if t.OuterType() == v.result {
		if typeArguments == nil {
			v.result = t
		} else {
			typeArguments.AcceptTypeArgumentVisitor(v)
			if v.result == typeArguments {
				v.result = t
			} else {
				v.result = t.CreateTypeWithArgs(v.result)
			}
		}
	} else {
		outerObjectType := v.result.(intmod.IObjectType)

		if typeArguments != nil {
			typeArguments.AcceptTypeArgumentVisitor(v)
			typeArguments = v.result
		}

		v.result = _type.NewInnerObjectTypeWithAll(t.InternalName(), t.QualifiedName(), t.Name(), typeArguments, t.Dimension(), outerObjectType)
	}
}

func (v *UpdateClassTypeArgumentsVisitor) VisitTypeArguments(arguments intmod.ITypeArguments) {
	size := arguments.Size()
	i := 0

	for i = 0; i < size; i++ {
		ta := arguments.Get(i).(intmod.ITypeArgument)
		ta.AcceptTypeArgumentVisitor(v)
		if v.result != ta {
			break
		}
	}

	if v.result != nil {
		if i == size {
			v.result = arguments
		} else {
			newTypes := _type.NewTypeArgumentsWithCapacity(size)
			newTypes.AddAll(arguments.ToSlice()[:i])
			newTypes.Add(v.result.(intmod.ITypeArgument))

			for i++; i < size; i++ {
				ta := arguments.Get(i).(intmod.ITypeArgument)
				ta.AcceptTypeArgumentVisitor(v)
				newTypes.Add(v.result.(intmod.ITypeArgument))
			}

			v.result = newTypes
		}
	}
}

// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$                                 $$$
// $$$ AbstractUpdateExpressionVisitor $$$
// $$$                                 $$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func (v *UpdateClassTypeArgumentsVisitor) UpdateExpression(_ intmod.IExpression) intmod.IExpression {
	return nil
}

func (v *UpdateClassTypeArgumentsVisitor) UpdateBaseExpression(baseExpression intmod.IExpression) intmod.IExpression {
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

func (v *UpdateClassTypeArgumentsVisitor) VisitAnnotationDeclaration(decl intmod.IAnnotationDeclaration) {
	v.SafeAcceptDeclaration(decl.AnnotationDeclarators())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *UpdateClassTypeArgumentsVisitor) VisitClassDeclaration(declaration intmod.IClassDeclaration) {
	v.VisitInterfaceDeclaration(declaration)
}

func (v *UpdateClassTypeArgumentsVisitor) VisitConstructorInvocationExpression(expression intmod.IConstructorInvocationExpression) {
	if expression.Parameters() != nil {
		expression.SetParameters(v.UpdateBaseExpression(expression.Parameters()))
		expression.Parameters().Accept(v)
	}
}

func (v *UpdateClassTypeArgumentsVisitor) VisitConstructorDeclaration(declaration intmod.IConstructorDeclaration) {
	v.SafeAcceptStatement(declaration.Statements())
}

func (v *UpdateClassTypeArgumentsVisitor) VisitEnumDeclaration(declaration intmod.IEnumDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *UpdateClassTypeArgumentsVisitor) VisitEnumDeclarationConstant(declaration intmod.IConstant) {
	if declaration.Arguments() != nil {
		declaration.SetArguments(v.UpdateBaseExpression(declaration.Arguments()))
		declaration.Arguments().Accept(v)
	}
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *UpdateClassTypeArgumentsVisitor) VisitExpressionVariableInitializer(declaration intmod.IExpressionVariableInitializer) {
	if declaration.Expression() != nil {
		declaration.SetExpression(v.UpdateExpression(declaration.Expression()))
		declaration.Expression().Accept(v)
	}
}

func (v *UpdateClassTypeArgumentsVisitor) VisitFieldDeclaration(declaration intmod.IFieldDeclaration) {
	declaration.FieldDeclarators().AcceptDeclaration(v)
}

func (v *UpdateClassTypeArgumentsVisitor) VisitFieldDeclarator(declaration intmod.IFieldDeclarator) {
	v.SafeAcceptDeclaration(declaration.VariableInitializer())
}

func (v *UpdateClassTypeArgumentsVisitor) VisitFormalParameter(_ intmod.IFormalParameter) {
}

func (v *UpdateClassTypeArgumentsVisitor) VisitInterfaceDeclaration(declaration intmod.IInterfaceDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *UpdateClassTypeArgumentsVisitor) VisitLocalVariableDeclaration(declaration intmod.ILocalVariableDeclaration) {
	declaration.LocalVariableDeclarators().AcceptDeclaration(v)
}

func (v *UpdateClassTypeArgumentsVisitor) VisitLocalVariableDeclarator(declarator intmod.ILocalVariableDeclarator) {
	v.SafeAcceptDeclaration(declarator.VariableInitializer())
}

func (v *UpdateClassTypeArgumentsVisitor) VisitMethodDeclaration(declaration intmod.IMethodDeclaration) {
	v.SafeAcceptReference(declaration.AnnotationReferences())
	v.SafeAcceptStatement(declaration.Statements())
}

func (v *UpdateClassTypeArgumentsVisitor) VisitArrayExpression(expression intmod.IArrayExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.SetIndex(v.UpdateExpression(expression.Index()))
	expression.Expression().Accept(v)
	expression.Index().Accept(v)
}

func (v *UpdateClassTypeArgumentsVisitor) VisitBinaryOperatorExpression(expression intmod.IBinaryOperatorExpression) {
	expression.SetLeftExpression(v.UpdateExpression(expression.LeftExpression()))
	expression.SetRightExpression(v.UpdateExpression(expression.RightExpression()))
	expression.LeftExpression().Accept(v)
	expression.RightExpression().Accept(v)
}

func (v *UpdateClassTypeArgumentsVisitor) VisitCastExpression(expression intmod.ICastExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.Expression().Accept(v)
}

func (v *UpdateClassTypeArgumentsVisitor) VisitFieldReferenceExpression(expression intmod.IFieldReferenceExpression) {
	if expression.Expression() != nil {
		expression.SetExpression(v.UpdateExpression(expression.Expression()))
		expression.Expression().Accept(v)
	}
}

func (v *UpdateClassTypeArgumentsVisitor) VisitInstanceOfExpression(expression intmod.IInstanceOfExpression) {
	if expression.Expression() != nil {
		expression.SetExpression(v.UpdateExpression(expression.Expression()))
		expression.Expression().Accept(v)
	}
}

func (v *UpdateClassTypeArgumentsVisitor) VisitLambdaFormalParametersExpression(expression intmod.ILambdaFormalParametersExpression) {
	expression.Statements().AcceptStatement(v)
}

func (v *UpdateClassTypeArgumentsVisitor) VisitLengthExpression(expression intmod.ILengthExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.Expression().Accept(v)
}

func (v *UpdateClassTypeArgumentsVisitor) VisitMethodInvocationExpression(expression intmod.IMethodInvocationExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	if expression.Parameters() != nil {
		expression.SetParameters(v.UpdateBaseExpression(expression.Parameters()))
		expression.Parameters().Accept(v)
	}
	expression.Expression().Accept(v)
}

func (v *UpdateClassTypeArgumentsVisitor) VisitMethodReferenceExpression(expression intmod.IMethodReferenceExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.Expression().Accept(v)
}

func (v *UpdateClassTypeArgumentsVisitor) VisitNewArray(expression intmod.INewArray) {
	if expression.DimensionExpressionList() != nil {
		expression.SetDimensionExpressionList(v.UpdateBaseExpression(expression.DimensionExpressionList()))
		expression.DimensionExpressionList().Accept(v)
	}
}

func (v *UpdateClassTypeArgumentsVisitor) VisitNewExpression(expression intmod.INewExpression) {
	if expression.Parameters() != nil {
		expression.SetParameters(v.UpdateBaseExpression(expression.Parameters()))
		expression.Parameters().Accept(v)
	}
	// v.SafeAccept(expression.BodyDeclaration());
}

func (v *UpdateClassTypeArgumentsVisitor) VisitNewInitializedArray(expression intmod.INewInitializedArray) {
	v.SafeAcceptDeclaration(expression.ArrayInitializer())
}

func (v *UpdateClassTypeArgumentsVisitor) VisitParenthesesExpression(expression intmod.IParenthesesExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.Expression().Accept(v)
}

func (v *UpdateClassTypeArgumentsVisitor) VisitPostOperatorExpression(expression intmod.IPostOperatorExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.Expression().Accept(v)
}

func (v *UpdateClassTypeArgumentsVisitor) VisitPreOperatorExpression(expression intmod.IPreOperatorExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.Expression().Accept(v)
}

func (v *UpdateClassTypeArgumentsVisitor) VisitSuperConstructorInvocationExpression(expression intmod.ISuperConstructorInvocationExpression) {
	if expression.Parameters() != nil {
		expression.SetParameters(v.UpdateBaseExpression(expression.Parameters()))
		expression.Parameters().Accept(v)
	}
}

func (v *UpdateClassTypeArgumentsVisitor) VisitTernaryOperatorExpression(expression intmod.ITernaryOperatorExpression) {
	expression.SetCondition(v.UpdateExpression(expression.Condition()))
	expression.SetTrueExpression(v.UpdateExpression(expression.TrueExpression()))
	expression.SetFalseExpression(v.UpdateExpression(expression.FalseExpression()))
	expression.Condition().Accept(v)
	expression.TrueExpression().Accept(v)
	expression.FalseExpression().Accept(v)
}

func (v *UpdateClassTypeArgumentsVisitor) VisitExpressionElementValue(reference intmod.IExpressionElementValue) {
	reference.SetExpression(v.UpdateExpression(reference.Expression()))
	reference.Expression().Accept(v)
}

func (v *UpdateClassTypeArgumentsVisitor) VisitAssertStatement(statement intmod.IAssertStatement) {
	statement.SetCondition(v.UpdateExpression(statement.Condition()))
	statement.Condition().Accept(v)
	v.SafeAcceptExpression(statement.Message())
}

func (v *UpdateClassTypeArgumentsVisitor) VisitDoWhileStatement(statement intmod.IDoWhileStatement) {
	statement.SetCondition(v.UpdateExpression(statement.Condition()))
	v.SafeAcceptExpression(statement.Condition())
	v.SafeAcceptStatement(statement.Statements())
}

func (v *UpdateClassTypeArgumentsVisitor) VisitExpressionStatement(statement intmod.IExpressionStatement) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
}

func (v *UpdateClassTypeArgumentsVisitor) VisitForEachStatement(statement intmod.IForEachStatement) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
	v.SafeAcceptStatement(statement.Statements())
}

func (v *UpdateClassTypeArgumentsVisitor) VisitForStatement(statement intmod.IForStatement) {
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

func (v *UpdateClassTypeArgumentsVisitor) VisitIfStatement(statement intmod.IIfStatement) {
	statement.SetCondition(v.UpdateExpression(statement.Condition()))
	statement.Condition().Accept(v)
	v.SafeAcceptStatement(statement.Statements())
}

func (v *UpdateClassTypeArgumentsVisitor) VisitIfElseStatement(statement intmod.IIfElseStatement) {
	statement.SetCondition(v.UpdateExpression(statement.Condition()))
	statement.Condition().Accept(v)
	v.SafeAcceptStatement(statement.Statements())
	statement.ElseStatements().AcceptStatement(v)
}

func (v *UpdateClassTypeArgumentsVisitor) VisitLambdaExpressionStatement(statement intmod.ILambdaExpressionStatement) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
}

func (v *UpdateClassTypeArgumentsVisitor) VisitReturnExpressionStatement(statement intmod.IReturnExpressionStatement) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
}

func (v *UpdateClassTypeArgumentsVisitor) VisitSwitchStatement(statement intmod.ISwitchStatement) {
	statement.SetCondition(v.UpdateExpression(statement.Condition()))
	statement.Condition().Accept(v)
	v.AcceptListStatement(statement.List())
}

func (v *UpdateClassTypeArgumentsVisitor) VisitSwitchStatementExpressionLabel(statement intmod.IExpressionLabel) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
}

func (v *UpdateClassTypeArgumentsVisitor) VisitSynchronizedStatement(statement intmod.ISynchronizedStatement) {
	statement.SetMonitor(v.UpdateExpression(statement.Monitor()))
	statement.Monitor().Accept(v)
	v.SafeAcceptStatement(statement.Statements())
}

func (v *UpdateClassTypeArgumentsVisitor) VisitThrowStatement(statement intmod.IThrowStatement) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
}

func (v *UpdateClassTypeArgumentsVisitor) VisitTryStatementCatchClause(statement intmod.ICatchClause) {
	v.SafeAcceptStatement(statement.Statements())
}

func (v *UpdateClassTypeArgumentsVisitor) VisitTryStatementResource(statement intmod.IResource) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
}

func (v *UpdateClassTypeArgumentsVisitor) VisitWhileStatement(statement intmod.IWhileStatement) {
	statement.SetCondition(v.UpdateExpression(statement.Condition()))
	statement.Condition().Accept(v)
	v.SafeAcceptStatement(statement.Statements())
}

func (v *UpdateClassTypeArgumentsVisitor) VisitConstructorReferenceExpression(_ intmod.IConstructorReferenceExpression) {
}
func (v *UpdateClassTypeArgumentsVisitor) VisitDoubleConstantExpression(_ intmod.IDoubleConstantExpression) {
}
func (v *UpdateClassTypeArgumentsVisitor) VisitEnumConstantReferenceExpression(_ intmod.IEnumConstantReferenceExpression) {
}
func (v *UpdateClassTypeArgumentsVisitor) VisitFloatConstantExpression(_ intmod.IFloatConstantExpression) {
}
func (v *UpdateClassTypeArgumentsVisitor) VisitIntegerConstantExpression(_ intmod.IIntegerConstantExpression) {
}
func (v *UpdateClassTypeArgumentsVisitor) VisitLocalVariableReferenceExpression(_ intmod.ILocalVariableReferenceExpression) {
}
func (v *UpdateClassTypeArgumentsVisitor) VisitLongConstantExpression(_ intmod.ILongConstantExpression) {
}
func (v *UpdateClassTypeArgumentsVisitor) VisitNullExpression(_ intmod.INullExpression) {
}
func (v *UpdateClassTypeArgumentsVisitor) VisitTypeReferenceDotClassExpression(_ intmod.ITypeReferenceDotClassExpression) {
}
func (v *UpdateClassTypeArgumentsVisitor) VisitObjectTypeReferenceExpression(_ intmod.IObjectTypeReferenceExpression) {
}
func (v *UpdateClassTypeArgumentsVisitor) VisitStringConstantExpression(_ intmod.IStringConstantExpression) {
}
func (v *UpdateClassTypeArgumentsVisitor) VisitSuperExpression(_ intmod.ISuperExpression) {
}
func (v *UpdateClassTypeArgumentsVisitor) VisitThisExpression(_ intmod.IThisExpression) {
}
func (v *UpdateClassTypeArgumentsVisitor) VisitAnnotationReference(_ intmod.IAnnotationReference) {
}
func (v *UpdateClassTypeArgumentsVisitor) VisitElementValueArrayInitializerElementValue(_ intmod.IElementValueArrayInitializerElementValue) {
}
func (v *UpdateClassTypeArgumentsVisitor) VisitAnnotationElementValue(_ intmod.IAnnotationElementValue) {
}
func (v *UpdateClassTypeArgumentsVisitor) VisitObjectReference(_ intmod.IObjectReference) {
}
func (v *UpdateClassTypeArgumentsVisitor) VisitBreakStatement(_ intmod.IBreakStatement) {}
func (v *UpdateClassTypeArgumentsVisitor) VisitByteCodeStatement(_ intmod.IByteCodeStatement) {
}
func (v *UpdateClassTypeArgumentsVisitor) VisitContinueStatement(_ intmod.IContinueStatement) {
}
func (v *UpdateClassTypeArgumentsVisitor) VisitReturnStatement(_ intmod.IReturnStatement) {
}
func (v *UpdateClassTypeArgumentsVisitor) VisitSwitchStatementDefaultLabel(_ intmod.IDefaultLabel) {
}

func (v *UpdateClassTypeArgumentsVisitor) VisitInnerObjectReference(_ intmod.IInnerObjectReference) {
}
func (v *UpdateClassTypeArgumentsVisitor) VisitTypes(_ intmod.ITypes) {}
func (v *UpdateClassTypeArgumentsVisitor) VisitTypeParameterWithTypeBounds(_ intmod.ITypeParameterWithTypeBounds) {
}

func (v *UpdateClassTypeArgumentsVisitor) AcceptListDeclaration(list []intmod.IDeclaration) {
	for _, value := range list {
		value.AcceptDeclaration(v)
	}
}

func (v *UpdateClassTypeArgumentsVisitor) AcceptListExpression(list []intmod.IExpression) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *UpdateClassTypeArgumentsVisitor) AcceptListReference(list []intmod.IReference) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *UpdateClassTypeArgumentsVisitor) AcceptListStatement(list []intmod.IStatement) {
	for _, value := range list {
		value.AcceptStatement(v)
	}
}

func (v *UpdateClassTypeArgumentsVisitor) SafeAcceptDeclaration(decl intmod.IDeclaration) {
	if decl != nil {
		decl.AcceptDeclaration(v)
	}
}

func (v *UpdateClassTypeArgumentsVisitor) SafeAcceptExpression(expr intmod.IExpression) {
	if expr != nil {
		expr.Accept(v)
	}
}

func (v *UpdateClassTypeArgumentsVisitor) SafeAcceptReference(ref intmod.IReference) {
	if ref != nil {
		ref.Accept(v)
	}
}

func (v *UpdateClassTypeArgumentsVisitor) SafeAcceptStatement(list intmod.IStatement) {
	if list != nil {
		list.AcceptStatement(v)
	}
}

func (v *UpdateClassTypeArgumentsVisitor) SafeAcceptType(list intmod.IType) {
	if list != nil {
		list.AcceptTypeVisitor(v)
	}
}

func (v *UpdateClassTypeArgumentsVisitor) SafeAcceptTypeParameter(list intmod.ITypeParameter) {
	if list != nil {
		list.AcceptTypeParameterVisitor(v)
	}
}

func (v *UpdateClassTypeArgumentsVisitor) SafeAcceptListDeclaration(list []intmod.IDeclaration) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *UpdateClassTypeArgumentsVisitor) SafeAcceptListConstant(list []intmod.IConstant) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *UpdateClassTypeArgumentsVisitor) SafeAcceptListStatement(list []intmod.IStatement) {
	if list != nil {
		for _, value := range list {
			value.AcceptStatement(v)
		}
	}
}
