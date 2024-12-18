package attribute

import intcls "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/classpath"

func NewLineNumber(startPc int, lineNumber int) intcls.ILineNumber {
	return &LineNumber{startPc, lineNumber}
}

type LineNumber struct {
	startPc    int
	lineNumber int
}

func (l LineNumber) StartPc() int {
	return l.startPc
}

func (l LineNumber) LineNumber() int {
	return l.lineNumber
}
