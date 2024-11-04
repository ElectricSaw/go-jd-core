package _type

func NewTypeParameter(identifier string) *TypeParameter {
	return &TypeParameter{identifier: identifier}
}

type TypeParameter struct {
	ITypeParameter

	identifier string
}

func (t *TypeParameter) Identifier() string {
	return t.identifier
}

func (t *TypeParameter) AcceptTypeParameterVisitor(visitor TypeParameterVisitor) {
	visitor.VisitTypeParameter(t)
}

func (t *TypeParameter) String() string {
	return "TypeParameter{ identifier=" + t.identifier + " }"
}
