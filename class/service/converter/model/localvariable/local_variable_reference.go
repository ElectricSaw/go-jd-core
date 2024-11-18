package localvariable

import _type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"

type ILocalVariableReference interface {
	LocalVariable() ILocalVariableReference
	SetLocalVariable(localVariable ILocalVariableReference)
}

type ILocalVariable interface {
	ILocalVariableReference

	Frame() *Frame
	Next() ILocalVariable
	IsDeclared() bool
	Index() int
	FromOffset() int
	ToOffset() int
	Type() _type.IType
	Name() string
	Dimension() int

	SetFrame(frame *Frame)
	SetNext(lv ILocalVariable)
	SetDeclared(declared bool)
	SetFromOffset(fromOffset int)
	SetToOffset(offset int)
	SetToOffsetWithForce(offset int, force bool)
	SetName(name string)

	Accept(visitor LocalVariableVisitor)
	References() []ILocalVariable
	AddReference(reference ILocalVariableReference)
	IsAssignableFrom(typeBounds map[string]_type.IType, otherType _type.IType) bool
	TypeOnRight(typeBounds map[string]_type.IType, typ _type.IType)
	TypeOnLeft(typeBounds map[string]_type.IType, typ _type.IType)
	IsAssignableFromWithVariable(typeBounds map[string]_type.IType, variable ILocalVariable) bool
	VariableOnRight(typeBounds map[string]_type.IType, variable ILocalVariable)
	VariableOnLeft(typeBounds map[string]_type.IType, variable ILocalVariable)
	FireChangeEvent(typeBounds map[string]_type.IType)

	AddVariableOnLeft(variable ILocalVariable)
	AddVariableOnRight(variable ILocalVariable)
	IsPrimitiveLocalVariable() bool
}

func NewAbstractLocalVariable(index, offset int, name string) *AbstractLocalVariable {
	return NewAbstractLocalVariableWithAll(index, offset, name, offset == 0)
}

func NewAbstractLocalVariableWithAll(index, offset int, name string, declared bool) *AbstractLocalVariable {
	return &AbstractLocalVariable{
		declared:         declared,
		index:            index,
		fromOffset:       offset,
		toOffset:         offset,
		name:             name,
		references:       make([]ILocalVariable, 0),
		variablesOnRight: make([]ILocalVariable, 0),
		variablesOnLeft:  make([]ILocalVariable, 0),
	}
}

type AbstractLocalVariable struct {
	frame            *Frame
	next             ILocalVariable
	declared         bool
	index            int
	fromOffset       int
	toOffset         int
	name             string
	references       []ILocalVariable
	variablesOnRight []ILocalVariable
	variablesOnLeft  []ILocalVariable
}

func (v *AbstractLocalVariable) Frame() *Frame {
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

func (v *AbstractLocalVariable) Type() _type.IType {
	return nil
}

func (v *AbstractLocalVariable) Name() string {
	return v.name
}

func (v *AbstractLocalVariable) Dimension() int {
	return 0
}

func (v *AbstractLocalVariable) SetFrame(frame *Frame) {
	v.frame = frame
}

func (v *AbstractLocalVariable) SetNext(lv ILocalVariable) {
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

func (v *AbstractLocalVariable) Accept(visitor LocalVariableVisitor) {
}

func (v *AbstractLocalVariable) References() []ILocalVariable {
	return v.references
}

func (v *AbstractLocalVariable) AddReference(reference ILocalVariableReference) {
	v.references = append(v.references, reference.(ILocalVariable))
}

func (v *AbstractLocalVariable) IsAssignableFrom(typeBounds map[string]_type.IType, otherType _type.IType) bool {
	return false
}

func (v *AbstractLocalVariable) TypeOnRight(typeBounds map[string]_type.IType, typ _type.IType) {
}

func (v *AbstractLocalVariable) TypeOnLeft(typeBounds map[string]_type.IType, typ _type.IType) {
}

func (v *AbstractLocalVariable) IsAssignableFromWithVariable(typeBounds map[string]_type.IType, variable ILocalVariable) bool {
	return false
}

func (v *AbstractLocalVariable) VariableOnRight(typeBounds map[string]_type.IType, variable ILocalVariable) {
}

func (v *AbstractLocalVariable) VariableOnLeft(typeBounds map[string]_type.IType, variable ILocalVariable) {
}

func (v *AbstractLocalVariable) FireChangeEvent(typeBounds map[string]_type.IType) {
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

func (v *AbstractLocalVariable) AddVariableOnLeft(variable ILocalVariable) {
	if v.variablesOnLeft == nil {
		v.variablesOnLeft = make([]ILocalVariable, 0)
		v.variablesOnLeft = append(v.variablesOnLeft, variable)
		variable.AddVariableOnRight(v)
	} else if !containsLv(v.variablesOnLeft, variable) {
		v.variablesOnLeft = append(v.variablesOnLeft, variable)
		variable.AddVariableOnRight(v)
	}
}

func (v *AbstractLocalVariable) AddVariableOnRight(variable ILocalVariable) {
	if v.variablesOnRight == nil {
		v.variablesOnRight = make([]ILocalVariable, 0)
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

func (v *AbstractLocalVariable) LocalVariable() ILocalVariableReference {
	return nil
}

func (v *AbstractLocalVariable) SetLocalVariable(localVariable ILocalVariableReference) {

}

type LocalVariableVisitor interface {
	VisitGenericLocalVariable(localVariable *GenericLocalVariable)
	VisitObjectLocalVariable(localVariable *ObjectLocalVariable)
	VisitPrimitiveLocalVariable(localVariable *PrimitiveLocalVariable)
}

type AbstractNopLocalVariableVisitor struct {
}

func (v *AbstractNopLocalVariableVisitor) VisitGenericLocalVariable(localVariable *GenericLocalVariable) {
}
func (v *AbstractNopLocalVariableVisitor) VisitObjectLocalVariable(localVariable *ObjectLocalVariable) {
}
func (v *AbstractNopLocalVariableVisitor) VisitPrimitiveLocalVariable(localVariable *PrimitiveLocalVariable) {
}

func containsLv(variables []ILocalVariable, variable ILocalVariable) bool {
	for _, v := range variables {
		if v == variable {
			return true
		}
	}
	return false
}
