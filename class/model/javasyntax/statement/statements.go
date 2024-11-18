package statement

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/model"
	"bitbucket.org/coontec/javaClass/class/util"
)

func NewStatements() intsyn.IStatements {
	return &Statements{}
}

type Statements struct {
	AbstractStatement
	util.DefaultList[intsyn.IStatement]
}

func (s *Statements) IsStatements() bool {
	return true
}

func (s *Statements) Accept(visitor intsyn.IStatementVisitor) {
	visitor.VisitStatements(s)
}
