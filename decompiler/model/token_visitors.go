package model

type AbstractNopTokenVisitor struct {
}

func (v *AbstractNopTokenVisitor) VisitBooleanConstantToken(_ *BooleanConstantToken) {}

func (v *AbstractNopTokenVisitor) VisitCharacterConstantToken(_ *CharacterConstantToken) {}

func (v *AbstractNopTokenVisitor) VisitDeclarationToken(_ *DeclarationToken) {}

func (v *AbstractNopTokenVisitor) VisitEndBlockToken(_ *EndBlockToken) {}

func (v *AbstractNopTokenVisitor) VisitEndMarkerToken(_ *EndMarkerToken) {}

func (v *AbstractNopTokenVisitor) VisitKeywordToken(_ *KeywordToken) {}

func (v *AbstractNopTokenVisitor) VisitLineNumberToken(_ *LineNumberToken) {}

func (v *AbstractNopTokenVisitor) VisitNewLineToken(_ *NewLineToken) {}

func (v *AbstractNopTokenVisitor) VisitNumericConstantToken(_ *NumericConstantToken) {}

func (v *AbstractNopTokenVisitor) VisitReferenceToken(_ *ReferenceToken) {}

func (v *AbstractNopTokenVisitor) VisitStartBlockToken(_ *StartBlockToken) {}

func (v *AbstractNopTokenVisitor) VisitStartMarkerToken(_ *StartMarkerToken) {}

func (v *AbstractNopTokenVisitor) VisitStringConstantToken(_ *StringConstantToken) {}

func (v *AbstractNopTokenVisitor) VisitTextToken(_ *TextToken) {}
