package statement

import intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"

func NewCommentStatement(text string) intmod.ICommentStatement {
	return &CommentStatement{
		text: text,
	}
}

type CommentStatement struct {
	AbstractStatement

	text string
}

func (s *CommentStatement) Text() string {
	return s.text
}

func (s *CommentStatement) IsContinueStatement() bool {
	return true
}

func (s *CommentStatement) AcceptStatement(visitor intmod.IStatementVisitor) {
	visitor.VisitCommentStatement(s)
}
