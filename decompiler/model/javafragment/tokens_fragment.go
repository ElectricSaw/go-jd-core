package javafragment

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/fragment"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/token"
)

func NewTokensFragment(tokens ...intmod.IToken) intmod.ITokensFragment {
	return NewTokensFragmentWithSlice(tokens)
}

func NewTokensFragmentWithSlice(tokens []intmod.IToken) intmod.ITokensFragment {
	return newTokensFragment(getLineCount(tokens), tokens)
}

func newTokensFragment(lineCount int, tokens []intmod.IToken) intmod.ITokensFragment {
	return &TokensFragment{
		FlexibleFragment: *fragment.NewFlexibleFragment(lineCount, lineCount, lineCount,
			0, "Tokens").(*fragment.FlexibleFragment),
		tokens: tokens,
	}
}

type TokensFragment struct {
	fragment.FlexibleFragment

	tokens []intmod.IToken
}

func (f *TokensFragment) TokenAt(index int) intmod.IToken {
	return f.tokens[index]
}

func (f *TokensFragment) Tokens() []intmod.IToken {
	return f.tokens
}

func (f *TokensFragment) Accept(visitor intmod.IJavaFragmentVisitor) {
	visitor.VisitTokensFragment(f)
}

func NewLineCountVisitor() intmod.ILineCountVisitor {
	return &LineCountVisitor{
		lineCount: 0,
	}
}

type LineCountVisitor struct {
	token.AbstractNopTokenVisitor

	lineCount int
}

func (v *LineCountVisitor) LineCount() int {
	return v.lineCount
}

func (v *LineCountVisitor) SetLineCount(lineCount int) {
	v.lineCount = lineCount
}

func (v *LineCountVisitor) VisitLineNumberToken(_ intmod.ILineNumberToken) {
	v.lineCount++
}

func getLineCount(tokens []intmod.IToken) int {
	visitor := NewLineCountVisitor()

	for _, tkn := range tokens {
		tkn.Accept(visitor)
	}

	return visitor.LineCount()
}
