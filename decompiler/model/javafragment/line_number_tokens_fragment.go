package javafragment

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/fragment"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/token"
)

func NewLineNumberTokensFragment(tokens []intmod.IToken) intmod.ILineNumberTokensFragment {
	return &LineNumberTokensFragment{
		tokens: tokens,
	}
}

func SearchFirstLineNumber(tokens []intmod.IToken) int {
	visitor := NewSearchLineNumberVisitor()

	for _, tkn := range tokens {
		tkn.Accept(visitor)

		if visitor.LineNumber() != intmod.UnknownLineNumberToken {
			return visitor.LineNumber() - visitor.NewLineNumber()
		}
	}

	return intmod.UnknownLineNumberToken
}

type LineNumberTokensFragment struct {
	fragment.FixedFragment

	tokens []intmod.IToken
}

func (f *LineNumberTokensFragment) TokenAt(index int) intmod.IToken {
	return f.tokens[index]
}

func (f *LineNumberTokensFragment) Tokens() []intmod.IToken {
	return f.tokens
}

func (f *LineNumberTokensFragment) Accept(visitor intmod.IJavaFragmentVisitor) {
	visitor.VisitLineNumberTokensFragment(f)
}

func NewSearchLineNumberVisitor() intmod.ISearchLineNumberVisitor {
	return &SearchLineNumberVisitor{}
}

type SearchLineNumberVisitor struct {
	token.AbstractNopTokenVisitor

	lineNumber     int
	newLineCounter int
}

func (v *SearchLineNumberVisitor) LineNumber() int {
	return v.lineNumber
}

func (v *SearchLineNumberVisitor) SetLineNumber(lineNumber int) {
	v.lineNumber = lineNumber
}

func (v *SearchLineNumberVisitor) NewLineNumber() int {
	return v.newLineCounter
}

func (v *SearchLineNumberVisitor) SetNewLineNumber(newLineCounter int) {
	v.newLineCounter = newLineCounter
}

func (v *SearchLineNumberVisitor) Reset() {
	v.lineNumber = intmod.UnknownLineNumberToken
	v.newLineCounter = 0
}

func (v *SearchLineNumberVisitor) VisitLineNumberToken(token intmod.ILineNumberToken) {
	v.lineNumber = token.LineNumber()
}

func (v *SearchLineNumberVisitor) VisitNewLineToken(_ intmod.INewLineToken) {
	v.newLineCounter++
}
