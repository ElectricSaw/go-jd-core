package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
	"fmt"
)

func NewCommentExpression(text string) intmod.ICommentExpression {
	return &CommentExpression{
		text: text,
	}
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
