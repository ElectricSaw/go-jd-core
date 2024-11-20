package expression

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func CreateItemType(expression intsyn.IExpression) intsyn.IType {
	typ := expression.Type()
	dimension := typ.Dimension()

	if dimension > 0 {
		return typ.CreateType(dimension - 1)
	}

	return typ.CreateType(0)
}

func NewArrayExpression(expression intsyn.IExpression, index intsyn.IExpression) intsyn.IArrayExpression {
	return &ArrayExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(CreateItemType(expression)),
		expression:                       expression,
		index:                            index,
	}
}

func NewArrayExpressionWithLineNumber(lineNumber int, expression intsyn.IExpression, index intsyn.IExpression) intsyn.IArrayExpression {
	return &ArrayExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, CreateItemType(expression)),
		expression:                       expression,
		index:                            index,
	}
}

type ArrayExpression struct {
	AbstractLineNumberTypeExpression

	expression intsyn.IExpression
	index      intsyn.IExpression
}

func (e *ArrayExpression) Expression() intsyn.IExpression {
	return e.expression
}

func (e *ArrayExpression) Index() intsyn.IExpression {
	return e.index
}

func (e *ArrayExpression) Priority() int {
	return 1
}

func (e *ArrayExpression) SetExpression(expression intsyn.IExpression) {
	e.expression = expression
}

func (e *ArrayExpression) SetIndex(index intsyn.IExpression) {
	e.index = index
}

func (e *ArrayExpression) IsArrayExpression() bool {
	return true
}

func (e *ArrayExpression) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitArrayExpression(e)
}

func (e *ArrayExpression) String() string {
	return fmt.Sprintf("ArrayExpression{%v[%v]}", e.expression, e.index)
}
