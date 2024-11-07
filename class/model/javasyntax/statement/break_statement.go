package statement

var Break = NewBreakStatement("")

func NewBreakStatement(label string) *BreakStatement {
	return &BreakStatement{
		label: label,
	}
}

type BreakStatement struct {
	AbstractStatement

	label string
}

func (s *BreakStatement) GetLabel() string {
	return s.label
}

func (s *BreakStatement) IsBreakStatement() bool {
	return true
}

func (s *BreakStatement) Accept(visitor StatementVisitor) {
	visitor.VisitBreakStatement(s)
}
