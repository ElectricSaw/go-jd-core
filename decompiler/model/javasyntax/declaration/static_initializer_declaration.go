package declaration

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewStaticInitializerDeclaration(descriptor string, statements Statement) StaticInitializerDeclaration {
	d := &StaticInitializerDeclaration{
		descriptor: descriptor,
		statements: statements,
	}
	d.SetValue(d)
	return d
}

type StaticInitializerDeclaration struct {
	AbstractMemberDeclaration

	descriptor string
	statements Statement
}

func (d *StaticInitializerDeclaration) Description() string {
	return d.descriptor
}

func (d *StaticInitializerDeclaration) Statements() Statement {
	return d.statements
}

func (d *StaticInitializerDeclaration) SetStatements(statements Statement) {
	d.statements = statements
}

func (d *StaticInitializerDeclaration) AcceptDeclaration(visitor DeclarationVisitor) {
	visitor.VisitStaticInitializerDeclaration(d)
}

func (d *StaticInitializerDeclaration) String() string {
	return "StaticInitializerDeclaration{}"
}
