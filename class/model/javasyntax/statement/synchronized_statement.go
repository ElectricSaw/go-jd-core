package statement

import "bitbucket.org/coontec/javaClass/class/model/javasyntax/expression"

func NewSynchronizedStatement(monitor expression.Expression, statements Statement) *SynchronizedStatement {
	return &SynchronizedStatement{
		monitor:    monitor,
		statements: statements,
	}
}

type SynchronizedStatement struct {
	AbstractStatement

	monitor    expression.Expression
	statements Statement
}

func (s *SynchronizedStatement) Monitor() expression.Expression {
	return s.monitor
}

func (s *SynchronizedStatement) SetMonitor(monitor expression.Expression) {
	s.monitor = monitor
}

func (s *SynchronizedStatement) Statements() Statement {
	return s.statements
}

func (s *SynchronizedStatement) Accept(visitor StatementVisitor) {
	visitor.VisitSynchronizedStatement(s)

}
