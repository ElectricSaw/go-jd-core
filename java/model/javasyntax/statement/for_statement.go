package statement

import (
	"bitbucket.org/coontec/javaClass/java/model/javasyntax/declaration"
	"bitbucket.org/coontec/javaClass/java/model/javasyntax/expression"
	"fmt"
)

func NewForStatementWithDeclaration(declaration declaration.LocalVariableDeclaration, condition expression.Expression, update expression.Expression, statements Statement) *ForStatement {
	return &ForStatement{
		declaration: declaration,
		condition:   condition,
		update:      update,
		statements:  statements,
	}
}

func NewForStatementWithInit(init expression.Expression, condition expression.Expression, update expression.Expression, statements Statement) *ForStatement {
	return &ForStatement{
		init:       init,
		condition:  condition,
		update:     update,
		statements: statements,
	}
}

type ForStatement struct {
	AbstractStatement

	declaration declaration.LocalVariableDeclaration
	init        expression.Expression
	condition   expression.Expression
	update      expression.Expression
	statements  Statement
}

func (s *ForStatement) GetDeclaration() declaration.LocalVariableDeclaration {
	return s.declaration
}

func (s *ForStatement) SetDeclaration(declaration declaration.LocalVariableDeclaration) {
	s.declaration = declaration
}

func (s *ForStatement) GetInit() expression.Expression {
	return s.init
}

func (s *ForStatement) SetInit(init expression.Expression) {
	s.init = init
}

func (s *ForStatement) GetCondition() expression.Expression {
	return s.condition
}

func (s *ForStatement) SetCondition(condition expression.Expression) {
	s.condition = condition
}

func (s *ForStatement) GetUpdate() expression.Expression {
	return s.update
}

func (s *ForStatement) SetUpdate(update expression.Expression) {
	s.update = update
}

func (s *ForStatement) GetStatements() Statement {
	return s.statements
}

func (s *ForStatement) Accept(visitor StatementVisitor) {
	visitor.VisitForStatement(s)
}

func (s *ForStatement) String() string {
	return fmt.Sprintf("ForStatement{%s or %s; %s; %s", s.declaration, s.init, s.condition, s.update)
}
