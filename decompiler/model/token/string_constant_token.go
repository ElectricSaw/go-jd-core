package token

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewStringConstantToken(text string, ownerInternalName string) intmod.IStringConstantToken {
	return &StringConstantToken{text, ownerInternalName}
}

type StringConstantToken struct {
	text              string
	ownerInternalName string
}

func (t *StringConstantToken) Text() string {
	return t.text
}

func (t *StringConstantToken) SetText(text string) {
	t.text = text
}

func (t *StringConstantToken) OwnerInternalName() string {
	return t.ownerInternalName
}

func (t *StringConstantToken) SetOwnerInternalName(ownerInternalName string) {
	t.ownerInternalName = ownerInternalName
}

func (t *StringConstantToken) Accept(visitor intmod.ITokenVisitor) {
	visitor.VisitStringConstantToken(t)
}

func (t *StringConstantToken) String() string {
	return fmt.Sprintf("StringConstantToken { '%s' }", t.text)
}
