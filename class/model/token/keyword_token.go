package token

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewKeywordToken(keyword string) intmod.IKeywordToken {
	return &KeywordToken{keyword}
}

type KeywordToken struct {
	keyword string
}

func (t *KeywordToken) Keyword() string {
	return t.keyword
}

func (t *KeywordToken) SetKeyword(keyword string) {
	t.keyword = keyword
}

func (t *KeywordToken) Accept(visitor intmod.ITokenVisitor) {
	visitor.VisitKeywordToken(t)
}

func (t *KeywordToken) String() string {
	return fmt.Sprintf("KeywordToken { '%s' }", t.keyword)
}
