package token

import "fmt"

func NewKeywordToken(keyword string) *KeywordToken {
	return &KeywordToken{keyword}
}

type KeywordToken struct {
	keyword string
}

func (t *KeywordToken) Keyword() string {
	return t.keyword
}

func (t *KeywordToken) Accept(visitor TokenVisitor) {
	visitor.VisitKeywordToken(t)
}

func (t *KeywordToken) String() string {
	return fmt.Sprintf("KeywordToken { '%s' }", t.keyword)
}
