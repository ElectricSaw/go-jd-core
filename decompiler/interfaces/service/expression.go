package service

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

type IClassFileCmpExpression interface {
	intmod.IBinaryOperatorExpression
}

type IClassFileConstructorInvocationExpression interface {
	intmod.IConstructorInvocationExpression

	ParameterTypes() intmod.IType
	SetParameterTypes(parameterTypes intmod.IType)
}

type IClassFileLocalVariableReferenceExpression interface {
	intmod.ILocalVariableReferenceExpression

	Offset() int
	LocalVariable() ILocalVariableReference
	SetLocalVariable(localVariable ILocalVariableReference)
}

type IClassFileMethodInvocationExpression interface {
	intmod.IMethodInvocationExpression

	TypeParameters() intmod.ITypeParameter
	ParameterTypes() intmod.IType
	SetParameterTypes(parameterTypes intmod.IType)
	IsBound() bool
	SetBound(bound bool)
}

type IClassFileNewExpression interface {
	intmod.INewExpression

	ParameterTypes() intmod.IType
	SetParameterTypes(parameterTypes intmod.IType)
	IsBound() bool
	SetBound(bound bool)
	Set(descriptor string, parameterTypes intmod.IType, parameters intmod.IExpression)
	String() string
}

type IClassFileSuperConstructorInvocationExpression interface {
	intmod.ISuperConstructorInvocationExpression

	ParameterTypes() intmod.IType
	SetParameterTypes(parameterTypes intmod.IType)
}
