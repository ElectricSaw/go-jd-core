package declaration

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/expression"
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

type Declaration interface {
	Accept(visitor DeclarationVisitor)
}

type DeclarationVisitor interface {
	VisitAnnotationDeclaration(declaration *AnnotationDeclaration)
	VisitArrayVariableInitializer(declaration *ArrayVariableInitializer)
	VisitBodyDeclaration(declaration *BodyDeclaration)
	VisitClassDeclaration(declaration *ClassDeclaration)
	VisitConstructorDeclaration(declaration *ConstructorDeclaration)
	VisitEnumDeclaration(declaration *EnumDeclaration)
	VisitEnumDeclarationConstant(declaration *Constant)
	VisitExpressionVariableInitializer(declaration *ExpressionVariableInitializer)
	VisitFieldDeclaration(declaration *FieldDeclaration)
	VisitFieldDeclarator(declaration *FieldDeclarator)
	VisitFieldDeclarators(declarations *FieldDeclarators)
	VisitFormalParameter(declaration *FormalParameter)
	VisitFormalParameters(declarations *FormalParameters)
	VisitInstanceInitializerDeclaration(declaration *InstanceInitializerDeclaration)
	VisitInterfaceDeclaration(declaration *InterfaceDeclaration)
	VisitLocalVariableDeclaration(declaration *LocalVariableDeclaration)
	VisitLocalVariableDeclarator(declarator *LocalVariableDeclarator)
	VisitLocalVariableDeclarators(declarators *LocalVariableDeclarators)
	VisitMethodDeclaration(declaration *MethodDeclaration)
	VisitMemberDeclarations(declarations *MemberDeclarations)
	VisitModuleDeclaration(declarations *ModuleDeclaration)
	VisitStaticInitializerDeclaration(declaration *StaticInitializerDeclaration)
	VisitTypeDeclarations(declarations *TypeDeclarations)
}

type IFieldDeclarator interface {
	Declaration

	SetFieldDeclaration(fieldDeclaration *FieldDeclaration)
}

type IFormalParameter interface {
	Declaration
}

type ILocalVariableDeclarator interface {
	Declaration

	LineNumber() int
}

type IMemberDeclaration interface {
	Declaration

	IsClassDeclaration() bool
}

type AbstractMemberDeclaration struct {
}

func (d *AbstractMemberDeclaration) IsClassDeclaration() bool {
	return false
}

func (d *AbstractMemberDeclaration) Accept(visitor DeclarationVisitor) {

}

type ITypeDeclaration interface {
	IMemberDeclaration
}

type AbstractTypeDeclaration struct {
	AbstractMemberDeclaration
}

type VariableInitializer interface {
	Declaration

	LineNumber() int
	IsExpressionVariableInitializer() bool
	Expression() expression.Expression
}

type AbstractVariableInitializer struct {
}

func (d *AbstractVariableInitializer) LineNumber() int {
	return -1
}

func (d *AbstractVariableInitializer) IsExpressionVariableInitializer() bool {
	return false
}

func (d *AbstractVariableInitializer) Expression() expression.Expression {
	return expression.NeNoExpression
}

func (d *AbstractVariableInitializer) Accept(visitor DeclarationVisitor) {

}
