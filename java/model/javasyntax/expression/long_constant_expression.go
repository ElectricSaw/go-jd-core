package expression

import (
	_type "bitbucket.org/coontec/javaClass/java/model/javasyntax/type"
	"fmt"
)

func NewLongConstantExpression(value int64) *LongConstantExpression {
	return &LongConstantExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(_type.PtTypeLong),
		value:                            value,
	}
}

func NewLongConstantExpressionWithAll(lineNumber int, value int64) *LongConstantExpression {
	return &LongConstantExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, _type.PtTypeLong),
		value:                            value,
	}
}

type LongConstantExpression struct {
	AbstractLineNumberTypeExpression

	value int64
}

func (e *LongConstantExpression) GetLongValue() int64 {
	return e.value
}

func (e *LongConstantExpression) IsLongConstantExpression() bool {
	return true
}

func (e *LongConstantExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitLongConstantExpression(e)
}

func (e *LongConstantExpression) String() string {
	return fmt.Sprintf("LongConstantExpression{%d}", e.value)
}
