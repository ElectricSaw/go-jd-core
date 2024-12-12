package declaration

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
)

func NewEnumDeclaration(flags int, internalTypeName, name string, constants []intmod.IConstant,
	bodyDeclaration intmod.IDeclaration) intmod.IEnumDeclaration {
	return NewEnumDeclarationWithAll(nil, flags,
		internalTypeName, name, nil, constants, bodyDeclaration)
}

func NewEnumDeclarationWithAll(annotationReferences intmod.IAnnotationReference,
	flags int, internalTypeName, name string, interfaces intmod.IType,
	constants []intmod.IConstant, bodyDeclaration intmod.IDeclaration) intmod.IEnumDeclaration {
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

	interfaces intmod.IType
	constants  []intmod.IConstant
}

func (d *EnumDeclaration) Interfaces() intmod.IType {
	return d.interfaces
}

func (d *EnumDeclaration) Constants() []intmod.IConstant {
	return d.constants
}

func (d *EnumDeclaration) SetConstants(constants []intmod.IConstant) {
	d.constants = constants
}

func (d *EnumDeclaration) AcceptDeclaration(visitor intmod.IDeclarationVisitor) {
	visitor.VisitEnumDeclaration(d)
}

func (d *EnumDeclaration) String() string {
	return fmt.Sprintf("EnumDeclaration{%s}", d.internalTypeName)
}

func NewConstant(name string) intmod.IConstant {
	return NewConstant6(-1, nil, name, nil, nil)
}

func NewConstant2(lineNumber int, name string) intmod.IConstant {
	return NewConstant6(lineNumber, nil, name, nil, nil)
}

func NewConstant3(name string, arguments intmod.IExpression) intmod.IConstant {
	return NewConstant6(-1, nil, name, arguments, nil)
}

func NewConstant4(lineNumber int, name string, arguments intmod.IExpression) intmod.IConstant {
	return NewConstant6(lineNumber, nil, name, arguments, nil)
}

func NewConstant5(lineNumber int, name string, arguments intmod.IExpression,
	bodyDeclaration intmod.IBodyDeclaration) intmod.IConstant {
	return NewConstant6(lineNumber, nil, name, arguments, bodyDeclaration)
}

func NewConstant6(lineNumber int, annotationReferences intmod.IAnnotationReference, name string,
	arguments intmod.IExpression, bodyDeclaration intmod.IBodyDeclaration) intmod.IConstant {
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
	annotationReferences intmod.IAnnotationReference
	name                 string
	arguments            intmod.IExpression
	bodyDeclaration      intmod.IBodyDeclaration
}

func (c *Constant) LineNumber() int {
	return c.lineNumber
}

func (c *Constant) AnnotationReferences() intmod.IAnnotationReference {
	return c.annotationReferences
}

func (c *Constant) Name() string {
	return c.name
}

func (c *Constant) Arguments() intmod.IExpression {
	return c.arguments
}

func (c *Constant) SetArguments(arguments intmod.IExpression) {
	c.arguments = arguments
}

func (c *Constant) BodyDeclaration() intmod.IBodyDeclaration {
	return c.bodyDeclaration
}

func (c *Constant) AcceptDeclaration(visitor intmod.IDeclarationVisitor) {
	visitor.VisitEnumDeclarationConstant(c)
}
