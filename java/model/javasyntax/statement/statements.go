package statement

type Statements struct {
	AbstractStatement

	Statements []Statement
}

func (s *Statements) IsStatements() bool {
	return true
}

func (s *Statements) Accept(visitor StatementVisitor) {
	visitor.VisitStatements(s)
}
