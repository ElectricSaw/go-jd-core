package javasyntax

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/declaration"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/expression"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/reference"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/statement"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
)

type AbstractJavaSyntaxVisitor struct {
	declaration.DeclarationVisitor
	expression.ExpressionVisitor
	reference.ReferenceVisitor
	statement.StatementVisitor
	_type.TypeVisitor
	_type.TypeParameterVisitor
	_type.AbstractTypeArgumentVisitor
}

// --- DeclarationVisitor ---
func (v *AbstractJavaSyntaxVisitor) VisitAnnotationDeclaration(declaration *declaration.AnnotationDeclaration) {
}
func (v *AbstractJavaSyntaxVisitor) VisitArrayVariableInitializer(declaration *declaration.ArrayVariableInitializer) {
}
func (v *AbstractJavaSyntaxVisitor) VisitBodyDeclaration(declaration *declaration.BodyDeclaration) {}
func (v *AbstractJavaSyntaxVisitor) VisitClassDeclaration(declaration *declaration.ClassDeclaration) {
}
func (v *AbstractJavaSyntaxVisitor) VisitConstructorDeclaration(declaration *declaration.ConstructorDeclaration) {
}
func (v *AbstractJavaSyntaxVisitor) VisitEnumDeclaration(declaration *declaration.EnumDeclaration)  {}
func (v *AbstractJavaSyntaxVisitor) VisitEnumDeclarationConstant(declaration *declaration.Constant) {}
func (v *AbstractJavaSyntaxVisitor) VisitExpressionVariableInitializer(declaration *declaration.ExpressionVariableInitializer) {
}
func (v *AbstractJavaSyntaxVisitor) VisitFieldDeclaration(declaration *declaration.FieldDeclaration) {
}
func (v *AbstractJavaSyntaxVisitor) VisitFieldDeclarator(declaration *declaration.FieldDeclarator) {}
func (v *AbstractJavaSyntaxVisitor) VisitFieldDeclarators(declarations *declaration.FieldDeclarators) {
}
func (v *AbstractJavaSyntaxVisitor) VisitFormalParameter(declaration *declaration.FormalParameter) {}
func (v *AbstractJavaSyntaxVisitor) VisitFormalParameters(declarations *declaration.FormalParameters) {
}
func (v *AbstractJavaSyntaxVisitor) VisitInstanceInitializerDeclaration(declaration *declaration.InstanceInitializerDeclaration) {
}
func (v *AbstractJavaSyntaxVisitor) VisitInterfaceDeclaration(declaration *declaration.InterfaceDeclaration) {
}
func (v *AbstractJavaSyntaxVisitor) VisitLocalVariableDeclaration(declaration *declaration.LocalVariableDeclaration) {
}
func (v *AbstractJavaSyntaxVisitor) VisitLocalVariableDeclarator(declarator *declaration.LocalVariableDeclarator) {
}
func (v *AbstractJavaSyntaxVisitor) VisitLocalVariableDeclarators(declarators *declaration.LocalVariableDeclarators) {
}
func (v *AbstractJavaSyntaxVisitor) VisitMethodDeclaration(declaration *declaration.MethodDeclaration) {
}
func (v *AbstractJavaSyntaxVisitor) VisitMemberDeclarations(declarations *declaration.MemberDeclarations) {
}
func (v *AbstractJavaSyntaxVisitor) VisitModuleDeclaration(declarations *declaration.ModuleDeclaration) {
}
func (v *AbstractJavaSyntaxVisitor) VisitStaticInitializerDeclaration(declaration *declaration.StaticInitializerDeclaration) {
}
func (v *AbstractJavaSyntaxVisitor) VisitTypeDeclarations(declarations *declaration.TypeDeclarations) {
}

// --- ExpressionVisitor ---
func (v *AbstractJavaSyntaxVisitor) VisitArrayExpression(expression *expression.ArrayExpression) {}
func (v *AbstractJavaSyntaxVisitor) VisitBinaryOperatorExpression(expression *expression.BinaryOperatorExpression) {
}
func (v *AbstractJavaSyntaxVisitor) VisitBooleanExpression(expression *expression.BooleanExpression) {
}
func (v *AbstractJavaSyntaxVisitor) VisitCastExpression(expression *expression.CastExpression) {}
func (v *AbstractJavaSyntaxVisitor) VisitCommentExpression(expression *expression.CommentExpression) {
}
func (v *AbstractJavaSyntaxVisitor) VisitConstructorInvocationExpression(expression *expression.ConstructorInvocationExpression) {
}
func (v *AbstractJavaSyntaxVisitor) VisitConstructorReferenceExpression(expression *expression.ConstructorReferenceExpression) {
}
func (v *AbstractJavaSyntaxVisitor) VisitDoubleConstantExpression(expression *expression.DoubleConstantExpression) {
}
func (v *AbstractJavaSyntaxVisitor) VisitEnumConstantReferenceExpression(expression *expression.EnumConstantReferenceExpression) {
}
func (v *AbstractJavaSyntaxVisitor) VisitExpressions(expression *expression.Expressions) {}
func (v *AbstractJavaSyntaxVisitor) VisitFieldReferenceExpression(expression *expression.FieldReferenceExpression) {
}
func (v *AbstractJavaSyntaxVisitor) VisitFloatConstantExpression(expression *expression.FloatConstantExpression) {
}
func (v *AbstractJavaSyntaxVisitor) VisitIntegerConstantExpression(expression *expression.IntegerConstantExpression) {
}
func (v *AbstractJavaSyntaxVisitor) VisitInstanceOfExpression(expression *expression.InstanceOfExpression) {
}
func (v *AbstractJavaSyntaxVisitor) VisitLambdaFormalParametersExpression(expression *expression.LambdaFormalParametersExpression) {
}
func (v *AbstractJavaSyntaxVisitor) VisitLambdaIdentifiersExpression(expression *expression.LambdaIdentifiersExpression) {
}
func (v *AbstractJavaSyntaxVisitor) VisitLengthExpression(expression *expression.LengthExpression) {}
func (v *AbstractJavaSyntaxVisitor) VisitLocalVariableReferenceExpression(expression *expression.LocalVariableReferenceExpression) {
}
func (v *AbstractJavaSyntaxVisitor) VisitLongConstantExpression(expression *expression.LongConstantExpression) {
}
func (v *AbstractJavaSyntaxVisitor) VisitMethodInvocationExpression(expression *expression.MethodInvocationExpression) {
}
func (v *AbstractJavaSyntaxVisitor) VisitMethodReferenceExpression(expression *expression.MethodReferenceExpression) {
}
func (v *AbstractJavaSyntaxVisitor) VisitNewArray(expression *expression.NewArray)           {}
func (v *AbstractJavaSyntaxVisitor) VisitNewExpression(expression *expression.NewExpression) {}
func (v *AbstractJavaSyntaxVisitor) VisitNewInitializedArray(expression *expression.NewInitializedArray) {
}
func (v *AbstractJavaSyntaxVisitor) VisitNoExpression(expression *expression.NoExpression)     {}
func (v *AbstractJavaSyntaxVisitor) VisitNullExpression(expression *expression.NullExpression) {}
func (v *AbstractJavaSyntaxVisitor) VisitObjectTypeReferenceExpression(expression *expression.ObjectTypeReferenceExpression) {
}
func (v *AbstractJavaSyntaxVisitor) VisitParenthesesExpression(expression *expression.ParenthesesExpression) {
}
func (v *AbstractJavaSyntaxVisitor) VisitPostOperatorExpression(expression *expression.PostOperatorExpression) {
}
func (v *AbstractJavaSyntaxVisitor) VisitPreOperatorExpression(expression *expression.PreOperatorExpression) {
}
func (v *AbstractJavaSyntaxVisitor) VisitStringConstantExpression(expression *expression.StringConstantExpression) {
}
func (v *AbstractJavaSyntaxVisitor) VisitSuperConstructorInvocationExpression(expression *expression.SuperConstructorInvocationExpression) {
}
func (v *AbstractJavaSyntaxVisitor) VisitSuperExpression(expression *expression.SuperExpression) {}
func (v *AbstractJavaSyntaxVisitor) VisitTernaryOperatorExpression(expression *expression.TernaryOperatorExpression) {
}
func (v *AbstractJavaSyntaxVisitor) VisitThisExpression(expression *expression.ThisExpression) {}
func (v *AbstractJavaSyntaxVisitor) VisitTypeReferenceDotClassExpression(expression *expression.TypeReferenceDotClassExpression) {
}

// --- ReferenceVisitor ---
func (v *AbstractJavaSyntaxVisitor) VisitAnnotationElementValue(reference *reference.AnnotationElementValue) {
}
func (v *AbstractJavaSyntaxVisitor) VisitAnnotationReference(reference *reference.AnnotationReference) {
}
func (v *AbstractJavaSyntaxVisitor) VisitAnnotationReferences(references *reference.AnnotationReferences) {
}
func (v *AbstractJavaSyntaxVisitor) VisitElementValueArrayInitializerElementValue(reference *reference.ElementValueArrayInitializerElementValue) {
}
func (v *AbstractJavaSyntaxVisitor) VisitElementValues(references *reference.ElementValues)         {}
func (v *AbstractJavaSyntaxVisitor) VisitElementValuePair(reference *reference.ElementValuePair)    {}
func (v *AbstractJavaSyntaxVisitor) VisitElementValuePairs(references *reference.ElementValuePairs) {}
func (v *AbstractJavaSyntaxVisitor) VisitExpressionElementValue(reference *reference.ExpressionElementValue) {
}
func (v *AbstractJavaSyntaxVisitor) VisitInnerObjectReference(reference *reference.InnerObjectReference) {
}
func (v *AbstractJavaSyntaxVisitor) VisitObjectReference(reference *reference.ObjectReference) {}

// --- StatementVisitor ---
func (v *AbstractJavaSyntaxVisitor) VisitAssertStatement(statement *statement.AssertStatement)     {}
func (v *AbstractJavaSyntaxVisitor) VisitBreakStatement(statement *statement.BreakStatement)       {}
func (v *AbstractJavaSyntaxVisitor) VisitByteCodeStatement(statement *statement.ByteCodeStatement) {}
func (v *AbstractJavaSyntaxVisitor) VisitCommentStatement(statement *statement.CommentStatement)   {}
func (v *AbstractJavaSyntaxVisitor) VisitContinueStatement(statement *statement.ContinueStatement) {}
func (v *AbstractJavaSyntaxVisitor) VisitDoWhileStatement(statement *statement.DoWhileStatement)   {}
func (v *AbstractJavaSyntaxVisitor) VisitExpressionStatement(statement *statement.ExpressionStatement) {
}
func (v *AbstractJavaSyntaxVisitor) VisitForEachStatement(statement *statement.ForEachStatement) {}
func (v *AbstractJavaSyntaxVisitor) VisitForStatement(statement *statement.ForStatement)         {}
func (v *AbstractJavaSyntaxVisitor) VisitIfStatement(statement *statement.IfStatement)           {}
func (v *AbstractJavaSyntaxVisitor) VisitIfElseStatement(statement *statement.IfElseStatement)   {}
func (v *AbstractJavaSyntaxVisitor) VisitLabelStatement(statement *statement.LabelStatement)     {}
func (v *AbstractJavaSyntaxVisitor) VisitLambdaExpressionStatement(statement *statement.LambdaExpressionStatement) {
}
func (v *AbstractJavaSyntaxVisitor) VisitLocalVariableDeclarationStatement(statement *statement.LocalVariableDeclarationStatement) {
}
func (v *AbstractJavaSyntaxVisitor) VisitNoStatement(statement *statement.NoStatement) {}
func (v *AbstractJavaSyntaxVisitor) VisitReturnExpressionStatement(statement *statement.ReturnExpressionStatement) {
}
func (v *AbstractJavaSyntaxVisitor) VisitReturnStatement(statement *statement.ReturnStatement) {}
func (v *AbstractJavaSyntaxVisitor) VisitStatements(statement *statement.Statements)           {}
func (v *AbstractJavaSyntaxVisitor) VisitSwitchStatement(statement *statement.SwitchStatement) {}
func (v *AbstractJavaSyntaxVisitor) VisitSwitchStatementDefaultLabel(statement *statement.DefaultLabe1) {
}
func (v *AbstractJavaSyntaxVisitor) VisitSwitchStatementExpressionLabel(statement *statement.ExpressionLabel) {
}
func (v *AbstractJavaSyntaxVisitor) VisitSwitchStatementLabelBlock(statement *statement.LabelBlock) {}
func (v *AbstractJavaSyntaxVisitor) VisitSwitchStatementMultiLabelsBlock(statement *statement.MultiLabelsBlock) {
}
func (v *AbstractJavaSyntaxVisitor) VisitSynchronizedStatement(statement *statement.SynchronizedStatement) {
}
func (v *AbstractJavaSyntaxVisitor) VisitThrowStatement(statement *statement.ThrowStatement)       {}
func (v *AbstractJavaSyntaxVisitor) VisitTryStatement(statement *statement.TryStatement)           {}
func (v *AbstractJavaSyntaxVisitor) VisitTryStatementResource(statement *statement.Resource)       {}
func (v *AbstractJavaSyntaxVisitor) VisitTryStatementCatchClause(statement *statement.CatchClause) {}
func (v *AbstractJavaSyntaxVisitor) VisitTypeDeclarationStatement(statement *statement.TypeDeclarationStatement) {
}
func (v *AbstractJavaSyntaxVisitor) VisitWhileStatement(statement *statement.WhileStatement) {}

// --- TypeVisitor ---
// func (v *AbstractJavaSyntaxVisitor) VisitPrimitiveType(y *PrimitiveType)     {}
// func (v *AbstractJavaSyntaxVisitor) VisitObjectType(y *ObjectType)           {}
// func (v *AbstractJavaSyntaxVisitor) VisitInnerObjectType(y *InnerObjectType) {}
func (v *AbstractJavaSyntaxVisitor) VisitTypes(types *_type.Types) {}

//func (v *AbstractJavaSyntaxVisitor) VisitGenericType(y *GenericType)         {}

// --- TypeParameterVisitor --- //
func (v *AbstractJavaSyntaxVisitor) VisitTypeParameter(parameter *_type.TypeParameter) {}
func (v *AbstractJavaSyntaxVisitor) VisitTypeParameterWithTypeBounds(parameter *_type.TypeParameterWithTypeBounds) {
}
func (v *AbstractJavaSyntaxVisitor) VisitTypeParameters(parameters *_type.TypeParameters) {}

// --- TypeArgumentVisitor ---
func (v *AbstractJavaSyntaxVisitor) VisitTypeArguments(arguments *_type.TypeArguments)            {}
func (v *AbstractJavaSyntaxVisitor) VisitDiamondTypeArgument(argument *_type.DiamondTypeArgument) {}
func (v *AbstractJavaSyntaxVisitor) VisitWildcardExtendsTypeArgument(argument *_type.WildcardExtendsTypeArgument) {
}
func (v *AbstractJavaSyntaxVisitor) VisitWildcardSuperTypeArgument(argument *_type.WildcardSuperTypeArgument) {
}
func (v *AbstractJavaSyntaxVisitor) VisitWildcardTypeArgument(argument *_type.WildcardTypeArgument) {}
func (v *AbstractJavaSyntaxVisitor) VisitPrimitiveType(t *_type.PrimitiveType)                      {}
func (v *AbstractJavaSyntaxVisitor) VisitObjectType(t *_type.ObjectType)                            {}
func (v *AbstractJavaSyntaxVisitor) VisitInnerObjectType(t *_type.InnerObjectType)                  {}
func (v *AbstractJavaSyntaxVisitor) VisitGenericType(t *_type.GenericType)                          {}
