package statement

import intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"

var Break = NewBreakStatement("")

func NewBreakStatement(label string) intmod.IBreakStatement {
	return &BreakStatement{
		label: label,
	}
}

type BreakStatement struct {
	AbstractStatement

	label string
}

func (s *BreakStatement) Text() string {
	return s.label
}

func (s *BreakStatement) IsBreakStatement() bool {
	return true
}

func (s *BreakStatement) AcceptStatement(visitor intmod.IStatementVisitor) {
	visitor.VisitBreakStatement(s)
}
