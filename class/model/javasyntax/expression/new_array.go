package expression

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	"fmt"
)

func NewNewArray(lineNumber int, typ intsyn.IType, dimensionExpressionList intsyn.IExpression) intsyn.INewArray {
	return &NewArray{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		dimensionExpressionList:          dimensionExpressionList,
	}
}

type NewArray struct {
	AbstractLineNumberTypeExpression

	dimensionExpressionList intsyn.IExpression
}

func (e *NewArray) DimensionExpressionList() intsyn.IExpression {
	return e.dimensionExpressionList
}

func (e *NewArray) SetDimensionExpressionList(dimensionExpressionList intsyn.IExpression) {
	e.dimensionExpressionList = dimensionExpressionList
}

func (e *NewArray) Priority() int {
	return 0
}

func (e *NewArray) IsNewArray() bool {
	return true
}

func (e *NewArray) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitNewArray(e)
}

func (e *NewArray) String() string {
	return fmt.Sprintf("NewArray{%s}", e.typ)
}
