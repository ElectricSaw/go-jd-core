package declaration

import (
	"bitbucket.org/coontec/javaClass/class/model/classfile"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/reference"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/statement"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewMethodDeclaration(flags int, name string, returnedType _type.IType, descriptor string) *MethodDeclaration {
	return &MethodDeclaration{
		flags:        flags,
		name:         name,
		returnedType: returnedType,
		descriptor:   descriptor,
	}
}

func NewMethodDeclaration2(flags int, name string, returnedType _type.IType, descriptor string, statements statement.Statement) *MethodDeclaration {
	return &MethodDeclaration{
		flags:        flags,
		name:         name,
		returnedType: returnedType,
		descriptor:   descriptor,
		statements:   statements,
	}
}

func NewMethodDeclaration3(flags int, name string, returnedType _type.IType, descriptor string, defaultAnnotationValue reference.IElementValue) *MethodDeclaration {
	return &MethodDeclaration{
		flags:                  flags,
		name:                   name,
		returnedType:           returnedType,
		descriptor:             descriptor,
		defaultAnnotationValue: defaultAnnotationValue,
	}
}

func NewMethodDeclaration4(flags int, name string, returnedType _type.IType, formalParameter IFormalParameter, descriptor string, statements statement.Statement) *MethodDeclaration {
	return &MethodDeclaration{
		flags:           flags,
		name:            name,
		returnedType:    returnedType,
		formalParameter: formalParameter,
		descriptor:      descriptor,
		statements:      statements,
	}
}

func NewMethodDeclaration5(flags int, name string, returnedType _type.IType, formalParameter IFormalParameter, descriptor string, defaultAnnotationValue reference.IElementValue) *MethodDeclaration {
	return &MethodDeclaration{
		flags:                  flags,
		name:                   name,
		returnedType:           returnedType,
		formalParameter:        formalParameter,
		descriptor:             descriptor,
		defaultAnnotationValue: defaultAnnotationValue,
	}
}

func NewMethodDeclaration6(annotationReferences reference.IAnnotationReference, flags int, name string, typeParameters *_type.TypeParameter, returnedType _type.IType, formalParameter IFormalParameter, exceptionTypes _type.IType, descriptor string, statements statement.Statement, defaultAnnotationValue reference.IElementValue) *MethodDeclaration {
	return &MethodDeclaration{
		annotationReferences:   annotationReferences,
		flags:                  flags,
		name:                   name,
		typeParameters:         typeParameters,
		returnedType:           returnedType,
		formalParameter:        formalParameter,
		exceptionTypes:         exceptionTypes,
		descriptor:             descriptor,
		statements:             statements,
		defaultAnnotationValue: defaultAnnotationValue,
	}
}

type MethodDeclaration struct {
	AbstractMemberDeclaration

	annotationReferences   reference.IAnnotationReference
	flags                  int
	name                   string
	typeParameters         *_type.TypeParameter
	returnedType           _type.IType
	formalParameter        IFormalParameter
	exceptionTypes         _type.IType
	descriptor             string
	statements             statement.Statement
	defaultAnnotationValue reference.IElementValue
}

func (d *MethodDeclaration) Flags() int {
	return d.flags
}

func (d *MethodDeclaration) SetFlags(flags int) {
	d.flags = flags
}

func (d *MethodDeclaration) AnnotationReferences() reference.IAnnotationReference {
	return d.annotationReferences
}

func (d *MethodDeclaration) IsStatic() bool {
	return d.flags&classfile.AccStatic != 0
}

func (d *MethodDeclaration) Name() string {
	return d.name
}

func (d *MethodDeclaration) TypeParameters() _type.ITypeParameter {
	return d.typeParameters
}

func (d *MethodDeclaration) ReturnedType() _type.IType {
	return d.returnedType
}

func (d *MethodDeclaration) FormalParameter() IFormalParameter {
	return d.formalParameter
}

func (d *MethodDeclaration) SetFormalParameters(formalParameter IFormalParameter) {
	d.formalParameter = formalParameter
}

func (d *MethodDeclaration) ExceptionTypes() _type.IType {
	return d.exceptionTypes
}

func (d *MethodDeclaration) Descriptor() string {
	return d.descriptor
}

func (d *MethodDeclaration) Statements() statement.Statement {
	return d.statements
}

func (d *MethodDeclaration) SetStatements(statements statement.Statement) {
	d.statements = statements
}

func (d *MethodDeclaration) DefaultAnnotationValue() reference.IElementValue {
	return d.defaultAnnotationValue
}

func (d *MethodDeclaration) Accept(visitor DeclarationVisitor) {
	visitor.VisitMethodDeclaration(d)
}

func (d *MethodDeclaration) String() string {
	return fmt.Sprintf("MethodDeclaration{name=%s, descriptor=%s}", d.name, d.descriptor)
}
