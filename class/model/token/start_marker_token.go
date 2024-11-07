package token

import "fmt"

var StartComment = NewStartMarkerToken(CommentToken)
var StartJavaDocToken = NewStartMarkerToken(JavaDocToken)
var StartErrorToken = NewStartMarkerToken(ErrorToken)
var StartImportStatementsToken = NewStartMarkerToken(ImportStatementsToken)

func NewStartMarkerToken(typ int) *StartMarkerToken {
	return &StartMarkerToken{typ}
}

type StartMarkerToken struct {
	typ int
}

func (t *StartMarkerToken) Type() int {
	return t.typ
}

func (t *StartMarkerToken) Accept(visitor TokenVisitor) {
	visitor.VisitStartMarkerToken(t)
}

func (t *StartMarkerToken) String() string {
	return fmt.Sprintf("StartMarkerToken { '%d' }", t.typ)
}
