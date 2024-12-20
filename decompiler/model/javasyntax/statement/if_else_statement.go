package statement

import intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"

func NewIfElseStatement(condition intmod.IExpression, statements intmod.IStatement, elseStatements intmod.IStatement) intmod.IIfElseStatement {
	return &IfElseStatement{
		IfStatement: IfStatement{
			condition:  condition,
			statements: statements,
		},
		elseStatements: elseStatements,
	}
}

type IfElseStatement struct {
	IfStatement

	elseStatements intmod.IStatement
}

func (s *IfElseStatement) ElseStatements() intmod.IStatement {
	return s.elseStatements
}

func (s *IfElseStatement) IsIfElseStatement() bool {
	return true
}

func (s *IfElseStatement) AcceptStatement(visitor intmod.IStatementVisitor) {
	visitor.VisitIfElseStatement(s)
}
