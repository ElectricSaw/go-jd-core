package model

import "fmt"

const (
	UnknownLineNumberToken = 0
)

const (
	TypeToken        = 1
	FieldToken       = 2
	MethodToken      = 3
	ConstructorToken = 4
	PackageToken     = 5
	ModuleToken      = 6
)

type IToken interface {
	Accept(visitor ITokenVisitor)
	String() string
}

type ITokenVisitor interface {
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

func NewBooleanConstantToken(value bool) BooleanConstantToken {
	return BooleanConstantToken{
		Value: value,
	}
}

func NewCharacterConstantToken(character string, ownerInternalName string) CharacterConstantToken {
	return CharacterConstantToken{
		Character:         character,
		OwnerInternalName: ownerInternalName,
	}
}

func NewDeclarationToken(typ int, internalTypeName, name, descriptor string) DeclarationToken {
	return DeclarationToken{
		Type:             typ,
		InternalTypeName: internalTypeName,
		Name:             name,
		Descriptor:       descriptor,
	}
}

func NewEndBlockToken(text string) EndBlockToken {
	return EndBlockToken{
		Text: text,
	}
}

func NewEndMarkerToken(typ int) EndMarkerToken {
	return EndMarkerToken{
		Type: typ,
	}
}

func NewKeywordToken(keyword string) KeywordToken {
	return KeywordToken{
		Keyword: keyword,
	}
}

func NewLineNumberToken(lineNumber int) LineNumberToken {
	return LineNumberToken{
		LineNumber: lineNumber,
	}
}

func NewNewLineToken(count int) NewLineToken {
	return NewLineToken{
		Count: count,
	}
}

func NewNumericConstantToken(text string) NumericConstantToken {
	return NumericConstantToken{
		Text: text,
	}
}

func NewReferenceToken(typ int, internalTypeName, name, descriptor, ownerInternalName string) ReferenceToken {
	return ReferenceToken{
		Type:              typ,
		InternalTypeName:  internalTypeName,
		Name:              name,
		Descriptor:        descriptor,
		OwnerInternalName: ownerInternalName,
	}
}

func NewStartBlockToken(text string) StartBlockToken {
	return StartBlockToken{
		Text: text,
	}
}

func NewStartMarkerToken(typ int) StartMarkerToken {
	return StartMarkerToken{
		Type: typ,
	}
}

func NewStringConstantToken(text string, ownerInternalName string) StringConstantToken {
	return StringConstantToken{
		Text:              text,
		OwnerInternalName: ownerInternalName,
	}
}

func NewTextToken(text string) TextToken {
	return TextToken{
		Text: text,
	}
}

type BooleanConstantToken struct {
	Value bool
}

func (t *BooleanConstantToken) Accept(visitor ITokenVisitor) {
	visitor.VisitBooleanConstantToken(t)
}

func (t *BooleanConstantToken) String() string {
	value := "false"

	if t.Value {
		value = "true"
	}

	return fmt.Sprintf("BooleanConstantToken { '%s' }", value)
}

type CharacterConstantToken struct {
	Character         string
	OwnerInternalName string
}

func (t *CharacterConstantToken) Accept(visitor ITokenVisitor) {
	visitor.VisitCharacterConstantToken(t)
}

func (t *CharacterConstantToken) String() string {
	return fmt.Sprintf("CharacterConstantToken { '%s' }", t.Character)
}

type DeclarationToken struct {
	Type             int
	InternalTypeName string
	Name             string
	Descriptor       string
}

func (t *DeclarationToken) Accept(visitor ITokenVisitor) {
	visitor.VisitDeclarationToken(t)
}

func (t *DeclarationToken) String() string {
	return fmt.Sprintf("DeclarationToken { declaration='%s' }", t.Name)
}

type EndBlockToken struct {
	Text string
}

func (t *EndBlockToken) Accept(visitor ITokenVisitor) {
	visitor.VisitEndBlockToken(t)
}

func (t *EndBlockToken) String() string {
	return fmt.Sprintf("EndBlockToken { '%s' }", t.Text)
}

type EndMarkerToken struct {
	Type int
}

func (t *EndMarkerToken) Accept(visitor ITokenVisitor) {
	visitor.VisitEndMarkerToken(t)
}

func (t *EndMarkerToken) String() string {
	return fmt.Sprintf("EndMarkerToken { '%d' }", t.Type)
}

type KeywordToken struct {
	Keyword string
}

func (t *KeywordToken) Accept(visitor ITokenVisitor) {
	visitor.VisitKeywordToken(t)
}

func (t *KeywordToken) String() string {
	return fmt.Sprintf("KeywordToken { '%s' }", t.Keyword)
}

type LineNumberToken struct {
	LineNumber int
}

func (t *LineNumberToken) Accept(visitor ITokenVisitor) {
	visitor.VisitLineNumberToken(t)
}

func (t *LineNumberToken) String() string {
	return fmt.Sprintf("LineNumberToken { '%d' }", t.LineNumber)
}

type NewLineToken struct {
	Count int
}

func (t *NewLineToken) Accept(visitor ITokenVisitor) {
	visitor.VisitNewLineToken(t)
}

func (t *NewLineToken) String() string {
	return fmt.Sprintf("NewLineToken { '%d' }", t.Count)
}

type NumericConstantToken struct {
	Text string
}

func (t *NumericConstantToken) Accept(visitor ITokenVisitor) {
	visitor.VisitNumericConstantToken(t)
}

func (t *NumericConstantToken) String() string {
	return fmt.Sprintf("NumericConstantToken { '%s' }", t.Text)
}

type ReferenceToken struct {
	//DeclarationToken
	Type              int
	InternalTypeName  string
	Name              string
	Descriptor        string
	OwnerInternalName string
}

func (t *ReferenceToken) Accept(visitor ITokenVisitor) {
	visitor.VisitReferenceToken(t)
}

func (t *ReferenceToken) String() string {
	return fmt.Sprintf("ReferenceToken { '%s' }", t.OwnerInternalName)
}

type StartBlockToken struct {
	Text string
}

func (t *StartBlockToken) Accept(visitor ITokenVisitor) {
	visitor.VisitStartBlockToken(t)
}

func (t *StartBlockToken) String() string {
	return fmt.Sprintf("StartBlockToken { '%s' }", t.Text)
}

type StartMarkerToken struct {
	Type int
}

func (t *StartMarkerToken) Accept(visitor ITokenVisitor) {
	visitor.VisitStartMarkerToken(t)
}

func (t *StartMarkerToken) String() string {
	return fmt.Sprintf("StartMarkerToken { '%d' }", t.Type)
}

type StringConstantToken struct {
	Text              string
	OwnerInternalName string
}

func (t *StringConstantToken) Accept(visitor ITokenVisitor) {
	visitor.VisitStringConstantToken(t)
}

func (t *StringConstantToken) String() string {
	return fmt.Sprintf("StringConstantToken { '%s' }", t.Text)
}

type TextToken struct {
	Text string
}

func (t *TextToken) Accept(visitor ITokenVisitor) {
	visitor.VisitTextToken(t)
}

func (t *TextToken) String() string {
	return fmt.Sprintf("TextToken { '%s' }", t.Text)
}
