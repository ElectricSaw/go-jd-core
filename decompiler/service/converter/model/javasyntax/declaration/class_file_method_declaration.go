package declaration

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/classfile"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewClassFileMethodDeclaration(bodyDeclaration *ClassFileBodyDeclaration, classFile *classfile.ClassFile,
	method *classfile.Method, name string, returnedType model.IType, parameterTypes model.IType,
	bindings map[string]model.ITypeArgument, typeBounds map[string]model.IType) ClassFileMethodDeclaration {
	return NewClassFileMethodDeclaration3(bodyDeclaration, classFile, method, nil, name,
		nil, returnedType, parameterTypes, nil,
		nil, bindings, typeBounds, -1)
}

func NewClassFileMethodDeclaration2(bodyDeclaration *ClassFileBodyDeclaration, classFile *classfile.ClassFile,
	method *classfile.Method, name string, returnedType model.IType, parameterTypes model.IType,
	bindings map[string]model.ITypeArgument, typeBounds map[string]model.IType, firstLineNumber int) ClassFileMethodDeclaration {
	return NewClassFileMethodDeclaration3(bodyDeclaration, classFile, method, nil, name,
		nil, returnedType, parameterTypes, nil,
		nil, bindings, typeBounds, firstLineNumber)
}

func NewClassFileMethodDeclaration3(bodyDeclaration *ClassFileBodyDeclaration, classFile *classfile.ClassFile,
	method *classfile.Method, annotationReferences *model.AnnotationReference, name string,
	typeParameters model.ITypeParameter, returnedType model.IType, parameterTypes model.IType,
	exceptionTypes model.IType, defaultAnnotationValue model.IElementValue,
	bindings map[string]model.ITypeArgument, typeBounds map[string]model.IType, firstLineNumber int) ClassFileMethodDeclaration {
	d := ClassFileMethodDeclaration{
		DefaultBase:            *util.NewDefaultBase[model.IMemberDeclaration]().(*util.DefaultBase[model.IMemberDeclaration]),
		AnnotationReferences:   annotationReferences,
		Flags:                  method.AccessFlags,
		Name:                   name,
		TypeParameters:         typeParameters,
		ReturnedType:           returnedType,
		ExceptionTypes:         exceptionTypes,
		Descriptor:             method.Descriptor,
		DefaultAnnotationValue: defaultAnnotationValue,
		BodyDeclaration:        bodyDeclaration,
		ClassFile:              classFile,
		ParameterTypes:         parameterTypes,
		Method:                 method,
		Bindings:               bindings,
		TypeBounds:             typeBounds,
		FirstLineNumber:        firstLineNumber,
	}
	d.SetValue(&d)
	return d
}

type ClassFileMethodDeclaration struct {
	util.DefaultBase[model.IMemberDeclaration]

	AnnotationReferences   *model.AnnotationReference
	Flags                  int
	Name                   string
	TypeParameters         model.ITypeParameter
	ReturnedType           model.IType
	FormalParameters       *model.FormalParameter
	ExceptionTypes         model.IType
	Descriptor             string
	Statements             model.IStatement
	DefaultAnnotationValue model.IElementValue
	BodyDeclaration        *ClassFileBodyDeclaration
	ClassFile              *classfile.ClassFile
	Method                 *classfile.Method
	ParameterTypes         model.IType
	Bindings               map[string]model.ITypeArgument
	TypeBounds             map[string]model.IType
	FirstLineNumber        int
}

func (d *ClassFileMethodDeclaration) GetDefaultAnnotationValue() model.IElementValue {
	return d.DefaultAnnotationValue
}

func (d *ClassFileMethodDeclaration) GetAnnotationReferences() *model.AnnotationReference {
	return d.AnnotationReferences
}

func (d *ClassFileMethodDeclaration) GetFlags() int {
	return d.Flags
}

func (d *ClassFileMethodDeclaration) GetName() string {
	return d.Name
}

func (d *ClassFileMethodDeclaration) GetTypeParameters() model.ITypeParameter {
	return d.TypeParameters
}

func (d *ClassFileMethodDeclaration) GetFormalParameters() *model.FormalParameter {
	return d.FormalParameters
}

func (d *ClassFileMethodDeclaration) GetExceptionTypes() model.IType {
	return d.ExceptionTypes
}

func (d *ClassFileMethodDeclaration) GetDescriptor() string {
	return d.Descriptor
}

func (d *ClassFileMethodDeclaration) GetStatements() model.IStatement {
	return d.Statements
}

func (d *ClassFileMethodDeclaration) GetBodyDeclaration() *ClassFileBodyDeclaration {
	return d.BodyDeclaration
}

func (d *ClassFileMethodDeclaration) GetClassFile() *classfile.ClassFile {
	return d.ClassFile
}

func (d *ClassFileMethodDeclaration) GetMethod() *classfile.Method {
	return d.Method
}

func (d *ClassFileMethodDeclaration) GetParameterTypes() model.IType {
	return d.ParameterTypes
}

func (d *ClassFileMethodDeclaration) GetBindings() map[string]model.ITypeArgument {
	return d.Bindings
}

func (d *ClassFileMethodDeclaration) GetTypeBounds() map[string]model.IType {
	return d.TypeBounds
}

func (d *ClassFileMethodDeclaration) GetReturnedType() model.IType {
	return d.ReturnedType
}

func (d *ClassFileMethodDeclaration) GetFirstLineNumber() int {
	return d.FirstLineNumber
}

func (d *ClassFileMethodDeclaration) SetFlags(flags int) {
	d.Flags = flags
}

func (d *ClassFileMethodDeclaration) SetFormalParameters(formalParameters *model.FormalParameter) {
	d.FormalParameters = formalParameters
}

func (d *ClassFileMethodDeclaration) SetStatements(statement model.IStatement) {
	d.Statements = statement
}

func (d *ClassFileMethodDeclaration) IsStatic() bool {
	return d.Flags&classfile.AccStatic != 0
}

func (d *ClassFileMethodDeclaration) IsClassDeclaration() bool {
	return false
}

func (d *ClassFileMethodDeclaration) AcceptDeclaration(visitor model.IDeclarationVisitor) {
	visitor.VisitMethodDeclaration(d)
}

func (d *ClassFileMethodDeclaration) String() string {
	return fmt.Sprintf("ClassFileMethodDeclaration{%s %s, firstLineNumber=%d}", d.Name, d.Descriptor, d.FirstLineNumber)
}
