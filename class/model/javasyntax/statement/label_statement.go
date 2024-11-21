package statement

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
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

func (s *LabelStatement) Label() string {
	return s.label
}

func (s *LabelStatement) Statement() intmod.IStatement {
	return s.statement
}

func (s *LabelStatement) IsLabelStatement() bool {
	return true
}

func (s *LabelStatement) Accept(visitor intmod.IStatementVisitor) {
	visitor.VisitLabelStatement(s)
}

func (s *LabelStatement) String() string {
	return fmt.Sprintf("LabelStatement{%s: %s}", s.label, s.statement)
}
