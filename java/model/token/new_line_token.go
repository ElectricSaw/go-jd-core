package token

import "fmt"

func NewNewLineToken(count int) NewLineToken {
	return NewLineToken{count}
}

type NewLineToken struct {
	count int
}

func (t *NewLineToken) Count() int {
	return t.count
}

func (t *NewLineToken) Accept(visitor TokenVisitor) {
	visitor.VisitNewLineToken(t)
}

func (t *NewLineToken) String() string {
	return fmt.Sprintf("NewLineToken { '%d' }", t.count)
}
