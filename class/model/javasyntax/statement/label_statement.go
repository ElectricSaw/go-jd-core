package statement

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/model"
	"fmt"
)

func NewLabelStatement(label string, statement intsyn.IStatement) intsyn.ILabelStatement {
	return &LabelStatement{
		label:     label,
		statement: statement,
	}
}

type LabelStatement struct {
	AbstractStatement

	label     string
	statement intsyn.IStatement
}

func (s *LabelStatement) Label() string {
	return s.label
}

func (s *LabelStatement) Statement() intsyn.IStatement {
	return s.statement
}

func (s *LabelStatement) IsLabelStatement() bool {
	return true
}

func (s *LabelStatement) Accept(visitor intsyn.IStatementVisitor) {
	visitor.VisitLabelStatement(s)
}

func (s *LabelStatement) String() string {
	return fmt.Sprintf("LabelStatement{%s: %s}", s.label, s.statement)
}
