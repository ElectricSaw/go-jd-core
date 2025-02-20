package declaration

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewClassFileAnnotationDeclaration(annotationReferences *model.AnnotationReference,
	flags int, internalTypeName string, name string, bodyDeclaration *ClassFileBodyDeclaration) ClassFileAnnotationDeclaration {
	d := ClassFileAnnotationDeclaration{
		DefaultBase:          *util.NewDefaultBase[model.IMemberDeclaration]().(*util.DefaultBase[model.IMemberDeclaration]),
		AnnotationReferences: annotationReferences,
		Flags:                flags,
		InternalTypeName:     internalTypeName,
		Name:                 name,
		BodyDeclaration:      bodyDeclaration,
	}

	if bodyDeclaration == nil {
		d.FirstLineNumber = 0
	} else {
		d.FirstLineNumber = bodyDeclaration.FirstLineNumber
	}
	d.SetValue(&d)

	return d
}

type ClassFileAnnotationDeclaration struct {
	util.DefaultBase[model.IMemberDeclaration]

	AnnotationReferences  *model.AnnotationReference
	Flags                 int
	InternalTypeName      string
	Name                  string
	BodyDeclaration       model.IBodyDeclaration
	AnnotationDeclarators *model.FieldDeclarator
	FirstLineNumber       int
}

func (d *ClassFileAnnotationDeclaration) GetAnnotationReferences() *model.AnnotationReference {
	return d.AnnotationReferences
}

func (d *ClassFileAnnotationDeclaration) GetName() string {
	return d.Name
}

func (d *ClassFileAnnotationDeclaration) GetFlag() int {
	return d.Flags
}

func (d *ClassFileAnnotationDeclaration) GetFirstLineNumber() int {
	return d.FirstLineNumber
}

func (d *ClassFileAnnotationDeclaration) GetInternalTypeName() string {
	return d.InternalTypeName
}

func (d *ClassFileAnnotationDeclaration) GetBodyDeclaration() model.IBodyDeclaration {
	return d.BodyDeclaration
}

func (d *ClassFileAnnotationDeclaration) GetAnnotationDeclarators() *model.FieldDeclarator {
	return d.AnnotationDeclarators
}

func (d *ClassFileAnnotationDeclaration) IsClassDeclaration() bool {
	return false
}

func (d *ClassFileAnnotationDeclaration) AcceptDeclaration(visitor model.IDeclarationVisitor) {
	visitor.VisitAnnotationDeclaration(d)
}

func (d *ClassFileAnnotationDeclaration) String() string {
	return fmt.Sprintf("ClassFileAnnotationDeclaration{%s, firstLineNumber=%d}", d.InternalTypeName, d.FirstLineNumber)
}
