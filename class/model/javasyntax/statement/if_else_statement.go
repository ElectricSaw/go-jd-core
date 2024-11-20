package statement

import intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"

func NewIfElseStatement(condition intsyn.IExpression, statements intsyn.IStatement, elseStatements intsyn.IStatement) intsyn.IIfElseStatement {
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

	elseStatements intsyn.IStatement
}

func (s *IfElseStatement) ElseStatements() intsyn.IStatement {
	return s.elseStatements
}

func (s *IfElseStatement) IsIfElseStatement() bool {
	return true
}

func (s *IfElseStatement) Accept(visitor intsyn.IStatementVisitor) {
	visitor.VisitIfElseStatement(s)
}
