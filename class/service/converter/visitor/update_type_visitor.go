package visitor

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/classfile/attribute"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
)

func NewUpdateTypeVisitor(localVariableSet intsrv.ILocalVariableSet) *UpdateTypeVisitor {
	return &UpdateTypeVisitor{
		localVariableSet: localVariableSet,
	}
}

type UpdateTypeVisitor struct {
	_type.AbstractNopTypeArgumentVisitor

	updateClassTypeArgumentsVisitor UpdateClassTypeArgumentsVisitor
	localVariableSet                intsrv.ILocalVariableSet
	localVariableType               attribute.LocalVariableType
}

func (v *UpdateTypeVisitor) SetLocalVariableType(localVariableType attribute.LocalVariableType) {
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
		typeArguments.AcceptTypeArgumentVisitor(&v.updateClassTypeArgumentsVisitor)

		if typeArguments != v.updateClassTypeArgumentsVisitor.TypeArgument() {
			t = t.CreateTypeWithArgs(v.updateClassTypeArgumentsVisitor.TypeArgument())
		}
	}

	return t
}
