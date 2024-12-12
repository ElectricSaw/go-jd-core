package utils

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	_type "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/type"
)

func NewAbstractTypeParametersToTypeArgumentsBinder() intsrv.ITypeParametersToTypeArgumentsBinder {
	return &AbstractTypeParametersToTypeArgumentsBinder{}
}

type AbstractTypeParametersToTypeArgumentsBinder struct {
}

func (b *AbstractTypeParametersToTypeArgumentsBinder) NewConstructorInvocationExpression(
	lineNumber int, objectType intmod.IObjectType, descriptor string,
	methodTypes intsrv.IMethodTypes, parameters intmod.IExpression) intsrv.IClassFileConstructorInvocationExpression {
	return nil
}

func (b *AbstractTypeParametersToTypeArgumentsBinder) NewSuperConstructorInvocationExpression(
	lineNumber int, objectType intmod.IObjectType, descriptor string,
	methodTypes intsrv.IMethodTypes, parameters intmod.IExpression) intsrv.IClassFileSuperConstructorInvocationExpression {
	return nil
}

func (b *AbstractTypeParametersToTypeArgumentsBinder) NewMethodInvocationExpression(
	lineNumber int, expression intmod.IExpression, objectType intmod.IObjectType, name, descriptor string,
	methodTypes intsrv.IMethodTypes, parameters intmod.IExpression) intsrv.IClassFileMethodInvocationExpression {
	return nil
}

func (b *AbstractTypeParametersToTypeArgumentsBinder) NewFieldReferenceExpression(
	lineNumber int, typ intmod.IType, expression intmod.IExpression,
	objectType intmod.IObjectType, name, descriptor string) intmod.IFieldReferenceExpression {
	return nil
}

func (b *AbstractTypeParametersToTypeArgumentsBinder) BindParameterTypesWithArgumentTypes(
	typ intmod.IType, expression intmod.IExpression) {
}

func (b *AbstractTypeParametersToTypeArgumentsBinder) UpdateNewExpression(ne intsrv.IClassFileNewExpression,
	descriptor string, methodTypes intsrv.IMethodTypes, parameters intmod.IExpression) {
	ne.Set(descriptor, Clone(methodTypes.ParameterTypes()), parameters)
}

func Clone(parameterTypes intmod.IType) intmod.IType {
	if (parameterTypes != nil) && parameterTypes.IsList() {
		switch parameterTypes.Size() {
		case 0:
			parameterTypes = nil
			break
		case 1:
			parameterTypes = parameterTypes.First()
			break
		default:
			parameterTypes = _type.NewTypesWithSlice(parameterTypes.ToSlice())
			break
		}
	}

	return parameterTypes
}
