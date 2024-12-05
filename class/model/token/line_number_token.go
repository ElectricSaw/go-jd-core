package token

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

var UnknownLineNumber = NewLineNumberToken(intmod.UnknownLineNumberToken)

func NewLineNumberToken(lineNumber int) intmod.ILineNumberToken {
	return &LineNumberToken{lineNumber}
}

type LineNumberToken struct {
	lineNumber int
}

func (t *LineNumberToken) LineNumber() int {
	return t.lineNumber
}

func (t *LineNumberToken) SetLineNumber(lineNumber int) {
	t.lineNumber = lineNumber
}

func (t *LineNumberToken) Accept(visitor intmod.ITokenVisitor) {
	visitor.VisitLineNumberToken(t)
}

func (t *LineNumberToken) String() string {
	return fmt.Sprintf("LineNumberToken { '%d' }", t.lineNumber)
}
