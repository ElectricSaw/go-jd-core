package model

type AbstractJavaSyntaxVisitor struct {
}

func (v *AbstractJavaSyntaxVisitor) VisitCompilationUnit(compilationUnit *CompilationUnit) {
	compilationUnit.TypeDeclarations.AcceptDeclaration(v)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////
// DeclarationVisitor
//////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (v *AbstractJavaSyntaxVisitor) VisitAnnotationDeclaration(declaration IAnnotationDeclaration) {
	v.SafeAcceptDeclaration(declaration.GetAnnotationDeclarators())
	v.SafeAcceptDeclaration(declaration.GetBodyDeclaration())
	v.SafeAcceptReference(declaration.GetAnnotationReferences())
}

func (v *AbstractJavaSyntaxVisitor) VisitArrayVariableInitializer(declaration *ArrayVariableInitializer) {
	list := make([]IDeclaration, declaration.Size())
	for i, element := range declaration.ToSlice() {
		list[i] = element
	}
	v.AcceptListDeclaration(list)
}

func (v *AbstractJavaSyntaxVisitor) VisitBodyDeclaration(declaration IBodyDeclaration) {
	v.SafeAcceptDeclaration(declaration.GetMemberDeclaration())
}

func (v *AbstractJavaSyntaxVisitor) VisitClassDeclaration(declaration IClassDeclaration) {
	superType := declaration.GetSuperType()

	if superType != nil {
		superType.AcceptTypeVisitor(v)
	}

	v.SafeAcceptTypeParameter(declaration.GetTypeParameters())
	v.SafeAcceptType(declaration.GetInterfaces())
	v.SafeAcceptReference(declaration.GetAnnotationReferences())
	v.SafeAcceptDeclaration(declaration.GetBodyDeclaration())
}

func (v *AbstractJavaSyntaxVisitor) VisitConstructorDeclaration(declaration IConstructorDeclaration) {
	v.SafeAcceptReference(declaration.GetAnnotationReferences())
	v.SafeAcceptDeclaration(declaration.GetFormalParameters())
	v.SafeAcceptType(declaration.GetExceptionTypes())
	v.SafeAcceptStatement(declaration.GetStatements())
}

func (v *AbstractJavaSyntaxVisitor) VisitEnumDeclaration(declaration IEnumDeclaration) {
	v.VisitTypeDeclaration(declaration)
	v.SafeAcceptType(declaration.GetInterfaces())
	v.SafeAcceptListConstant(declaration.GetConstants().ToSlice())
	v.SafeAcceptDeclaration(declaration.GetBodyDeclaration())
}

func (v *AbstractJavaSyntaxVisitor) VisitEnumDeclarationConstant(declaration IConstant) {
	v.SafeAcceptReference(declaration.GetAnnotationReferences())
	v.SafeAcceptExpression(declaration.GetArguments())
	v.SafeAcceptDeclaration(declaration.GetBodyDeclaration())
}

func (v *AbstractJavaSyntaxVisitor) VisitExpressionVariableInitializer(declaration *ExpressionVariableInitializer) {
	declaration.Expression.Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitFieldDeclaration(declaration IFieldDeclaration) {
	t := declaration.GetType()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptReference(declaration.GetAnnotationReferences())
	declaration.GetFieldDeclarators().AcceptDeclaration(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitFieldDeclarator(declaration *FieldDeclarator) {
	v.SafeAcceptDeclaration(declaration.VariableInitializer)
}

func (v *AbstractJavaSyntaxVisitor) VisitFieldDeclarators(declarations *FieldDeclarators) {
	list := make([]IDeclaration, declarations.Size())
	for i, element := range declarations.ToSlice() {
		list[i] = element
	}
	v.AcceptListDeclaration(list)
}

func (v *AbstractJavaSyntaxVisitor) VisitFormalParameter(declaration IFormalParameter) {
	t := declaration.GetType()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptReference(declaration.GetAnnotationReferences())
}

func (v *AbstractJavaSyntaxVisitor) VisitFormalParameters(declarations *FormalParameters) {
	list := make([]IDeclaration, declarations.Size())
	for i, element := range declarations.ToSlice() {
		list[i] = element
	}
	v.AcceptListDeclaration(list)
}

func (v *AbstractJavaSyntaxVisitor) VisitInstanceInitializerDeclaration(declaration *InstanceInitializerDeclaration) {
	v.SafeAcceptStatement(declaration.Statements)
}

func (v *AbstractJavaSyntaxVisitor) VisitInterfaceDeclaration(declaration IInterfaceDeclaration) {
	v.SafeAcceptType(declaration.GetInterfaces())
	v.SafeAcceptReference(declaration.GetAnnotationReferences())
	v.SafeAcceptDeclaration(declaration.GetBodyDeclaration())
}

func (v *AbstractJavaSyntaxVisitor) VisitLocalVariableDeclaration(declaration ILocalVariableDeclaration) {
	t := declaration.GetType()

	t.AcceptTypeVisitor(v)
	declaration.GetLocalVariableDeclarators().AcceptDeclaration(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitLocalVariableDeclarator(declarator ILocalVariableDeclarator) {
	v.SafeAcceptDeclaration(declarator.GetVariableInitializer())
}

func (v *AbstractJavaSyntaxVisitor) VisitLocalVariableDeclarators(declarators *LocalVariableDeclarators) {
	list := make([]IDeclaration, declarators.Size())
	for i, element := range declarators.ToSlice() {
		list[i] = element
	}
	v.AcceptListDeclaration(list)
}

func (v *AbstractJavaSyntaxVisitor) VisitMethodDeclaration(declaration IMethodDeclaration) {
	t := declaration.GetReturnedType()
	t.AcceptTypeVisitor(v)

	v.SafeAcceptReference(declaration.GetAnnotationReferences())
	v.SafeAcceptDeclaration(declaration.GetFormalParameters())
	v.SafeAcceptType(declaration.GetExceptionTypes())
	v.SafeAcceptStatement(declaration.GetStatements())
}

func (v *AbstractJavaSyntaxVisitor) VisitMemberDeclarations(declarations *MemberDeclarations) {
	list := make([]IDeclaration, declarations.Size())
	for i, element := range declarations.ToSlice() {
		list[i] = element
	}
	v.AcceptListDeclaration(list)
}

func (v *AbstractJavaSyntaxVisitor) VisitModuleDeclaration(_ *ModuleDeclaration) {
	// EMPTY
}

func (v *AbstractJavaSyntaxVisitor) VisitStaticInitializerDeclaration(declaration IStaticInitializerDeclaration) {
	v.SafeAcceptStatement(declaration.GetStatements())
}

func (v *AbstractJavaSyntaxVisitor) VisitTypeDeclarations(declarations *TypeDeclarations) {
	list := make([]IDeclaration, declarations.Size())
	for i, element := range declarations.ToSlice() {
		list[i] = element
	}
	v.AcceptListDeclaration(list)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////
// ExpressionVisitor
//////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (v *AbstractJavaSyntaxVisitor) VisitArrayExpression(expression *ArrayExpression) {
	t := expression.Type
	t.AcceptTypeVisitor(v)

	expression.Expression.Accept(v)
	expression.Index.Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitBinaryOperatorExpression(expression *BinaryOperatorExpression) {
	expression.LeftExpression.Accept(v)
	expression.RightExpression.Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitBooleanExpression(_ *BooleanExpression) {
	// EMPTY
}

func (v *AbstractJavaSyntaxVisitor) VisitCastExpression(expression *CastExpression) {
	t := expression.Type
	t.AcceptTypeVisitor(v)
	expression.Expression.Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitCommentExpression(_ *CommentExpression) {
	// EMPTY
}

func (v *AbstractJavaSyntaxVisitor) VisitConstructorInvocationExpression(expression *ConstructorInvocationExpression) {
	t := expression.Type
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expression.Parameters)
}

func (v *AbstractJavaSyntaxVisitor) VisitConstructorReferenceExpression(expression *ConstructorReferenceExpression) {
	t := expression.Type
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitDoubleConstantExpression(expression *DoubleConstantExpression) {
	t := expression.Type
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitEnumConstantReferenceExpression(expression *EnumConstantReferenceExpression) {
	t := expression.Type
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitExpressions(expression *Expressions) {
	list := make([]IExpression, expression.Size())
	for i, element := range expression.ToSlice() {
		list[i] = element
	}
	v.AcceptListExpression(list)
}

func (v *AbstractJavaSyntaxVisitor) VisitFieldReferenceExpression(expression *FieldReferenceExpression) {
	t := expression.Type
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expression.Expression)
}

func (v *AbstractJavaSyntaxVisitor) VisitFloatConstantExpression(expression *FloatConstantExpression) {
	t := expression.Type
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitIntegerConstantExpression(expression *IntegerConstantExpression) {
	t := expression.Type
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitInstanceOfExpression(expression *InstanceOfExpression) {
	t := expression.Type
	t.AcceptTypeVisitor(v)
	expression.Expression.Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitLambdaFormalParametersExpression(expression *LambdaFormalParametersExpression) {
	v.SafeAcceptDeclaration(expression.FormalParameters)
	expression.Statements.AcceptStatement(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitLambdaIdentifiersExpression(expression *LambdaIdentifiersExpression) {
	v.SafeAcceptStatement(expression.Statements)
}

func (v *AbstractJavaSyntaxVisitor) VisitLengthExpression(expression *LengthExpression) {
	expression.Expression.Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitLocalVariableReferenceExpression(expression *LocalVariableReferenceExpression) {
	t := expression.Type
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitLongConstantExpression(expression *LongConstantExpression) {
	t := expression.Type
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitMethodInvocationExpression(expression *MethodInvocationExpression) {
	expression.Expression.Accept(v)
	v.SafeAcceptTypeArgumentVisitable(expression.NonWildcardTypeArguments.(*WildcardSuperTypeArgument))
	v.SafeAcceptExpression(expression.Parameters)
}

func (v *AbstractJavaSyntaxVisitor) VisitMethodReferenceExpression(expression *MethodReferenceExpression) {
	expression.Expression.Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitNewArray(expression *NewArray) {
	t := expression.Type
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expression.DimensionExpressionList)
}

func (v *AbstractJavaSyntaxVisitor) VisitNewExpression(expression *NewExpression) {
	t := expression.Type
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expression.Parameters)
}

func (v *AbstractJavaSyntaxVisitor) VisitNewInitializedArray(expression *NewInitializedArray) {
	t := expression.Type
	t.AcceptTypeVisitor(v)
	v.SafeAcceptDeclaration(expression.ArrayInitializer)
}

func (v *AbstractJavaSyntaxVisitor) VisitNoExpression(_ *NoExpression) {
	// EMPTY
}

func (v *AbstractJavaSyntaxVisitor) VisitNullExpression(expression *NullExpression) {
	t := expression.Type
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitObjectTypeReferenceExpression(expression *ObjectTypeReferenceExpression) {
	t := expression.Type
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitParenthesesExpression(expression *ParenthesesExpression) {
	expression.Expression.Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitPostOperatorExpression(expression *PostOperatorExpression) {
	expression.Expression.Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitPreOperatorExpression(expression *PreOperatorExpression) {
	expression.Expression.Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitStringConstantExpression(_ *StringConstantExpression) {
	// EMPTY
}

func (v *AbstractJavaSyntaxVisitor) VisitSuperConstructorInvocationExpression(expression *SuperConstructorInvocationExpression) {
	t := expression.Type
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expression.Parameters)
}

func (v *AbstractJavaSyntaxVisitor) VisitSuperExpression(expression *SuperExpression) {
	t := expression.Type
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitTernaryOperatorExpression(expression *TernaryOperatorExpression) {
	expression.Condition.Accept(v)
	expression.TrueExpression.Accept(v)
	expression.FalseExpression.Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitThisExpression(expression *ThisExpression) {
	t := expression.Type
	t.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitTypeReferenceDotClassExpression(expression *TypeReferenceDotClassExpression) {
	t := expression.Type
	t.AcceptTypeVisitor(v)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////
// ReferenceVisitor
//////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (v *AbstractJavaSyntaxVisitor) VisitAnnotationElementValue(reference *AnnotationElementValue) {
	v.SafeAcceptReference(reference.ElementValue)
	v.SafeAcceptReference(reference.ElementValuePairs)
}

func (v *AbstractJavaSyntaxVisitor) VisitAnnotationReference(reference *AnnotationReference) {
	v.SafeAcceptReference(reference.ElementValue)
	v.SafeAcceptReference(reference.ElementValuePairs)
}

func (v *AbstractJavaSyntaxVisitor) VisitAnnotationReferences(references *AnnotationReferences) {
	list := make([]IReference, references.Size())
	for i, element := range references.ToSlice() {
		list[i] = element
	}
	v.AcceptListReference(list)
}

func (v *AbstractJavaSyntaxVisitor) VisitElementValueArrayInitializerElementValue(reference *ElementValueArrayInitializerElementValue) {
	v.SafeAcceptReference(reference.ElementValueArrayInitializer)
}

func (v *AbstractJavaSyntaxVisitor) VisitElementValues(references *ElementValues) {
	list := make([]IReference, references.Size())
	for i, element := range references.ToSlice() {
		list[i] = element
	}
	v.AcceptListReference(list)
}

func (v *AbstractJavaSyntaxVisitor) VisitElementValuePair(reference *ElementValuePair) {
	reference.ElementValue.Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitElementValuePairs(references *ElementValuePairs) {
	list := make([]IReference, references.Size())
	for i, element := range references.ToSlice() {
		list[i] = element
	}
	v.AcceptListReference(list)
}

func (v *AbstractJavaSyntaxVisitor) VisitExpressionElementValue(reference *ExpressionElementValue) {
	reference.Expression.Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitInnerObjectReference(reference *InnerObjectReference) {
	v.VisitInnerObjectType(reference)
}

func (v *AbstractJavaSyntaxVisitor) VisitObjectReference(reference *ObjectReference) {
	v.VisitObjectType(reference)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////
// StatementVisitor
//////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (v *AbstractJavaSyntaxVisitor) VisitAssertStatement(statement *AssertStatement) {
	statement.Condition.Accept(v)
	v.SafeAcceptExpression(statement.Message)
}

func (v *AbstractJavaSyntaxVisitor) VisitBreakStatement(_ *BreakStatement) {
	// EMPTY
}

func (v *AbstractJavaSyntaxVisitor) VisitByteCodeStatement(_ *ByteCodeStatement) {
	// EMPTY
}

func (v *AbstractJavaSyntaxVisitor) VisitCommentStatement(_ *CommentStatement) {
	// EMPTY
}

func (v *AbstractJavaSyntaxVisitor) VisitContinueStatement(_ *ContinueStatement) {
	// EMPTY
}

func (v *AbstractJavaSyntaxVisitor) VisitDoWhileStatement(statement *DoWhileStatement) {
	v.SafeAcceptExpression(statement.Condition)
	v.SafeAcceptStatement(statement.Statements)
}

func (v *AbstractJavaSyntaxVisitor) VisitExpressionStatement(statement *ExpressionStatement) {
	statement.Expression.Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitForEachStatement(statement *ForEachStatement) {
	t := statement.Type
	t.AcceptTypeVisitor(v)
	statement.Expression.Accept(v)
	v.SafeAcceptStatement(statement.Statements)
}

func (v *AbstractJavaSyntaxVisitor) VisitForStatement(statement *ForStatement) {
	v.SafeAcceptDeclaration(statement.Declaration)
	v.SafeAcceptExpression(statement.Init)
	v.SafeAcceptExpression(statement.Condition)
	v.SafeAcceptExpression(statement.Update)
	v.SafeAcceptStatement(statement.Statements)
}

func (v *AbstractJavaSyntaxVisitor) VisitIfStatement(statement *IfStatement) {
	statement.Condition.Accept(v)
	v.SafeAcceptStatement(statement.Statements)
}

func (v *AbstractJavaSyntaxVisitor) VisitIfElseStatement(statement *IfElseStatement) {
	statement.Condition.Accept(v)
	v.SafeAcceptStatement(statement.IfStatements)
	statement.ElseStatements.AcceptStatement(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitLabelStatement(statement *LabelStatement) {
	v.SafeAcceptStatement(statement.Statement)
}

func (v *AbstractJavaSyntaxVisitor) VisitLambdaExpressionStatement(statement *LambdaExpressionStatement) {
	statement.Expression.Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitLocalVariableDeclarationStatement(statement ILocalVariableDeclaration) {
	v.VisitLocalVariableDeclaration(statement)
}

func (v *AbstractJavaSyntaxVisitor) VisitNoStatement(_ *NoStatement) {
	// EMPTY
}

func (v *AbstractJavaSyntaxVisitor) VisitReturnExpressionStatement(statement *ReturnExpressionStatement) {
	statement.Expression.Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitReturnStatement(_ *ReturnStatement) {
	// EMPTY
}

func (v *AbstractJavaSyntaxVisitor) VisitStatements(statement *Statements) {
	list := make([]IStatement, statement.Size())
	for i, element := range statement.ToSlice() {
		list[i] = element
	}
	v.AcceptListStatement(list)
}

func (v *AbstractJavaSyntaxVisitor) VisitSwitchStatement(statement *SwitchStatement) {
	statement.Condition.Accept(v)
	list := make([]IStatement, statement.Blocks.Size())
	for i, element := range statement.Blocks.ToSlice() {
		list[i] = element
	}
	v.AcceptListStatement(list)
}

func (v *AbstractJavaSyntaxVisitor) VisitSwitchStatementDefaultLabel(_ *DefaultLabel) {
	// EMPTY
}

func (v *AbstractJavaSyntaxVisitor) VisitSwitchStatementExpressionLabel(statement *ExpressionLabel) {
	statement.Expression.Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitSwitchStatementLabelBlock(statement *LabelBlock) {
	statement.Label.AcceptStatement(v)
	statement.Statements.AcceptStatement(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitSwitchStatementMultiLabelsBlock(statement *MultiLabelsBlock) {
	list := make([]IStatement, statement.Labels.Size())
	for i, element := range statement.Labels.ToSlice() {
		list[i] = element
	}
	v.SafeAcceptListStatement(list)
	statement.Statements.AcceptStatement(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitSynchronizedStatement(statement *SynchronizedStatement) {
	statement.Monitor.Accept(v)
	v.SafeAcceptStatement(statement.Statements)
}

func (v *AbstractJavaSyntaxVisitor) VisitThrowStatement(statement *ThrowStatement) {
	statement.Expression.Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitTryStatement(statement *TryStatement) {
	list := make([]IStatement, statement.Resources.Size())
	for i, element := range statement.Resources.ToSlice() {
		list[i] = element
	}
	v.SafeAcceptListStatement(list)
	statement.TryStatements.AcceptStatement(v)

	list = make([]IStatement, statement.CatchClause.Size())
	for i, element := range statement.CatchClause.ToSlice() {
		list[i] = element
	}
	v.SafeAcceptListStatement(list)
	v.SafeAcceptStatement(statement.FinallyStatements)
}

func (v *AbstractJavaSyntaxVisitor) VisitTryStatementResource(statement *Resource) {
	t := statement.Type
	t.AcceptTypeVisitor(v)
	statement.Expression.Accept(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitTryStatementCatchClause(statement *CatchClause) {
	t := statement.Type
	t.AcceptTypeVisitor(v)
	v.SafeAcceptStatement(statement.Statements)
}

func (v *AbstractJavaSyntaxVisitor) VisitTypeDeclarationStatement(statement *TypeDeclarationStatement) {
	statement.TypeDeclaration.AcceptDeclaration(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitWhileStatement(statement *WhileStatement) {
	statement.Condition.Accept(v)
	v.SafeAcceptStatement(statement.Statements)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////
// TypeVisitor
//////////////////////////////////////////////////////////////////////////////////////////////////////////////

/* 중복 메소드.
func (v *AbstractJavaSyntaxVisitor) VisitPrimitiveType(y *PrimitiveType)     {}
func (v *AbstractJavaSyntaxVisitor) VisitObjectType(y *ObjectType)           {}
func (v *AbstractJavaSyntaxVisitor) VisitInnerObjectType(y *InnerObjectType) {}
func (v *AbstractJavaSyntaxVisitor) VisitGenericType(y *GenericType)         {}
*/

func (v *AbstractJavaSyntaxVisitor) VisitTypes(types *Types) {
	iterator := types.Iterator()
	for iterator.HasNext() {
		t := iterator.Next()
		t.AcceptTypeVisitor(v)
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////
// TypeParameterVisitor
//////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (v *AbstractJavaSyntaxVisitor) VisitTypeParameter(_ *TypeParameter) {
	// EMPTY
}

func (v *AbstractJavaSyntaxVisitor) VisitTypeParameterWithTypeBounds(parameter *TypeParameterWithTypeBounds) {
	parameter.TypeBounds.AcceptTypeVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitTypeParameters(parameters *TypeParameters) {
	iterator := parameters.Iterator()
	for iterator.HasNext() {
		t := iterator.Next()
		t.AcceptTypeParameterVisitor(v)
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////
// TypeArgumentVisitor
//////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (v *AbstractJavaSyntaxVisitor) VisitTypeArguments(arguments *TypeArguments) {
	for _, typeArgument := range arguments.ToSlice() {
		typeArgument.AcceptTypeArgumentVisitor(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) VisitDiamondTypeArgument(_ *DiamondTypeArgument) {
	// EMPTY
}

func (v *AbstractJavaSyntaxVisitor) VisitWildcardExtendsTypeArgument(argument *WildcardExtendsTypeArgument) {
	argument.Type.(ITypeArgument).AcceptTypeArgumentVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitWildcardSuperTypeArgument(argument *WildcardSuperTypeArgument) {
	argument.Type.(ITypeArgument).AcceptTypeArgumentVisitor(v)
}

func (v *AbstractJavaSyntaxVisitor) VisitWildcardTypeArgument(_ *WildcardTypeArgument) {
	// EMPTY
}

func (v *AbstractJavaSyntaxVisitor) VisitPrimitiveType(_ *PrimitiveType) {
	// EMPTY
}

func (v *AbstractJavaSyntaxVisitor) VisitObjectType(t IType) {
	switch meta := t.(type) {
	case *ObjectType:
		v.SafeAcceptTypeArgumentVisitable(meta.TypeArguments)
	case *ObjectReference:
		v.SafeAcceptTypeArgumentVisitable(meta.TypeArguments)
	}
}

func (v *AbstractJavaSyntaxVisitor) VisitInnerObjectType(t IType) {
	t.GetOuterType().AcceptTypeArgumentVisitor(v)
	switch meta := t.(type) {
	case *InnerObjectType:
		v.SafeAcceptTypeArgumentVisitable(meta.TypeArguments)
	case *InnerObjectReference:
		v.SafeAcceptTypeArgumentVisitable(meta.TypeArguments)
	}
}

func (v *AbstractJavaSyntaxVisitor) VisitGenericType(_ *GenericType) {
	// EMPTY
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptTypeArgumentVisitable(visitable ITypeArgumentVisitable) {
	if visitable != nil {
		visitable.AcceptTypeArgumentVisitor(v)
	}
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Utility Methods
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (v *AbstractJavaSyntaxVisitor) VisitTypeDeclaration(decl ITypeDeclaration) {
	v.SafeAcceptReference(decl.GetAnnotationReferences())
}

func (v *AbstractJavaSyntaxVisitor) AcceptListDeclaration(list []IDeclaration) {
	for _, value := range list {
		value.AcceptDeclaration(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) AcceptListExpression(list []IExpression) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) AcceptListReference(list []IReference) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) AcceptListStatement(list []IStatement) {
	for _, value := range list {
		value.AcceptStatement(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptDeclaration(decl IDeclaration) {
	if decl != nil {
		decl.AcceptDeclaration(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptExpression(expr IExpression) {
	if expr != nil {
		expr.Accept(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptReference(ref IReference) {
	if ref != nil {
		ref.Accept(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptStatement(list IStatement) {
	if list != nil {
		list.AcceptStatement(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptType(list IType) {
	if list != nil {
		list.AcceptTypeVisitor(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptTypeParameter(list ITypeParameter) {
	if list != nil {
		list.AcceptTypeParameterVisitor(v)
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptListConstant(list []IConstant) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *AbstractJavaSyntaxVisitor) SafeAcceptListStatement(list []IStatement) {
	if list != nil {
		for _, value := range list {
			value.AcceptStatement(v)
		}
	}
}
