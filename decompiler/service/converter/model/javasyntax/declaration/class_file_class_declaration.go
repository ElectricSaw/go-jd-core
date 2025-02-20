package declaration

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewClassFileClassDeclaration(annotationReferences *model.AnnotationReference,
	flags int, internalName, name string, typeParameters model.ITypeParameter,
	superType *model.ObjectType, interfaces model.IType, bodyDeclaration *ClassFileBodyDeclaration,
) ClassFileClassDeclaration {
	d := ClassFileClassDeclaration{
		DefaultBase:          *util.NewDefaultBase[model.IMemberDeclaration]().(*util.DefaultBase[model.IMemberDeclaration]),
		AnnotationReferences: annotationReferences,
		Flags:                flags,
		InternalTypeName:     internalName,
		Name:                 name,
		BodyDeclaration:      bodyDeclaration,
		TypeParameters:       typeParameters,
		Interfaces:           interfaces,
		SuperType:            superType,
	}

	if bodyDeclaration != nil {
		d.FirstLineNumber = bodyDeclaration.FirstLineNumber
	}
	d.SetValue(&d)

	return d
}

type ClassFileClassDeclaration struct {
	util.DefaultBase[model.IMemberDeclaration]

	AnnotationReferences *model.AnnotationReference
	Flags                int
	InternalTypeName     string
	Name                 string
	BodyDeclaration      model.IBodyDeclaration
	TypeParameters       model.ITypeParameter
	Interfaces           model.IType
	SuperType            *model.ObjectType
	FirstLineNumber      int
}

func (d *ClassFileClassDeclaration) GetFirstLineNumber() int {
	return d.FirstLineNumber
}

func (d *ClassFileClassDeclaration) GetAnnotationReferences() *model.AnnotationReference {
	return d.AnnotationReferences
}

func (d *ClassFileClassDeclaration) GetFlag() int {
	return d.Flags
}

func (d *ClassFileClassDeclaration) GetInternalTypeName() string {
	return d.InternalTypeName
}

func (d *ClassFileClassDeclaration) GetName() string {
	return d.Name
}

func (d *ClassFileClassDeclaration) GetBodyDeclaration() model.IBodyDeclaration {
	return d.BodyDeclaration
}

func (d *ClassFileClassDeclaration) GetTypeParameters() model.ITypeParameter {
	return d.TypeParameters
}

func (d *ClassFileClassDeclaration) GetInterfaces() model.IType {
	return d.Interfaces
}

func (d *ClassFileClassDeclaration) GetSuperType() *model.ObjectType {
	return d.SuperType
}

func (d *ClassFileClassDeclaration) IsClassDeclaration() bool {
	return true
}

func (d *ClassFileClassDeclaration) AcceptDeclaration(visitor model.IDeclarationVisitor) {
	visitor.VisitClassDeclaration(d)
}

func (d *ClassFileClassDeclaration) String() string {
	return fmt.Sprintf("ClassFileClassDeclaration{%s, firstLineNumber=%d}", d.InternalTypeName, d.FirstLineNumber)
}
