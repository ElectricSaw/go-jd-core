package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
)

func NewPopulateBindingsWithTypeParameterVisitor() intsrv.IPopulateBindingsWithTypeParameterVisitor {
	return &PopulateBindingsWithTypeParameterVisitor{}
}

type PopulateBindingsWithTypeParameterVisitor struct {
	bindings   map[string]intmod.ITypeArgument
	typeBounds map[string]intmod.IType
}

func (v *PopulateBindingsWithTypeParameterVisitor) Init(bindings map[string]intmod.ITypeArgument, typeBounds map[string]intmod.IType) {
	v.bindings = bindings
	v.typeBounds = typeBounds
}

func (v *PopulateBindingsWithTypeParameterVisitor) Bindings() map[string]intmod.ITypeArgument {
	return v.bindings
}

func (v *PopulateBindingsWithTypeParameterVisitor) TypeBounds() map[string]intmod.IType {
	return v.typeBounds
}

func (v *PopulateBindingsWithTypeParameterVisitor) VisitTypeParameter(parameter intmod.ITypeParameter) {
	v.bindings[parameter.Identifier()] = nil
}

func (v *PopulateBindingsWithTypeParameterVisitor) VisitTypeParameterWithTypeBounds(parameter intmod.ITypeParameterWithTypeBounds) {
	v.bindings[parameter.Identifier()] = nil
	v.typeBounds[parameter.Identifier()] = parameter.TypeBounds()
}

func (v *PopulateBindingsWithTypeParameterVisitor) VisitTypeParameters(parameters intmod.ITypeParameters) {
	for _, parameter := range parameters.ToSlice() {
		parameter.AcceptTypeParameterVisitor(v)
	}
}
