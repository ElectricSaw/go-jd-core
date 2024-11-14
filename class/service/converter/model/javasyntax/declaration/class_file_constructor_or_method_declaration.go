package declaration

import (
	"bitbucket.org/coontec/javaClass/class/model/classfile"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/declaration"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/statement"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
)

type ClassFileConstructorOrMethodDeclaration interface {
	ClassFileMemberDeclaration

	Flags() int
	SetFlags(flags int)

	ClassFile() *classfile.ClassFile
	Method() *classfile.Method
	TypeParameters() _type.ITypeParameter
	ParameterTypes() _type.IType
	ReturnedType() _type.IType
	BodyDeclaration() ClassFileBodyDeclaration
	Bindings() map[string]_type.ITypeArgument
	TypeBounds() map[string]_type.IType
	SetFormalParameters(formalParameters declaration.IFormalParameter)
	Statements() statement.Statement
	SetStatements(statement statement.Statement)
}
