package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
	"fmt"
)

func NewFloatConstantExpression(value float32) intmod.IFloatConstantExpression {
	return &FloatConstantExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(_type.PtTypeFloat.(intmod.IType)),
		value:                            value,
	}
}

func NewFloatConstantExpressionWithAll(lineNumber int, value float32) intmod.IFloatConstantExpression {
	return &FloatConstantExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, _type.PtTypeFloat.(intmod.IType)),
		value:                            value,
	}
}

type FloatConstantExpression struct {
	AbstractLineNumberTypeExpression

	value float32
}

func (e *FloatConstantExpression) FloatValue() float32 {
	return e.value
}

func (e *FloatConstantExpression) IsFloatConstantExpression() bool {
	return true
}

func (e *FloatConstantExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitFloatConstantExpression(e)
}

func (e *FloatConstantExpression) String() string {
	return fmt.Sprintf("FloatConstantExpression{%.2f}", e.value)
}
