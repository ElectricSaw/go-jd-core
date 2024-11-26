package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
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
	util.DefaultBase[intmod.IStaticInitializerDeclaration]

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

func (d *StaticInitializerDeclaration) Accept(visitor intmod.IDeclarationVisitor) {
	visitor.VisitStaticInitializerDeclaration(d)
}

func (d *StaticInitializerDeclaration) String() string {
	return "StaticInitializerDeclaration{}"
}
