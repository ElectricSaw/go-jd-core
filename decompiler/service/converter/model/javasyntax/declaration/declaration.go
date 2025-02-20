package declaration

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/classfile"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
)

/////////////////////////////////////////////////////////////////////////
//  Global Variable
/////////////////////////////////////////////////////////////////////////

type Magic uint32

const JavaMagicNumber Magic = 0xCAFEBABE

/////////////////////////////////////////////////////////////////////////
//  New Functions
/////////////////////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////////////////
//  Interfaces
/////////////////////////////////////////////////////////////////////////

type IClassFileConstructorOrMethodDeclaration interface {
	GetAnnotationReferences() *model.AnnotationReference
	GetFlags() int
	GetTypeParameters() model.ITypeParameter
	GetFormalParameters() *model.FormalParameter
	GetExceptionTypes() model.IType
	GetDescriptor() string
	GetStatements() model.IStatement
	GetBodyDeclaration() *ClassFileBodyDeclaration
	GetClassFile() *classfile.ClassFile
	GetMethod() *classfile.Method
	GetParameterTypes() model.IType
	GetBindings() map[string]model.ITypeArgument
	GetTypeBounds() map[string]model.IType
	GetReturnedType() model.IType
	GetFirstLineNumber() int

	SetFlags(flags int)
	SetFormalParameters(formalParameters *model.FormalParameter)
	SetStatements(statement model.IStatement)

	IsClassDeclaration() bool
	AcceptDeclaration(visitor model.IDeclarationVisitor)
	String() string
}

type IClassFileMemberDeclaration interface {
	GetFirstLineNumber() int
	IsClassDeclaration() bool
	AcceptDeclaration(visitor model.IDeclarationVisitor)
	String() string
}

type IClassFileTypeDeclaration interface {
	GetFirstLineNumber() int
	GetInternalTypeName() string
	GetBodyDeclaration() model.IBodyDeclaration
	IsClassDeclaration() bool
	AcceptDeclaration(visitor model.IDeclarationVisitor)
	String() string
}

/////////////////////////////////////////////////////////////////////////
//  Structures
/////////////////////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////////////////
//  Additional Structures
/////////////////////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////////////////
//  Functions
/////////////////////////////////////////////////////////////////////////
