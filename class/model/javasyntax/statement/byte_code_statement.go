package statement

import intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"

func NewByteCodeStatement(text string) intsyn.IByteCodeStatement {
	return &ByteCodeStatement{
		text: text,
	}
}

type ByteCodeStatement struct {
	AbstractStatement

	text string
}

func (s *ByteCodeStatement) Text() string {
	return s.text
}

func (s *ByteCodeStatement) Accept(visitor intsyn.IStatementVisitor) {
	visitor.VisitByteCodeStatement(s)
}
