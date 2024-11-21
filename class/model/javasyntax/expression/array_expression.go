package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func CreateItemType(expression intmod.IExpression) intmod.IType {
	typ := expression.Type()
	dimension := typ.Dimension()

	if dimension > 0 {
		return typ.CreateType(dimension - 1)
	}

	return typ.CreateType(0)
}

func NewArrayExpression(expression intmod.IExpression, index intmod.IExpression) intmod.IArrayExpression {
	return &ArrayExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(CreateItemType(expression)),
		expression:                       expression,
		index:                            index,
	}
}

func NewArrayExpressionWithLineNumber(lineNumber int, expression intmod.IExpression, index intmod.IExpression) intmod.IArrayExpression {
	return &ArrayExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, CreateItemType(expression)),
		expression:                       expression,
		index:                            index,
	}
}

type ArrayExpression struct {
	AbstractLineNumberTypeExpression

	expression intmod.IExpression
	index      intmod.IExpression
}

func (e *ArrayExpression) Expression() intmod.IExpression {
	return e.expression
}

func (e *ArrayExpression) Index() intmod.IExpression {
	return e.index
}

func (e *ArrayExpression) Priority() int {
	return 1
}

func (e *ArrayExpression) SetExpression(expression intmod.IExpression) {
	e.expression = expression
}

func (e *ArrayExpression) SetIndex(index intmod.IExpression) {
	e.index = index
}

func (e *ArrayExpression) IsArrayExpression() bool {
	return true
}

func (e *ArrayExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitArrayExpression(e)
}

func (e *ArrayExpression) String() string {
	return fmt.Sprintf("ArrayExpression{%v[%v]}", e.expression, e.index)
}
