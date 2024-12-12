package declaration

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
)

func NewInterfaceDeclaration(flags int, internalTypeName, name string, interfaces intmod.IType) intmod.IInterfaceDeclaration {
	return NewInterfaceDeclarationWithAll(nil, flags, internalTypeName, name, nil, nil, interfaces)
}

func NewInterfaceDeclarationWithAll(annotationReferences intmod.IAnnotationReference, flags int,
	internalTypeName, name string, bodyDeclaration intmod.IBodyDeclaration,
	typeParameters intmod.ITypeParameter, interfaces intmod.IType) intmod.IInterfaceDeclaration {
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

	typeParameters intmod.ITypeParameter
	interfaces     intmod.IType
}

func (d *InterfaceDeclaration) TypeParameters() intmod.ITypeParameter {
	return d.typeParameters
}

func (d *InterfaceDeclaration) Interfaces() intmod.IType {
	return d.interfaces
}

func (d *InterfaceDeclaration) AcceptDeclaration(visitor intmod.IDeclarationVisitor) {
	visitor.VisitInterfaceDeclaration(d)
}

func (d *InterfaceDeclaration) String() string {
	return fmt.Sprintf("InterfaceDeclaration{%v}", *d)
}
