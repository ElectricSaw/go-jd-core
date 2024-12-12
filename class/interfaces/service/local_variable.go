package service

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
)

type IFrame interface {
	Statements() intmod.IStatements
	AddLocalVariable(lv ILocalVariable)
	LocalVariable(index int) ILocalVariableReference
	Parent() IFrame
	SetExceptionLocalVariable(e ILocalVariable)
	MergeLocalVariable(typeBounds map[string]intmod.IType, localVariableMaker ILocalVariableMaker, lv ILocalVariable)
	RemoveLocalVariable(lv ILocalVariable)
	AddChild(child IFrame)
	Close()
	CreateNames(parentNames []string)
	UpdateLocalVariableInForStatements(typeMaker ITypeMaker)
	CreateDeclarations(containsLineNumber bool)
	AddIndex() int
}

type IRootFrame interface {
	IFrame
	CreateDeclarations(containsLineNumber bool)
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
	Update(index, offset int, typ intmod.IObjectType)
	Update2(index, offset int, typ intmod.IGenericType)
	Initialize(rootFrame IFrame) []ILocalVariable
}
