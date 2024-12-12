package expression

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax/expression"
)

func NewClassFileMethodInvocationExpression(lineNumber int, typeParameters intmod.ITypeParameter, typ intmod.IType,
	expr intmod.IExpression, internalTypeName, name, descriptor string, parameterTypes intmod.IType,
	parameters intmod.IExpression) intsrv.IClassFileMethodInvocationExpression {
	e := &ClassFileMethodInvocationExpression{
		MethodInvocationExpression: *expression.NewMethodInvocationExpressionWithAll(lineNumber,
			typ, expr, internalTypeName, name, descriptor, parameters).(*expression.MethodInvocationExpression),
		typeParameters: typeParameters,
		parameterTypes: parameterTypes,
		bound:          false,
	}
	e.SetValue(e)
	return e
}

type ClassFileMethodInvocationExpression struct {
	expression.MethodInvocationExpression

	typeParameters intmod.ITypeParameter
	parameterTypes intmod.IType
	bound          bool
}

func (e *ClassFileMethodInvocationExpression) TypeParameters() intmod.ITypeParameter {
	return e.typeParameters
}

func (e *ClassFileMethodInvocationExpression) ParameterTypes() intmod.IType {
	return e.parameterTypes
}

func (e *ClassFileMethodInvocationExpression) SetParameterTypes(parameterTypes intmod.IType) {
	e.parameterTypes = parameterTypes
}

func (e *ClassFileMethodInvocationExpression) IsBound() bool {
	return e.bound
}

func (e *ClassFileMethodInvocationExpression) SetBound(bound bool) {
	e.bound = bound
}
