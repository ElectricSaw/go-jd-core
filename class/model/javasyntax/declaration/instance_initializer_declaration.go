package declaration

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/statement"
)

func NewInstanceInitializerDeclaration(description string, statements statement.Statement) intsyn.IInstanceInitializerDeclaration {
	return &InstanceInitializerDeclaration{
		description: description,
		statements:  statements,
	}
}

type InstanceInitializerDeclaration struct {
	AbstractMemberDeclaration

	description string
	statements  statement.Statement
}

func (d *InstanceInitializerDeclaration) Description() string {
	return d.description
}

func (d *InstanceInitializerDeclaration) Statements() statement.Statement {
	return d.statements
}

func (d *InstanceInitializerDeclaration) Accept(visitor intsyn.IDeclarationVisitor) {
	visitor.VisitInstanceInitializerDeclaration(d)
}

func (d *InstanceInitializerDeclaration) String() string {
	return "InstanceInitializerDeclaration{}"
}
