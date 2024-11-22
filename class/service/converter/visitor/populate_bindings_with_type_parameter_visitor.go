package visitor

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
)

func NewPopulateBindingsWithTypeParameterVisitor() *PopulateBindingsWithTypeParameterVisitor {
	return &PopulateBindingsWithTypeParameterVisitor{}
}

type PopulateBindingsWithTypeParameterVisitor struct {
	Bindings   map[string]intmod.ITypeArgument
	TypeBounds map[string]intmod.IType
}

func (v *PopulateBindingsWithTypeParameterVisitor) Init(bindings map[string]intmod.ITypeArgument, typeBounds map[string]intmod.IType) {
	v.Bindings = bindings
	v.TypeBounds = typeBounds
}

func (v *PopulateBindingsWithTypeParameterVisitor) VisitTypeParameter(parameter intmod.ITypeParameter) {
	v.Bindings[parameter.Identifier()] = nil
}

func (v *PopulateBindingsWithTypeParameterVisitor) VisitTypeParameterWithTypeBounds(parameter intmod.ITypeParameterWithTypeBounds) {
	v.Bindings[parameter.Identifier()] = nil
	v.TypeBounds[parameter.Identifier()] = parameter.TypeBounds()
}

func (v *PopulateBindingsWithTypeParameterVisitor) VisitTypeParameters(parameters intmod.ITypeParameters) {
	for _, parameter := range parameters.Elements() {
		parameter.AcceptTypeParameterVisitor(v)
	}
}
