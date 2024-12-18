package token

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewEndMarkerToken(typ int) intmod.IEndMarkerToken {
	return &EndMarkerToken{typ}
}

type EndMarkerToken struct {
	typ int
}

func (t *EndMarkerToken) Type() int {
	return t.typ
}

func (t *EndMarkerToken) SetType(typ int) {
	t.typ = typ
}

func (t *EndMarkerToken) Accept(visitor intmod.ITokenVisitor) {
	visitor.VisitEndMarkerToken(t)
}

func (t *EndMarkerToken) String() string {
	return fmt.Sprintf("EndMarkerToken { '%d' }", t.typ)
}
