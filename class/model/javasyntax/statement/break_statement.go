package statement

import intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"

var Break = NewBreakStatement("")

func NewBreakStatement(label string) intsyn.IBreakStatement {
	return &BreakStatement{
		label: label,
	}
}

type BreakStatement struct {
	AbstractStatement

	label string
}

func (s *BreakStatement) Label() string {
	return s.label
}

func (s *BreakStatement) IsBreakStatement() bool {
	return true
}

func (s *BreakStatement) Accept(visitor intsyn.IStatementVisitor) {
	visitor.VisitBreakStatement(s)
}
