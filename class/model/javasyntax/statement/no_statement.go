package statement

var NoStmt = NewNoStatement()

func NewNoStatement() *NoStatement {
	return new(NoStatement)
}

type NoStatement struct {
	AbstractStatement
}

func (s *NoStatement) Accept(visitor StatementVisitor) {
	visitor.VisitNoStatement(s)
}

func (s *NoStatement) String() string {
	return "NoStatement"
}
