package declaration

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewInstanceInitializerDeclaration(description string, statements Statement) InstanceInitializerDeclaration {
	d := &InstanceInitializerDeclaration{
		description: description,
		statements:  statements,
	}
	d.SetValue(d)
	return d
}

type InstanceInitializerDeclaration struct {
	AbstractMemberDeclaration
	util.DefaultBase[MemberDeclaration]

	description string
	statements  Statement
}

func (d *InstanceInitializerDeclaration) Description() string {
	return d.description
}

func (d *InstanceInitializerDeclaration) Statements() Statement {
	return d.statements
}

func (d *InstanceInitializerDeclaration) AcceptDeclaration(visitor DeclarationVisitor) {
	visitor.VisitInstanceInitializerDeclaration(d)
}

func (d *InstanceInitializerDeclaration) String() string {
	return "InstanceInitializerDeclaration{}"
}
