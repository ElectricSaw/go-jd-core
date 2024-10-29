package attribute

func NewLineNumber(startPc int, lineNumber int) LineNumber {
	return LineNumber{startPc, lineNumber}
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
