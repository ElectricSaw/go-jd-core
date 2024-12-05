package token

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewNumericConstantToken(text string) intmod.INumericConstantToken {
	return &NumericConstantToken{text}
}

type NumericConstantToken struct {
	text string
}

func (t *NumericConstantToken) Text() string {
	return t.text
}

func (t *NumericConstantToken) SetText(text string) {
	t.text = text
}

func (t *NumericConstantToken) Accept(visitor intmod.ITokenVisitor) {
	visitor.VisitNumericConstantToken(t)
}

func (t *NumericConstantToken) String() string {
	return fmt.Sprintf("NumericConstantToken { '%s' }", t.text)
}
