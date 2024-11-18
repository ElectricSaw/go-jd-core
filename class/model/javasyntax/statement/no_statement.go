package statement

import intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"

var NoStmt = NewNoStatement()

func NewNoStatement() intsyn.INoStatement {
	return new(NoStatement)
}

type NoStatement struct {
	AbstractStatement
}

func (s *NoStatement) Accept(visitor intsyn.IStatementVisitor) {
	visitor.VisitNoStatement(s)
}

func (s *NoStatement) String() string {
	return "NoStatement"
}
