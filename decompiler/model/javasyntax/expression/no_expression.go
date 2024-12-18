package expression

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	_type "github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/type"
)

var NeNoExpression = NewNoExpression()

func NewNoExpression() intmod.INoExpression {
	e := &NoExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(_type.PtTypeVoid.(intmod.IType)),
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
