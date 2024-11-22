package visitor

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
	"bitbucket.org/coontec/go-jd-core/class/service/converter/model/localvariable"
	"bitbucket.org/coontec/go-jd-core/class/service/converter/utils"
)

func NewCreateParameterVisitor(typeMaker *utils.TypeMaker) *CreateParameterVisitor {
	return &CreateParameterVisitor{
		typeMaker: typeMaker,
	}
}

type CreateParameterVisitor struct {
	_type.AbstractNopTypeArgumentVisitor

	typeMaker     *utils.TypeMaker
	index         int
	name          string
	localVariable intsrv.ILocalVariable
}

func (v *CreateParameterVisitor) Init(index int, name string) {
	v.index = index
	v.name = name
}

func (v *CreateParameterVisitor) LocalVariable() intsrv.ILocalVariable {
	return v.localVariable
}

func (v *CreateParameterVisitor) VisitPrimitiveType(t intmod.IPrimitiveType) {
	if t.Dimension() == 0 {
		v.localVariable = localvariable.NewPrimitiveLocalVariable(v.index, 0, t, v.name)
	} else {
		v.localVariable = localvariable.NewObjectLocalVariable(v.typeMaker, v.index, 0, t, v.name)
	}
}

func (v *CreateParameterVisitor) VisitObjectType(t intmod.IObjectType) {
	v.localVariable = localvariable.NewObjectLocalVariable(v.typeMaker, v.index, 0, t.(intmod.IType), v.name)
}

func (v *CreateParameterVisitor) VisitInnerObjectType(t intmod.IInnerObjectType) {
	v.localVariable = localvariable.NewObjectLocalVariable(v.typeMaker, v.index, 0, t.(intmod.IType), v.name)
}

func (v *CreateParameterVisitor) VisitGenericType(t intmod.IGenericType) {
	v.localVariable = localvariable.NewGenericLocalVariableWithAll(v.index, 0, t, v.name)
}
