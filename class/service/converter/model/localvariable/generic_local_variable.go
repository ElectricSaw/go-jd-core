package localvariable

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
	"fmt"
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

func (v *GenericLocalVariable) IsAssignableFrom(typeBounds map[string]intmod.IType, otherType intmod.IType) bool {
	return v.typ.Equals(otherType.(*_type.GenericType))
}

func (v *GenericLocalVariable) TypeOnRight(typeBounds map[string]intmod.IType, typ intmod.IType) {
}

func (v *GenericLocalVariable) TypeOnLeft(typeBounds map[string]intmod.IType, typ intmod.IType) {
}

func (v *GenericLocalVariable) IsAssignableFromWithVariable(typeBounds map[string]intmod.IType, variable intsrv.ILocalVariable) bool {
	return v.IsAssignableFrom(typeBounds, variable.Type())
}

func (v *GenericLocalVariable) VariableOnRight(typeBounds map[string]intmod.IType, variable intsrv.ILocalVariable) {
}

func (v *GenericLocalVariable) VariableOnLeft(typeBounds map[string]intmod.IType, variable intsrv.ILocalVariable) {
}
