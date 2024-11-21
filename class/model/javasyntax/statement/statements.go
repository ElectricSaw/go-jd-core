package statement

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

func NewStatements() intmod.IStatements {
	return &Statements{}
}

type Statements struct {
	AbstractStatement
	util.DefaultList[intmod.IStatement]
}

func (s *Statements) IsStatements() bool {
	return true
}

func (s *Statements) Accept(visitor intmod.IStatementVisitor) {
	visitor.VisitStatements(s)
}
