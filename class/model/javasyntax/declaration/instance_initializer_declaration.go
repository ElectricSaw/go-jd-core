package declaration

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
)

func NewInstanceInitializerDeclaration(description string, statements intsyn.IStatement) intsyn.IInstanceInitializerDeclaration {
	return &InstanceInitializerDeclaration{
		description: description,
		statements:  statements,
	}
}

type InstanceInitializerDeclaration struct {
	AbstractMemberDeclaration

	description string
	statements  intsyn.IStatement
}

func (d *InstanceInitializerDeclaration) Description() string {
	return d.description
}

func (d *InstanceInitializerDeclaration) Statements() intsyn.IStatement {
	return d.statements
}

func (d *InstanceInitializerDeclaration) Accept(visitor intsyn.IDeclarationVisitor) {
	visitor.VisitInstanceInitializerDeclaration(d)
}

func (d *InstanceInitializerDeclaration) String() string {
	return "InstanceInitializerDeclaration{}"
}
