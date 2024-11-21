package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewInterfaceDeclaration(flags int, internalTypeName string, name string, interfaces intmod.IType) intmod.IInterfaceDeclaration {
	return &InterfaceDeclaration{
		TypeDeclaration: TypeDeclaration{
			annotationReferences: nil,
			flags:                flags,
			internalTypeName:     internalTypeName,
			name:                 name,
			bodyDeclaration:      nil,
		},
		interfaces: interfaces,
	}
}

func NewInterfaceDeclarationWithAll(annotationReferences intmod.IAnnotationReference, flags int,
	internalTypeName string, name string, bodyDeclaration intmod.IBodyDeclaration,
	typeParameters intmod.ITypeParameter, interfaces intmod.IType) intmod.IInterfaceDeclaration {
	return &InterfaceDeclaration{
		TypeDeclaration: TypeDeclaration{
			annotationReferences: annotationReferences,
			flags:                flags,
			internalTypeName:     internalTypeName,
			name:                 name,
			bodyDeclaration:      bodyDeclaration,
		},
		typeParameters: typeParameters,
		interfaces:     interfaces,
	}
}

type InterfaceDeclaration struct {
	TypeDeclaration

	typeParameters intmod.ITypeParameter
	interfaces     intmod.IType
}

func (d *InterfaceDeclaration) TypeParameters() intmod.ITypeParameter {
	return d.typeParameters
}

func (d *InterfaceDeclaration) Interfaces() intmod.IType {
	return d.interfaces
}

func (d *InterfaceDeclaration) Accept(visitor intmod.IDeclarationVisitor) {
	visitor.VisitInterfaceDeclaration(d)
}

func (d *InterfaceDeclaration) String() string {
	return fmt.Sprintf("InterfaceDeclaration{%v}", *d)
}
