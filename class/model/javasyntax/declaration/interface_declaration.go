package declaration

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/reference"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewInterfaceDeclaration(flags int, internalTypeName string, name string, interfaces _type.IType) *InterfaceDeclaration {
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

func NewInterfaceDeclarationWithAll(annotationReferences reference.IAnnotationReference, flags int, internalTypeName string, name string, bodyDeclaration *BodyDeclaration, typeParameters _type.ITypeParameter, interfaces _type.IType) *InterfaceDeclaration {
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

	typeParameters _type.ITypeParameter
	interfaces     _type.IType
}

func (d *InterfaceDeclaration) GetTypeParameters() _type.ITypeParameter {
	return d.typeParameters
}

func (d *InterfaceDeclaration) GetInterfaces() _type.IType {
	return d.interfaces
}

func (d *InterfaceDeclaration) Accept(visitor DeclarationVisitor) {
	visitor.VisitInterfaceDeclaration(d)
}

func (d *InterfaceDeclaration) String() string {
	return fmt.Sprintf("InterfaceDeclaration{%v}", *d)
}
