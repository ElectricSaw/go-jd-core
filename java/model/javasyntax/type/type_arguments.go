package _type

type TypeArguments struct {
	AbstractTypeArgument

	TypeArguments []ITypeArgument
}

func (t *TypeArguments) IsList() bool {
	return true
}

func (t *TypeArguments) Size() int {
	return len(t.TypeArguments)
}

func (t *TypeArguments) IsTypeArgumentAssignableFrom(typeBounds map[string]IType, typeArgument ITypeArgument) bool {
	ata, ok := typeArgument.(*TypeArguments)
	if !ok {
		return false
	}

	if t.Size() != ata.Size() {
		return false
	}

	for i := 0; i < t.Size(); i++ {
		if !t.TypeArguments[i].IsTypeArgumentAssignableFrom(typeBounds, ata.TypeArguments[i]) {
			return false
		}
	}

	return true
}

func (t *TypeArguments) IsTypeArgumentList() bool {
	return true
}

func (t *TypeArguments) GetTypeArgumentFirst() ITypeArgument {
	return t.TypeArguments[0]
}

func (t *TypeArguments) GetTypeArgumentLIst() []ITypeArgument {
	return t.TypeArguments
}

func (t *TypeArguments) TypeArgumentSize() int {
	return len(t.TypeArguments)
}

func (t *TypeArguments) AcceptTypeArgumentVisitor(visitor TypeArgumentVisitor) {
	visitor.VisitTypeArguments(t)
}
