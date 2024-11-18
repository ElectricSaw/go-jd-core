package declaration

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
)

func NewStaticInitializerDeclaration(descriptor string, statements intsyn.IStatement) intsyn.IStaticInitializerDeclaration {
	return &StaticInitializerDeclaration{
		descriptor: descriptor,
		statements: statements,
	}
}

type StaticInitializerDeclaration struct {
	AbstractMemberDeclaration

	descriptor string
	statements intsyn.IStatement
}

func (d *StaticInitializerDeclaration) Description() string {
	return d.descriptor
}

func (d *StaticInitializerDeclaration) Statements() intsyn.IStatement {
	return d.statements
}

func (d *StaticInitializerDeclaration) SetStatements(statements intsyn.IStatement) {
	d.statements = statements
}

func (d *StaticInitializerDeclaration) Accept(visitor intsyn.IDeclarationVisitor) {
	visitor.VisitStaticInitializerDeclaration(d)
}

func (d *StaticInitializerDeclaration) String() string {
	return "StaticInitializerDeclaration{}"
}
