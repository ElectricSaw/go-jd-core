package declaration

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/interfaces/classpath"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewMethodDeclaration(flags int, name string, returnedType Type,
	descriptor string) MethodDeclaration {
	return NewMethodDeclaration6(nil, flags, name, nil,
		returnedType, nil, nil, descriptor, nil, nil)
}

func NewMethodDeclaration2(flags int, name string, returnedType Type,
	descriptor string, statements Statement) MethodDeclaration {
	return NewMethodDeclaration6(nil, flags, name, nil,
		returnedType, nil, nil, descriptor, statements, nil)
}

func NewMethodDeclaration3(flags int, name string, returnedType Type,
	descriptor string, defaultAnnotationValue ElementValue) MethodDeclaration {
	return NewMethodDeclaration6(nil, flags, name, nil,
		returnedType, nil, nil, descriptor, nil, defaultAnnotationValue)
}

func NewMethodDeclaration4(flags int, name string, returnedType Type,
	formalParameter FormalParameter, descriptor string,
	statements Statement) MethodDeclaration {
	return NewMethodDeclaration6(nil, flags, name, nil,
		returnedType, formalParameter, nil, descriptor, statements, nil)
}

func NewMethodDeclaration5(flags int, name string, returnedType Type,
	formalParameter FormalParameter, descriptor string,
	defaultAnnotationValue ElementValue) MethodDeclaration {
	return NewMethodDeclaration6(nil, flags, name, nil,
		returnedType, formalParameter, nil, descriptor, nil, defaultAnnotationValue)
}

func NewMethodDeclaration6(annotationReferences AnnotationReference,
	flags int, name string, typeParameters TypeParameter, returnedType Type,
	formalParameter FormalParameter, exceptionTypes Type, descriptor string,
	statements Statement, defaultAnnotationValue ElementValue) MethodDeclaration {
	d := &MethodDeclaration{
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
	d.SetValue(d)
	return d
}

type MethodDeclaration struct {
	AbstractMemberDeclaration

	annotationReferences   AnnotationReference
	flags                  int
	name                   string
	typeParameters         TypeParameter
	returnedType           Type
	formalParameter        FormalParameter
	exceptionTypes         Type
	descriptor             string
	statements             Statement
	defaultAnnotationValue ElementValue
}

func (d *MethodDeclaration) Flags() int {
	return d.flags
}

func (d *MethodDeclaration) SetFlags(flags int) {
	d.flags = flags
}

func (d *MethodDeclaration) AnnotationReferences() AnnotationReference {
	return d.annotationReferences
}

func (d *MethodDeclaration) IsStatic() bool {
	return d.flags&classpath.AccStatic != 0
}

func (d *MethodDeclaration) Name() string {
	return d.name
}

func (d *MethodDeclaration) TypeParameters() TypeParameter {
	return d.typeParameters
}

func (d *MethodDeclaration) ReturnedType() Type {
	return d.returnedType
}

func (d *MethodDeclaration) FormalParameters() FormalParameter {
	return d.formalParameter
}

func (d *MethodDeclaration) SetFormalParameters(formalParameter FormalParameter) {
	d.formalParameter = formalParameter
}

func (d *MethodDeclaration) ExceptionTypes() Type {
	return d.exceptionTypes
}

func (d *MethodDeclaration) Descriptor() string {
	return d.descriptor
}

func (d *MethodDeclaration) Statements() Statement {
	return d.statements
}

func (d *MethodDeclaration) SetStatements(statements Statement) {
	d.statements = statements
}

func (d *MethodDeclaration) DefaultAnnotationValue() ElementValue {
	return d.defaultAnnotationValue
}

func (d *MethodDeclaration) AcceptDeclaration(visitor DeclarationVisitor) {
	visitor.VisitMethodDeclaration(d)
}

func (d *MethodDeclaration) String() string {
	return fmt.Sprintf("MethodDeclaration{name=%s, descriptor=%s}", d.name, d.descriptor)
}
