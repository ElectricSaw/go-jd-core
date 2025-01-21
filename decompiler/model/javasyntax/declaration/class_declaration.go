package declaration

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewClassDeclaration(flags int, internalTypeName, name string,
	bodyDeclaration BodyDeclaration) ClassDeclaration {
	return NewClassDeclarationWithAll(nil, flags, internalTypeName, name,
		bodyDeclaration, nil, nil, nil)
}

func NewClassDeclarationWithAll(annotationReferences AnnotationReference, flags int,
	internalTypeName, name string, bodyDeclaration BodyDeclaration,
	typeParameters TypeParameter, interfaces Type, superType ObjectType) ClassDeclaration {
	d := &ClassDeclaration{
		InterfaceDeclaration: *NewInterfaceDeclarationWithAll(annotationReferences, flags,
			internalTypeName, name, bodyDeclaration, typeParameters, interfaces).(*InterfaceDeclaration),
		superType: superType,
	}
	d.SetValue(d)
	return d
}

type ClassDeclaration struct {
	InterfaceDeclaration

	superType ObjectType
}

func (d *ClassDeclaration) SuperType() ObjectType {
	return d.superType
}

func (d *ClassDeclaration) IsClassDeclaration() bool {
	return true
}

func (d *ClassDeclaration) AcceptDeclaration(visitor DeclarationVisitor) {
	visitor.VisitClassDeclaration(d)
}

func (d *ClassDeclaration) String() string {
	return fmt.Sprintf("ClassDeclaration{%s}", d.internalTypeName)
}
