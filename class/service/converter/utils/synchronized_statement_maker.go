package utils

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
	modsts "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/statement"
)

func MakeSynchronizedStatementMaker(localVariableMaker intsrv.ILocalVariableMaker, statements intmod.IStatements, tryStatements intmod.IStatements) intmod.IStatement {
	// Remove monitor enter
	monitor := statements.RemoveLast().Monitor()
	var localVariable intsrv.ILocalVariable

	if monitor.IsLocalVariableReferenceExpression() {
		if !statements.IsEmpty() {
			expression := statements.RemoveLast().Expression()

			if expression.IsBinaryOperatorExpression() && expression.LeftExpression().IsLocalVariableReferenceExpression() {
				// m := monitor.(intsrv.IClassFileLocalVariableReferenceExpression)
				l := expression.LeftExpression().(intsrv.IClassFileLocalVariableReferenceExpression)
				// assert l.LocalVariable() == m.LocalVariable();
				// Update monitor
				monitor = expression.RightExpression()
				// Store synthetic local variable
				localVariable = l.LocalVariable().(intsrv.ILocalVariable)
			}
		}
	} else if monitor.IsBinaryOperatorExpression() {
		if monitor.LeftExpression().IsLocalVariableReferenceExpression() {
			l := monitor.LeftExpression().(intsrv.IClassFileLocalVariableReferenceExpression)
			// Update monitor
			monitor = monitor.RightExpression()
			// Store synthetic local variable
			localVariable = l.LocalVariable().(intsrv.ILocalVariable)
		}
	}

	NewRemoveMonitorExitVisitor(localVariable).VisitStatements(tryStatements)

	// Remove synthetic local variable
	localVariableMaker.RemoveLocalVariable(localVariable)

	return modsts.NewSynchronizedStatement(monitor, tryStatements)
}

func NewRemoveMonitorExitVisitor(localVariable intsrv.ILocalVariable) *RemoveMonitorExitVisitor {
	return &RemoveMonitorExitVisitor{
		localVariable: localVariable,
	}
}

type RemoveMonitorExitVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	localVariable intsrv.ILocalVariable
}

func (v *RemoveMonitorExitVisitor) VisitStatements(list intmod.IStatements) {
	if !list.IsEmpty() {
		iterator := list.Iterator()

		for iterator.HasNext() {
			statement := iterator.Next()

			if statement.IsMonitorExitStatement() {
				if statement.Monitor().IsLocalVariableReferenceExpression() {
					cflvre := statement.Monitor().(intsrv.IClassFileLocalVariableReferenceExpression)
					if cflvre.LocalVariable() == v.localVariable {
						_ = iterator.Remove()
					}
				}
			} else {
				statement.Accept(v)
			}
		}
	}
}
