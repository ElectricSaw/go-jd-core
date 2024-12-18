package statement

import intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"

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
