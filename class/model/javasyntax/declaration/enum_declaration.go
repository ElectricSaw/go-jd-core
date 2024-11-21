package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewEnumDeclaration(flags int, internalTypeName, name string, constants []intmod.IConstant,
	bodyDeclaration intmod.IDeclaration) intmod.IEnumDeclaration {
	return &EnumDeclaration{
		TypeDeclaration: *NewTypeDeclaration(nil, flags, internalTypeName, name, bodyDeclaration).(*TypeDeclaration),
		constants:       constants,
	}
}

func NewEnumDeclarationWithAll(annotationReferences intmod.IAnnotationReference, flags int, internalTypeName,
	name string, interfaces intmod.IType, constants []intmod.IConstant, bodyDeclaration intmod.IDeclaration) intmod.IEnumDeclaration {
	return &EnumDeclaration{
		TypeDeclaration: *NewTypeDeclaration(annotationReferences, flags, internalTypeName, name, bodyDeclaration).(*TypeDeclaration),
		interfaces:      interfaces,
		constants:       constants,
	}
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

func (d *EnumDeclaration) Accept(visitor intmod.IDeclarationVisitor) {
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

func NewConstant3(name string, arguments intmod.IExpression) *Constant {
	return &Constant{
		name:      name,
		arguments: arguments,
	}
}

func NewConstant4(lineNumber int, name string, arguments intmod.IExpression) *Constant {
	return &Constant{
		lineNumber: lineNumber,
		name:       name,
		arguments:  arguments,
	}
}

func NewConstant5(lineNumber int, name string, arguments intmod.IExpression,
	bodyDeclaration intmod.IBodyDeclaration) intmod.IConstant {
	return &Constant{
		lineNumber:      lineNumber,
		name:            name,
		arguments:       arguments,
		bodyDeclaration: bodyDeclaration,
	}
}

func NewConstant6(lineNumber int, annotationReferences intmod.IAnnotationReference, name string,
	arguments intmod.IExpression, bodyDeclaration intmod.IBodyDeclaration) intmod.IConstant {
	return &Constant{
		lineNumber:           lineNumber,
		annotationReferences: annotationReferences,
		name:                 name,
		arguments:            arguments,
		bodyDeclaration:      bodyDeclaration,
	}
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

func (c *Constant) Accept(visitor intmod.IDeclarationVisitor) {
	visitor.VisitEnumDeclarationConstant(c)
}
