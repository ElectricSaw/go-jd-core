package statement

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/class/util"
)

func NewStatements() intmod.IStatements {
	s := &Statements{
		DefaultList: *util.NewDefaultList[intmod.IStatement]().(*util.DefaultList[intmod.IStatement]),
	}
	s.SetValue(s)
	return s
}

func NewStatementsWithList(list util.IList[intmod.IStatement]) intmod.IStatements {
	return NewStatementsWithSlice(list.ToSlice())
}

func NewStatementsWithSlice(slice []intmod.IStatement) intmod.IStatements {
	s := &Statements{
		DefaultList: *util.NewDefaultListWithSlice(slice).(*util.DefaultList[intmod.IStatement]),
	}
	s.SetValue(s)
	return s
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

func (s *Statements) AcceptStatement(visitor intmod.IStatementVisitor) {
	visitor.VisitStatements(s)
}
