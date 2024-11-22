package localvariable

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
	"bitbucket.org/coontec/go-jd-core/class/service/converter/utils"
	"fmt"
)

func NewObjectLocalVariable(typeMaker *utils.TypeMaker, index, offset int, typ intmod.IType, name string) intsrv.IObjectLocalVariable {
	return &ObjectLocalVariable{
		AbstractLocalVariable: *NewAbstractLocalVariable(index, offset, name).(*AbstractLocalVariable),
		typeMaker:             typeMaker,
		typ:                   typ,
	}
}

func NewObjectLocalVariable2(typeMaker *utils.TypeMaker, index, offset int, typ intmod.IType, name string, declared bool) intsrv.IObjectLocalVariable {
	v := NewObjectLocalVariable(typeMaker, index, offset, typ, name)
	v.SetDeclared(declared)
	return v
}

func NewObjectLocalVariable3(typeMaker *utils.TypeMaker, index, offset int, objectLocalVariable intsrv.IObjectLocalVariable) intsrv.IObjectLocalVariable {
	return &ObjectLocalVariable{
		AbstractLocalVariable: *NewAbstractLocalVariable(index, offset, "").(*AbstractLocalVariable),
		typeMaker:             typeMaker,
		typ:                   objectLocalVariable.Type(),
	}
}

type ObjectLocalVariable struct {
	AbstractLocalVariable

	typeMaker *utils.TypeMaker
	typ       intmod.IType
}

func (v *ObjectLocalVariable) Type() intmod.IType {
	return v.typ
}

func (v *ObjectLocalVariable) SetType(typeBounds map[string]intmod.IType, t intmod.IType) {
	if !(v.typ == t) {
		v.typ = t
		v.FireChangeEvent(typeBounds)
	}
}

func (v *ObjectLocalVariable) Dimension() int {
	return v.typ.Dimension()
}

func (v *ObjectLocalVariable) Accept(visitor intsrv.ILocalVariableVisitor) {
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

func (v *ObjectLocalVariable) IsAssignableFrom(typeBounds map[string]intmod.IType, typ intmod.IType) bool {
	if v.typ.IsObjectType() {
		if v.typ == _type.OtTypeObject.(intmod.IType) {
			if typ.Dimension() > 0 || !typ.IsPrimitiveType() {
				return true
			}
		}

		if typ.IsObjectType() {
			return v.typeMaker.IsAssignable(typeBounds, v.typ.(intmod.IObjectType), typ.(intmod.IObjectType))
		}
	}
	return false
}

func (v *ObjectLocalVariable) TypeOnRight(typeBounds map[string]intmod.IType, typ intmod.IType) {
	if typ != _type.OtTypeUndefinedObject.(intmod.IType) {
		if v.typ == _type.OtTypeUndefinedObject.(intmod.IType) {
			v.typ = typ
			v.FireChangeEvent(typeBounds)
		} else if v.typ.Dimension() == 0 && typ.Dimension() == 0 {
			if v.typ.IsObjectType() {
				thisObjectType := v.typ.(intmod.IObjectType)

				if typ.IsObjectType() {
					otherObjectType := typ.(intmod.IObjectType)

					if thisObjectType.InternalName() == otherObjectType.InternalName() {
						if thisObjectType.TypeArguments() == nil && otherObjectType.TypeArguments() != nil {
							v.typ = otherObjectType.(intmod.IType)
							v.FireChangeEvent(typeBounds)
						}
					} else if v.typeMaker.IsAssignable(typeBounds, thisObjectType, otherObjectType) {
						if thisObjectType.TypeArguments() == nil && otherObjectType.TypeArguments() != nil {
							v.typ = otherObjectType.CreateTypeWithArgs(otherObjectType.TypeArguments()).(intmod.IType)
							v.FireChangeEvent(typeBounds)
						}
					}
				}
			} else if v.typ.IsGenericType() {
				if typ.IsGenericType() {
					v.typ = typ
					v.FireChangeEvent(typeBounds)
				}
			}
		}
	}
}

func (v *ObjectLocalVariable) TypeOnLeft(typeBounds map[string]intmod.IType, typ intmod.IType) {
	if typ != _type.OtTypeUndefinedObject.(intmod.IType) && typ != _type.OtTypeObject.(intmod.IType) {
		if v.typ == _type.OtTypeUndefinedObject.(intmod.IType) {
			v.typ = typ
			v.FireChangeEvent(typeBounds)
		} else if v.typ.Dimension() == 0 && typ.Dimension() == 0 && v.typ.IsObjectType() && typ.IsObjectType() {
			thisObjectType := v.typ.(intmod.IObjectType)
			otherObjectType := typ.(intmod.IObjectType)

			if thisObjectType.InternalName() == otherObjectType.InternalName() {
				if thisObjectType.TypeArguments() == nil && otherObjectType.TypeArguments() != nil {
					v.typ = otherObjectType.(intmod.IType)
					v.FireChangeEvent(typeBounds)
				}
			} else if v.typeMaker.IsAssignable(typeBounds, thisObjectType, otherObjectType) {
				if thisObjectType.TypeArguments() == nil && otherObjectType.TypeArguments() != nil {
					v.typ = thisObjectType.CreateTypeWithArgs(otherObjectType.TypeArguments()).(intmod.IType)
					v.FireChangeEvent(typeBounds)
				}
			}
		}
	}
}

func (v *ObjectLocalVariable) IsAssignableFromWithVariable(typeBounds map[string]intmod.IType, variable intsrv.ILocalVariable) bool {
	return v.IsAssignableFrom(typeBounds, variable.Type())
}

func (v *ObjectLocalVariable) VariableOnRight(typeBounds map[string]intmod.IType, variable intsrv.ILocalVariable) {
	v.AddVariableOnRight(variable)
	v.TypeOnRight(typeBounds, variable.Type())
}

func (v *ObjectLocalVariable) VariableOnLeft(typeBounds map[string]intmod.IType, variable intsrv.ILocalVariable) {
	v.AddVariableOnLeft(variable)
	v.TypeOnLeft(typeBounds, variable.Type())
}
