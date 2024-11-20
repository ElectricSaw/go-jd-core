package expression

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewThisExpression(typ intsyn.IType) intsyn.IThisExpression {
	return &ThisExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpressionEmpty(),
		typ:                          typ,
		explicit:                     true,
	}
}

func NewThisExpressionWithAll(lineNumber int, typ intsyn.IType) intsyn.IThisExpression {
	return &ThisExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpression(lineNumber),
		typ:                          typ,
		explicit:                     true,
	}
}

type ThisExpression struct {
	AbstractLineNumberExpression

	typ      intsyn.IType
	explicit bool
}

func (e *ThisExpression) Type() intsyn.IType {
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

func (e *ThisExpression) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitThisExpression(e)
}

func (e *ThisExpression) String() string {
	return fmt.Sprintf("ThisExpression{%s}", e.typ)
}
