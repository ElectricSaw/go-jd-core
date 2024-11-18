package declaration

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/reference"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewInterfaceDeclaration(flags int, internalTypeName string, name string, interfaces _type.IType) intsyn.IInterfaceDeclaration {
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

func NewInterfaceDeclarationWithAll(annotationReferences reference.IAnnotationReference, flags int, internalTypeName string, name string, bodyDeclaration intsyn.IBodyDeclaration, typeParameters _type.ITypeParameter, interfaces _type.IType) intsyn.IInterfaceDeclaration {
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

func (d *InterfaceDeclaration) TypeParameters() _type.ITypeParameter {
	return d.typeParameters
}

func (d *InterfaceDeclaration) Interfaces() _type.IType {
	return d.interfaces
}

func (d *InterfaceDeclaration) Accept(visitor intsyn.IDeclarationVisitor) {
	visitor.VisitInterfaceDeclaration(d)
}

func (d *InterfaceDeclaration) String() string {
	return fmt.Sprintf("InterfaceDeclaration{%v}", *d)
}
