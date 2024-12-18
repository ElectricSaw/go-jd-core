package declaration

import (
	"fmt"
	intcls "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/classpath"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/declaration"
)

func NewClassFileStaticInitializerDeclaration(bodyDeclaration intsrv.IClassFileBodyDeclaration,
	classFile intcls.IClassFile, method intcls.IMethod,
	bindings map[string]intmod.ITypeArgument, typeBounds map[string]intmod.IType,
	firstLineNumber int) intsrv.IClassFileStaticInitializerDeclaration {
	return NewClassFileStaticInitializerDeclaration2(bodyDeclaration, classFile, method,
		bindings, typeBounds, firstLineNumber, nil)
}

func NewClassFileStaticInitializerDeclaration2(bodyDeclaration intsrv.IClassFileBodyDeclaration, classFile intcls.IClassFile,
	method intcls.IMethod, bindings map[string]intmod.ITypeArgument, typeBounds map[string]intmod.IType,
	firstLineNumber int, statements intmod.IStatement) intsrv.IClassFileStaticInitializerDeclaration {
	d := &ClassFileStaticInitializerDeclaration{
		StaticInitializerDeclaration: *declaration.NewStaticInitializerDeclaration(method.Descriptor(), statements).(*declaration.StaticInitializerDeclaration),
		bodyDeclaration:              bodyDeclaration,
		classFile:                    classFile,
		method:                       method,
		bindings:                     bindings,
		typeBounds:                   typeBounds,
		firstLineNumber:              firstLineNumber,
	}
	d.SetValue(d)
	return d
}

type ClassFileStaticInitializerDeclaration struct {
	declaration.StaticInitializerDeclaration

	bodyDeclaration intsrv.IClassFileBodyDeclaration
	classFile       intcls.IClassFile
	method          intcls.IMethod
	bindings        map[string]intmod.ITypeArgument
	typeBounds      map[string]intmod.IType
	firstLineNumber int
}

func (d *ClassFileStaticInitializerDeclaration) Flags() int {
	return 0
}

func (d *ClassFileStaticInitializerDeclaration) ClassFile() intcls.IClassFile {
	return d.classFile
}

func (d *ClassFileStaticInitializerDeclaration) Method() intcls.IMethod {
	return d.method
}

func (d *ClassFileStaticInitializerDeclaration) TypeParameters() intmod.ITypeParameter {
	return nil
}

func (d *ClassFileStaticInitializerDeclaration) ParameterTypes() intmod.IType {
	return nil
}

func (d *ClassFileStaticInitializerDeclaration) ReturnedType() intmod.IType {
	return nil
}

func (d *ClassFileStaticInitializerDeclaration) BodyDeclaration() intsrv.IClassFileBodyDeclaration {
	return d.bodyDeclaration
}

func (d *ClassFileStaticInitializerDeclaration) Bindings() map[string]intmod.ITypeArgument {
	return d.bindings
}

func (d *ClassFileStaticInitializerDeclaration) TypeBounds() map[string]intmod.IType {
	return d.typeBounds
}

func (d *ClassFileStaticInitializerDeclaration) SetFlags(flags int) {
}

func (d *ClassFileStaticInitializerDeclaration) SetFormalParameters(formalParameters intmod.IFormalParameter) {
}

func (d *ClassFileStaticInitializerDeclaration) FirstLineNumber() int {
	return d.firstLineNumber
}

func (d *ClassFileStaticInitializerDeclaration) SetFirstLineNumber(lineNumber int) {
	d.firstLineNumber = lineNumber
}

func (d *ClassFileStaticInitializerDeclaration) AcceptDeclaration(visitor intmod.IDeclarationVisitor) {
	visitor.VisitStaticInitializerDeclaration(d)
}

func (d *ClassFileStaticInitializerDeclaration) String() string {
	return fmt.Sprintf("ClassFileStaticInitializerDeclaration{%s, firstLineNumber=%d}", d.Description(), d.firstLineNumber)
}
