package expression

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	"fmt"
)

func NewNullExpression(typ intsyn.IType) intsyn.INullExpression {
	return &NullExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(typ),
	}
}

func NewNullExpressionWithAll(lineNumber int, typ intsyn.IType) intsyn.INullExpression {
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

func (e *NullExpression) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitNullExpression(e)
}

func (e *NullExpression) String() string {
	return fmt.Sprintf("NullExpression{type=%s}", e.typ)
}
