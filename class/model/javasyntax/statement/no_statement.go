package statement

import intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"

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
