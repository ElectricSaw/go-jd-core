package statement

import intsyn "bitbucket.org/coontec/javaClass/class/interfaces/model"

var Return = NewReturnStatement()

func NewReturnStatement() intsyn.IReturnStatement {
	return &ReturnStatement{}
}

type ReturnStatement struct {
	AbstractStatement
}

func (s *ReturnStatement) IsReturnStatement() bool {
	return true
}

func (s *ReturnStatement) Accept(visitor intsyn.IStatementVisitor) {
	visitor.VisitReturnStatement(s)
}

func (s *ReturnStatement) String() string {
	return "ReturnStatement{}"
}
