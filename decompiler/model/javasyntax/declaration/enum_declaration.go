package declaration

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewEnumDeclaration(flags int, internalTypeName, name string, constants []Constant,
	bodyDeclaration Declaration) EnumDeclaration {
	return NewEnumDeclarationWithAll(nil, flags,
		internalTypeName, name, nil, constants, bodyDeclaration)
}

func NewEnumDeclarationWithAll(annotationReferences AnnotationReference,
	flags int, internalTypeName, name string, interfaces Type,
	constants []Constant, bodyDeclaration Declaration) EnumDeclaration {
	d := &EnumDeclaration{
		TypeDeclaration: *NewTypeDeclaration(annotationReferences, flags, internalTypeName, name, bodyDeclaration).(*TypeDeclaration),
		interfaces:      interfaces,
		constants:       constants,
	}
	d.SetValue(d)
	return d
}

type EnumDeclaration struct {
	TypeDeclaration

	interfaces Type
	constants  []Constant
}

func (d *EnumDeclaration) Interfaces() Type {
	return d.interfaces
}

func (d *EnumDeclaration) Constants() []Constant {
	return d.constants
}

func (d *EnumDeclaration) SetConstants(constants []Constant) {
	d.constants = constants
}

func (d *EnumDeclaration) AcceptDeclaration(visitor DeclarationVisitor) {
	visitor.VisitEnumDeclaration(d)
}

func (d *EnumDeclaration) String() string {
	return fmt.Sprintf("EnumDeclaration{%s}", d.internalTypeName)
}

func NewConstant(name string) Constant {
	return NewConstant6(-1, nil, name, nil, nil)
}

func NewConstant2(lineNumber int, name string) Constant {
	return NewConstant6(lineNumber, nil, name, nil, nil)
}

func NewConstant3(name string, arguments Expression) Constant {
	return NewConstant6(-1, nil, name, arguments, nil)
}

func NewConstant4(lineNumber int, name string, arguments Expression) Constant {
	return NewConstant6(lineNumber, nil, name, arguments, nil)
}

func NewConstant5(lineNumber int, name string, arguments Expression,
	bodyDeclaration BodyDeclaration) Constant {
	return NewConstant6(lineNumber, nil, name, arguments, bodyDeclaration)
}

func NewConstant6(lineNumber int, annotationReferences AnnotationReference, name string,
	arguments Expression, bodyDeclaration BodyDeclaration) Constant {
	c := &Constant{
		lineNumber:           lineNumber,
		annotationReferences: annotationReferences,
		name:                 name,
		arguments:            arguments,
		bodyDeclaration:      bodyDeclaration,
	}
	c.SetValue(c)
	return c
}

type Constant struct {
	TypeDeclaration

	lineNumber           int
	annotationReferences AnnotationReference
	name                 string
	arguments            Expression
	bodyDeclaration      BodyDeclaration
}

func (c *Constant) LineNumber() int {
	return c.lineNumber
}

func (c *Constant) AnnotationReferences() AnnotationReference {
	return c.annotationReferences
}

func (c *Constant) Name() string {
	return c.name
}

func (c *Constant) Arguments() Expression {
	return c.arguments
}

func (c *Constant) SetArguments(arguments Expression) {
	c.arguments = arguments
}

func (c *Constant) BodyDeclaration() BodyDeclaration {
	return c.bodyDeclaration
}

func (c *Constant) AcceptDeclaration(visitor DeclarationVisitor) {
	visitor.VisitEnumDeclarationConstant(c)
}
