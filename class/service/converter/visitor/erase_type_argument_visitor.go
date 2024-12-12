package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	_type "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/type"
)

func NewEraseTypeArgumentVisitor() *EraseTypeArgumentVisitor {
	return &EraseTypeArgumentVisitor{}
}

type EraseTypeArgumentVisitor struct {
	result intmod.IType
}

func (v *EraseTypeArgumentVisitor) Init() {
	v.result = nil
}

func (v *EraseTypeArgumentVisitor) Type() intmod.IType {
	return v.result
}

func (v *EraseTypeArgumentVisitor) VisitPrimitiveType(typ intmod.IPrimitiveType) {
	v.result = typ
}

func (v *EraseTypeArgumentVisitor) VisitObjectType(typ intmod.IObjectType) {
	v.result = typ.CreateTypeWithArgs(nil)
}

func (v *EraseTypeArgumentVisitor) VisitGenericType(_ intmod.IGenericType) {
	v.result = _type.OtTypeObject
}

func (v *EraseTypeArgumentVisitor) VisitInnerObjectType(typ intmod.IInnerObjectType) {
	t := typ.OuterType()

	t.AcceptTypeVisitor(v)

	if v.result == t {
		v.result = t.CreateTypeWithArgs(nil)
	} else {
		v.result = _type.NewInnerObjectTypeWithAll(t.InternalName(), t.QualifiedName(),
			t.Name(), nil, t.Dimension(), v.result.(intmod.IObjectType))
	}
}

func (v *EraseTypeArgumentVisitor) VisitTypes(types intmod.ITypes) {
	size := types.Size()
	i := 0

	for i = 0; i < size; i++ {
		t := types.Get(i)
		t.AcceptTypeVisitor(v)
		if v.result != t {
			break
		}
	}

	if i == size {
		v.result = types
	} else {
		newTypes := _type.NewTypes()

		newTypes.AddAll(types.SubList(0, i).ToSlice())
		newTypes.Add(v.result)

		for i++; i < size; i++ {
			t := types.Get(i)
			t.AcceptTypeVisitor(v)
			newTypes.Add(v.result)
		}

		v.result = newTypes
	}
}
