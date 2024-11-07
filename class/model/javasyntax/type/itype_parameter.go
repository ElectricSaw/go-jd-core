package _type

type ITypeParameter interface {
	TypeParameterVisitable
}

type TypeParameterVisitable interface {
	AcceptTypeParameterVisitor(visitor TypeParameterVisitor)
}

type TypeParameterVisitor interface {
	VisitTypeParameter(parameter *TypeParameter)
	VisitTypeParameterWithTypeBounds(parameter *TypeParameterWithTypeBounds)
	VisitTypeParameters(parameters *TypeParameters)
}
