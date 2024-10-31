package token

import "fmt"

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

func NewEndMarkerToken(typ int) *EndMarkerToken {
	return &EndMarkerToken{typ}
}

type EndMarkerToken struct {
	typ int
}

func (t *EndMarkerToken) Type() int {
	return t.typ
}

func (t *EndMarkerToken) Accept(visitor TokenVisitor) {
	visitor.VisitEndMarkerToken(t)
}

func (t *EndMarkerToken) String() string {
	return fmt.Sprintf("EndMarkerToken { '%d' }", t.typ)
}
