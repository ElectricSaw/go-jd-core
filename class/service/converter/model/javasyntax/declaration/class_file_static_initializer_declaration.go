package declaration

import (
	"bitbucket.org/coontec/javaClass/class/model/classfile"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/declaration"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/statement"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewClassFileStaticInitializerDeclaration(bodyDeclaration ClassFileBodyDeclaration, classFile *classfile.ClassFile,
	method *classfile.Method, bindings map[string]_type.ITypeArgument, typeBounds map[string]_type.IType, firstLineNumber int) *ClassFileStaticInitializerDeclaration {
	return &ClassFileStaticInitializerDeclaration{
		StaticInitializerDeclaration: *declaration.NewStaticInitializerDeclaration(method.Descriptor(), nil),
		bodyDeclaration:              bodyDeclaration,
		classFile:                    classFile,
		method:                       method,
		bindings:                     bindings,
		typeBounds:                   typeBounds,
		firstLineNumber:              firstLineNumber,
	}
}

func NewClassFileStaticInitializerDeclaration2(bodyDeclaration ClassFileBodyDeclaration, classFile *classfile.ClassFile,
	method *classfile.Method, bindings map[string]_type.ITypeArgument, typeBounds map[string]_type.IType,
	firstLineNumber int, statements statement.IStatement) *ClassFileStaticInitializerDeclaration {
	return &ClassFileStaticInitializerDeclaration{
		StaticInitializerDeclaration: *declaration.NewStaticInitializerDeclaration(method.Descriptor(), statements),
		bodyDeclaration:              bodyDeclaration,
		classFile:                    classFile,
		method:                       method,
		bindings:                     bindings,
		typeBounds:                   typeBounds,
		firstLineNumber:              firstLineNumber,
	}
}

type ClassFileStaticInitializerDeclaration struct {
	declaration.StaticInitializerDeclaration

	bodyDeclaration ClassFileBodyDeclaration
	classFile       *classfile.ClassFile
	method          *classfile.Method
	bindings        map[string]_type.ITypeArgument
	typeBounds      map[string]_type.IType
	firstLineNumber int
}

func (d *ClassFileStaticInitializerDeclaration) Flags() int {
	return 0
}

func (d *ClassFileStaticInitializerDeclaration) ClassFile() *classfile.ClassFile {
	return d.classFile
}

func (d *ClassFileStaticInitializerDeclaration) Method() *classfile.Method {
	return d.method
}

func (d *ClassFileStaticInitializerDeclaration) TypeParameters() _type.ITypeParameter {
	return nil
}

func (d *ClassFileStaticInitializerDeclaration) ParameterTypes() _type.IType {
	return nil
}

func (d *ClassFileStaticInitializerDeclaration) ReturnedType() _type.IType {
	return nil
}

func (d *ClassFileStaticInitializerDeclaration) BodyDeclaration() ClassFileBodyDeclaration {
	return d.bodyDeclaration
}

func (d *ClassFileStaticInitializerDeclaration) Bindings() map[string]_type.ITypeArgument {
	return d.bindings
}

func (d *ClassFileStaticInitializerDeclaration) TypeBounds() map[string]_type.IType {
	return d.typeBounds
}

func (d *ClassFileStaticInitializerDeclaration) SetFlags(flags int) {
}

func (d *ClassFileStaticInitializerDeclaration) SetFormalParameters(formalParameters declaration.IFormalParameter) {
}

func (d *ClassFileStaticInitializerDeclaration) FirstLineNumber() int {
	return d.firstLineNumber
}

func (d *ClassFileStaticInitializerDeclaration) SetFirstLineNumber(lineNumber int) {
	d.firstLineNumber = lineNumber
}

func (d *ClassFileStaticInitializerDeclaration) String() string {
	return fmt.Sprintf("ClassFileStaticInitializerDeclaration{%s, firstLineNumber=%d}", d.Description(), d.firstLineNumber)
}
