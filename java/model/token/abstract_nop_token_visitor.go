package token

type AbstractNopTokenVisitor struct {
}

func (v *AbstractNopTokenVisitor) VisitBooleanConstantToken(token *BooleanConstantToken) {

}

func (v *AbstractNopTokenVisitor) VisitCharacterConstantToken(token *CharacterConstantToken) {

}

func (v *AbstractNopTokenVisitor) VisitDeclarationToken(token *DeclarationToken) {

}

func (v *AbstractNopTokenVisitor) VisitEndBlockToken(token *EndBlockToken) {

}

func (v *AbstractNopTokenVisitor) VisitEndMarkerToken(token *EndMarkerToken) {

}

func (v *AbstractNopTokenVisitor) VisitKeywordToken(token *KeywordToken) {

}

func (v *AbstractNopTokenVisitor) VisitLineNumberToken(token *LineNumberToken) {

}

func (v *AbstractNopTokenVisitor) VisitNewLineToken(token *NewLineToken) {

}

func (v *AbstractNopTokenVisitor) VisitNumericConstantToken(token *NumericConstantToken) {

}

func (v *AbstractNopTokenVisitor) VisitReferenceToken(token *ReferenceToken) {

}

func (v *AbstractNopTokenVisitor) VisitStartBlockToken(token *StartBlockToken) {

}

func (v *AbstractNopTokenVisitor) VisitStartMarkerToken(token *StartMarkerToken) {

}

func (v *AbstractNopTokenVisitor) VisitStringConstantToken(token *StringConstantToken) {

}

func (v *AbstractNopTokenVisitor) VisitTextToken(token *TextToken) {

}
