package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/expression"
)

func NewClassFileConstructorInvocationExpression(lineNumber int, typ intmod.IObjectType, descriptor string,
	parameterTypes intmod.IType, parameters intmod.IExpression) intsrv.IClassFileConstructorInvocationExpression {
	return &ClassFileConstructorInvocationExpression{
		ConstructorInvocationExpression: *expression.NewConstructorInvocationExpressionWithAll(
			lineNumber, typ, descriptor, parameters).(*expression.ConstructorInvocationExpression),
		parameterTypes: parameterTypes,
	}
}

type ClassFileConstructorInvocationExpression struct {
	expression.ConstructorInvocationExpression
	parameterTypes intmod.IType
}

func (e *ClassFileConstructorInvocationExpression) ParameterTypes() intmod.IType {
	return e.parameterTypes
}

func (e *ClassFileConstructorInvocationExpression) SetParameterTypes(parameterTypes intmod.IType) {
	e.parameterTypes = parameterTypes
}
