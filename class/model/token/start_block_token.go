package token

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"fmt"
)

var StartBlock = NewStartBlockToken("{")
var StartArrayBlock = NewStartBlockToken("[")
var StartArrayInitializerBlock = NewStartBlockToken("{")
var StartParametersBlock = NewStartBlockToken("(")
var StartResourcesBlock = NewStartBlockToken("(")
var StartDeclarationOrStatementBlock = NewStartBlockToken("")

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
