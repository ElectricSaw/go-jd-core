package statement

import "bitbucket.org/coontec/javaClass/java/model/javasyntax/expression"

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

func (s *SynchronizedStatement) GetMonitor() expression.Expression {
	return s.monitor
}

func (s *SynchronizedStatement) SetMonitor(monitor expression.Expression) {
	s.monitor = monitor
}

func (s *SynchronizedStatement) GetStatements() Statement {
	return s.statements
}

func (s *SynchronizedStatement) Accept(visitor StatementVisitor) {
	visitor.VisitSynchronizedStatement(s)

}
