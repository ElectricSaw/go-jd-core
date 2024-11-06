package statement

import "bitbucket.org/coontec/javaClass/java/model/javasyntax/expression"

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

func (s *IfStatement) GetCondition() expression.Expression {
	return s.condition
}

func (s *IfStatement) SetCondition(condition expression.Expression) {
	s.condition = condition
}

func (s *IfStatement) GetStatements() Statement {
	return s.statements
}

func (s *IfStatement) IsIfStatement() bool {
	return s.condition != nil
}

func (s *IfStatement) Accept(visitor StatementVisitor) {
	visitor.VisitIfStatement(s)
}
