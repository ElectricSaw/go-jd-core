package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewMethodDeclaration(flags int, name string, returnedType intmod.IType,
	descriptor string) intmod.IMethodDeclaration {
	return &MethodDeclaration{
		flags:        flags,
		name:         name,
		returnedType: returnedType,
		descriptor:   descriptor,
	}
}

func NewMethodDeclaration2(flags int, name string, returnedType intmod.IType,
	descriptor string, statements intmod.IStatement) intmod.IMethodDeclaration {
	return &MethodDeclaration{
		flags:        flags,
		name:         name,
		returnedType: returnedType,
		descriptor:   descriptor,
		statements:   statements,
	}
}

func NewMethodDeclaration3(flags int, name string, returnedType intmod.IType,
	descriptor string, defaultAnnotationValue intmod.IElementValue) intmod.IMethodDeclaration {
	return &MethodDeclaration{
		flags:                  flags,
		name:                   name,
		returnedType:           returnedType,
		descriptor:             descriptor,
		defaultAnnotationValue: defaultAnnotationValue,
	}
}

func NewMethodDeclaration4(flags int, name string, returnedType intmod.IType,
	formalParameter intmod.IFormalParameter, descriptor string,
	statements intmod.IStatement) intmod.IMethodDeclaration {
	return &MethodDeclaration{
		flags:           flags,
		name:            name,
		returnedType:    returnedType,
		formalParameter: formalParameter,
		descriptor:      descriptor,
		statements:      statements,
	}
}

func NewMethodDeclaration5(flags int, name string, returnedType intmod.IType,
	formalParameter intmod.IFormalParameter, descriptor string,
	defaultAnnotationValue intmod.IElementValue) intmod.IMethodDeclaration {
	return &MethodDeclaration{
		flags:                  flags,
		name:                   name,
		returnedType:           returnedType,
		formalParameter:        formalParameter,
		descriptor:             descriptor,
		defaultAnnotationValue: defaultAnnotationValue,
	}
}

func NewMethodDeclaration6(annotationReferences intmod.IAnnotationReference,
	flags int, name string, typeParameters intmod.ITypeParameter, returnedType intmod.IType,
	formalParameter intmod.IFormalParameter, exceptionTypes intmod.IType, descriptor string,
	statements intmod.IStatement, defaultAnnotationValue intmod.IElementValue) intmod.IMethodDeclaration {
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

	annotationReferences   intmod.IAnnotationReference
	flags                  int
	name                   string
	typeParameters         intmod.ITypeParameter
	returnedType           intmod.IType
	formalParameter        intmod.IFormalParameter
	exceptionTypes         intmod.IType
	descriptor             string
	statements             intmod.IStatement
	defaultAnnotationValue intmod.IElementValue
}

func (d *MethodDeclaration) Flags() int {
	return d.flags
}

func (d *MethodDeclaration) SetFlags(flags int) {
	d.flags = flags
}

func (d *MethodDeclaration) AnnotationReferences() intmod.IAnnotationReference {
	return d.annotationReferences
}

func (d *MethodDeclaration) IsStatic() bool {
	return d.flags&intmod.AccStatic != 0
}

func (d *MethodDeclaration) Name() string {
	return d.name
}

func (d *MethodDeclaration) TypeParameters() intmod.ITypeParameter {
	return d.typeParameters
}

func (d *MethodDeclaration) ReturnedType() intmod.IType {
	return d.returnedType
}

func (d *MethodDeclaration) FormalParameter() intmod.IFormalParameter {
	return d.formalParameter
}

func (d *MethodDeclaration) SetFormalParameters(formalParameter intmod.IFormalParameter) {
	d.formalParameter = formalParameter
}

func (d *MethodDeclaration) ExceptionTypes() intmod.IType {
	return d.exceptionTypes
}

func (d *MethodDeclaration) Descriptor() string {
	return d.descriptor
}

func (d *MethodDeclaration) Statements() intmod.IStatement {
	return d.statements
}

func (d *MethodDeclaration) SetStatements(statements intmod.IStatement) {
	d.statements = statements
}

func (d *MethodDeclaration) DefaultAnnotationValue() intmod.IElementValue {
	return d.defaultAnnotationValue
}

func (d *MethodDeclaration) Accept(visitor intmod.IDeclarationVisitor) {
	visitor.VisitMethodDeclaration(d)
}

func (d *MethodDeclaration) String() string {
	return fmt.Sprintf("MethodDeclaration{name=%s, descriptor=%s}", d.name, d.descriptor)
}
