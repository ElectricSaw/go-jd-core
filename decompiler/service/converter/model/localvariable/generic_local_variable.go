package localvariable

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	_type "github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/type"
)

func NewGenericLocalVariable(index, offset int, typ intmod.IGenericType) intsrv.IGenericLocalVariable {
	return &GenericLocalVariable{
		AbstractLocalVariable: *NewAbstractLocalVariable(index, offset, "").(*AbstractLocalVariable),
		typ:                   typ,
	}
}

func NewGenericLocalVariableWithAll(index, offset int, typ intmod.IGenericType, name string) intsrv.IGenericLocalVariable {
	return &GenericLocalVariable{
		AbstractLocalVariable: *NewAbstractLocalVariable(index, offset, name).(*AbstractLocalVariable),
		typ:                   typ,
	}
}

type GenericLocalVariable struct {
	AbstractLocalVariable

	typ intmod.IGenericType
}

func (v *GenericLocalVariable) Type() intmod.IType {
	return v.typ.(intmod.IType)
}

func (v *GenericLocalVariable) SetType(typ intmod.IGenericType) {
	v.typ = typ
}

func (v *GenericLocalVariable) Dimension() int {
	return v.typ.Dimension()
}

func (v *GenericLocalVariable) Accept(visitor intsrv.ILocalVariableVisitor) {
	visitor.VisitGenericLocalVariable(v)
}

func (v *GenericLocalVariable) String() string {
	sb := fmt.Sprintf("GenericLocalVariable{%s", v.typ.Name())

	if v.typ.Dimension() > 0 {
		for i := 0; i < v.typ.Dimension(); i++ {
			sb += "[]"
		}
	}

	sb += fmt.Sprintf(" %s, index=%d", v.Name(), v.Index())

	if v.Next() != nil {
		sb += fmt.Sprintf(", next=%d", v.Next())
	}

	sb += "}"

	return sb
}

func (v *GenericLocalVariable) IsAssignableFrom(_ map[string]intmod.IType, otherType intmod.IType) bool {
	return v.typ.Equals(otherType.(*_type.GenericType))
}

func (v *GenericLocalVariable) TypeOnRight(_ map[string]intmod.IType, _ intmod.IType) {
}

func (v *GenericLocalVariable) TypeOnLeft(_ map[string]intmod.IType, _ intmod.IType) {
}

func (v *GenericLocalVariable) IsAssignableFromWithVariable(typeBounds map[string]intmod.IType, variable intsrv.ILocalVariable) bool {
	return v.IsAssignableFrom(typeBounds, variable.Type())
}

func (v *GenericLocalVariable) VariableOnRight(_ map[string]intmod.IType, _ intsrv.ILocalVariable) {
}

func (v *GenericLocalVariable) VariableOnLeft(_ map[string]intmod.IType, _ intsrv.ILocalVariable) {
}
