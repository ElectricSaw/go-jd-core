package service

import intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"

type IFrame interface {
	// TODO: AddLocalVariable(ILocalVariable) 메소드는 호출을 위해 심시 추가.
	AddLocalVariable(ILocalVariable)
}

type ILocalVariableReference interface {
	LocalVariable() ILocalVariableReference
	SetLocalVariable(localVariable ILocalVariableReference)
}

type ILocalVariable interface {
	ILocalVariableReference

	Frame() IFrame
	Next() ILocalVariable
	IsDeclared() bool
	Index() int
	FromOffset() int
	ToOffset() int
	Type() intmod.IType
	Name() string
	Dimension() int

	SetFrame(frame IFrame)
	SetNext(lv ILocalVariable)
	SetDeclared(declared bool)
	SetFromOffset(fromOffset int)
	SetToOffset(offset int)
	SetToOffsetWithForce(offset int, force bool)
	SetName(name string)

	Accept(visitor ILocalVariableVisitor)
	References() []ILocalVariable
	AddReference(reference ILocalVariableReference)
	IsAssignableFrom(typeBounds map[string]intmod.IType, otherType intmod.IType) bool
	TypeOnRight(typeBounds map[string]intmod.IType, typ intmod.IType)
	TypeOnLeft(typeBounds map[string]intmod.IType, typ intmod.IType)
	IsAssignableFromWithVariable(typeBounds map[string]intmod.IType, variable ILocalVariable) bool
	VariableOnRight(typeBounds map[string]intmod.IType, variable ILocalVariable)
	VariableOnLeft(typeBounds map[string]intmod.IType, variable ILocalVariable)
	FireChangeEvent(typeBounds map[string]intmod.IType)

	AddVariableOnLeft(variable ILocalVariable)
	AddVariableOnRight(variable ILocalVariable)
	IsPrimitiveLocalVariable() bool
}

type ILocalVariableVisitor interface {
	VisitGenericLocalVariable(localVariable IGenericLocalVariable)
	VisitObjectLocalVariable(localVariable IObjectLocalVariable)
	VisitPrimitiveLocalVariable(localVariable IPrimitiveLocalVariable)
}

type IGenericLocalVariable interface {
	ILocalVariable

	Type() intmod.IType
	SetType(typ intmod.IGenericType)
	Dimension() int
	Accept(visitor ILocalVariableVisitor)
	String() string
	IsAssignableFrom(typeBounds map[string]intmod.IType, otherType intmod.IType) bool
	TypeOnRight(typeBounds map[string]intmod.IType, typ intmod.IType)
	TypeOnLeft(typeBounds map[string]intmod.IType, typ intmod.IType)
	IsAssignableFromWithVariable(typeBounds map[string]intmod.IType, variable ILocalVariable) bool
	VariableOnRight(typeBounds map[string]intmod.IType, variable ILocalVariable)
	VariableOnLeft(typeBounds map[string]intmod.IType, variable ILocalVariable)
}

type IObjectLocalVariable interface {
	ILocalVariable

	Type() intmod.IType
	SetType(typeBounds map[string]intmod.IType, t intmod.IType)
	Dimension() int
	Accept(visitor ILocalVariableVisitor)
	String() string
	IsAssignableFrom(typeBounds map[string]intmod.IType, typ intmod.IType) bool
	TypeOnRight(typeBounds map[string]intmod.IType, typ intmod.IType)
	TypeOnLeft(typeBounds map[string]intmod.IType, typ intmod.IType)
	IsAssignableFromWithVariable(typeBounds map[string]intmod.IType, variable ILocalVariable) bool
	VariableOnRight(typeBounds map[string]intmod.IType, variable ILocalVariable)
	VariableOnLeft(typeBounds map[string]intmod.IType, variable ILocalVariable)
}

type IPrimitiveLocalVariable interface {
	ILocalVariable

	Flags() int
	SetFlags(flags int)
	Type() intmod.IType
	Dimension() int
	SetType(typ intmod.IPrimitiveType)
	Accept(visitor ILocalVariableVisitor)
	String() string
	IsAssignableFrom(typeBounds map[string]intmod.IType, typ intmod.IType) bool
	TypeOnRight(typeBounds map[string]intmod.IType, typ intmod.IType)
	TypeOnLeft(typeBounds map[string]intmod.IType, typ intmod.IType)
	IsAssignableFromWithVariable(typeBounds map[string]intmod.IType, variable ILocalVariable) bool
	VariableOnRight(typeBounds map[string]intmod.IType, variable ILocalVariable)
	VariableOnLeft(typeBounds map[string]intmod.IType, variable ILocalVariable)
	IsPrimitiveLocalVariable() bool
}

type ILocalVariableSet interface {
	Add(index int, newLV ILocalVariable)
	Root(index int) ILocalVariable
	Remove(index, offset int) ILocalVariable
	Get(index, offset int) ILocalVariable
	IsEmpty() bool
}
