package _type

type TypeParameters struct {
	ITypeParameter

	TypeParameters []TypeParameter
}

func (t *TypeParameters) AcceptTypeParameterVisitor(visitor TypeParameterVisitor) {
	visitor.VisitTypeParameters(t)
}
