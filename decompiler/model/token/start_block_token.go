package token

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewStartBlockToken(text string) intmod.IStartBlockToken {
	return &StartBlockToken{text}
}

type StartBlockToken struct {
	text string
}

func (t *StartBlockToken) Text() string {
	return t.text
}

func (t *StartBlockToken) SetText(text string) {
	t.text = text
}

func (t *StartBlockToken) Accept(visitor intmod.ITokenVisitor) {
	visitor.VisitStartBlockToken(t)
}

func (t *StartBlockToken) String() string {
	return fmt.Sprintf("StartBlockToken { '%s' }", t.text)
}
