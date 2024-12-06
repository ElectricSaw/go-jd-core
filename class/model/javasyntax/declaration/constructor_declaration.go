package declaration

import (
	"bitbucket.org/coontec/go-jd-core/class/interfaces/classpath"
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
	"fmt"
)

func NewConstructorDeclaration(flags int, formalParameter intmod.IFormalParameter, descriptor string,
	statements intmod.IStatement) intmod.IConstructorDeclaration {
	return NewConstructorDeclarationWithAll(nil, flags, nil,
		formalParameter, nil, descriptor, statements)
}

func NewConstructorDeclarationWithAll(annotationReferences intmod.IReference, flags int,
	typeParameters intmod.ITypeParameter, formalParameter intmod.IFormalParameter,
	exceptionTypes intmod.IType, descriptor string, statements intmod.IStatement) intmod.IConstructorDeclaration {
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
	util.DefaultBase[intmod.IMemberDeclaration]

	annotationReferences intmod.IReference
	flags                int
	typeParameters       intmod.ITypeParameter
	formalParameter      intmod.IFormalParameter
	exceptionTypes       intmod.IType
	descriptor           string
	statements           intmod.IStatement
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

func (d *ConstructorDeclaration) AnnotationReferences() intmod.IReference {
	return d.annotationReferences
}

func (d *ConstructorDeclaration) TypeParameters() intmod.ITypeParameter {
	return d.typeParameters
}

func (d *ConstructorDeclaration) FormalParameters() intmod.IFormalParameter {
	return d.formalParameter
}

func (d *ConstructorDeclaration) SetFormalParameters(formalParameter intmod.IFormalParameter) {
	d.formalParameter = formalParameter
}

func (d *ConstructorDeclaration) ExceptionTypes() intmod.IType {
	return d.exceptionTypes
}

func (d *ConstructorDeclaration) Descriptor() string {
	return d.descriptor
}

func (d *ConstructorDeclaration) Statements() intmod.IStatement {
	return d.statements
}

func (d *ConstructorDeclaration) SetStatements(state intmod.IStatement) {
	d.statements = state
}

func (d *ConstructorDeclaration) AcceptDeclaration(visitor intmod.IDeclarationVisitor) {
	visitor.VisitConstructorDeclaration(d)
}

func (d *ConstructorDeclaration) String() string {
	return fmt.Sprintf("ConstructorDeclaration{%s}", d.descriptor)
}
