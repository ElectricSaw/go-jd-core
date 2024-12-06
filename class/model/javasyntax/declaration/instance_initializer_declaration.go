package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

func NewInstanceInitializerDeclaration(description string, statements intmod.IStatement) intmod.IInstanceInitializerDeclaration {
	d := &InstanceInitializerDeclaration{
		description: description,
		statements:  statements,
	}
	d.SetValue(d)
	return d
}

type InstanceInitializerDeclaration struct {
	AbstractMemberDeclaration
	util.DefaultBase[intmod.IMemberDeclaration]

	description string
	statements  intmod.IStatement
}

func (d *InstanceInitializerDeclaration) Description() string {
	return d.description
}

func (d *InstanceInitializerDeclaration) Statements() intmod.IStatement {
	return d.statements
}

func (d *InstanceInitializerDeclaration) AcceptDeclaration(visitor intmod.IDeclarationVisitor) {
	visitor.VisitInstanceInitializerDeclaration(d)
}

func (d *InstanceInitializerDeclaration) String() string {
	return "InstanceInitializerDeclaration{}"
}
