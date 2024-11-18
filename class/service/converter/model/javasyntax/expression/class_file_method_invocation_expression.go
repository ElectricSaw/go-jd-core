package expression

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/expression"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
)

func NewClassFileMethodInvocationExpression(lineNumber int, typeParameters _type.ITypeParameter, typ _type.IType,
	expr expression.IExpression, internalTypeName, name, descriptor string, parameterTypes _type.IType,
	parameters expression.IExpression) *ClassFileMethodInvocationExpression {
	return &ClassFileMethodInvocationExpression{
		MethodInvocationExpression: *expression.NewMethodInvocationExpressionWithAll(lineNumber, typ, expr, internalTypeName, name, descriptor, parameters),
		typeParameters:             typeParameters,
		parameterTypes:             parameterTypes,
		bound:                      false,
	}
}

type ClassFileMethodInvocationExpression struct {
	expression.MethodInvocationExpression

	typeParameters _type.ITypeParameter
	parameterTypes _type.IType
	bound          bool
}

func (e *ClassFileMethodInvocationExpression) TypeParameters() _type.ITypeParameter {
	return e.typeParameters
}

func (e *ClassFileMethodInvocationExpression) ParameterTypes() _type.IType {
	return e.parameterTypes
}

func (e *ClassFileMethodInvocationExpression) SetParameterTypes(parameterTypes _type.IType) {
	e.parameterTypes = parameterTypes
}

func (e *ClassFileMethodInvocationExpression) IsBound() bool {
	return e.bound
}

func (e *ClassFileMethodInvocationExpression) SetBound(bound bool) {
	e.bound = bound
}
