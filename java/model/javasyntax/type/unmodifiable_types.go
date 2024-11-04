package _type

func NewUnmodifiableTypes(types ...IType) *UnmodifiableTypes {
	return NewUnmodifiableTypesWithSlice(types)
}

func NewUnmodifiableTypesWithSlice(types []IType) *UnmodifiableTypes {
	t := &UnmodifiableTypes{}

	for _, i := range types {
		t.Types.Types = append(t.Types.Types, i)
	}

	return t
}

type UnmodifiableTypes struct {
	Types
}

func (t *UnmodifiableTypes) ListIterator(i int) []IType {
	return t.Types.Types
}
