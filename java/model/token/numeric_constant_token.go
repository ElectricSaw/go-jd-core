package token

import "fmt"

func NewNumericConstantToken(text string) NumericConstantToken {
	return NumericConstantToken{text}
}

type NumericConstantToken struct {
	text string
}

func (t *NumericConstantToken) Text() string {
	return t.text
}

func (t *NumericConstantToken) Accept(visitor TokenVisitor) {
	visitor.VisitNumericConstantToken(t)
}

func (t *NumericConstantToken) String() string {
	return fmt.Sprintf("NumericConstantToken { '%s' }", t.text)
}
