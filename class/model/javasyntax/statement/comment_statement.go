package statement

import intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"

func NewCommentStatement(text string) intmod.ICommentStatement {
	return &CommentStatement{
		text: text,
	}
}

type CommentStatement struct {
	AbstractStatement

	text string
}

func (s *CommentStatement) Label() string {
	return s.text
}

func (s *CommentStatement) IsContinueStatement() bool {
	return true
}

func (s *CommentStatement) Accept(visitor intmod.IStatementVisitor) {
	visitor.VisitCommentStatement(s)
}
