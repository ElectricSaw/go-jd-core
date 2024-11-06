package declaration

import (
	"bitbucket.org/coontec/javaClass/java/model/classfile"
	"bitbucket.org/coontec/javaClass/java/model/javasyntax/reference"
	"bitbucket.org/coontec/javaClass/java/model/javasyntax/statement"
	_type "bitbucket.org/coontec/javaClass/java/model/javasyntax/type"
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

func NewMethodDeclaration6(annotationReferences reference.IAnnotationReference, flags int, name string, typeParameters _type.TypeParameter, returnedType _type.IType, formalParameter IFormalParameter, exceptionTypes _type.IType, descriptor string, statements statement.Statement, defaultAnnotationValue reference.IElementValue) *MethodDeclaration {
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
	typeParameters         _type.TypeParameter
	returnedType           _type.IType
	formalParameter        IFormalParameter
	exceptionTypes         _type.IType
	descriptor             string
	statements             statement.Statement
	defaultAnnotationValue reference.IElementValue
}

func (d *MethodDeclaration) GetFlags() int {
	return d.flags
}

func (d *MethodDeclaration) IsStatic() bool {
	return d.flags&classfile.AccStatic != 0
}

func (d *MethodDeclaration) GetName() string {
	return d.name
}

func (d *MethodDeclaration) GetTypeParameters() _type.TypeParameter {
	return d.typeParameters
}

func (d *MethodDeclaration) GetReturnType() _type.IType {
	return d.returnedType
}

func (d *MethodDeclaration) GetFormalParameter() IFormalParameter {
	return d.formalParameter
}

func (d *MethodDeclaration) GetExceptionTypes() _type.IType {
	return d.exceptionTypes
}

func (d *MethodDeclaration) GetDescriptor() string {
	return d.descriptor
}

func (d *MethodDeclaration) GetStatements() statement.Statement {
	return d.statements
}

func (d *MethodDeclaration) GetDefaultAnnotationValue() reference.IElementValue {
	return d.defaultAnnotationValue
}

func (d *MethodDeclaration) Accept(visitor DeclarationVisitor) {
	visitor.VisitMethodDeclaration(d)
}

func (d *MethodDeclaration) String() string {
	return fmt.Sprintf("MethodDeclaration{name=%s, descriptor=%s}", d.name, d.descriptor)
}
