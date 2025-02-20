package localvariable

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/service/converter/visitor"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewObjectLocalVariable(typeMaker *visitor.TypeMaker, index, offset int, typ model.IType, name string) ObjectLocalVariable {
	return ObjectLocalVariable{
		declared:         offset == 0,
		index:            index,
		fromOffset:       offset,
		toOffset:         offset,
		name:             name,
		references:       util.NewDefaultList[ILocalVariable](),
		variablesOnRight: util.NewSet[ILocalVariable](),
		variablesOnLeft:  util.NewSet[ILocalVariable](),
		typeMaker:        typeMaker,
		typ:              typ,
	}
}

func NewObjectLocalVariable2(typeMaker *visitor.TypeMaker, index, offset int, typ model.IType, name string, declared bool) ObjectLocalVariable {
	return ObjectLocalVariable{
		declared:         declared,
		index:            index,
		fromOffset:       offset,
		toOffset:         offset,
		name:             name,
		references:       util.NewDefaultList[ILocalVariable](),
		variablesOnRight: util.NewSet[ILocalVariable](),
		variablesOnLeft:  util.NewSet[ILocalVariable](),
		typeMaker:        typeMaker,
		typ:              typ,
	}

}

func NewObjectLocalVariable3(typeMaker *visitor.TypeMaker, index, offset int, objectLocalVariable ObjectLocalVariable) ObjectLocalVariable {
	return ObjectLocalVariable{
		declared:         offset == 0,
		index:            index,
		fromOffset:       offset,
		toOffset:         offset,
		name:             "",
		references:       util.NewDefaultList[ILocalVariable](),
		variablesOnRight: util.NewSet[ILocalVariable](),
		variablesOnLeft:  util.NewSet[ILocalVariable](),
		typeMaker:        typeMaker,
		typ:              objectLocalVariable.Type(),
	}
}

type ObjectLocalVariable struct {
	frame            IFrame
	next             ILocalVariable
	declared         bool
	index            int
	fromOffset       int
	toOffset         int
	name             string
	references       util.IList[ILocalVariable]
	variablesOnRight util.ISet[ILocalVariable]
	variablesOnLeft  util.ISet[ILocalVariable]
	typeMaker        *visitor.TypeMaker
	typ              model.IType
}

func (v *ObjectLocalVariable) Frame() IFrame {
	return v.frame
}

func (v *ObjectLocalVariable) Next() ILocalVariable {
	return v.next
}

func (v *ObjectLocalVariable) IsDeclared() bool {
	return v.declared
}

func (v *ObjectLocalVariable) Index() int {
	return v.index
}

func (v *ObjectLocalVariable) FromOffset() int {
	return v.fromOffset
}

func (v *ObjectLocalVariable) ToOffset() int {
	return v.toOffset
}

func (v *ObjectLocalVariable) Type() model.IType {
	return v.typ
}

func (v *ObjectLocalVariable) Name() string {
	return v.name
}

func (v *ObjectLocalVariable) Dimension() int {
	return v.typ.GetDimension()
}

func (v *ObjectLocalVariable) References() util.IList[ILocalVariable] {
	return v.references
}

func (v *ObjectLocalVariable) SetFrame(frame IFrame) {
	v.frame = frame
}

func (v *ObjectLocalVariable) SetNext(localVariable ILocalVariable) {
	v.next = localVariable
}

func (v *ObjectLocalVariable) SetDeclared(declared bool) {
	v.declared = declared

}

func (v *ObjectLocalVariable) SetIndex(index int) {
	v.index = index
}

func (v *ObjectLocalVariable) SetFromOffset(fromOffset int) {
	v.fromOffset = fromOffset
}

func (v *ObjectLocalVariable) SetToOffset(offset int) {
	if v.fromOffset > offset {
		v.fromOffset = offset
	}
	if v.toOffset < offset {
		v.toOffset = offset
	}
}

func (v *ObjectLocalVariable) SetToOffsetForce(offset int, force bool) {
	v.toOffset = offset
}

func (v *ObjectLocalVariable) SetType(t model.IType) {
}

func (v *ObjectLocalVariable) SetTypeWithTypeBounds(typeBounds map[string]model.IType, t model.IType) {
	if !(v.typ == t) {
		v.typ = t
		v.FireChangeEvent(typeBounds)
	}
}

func (v *ObjectLocalVariable) SetName(name string) {
	v.name = name
}

func (v *ObjectLocalVariable) SetDimension(dimension int) {
}

func (v *ObjectLocalVariable) Accept(visitor ILocalVariableVisitor) {
	visitor.VisitObjectLocalVariable(v)
}

func (v *ObjectLocalVariable) AddReference(reference ILocalVariableReference) {
	v.references.Add(reference.(ILocalVariable))
}

func (v *ObjectLocalVariable) IsAssignableFrom(typeBounds map[string]model.IType, typ model.IType) bool {
	if v.typ.IsObjectType() {
		if v.typ == &model.OtTypeObject {
			if typ.GetDimension() > 0 || !typ.IsPrimitiveType() {
				return true
			}
		}

		if typ.IsObjectType() {
			return v.typeMaker.IsAssignable(typeBounds, v.typ.(*model.ObjectType), typ.(*model.ObjectType))
		}
	}
	return false
}

func (v *ObjectLocalVariable) TypeOnRight(typeBounds map[string]model.IType, typ model.IType) {
	if typ != &model.OtTypeUndefinedObject {
		if v.typ == &model.OtTypeUndefinedObject {
			v.typ = typ
			v.FireChangeEvent(typeBounds)
		} else if v.typ.GetDimension() == 0 && typ.GetDimension() == 0 {
			if v.typ.IsObjectType() {
				thisObjectType := v.typ.(*model.ObjectType)

				if typ.IsObjectType() {
					otherObjectType := typ.(*model.ObjectType)

					if thisObjectType.InternalName == otherObjectType.InternalName {
						if thisObjectType.TypeArguments == nil && otherObjectType.TypeArguments != nil {
							v.typ = otherObjectType
							v.FireChangeEvent(typeBounds)
						}
					} else if v.typeMaker.IsAssignable(typeBounds, thisObjectType, otherObjectType) {
						if thisObjectType.TypeArguments == nil && otherObjectType.TypeArguments != nil {
							v.typ = otherObjectType.CreateTypeWithArgs(otherObjectType.TypeArguments).(model.IType)
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

func (v *ObjectLocalVariable) TypeOnLeft(typeBounds map[string]model.IType, typ model.IType) {
	if typ != &model.OtTypeUndefinedObject && typ != &model.OtTypeObject {
		if v.typ == &model.OtTypeUndefinedObject {
			v.typ = typ
			v.FireChangeEvent(typeBounds)
		} else if v.typ.GetDimension() == 0 && typ.GetDimension() == 0 && v.typ.IsObjectType() && typ.IsObjectType() {
			thisObjectType := v.typ.(*model.ObjectType)
			otherObjectType := typ.(*model.ObjectType)

			if thisObjectType.InternalName == otherObjectType.InternalName {
				if thisObjectType.TypeArguments == nil && otherObjectType.TypeArguments != nil {
					v.typ = otherObjectType
					v.FireChangeEvent(typeBounds)
				}
			} else if v.typeMaker.IsAssignable(typeBounds, thisObjectType, otherObjectType) {
				if thisObjectType.TypeArguments == nil && otherObjectType.TypeArguments != nil {
					v.typ = thisObjectType.CreateTypeWithArgs(otherObjectType.TypeArguments).(model.IType)
					v.FireChangeEvent(typeBounds)
				}
			}
		}
	}
}

func (v *ObjectLocalVariable) IsAssignableFromWithVariable(typeBounds map[string]model.IType, variable ILocalVariable) bool {
	return v.IsAssignableFrom(typeBounds, variable.Type())
}

func (v *ObjectLocalVariable) VariableOnRight(typeBounds map[string]model.IType, variable ILocalVariable) {
	v.AddVariableOnRight(variable)
	v.TypeOnRight(typeBounds, variable.Type())
}

func (v *ObjectLocalVariable) VariableOnLeft(typeBounds map[string]model.IType, variable ILocalVariable) {
	v.AddVariableOnLeft(variable)
	v.TypeOnLeft(typeBounds, variable.Type())
}

func (v *ObjectLocalVariable) FireChangeEvent(typeBounds map[string]model.IType) {
	if v.variablesOnLeft != nil {
		for _, variable := range v.variablesOnLeft.ToSlice() {
			v.VariableOnRight(typeBounds, variable)
		}
	}
	if v.variablesOnRight != nil {
		for _, variable := range v.variablesOnRight.ToSlice() {
			v.VariableOnLeft(typeBounds, variable)
		}
	}
}

func (v *ObjectLocalVariable) AddVariableOnLeft(variable ILocalVariable) {
	if v.variablesOnLeft == nil {
		v.variablesOnLeft = util.NewSet[ILocalVariable]()
		v.variablesOnLeft.Add(variable)
		variable.AddVariableOnRight(v)
	} else if !v.variablesOnLeft.Contains(variable) {
		v.variablesOnLeft.Add(variable)
		variable.AddVariableOnRight(v)
	}
}

func (v *ObjectLocalVariable) AddVariableOnRight(variable ILocalVariable) {
	if v.variablesOnRight == nil {
		v.variablesOnRight = util.NewSet[ILocalVariable]()
		v.variablesOnRight.Add(variable)
		variable.AddVariableOnLeft(v)
	} else if !v.variablesOnRight.Contains(variable) {
		v.variablesOnRight.Add(variable)
		variable.AddVariableOnLeft(v)
	}
}

func (v *ObjectLocalVariable) IsPrimitiveLocalVariable() bool {
	return false
}

func (v *ObjectLocalVariable) LocalVariable() ILocalVariableReference {
	return nil
}

func (v *ObjectLocalVariable) SetLocalVariable(_ ILocalVariableReference) {

}

func (v *ObjectLocalVariable) String() string {
	sb := "ObjectLocalVariable{"

	if v.typ.GetName() == "" {
		sb += v.typ.GetInternalName()
	} else {
		sb += v.typ.GetName()
	}

	if v.typ.GetDimension() > 0 {
		for i := 0; i < v.typ.GetDimension(); i++ {
			sb += "[]"
		}
	}

	sb += fmt.Sprintf(" %s, index=%d", v.Name(), v.Index())

	if v.Next() != nil {
		sb += fmt.Sprintf(", next=%s", v.Next())
	}

	return sb + "}"
}
