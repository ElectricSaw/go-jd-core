package localvariable

import (
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewGenericLocalVariable(index, offset int, typ *_type.GenericType) *GenericLocalVariable {
	return &GenericLocalVariable{
		AbstractLocalVariable: *NewAbstractLocalVariable(index, offset, ""),
		typ:                   typ,
	}
}

func NewGenericLocalVariableWithAll(index, offset int, typ *_type.GenericType, name string) *GenericLocalVariable {
	return &GenericLocalVariable{
		AbstractLocalVariable: *NewAbstractLocalVariable(index, offset, name),
		typ:                   typ,
	}
}

type GenericLocalVariable struct {
	AbstractLocalVariable

	typ *_type.GenericType
}

func (v *GenericLocalVariable) Type() _type.IType {
	return v.typ
}

func (v *GenericLocalVariable) SetType(typ *_type.GenericType) {
	v.typ = typ
}

func (v *GenericLocalVariable) Dimension() int {
	return v.typ.Dimension()
}

func (v *GenericLocalVariable) Accept(visitor LocalVariableVisitor) {
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

func (v *GenericLocalVariable) IsAssignableFrom(typeBounds map[string]_type.IType, otherType _type.IType) bool {
	return v.typ.Equals(otherType.(*_type.GenericType))
}

func (v *GenericLocalVariable) TypeOnRight(typeBounds map[string]_type.IType, typ _type.IType) {
}

func (v *GenericLocalVariable) TypeOnLeft(typeBounds map[string]_type.IType, typ _type.IType) {
}

func (v *GenericLocalVariable) IsAssignableFromWithVariable(typeBounds map[string]_type.IType, variable *AbstractLocalVariable) bool {
	return v.IsAssignableFrom(typeBounds, variable.Type())
}

func (v *GenericLocalVariable) VariableOnRight(typeBounds map[string]_type.IType, variable ILocalVariableReference) {
}

func (v *GenericLocalVariable) VariableOnLeft(typeBounds map[string]_type.IType, variable ILocalVariableReference) {
}
