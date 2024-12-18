package visitor

import (
	intcls "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/classpath"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	_type "github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/type"
)

func NewUpdateTypeVisitor(localVariableSet intsrv.ILocalVariableSet) intsrv.IUpdateTypeVisitor {
	return &UpdateTypeVisitor{
		updateClassTypeArgumentsVisitor: NewUpdateClassTypeArgumentsVisitor(),
		localVariableSet:                localVariableSet,
	}
}

type UpdateTypeVisitor struct {
	_type.AbstractNopTypeArgumentVisitor

	updateClassTypeArgumentsVisitor intsrv.IUpdateClassTypeArgumentsVisitor
	localVariableSet                intsrv.ILocalVariableSet
	localVariableType               intcls.ILocalVariableType
}

func (v *UpdateTypeVisitor) SetLocalVariableType(localVariableType intcls.ILocalVariableType) {
	v.localVariableType = localVariableType
}

func (v *UpdateTypeVisitor) VisitObjectType(t intmod.IObjectType) {
	v.localVariableSet.Update(v.localVariableType.Index(), v.localVariableType.StartPc(), v.updateType(t))
}

func (v *UpdateTypeVisitor) VisitInnerObjectType(t intmod.IInnerObjectType) {
	v.localVariableSet.Update(v.localVariableType.Index(), v.localVariableType.StartPc(), v.updateType(t))
}

func (v *UpdateTypeVisitor) VisitGenericType(t intmod.IGenericType) {
	v.localVariableSet.Update2(v.localVariableType.Index(), v.localVariableType.StartPc(), t)
}

func (v *UpdateTypeVisitor) updateType(t intmod.IObjectType) intmod.IObjectType {
	typeArguments := t.TypeArguments()

	if typeArguments != nil {
		v.updateClassTypeArgumentsVisitor.Init()
		typeArguments.AcceptTypeArgumentVisitor(v.updateClassTypeArgumentsVisitor)

		if typeArguments != v.updateClassTypeArgumentsVisitor.TypeArgument() {
			t = t.CreateTypeWithArgs(v.updateClassTypeArgumentsVisitor.TypeArgument())
		}
	}

	return t
}
