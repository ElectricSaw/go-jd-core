package declaration

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/statement"
)

func NewStaticInitializerDeclaration(descriptor string, statements statement.Statement) intsyn.IStaticInitializerDeclaration {
	return &StaticInitializerDeclaration{
		descriptor: descriptor,
		statements: statements,
	}
}

type StaticInitializerDeclaration struct {
	AbstractMemberDeclaration

	descriptor string
	statements statement.Statement
}

func (d *StaticInitializerDeclaration) Description() string {
	return d.descriptor
}

func (d *StaticInitializerDeclaration) Statements() statement.Statement {
	return d.statements
}

func (d *StaticInitializerDeclaration) SetStatements(statements statement.Statement) {
	d.statements = statements
}

func (d *StaticInitializerDeclaration) Accept(visitor intsyn.IDeclarationVisitor) {
	visitor.VisitStaticInitializerDeclaration(d)
}

func (d *StaticInitializerDeclaration) String() string {
	return "StaticInitializerDeclaration{}"
}
