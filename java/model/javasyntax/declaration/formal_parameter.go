package declaration

import (
	"bitbucket.org/coontec/javaClass/java/model/javasyntax/reference"
	_type "bitbucket.org/coontec/javaClass/java/model/javasyntax/type"
	"fmt"
)

func NewFormalParameter(typ _type.IType, name string) *FormalParameter {
	return &FormalParameter{
		typ:  typ,
		name: name,
	}
}

func NewFormalParameter2(annotationReferences reference.IAnnotationReference, typ _type.IType, name string) *FormalParameter {
	return &FormalParameter{
		annotationReferences: annotationReferences,
		typ:                  typ,
		name:                 name,
	}
}

func NewFormalParameter3(typ _type.IType, varargs bool, name string) *FormalParameter {
	return &FormalParameter{
		typ:     typ,
		varargs: varargs,
		name:    name,
	}
}

func NewFormalParameter4(annotationReferences reference.IAnnotationReference, typ _type.IType, varargs bool, name string) *FormalParameter {
	return &FormalParameter{
		annotationReferences: annotationReferences,
		typ:                  typ,
		varargs:              varargs,
		name:                 name,
	}
}

type FormalParameter struct {
	annotationReferences reference.IAnnotationReference
	final                bool
	typ                  _type.IType
	varargs              bool
	name                 string
}

func (d *FormalParameter) GetAnnotationReferences() reference.IAnnotationReference {
	return d.annotationReferences
}

func (d *FormalParameter) IsFinal() bool {
	return d.final
}

func (d *FormalParameter) SetFinal(final bool) {
	d.final = final
}

func (d *FormalParameter) GetType() _type.IType {
	return d.typ
}

func (d *FormalParameter) IsVarargs() bool {
	return d.varargs
}

func (d *FormalParameter) GetName() string {
	return d.name
}

func (d *FormalParameter) SetName(name string) {
	d.name = name
}

func (d *FormalParameter) Accept(visitor DeclarationVisitor) {
	visitor.VisitFormalParameter(d)
}

func (d *FormalParameter) String() string {
	msg := "FormalParameter{"

	if d.annotationReferences != nil {
		msg += fmt.Sprintf("%v ", d.annotationReferences)
	}

	if d.varargs {
		msg += fmt.Sprintf("%v... ", d.typ.CreateType(d.typ.GetDimension()-1))
	} else {
		msg += fmt.Sprintf("%v ", d.typ)
	}
	msg += "}"

	return msg
}
