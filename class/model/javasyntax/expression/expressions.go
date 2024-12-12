package expression

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/class/util"
)

func NewExpressions() intmod.IExpressions {
	return &Expressions{}
}

type Expressions struct {
	AbstractExpression
	util.DefaultList[intmod.IExpression]
}

func (e *Expressions) IsList() bool {
	return e.DefaultList.IsList()
}

func (e *Expressions) Size() int {
	return e.DefaultList.Size()
}

func (e *Expressions) ToSlice() []intmod.IExpression {
	return e.DefaultList.ToSlice()
}

func (e *Expressions) ToList() *util.DefaultList[intmod.IExpression] {
	return e.DefaultList.ToList()
}

func (e *Expressions) First() intmod.IExpression {
	return e.DefaultList.First()
}

func (e *Expressions) Last() intmod.IExpression {
	return e.DefaultList.Last()
}

func (e *Expressions) Iterator() util.IIterator[intmod.IExpression] {
	return e.DefaultList.Iterator()
}

func (e *Expressions) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitExpressions(e)
}
