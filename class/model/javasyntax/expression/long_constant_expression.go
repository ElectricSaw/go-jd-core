package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
	"fmt"
)

func NewLongConstantExpression(value int64) intmod.ILongConstantExpression {
	return &LongConstantExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(_type.PtTypeLong.(intmod.IType)),
		value:                            value,
	}
}

func NewLongConstantExpressionWithAll(lineNumber int, value int64) intmod.ILongConstantExpression {
	return &LongConstantExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, _type.PtTypeLong.(intmod.IType)),
		value:                            value,
	}
}

type LongConstantExpression struct {
	AbstractLineNumberTypeExpression

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
