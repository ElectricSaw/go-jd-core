package declaration

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/service/converter/model/localvariable"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewClassFileFormalParameter(localVariable localvariable.ILocalVariable) ClassFileFormalParameter {
	return NewClassFileFormalParameter3(nil, localVariable, false)
}

func NewClassFileFormalParameter2(localVariable localvariable.ILocalVariable, varargs bool) ClassFileFormalParameter {
	return NewClassFileFormalParameter3(nil, localVariable, varargs)
}

func NewClassFileFormalParameter3(annotationReferences *model.AnnotationReference,
	localVariable localvariable.ILocalVariable, varargs bool) ClassFileFormalParameter {
	p := ClassFileFormalParameter{
		DefaultBase:          *util.NewDefaultBase[model.IFormalParameter]().(*util.DefaultBase[model.IFormalParameter]),
		AnnotationReferences: annotationReferences,
		Varargs:              varargs,
		LocalVariable:        localVariable,
	}
	p.SetValue(&p)
	return p
}

type ClassFileFormalParameter struct {
	util.DefaultBase[model.IFormalParameter]

	AnnotationReferences *model.AnnotationReference
	Final                bool
	Type                 model.IType
	Varargs              bool
	Name                 string
	LocalVariable        localvariable.ILocalVariable
}

func (p *ClassFileFormalParameter) GetAnnotationReferences() *model.AnnotationReference {
	return p.AnnotationReferences
}

func (p *ClassFileFormalParameter) IsFinal() bool {
	return p.Final
}

func (p *ClassFileFormalParameter) GetType() model.IType {
	return p.Type
}

func (p *ClassFileFormalParameter) IsVarargs() bool {
	return p.Varargs
}

func (p *ClassFileFormalParameter) GetName() string {
	return p.Name
}

func (p *ClassFileFormalParameter) GetLocalVariable() localvariable.ILocalVariableReference {
	return p.LocalVariable
}

func (p *ClassFileFormalParameter) SetLocalVariable(localVariable localvariable.ILocalVariableReference) {
	p.LocalVariable = localVariable.(localvariable.ILocalVariable)
}

func (p *ClassFileFormalParameter) AcceptDeclaration(visitor model.IDeclarationVisitor) {
	visitor.VisitFormalParameter(p)
}

func (p *ClassFileFormalParameter) String() string {
	s := "ClassFileFormalParameter{"

	if p.AnnotationReferences != nil {
		s += fmt.Sprintf("%s ", p.AnnotationReferences)
	}

	t := p.LocalVariable.Type()

	if p.Varargs {
		s += fmt.Sprintf("%s... ", t.CreateType(t.GetDimension()-1))
	} else {
		s += fmt.Sprintf("%s ", t)
	}

	s += fmt.Sprintf("%s}", p.LocalVariable.Name())

	return s
}
