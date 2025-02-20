package declaration

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewClassFileInterfaceDeclaration(
	annotationReferences *model.AnnotationReference,
	flags int,
	internalTypeName string,
	name string,
	typeParameters model.ITypeParameter,
	interfaces model.IType,
	bodyDeclaration *ClassFileBodyDeclaration,
) ClassFileInterfaceDeclaration {
	d := ClassFileInterfaceDeclaration{
		DefaultBase:          *util.NewDefaultBase[model.IMemberDeclaration]().(*util.DefaultBase[model.IMemberDeclaration]),
		AnnotationReferences: annotationReferences,
		Flags:                flags,
		InternalTypeName:     internalTypeName,
		Name:                 name,
		BodyDeclaration:      bodyDeclaration,
		TypeParameters:       typeParameters,
		Interfaces:           interfaces,
		FirstLineNumber:      bodyDeclaration.FirstLineNumber,
	}
	d.SetValue(&d)
	return d
}

type ClassFileInterfaceDeclaration struct {
	util.DefaultBase[model.IMemberDeclaration]

	AnnotationReferences *model.AnnotationReference
	Flags                int
	InternalTypeName     string
	Name                 string
	BodyDeclaration      model.IBodyDeclaration
	TypeParameters       model.ITypeParameter
	Interfaces           model.IType
	FirstLineNumber      int
}

func (d *ClassFileInterfaceDeclaration) GetAnnotationReferences() *model.AnnotationReference {
	return d.AnnotationReferences
}

func (d *ClassFileInterfaceDeclaration) GetFlag() int {
	return d.Flags
}

func (d *ClassFileInterfaceDeclaration) GetInternalTypeName() string {
	return d.InternalTypeName
}

func (d *ClassFileInterfaceDeclaration) GetName() string {
	return d.Name
}

func (d *ClassFileInterfaceDeclaration) GetBodyDeclaration() model.IBodyDeclaration {
	return d.BodyDeclaration
}

func (d *ClassFileInterfaceDeclaration) GetTypeParameters() model.ITypeParameter {
	return d.TypeParameters
}

func (d *ClassFileInterfaceDeclaration) GetInterfaces() model.IType {
	return d.Interfaces
}

func (d *ClassFileInterfaceDeclaration) GetFirstLineNumber() int {
	return d.FirstLineNumber
}

func (d *ClassFileInterfaceDeclaration) IsClassDeclaration() bool {
	return false
}

func (d *ClassFileInterfaceDeclaration) AcceptDeclaration(visitor model.IDeclarationVisitor) {
	visitor.VisitInterfaceDeclaration(d)
}

func (d *ClassFileInterfaceDeclaration) String() string {
	return fmt.Sprintf("ClassFileInterfaceDeclaration{%s, firstLineNumber=%d}", d.InternalTypeName, d.FirstLineNumber)
}
