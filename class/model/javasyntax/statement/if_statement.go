package statement

import "bitbucket.org/coontec/javaClass/class/model/javasyntax/expression"

func NewIfStatement(condition expression.Expression, statements Statement) *IfStatement {
	return &IfStatement{
		condition:  condition,
		statements: statements,
	}
}

type IfStatement struct {
	AbstractStatement

	condition  expression.Expression
	statements Statement
}

func (s *IfStatement) Condition() expression.Expression {
	return s.condition
}

func (s *IfStatement) SetCondition(condition expression.Expression) {
	s.condition = condition
}

func (s *IfStatement) Statements() Statement {
	return s.statements
}

func (s *IfStatement) IsIfStatement() bool {
	return s.condition != nil
}

func (s *IfStatement) Accept(visitor StatementVisitor) {
	visitor.VisitIfStatement(s)
}
