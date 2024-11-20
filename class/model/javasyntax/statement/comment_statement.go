package statement

import intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"

func NewCommentStatement(text string) intsyn.ICommentStatement {
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

func (s *CommentStatement) Accept(visitor intsyn.IStatementVisitor) {
	visitor.VisitCommentStatement(s)
}
