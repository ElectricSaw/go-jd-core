package declaration

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	"bitbucket.org/coontec/javaClass/class/model/classfile"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/reference"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/statement"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewConstructorDeclaration(flags int, formalParameter intsyn.IFormalParameter, descriptor string, statements statement.Statement) intsyn.IConstructorDeclaration {
	return &ConstructorDeclaration{
		flags:           flags,
		formalParameter: formalParameter,
		descriptor:      descriptor,
		statements:      statements,
	}
}

func NewConstructorDeclarationWithAll(annotationReferences reference.IReference, flags int, typeParameters _type.ITypeParameter, formalParameter intsyn.IFormalParameter, exceptionTypes _type.IType, descriptor string, statements statement.Statement) intsyn.IConstructorDeclaration {
	return &ConstructorDeclaration{
		annotationReferences: annotationReferences,
		flags:                flags,
		typeParameters:       typeParameters,
		formalParameter:      formalParameter,
		exceptionTypes:       exceptionTypes,
		descriptor:           descriptor,
		statements:           statements,
	}
}

type ConstructorDeclaration struct {
	AbstractMemberDeclaration

	annotationReferences reference.IReference
	flags                int
	typeParameters       _type.ITypeParameter
	formalParameter      intsyn.IFormalParameter
	exceptionTypes       _type.IType
	descriptor           string
	statements           statement.Statement
}

func (d *ConstructorDeclaration) Flags() int {
	return d.flags
}

func (d *ConstructorDeclaration) SetFlags(flags int) {
	d.flags = flags
}

func (d *ConstructorDeclaration) IsStatic() bool {
	return d.flags&classfile.AccStatic != 0
}

func (d *ConstructorDeclaration) AnnotationReferences() reference.IReference {
	return d.annotationReferences
}

func (d *ConstructorDeclaration) TypeParameters() _type.ITypeParameter {
	return d.typeParameters
}

func (d *ConstructorDeclaration) FormalParameters() intsyn.IFormalParameter {
	return d.formalParameter
}

func (d *ConstructorDeclaration) SetFormalParameters(formalParameter intsyn.IFormalParameter) {
	d.formalParameter = formalParameter
}

func (d *ConstructorDeclaration) ExceptionTypes() _type.IType {
	return d.exceptionTypes
}

func (d *ConstructorDeclaration) Descriptor() string {
	return d.descriptor
}

func (d *ConstructorDeclaration) Statements() statement.Statement {
	return d.statements
}

func (d *ConstructorDeclaration) SetStatements(state statement.Statement) {
	d.statements = state
}

func (d *ConstructorDeclaration) Accept(visitor intsyn.IDeclarationVisitor) {
	visitor.VisitConstructorDeclaration(d)
}

func (d *ConstructorDeclaration) String() string {
	return fmt.Sprintf("ConstructorDeclaration{%s}", d.descriptor)
}
