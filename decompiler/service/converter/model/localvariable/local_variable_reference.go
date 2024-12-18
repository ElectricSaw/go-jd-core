package localvariable

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
)

func NewAbstractLocalVariable(index, offset int, name string) intsrv.ILocalVariable {
	return NewAbstractLocalVariableWithAll(index, offset, name, offset == 0)
}

func NewAbstractLocalVariableWithAll(index, offset int, name string, declared bool) intsrv.ILocalVariable {
	return &AbstractLocalVariable{
		declared:         declared,
		index:            index,
		fromOffset:       offset,
		toOffset:         offset,
		name:             name,
		references:       make([]intsrv.ILocalVariable, 0),
		variablesOnRight: make([]intsrv.ILocalVariable, 0),
		variablesOnLeft:  make([]intsrv.ILocalVariable, 0),
	}
}

type AbstractLocalVariable struct {
	frame            intsrv.IFrame
	next             intsrv.ILocalVariable
	declared         bool
	index            int
	fromOffset       int
	toOffset         int
	name             string
	references       []intsrv.ILocalVariable
	variablesOnRight []intsrv.ILocalVariable
	variablesOnLeft  []intsrv.ILocalVariable
}

func (v *AbstractLocalVariable) Frame() intsrv.IFrame {
	return v.frame
}

func (v *AbstractLocalVariable) Next() intsrv.ILocalVariable {
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

func (v *AbstractLocalVariable) Type() intmod.IType {
	return nil
}

func (v *AbstractLocalVariable) Name() string {
	return v.name
}

func (v *AbstractLocalVariable) Dimension() int {
	return 0
}

func (v *AbstractLocalVariable) SetFrame(frame intsrv.IFrame) {
	v.frame = frame
}

func (v *AbstractLocalVariable) SetNext(lv intsrv.ILocalVariable) {
	v.next = lv
}

func (v *AbstractLocalVariable) SetDeclared(declared bool) {
	v.declared = declared
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

func (v *AbstractLocalVariable) SetToOffsetWithForce(offset int, force bool) {
	v.toOffset = offset
}

func (v *AbstractLocalVariable) SetName(name string) {
	v.name = name
}

func (v *AbstractLocalVariable) Accept(_ intsrv.ILocalVariableVisitor) {
}

func (v *AbstractLocalVariable) References() []intsrv.ILocalVariable {
	return v.references
}

func (v *AbstractLocalVariable) AddReference(reference intsrv.ILocalVariableReference) {
	v.references = append(v.references, reference.(intsrv.ILocalVariable))
}

func (v *AbstractLocalVariable) IsAssignableFrom(_ map[string]intmod.IType, _ intmod.IType) bool {
	return false
}

func (v *AbstractLocalVariable) TypeOnRight(_ map[string]intmod.IType, _ intmod.IType) {
}

func (v *AbstractLocalVariable) TypeOnLeft(_ map[string]intmod.IType, _ intmod.IType) {
}

func (v *AbstractLocalVariable) IsAssignableFromWithVariable(_ map[string]intmod.IType, _ intsrv.ILocalVariable) bool {
	return false
}

func (v *AbstractLocalVariable) VariableOnRight(_ map[string]intmod.IType, _ intsrv.ILocalVariable) {
}

func (v *AbstractLocalVariable) VariableOnLeft(_ map[string]intmod.IType, _ intsrv.ILocalVariable) {
}

func (v *AbstractLocalVariable) FireChangeEvent(typeBounds map[string]intmod.IType) {
	if v.variablesOnLeft != nil {
		for _, variable := range v.variablesOnLeft {
			v.VariableOnRight(typeBounds, variable)
		}
	}
	if v.variablesOnRight != nil {
		for _, variable := range v.variablesOnRight {
			v.VariableOnLeft(typeBounds, variable)
		}
	}
}

func (v *AbstractLocalVariable) AddVariableOnLeft(variable intsrv.ILocalVariable) {
	if v.variablesOnLeft == nil {
		v.variablesOnLeft = make([]intsrv.ILocalVariable, 0)
		v.variablesOnLeft = append(v.variablesOnLeft, variable)
		variable.AddVariableOnRight(v)
	} else if !containsLv(v.variablesOnLeft, variable) {
		v.variablesOnLeft = append(v.variablesOnLeft, variable)
		variable.AddVariableOnRight(v)
	}
}

func (v *AbstractLocalVariable) AddVariableOnRight(variable intsrv.ILocalVariable) {
	if v.variablesOnRight == nil {
		v.variablesOnRight = make([]intsrv.ILocalVariable, 0)
		v.variablesOnRight = append(v.variablesOnRight, variable)
		variable.AddVariableOnLeft(v)
	} else if !containsLv(v.variablesOnRight, variable) {
		v.variablesOnRight = append(v.variablesOnRight, variable)
		variable.AddVariableOnLeft(v)
	}
}

func (v *AbstractLocalVariable) IsPrimitiveLocalVariable() bool {
	return false
}

func (v *AbstractLocalVariable) LocalVariable() intsrv.ILocalVariableReference {
	return nil
}

func (v *AbstractLocalVariable) SetLocalVariable(_ intsrv.ILocalVariableReference) {

}

type AbstractNopLocalVariableVisitor struct {
}

func (v *AbstractNopLocalVariableVisitor) VisitGenericLocalVariable(_ intsrv.IGenericLocalVariable) {
}
func (v *AbstractNopLocalVariableVisitor) VisitObjectLocalVariable(_ intsrv.IObjectLocalVariable) {
}
func (v *AbstractNopLocalVariableVisitor) VisitPrimitiveLocalVariable(_ intsrv.IPrimitiveLocalVariable) {
}

func containsLv(variables []intsrv.ILocalVariable, variable intsrv.ILocalVariable) bool {
	for _, v := range variables {
		if v == variable {
			return true
		}
	}
	return false
}
