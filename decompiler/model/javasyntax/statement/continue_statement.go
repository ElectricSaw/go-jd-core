package statement

import intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"

var Continue = NewContinueStatement("")

func NewContinueStatement(label string) intmod.IContinueStatement {
	return &ContinueStatement{
		label: label,
	}
}

type ContinueStatement struct {
	AbstractStatement

	label string
}

func (s *ContinueStatement) Text() string {
	return s.label
}

func (s *ContinueStatement) IsContinueStatement() bool {
	return true
}

func (s *ContinueStatement) AcceptStatement(visitor intmod.IStatementVisitor) {
	visitor.VisitContinueStatement(s)
}
