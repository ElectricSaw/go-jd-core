package declaration

import "bitbucket.org/coontec/javaClass/java/model/javasyntax/statement"

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

func (d *StaticInitializerDeclaration) GetDescription() string {
	return d.descriptor
}

func (d *StaticInitializerDeclaration) GetStatement() statement.Statement {
	return d.statements
}

func (d *StaticInitializerDeclaration) Accept(visitor DeclarationVisitor) {
	visitor.VisitStaticInitializerDeclaration(d)
}

func (d *StaticInitializerDeclaration) String() string {
	return "StaticInitializerDeclaration{}"
}
