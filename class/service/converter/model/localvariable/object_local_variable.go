package localvariable

import (
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"bitbucket.org/coontec/javaClass/class/service/converter/utils"
	"fmt"
)

func NewObjectLocalVariable(typeMaker *utils.TypeMaker, index, offset int, typ _type.IType, name string) *ObjectLocalVariable {
	return &ObjectLocalVariable{
		AbstractLocalVariable: *NewAbstractLocalVariable(index, offset, name),
		typeMaker:             typeMaker,
		typ:                   typ,
	}
}

func NewObjectLocalVariable2(typeMaker *utils.TypeMaker, index, offset int, typ _type.IType, name string, declared bool) *ObjectLocalVariable {
	v := NewObjectLocalVariable(typeMaker, index, offset, typ, name)
	v.declared = declared
	return v
}

func NewObjectLocalVariable3(typeMaker *utils.TypeMaker, index, offset int, objectLocalVariable *ObjectLocalVariable) *ObjectLocalVariable {
	return &ObjectLocalVariable{
		AbstractLocalVariable: *NewAbstractLocalVariable(index, offset, ""),
		typeMaker:             typeMaker,
		typ:                   objectLocalVariable.typ,
	}
}

type ObjectLocalVariable struct {
	AbstractLocalVariable

	typeMaker *utils.TypeMaker
	typ       _type.IType
}

func (v *ObjectLocalVariable) Type() _type.IType {
	return v.typ
}

func (v *ObjectLocalVariable) SetType(typeBounds map[string]_type.IType, t _type.IType) {
	if !(v.typ == t) {
		v.typ = t
		v.FireChangeEvent(typeBounds)
	}
}

func (v *ObjectLocalVariable) Dimension() int {
	return v.typ.Dimension()
}

func (v *ObjectLocalVariable) Accept(visitor LocalVariableVisitor) {
	visitor.VisitObjectLocalVariable(v)
}

func (v *ObjectLocalVariable) String() string {
	sb := "ObjectLocalVariable{"

	if v.typ.Name() == "" {
		sb += v.typ.InternalName()
	} else {
		sb += v.typ.Name()
	}

	if v.typ.Dimension() > 0 {
		for i := 0; i < v.typ.Dimension(); i++ {
			sb += "[]"
		}
	}

	sb += fmt.Sprintf(" %s, index=%d", v.Name(), v.Index())

	if v.Next() != nil {
		sb += fmt.Sprintf(", next=%s", v.Next())
	}

	return sb + "}"
}

func (v *ObjectLocalVariable) IsAssignableFrom(typeBounds map[string]_type.IType, typ _type.IType) bool {
	if v.typ.IsObjectType() {
		if v.typ == _type.OtTypeObject {
			if typ.Dimension() > 0 || !typ.IsPrimitiveType() {
				return true
			}
		}

		if typ.IsObjectType() {
			return v.typeMaker.IsAssignable(typeBounds, typ)
		}
	}
	return false
}

func (v *ObjectLocalVariable) TypeOnRight(typeBounds map[string]_type.IType, typ _type.IType) {
}

func (v *ObjectLocalVariable) TypeOnLeft(typeBounds map[string]_type.IType, typ _type.IType) {
}

func (v *ObjectLocalVariable) IsAssignableFromWithVariable(typeBounds map[string]_type.IType, variable ILocalVariableReference) bool {
	return false
}

func (v *ObjectLocalVariable) VariableOnRight(typeBounds map[string]_type.IType, variable ILocalVariableReference) {

}

func (v *ObjectLocalVariable) VariableOnLeft(typeBounds map[string]_type.IType, variable ILocalVariableReference) {

}