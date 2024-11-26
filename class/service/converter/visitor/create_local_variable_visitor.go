package visitor

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
	"bitbucket.org/coontec/go-jd-core/class/service/converter/model/localvariable"
	"bitbucket.org/coontec/go-jd-core/class/service/converter/utils"
)

func NewCreateLocalVariableVisitor(typeMaker *utils.TypeMaker) *CreateLocalVariableVisitor {
	return &CreateLocalVariableVisitor{
		typeMaker: typeMaker,
	}
}

type CreateLocalVariableVisitor struct {
	_type.AbstractNopTypeArgumentVisitor

	typeMaker     *utils.TypeMaker
	index         int
	offset        int
	localVariable intsrv.ILocalVariable
}

func (v *CreateLocalVariableVisitor) Init(index, offset int) {
	v.index = index
	v.offset = offset
}

func (v *CreateLocalVariableVisitor) LocalVariable() intsrv.ILocalVariable {
	return v.localVariable
}

func (v *CreateLocalVariableVisitor) VisitPrimitiveType(t intmod.IPrimitiveType) {
	if t.Dimension() == 0 {
		v.localVariable = localvariable.NewPrimitiveLocalVariable(v.index, v.offset, t, "")
	} else {
		v.localVariable = localvariable.NewObjectLocalVariable(v.typeMaker, v.index, v.offset, t, "")
	}
}

func (v *CreateLocalVariableVisitor) VisitObjectType(t intmod.IObjectType) {
	v.localVariable = localvariable.NewObjectLocalVariable(v.typeMaker, v.index, v.offset, t.(intmod.IType), "")
}

func (v *CreateLocalVariableVisitor) VisitInnerObjectType(t intmod.IInnerObjectType) {
	v.localVariable = localvariable.NewObjectLocalVariable(v.typeMaker, v.index, v.offset, t.(intmod.IType), "")
}

func (v *CreateLocalVariableVisitor) VisitGenericType(t intmod.IGenericType) {
	v.localVariable = localvariable.NewGenericLocalVariable(v.index, v.offset, t)
}

func (v *CreateLocalVariableVisitor) VisitGenericLocalVariable(lv intsrv.IGenericLocalVariable) {
	v.localVariable = localvariable.NewGenericLocalVariable(v.index, v.offset, lv.Type().(intmod.IGenericType))
}

func (v *CreateLocalVariableVisitor) VisitObjectLocalVariable(lv intsrv.IObjectLocalVariable) {
	v.localVariable = localvariable.NewObjectLocalVariable3(v.typeMaker, v.index, v.offset, lv)
}

func (v *CreateLocalVariableVisitor) VisitPrimitiveLocalVariable(lv intsrv.IPrimitiveLocalVariable) {
	if lv.Dimension() == 0 {
		v.localVariable = localvariable.NewPrimitiveLocalVariableWithVar(v.index, v.offset, lv)
	} else {
		v.localVariable = localvariable.NewObjectLocalVariable(v.typeMaker, v.index, v.offset, lv.Type(), "")
	}
}
