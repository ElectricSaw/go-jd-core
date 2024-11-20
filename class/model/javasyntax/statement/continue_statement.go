package statement

import intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"

var Continue = NewContinueStatement("")

func NewContinueStatement(label string) intsyn.IContinueStatement {
	return &ContinueStatement{
		label: label,
	}
}

type ContinueStatement struct {
	AbstractStatement

	label string
}

func (s *ContinueStatement) Label() string {
	return s.label
}

func (s *ContinueStatement) IsContinueStatement() bool {
	return true
}

func (s *ContinueStatement) Accept(visitor intsyn.IStatementVisitor) {
	visitor.VisitContinueStatement(s)
}
