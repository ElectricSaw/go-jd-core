package statement

func NewByteCodeStatement(text string) *ByteCodeStatement {
	return &ByteCodeStatement{
		text: text,
	}
}

type ByteCodeStatement struct {
	AbstractStatement

	text string
}

func (s *ByteCodeStatement) GetText() string {
	return s.text
}

func (s *ByteCodeStatement) Accept(visitor StatementVisitor) {
	visitor.VisitByteCodeStatement(s)
}
