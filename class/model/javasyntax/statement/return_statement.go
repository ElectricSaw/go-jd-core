package statement

var Return = NewReturnStatement()

func NewReturnStatement() *ReturnStatement {
	return &ReturnStatement{}
}

type ReturnStatement struct {
	AbstractStatement
}

func (s *ReturnStatement) IsReturnStatement() bool {
	return true
}

func (s *ReturnStatement) Accept(visitor StatementVisitor) {
	visitor.VisitReturnStatement(s)
}

func (s *ReturnStatement) String() string {
	return "ReturnStatement{}"
}
