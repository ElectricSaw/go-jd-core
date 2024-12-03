package service

import (
	intcls "bitbucket.org/coontec/go-jd-core/class/interfaces/classpath"
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
)

type Magic uint32

const JavaMagicNumber Magic = 0xCAFEBABE

type IClassFileReader interface {
	Offset() int
	Skip(length int)
	Read() byte
	ReadUnsignedByte() int
	ReadUnsignedShort() int
	ReadMagic() Magic
	ReadInt() int
	ReadFloat() float32
	ReadLong() int64
	ReadDouble() float64
	ReadFully(length int) []byte
	ReadUTF8() string
}

type IClassFileAnnotationDeclaration interface {
	IClassFileTypeDeclaration
	intmod.IAnnotationDeclaration

	FirstLineNumber() int
	String() string
}

type IClassFileBodyDeclaration interface {
	IClassFileMemberDeclaration
	intmod.IBodyDeclaration

	FieldDeclarations() []IClassFileFieldDeclaration
	SetFieldDeclarations(fieldDeclarations []IClassFileFieldDeclaration)
	MethodDeclarations() []IClassFileConstructorOrMethodDeclaration
	SetMethodDeclarations(methodDeclarations []IClassFileConstructorOrMethodDeclaration)
	InnerTypeDeclarations() []IClassFileTypeDeclaration
	SetInnerTypeDeclarations(innerTypeDeclarations []IClassFileTypeDeclaration)
	InnerTypeDeclaration(internalName string) IClassFileTypeDeclaration
	RemoveInnerTypeDeclaration(internalName string) IClassFileTypeDeclaration
	UpdateFirstLineNumber(members []IClassFileMemberDeclaration)
	ClassFile() intcls.IClassFile
	FirstLineNumber() int
	OuterTypeFieldName() string
	SetOuterTypeFieldName(outerTypeFieldName string)
	SyntheticInnerFieldNames() []string
	SetSyntheticInnerFieldNames(names []string)
	OuterBodyDeclaration() IClassFileBodyDeclaration
	SetOuterBodyDeclaration(bodyDeclaration IClassFileBodyDeclaration)
	Bindings() map[string]intmod.ITypeArgument
	TypeBounds() map[string]intmod.IType
	IsClassDeclaration() bool
	String() string
}

type IClassFileClassDeclaration interface {
	IClassFileTypeDeclaration
	intmod.IClassDeclaration

	FirstLineNumber() int
	String() string
}

type IClassFileConstructorDeclaration interface {
	IClassFileConstructorOrMethodDeclaration
	intmod.IConstructorDeclaration

	ClassFile() intcls.IClassFile
	Method() intcls.IMethod
	ParameterTypes() intmod.IType
	ReturnedType() intmod.IType
	BodyDeclaration() IClassFileBodyDeclaration
	Bindings() map[string]intmod.ITypeArgument
	TypeBounds() map[string]intmod.IType
	FirstLineNumber() int
}

type IClassFileConstructorOrMethodDeclaration interface {
	IClassFileMemberDeclaration

	Flags() int
	ClassFile() intcls.IClassFile
	Method() intcls.IMethod
	TypeParameters() intmod.ITypeParameter
	ParameterTypes() intmod.IType
	ReturnedType() intmod.IType
	BodyDeclaration() IClassFileBodyDeclaration
	Bindings() map[string]intmod.ITypeArgument
	TypeBounds() map[string]intmod.IType
	Statements() intmod.IStatement

	SetFlags(flags int)
	SetFormalParameters(formalParameters intmod.IFormalParameter)
	SetStatements(statement intmod.IStatement)
}

type IClassFileEnumDeclaration interface {
	IClassFileTypeDeclaration
	intmod.IEnumDeclaration

	FirstLineNumber() int
	String() string
}

type IClassFileConstant interface {
	intmod.IConstant

	Index() int
	String() string
}

type IClassFileFieldDeclaration interface {
	IClassFileMethodDeclaration
	intmod.IFieldDeclaration

	FirstLineNumber() int
	SetFirstLineNumber(firstLineNumber int)
	String() string
}

type IClassFileFormalParameter interface {
	ILocalVariableReference
	intmod.IFormalParameter

	Type() intmod.IType
	Name() string
	LocalVariable() ILocalVariableReference
	SetLocalVariable(localVariable ILocalVariableReference)
	String() string
}

type IClassFileInterfaceDeclaration interface {
	IClassFileTypeDeclaration
	intmod.IInterfaceDeclaration

	FirstLineNumber() int
	String() string
}

type IClassFileLocalVariableDeclarator interface {
	ILocalVariableReference
	intmod.ILocalVariableDeclarator

	Name() string
	SetName(name string)
	LocalVariable() ILocalVariableReference
	SetLocalVariable(localVariable ILocalVariableReference)
}

type IClassFileMemberDeclaration interface {
	intmod.IMemberDeclaration

	FirstLineNumber() int
}

type IClassFileMethodDeclaration interface {
	IClassFileConstructorOrMethodDeclaration
	intmod.IMethodDeclaration

	ClassFile() intcls.IClassFile
	Method() intcls.IMethod
	ParameterTypes() intmod.IType
	BodyDeclaration() IClassFileBodyDeclaration
	Bindings() map[string]intmod.ITypeArgument
	TypeBounds() map[string]intmod.IType
	FirstLineNumber() int
	String() string
}

type IClassFileStaticInitializerDeclaration interface {
	IClassFileConstructorOrMethodDeclaration
	intmod.IStaticInitializerDeclaration

	Flags() int
	ClassFile() intcls.IClassFile
	Method() intcls.IMethod
	TypeParameters() intmod.ITypeParameter
	ParameterTypes() intmod.IType
	ReturnedType() intmod.IType
	BodyDeclaration() IClassFileBodyDeclaration
	Bindings() map[string]intmod.ITypeArgument
	TypeBounds() map[string]intmod.IType
	SetFlags(flags int)
	SetFormalParameters(formalParameters intmod.IFormalParameter)
	SetFirstLineNumber(lineNumber int)
	String() string
}

type IClassFileTypeDeclaration interface {
	IClassFileMemberDeclaration

	InternalTypeName() string
	BodyDeclaration() intmod.IBodyDeclaration
}
