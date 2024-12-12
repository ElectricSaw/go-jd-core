package token

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
)

var StartComment = NewStartMarkerToken(CommentToken)
var StartJavaDocToken = NewStartMarkerToken(JavaDocToken)
var StartErrorToken = NewStartMarkerToken(ErrorToken)
var StartImportStatementsToken = NewStartMarkerToken(ImportStatementsToken)

func NewStartMarkerToken(typ int) intmod.IStartMarkerToken {
	return &StartMarkerToken{typ}
}

type StartMarkerToken struct {
	typ int
}

func (t *StartMarkerToken) Type() int {
	return t.typ
}

func (t *StartMarkerToken) SetType(typ int) {
	t.typ = typ
}

func (t *StartMarkerToken) Accept(visitor intmod.ITokenVisitor) {
	visitor.VisitStartMarkerToken(t)
}

func (t *StartMarkerToken) String() string {
	return fmt.Sprintf("StartMarkerToken { '%d' }", t.typ)
}
