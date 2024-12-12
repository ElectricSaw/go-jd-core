package model

type IJavaSyntaxVisitor interface {
	IDeclarationVisitor
	IExpressionVisitor
	IReferenceVisitor
	IStatementVisitor
	ITypeVisitor
	ITypeParameterVisitor
	ITypeArgumentVisitor

	VisitTypeDeclaration(decl ITypeDeclaration)
	AcceptListDeclaration(list []IDeclaration)
	AcceptListExpression(list []IExpression)
	AcceptListReference(list []IReference)
	AcceptListStatement(list []IStatement)
	SafeAcceptDeclaration(decl IDeclaration)
	SafeAcceptExpression(expr IExpression)
	SafeAcceptReference(ref IReference)
	SafeAcceptStatement(list IStatement)
	SafeAcceptType(list IType)
	SafeAcceptTypeParameter(list ITypeParameter)
	SafeAcceptListDeclaration(list []IDeclaration)
	SafeAcceptListConstant(list []IConstant)
	SafeAcceptListStatement(list []IStatement)
}

type ICompilationUnit interface {
	TypeDeclarations() ITypeDeclaration
}
