package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/expression"
)

func NewClassFileSuperConstructorInvocationExpression(lineNumber int, typ intmod.IObjectType, descriptor string,
	parameterTypes intmod.IType, parameters intmod.IExpression) intsrv.IClassFileSuperConstructorInvocationExpression {
	return &ClassFileSuperConstructorInvocationExpression{
		SuperConstructorInvocationExpression: *expression.NewSuperConstructorInvocationExpressionWithAll(
			lineNumber, typ, descriptor, parameters).(*expression.SuperConstructorInvocationExpression),
		parameterTypes: parameterTypes,
	}
}

type ClassFileSuperConstructorInvocationExpression struct {
	expression.SuperConstructorInvocationExpression

	parameterTypes intmod.IType
}

func (e *ClassFileSuperConstructorInvocationExpression) ParameterTypes() intmod.IType {
	return e.parameterTypes
}

func (e *ClassFileSuperConstructorInvocationExpression) SetParameterTypes(parameterTypes intmod.IType) {
	e.parameterTypes = parameterTypes
}
