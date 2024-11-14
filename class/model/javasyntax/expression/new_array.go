package expression

import (
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewNewArray(lineNumber int, typ _type.IType, dimensionExpressionList Expression) *NewArray {
	return &NewArray{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		dimensionExpressionList:          dimensionExpressionList,
	}
}

type NewArray struct {
	AbstractLineNumberTypeExpression

	dimensionExpressionList Expression
}

func (e *NewArray) DimensionExpressionList() Expression {
	return e.dimensionExpressionList
}

func (e *NewArray) SetDimensionExpressionList(dimensionExpressionList Expression) {
	e.dimensionExpressionList = dimensionExpressionList
}

func (e *NewArray) Priority() int {
	return 0
}

func (e *NewArray) IsNewArray() bool {
	return true
}

func (e *NewArray) Accept(visitor ExpressionVisitor) {
	visitor.VisitNewArray(e)
}

func (e *NewArray) String() string {
	return fmt.Sprintf("NewArray{%s}", e.typ)
}
