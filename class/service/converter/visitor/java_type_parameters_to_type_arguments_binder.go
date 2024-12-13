package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	modexpr "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/expression"
	srvexpr "github.com/ElectricSaw/go-jd-core/class/service/converter/model/javasyntax/expression"
)

func NewJavaTypeParametersToTypeArgumentsBinder() intsrv.ITypeParametersToTypeArgumentsBinder {
	return &JavaTypeParametersToTypeArgumentsBinder{}
}

type JavaTypeParametersToTypeArgumentsBinder struct {
	AbstractTypeParametersToTypeArgumentsBinder
}

func (b *JavaTypeParametersToTypeArgumentsBinder) NewConstructorInvocationExpression(
	lineNumber int, objectType intmod.IObjectType, descriptor string,
	methodTypes intsrv.IMethodTypes, parameters intmod.IExpression) intsrv.IClassFileConstructorInvocationExpression {
	return srvexpr.NewClassFileConstructorInvocationExpression(lineNumber, objectType,
		descriptor, Clone(methodTypes.ParameterTypes()), parameters)
}

func (b *JavaTypeParametersToTypeArgumentsBinder) NewSuperConstructorInvocationExpression(
	lineNumber int, objectType intmod.IObjectType, descriptor string,
	methodTypes intsrv.IMethodTypes, parameters intmod.IExpression) intsrv.IClassFileSuperConstructorInvocationExpression {
	return srvexpr.NewClassFileSuperConstructorInvocationExpression(lineNumber, objectType,
		descriptor, Clone(methodTypes.ParameterTypes()), parameters)
}

func (b *JavaTypeParametersToTypeArgumentsBinder) NewMethodInvocationExpression(
	lineNumber int, expr intmod.IExpression, objectType intmod.IObjectType, name, descriptor string,
	methodTypes intsrv.IMethodTypes, parameters intmod.IExpression) intsrv.IClassFileMethodInvocationExpression {
	return srvexpr.NewClassFileMethodInvocationExpression(lineNumber, methodTypes.TypeParameters(),
		methodTypes.ReturnedType(), expr, objectType.InternalName(), name, descriptor,
		Clone(methodTypes.ParameterTypes()), parameters)
}

func (b *JavaTypeParametersToTypeArgumentsBinder) NewFieldReferenceExpression(
	lineNumber int, typ intmod.IType, expr intmod.IExpression,
	objectType intmod.IObjectType, name, descriptor string) intmod.IFieldReferenceExpression {
	return modexpr.NewFieldReferenceExpressionWithAll(
		lineNumber, typ, expr, objectType.InternalName(), name, descriptor)
}

func (b *JavaTypeParametersToTypeArgumentsBinder) BindParameterTypesWithArgumentTypes(
	typ intmod.IType, expression intmod.IExpression) {
}
