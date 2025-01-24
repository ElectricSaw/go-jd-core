package expression

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
)

func CreateItemType(expression IExpression) model.IType {
	typ := expression.Type()
	dimension := typ.Dimension()

	if dimension > 0 {
		return typ.CreateType(dimension - 1)
	}

	return typ.CreateType(0)
}

func NewArrayExpression(expression IExpression, index IExpression) IArrayExpression {
	return NewArrayExpressionWithAll(0, expression, index)
}

func NewArrayExpressionWithAll(lineNumber int, expression IExpression, index IExpression) IArrayExpression {
	e := &ArrayExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, CreateItemType(expression)),
		expression:                       expression,
		index:                            index,
	}
	e.SetValue(e)
	return e
}

type ArrayExpression struct {
	AbstractLineNumberTypeExpression

	expression IExpression
	index      IExpression
}

func (e *ArrayExpression) Expression() IExpression {
	return e.expression
}

func (e *ArrayExpression) Index() IExpression {
	return e.index
}

func (e *ArrayExpression) Priority() int {
	return 1
}

func (e *ArrayExpression) SetExpression(expression IExpression) {
	e.expression = expression
}

func (e *ArrayExpression) SetIndex(index IExpression) {
	e.index = index
}

func (e *ArrayExpression) IsArrayExpression() bool {
	return true
}

func (e *ArrayExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitArrayExpression(e)
}

func (e *ArrayExpression) String() string {
	return fmt.Sprintf("ArrayExpression{%v[%v]}", e.expression, e.index)
}
