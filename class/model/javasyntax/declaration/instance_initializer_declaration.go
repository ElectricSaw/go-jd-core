package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
)

func NewInstanceInitializerDeclaration(description string, statements intmod.IStatement) intmod.IInstanceInitializerDeclaration {
	return &InstanceInitializerDeclaration{
		description: description,
		statements:  statements,
	}
}

type InstanceInitializerDeclaration struct {
	AbstractMemberDeclaration

	description string
	statements  intmod.IStatement
}

func (d *InstanceInitializerDeclaration) Description() string {
	return d.description
}

func (d *InstanceInitializerDeclaration) Statements() intmod.IStatement {
	return d.statements
}

func (d *InstanceInitializerDeclaration) Accept(visitor intmod.IDeclarationVisitor) {
	visitor.VisitInstanceInitializerDeclaration(d)
}

func (d *InstanceInitializerDeclaration) String() string {
	return "InstanceInitializerDeclaration{}"
}
