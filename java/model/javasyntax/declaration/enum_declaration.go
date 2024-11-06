package declaration

import (
	"bitbucket.org/coontec/javaClass/java/model/javasyntax/expression"
	"bitbucket.org/coontec/javaClass/java/model/javasyntax/reference"
	_type "bitbucket.org/coontec/javaClass/java/model/javasyntax/type"
	"fmt"
)

func NewEnumDeclaration(flags int, internalTypeName string, name string, interfaces _type.IType) *InterfaceDeclaration {
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

func NewEnumDeclarationWithAll(annotationReferences reference.IAnnotationReference, flags int, internalTypeName string, name string, bodyDeclaration *BodyDeclaration, typeParameters _type.ITypeParameter, interfaces _type.IType) *InterfaceDeclaration {
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

type EnumDeclaration struct {
	TypeDeclaration

	interfaces _type.IType
	constants  []Constant
}

func (d *EnumDeclaration) GetInterfaces() _type.IType {
	return d.interfaces
}

func (d *EnumDeclaration) GetConstants() []Constant {
	return d.constants
}

func (d *EnumDeclaration) Accept(visitor DeclarationVisitor) {
	visitor.VisitEnumDeclaration(d)
}

func (d *EnumDeclaration) String() string {
	return fmt.Sprintf("EnumDeclaration{%s}", d.internalTypeName)
}

func NewConstant(name string) *Constant {
	return &Constant{
		name: name,
	}
}

func NewConstant2(lineNumber int, name string) *Constant {
	return &Constant{
		lineNumber: lineNumber,
		name:       name,
	}
}

func NewConstant3(name string, arguments expression.Expression) *Constant {
	return &Constant{
		name:      name,
		arguments: arguments,
	}
}

func NewConstant4(lineNumber int, name string, arguments expression.Expression) *Constant {
	return &Constant{
		lineNumber: lineNumber,
		name:       name,
		arguments:  arguments,
	}
}

func NewConstant5(lineNumber int, name string, arguments expression.Expression, bodyDeclaration *BodyDeclaration) *Constant {
	return &Constant{
		lineNumber:      lineNumber,
		name:            name,
		arguments:       arguments,
		bodyDeclaration: bodyDeclaration,
	}
}

func NewConstant6(lineNumber int, annotationReferences reference.IAnnotationReference, name string, arguments expression.Expression, bodyDeclaration *BodyDeclaration) *Constant {
	return &Constant{
		lineNumber:           lineNumber,
		annotationReferences: annotationReferences,
		name:                 name,
		arguments:            arguments,
		bodyDeclaration:      bodyDeclaration,
	}
}

type Constant struct {
	lineNumber           int
	annotationReferences reference.IAnnotationReference
	name                 string
	arguments            expression.Expression
	bodyDeclaration      *BodyDeclaration
}

func (c *Constant) GetLineNumber() int {
	return c.lineNumber
}

func (c *Constant) GetAnnotationReferences() reference.IAnnotationReference {
	return c.annotationReferences
}

func (c *Constant) GetName() string {
	return c.name
}

func (c *Constant) GetArguments() expression.Expression {
	return c.arguments
}

func (c *Constant) SetArguments(arguments expression.Expression) {
	c.arguments = arguments
}

func (c *Constant) GetBodyDeclaration() *BodyDeclaration {
	return c.bodyDeclaration
}

func (c *Constant) Accept(visitor DeclarationVisitor) {
	visitor.VisitEnumDeclarationConstant(c)
}
