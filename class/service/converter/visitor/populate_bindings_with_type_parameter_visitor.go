package visitor

import _type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"

type PopulateBindingsWithTypeParameterVisitor struct {
	Bindings   map[string]_type.ITypeArgument
	TypeBounds map[string]_type.IType
}

func (v *PopulateBindingsWithTypeParameterVisitor) Init(bindings map[string]_type.ITypeArgument, typeBounds map[string]_type.IType) {
	v.Bindings = bindings
	v.TypeBounds = typeBounds
}

func (v *PopulateBindingsWithTypeParameterVisitor) VisitTypeParameter(parameter *_type.TypeParameter) {
	v.Bindings[parameter.Identifier()] = nil
}

func (v *PopulateBindingsWithTypeParameterVisitor) VisitTypeParameterWithTypeBounds(parameter *_type.TypeParameterWithTypeBounds) {
	v.Bindings[parameter.Identifier()] = nil
	v.TypeBounds[parameter.Identifier()] = parameter.TypeBounds()
}

func (v *PopulateBindingsWithTypeParameterVisitor) VisitTypeParameters(parameters *_type.TypeParameters) {
	for _, parameter := range parameters.TypeParameters {
		parameter.AcceptTypeParameterVisitor(v)
	}
}
