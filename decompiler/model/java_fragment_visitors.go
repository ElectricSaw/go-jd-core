package model

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/token"
)

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

func (v *SearchLineNumberVisitor) String() string {
	return fmt.Sprintf("SearchLineNumberVisitor{lineNumber: %d, newLineCounter: %d}", v.LineNumber, v.NewLineCounter)
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
