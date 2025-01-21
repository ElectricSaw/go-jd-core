package declaration

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewInterfaceDeclaration(flags int, internalTypeName, name string, interfaces Type) InterfaceDeclaration {
	return NewInterfaceDeclarationWithAll(nil, flags, internalTypeName, name, nil, nil, interfaces)
}

func NewInterfaceDeclarationWithAll(annotationReferences AnnotationReference, flags int,
	internalTypeName, name string, bodyDeclaration BodyDeclaration,
	typeParameters TypeParameter, interfaces Type) InterfaceDeclaration {
	d := &InterfaceDeclaration{
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
	d.SetValue(d)
	return d
}

type InterfaceDeclaration struct {
	TypeDeclaration

	typeParameters TypeParameter
	interfaces     Type
}

func (d *InterfaceDeclaration) TypeParameters() TypeParameter {
	return d.typeParameters
}

func (d *InterfaceDeclaration) Interfaces() Type {
	return d.interfaces
}

func (d *InterfaceDeclaration) AcceptDeclaration(visitor DeclarationVisitor) {
	visitor.VisitInterfaceDeclaration(d)
}

func (d *InterfaceDeclaration) String() string {
	return fmt.Sprintf("InterfaceDeclaration{%v}", *d)
}
