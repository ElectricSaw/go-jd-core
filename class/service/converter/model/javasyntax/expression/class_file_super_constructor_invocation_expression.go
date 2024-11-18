package expression

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/expression"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
)

func NewClassFileSuperConstructorInvocationExpression(lineNumber int, typ _type.IObjectType, descriptor string,
	parameterTypes _type.IType, parameters expression.IExpression) *ClassFileSuperConstructorInvocationExpression {
	return &ClassFileSuperConstructorInvocationExpression{
		SuperConstructorInvocationExpression: *expression.NewSuperConstructorInvocationExpressionWithAll(lineNumber, typ, descriptor, parameters),
		parameterTypes:                       parameterTypes,
	}
}

type ClassFileSuperConstructorInvocationExpression struct {
	expression.SuperConstructorInvocationExpression

	parameterTypes _type.IType
}

func (e *ClassFileSuperConstructorInvocationExpression) ParameterTypes() _type.IType {
	return e.parameterTypes
}

func (e *ClassFileSuperConstructorInvocationExpression) SetParameterTypes(parameterTypes _type.IType) {
	e.parameterTypes = parameterTypes
}
