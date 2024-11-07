package expression

import (
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewNullExpression(typ _type.IType) *NullExpression {
	return &NullExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(typ),
	}
}

func NewNullExpressionWithAll(lineNumber int, typ _type.IType) *NullExpression {
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

func (e *NullExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitNullExpression(e)
}

func (e *NullExpression) String() string {
	return fmt.Sprintf("NullExpression{type=%s}", e.typ)
}
