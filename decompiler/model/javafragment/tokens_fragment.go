package javafragment

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/token"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewTokensFragment(tokens ...intmod.IToken) TokensFragment {
	return NewTokensFragmentWithSlice(tokens)
}

func NewTokensFragmentWithSlice(tokens []intmod.IToken) TokensFragment {
	return newTokensFragment(getLineCount(tokens), tokens)
}

func newTokensFragment(lineCount int, tokens []intmod.IToken) TokensFragment {
	return TokensFragment{
		FlexibleFragment: NewFlexibleFragment(lineCount, lineCount, lineCount,
			0, "Tokens"),
		Tokens: util.NewDefaultListWithElements[intmod.IToken](tokens...),
	}
}

type TokensFragment struct {
	FlexibleFragment

	Tokens util.IList[intmod.IToken]
}

func (f *TokensFragment) TokenAt(index int) intmod.IToken {
	return f.Tokens.Get(index)
}

func (f *TokensFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitTokensFragment(f)
}

func NewLineCountVisitor() LineCountVisitor {
	return LineCountVisitor{
		LineCount: 0,
	}
}

type LineCountVisitor struct {
	token.AbstractNopTokenVisitor

	LineCount int
}

func (v *LineCountVisitor) VisitLineNumberToken(_ intmod.ILineNumberToken) {
	v.LineCount++
}

func getLineCount(tokens []intmod.IToken) int {
	visitor := NewLineCountVisitor()

	for _, tkn := range tokens {
		tkn.Accept(&visitor)
	}

	return visitor.LineCount
}
