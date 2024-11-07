package statement

import "bitbucket.org/coontec/javaClass/class/model/javasyntax/expression"

func NewDoWhileStatement(condition expression.Expression, statements Statement) *DoWhileStatement {
	return &DoWhileStatement{
		condition:  condition,
		statements: statements,
	}
}

type DoWhileStatement struct {
	AbstractStatement

	condition  expression.Expression
	statements Statement
}

func (s *DoWhileStatement) GetCondition() expression.Expression {
	return s.condition
}

func (s *DoWhileStatement) SetCondition(condition expression.Expression) {
	s.condition = condition
}

func (s *DoWhileStatement) GetStatements() Statement {
	return s.statements
}

func (s *DoWhileStatement) Accept(visitor StatementVisitor) {
	visitor.VisitDoWhileStatement(s)
}
