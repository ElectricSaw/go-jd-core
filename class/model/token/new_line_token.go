package token

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"fmt"
)

var NewLine1 = NewNewLineToken(1)
var NewLine2 = NewNewLineToken(2)

func NewNewLineToken(count int) intmod.INewLineToken {
	return &NewLineToken{count}
}

type NewLineToken struct {
	count int
}

func (t *NewLineToken) Count() int {
	return t.count
}

func (t *NewLineToken) SetCount(count int) {
	t.count = count
}

func (t *NewLineToken) Accept(visitor intmod.ITokenVisitor) {
	visitor.VisitNewLineToken(t)
}

func (t *NewLineToken) String() string {
	return fmt.Sprintf("NewLineToken { '%d' }", t.count)
}
