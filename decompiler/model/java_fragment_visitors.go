package model

import "fmt"

func NewSearchLineNumberVisitor() SearchLineNumberVisitor {
	return SearchLineNumberVisitor{}
}

type SearchLineNumberVisitor struct {
	LineNumber     int
	NewLineCounter int
}

func (v *SearchLineNumberVisitor) Reset() {
	v.LineNumber = UnknownLineNumberToken
	v.NewLineCounter = 0
}

func (v *SearchLineNumberVisitor) VisitBooleanConstantToken(_ *BooleanConstantToken) {}

func (v *SearchLineNumberVisitor) VisitCharacterConstantToken(_ *CharacterConstantToken) {}

func (v *SearchLineNumberVisitor) VisitDeclarationToken(_ *DeclarationToken) {}

func (v *SearchLineNumberVisitor) VisitEndBlockToken(_ *EndBlockToken) {}

func (v *SearchLineNumberVisitor) VisitEndMarkerToken(_ *EndMarkerToken) {}

func (v *SearchLineNumberVisitor) VisitKeywordToken(_ *KeywordToken) {}

func (v *SearchLineNumberVisitor) VisitLineNumberToken(token *LineNumberToken) {
	v.LineNumber = token.LineNumber
}

func (v *SearchLineNumberVisitor) VisitNewLineToken(_ *NewLineToken) {
	v.NewLineCounter++
}

func (v *SearchLineNumberVisitor) VisitNumericConstantToken(_ *NumericConstantToken) {}

func (v *SearchLineNumberVisitor) VisitReferenceToken(_ *ReferenceToken) {}

func (v *SearchLineNumberVisitor) VisitStartBlockToken(_ *StartBlockToken) {}

func (v *SearchLineNumberVisitor) VisitStartMarkerToken(_ *StartMarkerToken) {}

func (v *SearchLineNumberVisitor) VisitStringConstantToken(_ *StringConstantToken) {}

func (v *SearchLineNumberVisitor) VisitTextToken(_ *TextToken) {}

func (v *SearchLineNumberVisitor) String() string {
	return fmt.Sprintf("SearchLineNumberVisitor{lineNumber: %d, newLineCounter: %d}", v.LineNumber, v.NewLineCounter)
}

func NewLineCountVisitor() LineCountVisitor {
	return LineCountVisitor{
		LineCount: 0,
	}
}

type LineCountVisitor struct {
	LineCount int
}

func (v *LineCountVisitor) VisitBooleanConstantToken(_ *BooleanConstantToken) {}

func (v *LineCountVisitor) VisitCharacterConstantToken(_ *CharacterConstantToken) {}

func (v *LineCountVisitor) VisitDeclarationToken(_ *DeclarationToken) {}

func (v *LineCountVisitor) VisitEndBlockToken(_ *EndBlockToken) {}

func (v *LineCountVisitor) VisitEndMarkerToken(_ *EndMarkerToken) {}

func (v *LineCountVisitor) VisitKeywordToken(_ *KeywordToken) {}

func (v *LineCountVisitor) VisitLineNumberToken(_ *LineNumberToken) {
	v.LineCount++
}

func (v *LineCountVisitor) VisitNewLineToken(_ *NewLineToken) {}

func (v *LineCountVisitor) VisitNumericConstantToken(_ *NumericConstantToken) {}

func (v *LineCountVisitor) VisitReferenceToken(_ *ReferenceToken) {}

func (v *LineCountVisitor) VisitStartBlockToken(_ *StartBlockToken) {}

func (v *LineCountVisitor) VisitStartMarkerToken(_ *StartMarkerToken) {}

func (v *LineCountVisitor) VisitStringConstantToken(_ *StringConstantToken) {}

func (v *LineCountVisitor) VisitTextToken(_ *TextToken) {}

func (v *LineCountVisitor) String() string {
	return fmt.Sprintf("LineCountVisitor{ lineCounter: %d }", v.LineCount)
}
