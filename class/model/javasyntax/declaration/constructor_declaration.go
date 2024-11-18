package declaration

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	"bitbucket.org/coontec/javaClass/class/model/classfile"
	"fmt"
)

func NewConstructorDeclaration(flags int, formalParameter intsyn.IFormalParameter, descriptor string,
	statements intsyn.IStatement) intsyn.IConstructorDeclaration {
	return &ConstructorDeclaration{
		flags:           flags,
		formalParameter: formalParameter,
		descriptor:      descriptor,
		statements:      statements,
	}
}

func NewConstructorDeclarationWithAll(annotationReferences intsyn.IReference, flags int,
	typeParameters intsyn.ITypeParameter, formalParameter intsyn.IFormalParameter,
	exceptionTypes intsyn.IType, descriptor string, statements intsyn.IStatement) intsyn.IConstructorDeclaration {
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

	annotationReferences intsyn.IReference
	flags                int
	typeParameters       intsyn.ITypeParameter
	formalParameter      intsyn.IFormalParameter
	exceptionTypes       intsyn.IType
	descriptor           string
	statements           intsyn.IStatement
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

func (d *ConstructorDeclaration) AnnotationReferences() intsyn.IReference {
	return d.annotationReferences
}

func (d *ConstructorDeclaration) TypeParameters() intsyn.ITypeParameter {
	return d.typeParameters
}

func (d *ConstructorDeclaration) FormalParameters() intsyn.IFormalParameter {
	return d.formalParameter
}

func (d *ConstructorDeclaration) SetFormalParameters(formalParameter intsyn.IFormalParameter) {
	d.formalParameter = formalParameter
}

func (d *ConstructorDeclaration) ExceptionTypes() intsyn.IType {
	return d.exceptionTypes
}

func (d *ConstructorDeclaration) Descriptor() string {
	return d.descriptor
}

func (d *ConstructorDeclaration) Statements() intsyn.IStatement {
	return d.statements
}

func (d *ConstructorDeclaration) SetStatements(state intsyn.IStatement) {
	d.statements = state
}

func (d *ConstructorDeclaration) Accept(visitor intsyn.IDeclarationVisitor) {
	visitor.VisitConstructorDeclaration(d)
}

func (d *ConstructorDeclaration) String() string {
	return fmt.Sprintf("ConstructorDeclaration{%s}", d.descriptor)
}
