package declaration

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/expression"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/reference"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewEnumDeclaration(flags int, internalTypeName, name string, constants []Constant, bodyDeclaration Declaration) *EnumDeclaration {
	return &EnumDeclaration{
		TypeDeclaration: *NewTypeDeclaration(nil, flags, internalTypeName, name, bodyDeclaration),
		constants:       constants,
	}
}

func NewEnumDeclarationWithAll(annotationReferences reference.IAnnotationReference, flags int, internalTypeName, name string, interfaces _type.IType, constants []Constant, bodyDeclaration Declaration) *EnumDeclaration {
	return &EnumDeclaration{
		TypeDeclaration: *NewTypeDeclaration(annotationReferences, flags, internalTypeName, name, bodyDeclaration),
		interfaces:      interfaces,
		constants:       constants,
	}
}

type EnumDeclaration struct {
	TypeDeclaration

	interfaces _type.IType
	constants  []Constant
}

func (d *EnumDeclaration) Interfaces() _type.IType {
	return d.interfaces
}

func (d *EnumDeclaration) Constants() []Constant {
	return d.constants
}

func (d *EnumDeclaration) SetConstants(constants []Constant) {
	d.constants = constants
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

func (c *Constant) LineNumber() int {
	return c.lineNumber
}

func (c *Constant) AnnotationReferences() reference.IAnnotationReference {
	return c.annotationReferences
}

func (c *Constant) Name() string {
	return c.name
}

func (c *Constant) Arguments() expression.Expression {
	return c.arguments
}

func (c *Constant) SetArguments(arguments expression.Expression) {
	c.arguments = arguments
}

func (c *Constant) BodyDeclaration() *BodyDeclaration {
	return c.bodyDeclaration
}

func (c *Constant) Accept(visitor DeclarationVisitor) {
	visitor.VisitEnumDeclarationConstant(c)
}
