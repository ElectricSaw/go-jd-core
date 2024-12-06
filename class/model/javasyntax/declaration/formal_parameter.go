package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
	"fmt"
)

func NewFormalParameter(typ intmod.IType, name string) intmod.IFormalParameter {
	return NewFormalParameter4(nil, typ, false, name)
}

func NewFormalParameter2(annotationReferences intmod.IAnnotationReference, typ intmod.IType, name string) intmod.IFormalParameter {
	return NewFormalParameter4(annotationReferences, typ, false, name)
}

func NewFormalParameter3(typ intmod.IType, varargs bool, name string) intmod.IFormalParameter {
	return NewFormalParameter4(nil, typ, varargs, name)
}

func NewFormalParameter4(annotationReferences intmod.IAnnotationReference, typ intmod.IType, varargs bool, name string) intmod.IFormalParameter {
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
	util.DefaultBase[intmod.IFormalParameter]

	annotationReferences intmod.IAnnotationReference
	final                bool
	typ                  intmod.IType
	varargs              bool
	name                 string
}

func (d *FormalParameter) AnnotationReferences() intmod.IAnnotationReference {
	return d.annotationReferences
}

func (d *FormalParameter) IsFinal() bool {
	return d.final
}

func (d *FormalParameter) SetFinal(final bool) {
	d.final = final
}

func (d *FormalParameter) Type() intmod.IType {
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

func (d *FormalParameter) AcceptDeclaration(visitor intmod.IDeclarationVisitor) {
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
