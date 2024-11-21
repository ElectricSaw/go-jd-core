package visitor

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
)

func NewSearchUndeclaredLocalVariableVisitor() *SearchUndeclaredLocalVariableVisitor {
	return &SearchUndeclaredLocalVariableVisitor{
		variables: make([]intsrv.ILocalVariable, 0),
	}
}

type SearchUndeclaredLocalVariableVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	variables []intsrv.ILocalVariable
}

func (v *SearchUndeclaredLocalVariableVisitor) Init() {
	if v.variables == nil {
		v.variables = make([]intsrv.ILocalVariable, 0)
	}

	v.variables = v.variables[:0]
}

func (v *SearchUndeclaredLocalVariableVisitor) Variables() []intsrv.ILocalVariable {
	return v.variables
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitBinaryOperatorExpression(expression intmod.IBinaryOperatorExpression) {
	if expression.LeftExpression().IsLocalVariableReferenceExpression() && expression.Operator() == "=" {
		lv := expression.LeftExpression().(intsrv.IClassFileLocalVariableReferenceExpression).
			LocalVariable().(intsrv.ILocalVariable)

		if !lv.IsDeclared() {
			v.variables = append(v.variables, lv)
		}
	}

	expression.LeftExpression().Accept(v)
	expression.RightExpression().Accept(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitDoWhileStatement(statement intmod.IDoWhileStatement) {
	v.SafeAcceptExpression(statement.Condition())
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitForEachStatement(statement intmod.IForEachStatement) {
	statement.Expression().Accept(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitForStatement(statement intmod.IForStatement) {
	v.SafeAcceptDeclaration(statement.Declaration())
	v.SafeAcceptExpression(statement.Init())
	v.SafeAcceptExpression(statement.Condition())
	v.SafeAcceptExpression(statement.Update())
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitIfStatement(statement intmod.IIfStatement) {
	statement.Condition().Accept(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitIfElseStatement(statement intmod.IIfElseStatement) {
	statement.Condition().Accept(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitLambdaExpressionStatement(statement intmod.ILambdaExpressionStatement) {
	statement.Expression().Accept(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitSwitchStatement(statement intmod.ISwitchStatement) {
	statement.Condition().Accept(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitSynchronizedStatement(statement intmod.ISynchronizedStatement) {
	statement.Monitor().Accept(v)
}

func (v *SearchUndeclaredLocalVariableVisitor) VisitTryStatement(statement intmod.ITryStatement) {}

func (v *SearchUndeclaredLocalVariableVisitor) VisitWhileStatement(statement intmod.IWhileStatement) {
	statement.Condition().Accept(v)
}
