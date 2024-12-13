package token

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
)

func NewEndBlockToken(text string) intmod.IEndBlockToken {
	return &EndBlockToken{text}
}

type EndBlockToken struct {
	text string
}

func (t *EndBlockToken) Text() string {
	return t.text
}

func (t *EndBlockToken) SetText(text string) {
	t.text = text
}

func (t *EndBlockToken) Accept(visitor intmod.ITokenVisitor) {
	visitor.VisitEndBlockToken(t)
}

func (t *EndBlockToken) String() string {
	return fmt.Sprintf("EndBlockToken { '%s' }", t.text)
}
