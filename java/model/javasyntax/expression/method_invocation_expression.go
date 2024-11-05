package expression

import (
	_type "bitbucket.org/coontec/javaClass/java/model/javasyntax/type"
	"fmt"
)

func NewMethodInvocationExpression(typ _type.IType, expression Expression, internalTypeName, name, descriptor string) *MethodInvocationExpression {
	return &MethodInvocationExpression{
		MethodReferenceExpression: *NewMethodReferenceExpression(typ, expression, internalTypeName, name, descriptor),
	}
}

func NewMethodInvocationExpressionWithLineNumber(lineNumber int, typ _type.IType, expression Expression, internalTypeName, name, descriptor string) *MethodInvocationExpression {
	return &MethodInvocationExpression{
		MethodReferenceExpression: *NewMethodReferenceExpressionWithAll(lineNumber, typ, expression, internalTypeName, name, descriptor),
	}
}

func NewMethodInvocationExpressionWithParam(typ _type.IType, expression Expression, internalTypeName, name, descriptor string, parameters Expression) *MethodInvocationExpression {
	return &MethodInvocationExpression{
		MethodReferenceExpression: *NewMethodReferenceExpression(typ, expression, internalTypeName, name, descriptor),
		parameters:                parameters,
	}
}

func NewMethodInvocationExpressionWithAll(lineNumber int, typ _type.IType, expression Expression, internalTypeName, name, descriptor string, parameters Expression) *MethodInvocationExpression {
	return &MethodInvocationExpression{
		MethodReferenceExpression: *NewMethodReferenceExpressionWithAll(lineNumber, typ, expression, internalTypeName, name, descriptor),
		parameters:                parameters,
	}
}

type MethodInvocationExpression struct {
	MethodReferenceExpression

	nonWildcardTypeArguments _type.ITypeArgument
	parameters               Expression
}

func (e *MethodInvocationExpression) GetNonWildcardTypeArguments() _type.ITypeArgument {
	return e.nonWildcardTypeArguments
}

func (e *MethodInvocationExpression) SetNonWildcardTypeArguments(arguments _type.ITypeArgument) {
	e.nonWildcardTypeArguments = arguments
}

func (e *MethodInvocationExpression) GetParameters() Expression {
	return e.parameters
}

func (e *MethodInvocationExpression) SetParameters(params Expression) {
	e.parameters = params
}

func (e *MethodInvocationExpression) GetPriority() int {
	return 1
}

func (e *MethodInvocationExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitMethodInvocationExpression(e)
}

func (e *MethodInvocationExpression) String() string {
	return fmt.Sprintf("MethodInvocationExpression{call %s . %s (%s)}", e.expression, e.name, e.descriptor)
}
