package declaration

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewFormalParameter(typ Type, name string) FormalParameter {
	return NewFormalParameter4(nil, typ, false, name)
}

func NewFormalParameter2(annotationReferences AnnotationReference, typ Type, name string) FormalParameter {
	return NewFormalParameter4(annotationReferences, typ, false, name)
}

func NewFormalParameter3(typ Type, varargs bool, name string) FormalParameter {
	return NewFormalParameter4(nil, typ, varargs, name)
}

func NewFormalParameter4(annotationReferences AnnotationReference, typ Type, varargs bool, name string) FormalParameter {
	p := &FormalParameter{
		annotationReferences: annotationReferences,
		typ:                  typ,
		varargs:              varargs,
		name:                 name,
	}
	p.SetValue(p)
	return p
}

type FormalParameter struct {
	util.DefaultBase[FormalParameter]

	annotationReferences AnnotationReference
	final                bool
	typ                  Type
	varargs              bool
	name                 string
}

func (d *FormalParameter) AnnotationReferences() AnnotationReference {
	return d.annotationReferences
}

func (d *FormalParameter) IsFinal() bool {
	return d.final
}

func (d *FormalParameter) SetFinal(final bool) {
	d.final = final
}

func (d *FormalParameter) Type() Type {
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

func (d *FormalParameter) AcceptDeclaration(visitor DeclarationVisitor) {
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
