package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewNewArray(lineNumber int, typ intmod.IType, dimensionExpressionList intmod.IExpression) intmod.INewArray {
	return &NewArray{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		dimensionExpressionList:          dimensionExpressionList,
	}
}

type NewArray struct {
	AbstractLineNumberTypeExpression

	dimensionExpressionList intmod.IExpression
}

func (e *NewArray) DimensionExpressionList() intmod.IExpression {
	return e.dimensionExpressionList
}

func (e *NewArray) SetDimensionExpressionList(dimensionExpressionList intmod.IExpression) {
	e.dimensionExpressionList = dimensionExpressionList
}

func (e *NewArray) Priority() int {
	return 0
}

func (e *NewArray) IsNewArray() bool {
	return true
}

func (e *NewArray) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitNewArray(e)
}

func (e *NewArray) String() string {
	return fmt.Sprintf("NewArray{%s}", e.typ)
}
