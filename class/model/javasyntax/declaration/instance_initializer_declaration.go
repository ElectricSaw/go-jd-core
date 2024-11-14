package declaration

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/statement"
)

func NewInstanceInitializerDeclaration(description string, statements statement.Statement) *InstanceInitializerDeclaration {
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

func (d *InstanceInitializerDeclaration) Accept(visitor DeclarationVisitor) {
	visitor.VisitInstanceInitializerDeclaration(d)
}

func (d *InstanceInitializerDeclaration) String() string {
	return "InstanceInitializerDeclaration{}"
}
