package declaration

import (
	"bitbucket.org/coontec/javaClass/class/model/classfile"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/reference"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/statement"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewConstructorDeclaration(flags int, formalParameter IFormalParameter, descriptor string, statements statement.Statement) *ConstructorDeclaration {
	return &ConstructorDeclaration{
		flags:           flags,
		formalParameter: formalParameter,
		descriptor:      descriptor,
		statements:      statements,
	}
}

func NewConstructorDeclarationWithAll(annotationReferences reference.AnnotationReference, flags int, typeParameters _type.ITypeParameter, formalParameter IFormalParameter, exceptionTypes _type.IType, descriptor string, statements statement.Statement) *ConstructorDeclaration {
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

	annotationReferences reference.AnnotationReference
	flags                int
	typeParameters       _type.ITypeParameter
	formalParameter      IFormalParameter
	exceptionTypes       _type.IType
	descriptor           string
	statements           statement.Statement
}

func (d *ConstructorDeclaration) GetFlags() int {
	return d.flags
}

func (d *ConstructorDeclaration) SetFlags(flags int) {
	d.flags = flags
}

func (d *ConstructorDeclaration) IsStatic() bool {
	return d.flags&classfile.AccStatic != 0
}

func (d *ConstructorDeclaration) GetAnnotationReferences() reference.AnnotationReference {
	return d.annotationReferences
}

func (d *ConstructorDeclaration) GetTypeParameters() _type.ITypeParameter {
	return d.typeParameters
}

func (d *ConstructorDeclaration) GetFormalParameters() IFormalParameter {
	return d.formalParameter
}

func (d *ConstructorDeclaration) GetExceptionTypes() _type.IType {
	return d.exceptionTypes
}

func (d *ConstructorDeclaration) GetDescriptor() string {
	return d.descriptor
}

func (d *ConstructorDeclaration) GetStatements() statement.Statement {
	return d.statements
}

func (d *ConstructorDeclaration) Accept(visitor DeclarationVisitor) {
	visitor.VisitConstructorDeclaration(d)
}

func (d *ConstructorDeclaration) String() string {
	return fmt.Sprintf("ConstructorDeclaration{%s}", d.descriptor)
}
