package localvariable

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/service/converter/visitor"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

type IFrame interface {
	GetStatements() model.Statements
	GetLocalVariable(index int) ILocalVariableReference
	GetParent() IFrame
	AddLocalVariable(lv ILocalVariable)
	MergeLocalVariable(typeBounds map[string]model.IType, localVariableMaker *visitor.LocalVariableMaker, lv ILocalVariable)
	RemoveLocalVariable(lv ILocalVariable)
	AddChild(child IFrame)
	Close()
	CreateNames(parentNames []string)
	UpdateLocalVariableInForStatements(typeMaker *visitor.TypeMaker)
	CreateDeclarations(containsLineNumber bool)
	AddIndex() int
}

type ILocalVariableReference interface {
	LocalVariable() ILocalVariableReference
	SetLocalVariable(localVariable ILocalVariableReference)
}

type ILocalVariable interface {
	Frame() IFrame
	Next() ILocalVariable
	IsDeclared() bool
	Index() int
	FromOffset() int
	ToOffset() int
	Type() model.IType
	Name() string
	Dimension() int
	References() util.IList[ILocalVariable]

	SetFrame(frame IFrame)
	SetNext(localVariable ILocalVariable)
	SetDeclared(declared bool)
	SetIndex(index int)
	SetFromOffset(fromOffset int)
	SetToOffset(offset int)
	SetType(t model.IType)
	SetName(name string)
	SetDimension(dimension int)

	Accept(visitor ILocalVariableVisitor)
	AddReference(reference ILocalVariableReference)
	IsAssignableFrom(typeBounds map[string]model.IType, otherType model.IType) bool
	TypeOnRight(typeBounds map[string]model.IType, typ model.IType)
	TypeOnLeft(typeBounds map[string]model.IType, typ model.IType)
	IsAssignableFromWithVariable(typeBounds map[string]model.IType, variable ILocalVariable) bool
	VariableOnRight(typeBounds map[string]model.IType, variable ILocalVariable)
	VariableOnLeft(typeBounds map[string]model.IType, variable ILocalVariable)
	FireChangeEvent(typeBounds map[string]model.IType)

	AddVariableOnLeft(variable ILocalVariable)
	AddVariableOnRight(variable ILocalVariable)
	IsPrimitiveLocalVariable() bool

	LocalVariable() ILocalVariableReference
	SetLocalVariable(localVariable ILocalVariableReference)

	String() string
}

type ILocalVariableVisitor interface {
	VisitGenericLocalVariable(localVariable *GenericLocalVariable)
	VisitObjectLocalVariable(localVariable *ObjectLocalVariable)
	VisitPrimitiveLocalVariable(localVariable *PrimitiveLocalVariable)
}
