package declaration

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/classfile"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewClassFileConstructorDeclaration(
	bodyDeclaration *ClassFileBodyDeclaration,
	classFile *classfile.ClassFile,
	method *classfile.Method,
	annotationReferences *model.AnnotationReference,
	typeParameters model.ITypeParameter,
	parameterTypes, exceptionTypes model.IType,
	bindings map[string]model.ITypeArgument,
	typeBounds map[string]model.IType,
	firstLineNumber int) ClassFileConstructorDeclaration {
	d := ClassFileConstructorDeclaration{
		DefaultBase:          *util.NewDefaultBase[model.IMemberDeclaration]().(*util.DefaultBase[model.IMemberDeclaration]),
		AnnotationReferences: annotationReferences,
		Flags:                method.AccessFlags,
		TypeParameters:       typeParameters,
		ExceptionTypes:       exceptionTypes,
		Descriptor:           method.Descriptor,
		BodyDeclaration:      bodyDeclaration,
		ClassFile:            classFile,
		Method:               method,
		ParameterTypes:       parameterTypes,
		Bindings:             bindings,
		TypeBounds:           typeBounds,
		FirstLineNumber:      firstLineNumber,
	}
	d.SetValue(&d)
	return d
}

type ClassFileConstructorDeclaration struct {
	util.DefaultBase[model.IMemberDeclaration]

	AnnotationReferences *model.AnnotationReference
	Flags                int
	TypeParameters       model.ITypeParameter
	FormalParameters     *model.FormalParameter
	ExceptionTypes       model.IType
	Descriptor           string
	Statements           model.IStatement
	BodyDeclaration      *ClassFileBodyDeclaration
	ClassFile            *classfile.ClassFile
	Method               *classfile.Method
	ParameterTypes       model.IType
	Bindings             map[string]model.ITypeArgument
	TypeBounds           map[string]model.IType
	FirstLineNumber      int
}

func (d *ClassFileConstructorDeclaration) GetAnnotationReferences() *model.AnnotationReference {
	return d.AnnotationReferences
}

func (d *ClassFileConstructorDeclaration) GetFlags() int {
	return d.Flags
}

func (d *ClassFileConstructorDeclaration) GetTypeParameters() model.ITypeParameter {
	return d.TypeParameters
}

func (d *ClassFileConstructorDeclaration) GetFormalParameters() *model.FormalParameter {
	return d.FormalParameters
}

func (d *ClassFileConstructorDeclaration) GetExceptionTypes() model.IType {
	return d.ExceptionTypes
}

func (d *ClassFileConstructorDeclaration) GetDescriptor() string {
	return d.Descriptor
}

func (d *ClassFileConstructorDeclaration) GetStatements() model.IStatement {
	return d.Statements
}

func (d *ClassFileConstructorDeclaration) GetBodyDeclaration() *ClassFileBodyDeclaration {
	return d.BodyDeclaration
}

func (d *ClassFileConstructorDeclaration) GetClassFile() *classfile.ClassFile {
	return d.ClassFile
}

func (d *ClassFileConstructorDeclaration) GetMethod() *classfile.Method {
	return d.Method
}

func (d *ClassFileConstructorDeclaration) GetParameterTypes() model.IType {
	return d.ParameterTypes
}

func (d *ClassFileConstructorDeclaration) GetBindings() map[string]model.ITypeArgument {
	return d.Bindings
}

func (d *ClassFileConstructorDeclaration) GetTypeBounds() map[string]model.IType {
	return d.TypeBounds
}

func (d *ClassFileConstructorDeclaration) GetReturnedType() model.IType {
	return nil
}

func (d *ClassFileConstructorDeclaration) GetFirstLineNumber() int {
	return d.FirstLineNumber
}

func (d *ClassFileConstructorDeclaration) SetFlags(flags int) {
	d.Flags = flags
}

func (d *ClassFileConstructorDeclaration) SetFormalParameters(formalParameters *model.FormalParameter) {
	d.FormalParameters = formalParameters
}

func (d *ClassFileConstructorDeclaration) SetStatements(statement model.IStatement) {
	d.Statements = statement
}

func (d *ClassFileConstructorDeclaration) IsClassDeclaration() bool {
	return false
}

func (d *ClassFileConstructorDeclaration) AcceptDeclaration(visitor model.IDeclarationVisitor) {
	visitor.VisitConstructorDeclaration(d)
}

func (d *ClassFileConstructorDeclaration) String() string {
	return fmt.Sprintf("ClassFileConstructorDeclaration{ %s %s, first-line-number=%d }",
		d.ClassFile.InternalTypeName, d.Descriptor, d.FirstLineNumber)
}
