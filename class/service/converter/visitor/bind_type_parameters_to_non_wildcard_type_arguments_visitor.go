package visitor

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
)

type BindTypeParametersToNonWildcardTypeArgumentsVisitor struct {
	bindings map[string]intmod.ITypeArgument
	result   intmod.ITypeArgument
}

func (v *BindTypeParametersToNonWildcardTypeArgumentsVisitor) Init(bindings map[string]intmod.ITypeArgument) {
	v.bindings = bindings
	v.result = nil
}

func (v *BindTypeParametersToNonWildcardTypeArgumentsVisitor) TypeArgument() intmod.ITypeArgument {
	return v.result
}

func (v *BindTypeParametersToNonWildcardTypeArgumentsVisitor) VisitTypeParameter(parameter intmod.ITypeParameter) {
	v.result = v.bindings[parameter.Identifier()]

	if v.result != nil {
		v.result.AcceptTypeArgumentVisitor(v)
	}
}

func (v *BindTypeParametersToNonWildcardTypeArgumentsVisitor) VisitTypeParameterWithTypeBounds(parameter intmod.ITypeParameterWithTypeBounds) {
	v.result = v.bindings[parameter.Identifier()]

	if v.result != nil {
		v.result.AcceptTypeArgumentVisitor(v)
	}
}

func (v *BindTypeParametersToNonWildcardTypeArgumentsVisitor) VisitTypeParameters(parameters intmod.ITypeParameters) {
	size := parameters.Size()
	arguments := _type.NewTypeArgumentsWithSize(size)

	for _, parameter := range parameters.Elements() {
		parameter.AcceptTypeParameterVisitor(v)

		if v.result == nil {
			return
		}

		arguments.Add(v.result)
	}

	v.result = arguments
}

func (v *BindTypeParametersToNonWildcardTypeArgumentsVisitor) VisitWildcardExtendsTypeArgument(argument intmod.IWildcardExtendsTypeArgument) {
	v.result = argument.Type()
}

func (v *BindTypeParametersToNonWildcardTypeArgumentsVisitor) VisitWildcardSuperTypeArgument(argument intmod.IWildcardSuperTypeArgument) {
	v.result = argument.Type()
}

func (v *BindTypeParametersToNonWildcardTypeArgumentsVisitor) VisitDiamondTypeArgument(argument intmod.IDiamondTypeArgument) {
	v.result = nil
}

func (v *BindTypeParametersToNonWildcardTypeArgumentsVisitor) VisitWildcardTypeArgument(argument intmod.IWildcardTypeArgument) {
	v.result = nil
}

func (v *BindTypeParametersToNonWildcardTypeArgumentsVisitor) VisitTypeArguments(arguments intmod.ITypeArguments) {
}
func (v *BindTypeParametersToNonWildcardTypeArgumentsVisitor) VisitPrimitiveType(t intmod.IPrimitiveType) {
}
func (v *BindTypeParametersToNonWildcardTypeArgumentsVisitor) VisitObjectType(t intmod.IObjectType) {
}
func (v *BindTypeParametersToNonWildcardTypeArgumentsVisitor) VisitInnerObjectType(t intmod.IInnerObjectType) {
}
func (v *BindTypeParametersToNonWildcardTypeArgumentsVisitor) VisitGenericType(t intmod.IGenericType) {
}
