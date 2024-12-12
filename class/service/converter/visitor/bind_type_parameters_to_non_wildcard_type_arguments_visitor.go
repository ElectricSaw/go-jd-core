package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	_type "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/type"
)

func NewBindTypeParametersToNonWildcardTypeArgumentsVisitor() intsrv.IBindTypeParametersToNonWildcardTypeArgumentsVisitor {
	return &BindTypeParametersToNonWildcardTypeArgumentsVisitor{}
}

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
	arguments := _type.NewTypeArgumentsWithCapacity(size)

	for _, parameter := range parameters.ToSlice() {
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

func (v *BindTypeParametersToNonWildcardTypeArgumentsVisitor) VisitDiamondTypeArgument(_ intmod.IDiamondTypeArgument) {
	v.result = nil
}

func (v *BindTypeParametersToNonWildcardTypeArgumentsVisitor) VisitWildcardTypeArgument(_ intmod.IWildcardTypeArgument) {
	v.result = nil
}

func (v *BindTypeParametersToNonWildcardTypeArgumentsVisitor) VisitTypeArguments(_ intmod.ITypeArguments) {
}
func (v *BindTypeParametersToNonWildcardTypeArgumentsVisitor) VisitPrimitiveType(_ intmod.IPrimitiveType) {
}
func (v *BindTypeParametersToNonWildcardTypeArgumentsVisitor) VisitObjectType(_ intmod.IObjectType) {
}
func (v *BindTypeParametersToNonWildcardTypeArgumentsVisitor) VisitInnerObjectType(_ intmod.IInnerObjectType) {
}
func (v *BindTypeParametersToNonWildcardTypeArgumentsVisitor) VisitGenericType(_ intmod.IGenericType) {
}
