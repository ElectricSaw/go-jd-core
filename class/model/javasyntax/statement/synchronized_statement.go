package statement

import intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"

func NewSynchronizedStatement(monitor intsyn.IExpression, statements intsyn.IStatement) intsyn.ISynchronizedStatement {
	return &SynchronizedStatement{
		monitor:    monitor,
		statements: statements,
	}
}

type SynchronizedStatement struct {
	AbstractStatement

	monitor    intsyn.IExpression
	statements intsyn.IStatement
}

func (s *SynchronizedStatement) Monitor() intsyn.IExpression {
	return s.monitor
}

func (s *SynchronizedStatement) SetMonitor(monitor intsyn.IExpression) {
	s.monitor = monitor
}

func (s *SynchronizedStatement) Statements() intsyn.IStatement {
	return s.statements
}

func (s *SynchronizedStatement) Accept(visitor intsyn.IStatementVisitor) {
	visitor.VisitSynchronizedStatement(s)

}
