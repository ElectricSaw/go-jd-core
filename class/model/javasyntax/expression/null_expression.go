package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewNullExpression(typ intmod.IType) intmod.INullExpression {
	return &NullExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(typ),
	}
}

func NewNullExpressionWithAll(lineNumber int, typ intmod.IType) intmod.INullExpression {
	return &NullExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
	}
}

type NullExpression struct {
	AbstractLineNumberTypeExpression
}

func (e *NullExpression) IsNullExpression() bool {
	return true
}

func (e *NullExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitNullExpression(e)
}

func (e *NullExpression) String() string {
	return fmt.Sprintf("NullExpression{type=%s}", e.typ)
}
