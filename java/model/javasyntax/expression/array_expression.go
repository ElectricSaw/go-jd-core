package expression

import (
	_type "bitbucket.org/coontec/javaClass/java/model/javasyntax/type"
	"fmt"
)

func NewArrayExpression(expression Expression, index Expression) *ArrayExpression {
	return &ArrayExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(CreateItemType(expression)),
		expression:                       expression,
		index:                            index,
	}
}

func NewArrayExpressionWithLineNumber(lineNumber int, expression Expression, index Expression) *ArrayExpression {
	return &ArrayExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, CreateItemType(expression)),
		expression:                       expression,
		index:                            index,
	}
}

func CreateItemType(expression Expression) _type.IType {
	typ := expression.GetType()
	dimension := typ.Dimension()

	if dimension > 0 {
		return typ.CreateType(dimension - 1)
	}

	return typ.CreateType(0)
}

type ArrayExpression struct {
	AbstractLineNumberTypeExpression

	expression Expression
	index      Expression
}

func (e *ArrayExpression) GetExpression() Expression {
	return e.expression
}

func (e *ArrayExpression) GetIndex() Expression {
	return e.index
}

func (e *ArrayExpression) GetPriority() int {
	return 1
}

func (e *ArrayExpression) SetExpression(expression Expression) {
	e.expression = expression
}

func (e *ArrayExpression) SetIndex(index Expression) {
	e.index = index
}

func (e *ArrayExpression) IsArrayExpression() bool {
	return true
}

func (e *ArrayExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitArrayExpression(e)
}

func (e *ArrayExpression) String() string {
	return fmt.Sprintf("ArrayExpression{%v[%v]}", e.expression, e.index)
}
