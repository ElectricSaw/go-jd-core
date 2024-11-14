package expression

import (
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewCommentExpression(text string) *CommentExpression {
	return &CommentExpression{
		text: text,
	}
}

type CommentExpression struct {
	AbstractExpression

	text string
}

func (e *CommentExpression) LineNumber() int {
	return UnknownLineNumber
}

func (e *CommentExpression) Type() _type.IType {
	return _type.PtTypeVoid
}

func (e *CommentExpression) Priority() int {
	return 0
}

func (e *CommentExpression) GetText() string {
	return e.text
}

func (e *CommentExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitCommentExpression(e)
}

func (e *CommentExpression) String() string {
	return fmt.Sprintf("CommentExpression{%s}", e.text)
}
