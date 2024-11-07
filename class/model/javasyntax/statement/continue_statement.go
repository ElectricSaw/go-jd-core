package statement

var Continue = NewContinueStatement("")

func NewContinueStatement(label string) *ContinueStatement {
	return &ContinueStatement{
		label: label,
	}
}

type ContinueStatement struct {
	AbstractStatement

	label string
}

func (s *ContinueStatement) GetLabel() string {
	return s.label
}

func (s *ContinueStatement) IsContinueStatement() bool {
	return true
}

func (s *ContinueStatement) Accept(visitor StatementVisitor) {
	visitor.VisitContinueStatement(s)
}
