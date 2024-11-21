package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewMethodInvocationExpression(typ intmod.IType, expression intmod.IExpression,
	internalTypeName, name, descriptor string) intmod.IMethodInvocationExpression {
	return &MethodInvocationExpression{
		MethodReferenceExpression: *NewMethodReferenceExpression(typ, expression,
			internalTypeName, name, descriptor).(*MethodReferenceExpression),
	}
}

func NewMethodInvocationExpressionWithLineNumber(lineNumber int, typ intmod.IType,
	expression intmod.IExpression, internalTypeName, name, descriptor string) intmod.IMethodInvocationExpression {
	return &MethodInvocationExpression{
		MethodReferenceExpression: *NewMethodReferenceExpressionWithAll(lineNumber,
			typ, expression, internalTypeName, name, descriptor).(*MethodReferenceExpression),
	}
}

func NewMethodInvocationExpressionWithParam(typ intmod.IType, expression intmod.IExpression,
	internalTypeName, name, descriptor string, parameters intmod.IExpression) intmod.IMethodInvocationExpression {
	return &MethodInvocationExpression{
		MethodReferenceExpression: *NewMethodReferenceExpression(typ, expression,
			internalTypeName, name, descriptor).(*MethodReferenceExpression),
		parameters: parameters,
	}
}

func NewMethodInvocationExpressionWithAll(lineNumber int, typ intmod.IType, expression intmod.IExpression,
	internalTypeName, name, descriptor string, parameters intmod.IExpression) intmod.IMethodInvocationExpression {
	return &MethodInvocationExpression{
		MethodReferenceExpression: *NewMethodReferenceExpressionWithAll(lineNumber, typ,
			expression, internalTypeName, name, descriptor).(*MethodReferenceExpression),
		parameters: parameters,
	}
}

type MethodInvocationExpression struct {
	MethodReferenceExpression

	nonWildcardTypeArguments intmod.ITypeArgument
	parameters               intmod.IExpression
}

func (e *MethodInvocationExpression) NonWildcardTypeArguments() intmod.ITypeArgument {
	return e.nonWildcardTypeArguments
}

func (e *MethodInvocationExpression) SetNonWildcardTypeArguments(arguments intmod.ITypeArgument) {
	e.nonWildcardTypeArguments = arguments
}

func (e *MethodInvocationExpression) Parameters() intmod.IExpression {
	return e.parameters
}

func (e *MethodInvocationExpression) SetParameters(params intmod.IExpression) {
	e.parameters = params
}

func (e *MethodInvocationExpression) Priority() int {
	return 1
}

func (e *MethodInvocationExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitMethodInvocationExpression(e)
}

func (e *MethodInvocationExpression) String() string {
	return fmt.Sprintf("MethodInvocationExpression{call %s . %s (%s)}", e.expression, e.name, e.descriptor)
}
