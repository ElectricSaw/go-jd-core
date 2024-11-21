package statement

import intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"

var NoStmt = NewNoStatement()

func NewNoStatement() intmod.INoStatement {
	return new(NoStatement)
}

type NoStatement struct {
	AbstractStatement
}

func (s *NoStatement) Accept(visitor intmod.IStatementVisitor) {
	visitor.VisitNoStatement(s)
}

func (s *NoStatement) String() string {
	return "NoStatement"
}
