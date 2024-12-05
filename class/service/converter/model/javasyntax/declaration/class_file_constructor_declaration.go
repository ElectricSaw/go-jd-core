package declaration

import (
	intcls "bitbucket.org/coontec/go-jd-core/class/interfaces/classpath"
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/declaration"
)

func NewClassFileConstructorDeclaration(
	bodyDeclaration intsrv.IClassFileBodyDeclaration,
	classFile intcls.IClassFile,
	method intcls.IMethod,
	annotationReferences intmod.IAnnotationReference,
	typeParameters intmod.ITypeParameter,
	parameterTypes intmod.IType,
	exceptionTypes intmod.IType,
	bindings map[string]intmod.ITypeArgument,
	typeBounds map[string]intmod.IType,
	firstLineNumber int) intsrv.IClassFileConstructorDeclaration {
	d := &ClassFileConstructorDeclaration{
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
	d.SetValue(d)
	return d
}

type ClassFileConstructorDeclaration struct {
	declaration.ConstructorDeclaration

	bodyDeclaration intsrv.IClassFileBodyDeclaration
	classFile       intcls.IClassFile
	method          intcls.IMethod
	parameterTypes  intmod.IType
	bindings        map[string]intmod.ITypeArgument
	typeBounds      map[string]intmod.IType
	firstLineNumber int
}

func (d *ClassFileConstructorDeclaration) ClassFile() intcls.IClassFile {
	return d.classFile
}

func (d *ClassFileConstructorDeclaration) Method() intcls.IMethod {
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

func (d *ClassFileConstructorDeclaration) SetFirstLineNumber(firstLineNumber int) {
	d.firstLineNumber = firstLineNumber
}
