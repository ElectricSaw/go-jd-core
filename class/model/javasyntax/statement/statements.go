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

func (s *Statements) IsList() bool {
	return s.DefaultList.IsList()
}

func (s *Statements) Size() int {
	return s.DefaultList.Size()
}

func (s *Statements) ToSlice() []intmod.IStatement {
	return s.DefaultList.ToSlice()
}

func (s *Statements) ToList() *util.DefaultList[intmod.IStatement] {
	return s.DefaultList.ToList()
}

func (s *Statements) First() intmod.IStatement {
	return s.DefaultList.First()
}

func (s *Statements) Last() intmod.IStatement {
	return s.DefaultList.Last()
}

func (s *Statements) Iterator() util.IIterator[intmod.IStatement] {
	return s.DefaultList.Iterator()
}

func (s *Statements) IsStatements() bool {
	return true
}

func (s *Statements) Accept(visitor intmod.IStatementVisitor) {
	visitor.VisitStatements(s)
}
