package localvariable

import _type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"

type ILocalVariableReference interface {
	LocalVariable() ILocalVariableReference
	SetLocalVariable(localVariable ILocalVariableReference)

	Frame() *Frame
	Next() ILocalVariableReference
	IsDeclared() bool
	Index() int
	FromOffset() int
	ToOffset() int
	Type() _type.IType
	Name() string
	Dimension() int

	SetFrame(frame *Frame)
	SetNext(lv ILocalVariableReference)
	SetDeclared(declared bool)
	SetFromOffset(fromOffset int)
	SetToOffset(offset int)
	SetToOffsetWithForce(offset int, force bool)
	SetName(name string)

	Accept(visitor LocalVariableVisitor)
	References() []ILocalVariableReference
	AddReference(reference ILocalVariableReference)
	IsAssignableFrom(typeBounds map[string]_type.IType, otherType _type.IType) bool
	TypeOnRight(typeBounds map[string]_type.IType, typ _type.IType)
	TypeOnLeft(typeBounds map[string]_type.IType, typ _type.IType)
	IsAssignableFromWithVariable(typeBounds map[string]_type.IType, variable ILocalVariableReference) bool
	VariableOnRight(typeBounds map[string]_type.IType, variable ILocalVariableReference)
	VariableOnLeft(typeBounds map[string]_type.IType, variable ILocalVariableReference)
	FireChangeEvent(typeBounds map[string]_type.IType)
	AddVariableOnLeft(variable ILocalVariableReference)
	AddVariableOnRight(variable ILocalVariableReference)
	IsPrimitiveLocalVariable() bool
}

func NewAbstractLocalVariable(index, offset int, name string) *AbstractLocalVariable {
	return NewAbstractLocalVariableWithAll(index, offset, name, offset == 0)
}

func NewAbstractLocalVariableWithAll(index, offset int, name string, declared bool) *AbstractLocalVariable {
	return &AbstractLocalVariable{
		declared:   declared,
		index:      index,
		fromOffset: offset,
		toOffset:   offset,
		name:       name,
		references: make([]ILocalVariableReference, 0),
	}
}

type AbstractLocalVariable struct {
	frame            *Frame
	next             ILocalVariableReference
	declared         bool
	index            int
	fromOffset       int
	toOffset         int
	name             string
	references       []ILocalVariableReference
	variablesOnRight []ILocalVariableReference
	variablesOnLeft  []ILocalVariableReference
}

func (v *AbstractLocalVariable) Frame() *Frame {
	return v.frame
}

func (v *AbstractLocalVariable) SetFrame(frame *Frame) {
	v.frame = frame
}

func (v *AbstractLocalVariable) Next() ILocalVariableReference {
	return v.next
}

func (v *AbstractLocalVariable) SetNext(lv ILocalVariableReference) {
	v.next = lv
}

func (v *AbstractLocalVariable) IsDeclared() bool {
	return v.declared
}

func (v *AbstractLocalVariable) SetDeclared(declared bool) {
	v.declared = declared
}

func (v *AbstractLocalVariable) Index() int {
	return v.index
}

func (v *AbstractLocalVariable) FromOffset() int {
	return v.fromOffset
}

func (v *AbstractLocalVariable) SetFromOffset(fromOffset int) {
	v.fromOffset = fromOffset
}

func (v *AbstractLocalVariable) ToOffset() int {
	return v.toOffset
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

func (v *AbstractLocalVariable) Type() _type.IType {
	return nil
}

func (v *AbstractLocalVariable) Name() string {
	return v.name
}

func (v *AbstractLocalVariable) SetName(name string) {
	v.name = name
}

func (v *AbstractLocalVariable) Dimension() int {
	return 0
}

func (v *AbstractLocalVariable) Accept(visitor LocalVariableVisitor) {

}

func (v *AbstractLocalVariable) References() []ILocalVariableReference {
	return v.references
}

func (v *AbstractLocalVariable) AddReference(reference ILocalVariableReference) {
	v.references = append(v.references, reference)
}

/**
 * Determines if the local variable represented by v object is either the same as, or is a super type variable
 * of, the local variable represented by the specified parameter.
 */
func (v *AbstractLocalVariable) IsAssignableFrom(typeBounds map[string]_type.IType, otherType _type.IType) bool {
	return false
}

func (v *AbstractLocalVariable) TypeOnRight(typeBounds map[string]_type.IType, typ _type.IType) {
}

func (v *AbstractLocalVariable) TypeOnLeft(typeBounds map[string]_type.IType, typ _type.IType) {
}

func (v *AbstractLocalVariable) IsAssignableFromWithVariable(typeBounds map[string]_type.IType, variable ILocalVariableReference) bool {
	return false
}

func (v *AbstractLocalVariable) VariableOnRight(typeBounds map[string]_type.IType, variable ILocalVariableReference) {

}

func (v *AbstractLocalVariable) VariableOnLeft(typeBounds map[string]_type.IType, variable ILocalVariableReference) {

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

func (v *AbstractLocalVariable) AddVariableOnLeft(variable ILocalVariableReference) {
	if v.variablesOnLeft == nil {
		v.variablesOnLeft = make([]ILocalVariableReference, 0)
		v.variablesOnLeft = append(v.variablesOnLeft, variable)
		variable.AddVariableOnRight(v)
	} else if !containsLv(v.variablesOnLeft, variable) {
		v.variablesOnLeft = append(v.variablesOnLeft, variable)
		variable.AddVariableOnRight(v)
	}
}

func (v *AbstractLocalVariable) AddVariableOnRight(variable ILocalVariableReference) {
	if v.variablesOnRight == nil {
		v.variablesOnRight = make([]ILocalVariableReference, 0)
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

func containsLv(variables []ILocalVariableReference, variable ILocalVariableReference) bool {
	for _, v := range variables {
		if v == variable {
			return true
		}
	}
	return false
}
