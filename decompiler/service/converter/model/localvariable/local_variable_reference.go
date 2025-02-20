package localvariable

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewAbstractLocalVariable(index, offset int, name string) ILocalVariable {
	return NewAbstractLocalVariableWithAll(index, offset, name, offset == 0)
}

func NewAbstractLocalVariableWithAll(index, offset int, name string, declared bool) ILocalVariable {
	return AbstractLocalVariable{
		declared:         declared,
		index:            index,
		fromOffset:       offset,
		toOffset:         offset,
		name:             name,
		references:       util.NewDefaultList[ILocalVariable](),
		variablesOnRight: util.NewSet[ILocalVariable](),
		variablesOnLeft:  util.NewSet[ILocalVariable](),
	}
}

type AbstractLocalVariable struct {
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
}

func (v *AbstractLocalVariable) Frame() IFrame {
	return v.frame
}

func (v *AbstractLocalVariable) Next() ILocalVariable {
	return v.next
}

func (v *AbstractLocalVariable) IsDeclared() bool {
	return v.declared
}

func (v *AbstractLocalVariable) Index() int {
	return v.index
}

func (v *AbstractLocalVariable) FromOffset() int {
	return v.fromOffset
}

func (v *AbstractLocalVariable) ToOffset() int {
	return v.toOffset
}

func (v *AbstractLocalVariable) Type() model.IType {
	return nil
}

func (v *AbstractLocalVariable) Name() string {
	return v.name
}

func (v *AbstractLocalVariable) Dimension() int {
	return 0
}

func (v *AbstractLocalVariable) References() util.IList[ILocalVariable] {
	return v.references
}

func (v *AbstractLocalVariable) SetFrame(frame IFrame) {
	v.frame = frame
}

func (v *AbstractLocalVariable) SetNext(localVariable ILocalVariable) {
	v.next = localVariable
}

func (v *AbstractLocalVariable) SetDeclared(declared bool) {
	v.declared = declared

}

func (v *AbstractLocalVariable) SetIndex(index int) {
	v.index = index
}

func (v *AbstractLocalVariable) SetFromOffset(fromOffset int) {
	v.fromOffset = fromOffset
}

func (v *AbstractLocalVariable) SetToOffset(offset int) {
	if v.fromOffset > offset {
		v.fromOffset = offset
	}
	if v.toOffset < offset {
		v.toOffset = offset
	}
}

func (v *AbstractLocalVariable) SetToOffsetForce(offset int, force bool) {
	v.toOffset = offset
}

func (v *AbstractLocalVariable) SetType(t model.IType) {
}

func (v *AbstractLocalVariable) SetName(name string) {
	v.name = name
}

func (v *AbstractLocalVariable) SetDimension(dimension int) {
}

func (v *AbstractLocalVariable) Accept(_ ILocalVariableVisitor) {
}

func (v *AbstractLocalVariable) AddReference(reference ILocalVariableReference) {
	v.references.Add(reference.(ILocalVariable))
}

func (v *AbstractLocalVariable) IsAssignableFrom(_ map[string]model.IType, _ model.IType) bool {
	return false
}

func (v *AbstractLocalVariable) TypeOnRight(_ map[string]model.IType, _ model.IType) {
}

func (v *AbstractLocalVariable) TypeOnLeft(_ map[string]model.IType, _ model.IType) {
}

func (v *AbstractLocalVariable) IsAssignableFromWithVariable(_ map[string]model.IType, _ ILocalVariable) bool {
	return false
}

func (v *AbstractLocalVariable) VariableOnRight(_ map[string]model.IType, _ ILocalVariable) {
}

func (v *AbstractLocalVariable) VariableOnLeft(_ map[string]model.IType, _ ILocalVariable) {
}

func (v *AbstractLocalVariable) FireChangeEvent(typeBounds map[string]model.IType) {
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

func (v *AbstractLocalVariable) AddVariableOnLeft(variable ILocalVariable) {
	if v.variablesOnLeft == nil {
		v.variablesOnLeft = util.NewSet[ILocalVariable]()
		v.variablesOnLeft.Add(variable)
		variable.AddVariableOnRight(v)
	} else if !v.variablesOnLeft.Contains(variable) {
		v.variablesOnLeft.Add(variable)
		variable.AddVariableOnRight(v)
	}
}

func (v *AbstractLocalVariable) AddVariableOnRight(variable ILocalVariable) {
	if v.variablesOnRight == nil {
		v.variablesOnRight = util.NewSet[ILocalVariable]()
		v.variablesOnRight.Add(variable)
		variable.AddVariableOnLeft(v)
	} else if !v.variablesOnRight.Contains(variable) {
		v.variablesOnRight.Add(variable)
		variable.AddVariableOnLeft(v)
	}
}

func (v *AbstractLocalVariable) IsPrimitiveLocalVariable() bool {
	return false
}

func (v *AbstractLocalVariable) LocalVariable() ILocalVariableReference {
	return nil
}

func (v *AbstractLocalVariable) SetLocalVariable(_ ILocalVariableReference) {

}

func (v *AbstractLocalVariable) String() string {
	return "AbstractLocalVariable{}"
}

type AbstractNopLocalVariableVisitor struct {
}

func (v *AbstractNopLocalVariableVisitor) VisitGenericLocalVariable(_ *GenericLocalVariable) {
}
func (v *AbstractNopLocalVariableVisitor) VisitObjectLocalVariable(_ *ObjectLocalVariable) {
}
func (v *AbstractNopLocalVariableVisitor) VisitPrimitiveLocalVariable(_ *PrimitiveLocalVariable) {
}
