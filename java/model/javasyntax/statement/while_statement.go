package statement

import "bitbucket.org/coontec/javaClass/java/model/javasyntax/expression"

func NewWhileStatement(condition expression.Expression, statements Statement) *WhileStatement {
	return &WhileStatement{
		condition:  condition,
		statements: statements,
	}
}

type WhileStatement struct {
	AbstractStatement

	condition  expression.Expression
	statements Statement
}

func (s *WhileStatement) GetCondition() expression.Expression {
	return s.condition
}

func (s *WhileStatement) SetCondition(condition expression.Expression) {
	s.condition = condition
}

func (s *WhileStatement) GetStatements() Statement {
	return s.statements
}

func (s *WhileStatement) IsWhileStatement() bool {
	return true
}

func (s *WhileStatement) Accept(visitor StatementVisitor) {
	visitor.VisitWhileStatement(s)
}
