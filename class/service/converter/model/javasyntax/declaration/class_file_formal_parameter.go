package declaration

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/declaration"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/reference"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"bitbucket.org/coontec/javaClass/class/service/converter/model/localvariable"
	"fmt"
)

func NewClassFileFormalParameter(localVariable localvariable.ILocalVariable) *ClassFileFormalParameter {
	return &ClassFileFormalParameter{
		FormalParameter: *declaration.NewFormalParameter(nil, ""),
		localVariable:   localVariable,
	}
}

func NewClassFileFormalParameter2(localVariable localvariable.ILocalVariable, varargs bool) *ClassFileFormalParameter {
	return &ClassFileFormalParameter{
		FormalParameter: *declaration.NewFormalParameter3(nil, varargs, ""),
		localVariable:   localVariable,
	}
}

func NewClassFileFormalParameter3(annotationReferences reference.IAnnotationReference,
	localVariable localvariable.ILocalVariable, varargs bool) *ClassFileFormalParameter {
	return &ClassFileFormalParameter{
		FormalParameter: *declaration.NewFormalParameter4(annotationReferences, nil, varargs, ""),
		localVariable:   localVariable,
	}
}

type ClassFileFormalParameter struct {
	declaration.FormalParameter

	localVariable localvariable.ILocalVariable
}

func (p *ClassFileFormalParameter) Type() _type.IType {
	return p.localVariable.Type()
}

func (p *ClassFileFormalParameter) Name() string {
	return p.localVariable.Name()
}

func (p *ClassFileFormalParameter) LocalVariable() localvariable.ILocalVariableReference {
	return p.localVariable
}

func (p *ClassFileFormalParameter) SetLocalVariable(localVariable localvariable.ILocalVariableReference) {
	p.localVariable = localVariable.(localvariable.ILocalVariable)
}

func (p *ClassFileFormalParameter) String() string {
	s := "ClassFileFormalParameter{"

	if p.AnnotationReferences() != nil {
		s += fmt.Sprintf("%s ", p.AnnotationReferences())
	}

	t := p.localVariable.Type()

	if p.IsVarargs() {
		s += fmt.Sprintf("%s... ", t.CreateType(t.Dimension()-1))
	} else {
		s += fmt.Sprintf("%s ", t)
	}

	s += fmt.Sprintf("%s}", p.localVariable.Name())

	return s
}
