package declaration

import (
	"fmt"
	intcls "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/classpath"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/declaration"
)

func NewClassFileMethodDeclaration(bodyDeclaration intsrv.IClassFileBodyDeclaration, classFile intcls.IClassFile,
	method intcls.IMethod, name string, returnedType intmod.IType, parameterTypes intmod.IType,
	bindings map[string]intmod.ITypeArgument, typeBounds map[string]intmod.IType) intsrv.IClassFileMethodDeclaration {
	return NewClassFileMethodDeclaration3(bodyDeclaration, classFile, method, nil, name,
		nil, returnedType, parameterTypes, nil,
		nil, bindings, typeBounds, -1)
}

func NewClassFileMethodDeclaration2(bodyDeclaration intsrv.IClassFileBodyDeclaration, classFile intcls.IClassFile,
	method intcls.IMethod, name string, returnedType intmod.IType, parameterTypes intmod.IType,
	bindings map[string]intmod.ITypeArgument, typeBounds map[string]intmod.IType, firstLineNumber int) intsrv.IClassFileMethodDeclaration {
	return NewClassFileMethodDeclaration3(bodyDeclaration, classFile, method, nil, name,
		nil, returnedType, parameterTypes, nil,
		nil, bindings, typeBounds, firstLineNumber)
}

func NewClassFileMethodDeclaration3(bodyDeclaration intsrv.IClassFileBodyDeclaration, classFile intcls.IClassFile,
	method intcls.IMethod, annotationReferences intmod.IAnnotationReference, name string,
	typeParameters intmod.ITypeParameter, returnedType intmod.IType, parameterTypes intmod.IType,
	exceptionTypes intmod.IType, defaultAnnotationValue intmod.IElementValue,
	bindings map[string]intmod.ITypeArgument, typeBounds map[string]intmod.IType, firstLineNumber int) intsrv.IClassFileMethodDeclaration {
	d := &ClassFileMethodDeclaration{
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
	d.SetValue(d)
	return d
}

type ClassFileMethodDeclaration struct {
	declaration.MethodDeclaration

	bodyDeclaration intsrv.IClassFileBodyDeclaration
	classFile       intcls.IClassFile
	method          intcls.IMethod
	parameterTypes  intmod.IType
	bindings        map[string]intmod.ITypeArgument
	typeBounds      map[string]intmod.IType
	firstLineNumber int
}

func (d *ClassFileMethodDeclaration) ClassFile() intcls.IClassFile {
	return d.classFile
}

func (d *ClassFileMethodDeclaration) Method() intcls.IMethod {
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

func (d *ClassFileMethodDeclaration) AcceptDeclaration(visitor intmod.IDeclarationVisitor) {
	visitor.VisitMethodDeclaration(d)
}

func (d *ClassFileMethodDeclaration) String() string {
	return fmt.Sprintf("ClassFileMethodDeclaration{%s %s, firstLineNumber=%d}", d.Name(), d.Descriptor(), d.FirstLineNumber())
}
