package token

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
)

const (
	CommentToken          = 1
	JavaDocToken          = 2
	ErrorToken            = 3
	ImportStatementsToken = 4
)

var EndComment = NewEndMarkerToken(CommentToken)
var EndJavaDoc = NewEndMarkerToken(JavaDocToken)
var EndError = NewEndMarkerToken(ErrorToken)
var EndImportStatements = NewEndMarkerToken(ImportStatementsToken)

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
