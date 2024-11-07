package token

import "fmt"

var EndBlock = NewEndBlockToken("}")
var EndArrayBlock = NewEndBlockToken("]")
var EndArrayInitializerBlock = NewEndBlockToken("}")
var EndParametersBlock = NewEndBlockToken(")")
var EndResourcesBlock = NewEndBlockToken(")")
var EndDeclarationOrStatementBlock = NewEndBlockToken("")

func NewEndBlockToken(text string) *EndBlockToken {
	return &EndBlockToken{text}
}

type EndBlockToken struct {
	text string
}

func (t *EndBlockToken) Text() string {
	return t.text
}

func (t *EndBlockToken) Accept(visitor TokenVisitor) {
	visitor.VisitEndBlockToken(t)
}

func (t *EndBlockToken) String() string {
	return fmt.Sprintf("EndBlockToken { '%s' }", t.text)
}
