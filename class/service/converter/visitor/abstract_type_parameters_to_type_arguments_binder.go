package visitor

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
	_ int, _ intmod.IObjectType, _ string, _ intsrv.IMethodTypes,
	_ intmod.IExpression) intsrv.IClassFileConstructorInvocationExpression {
	return nil
}

func (b *AbstractTypeParametersToTypeArgumentsBinder) NewSuperConstructorInvocationExpression(
	_ int, _ intmod.IObjectType, _ string,
	_ intsrv.IMethodTypes, _ intmod.IExpression) intsrv.IClassFileSuperConstructorInvocationExpression {
	return nil
}

func (b *AbstractTypeParametersToTypeArgumentsBinder) NewMethodInvocationExpression(
	_ int, _ intmod.IExpression, _ intmod.IObjectType, _, _ string,
	_ intsrv.IMethodTypes, _ intmod.IExpression) intsrv.IClassFileMethodInvocationExpression {
	return nil
}

func (b *AbstractTypeParametersToTypeArgumentsBinder) NewFieldReferenceExpression(
	_ int, _ intmod.IType, _ intmod.IExpression,
	_ intmod.IObjectType, _, _ string) intmod.IFieldReferenceExpression {
	return nil
}

func (b *AbstractTypeParametersToTypeArgumentsBinder) BindParameterTypesWithArgumentTypes(
	_ intmod.IType, _ intmod.IExpression) {
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
