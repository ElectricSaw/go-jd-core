package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
)

func NewStaticInitializerDeclaration(descriptor string, statements intmod.IStatement) intmod.IStaticInitializerDeclaration {
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
	statements intmod.IStatement
}

func (d *StaticInitializerDeclaration) Description() string {
	return d.descriptor
}

func (d *StaticInitializerDeclaration) Statements() intmod.IStatement {
	return d.statements
}

func (d *StaticInitializerDeclaration) SetStatements(statements intmod.IStatement) {
	d.statements = statements
}

func (d *StaticInitializerDeclaration) AcceptDeclaration(visitor intmod.IDeclarationVisitor) {
	visitor.VisitStaticInitializerDeclaration(d)
}

func (d *StaticInitializerDeclaration) String() string {
	return "StaticInitializerDeclaration{}"
}
