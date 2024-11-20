package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/classfile"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/declaration"
	"fmt"
)

func NewClassFileMethodDeclaration(bodyDeclaration intsrv.IClassFileBodyDeclaration, classFile *classfile.ClassFile,
	method *classfile.Method, name string, returnedType intmod.IType, parameterTypes intmod.IType,
	bindings map[string]intmod.ITypeArgument, typeBounds map[string]intmod.IType) intsrv.IClassFileMethodDeclaration {
	return &ClassFileMethodDeclaration{
		MethodDeclaration: *declaration.NewMethodDeclaration6(nil, method.AccessFlags(), name, nil, returnedType,
			nil, nil, method.Descriptor(), nil, nil).(*declaration.MethodDeclaration),
		bodyDeclaration: bodyDeclaration,
		classFile:       classFile,
		parameterTypes:  parameterTypes,
		method:          method,
		bindings:        bindings,
		typeBounds:      typeBounds,
	}
}

func NewClassFileMethodDeclaration2(bodyDeclaration intsrv.IClassFileBodyDeclaration, classFile *classfile.ClassFile,
	method *classfile.Method, name string, returnedType intmod.IType, parameterTypes intmod.IType,
	bindings map[string]intmod.ITypeArgument, typeBounds map[string]intmod.IType, firstLineNumber int) intsrv.IClassFileMethodDeclaration {
	return &ClassFileMethodDeclaration{
		MethodDeclaration: *declaration.NewMethodDeclaration6(nil, method.AccessFlags(), name, nil, returnedType,
			nil, nil, method.Descriptor(), nil, nil).(*declaration.MethodDeclaration),
		bodyDeclaration: bodyDeclaration,
		classFile:       classFile,
		parameterTypes:  parameterTypes,
		method:          method,
		bindings:        bindings,
		typeBounds:      typeBounds,
		firstLineNumber: firstLineNumber,
	}
}

func NewClassFileMethodDeclaration3(bodyDeclaration intsrv.IClassFileBodyDeclaration, classFile *classfile.ClassFile,
	method *classfile.Method, annotationReferences intmod.IAnnotationReference, name string,
	typeParameters intmod.ITypeParameter, returnedType intmod.IType, parameterTypes intmod.IType,
	exceptionTypes intmod.IType, defaultAnnotationValue intmod.IElementValue,
	bindings map[string]intmod.ITypeArgument, typeBounds map[string]intmod.IType, firstLineNumber int) intsrv.IClassFileMethodDeclaration {
	return &ClassFileMethodDeclaration{
		MethodDeclaration: *declaration.NewMethodDeclaration6(annotationReferences, method.AccessFlags(), name, typeParameters,
			returnedType, nil, exceptionTypes, method.Descriptor(), nil, defaultAnnotationValue).(*declaration.MethodDeclaration),
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

	bodyDeclaration intsrv.IClassFileBodyDeclaration
	classFile       *classfile.ClassFile
	method          *classfile.Method
	parameterTypes  intmod.IType
	bindings        map[string]intmod.ITypeArgument
	typeBounds      map[string]intmod.IType
	firstLineNumber int
}

func (d *ClassFileMethodDeclaration) ClassFile() *classfile.ClassFile {
	return d.classFile
}

func (d *ClassFileMethodDeclaration) Method() *classfile.Method {
	return d.method
}

func (d *ClassFileMethodDeclaration) ParameterTypes() intmod.IType {
	return d.parameterTypes
}

func (d *ClassFileMethodDeclaration) BodyDeclaration() intsrv.IClassFileBodyDeclaration {
	return d.bodyDeclaration
}

func (d *ClassFileMethodDeclaration) Bindings() map[string]intmod.ITypeArgument {
	return d.bindings
}

func (d *ClassFileMethodDeclaration) TypeBounds() map[string]intmod.IType {
	return d.typeBounds
}

func (d *ClassFileMethodDeclaration) FirstLineNumber() int {
	return d.firstLineNumber
}

func (d *ClassFileMethodDeclaration) String() string {
	return fmt.Sprintf("ClassFileMethodDeclaration{%s %s, firstLineNumber=%d}", d.Name(), d.Descriptor(), d.FirstLineNumber())
}
