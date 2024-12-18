package declaration

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewClassDeclaration(flags int, internalTypeName, name string,
	bodyDeclaration intmod.IBodyDeclaration) intmod.IClassDeclaration {
	return NewClassDeclarationWithAll(nil, flags, internalTypeName, name,
		bodyDeclaration, nil, nil, nil)
}

func NewClassDeclarationWithAll(annotationReferences intmod.IAnnotationReference, flags int,
	internalTypeName, name string, bodyDeclaration intmod.IBodyDeclaration,
	typeParameters intmod.ITypeParameter, interfaces intmod.IType, superType intmod.IObjectType) intmod.IClassDeclaration {
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

	superType intmod.IObjectType
}

func (d *ClassDeclaration) SuperType() intmod.IObjectType {
	return d.superType
}

func (d *ClassDeclaration) IsClassDeclaration() bool {
	return true
}

func (d *ClassDeclaration) AcceptDeclaration(visitor intmod.IDeclarationVisitor) {
	visitor.VisitClassDeclaration(d)
}

func (d *ClassDeclaration) String() string {
	return fmt.Sprintf("ClassDeclaration{%s}", d.internalTypeName)
}
