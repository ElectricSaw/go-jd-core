package declaration

import "bitbucket.org/coontec/javaClass/class/model/javasyntax/statement"

func NewStaticInitializerDeclaration(descriptor string, statements statement.Statement) *StaticInitializerDeclaration {
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

func (d *StaticInitializerDeclaration) Accept(visitor DeclarationVisitor) {
	visitor.VisitStaticInitializerDeclaration(d)
}

func (d *StaticInitializerDeclaration) String() string {
	return "StaticInitializerDeclaration{}"
}
