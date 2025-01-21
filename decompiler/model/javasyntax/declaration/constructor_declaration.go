package declaration

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/interfaces/classpath"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewConstructorDeclaration(flags int, formalParameter FormalParameter, descriptor string,
	statements Statement) ConstructorDeclaration {
	return NewConstructorDeclarationWithAll(nil, flags, nil,
		formalParameter, nil, descriptor, statements)
}

func NewConstructorDeclarationWithAll(annotationReferences Reference, flags int,
	typeParameters TypeParameter, formalParameter FormalParameter,
	exceptionTypes Type, descriptor string, statements Statement) ConstructorDeclaration {
	d := &ConstructorDeclaration{
		annotationReferences: annotationReferences,
		flags:                flags,
		typeParameters:       typeParameters,
		formalParameter:      formalParameter,
		exceptionTypes:       exceptionTypes,
		descriptor:           descriptor,
		statements:           statements,
	}
	d.SetValue(d)
	return d
}

type ConstructorDeclaration struct {
	AbstractMemberDeclaration
	util.DefaultBase[MemberDeclaration]

	annotationReferences Reference
	flags                int
	typeParameters       TypeParameter
	formalParameter      FormalParameter
	exceptionTypes       Type
	descriptor           string
	statements           Statement
}

func (d *ConstructorDeclaration) Flags() int {
	return d.flags
}

func (d *ConstructorDeclaration) SetFlags(flags int) {
	d.flags = flags
}

func (d *ConstructorDeclaration) IsStatic() bool {
	return d.flags&classpath.AccStatic != 0
}

func (d *ConstructorDeclaration) AnnotationReferences() Reference {
	return d.annotationReferences
}

func (d *ConstructorDeclaration) TypeParameters() TypeParameter {
	return d.typeParameters
}

func (d *ConstructorDeclaration) FormalParameters() FormalParameter {
	return d.formalParameter
}

func (d *ConstructorDeclaration) SetFormalParameters(formalParameter FormalParameter) {
	d.formalParameter = formalParameter
}

func (d *ConstructorDeclaration) ExceptionTypes() Type {
	return d.exceptionTypes
}

func (d *ConstructorDeclaration) Descriptor() string {
	return d.descriptor
}

func (d *ConstructorDeclaration) Statements() Statement {
	return d.statements
}

func (d *ConstructorDeclaration) SetStatements(state Statement) {
	d.statements = state
}

func (d *ConstructorDeclaration) AcceptDeclaration(visitor DeclarationVisitor) {
	visitor.VisitConstructorDeclaration(d)
}

func (d *ConstructorDeclaration) String() string {
	return fmt.Sprintf("ConstructorDeclaration{%s}", d.descriptor)
}
