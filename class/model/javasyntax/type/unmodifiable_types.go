package _type

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
)

func NewUnmodifiableTypes(types ...intmod.IType) intmod.IUnmodifiableTypes {
	return NewUnmodifiableTypesWithSlice(types)
}

func NewUnmodifiableTypesWithSlice(types []intmod.IType) intmod.IUnmodifiableTypes {
	t := &UnmodifiableTypes{}
	t.AddAll(types)

	return t
}

type UnmodifiableTypes struct {
	Types
}

func (t *UnmodifiableTypes) IsUnmodifiableTypes() bool {
	return true
}
