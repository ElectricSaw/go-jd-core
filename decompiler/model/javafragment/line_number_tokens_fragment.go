package javafragment

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/token"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewLineNumberTokensFragment(tokens ...intmod.IToken) LineNumberTokensFragment {
	frag := LineNumberTokensFragment{
		Tokens: util.NewDefaultListWithElements[intmod.IToken](tokens...),
	}
	return frag
}

func SearchFirstLineNumber(tokens util.IList[intmod.IToken]) int {
	visitor := NewSearchLineNumberVisitor()

	for _, tkn := range tokens.ToSlice() {
		tkn.Accept(&visitor)

		if visitor.LineNumber != intmod.UnknownLineNumberToken {
			return visitor.LineNumber - visitor.NewLineCounter
		}
	}

	return intmod.UnknownLineNumberToken
}

func searchLastLineNumber(tokens util.IList[intmod.IToken]) int {
	visitor := NewSearchLineNumberVisitor()
	index := tokens.Size()

	for index > 0 {
		index--
		tokens.Get(index).Accept(&visitor)

		if visitor.LineNumber != intmod.UnknownLineNumberToken {
			return visitor.LineNumber + visitor.NewLineCounter
		}
	}

	return intmod.UnknownLineNumberToken
}

type LineNumberTokensFragment struct {
	FixedFragment

	Tokens util.IList[intmod.IToken]
}

func (f *LineNumberTokensFragment) TokenAt(index int) intmod.IToken {
	return f.Tokens.Get(index)
}

func (f *LineNumberTokensFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitLineNumberTokensFragment(f)
}

func NewSearchLineNumberVisitor() SearchLineNumberVisitor {
	return SearchLineNumberVisitor{}
}

type SearchLineNumberVisitor struct {
	token.AbstractNopTokenVisitor

	LineNumber     int
	NewLineCounter int
}

func (v *SearchLineNumberVisitor) Reset() {
	v.LineNumber = intmod.UnknownLineNumberToken
	v.NewLineCounter = 0
}

func (v *SearchLineNumberVisitor) VisitBooleanConstantToken(_ intmod.IBooleanConstantToken) {}

func (v *SearchLineNumberVisitor) VisitCharacterConstantToken(_ intmod.ICharacterConstantToken) {}

func (v *SearchLineNumberVisitor) VisitDeclarationToken(_ intmod.IDeclarationToken) {}

func (v *SearchLineNumberVisitor) VisitEndBlockToken(_ intmod.IEndBlockToken) {}

func (v *SearchLineNumberVisitor) VisitEndMarkerToken(_ intmod.IEndMarkerToken) {}

func (v *SearchLineNumberVisitor) VisitKeywordToken(_ intmod.IKeywordToken) {}

func (v *SearchLineNumberVisitor) VisitLineNumberToken(token intmod.ILineNumberToken) {
	v.LineNumber = token.LineNumber()
}

func (v *SearchLineNumberVisitor) VisitNewLineToken(_ intmod.INewLineToken) {
	v.NewLineCounter++
}

func (v *SearchLineNumberVisitor) VisitNumericConstantToken(_ intmod.INumericConstantToken) {}

func (v *SearchLineNumberVisitor) VisitReferenceToken(_ intmod.IReferenceToken) {}

func (v *SearchLineNumberVisitor) VisitStartBlockToken(_ intmod.IStartBlockToken) {}

func (v *SearchLineNumberVisitor) VisitStartMarkerToken(_ intmod.IStartMarkerToken) {}

func (v *SearchLineNumberVisitor) VisitStringConstantToken(_ intmod.IStringConstantToken) {}

func (v *SearchLineNumberVisitor) VisitTextToken(_ intmod.ITextToken) {}
