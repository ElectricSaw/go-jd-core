package token

import "fmt"

const (
	UnknownLineNumberToken = 0
)

var UnknownLineNumber = NewLineNumberToken(UnknownLineNumberToken)

func NewLineNumberToken(lineNumber int) *LineNumberToken {
	return &LineNumberToken{lineNumber}
}

type LineNumberToken struct {
	lineNumber int
}

func (t *LineNumberToken) LineNumber() int {
	return t.lineNumber
}

func (t *LineNumberToken) Accept(visitor TokenVisitor) {
	visitor.VisitLineNumberToken(t)
}

func (t *LineNumberToken) String() string {
	return fmt.Sprintf("LineNumberToken { '%d' }", t.lineNumber)
}
