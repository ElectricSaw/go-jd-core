package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/classfile"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/declaration"
	"fmt"
)

func NewClassFileStaticInitializerDeclaration(bodyDeclaration intsrv.IClassFileBodyDeclaration,
	classFile *classfile.ClassFile, method *classfile.Method,
	bindings map[string]intmod.ITypeArgument, typeBounds map[string]intmod.IType,
	firstLineNumber int) intsrv.IClassFileStaticInitializerDeclaration {
	return &ClassFileStaticInitializerDeclaration{
		StaticInitializerDeclaration: *declaration.NewStaticInitializerDeclaration(method.Descriptor(), nil).(*declaration.StaticInitializerDeclaration),
		bodyDeclaration:              bodyDeclaration,
		classFile:                    classFile,
		method:                       method,
		bindings:                     bindings,
		typeBounds:                   typeBounds,
		firstLineNumber:              firstLineNumber,
	}
}

func NewClassFileStaticInitializerDeclaration2(bodyDeclaration intsrv.IClassFileBodyDeclaration, classFile *classfile.ClassFile,
	method *classfile.Method, bindings map[string]intmod.ITypeArgument, typeBounds map[string]intmod.IType,
	firstLineNumber int, statements intmod.IStatement) intsrv.IClassFileStaticInitializerDeclaration {
	return &ClassFileStaticInitializerDeclaration{
		StaticInitializerDeclaration: *declaration.NewStaticInitializerDeclaration(method.Descriptor(), statements).(*declaration.StaticInitializerDeclaration),
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

	bodyDeclaration intsrv.IClassFileBodyDeclaration
	classFile       *classfile.ClassFile
	method          *classfile.Method
	bindings        map[string]intmod.ITypeArgument
	typeBounds      map[string]intmod.IType
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

func (d *ClassFileStaticInitializerDeclaration) String() string {
	return fmt.Sprintf("ClassFileStaticInitializerDeclaration{%s, firstLineNumber=%d}", d.Description(), d.firstLineNumber)
}
