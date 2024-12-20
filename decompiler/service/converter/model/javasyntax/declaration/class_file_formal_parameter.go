package declaration

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/declaration"
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

func (d *ClassFileFormalParameter) AcceptDeclaration(visitor intmod.IDeclarationVisitor) {
	visitor.VisitFormalParameter(d)
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
