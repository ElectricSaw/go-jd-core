package expression

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
)

var NeNoExpression = NewNoExpression()

func NewNoExpression() intmod.INoExpression {
	e := &NoExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(model.PtTypeVoid.(intmod.IType)),
	}
	e.SetValue(e)
	return e
}

type NoExpression struct {
	AbstractLineNumberTypeExpression
}

func (e *NoExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitNoExpression(e)
}

func (e *NoExpression) String() string {
	return "NoExpression"
}
