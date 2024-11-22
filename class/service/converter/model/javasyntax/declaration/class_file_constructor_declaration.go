package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/declaration"
)

func NewClassFileConstructorDeclaration(
	bodyDeclaration intsrv.IClassFileBodyDeclaration,
	classFile intmod.IClassFile,
	method intmod.IMethod,
	annotationReferences intmod.IAnnotationReference,
	typeParameters intmod.ITypeParameter,
	parameterTypes intmod.IType,
	exceptionTypes intmod.IType,
	bindings map[string]intmod.ITypeArgument,
	typeBounds map[string]intmod.IType,
	firstLineNumber int) intsrv.IClassFileConstructorDeclaration {
	return &ClassFileConstructorDeclaration{
		ConstructorDeclaration: *declaration.NewConstructorDeclarationWithAll(
			annotationReferences, method.AccessFlags(), typeParameters,
			nil, exceptionTypes, method.Descriptor(), nil).(*declaration.ConstructorDeclaration),
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

	bodyDeclaration intsrv.IClassFileBodyDeclaration
	classFile       intmod.IClassFile
	method          intmod.IMethod
	parameterTypes  intmod.IType
	bindings        map[string]intmod.ITypeArgument
	typeBounds      map[string]intmod.IType
	firstLineNumber int
}

func (d *ClassFileConstructorDeclaration) ClassFile() intmod.IClassFile {
	return d.classFile
}

func (d *ClassFileConstructorDeclaration) Method() intmod.IMethod {
	return d.method
}

func (d *ClassFileConstructorDeclaration) ParameterTypes() intmod.IType {
	return d.parameterTypes
}

func (d *ClassFileConstructorDeclaration) ReturnedType() intmod.IType {
	return nil
}

func (d *ClassFileConstructorDeclaration) BodyDeclaration() intsrv.IClassFileBodyDeclaration {
	return d.bodyDeclaration
}

func (d *ClassFileConstructorDeclaration) Bindings() map[string]intmod.ITypeArgument {
	return d.bindings
}

func (d *ClassFileConstructorDeclaration) TypeBounds() map[string]intmod.IType {
	return d.typeBounds
}

func (d *ClassFileConstructorDeclaration) FirstLineNumber() int {
	return d.firstLineNumber
}
