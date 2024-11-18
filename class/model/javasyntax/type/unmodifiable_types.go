package _type

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/model"
)

func NewUnmodifiableTypes(types ...intsyn.IType) intsyn.IUnmodifiableTypes {
	return NewUnmodifiableTypesWithSlice(types)
}

func NewUnmodifiableTypesWithSlice(types []intsyn.IType) intsyn.IUnmodifiableTypes {
	t := &UnmodifiableTypes{}
	t.AddAll(types)

	return t
}

type UnmodifiableTypes struct {
	Types
}

func (t *UnmodifiableTypes) ListIterator(i int) []intsyn.IType {
	return t.Elements()
}
