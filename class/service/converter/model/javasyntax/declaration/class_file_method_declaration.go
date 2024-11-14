package declaration

import (
	"bitbucket.org/coontec/javaClass/class/model/classfile"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/declaration"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/reference"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewClassFileMethodDeclaration(bodyDeclaration ClassFileBodyDeclaration, classFile *classfile.ClassFile,
	method *classfile.Method, name string, returnedType _type.IType, parameterTypes _type.IType,
	bindings map[string]_type.ITypeArgument, typeBounds map[string]_type.IType) *ClassFileMethodDeclaration {
	return &ClassFileMethodDeclaration{
		MethodDeclaration: *declaration.NewMethodDeclaration6(nil, method.AccessFlags(), name, nil, returnedType,
			nil, nil, method.Descriptor(), nil, nil),
		bodyDeclaration: bodyDeclaration,
		classFile:       classFile,
		parameterTypes:  parameterTypes,
		method:          method,
		bindings:        bindings,
		typeBounds:      typeBounds,
	}
}

func NewClassFileMethodDeclaration2(bodyDeclaration ClassFileBodyDeclaration, classFile *classfile.ClassFile,
	method *classfile.Method, name string, returnedType _type.IType, parameterTypes _type.IType,
	bindings map[string]_type.ITypeArgument, typeBounds map[string]_type.IType, firstLineNumber int) *ClassFileMethodDeclaration {
	return &ClassFileMethodDeclaration{
		MethodDeclaration: *declaration.NewMethodDeclaration6(nil, method.AccessFlags(), name, nil, returnedType,
			nil, nil, method.Descriptor(), nil, nil),
		bodyDeclaration: bodyDeclaration,
		classFile:       classFile,
		parameterTypes:  parameterTypes,
		method:          method,
		bindings:        bindings,
		typeBounds:      typeBounds,
		firstLineNumber: firstLineNumber,
	}
}

func NewClassFileMethodDeclaration3(bodyDeclaration ClassFileBodyDeclaration, classFile *classfile.ClassFile,
	method *classfile.Method, annotationReferences reference.IAnnotationReference, name string,
	typeParameters *_type.TypeParameter, returnedType _type.IType, parameterTypes _type.IType,
	exceptionTypes _type.IType, defaultAnnotationValue reference.IElementValue,
	bindings map[string]_type.ITypeArgument, typeBounds map[string]_type.IType, firstLineNumber int) *ClassFileMethodDeclaration {
	return &ClassFileMethodDeclaration{
		MethodDeclaration: *declaration.NewMethodDeclaration6(annotationReferences, method.AccessFlags(), name, typeParameters,
			returnedType, nil, exceptionTypes, method.Descriptor(), nil, defaultAnnotationValue),
		bodyDeclaration: bodyDeclaration,
		classFile:       classFile,
		parameterTypes:  parameterTypes,
		method:          method,
		bindings:        bindings,
		typeBounds:      typeBounds,
		firstLineNumber: firstLineNumber,
	}
}

type ClassFileMethodDeclaration struct {
	declaration.MethodDeclaration

	bodyDeclaration ClassFileBodyDeclaration
	classFile       *classfile.ClassFile
	method          *classfile.Method
	parameterTypes  _type.IType
	bindings        map[string]_type.ITypeArgument
	typeBounds      map[string]_type.IType
	firstLineNumber int
}

func (d *ClassFileMethodDeclaration) ClassFile() *classfile.ClassFile {
	return d.classFile
}

func (d *ClassFileMethodDeclaration) Method() *classfile.Method {
	return d.method
}

func (d *ClassFileMethodDeclaration) ParameterTypes() _type.IType {
	return d.parameterTypes
}

func (d *ClassFileMethodDeclaration) BodyDeclaration() ClassFileBodyDeclaration {
	return d.bodyDeclaration
}

func (d *ClassFileMethodDeclaration) Bindings() map[string]_type.ITypeArgument {
	return d.bindings
}

func (d *ClassFileMethodDeclaration) TypeBounds() map[string]_type.IType {
	return d.typeBounds
}

func (d *ClassFileMethodDeclaration) FirstLineNumber() int {
	return d.firstLineNumber
}

func (d *ClassFileMethodDeclaration) String() string {
	return fmt.Sprintf("ClassFileMethodDeclaration{%s %s, firstLineNumber=%d}", d.Name(), d.Descriptor(), d.FirstLineNumber())
}
