package statement

import intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"

var NoStmt = NewNoStatement()

func NewNoStatement() intmod.INoStatement {
	return new(NoStatement)
}

type NoStatement struct {
	AbstractStatement
}

func (s *NoStatement) AcceptStatement(visitor intmod.IStatementVisitor) {
	visitor.VisitNoStatement(s)
}

func (s *NoStatement) String() string {
	return "NoStatement"
}
