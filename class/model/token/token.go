package token

type Token interface {
	Accept(visitor TokenVisitor)
}

type TokenVisitor interface {
	VisitBooleanConstantToken(token *BooleanConstantToken)
	VisitCharacterConstantToken(token *CharacterConstantToken)
	VisitDeclarationToken(token *DeclarationToken)
	VisitEndBlockToken(token *EndBlockToken)
	VisitEndMarkerToken(token *EndMarkerToken)
	VisitKeywordToken(token *KeywordToken)
	VisitLineNumberToken(token *LineNumberToken)
	VisitNewLineToken(token *NewLineToken)
	VisitNumericConstantToken(token *NumericConstantToken)
	VisitReferenceToken(token *ReferenceToken)
	VisitStartBlockToken(token *StartBlockToken)
	VisitStartMarkerToken(token *StartMarkerToken)
	VisitStringConstantToken(token *StringConstantToken)
	VisitTextToken(token *TextToken)
}
