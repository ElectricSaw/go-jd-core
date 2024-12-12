package expression

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	_type "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/type"
)

func NewCommentExpression(text string) intmod.ICommentExpression {
	e := &CommentExpression{
		text: text,
	}
	e.SetValue(e)
	return e
}

type CommentExpression struct {
	AbstractExpression

	text string
}

func (e *CommentExpression) LineNumber() int {
	return intmod.UnknownLineNumber
}

func (e *CommentExpression) Type() intmod.IType {
	return _type.PtTypeVoid.(intmod.IType)
}

func (e *CommentExpression) Priority() int {
	return 0
}

func (e *CommentExpression) Text() string {
	return e.text
}

func (e *CommentExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitCommentExpression(e)
}

func (e *CommentExpression) String() string {
	return fmt.Sprintf("CommentExpression{%s}", e.text)
}
