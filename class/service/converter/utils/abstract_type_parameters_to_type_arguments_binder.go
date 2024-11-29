package utils

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
)

type ITypeParametersToTypeArgumentsBinder interface {
	newConstructorInvocationExpression(lineNumber int, objectType intmod.IObjectType, descriptor string,
		methodTypes intsrv.IMethodTypes, parameters intmod.IExpression) intsrv.IClassFileConstructorInvocationExpression
	newSuperConstructorInvocationExpression(lineNumber int, objectType intmod.IObjectType, descriptor string,
		methodTypes intsrv.IMethodTypes, parameters intmod.IExpression) intsrv.IClassFileSuperConstructorInvocationExpression
	newMethodInvocationExpression(lineNumber int, expression intmod.IExpression, objectType intmod.IObjectType, name, descriptor string,
		methodTypes intsrv.IMethodTypes, parameters intmod.IExpression) intsrv.IClassFileMethodInvocationExpression
	newFieldReferenceExpression(lineNumber int, typ intmod.IType, expression intmod.IExpression, objectType intmod.IObjectType, name, descriptor string) intmod.IFieldReferenceExpression
	bindParameterTypesWithArgumentTypes(typ intmod.IType, expression intmod.IExpression)
	updateNewExpression(ne intsrv.IClassFileNewExpression, descriptor string, methodTypes intsrv.IMethodTypes, parameters intmod.IExpression)
}

type AbstractTypeParametersToTypeArgumentsBinder struct {
}

func (b *AbstractTypeParametersToTypeArgumentsBinder) newConstructorInvocationExpression(
	lineNumber int, objectType intmod.IObjectType, descriptor string,
	methodTypes intsrv.IMethodTypes, parameters intmod.IExpression) intsrv.IClassFileConstructorInvocationExpression {
	return nil
}

func (b *AbstractTypeParametersToTypeArgumentsBinder) newSuperConstructorInvocationExpression(lineNumber int, objectType intmod.IObjectType, descriptor string,
	methodTypes intsrv.IMethodTypes, parameters intmod.IExpression) intsrv.IClassFileSuperConstructorInvocationExpression {
	return nil
}

func (b *AbstractTypeParametersToTypeArgumentsBinder) newMethodInvocationExpression(
	lineNumber int, expression intmod.IExpression, objectType intmod.IObjectType, name, descriptor string,
	methodTypes intsrv.IMethodTypes, parameters intmod.IExpression) intsrv.IClassFileMethodInvocationExpression {
	return nil
}

func (b *AbstractTypeParametersToTypeArgumentsBinder) newFieldReferenceExpression(
	lineNumber int, typ intmod.IType, expression intmod.IExpression, objectType intmod.IObjectType, name, descriptor string) intmod.IFieldReferenceExpression {
	return nil
}

func (b *AbstractTypeParametersToTypeArgumentsBinder) bindParameterTypesWithArgumentTypes(typ intmod.IType, expression intmod.IExpression) {
}

func (b *AbstractTypeParametersToTypeArgumentsBinder) updateNewExpression(ne intsrv.IClassFileNewExpression,
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
