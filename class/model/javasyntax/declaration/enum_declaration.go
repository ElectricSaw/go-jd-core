package declaration

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewEnumDeclaration(flags int, internalTypeName, name string, constants []intsyn.IConstant,
	bodyDeclaration intsyn.IDeclaration) intsyn.IEnumDeclaration {
	return &EnumDeclaration{
		TypeDeclaration: *NewTypeDeclaration(nil, flags, internalTypeName, name, bodyDeclaration).(*TypeDeclaration),
		constants:       constants,
	}
}

func NewEnumDeclarationWithAll(annotationReferences intsyn.IAnnotationReference, flags int, internalTypeName,
	name string, interfaces intsyn.IType, constants []intsyn.IConstant, bodyDeclaration intsyn.IDeclaration) intsyn.IEnumDeclaration {
	return &EnumDeclaration{
		TypeDeclaration: *NewTypeDeclaration(annotationReferences, flags, internalTypeName, name, bodyDeclaration).(*TypeDeclaration),
		interfaces:      interfaces,
		constants:       constants,
	}
}

type EnumDeclaration struct {
	TypeDeclaration

	interfaces intsyn.IType
	constants  []intsyn.IConstant
}

func (d *EnumDeclaration) Interfaces() intsyn.IType {
	return d.interfaces
}

func (d *EnumDeclaration) Constants() []intsyn.IConstant {
	return d.constants
}

func (d *EnumDeclaration) SetConstants(constants []intsyn.IConstant) {
	d.constants = constants
}

func (d *EnumDeclaration) Accept(visitor intsyn.IDeclarationVisitor) {
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

func NewConstant3(name string, arguments intsyn.IExpression) *Constant {
	return &Constant{
		name:      name,
		arguments: arguments,
	}
}

func NewConstant4(lineNumber int, name string, arguments intsyn.IExpression) *Constant {
	return &Constant{
		lineNumber: lineNumber,
		name:       name,
		arguments:  arguments,
	}
}

func NewConstant5(lineNumber int, name string, arguments intsyn.IExpression,
	bodyDeclaration intsyn.IBodyDeclaration) intsyn.IConstant {
	return &Constant{
		lineNumber:      lineNumber,
		name:            name,
		arguments:       arguments,
		bodyDeclaration: bodyDeclaration,
	}
}

func NewConstant6(lineNumber int, annotationReferences intsyn.IAnnotationReference, name string,
	arguments intsyn.IExpression, bodyDeclaration intsyn.IBodyDeclaration) intsyn.IConstant {
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
	annotationReferences intsyn.IAnnotationReference
	name                 string
	arguments            intsyn.IExpression
	bodyDeclaration      intsyn.IBodyDeclaration
}

func (c *Constant) LineNumber() int {
	return c.lineNumber
}

func (c *Constant) AnnotationReferences() intsyn.IAnnotationReference {
	return c.annotationReferences
}

func (c *Constant) Name() string {
	return c.name
}

func (c *Constant) Arguments() intsyn.IExpression {
	return c.arguments
}

func (c *Constant) SetArguments(arguments intsyn.IExpression) {
	c.arguments = arguments
}

func (c *Constant) BodyDeclaration() intsyn.IBodyDeclaration {
	return c.bodyDeclaration
}

func (c *Constant) Accept(visitor intsyn.IDeclarationVisitor) {
	visitor.VisitEnumDeclarationConstant(c)
}
