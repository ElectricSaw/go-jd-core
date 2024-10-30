package token

import "fmt"

func NewBooleanConstantToken(value bool) BooleanConstantToken {
	return BooleanConstantToken{value}
}

type BooleanConstantToken struct {
	value bool
}

func (t *BooleanConstantToken) Value() bool {
	return t.value
}

func (t *BooleanConstantToken) Accept(visitor TokenVisitor) {
	visitor.VisitBooleanConstantToken(t)
}

func (t *BooleanConstantToken) String() string {
	value := "false"

	if t.Value() {
		value = "true"
	}

	return fmt.Sprintf("BooleanConstantToken { '%s' }", value)
}
