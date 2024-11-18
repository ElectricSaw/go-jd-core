package expression

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewArrayExpression(expression intsyn.IExpression, index intsyn.IExpression) intsyn.IArrayExpression {
	return &ArrayExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(_type.CreateItemType(expression)),
		expression:                       expression,
		index:                            index,
	}
}

func NewArrayExpressionWithLineNumber(lineNumber int, expression intsyn.IExpression, index intsyn.IExpression) intsyn.IArrayExpression {
	return &ArrayExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, _type.CreateItemType(expression)),
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
