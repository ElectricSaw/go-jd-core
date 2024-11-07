package statement

import "fmt"

func NewLabelStatement(label string, statement Statement) *LabelStatement {
	return &LabelStatement{
		label:     label,
		statement: statement,
	}
}

type LabelStatement struct {
	AbstractStatement

	label     string
	statement Statement
}

func (s *LabelStatement) GetLabel() string {
	return s.label
}

func (s *LabelStatement) GetStatement() Statement {
	return s.statement
}

func (s *LabelStatement) IsLabelStatement() bool {
	return true
}

func (s *LabelStatement) Accept(visitor StatementVisitor) {
	visitor.VisitLabelStatement(s)
}

func (s *LabelStatement) String() string {
	return fmt.Sprintf("LabelStatement{%s: %s}", s.label, s.statement)
}
