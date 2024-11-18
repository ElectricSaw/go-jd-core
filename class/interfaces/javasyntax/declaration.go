package javasyntax

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/expression"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/reference"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/statement"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"bitbucket.org/coontec/javaClass/class/util"
)

const (
	// Access flags for Class, Field, Method, Nested class, Module, Module Requires, Module Exports, Module Opens
	FlagPublic       = 0x0001 // C  F  M  N  .  .  .  .
	FlagPrivate      = 0x0002 // .  F  M  N  .  .  .  .
	FlagProtected    = 0x0004 // .  F  M  N  .  .  .  .
	FlagStatic       = 0x0008 // C  F  M  N  .  .  .  .
	FlagFinal        = 0x0010 // C  F  M  N  .  .  .  .
	FlagSynchronized = 0x0020 // .  .  M  .  .  .  .  .
	FlagSuper        = 0x0020 // C  .  .  .  .  .  .  .
	FlagOpen         = 0x0020 // .  .  .  .  Mo .  .  .
	FlagTransitive   = 0x0020 // .  .  .  .  .  MR .  .
	FlagVolatile     = 0x0040 // .  F  .  .  .  .  .  .
	FlagBridge       = 0x0040 // .  .  M  .  .  .  .  .
	FlagStaticPhase  = 0x0040 // .  .  .  .  .  MR .  .
	FlagTransient    = 0x0080 // .  F  .  .  .  .  .  .
	FlagVarArgs      = 0x0080 // .  .  M  .  .  .  .  .
	FlagNative       = 0x0100 // .  .  M  .  .  .  .  .
	FlagInterface    = 0x0200 // C  .  .  N  .  .  .  .
	FlagAnonymous    = 0x0200 // .  .  M  .  .  .  .  . // Custom flag
	FlagAbstract     = 0x0400 // C  .  M  N  .  .  .  .
	FlagStrict       = 0x0800 // .  .  M  .  .  .  .  .
	FlagSynthetic    = 0x1000 // C  F  M  N  Mo MR ME MO
	FlagAnnotation   = 0x2000 // C  .  .  N  .  .  .  .
	FlagEnum         = 0x4000 // C  F  .  N  .  .  .  .
	FlagModule       = 0x8000 // C  .  .  .  .  .  .  .
	FlagMandated     = 0x8000 // .  .  .  .  Mo MR ME MO

	// Extension
	FlagDefault = 0x10000 // .  .  M  .  .  .  .  .
)

type IAnnotationDeclaration interface {
	AnnotationDeclarators() IFieldDeclarator
	Accept(visitor IDeclarationVisitor)
	String() string
}

type IArrayVariableInitializer interface {
	Type() _type.IType
	LineNumber() int
	Accept(visitor IDeclarationVisitor)
}

type IBodyDeclaration interface {
	IDeclaration
	InternalTypeName() string
	MemberDeclarations() IMemberDeclaration
}

type IClassDeclaration interface {
	SuperType() _type.IObjectType
	IsClassDeclaration() bool
	Accept(visitor IDeclarationVisitor)
	String() string
}

type IConstructorDeclaration interface {
	Flags() int
	SetFlags(flags int)
	IsStatic() bool
	AnnotationReferences() reference.IReference
	TypeParameters() _type.ITypeParameter
	FormalParameters() IFormalParameter
	SetFormalParameters(formalParameter IFormalParameter)
	ExceptionTypes() _type.IType
	Descriptor() string
	Statements() statement.Statement
	SetStatements(state statement.Statement)
	Accept(visitor IDeclarationVisitor)
	String() string
}

type IDeclaration interface {
	Accept(visitor IDeclarationVisitor)
}

type IDeclarationVisitor interface {
	VisitAnnotationDeclaration(declaration IAnnotationDeclaration)
	VisitArrayVariableInitializer(declaration IArrayVariableInitializer)
	VisitBodyDeclaration(declaration IBodyDeclaration)
	VisitClassDeclaration(declaration IClassDeclaration)
	VisitConstructorDeclaration(declaration IConstructorDeclaration)
	VisitEnumDeclaration(declaration IEnumDeclaration)
	VisitEnumDeclarationConstant(declaration IConstant)
	VisitExpressionVariableInitializer(declaration IExpressionVariableInitializer)
	VisitFieldDeclaration(declaration IFieldDeclaration)
	VisitFieldDeclarator(declaration IFieldDeclarator)
	VisitFieldDeclarators(declarations IFieldDeclarators)
	VisitFormalParameter(declaration IFormalParameter)
	VisitFormalParameters(declarations IFormalParameters)
	VisitInstanceInitializerDeclaration(declaration IInstanceInitializerDeclaration)
	VisitInterfaceDeclaration(declaration IInterfaceDeclaration)
	VisitLocalVariableDeclaration(declaration ILocalVariableDeclaration)
	VisitLocalVariableDeclarator(declarator ILocalVariableDeclarator)
	VisitLocalVariableDeclarators(declarators ILocalVariableDeclarators)
	VisitMethodDeclaration(declaration IMethodDeclaration)
	VisitMemberDeclarations(declarations IMemberDeclarations)
	VisitModuleDeclaration(declarations IModuleDeclaration)
	VisitStaticInitializerDeclaration(declaration IStaticInitializerDeclaration)
	VisitTypeDeclarations(declarations ITypeDeclarations)
}

type IFieldDeclarator interface {
	IDeclaration
	util.Base[IFieldDeclarator]

	SetFieldDeclaration(fieldDeclaration IFieldDeclaration)
	FieldDeclaration() IFieldDeclaration
	Name() string
	Dimension() int
	VariableInitializer() IVariableInitializer
	SetVariableInitializer(variableInitializer IVariableInitializer)
	Accept(visitor IDeclarationVisitor)
	String() string
}

type IFormalParameter interface {
	IDeclaration

	AnnotationReferences() reference.IAnnotationReference
	IsFinal() bool
	SetFinal(final bool)
	Type() _type.IType
	IsVarargs() bool
	Name() string
	SetName(name string)
	Accept(visitor IDeclarationVisitor)
	String() string
}

type ILocalVariableDeclarator interface {
	IDeclaration

	Name() string
	SetName(name string)
	Dimension() int
	SetDimension(dimension int)
	LineNumber() int
	VariableInitializer() IVariableInitializer
	Accept(visitor IDeclarationVisitor)
	String() string
}

type IMemberDeclaration interface {
	IDeclaration

	IsClassDeclaration() bool
}

type ITypeDeclaration interface {
	IMemberDeclaration

	AnnotationReferences() reference.IAnnotationReference
	Flags() int
	SetFlags(flags int)
	InternalTypeName() string
	Name() string
	BodyDeclaration() IBodyDeclaration
}

type IVariableInitializer interface {
	IDeclaration

	LineNumber() int
	IsExpressionVariableInitializer() bool
	Expression() expression.Expression
}

type IEnumDeclaration interface {
	Interfaces() _type.IType
	Constants() []IConstant
	SetConstants(constants []IConstant)
	Accept(visitor IDeclarationVisitor)
	String() string
}

type IConstant interface {
	LineNumber() int
	AnnotationReferences() reference.IAnnotationReference
	Name() string
	Arguments() expression.Expression
	SetArguments(arguments expression.Expression)
	BodyDeclaration() IBodyDeclaration
	Accept(visitor IDeclarationVisitor)
}

type IExpressionVariableInitializer interface {
	Expression() expression.Expression
	LineNumber() int
	SetExpression(expression expression.Expression)
	IsExpressionVariableInitializer() bool
	Accept(visitor IDeclarationVisitor)
}

type IFieldDeclaration interface {
	Flags() int
	SetFlags(flags int)
	AnnotationReferences() reference.IAnnotationReference
	Type() _type.IType
	SetType(t _type.IType)
	FieldDeclarators() IFieldDeclarator
	SetFieldDeclarators(fd IFieldDeclarator)
	Accept(visitor IDeclarationVisitor)
}
type IFieldDeclarators interface {
	SetFieldDeclaration(fieldDeclaration IFieldDeclaration)
	Accept(visitor IDeclarationVisitor)
}
type IFormalParameters interface {
	Accept(visitor IDeclarationVisitor)
}

type IInstanceInitializerDeclaration interface {
	Description() string
	Statements() statement.Statement
	Accept(visitor IDeclarationVisitor)
	String() string
}

type IInterfaceDeclaration interface {
	TypeParameters() _type.ITypeParameter
	Interfaces() _type.IType
	Accept(visitor IDeclarationVisitor)
	String() string
}

type ILocalVariableDeclaration interface {
	IsFinal() bool
	SetFinal(final bool)
	Type() _type.IType
	LocalVariableDeclarators() ILocalVariableDeclarator
	SetLocalVariableDeclarators(localVariableDeclarators ILocalVariableDeclarator)
	Accept(visitor IDeclarationVisitor)
}
type ILocalVariableDeclarators interface {
	LineNumber() int
	VariableInitializer() IVariableInitializer
	Accept(visitor IDeclarationVisitor)
	String() string
}
type IMemberDeclarations interface {
	Accept(visitor IDeclarationVisitor)
}

type IMethodDeclaration interface {
	Flags() int
	SetFlags(flags int)
	AnnotationReferences() reference.IAnnotationReference
	IsStatic() bool
	Name() string
	TypeParameters() _type.ITypeParameter
	ReturnedType() _type.IType
	FormalParameter() IFormalParameter
	SetFormalParameters(formalParameter IFormalParameter)
	ExceptionTypes() _type.IType
	Descriptor() string
	Statements() statement.Statement
	SetStatements(statements statement.Statement)
	DefaultAnnotationValue() reference.IElementValue
	Accept(visitor IDeclarationVisitor)
	String() string
}

type IModuleDeclaration interface {
	Version() string
	Requires() []IModuleInfo
	Exports() []IPackageInfo
	Opens() []IPackageInfo
	Uses() []string
	Provides() []IServiceInfo
	Accept(visitor IDeclarationVisitor)
	String() string
}

type IModuleInfo interface {
	Name() string
	Flags() int
	Version() string
	String() string
}

type IPackageInfo interface {
	InternalName() string
	Flags() int
	ModuleInfoNames() []string
	String() string
}

type IServiceInfo interface {
	InternalTypeName() string
	ImplementationTypeNames() []string
	String() string
}

type IStaticInitializerDeclaration interface {
	Description() string
	Statements() statement.Statement
	SetStatements(statements statement.Statement)
	Accept(visitor IDeclarationVisitor)
	String() string
}
type ITypeDeclarations interface {
	Accept(visitor IDeclarationVisitor)
}
