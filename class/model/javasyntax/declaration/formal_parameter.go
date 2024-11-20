package declaration

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewFormalParameter(typ intsyn.IType, name string) intsyn.IFormalParameter {
	return &FormalParameter{
		typ:  typ,
		name: name,
	}
}

func NewFormalParameter2(annotationReferences intsyn.IAnnotationReference, typ intsyn.IType, name string) intsyn.IFormalParameter {
	return &FormalParameter{
		annotationReferences: annotationReferences,
		typ:                  typ,
		name:                 name,
	}
}

func NewFormalParameter3(typ intsyn.IType, varargs bool, name string) intsyn.IFormalParameter {
	return &FormalParameter{
		typ:     typ,
		varargs: varargs,
		name:    name,
	}
}

func NewFormalParameter4(annotationReferences intsyn.IAnnotationReference, typ intsyn.IType, varargs bool, name string) intsyn.IFormalParameter {
	return &FormalParameter{
		annotationReferences: annotationReferences,
		typ:                  typ,
		varargs:              varargs,
		name:                 name,
	}
}

type FormalParameter struct {
	annotationReferences intsyn.IAnnotationReference
	final                bool
	typ                  intsyn.IType
	varargs              bool
	name                 string
}

func (d *FormalParameter) AnnotationReferences() intsyn.IAnnotationReference {
	return d.annotationReferences
}

func (d *FormalParameter) IsFinal() bool {
	return d.final
}

func (d *FormalParameter) SetFinal(final bool) {
	d.final = final
}

func (d *FormalParameter) Type() intsyn.IType {
	return d.typ
}

func (d *FormalParameter) IsVarargs() bool {
	return d.varargs
}

func (d *FormalParameter) Name() string {
	return d.name
}

func (d *FormalParameter) SetName(name string) {
	d.name = name
}

func (d *FormalParameter) Accept(visitor intsyn.IDeclarationVisitor) {
	visitor.VisitFormalParameter(d)
}

func (d *FormalParameter) String() string {
	msg := "FormalParameter{"

	if d.annotationReferences != nil {
		msg += fmt.Sprintf("%v ", d.annotationReferences)
	}

	if d.varargs {
		msg += fmt.Sprintf("%v... ", d.typ.CreateType(d.typ.Dimension()-1))
	} else {
		msg += fmt.Sprintf("%v ", d.typ)
	}
	msg += "}"

	return msg
}
