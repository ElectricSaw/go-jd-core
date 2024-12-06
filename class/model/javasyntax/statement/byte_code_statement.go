package statement

import intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"

func NewByteCodeStatement(text string) intmod.IByteCodeStatement {
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

func (s *ByteCodeStatement) AcceptStatement(visitor intmod.IStatementVisitor) {
	visitor.VisitByteCodeStatement(s)
}
