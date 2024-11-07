package _type

func NewTypeParameters() *TypeParameters {
	return &TypeParameters{}
}

type TypeParameters struct {
	ITypeParameter

	TypeParameters []ITypeParameter
}

func (t *TypeParameters) Add(typeParameter ITypeParameter) {
	t.TypeParameters = append(t.TypeParameters, typeParameter)
}

func (t *TypeParameters) AcceptTypeParameterVisitor(visitor TypeParameterVisitor) {
	visitor.VisitTypeParameters(t)
}
