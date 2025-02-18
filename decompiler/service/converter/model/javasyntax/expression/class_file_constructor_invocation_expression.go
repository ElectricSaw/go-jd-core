package expression

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
)

func NewClassFileConstructorInvocationExpression(lineNumber int, typ intmod.IObjectType, descriptor string,
	parameterTypes intmod.IType, parameters intmod.IExpression) intsrv.IClassFileConstructorInvocationExpression {
	e := &ClassFileConstructorInvocationExpression{
		ConstructorInvocationExpression: *model.NewConstructorInvocationExpressionWithAll(
			lineNumber, typ, descriptor, parameters).(*model.ConstructorInvocationExpression),
		parameterTypes: parameterTypes,
	}
	e.SetValue(e)
	return e
}

type ClassFileConstructorInvocationExpression struct {
	model.ConstructorInvocationExpression
	parameterTypes intmod.IType
}

func (e *ClassFileConstructorInvocationExpression) ParameterTypes() intmod.IType {
	return e.parameterTypes
}

func (e *ClassFileConstructorInvocationExpression) SetParameterTypes(parameterTypes intmod.IType) {
	e.parameterTypes = parameterTypes
}
