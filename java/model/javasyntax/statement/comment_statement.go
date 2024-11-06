package statement

func NewCommentStatement(text string) *CommentStatement {
	return &CommentStatement{
		text: text,
	}
}

type CommentStatement struct {
	AbstractStatement

	text string
}

func (s *CommentStatement) GetLabel() string {
	return s.text
}

func (s *CommentStatement) IsContinueStatement() bool {
	return true
}

func (s *CommentStatement) Accept(visitor StatementVisitor) {
	visitor.VisitCommentStatement(s)
}
