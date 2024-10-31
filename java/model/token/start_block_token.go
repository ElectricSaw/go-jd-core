package token

import "fmt"

var StartBlock = NewStartBlockToken("{")
var StartArrayBlock = NewStartBlockToken("[")
var StartArrayInitializerBlock = NewStartBlockToken("{")
var StartParametersBlock = NewStartBlockToken("(")
var StartResourcesBlock = NewStartBlockToken("(")
var StartDeclarationOrStatementBlock = NewStartBlockToken("")

func NewStartBlockToken(text string) *StartBlockToken {
	return &StartBlockToken{text}
}

type StartBlockToken struct {
	text string
}

func (t *StartBlockToken) Text() string {
	return t.text
}

func (t *StartBlockToken) Accept(visitor TokenVisitor) {
	visitor.VisitStartBlockToken(t)
}

func (t *StartBlockToken) String() string {
	return fmt.Sprintf("StartBlockToken { '%s' }", t.text)
}
