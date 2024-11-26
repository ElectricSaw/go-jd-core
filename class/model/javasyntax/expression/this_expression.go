package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
	"fmt"
)

func NewThisExpression(typ intmod.IType) intmod.IThisExpression {
	return NewThisExpressionWithAll(0, typ)
}

func NewThisExpressionWithAll(lineNumber int, typ intmod.IType) intmod.IThisExpression {
	e := &ThisExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpression(lineNumber),
		typ:                          typ,
		explicit:                     true,
	}
	e.SetValue(e)
	return e
}

type ThisExpression struct {
	AbstractLineNumberExpression
	util.DefaultBase[intmod.IThisExpression]

	typ      intmod.IType
	explicit bool
}

func (e *ThisExpression) Type() intmod.IType {
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

func (e *ThisExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitThisExpression(e)
}

func (e *ThisExpression) String() string {
	return fmt.Sprintf("ThisExpression{%s}", e.typ)
}
