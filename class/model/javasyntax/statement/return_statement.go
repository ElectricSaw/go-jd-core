package statement

import intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"

var Return = NewReturnStatement()

func NewReturnStatement() intmod.IReturnStatement {
	return &ReturnStatement{}
}

type ReturnStatement struct {
	AbstractStatement
}

func (s *ReturnStatement) IsReturnStatement() bool {
	return true
}

func (s *ReturnStatement) AcceptStatement(visitor intmod.IStatementVisitor) {
	visitor.VisitReturnStatement(s)
}

func (s *ReturnStatement) String() string {
	return "ReturnStatement{}"
}
