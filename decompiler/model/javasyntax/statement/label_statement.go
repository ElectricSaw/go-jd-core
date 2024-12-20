package statement

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewLabelStatement(label string, statement intmod.IStatement) intmod.ILabelStatement {
	return &LabelStatement{
		label:     label,
		statement: statement,
	}
}

type LabelStatement struct {
	AbstractStatement

	label     string
	statement intmod.IStatement
}

func (s *LabelStatement) Text() string {
	return s.label
}

func (s *LabelStatement) Statement() intmod.IStatement {
	return s.statement
}

func (s *LabelStatement) IsLabelStatement() bool {
	return true
}

func (s *LabelStatement) AcceptStatement(visitor intmod.IStatementVisitor) {
	visitor.VisitLabelStatement(s)
}

func (s *LabelStatement) String() string {
	return fmt.Sprintf("LabelStatement{%s: %s}", s.label, s.statement)
}
