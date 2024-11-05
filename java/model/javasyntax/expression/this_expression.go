package expression

import (
	_type "bitbucket.org/coontec/javaClass/java/model/javasyntax/type"
	"fmt"
)

func NewThisExpression(typ _type.IType) *ThisExpression {
	return &ThisExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpressionEmpty(),
		typ:                          typ,
		explicit:                     true,
	}
}

func NewThisExpressionWithAll(lineNumber int, typ _type.IType) *ThisExpression {
	return &ThisExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpression(lineNumber),
		typ:                          typ,
		explicit:                     true,
	}
}

type ThisExpression struct {
	AbstractLineNumberExpression

	typ      _type.IType
	explicit bool
}

func (e *ThisExpression) GetType() _type.IType {
	return e.typ
}

func (e *ThisExpression) IsExplicit() bool {
	return e.explicit
}

func (e *ThisExpression) SetExplicit(explicit bool) {
	e.explicit = explicit
}

func (e *ThisExpression) IsThisExpression() bool {
	return true
}

func (e *ThisExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitThisExpression(e)
}

func (e *ThisExpression) String() string {
	return fmt.Sprintf("ThisExpression{%s}", e.typ)
}
