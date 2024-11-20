package declaration

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewInterfaceDeclaration(flags int, internalTypeName string, name string, interfaces intsyn.IType) intsyn.IInterfaceDeclaration {
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

func NewInterfaceDeclarationWithAll(annotationReferences intsyn.IAnnotationReference, flags int,
	internalTypeName string, name string, bodyDeclaration intsyn.IBodyDeclaration,
	typeParameters intsyn.ITypeParameter, interfaces intsyn.IType) intsyn.IInterfaceDeclaration {
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

	typeParameters intsyn.ITypeParameter
	interfaces     intsyn.IType
}

func (d *InterfaceDeclaration) TypeParameters() intsyn.ITypeParameter {
	return d.typeParameters
}

func (d *InterfaceDeclaration) Interfaces() intsyn.IType {
	return d.interfaces
}

func (d *InterfaceDeclaration) Accept(visitor intsyn.IDeclarationVisitor) {
	visitor.VisitInterfaceDeclaration(d)
}

func (d *InterfaceDeclaration) String() string {
	return fmt.Sprintf("InterfaceDeclaration{%v}", *d)
}
