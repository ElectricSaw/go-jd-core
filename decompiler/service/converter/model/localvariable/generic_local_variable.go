package localvariable

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewGenericLocalVariable(index, offset int, typ *model.GenericType) GenericLocalVariable {
	return NewGenericLocalVariableWithAll(index, offset, typ, "")
}

func NewGenericLocalVariableWithAll(index, offset int, typ *model.GenericType, name string) GenericLocalVariable {
	return GenericLocalVariable{
		declared:         offset == 0,
		index:            index,
		fromOffset:       offset,
		toOffset:         offset,
		name:             name,
		references:       util.NewDefaultList[ILocalVariable](),
		variablesOnRight: util.NewSet[ILocalVariable](),
		variablesOnLeft:  util.NewSet[ILocalVariable](),
		typ:              typ,
	}
}

type GenericLocalVariable struct {
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
	typ              *model.GenericType
}

func (v *GenericLocalVariable) Frame() IFrame {
	return v.frame
}

func (v *GenericLocalVariable) Next() ILocalVariable {
	return v.next
}

func (v *GenericLocalVariable) IsDeclared() bool {
	return v.declared
}

func (v *GenericLocalVariable) Index() int {
	return v.index
}

func (v *GenericLocalVariable) FromOffset() int {
	return v.fromOffset
}

func (v *GenericLocalVariable) ToOffset() int {
	return v.toOffset
}

func (v *GenericLocalVariable) Type() model.IType {
	return v.typ
}

func (v *GenericLocalVariable) Name() string {
	return v.name
}

func (v *GenericLocalVariable) Dimension() int {
	return v.typ.Dimension
}

func (v *GenericLocalVariable) References() util.IList[ILocalVariable] {
	return v.references
}

func (v *GenericLocalVariable) SetFrame(frame IFrame) {
	v.frame = frame
}

func (v *GenericLocalVariable) SetNext(localVariable ILocalVariable) {
	v.next = localVariable
}

func (v *GenericLocalVariable) SetDeclared(declared bool) {
	v.declared = declared

}

func (v *GenericLocalVariable) SetIndex(index int) {
	v.index = index
}

func (v *GenericLocalVariable) SetFromOffset(fromOffset int) {
	v.fromOffset = fromOffset
}

func (v *GenericLocalVariable) SetToOffset(offset int) {
	if v.fromOffset > offset {
		v.fromOffset = offset
	}
	if v.toOffset < offset {
		v.toOffset = offset
	}
}

func (v *GenericLocalVariable) SetToOffsetForce(offset int, force bool) {
	v.toOffset = offset
}

func (v *GenericLocalVariable) SetType(t model.IType) {
	v.typ = t.(*model.GenericType)
}

func (v *GenericLocalVariable) SetName(name string) {
	v.name = name
}

func (v *GenericLocalVariable) SetDimension(dimension int) {
}

func (v *GenericLocalVariable) Accept(visitor ILocalVariableVisitor) {
	visitor.VisitGenericLocalVariable(v)
}

func (v *GenericLocalVariable) AddReference(reference ILocalVariableReference) {
	v.references.Add(reference.(ILocalVariable))
}

func (v *GenericLocalVariable) IsAssignableFrom(_ map[string]model.IType, otherType model.IType) bool {
	return v.typ.Equals(otherType)
}

func (v *GenericLocalVariable) TypeOnRight(_ map[string]model.IType, _ model.IType) {
}

func (v *GenericLocalVariable) TypeOnLeft(_ map[string]model.IType, _ model.IType) {
}

func (v *GenericLocalVariable) IsAssignableFromWithVariable(typeBounds map[string]model.IType, variable ILocalVariable) bool {
	return v.IsAssignableFrom(typeBounds, variable.Type())
}

func (v *GenericLocalVariable) VariableOnRight(_ map[string]model.IType, _ ILocalVariable) {
}

func (v *GenericLocalVariable) VariableOnLeft(_ map[string]model.IType, _ ILocalVariable) {
}

func (v *GenericLocalVariable) FireChangeEvent(typeBounds map[string]model.IType) {
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

func (v *GenericLocalVariable) AddVariableOnLeft(variable ILocalVariable) {
	if v.variablesOnLeft == nil {
		v.variablesOnLeft = util.NewSet[ILocalVariable]()
		v.variablesOnLeft.Add(variable)
		variable.AddVariableOnRight(v)
	} else if !v.variablesOnLeft.Contains(variable) {
		v.variablesOnLeft.Add(variable)
		variable.AddVariableOnRight(v)
	}
}

func (v *GenericLocalVariable) AddVariableOnRight(variable ILocalVariable) {
	if v.variablesOnRight == nil {
		v.variablesOnRight = util.NewSet[ILocalVariable]()
		v.variablesOnRight.Add(variable)
		variable.AddVariableOnLeft(v)
	} else if !v.variablesOnRight.Contains(variable) {
		v.variablesOnRight.Add(variable)
		variable.AddVariableOnLeft(v)
	}
}

func (v *GenericLocalVariable) IsPrimitiveLocalVariable() bool {
	return false
}

func (v *GenericLocalVariable) LocalVariable() ILocalVariableReference {
	return nil
}

func (v *GenericLocalVariable) SetLocalVariable(_ ILocalVariableReference) {

}

func (v *GenericLocalVariable) String() string {
	sb := fmt.Sprintf("GenericLocalVariable{%s", v.typ.Name)

	if v.typ.Dimension > 0 {
		for i := 0; i < v.typ.Dimension; i++ {
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
