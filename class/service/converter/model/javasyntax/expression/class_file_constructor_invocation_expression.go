package expression

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/expression"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
)

func NewClassFileConstructorInvocationExpression(lineNumber int, typ _type.IObjectType, descriptor string,
	parameterTypes _type.IType, parameters expression.Expression) *ClassFileConstructorInvocationExpression {
	return &ClassFileConstructorInvocationExpression{
		ConstructorInvocationExpression: *expression.NewConstructorInvocationExpressionWithAll(
			lineNumber, typ, descriptor, parameters),
		parameterTypes: parameterTypes,
	}
}

type ClassFileConstructorInvocationExpression struct {
	expression.ConstructorInvocationExpression
	parameterTypes _type.IType
}

func (e *ClassFileConstructorInvocationExpression) ParameterTypes() _type.IType {
	return e.parameterTypes
}

func (e *ClassFileConstructorInvocationExpression) SetParameterTypes(parameterTypes _type.IType) {
	e.parameterTypes = parameterTypes
}
