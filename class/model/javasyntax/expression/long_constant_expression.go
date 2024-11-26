package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
	"bitbucket.org/coontec/go-jd-core/class/util"
	"fmt"
)

func NewLongConstantExpression(value int64) intmod.ILongConstantExpression {
	return NewLongConstantExpressionWithAll(0, value)
}

func NewLongConstantExpressionWithAll(lineNumber int, value int64) intmod.ILongConstantExpression {
	e := &LongConstantExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, _type.PtTypeLong.(intmod.IType)),
		value:                            value,
	}
	e.SetValue(e)
	return e
}

type LongConstantExpression struct {
	AbstractLineNumberTypeExpression
	util.DefaultBase[intmod.ILongConstantExpression]

	value int64
}

func (e *LongConstantExpression) LongValue() int64 {
	return e.value
}

func (e *LongConstantExpression) IsLongConstantExpression() bool {
	return true
}

func (e *LongConstantExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitLongConstantExpression(e)
}

func (e *LongConstantExpression) String() string {
	return fmt.Sprintf("LongConstantExpression{%d}", e.value)
}
