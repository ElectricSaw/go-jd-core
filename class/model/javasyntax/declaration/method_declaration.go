package declaration

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/model/classfile"
	"fmt"
)

func NewMethodDeclaration(flags int, name string, returnedType intsyn.IType,
	descriptor string) intsyn.IMethodDeclaration {
	return &MethodDeclaration{
		flags:        flags,
		name:         name,
		returnedType: returnedType,
		descriptor:   descriptor,
	}
}

func NewMethodDeclaration2(flags int, name string, returnedType intsyn.IType,
	descriptor string, statements intsyn.IStatement) intsyn.IMethodDeclaration {
	return &MethodDeclaration{
		flags:        flags,
		name:         name,
		returnedType: returnedType,
		descriptor:   descriptor,
		statements:   statements,
	}
}

func NewMethodDeclaration3(flags int, name string, returnedType intsyn.IType,
	descriptor string, defaultAnnotationValue intsyn.IElementValue) intsyn.IMethodDeclaration {
	return &MethodDeclaration{
		flags:                  flags,
		name:                   name,
		returnedType:           returnedType,
		descriptor:             descriptor,
		defaultAnnotationValue: defaultAnnotationValue,
	}
}

func NewMethodDeclaration4(flags int, name string, returnedType intsyn.IType,
	formalParameter intsyn.IFormalParameter, descriptor string,
	statements intsyn.IStatement) intsyn.IMethodDeclaration {
	return &MethodDeclaration{
		flags:           flags,
		name:            name,
		returnedType:    returnedType,
		formalParameter: formalParameter,
		descriptor:      descriptor,
		statements:      statements,
	}
}

func NewMethodDeclaration5(flags int, name string, returnedType intsyn.IType,
	formalParameter intsyn.IFormalParameter, descriptor string,
	defaultAnnotationValue intsyn.IElementValue) intsyn.IMethodDeclaration {
	return &MethodDeclaration{
		flags:                  flags,
		name:                   name,
		returnedType:           returnedType,
		formalParameter:        formalParameter,
		descriptor:             descriptor,
		defaultAnnotationValue: defaultAnnotationValue,
	}
}

func NewMethodDeclaration6(annotationReferences intsyn.IAnnotationReference,
	flags int, name string, typeParameters intsyn.ITypeParameter, returnedType intsyn.IType,
	formalParameter intsyn.IFormalParameter, exceptionTypes intsyn.IType, descriptor string,
	statements intsyn.IStatement, defaultAnnotationValue intsyn.IElementValue) intsyn.IMethodDeclaration {
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

	annotationReferences   intsyn.IAnnotationReference
	flags                  int
	name                   string
	typeParameters         intsyn.ITypeParameter
	returnedType           intsyn.IType
	formalParameter        intsyn.IFormalParameter
	exceptionTypes         intsyn.IType
	descriptor             string
	statements             intsyn.IStatement
	defaultAnnotationValue intsyn.IElementValue
}

func (d *MethodDeclaration) Flags() int {
	return d.flags
}

func (d *MethodDeclaration) SetFlags(flags int) {
	d.flags = flags
}

func (d *MethodDeclaration) AnnotationReferences() intsyn.IAnnotationReference {
	return d.annotationReferences
}

func (d *MethodDeclaration) IsStatic() bool {
	return d.flags&classfile.AccStatic != 0
}

func (d *MethodDeclaration) Name() string {
	return d.name
}

func (d *MethodDeclaration) TypeParameters() intsyn.ITypeParameter {
	return d.typeParameters
}

func (d *MethodDeclaration) ReturnedType() intsyn.IType {
	return d.returnedType
}

func (d *MethodDeclaration) FormalParameter() intsyn.IFormalParameter {
	return d.formalParameter
}

func (d *MethodDeclaration) SetFormalParameters(formalParameter intsyn.IFormalParameter) {
	d.formalParameter = formalParameter
}

func (d *MethodDeclaration) ExceptionTypes() intsyn.IType {
	return d.exceptionTypes
}

func (d *MethodDeclaration) Descriptor() string {
	return d.descriptor
}

func (d *MethodDeclaration) Statements() intsyn.IStatement {
	return d.statements
}

func (d *MethodDeclaration) SetStatements(statements intsyn.IStatement) {
	d.statements = statements
}

func (d *MethodDeclaration) DefaultAnnotationValue() intsyn.IElementValue {
	return d.defaultAnnotationValue
}

func (d *MethodDeclaration) Accept(visitor intsyn.IDeclarationVisitor) {
	visitor.VisitMethodDeclaration(d)
}

func (d *MethodDeclaration) String() string {
	return fmt.Sprintf("MethodDeclaration{name=%s, descriptor=%s}", d.name, d.descriptor)
}
