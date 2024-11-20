package javafragment

import (
	"bitbucket.org/coontec/go-jd-core/class/model/fragment"
	"bitbucket.org/coontec/go-jd-core/class/model/token"
)

var Comma = NewTokensFragment(token.Comma)

func NewTokensFragment(tokens ...token.Token) *TokensFragment {
	return NewTokensFragmentWithSlice(tokens)
}

func NewTokensFragmentWithSlice(tokens []token.Token) *TokensFragment {
	return newTokensFragment(getLineCount(tokens), tokens)
}

func newTokensFragment(lineCount int, tokens []token.Token) *TokensFragment {
	return &TokensFragment{
		FlexibleFragment: *fragment.NewFlexibleFragment(lineCount, lineCount, lineCount, 0, "Tokens"),
		tokens:           tokens,
	}
}

func getLineCount(tokens []token.Token) int {
	visitor := &LineCountVisitor{}

	for _, tkn := range tokens {
		tkn.Accept(visitor)
	}

	return visitor.LineCount()
}

type TokensFragment struct {
	fragment.FlexibleFragment

	tokens []token.Token
}

func (f *TokensFragment) Tokens() []token.Token {
	return f.tokens
}

func (f *TokensFragment) Accept(visitor JavaFragmentVisitor) {
	visitor.VisitTokensFragment(f)
}

type LineCountVisitor struct {
	token.AbstractNopTokenVisitor

	lineCount int
}

func (v *LineCountVisitor) LineCount() int {
	return v.lineCount
}

func (v *LineCountVisitor) VisitLineNumberToken(token *token.LineNumberToken) {
	v.lineCount++
}
