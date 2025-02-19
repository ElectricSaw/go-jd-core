package localvariable

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
)

func NewGenericLocalVariable(index, offset int, typ model.GenericType) GenericLocalVariable {
	return &GenericLocalVariable{
		AbstractLocalVariable: *NewAbstractLocalVariable(index, offset, "").(*AbstractLocalVariable),
		typ:                   typ,
	}
}

func NewGenericLocalVariableWithAll(index, offset int, typ model.GenericType, name string) GenericLocalVariable {
	return &GenericLocalVariable{
		AbstractLocalVariable: *NewAbstractLocalVariable(index, offset, name).(*AbstractLocalVariable),
		typ:                   typ,
	}
}

type GenericLocalVariable struct {
	AbstractLocalVariable

	Type model.GenericType
}

func (v *GenericLocalVariable) SetType(typ model.GenericType) {
	v.typ = typ
}

func (v *GenericLocalVariable) Dimension() int {
	return v.typ.Dimension()
}

func (v *GenericLocalVariable) Accept(visitor model.LocalVariableVisitor) {
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

func (v *GenericLocalVariable) IsAssignableFrom(_ map[string]model.Type, otherType model.Type) bool {
	return v.typ.Equals(otherType.(*_type.GenericType))
}

func (v *GenericLocalVariable) TypeOnRight(_ map[string]model.Type, _ model.Type) {
}

func (v *GenericLocalVariable) TypeOnLeft(_ map[string]model.Type, _ model.Type) {
}

func (v *GenericLocalVariable) IsAssignableFromWithVariable(typeBounds map[string]model.Type, variable model.LocalVariable) bool {
	return v.IsAssignableFrom(typeBounds, variable.Type())
}

func (v *GenericLocalVariable) VariableOnRight(_ map[string]model.Type, _ model.LocalVariable) {
}

func (v *GenericLocalVariable) VariableOnLeft(_ map[string]model.Type, _ model.LocalVariable) {
}
