package declaration

import (
	"bitbucket.org/coontec/javaClass/class/model/classfile"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/declaration"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/reference"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
)

func NewClassFileConstructorDeclaration(
	bodyDeclaration ClassFileBodyDeclaration,
	classFile *classfile.ClassFile,
	method *classfile.Method,
	annotationReferences reference.AnnotationReference,
	typeParameters _type.ITypeParameter,
	parameterTypes _type.IType,
	exceptionTypes _type.IType,
	bindings map[string]_type.ITypeArgument,
	typeBounds map[string]_type.IType,
	firstLineNumber int) *ClassFileConstructorDeclaration {
	return &ClassFileConstructorDeclaration{
		ConstructorDeclaration: *declaration.NewConstructorDeclarationWithAll(
			annotationReferences, method.AccessFlags(), typeParameters,
			nil, exceptionTypes, method.Descriptor(), nil),
		bodyDeclaration: bodyDeclaration,
		classFile:       classFile,
		method:          method,
		parameterTypes:  parameterTypes,
		bindings:        bindings,
		typeBounds:      typeBounds,
		firstLineNumber: firstLineNumber,
	}
}

type ClassFileConstructorDeclaration struct {
	declaration.ConstructorDeclaration

	bodyDeclaration ClassFileBodyDeclaration
	classFile       *classfile.ClassFile
	method          *classfile.Method
	parameterTypes  _type.IType
	bindings        map[string]_type.ITypeArgument
	typeBounds      map[string]_type.IType
	firstLineNumber int
}

func (d *ClassFileConstructorDeclaration) ClassFile() *classfile.ClassFile {
	return d.classFile
}

func (d *ClassFileConstructorDeclaration) Method() *classfile.Method {
	return d.method
}

func (d *ClassFileConstructorDeclaration) ParameterTypes() _type.IType {
	return d.parameterTypes
}

func (d *ClassFileConstructorDeclaration) ReturnedType() _type.IType {
	return nil
}

func (d *ClassFileConstructorDeclaration) BodyDeclaration() ClassFileBodyDeclaration {
	return d.bodyDeclaration
}

func (d *ClassFileConstructorDeclaration) Bindings() map[string]_type.ITypeArgument {
	return d.bindings
}

func (d *ClassFileConstructorDeclaration) TypeBounds() map[string]_type.IType {
	return d.typeBounds
}

func (d *ClassFileConstructorDeclaration) FirstLineNumber() int {
	return d.firstLineNumber
}
