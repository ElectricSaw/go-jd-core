package expression

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	"fmt"
)

func NewMethodInvocationExpression(typ intsyn.IType, expression intsyn.IExpression,
	internalTypeName, name, descriptor string) intsyn.IMethodInvocationExpression {
	return &MethodInvocationExpression{
		MethodReferenceExpression: *NewMethodReferenceExpression(typ, expression,
			internalTypeName, name, descriptor).(*MethodReferenceExpression),
	}
}

func NewMethodInvocationExpressionWithLineNumber(lineNumber int, typ intsyn.IType,
	expression intsyn.IExpression, internalTypeName, name, descriptor string) intsyn.IMethodInvocationExpression {
	return &MethodInvocationExpression{
		MethodReferenceExpression: *NewMethodReferenceExpressionWithAll(lineNumber,
			typ, expression, internalTypeName, name, descriptor).(*MethodReferenceExpression),
	}
}

func NewMethodInvocationExpressionWithParam(typ intsyn.IType, expression intsyn.IExpression,
	internalTypeName, name, descriptor string, parameters intsyn.IExpression) intsyn.IMethodInvocationExpression {
	return &MethodInvocationExpression{
		MethodReferenceExpression: *NewMethodReferenceExpression(typ, expression,
			internalTypeName, name, descriptor).(*MethodReferenceExpression),
		parameters: parameters,
	}
}

func NewMethodInvocationExpressionWithAll(lineNumber int, typ intsyn.IType, expression intsyn.IExpression,
	internalTypeName, name, descriptor string, parameters intsyn.IExpression) intsyn.IMethodInvocationExpression {
	return &MethodInvocationExpression{
		MethodReferenceExpression: *NewMethodReferenceExpressionWithAll(lineNumber, typ,
			expression, internalTypeName, name, descriptor).(*MethodReferenceExpression),
		parameters: parameters,
	}
}

type MethodInvocationExpression struct {
	MethodReferenceExpression

	nonWildcardTypeArguments intsyn.ITypeArgument
	parameters               intsyn.IExpression
}

func (e *MethodInvocationExpression) NonWildcardTypeArguments() intsyn.ITypeArgument {
	return e.nonWildcardTypeArguments
}

func (e *MethodInvocationExpression) SetNonWildcardTypeArguments(arguments intsyn.ITypeArgument) {
	e.nonWildcardTypeArguments = arguments
}

func (e *MethodInvocationExpression) Parameters() intsyn.IExpression {
	return e.parameters
}

func (e *MethodInvocationExpression) SetParameters(params intsyn.IExpression) {
	e.parameters = params
}

func (e *MethodInvocationExpression) Priority() int {
	return 1
}

func (e *MethodInvocationExpression) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitMethodInvocationExpression(e)
}

func (e *MethodInvocationExpression) String() string {
	return fmt.Sprintf("MethodInvocationExpression{call %s . %s (%s)}", e.expression, e.name, e.descriptor)
}
