package token

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"fmt"
)

var EndBlock = NewEndBlockToken("}")
var EndArrayBlock = NewEndBlockToken("]")
var EndArrayInitializerBlock = NewEndBlockToken("}")
var EndParametersBlock = NewEndBlockToken(")")
var EndResourcesBlock = NewEndBlockToken(")")
var EndDeclarationOrStatementBlock = NewEndBlockToken("")

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
