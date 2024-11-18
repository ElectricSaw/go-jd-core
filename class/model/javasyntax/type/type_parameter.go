package _type

import intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"

func NewTypeParameter(identifier string) intsyn.ITypeParameter {
	return &TypeParameter{identifier: identifier}
}

type TypeParameter struct {
	AbstractTypeParameter

	identifier string
}

func (t *TypeParameter) Identifier() string {
	return t.identifier
}

func (t *TypeParameter) AcceptTypeParameterVisitor(visitor intsyn.ITypeParameterVisitor) {
	visitor.VisitTypeParameter(t)
}

func (t *TypeParameter) String() string {
	return "TypeParameter{ identifier=" + t.identifier + " }"
}
