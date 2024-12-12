package token

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
)

func NewBooleanConstantToken(value bool) intmod.IBooleanConstantToken {
	return &BooleanConstantToken{value}
}

type BooleanConstantToken struct {
	value bool
}

func (t *BooleanConstantToken) Value() bool {
	return t.value
}

func (t *BooleanConstantToken) SetValue(value bool) {
	t.value = value
}

func (t *BooleanConstantToken) Accept(visitor intmod.ITokenVisitor) {
	visitor.VisitBooleanConstantToken(t)
}

func (t *BooleanConstantToken) String() string {
	value := "false"

	if t.Value() {
		value = "true"
	}

	return fmt.Sprintf("BooleanConstantToken { '%s' }", value)
}
