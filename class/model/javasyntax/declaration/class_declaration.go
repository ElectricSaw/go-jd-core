package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewClassDeclaration(flags int, internalTypeName string, name string,
	bodyDeclaration intmod.IBodyDeclaration) intmod.IClassDeclaration {
	return &ClassDeclaration{
		InterfaceDeclaration: *NewInterfaceDeclarationWithAll(nil, flags,
			internalTypeName, name, bodyDeclaration, nil, nil).(*InterfaceDeclaration),
	}
}

func NewClassDeclarationWithAll(annotationReferences intmod.IAnnotationReference, flags int,
	internalTypeName string, name string, bodyDeclaration intmod.IBodyDeclaration,
	typeParameters intmod.ITypeParameter, interfaces intmod.IType, superType intmod.IObjectType) intmod.IClassDeclaration {
	return &ClassDeclaration{
		InterfaceDeclaration: *NewInterfaceDeclarationWithAll(annotationReferences, flags,
			internalTypeName, name, bodyDeclaration, typeParameters, interfaces).(*InterfaceDeclaration),
		superType: superType,
	}
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

func (d *ClassDeclaration) Accept(visitor intmod.IDeclarationVisitor) {
	visitor.VisitClassDeclaration(d)
}

func (d *ClassDeclaration) String() string {
	return fmt.Sprintf("ClassDeclaration{%s}", d.internalTypeName)
}
