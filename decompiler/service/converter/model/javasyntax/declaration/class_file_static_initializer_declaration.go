package declaration

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/classfile"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewClassFileStaticInitializerDeclaration(bodyDeclaration *ClassFileBodyDeclaration,
	classFile *classfile.ClassFile, method *classfile.Method,
	bindings map[string]model.ITypeArgument, typeBounds map[string]model.IType,
	firstLineNumber int) ClassFileStaticInitializerDeclaration {
	return NewClassFileStaticInitializerDeclaration2(bodyDeclaration, classFile, method,
		bindings, typeBounds, firstLineNumber, nil)
}

func NewClassFileStaticInitializerDeclaration2(bodyDeclaration *ClassFileBodyDeclaration, classFile *classfile.ClassFile,
	method *classfile.Method, bindings map[string]model.ITypeArgument, typeBounds map[string]model.IType,
	firstLineNumber int, statements model.IStatement) ClassFileStaticInitializerDeclaration {
	d := ClassFileStaticInitializerDeclaration{
		Descriptor:      method.Descriptor,
		Statements:      statements,
		BodyDeclaration: bodyDeclaration,
		ClassFile:       classFile,
		Method:          method,
		Bindings:        bindings,
		TypeBounds:      typeBounds,
		FirstLineNumber: firstLineNumber,
	}
	d.SetValue(&d)
	return d
}

type ClassFileStaticInitializerDeclaration struct {
	util.DefaultBase[model.IMemberDeclaration]

	Descriptor      string
	Statements      model.IStatement
	BodyDeclaration *ClassFileBodyDeclaration
	ClassFile       *classfile.ClassFile
	Method          *classfile.Method
	Bindings        map[string]model.ITypeArgument
	TypeBounds      map[string]model.IType
	FirstLineNumber int
}

func (d *ClassFileStaticInitializerDeclaration) GetAnnotationReferences() *model.AnnotationReference {
	return nil
}

func (d *ClassFileStaticInitializerDeclaration) GetFlags() int {
	return 0
}

func (d *ClassFileStaticInitializerDeclaration) GetTypeParameters() model.ITypeParameter {
	return nil
}

func (d *ClassFileStaticInitializerDeclaration) GetFormalParameters() *model.FormalParameter {
	return nil
}

func (d *ClassFileStaticInitializerDeclaration) GetExceptionTypes() model.IType {
	return nil
}

func (d *ClassFileStaticInitializerDeclaration) GetDescriptor() string {
	return d.Descriptor
}

func (d *ClassFileStaticInitializerDeclaration) GetStatements() model.IStatement {
	return d.Statements
}

func (d *ClassFileStaticInitializerDeclaration) GetBodyDeclaration() *ClassFileBodyDeclaration {
	return d.BodyDeclaration
}

func (d *ClassFileStaticInitializerDeclaration) GetClassFile() *classfile.ClassFile {
	return d.ClassFile
}

func (d *ClassFileStaticInitializerDeclaration) GetMethod() *classfile.Method {
	return d.Method
}

func (d *ClassFileStaticInitializerDeclaration) GetParameterTypes() model.IType {
	return nil
}

func (d *ClassFileStaticInitializerDeclaration) GetBindings() map[string]model.ITypeArgument {
	return d.Bindings
}

func (d *ClassFileStaticInitializerDeclaration) GetTypeBounds() map[string]model.IType {
	return d.TypeBounds
}

func (d *ClassFileStaticInitializerDeclaration) GetReturnedType() model.IType {
	return nil
}

func (d *ClassFileStaticInitializerDeclaration) GetFirstLineNumber() int {
	return d.FirstLineNumber
}

func (d *ClassFileStaticInitializerDeclaration) SetFlags(flags int) {
	// EMPTY
}

func (d *ClassFileStaticInitializerDeclaration) SetFormalParameters(formalParameters *model.FormalParameter) {
	// EMPTY
}

func (d *ClassFileStaticInitializerDeclaration) SetStatements(statement model.IStatement) {
	d.Statements = statement
}

func (d *ClassFileStaticInitializerDeclaration) IsClassDeclaration() bool {
	return false
}

func (d *ClassFileStaticInitializerDeclaration) AcceptDeclaration(visitor model.IDeclarationVisitor) {
	visitor.VisitStaticInitializerDeclaration(d)
}

func (d *ClassFileStaticInitializerDeclaration) String() string {
	return fmt.Sprintf("ClassFileStaticInitializerDeclaration{%s, firstLineNumber=%d}", d.Descriptor, d.FirstLineNumber)
}
