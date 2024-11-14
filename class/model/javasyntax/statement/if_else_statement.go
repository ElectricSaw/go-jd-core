package statement

import "bitbucket.org/coontec/javaClass/class/model/javasyntax/expression"

func NewIfElseStatement(condition expression.Expression, statements Statement, elseStatements Statement) *IfElseStatement {
	return &IfElseStatement{
		IfStatement: IfStatement{
			condition:  condition,
			statements: statements,
		},
		elseStatements: elseStatements,
	}
}

type IfElseStatement struct {
	IfStatement

	elseStatements Statement
}

func (s *IfElseStatement) ElseStatements() Statement {
	return s.elseStatements
}

func (s *IfElseStatement) IsIfElseStatement() bool {
	return true
}

func (s *IfElseStatement) Accept(visitor StatementVisitor) {
	visitor.VisitIfElseStatement(s)
}
