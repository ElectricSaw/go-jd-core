package expression

import (
	_type "bitbucket.org/coontec/javaClass/java/model/javasyntax/type"
	"fmt"
)

func NewFloatConstantExpression(value float32) *FloatConstantExpression {
	return &FloatConstantExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(_type.PtTypeFloat),
		value:                            value,
	}
}

func NewFloatConstantExpressionWithAll(lineNumber int, value float32) *FloatConstantExpression {
	return &FloatConstantExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, _type.PtTypeFloat),
		value:                            value,
	}
}

type FloatConstantExpression struct {
	AbstractLineNumberTypeExpression

	value float32
}

func (e *FloatConstantExpression) GetFloatValue() float32 {
	return e.value
}

func (e *FloatConstantExpression) IsFloatConstantExpression() bool {
	return true
}

func (e *FloatConstantExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitFloatConstantExpression(e)
}

func (e *FloatConstantExpression) String() string {
	return fmt.Sprintf("FloatConstantExpression{%.2f}", e.value)
}
