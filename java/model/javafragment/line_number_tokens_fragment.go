package javafragment

import (
	"bitbucket.org/coontec/javaClass/java/model/fragment"
	"bitbucket.org/coontec/javaClass/java/model/token"
)

func NewLineNumberTokensFragment(tokens []token.Token) *LineNumberTokensFragment {
	return &LineNumberTokensFragment{
		tokens: tokens,
	}
}

func SearchFirstLineNumber(tokens []token.Token) int {
	visitor := new(SearchLineNumberVisitor)

	for _, tkn := range tokens {
		tkn.Accept(visitor)

		if visitor.LineNumber != token.UnknownLineNumberToken {
			return visitor.LineNumber - visitor.NewLineCounter
		}
	}

	return token.UnknownLineNumberToken
}

type LineNumberTokensFragment struct {
	fragment.FixedFragment

	tokens []token.Token
}

func (f *LineNumberTokensFragment) Tokens() []token.Token {
	return f.tokens
}

func (f *LineNumberTokensFragment) Accept(visitor JavaFragmentVisitor) {
	visitor.VisitLineNumberTokensFragment(f)
}

type SearchLineNumberVisitor struct {
	token.AbstractNopTokenVisitor

	LineNumber     int
	NewLineCounter int
}

func (v *SearchLineNumberVisitor) Reset() {
	v.LineNumber = token.UnknownLineNumberToken
	v.NewLineCounter = 0
}

func (v *SearchLineNumberVisitor) VisitLineNumberToken(token *token.LineNumberToken) {
	v.LineNumber = token.LineNumber()
}

func (v *SearchLineNumberVisitor) VisitNewLineToken(token *token.NewLineToken) {
	v.NewLineCounter++
}
