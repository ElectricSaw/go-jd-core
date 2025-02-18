package expression

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
)

func NewClassFileSuperConstructorInvocationExpression(lineNumber int, typ intmod.IObjectType, descriptor string,
	parameterTypes intmod.IType, parameters intmod.IExpression) intsrv.IClassFileSuperConstructorInvocationExpression {
	e := &ClassFileSuperConstructorInvocationExpression{
		SuperConstructorInvocationExpression: *model.NewSuperConstructorInvocationExpressionWithAll(
			lineNumber, typ, descriptor, parameters).(*model.SuperConstructorInvocationExpression),
		parameterTypes: parameterTypes,
	}
	e.SetValue(e)
	return e
}

type ClassFileSuperConstructorInvocationExpression struct {
	model.SuperConstructorInvocationExpression

	parameterTypes intmod.IType
}

func (e *ClassFileSuperConstructorInvocationExpression) ParameterTypes() intmod.IType {
	return e.parameterTypes
}

func (e *ClassFileSuperConstructorInvocationExpression) SetParameterTypes(parameterTypes intmod.IType) {
	e.parameterTypes = parameterTypes
}
