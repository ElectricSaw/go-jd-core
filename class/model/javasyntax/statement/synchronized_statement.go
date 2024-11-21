package statement

import intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"

func NewSynchronizedStatement(monitor intmod.IExpression, statements intmod.IStatement) intmod.ISynchronizedStatement {
	return &SynchronizedStatement{
		monitor:    monitor,
		statements: statements,
	}
}

type SynchronizedStatement struct {
	AbstractStatement

	monitor    intmod.IExpression
	statements intmod.IStatement
}

func (s *SynchronizedStatement) Monitor() intmod.IExpression {
	return s.monitor
}

func (s *SynchronizedStatement) SetMonitor(monitor intmod.IExpression) {
	s.monitor = monitor
}

func (s *SynchronizedStatement) Statements() intmod.IStatement {
	return s.statements
}

func (s *SynchronizedStatement) Accept(visitor intmod.IStatementVisitor) {
	visitor.VisitSynchronizedStatement(s)

}
