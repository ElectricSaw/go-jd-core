package token

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
)

func NewTextToken(text string) intmod.ITextToken {
	return &TextToken{text}
}

type TextToken struct {
	text string
}

func (t *TextToken) Text() string {
	return t.text
}

func (t *TextToken) SetText(text string) {
	t.text = text
}

func (t *TextToken) Accept(visitor intmod.ITokenVisitor) {
	visitor.VisitTextToken(t)
}

func (t *TextToken) String() string {
	return fmt.Sprintf("TextToken { '%s' }", t.text)
}
