package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/declaration"
	"fmt"
)

func NewClassFileFormalParameter(localVariable intsrv.ILocalVariable) intsrv.IClassFileFormalParameter {
	return NewClassFileFormalParameter3(nil, localVariable, false)
}

func NewClassFileFormalParameter2(localVariable intsrv.ILocalVariable, varargs bool) intsrv.IClassFileFormalParameter {
	return NewClassFileFormalParameter3(nil, localVariable, varargs)
}

func NewClassFileFormalParameter3(annotationReferences intmod.IAnnotationReference,
	localVariable intsrv.ILocalVariable, varargs bool) intsrv.IClassFileFormalParameter {
	p := &ClassFileFormalParameter{
		FormalParameter: *declaration.NewFormalParameter4(annotationReferences, nil, varargs, "").(*declaration.FormalParameter),
		localVariable:   localVariable,
	}
	p.SetValue(p)
	return p
}

type ClassFileFormalParameter struct {
	declaration.FormalParameter

	localVariable intsrv.ILocalVariable
}

func (p *ClassFileFormalParameter) Type() intmod.IType {
	return p.localVariable.Type()
}

func (p *ClassFileFormalParameter) Name() string {
	return p.localVariable.Name()
}

func (p *ClassFileFormalParameter) LocalVariable() intsrv.ILocalVariableReference {
	return p.localVariable
}

func (p *ClassFileFormalParameter) SetLocalVariable(localVariable intsrv.ILocalVariableReference) {
	p.localVariable = localVariable.(intsrv.ILocalVariable)
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
