package token

import "fmt"

func NewStringConstantToken(text string, ownerInternalName string) *StringConstantToken {
	return &StringConstantToken{text, ownerInternalName}
}

type StringConstantToken struct {
	text              string
	ownerInternalName string
}

func (t *StringConstantToken) Text() string {
	return t.text
}

func (t *StringConstantToken) OwnerInternalName() string {
	return t.ownerInternalName
}

func (t *StringConstantToken) Accept(visitor TokenVisitor) {
	visitor.VisitStringConstantToken(t)
}

func (t *StringConstantToken) String() string {
	return fmt.Sprintf("StringConstantToken { '%s' }", t.text)
}
