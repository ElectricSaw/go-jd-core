package expression

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/expression"
)

func NewClassFileConstructorInvocationExpression(lineNumber int, typ intmod.IObjectType, descriptor string,
	parameterTypes intmod.IType, parameters intmod.IExpression) intsrv.IClassFileConstructorInvocationExpression {
	e := &ClassFileConstructorInvocationExpression{
		ConstructorInvocationExpression: *expression.NewConstructorInvocationExpressionWithAll(
			lineNumber, typ, descriptor, parameters).(*expression.ConstructorInvocationExpression),
		parameterTypes: parameterTypes,
	}
	e.SetValue(e)
	return e
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
