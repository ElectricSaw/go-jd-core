package visitor

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
)

func NewRemoveBinaryOpReturnStatementsVisitor(localVariableMaker intsrv.ILocalVariableMaker) *RemoveBinaryOpReturnStatementsVisitor {
	return &RemoveBinaryOpReturnStatementsVisitor{
		localVariableMaker: localVariableMaker,
	}
}

type RemoveBinaryOpReturnStatementsVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	localVariableMaker intsrv.ILocalVariableMaker
}

func (v *RemoveBinaryOpReturnStatementsVisitor) VisitStatements(state intmod.IStatements) {
	if state.Size() > 1 {
		// Replace pattern "local_var_2 = ...; return local_var_2;" with "return ...;"
		lastStatement := state.Last()

		if lastStatement.IsReturnExpressionStatement() && lastStatement.Expression().IsLocalVariableReferenceExpression() {
			lvr1 := lastStatement.Expression().(intsrv.IClassFileLocalVariableReferenceExpression)

			if lvr1.Name() == "" {
				statement := state.Get(state.Size() - 2)

				if statement.Expression().IsBinaryOperatorExpression() {
					boe := statement.Expression()
					leftExpression := boe.LeftExpression()

					if leftExpression.IsLocalVariableReferenceExpression() {
						lvr2 := leftExpression.(intsrv.IClassFileLocalVariableReferenceExpression)

						if (lvr1.LocalVariable() == lvr2.LocalVariable()) &&
							(len(lvr1.LocalVariable().(intsrv.ILocalVariable).References()) == 2) {
							res := lastStatement.(intmod.IReturnExpressionStatement)

							// Remove synthetic assignment statement
							state.RemoveAt(state.Size() - 2)
							// Replace synthetic local variable with expression
							res.SetExpression(boe.RightExpression())
							// Check line number
							expressionLineNumber := boe.RightExpression().LineNumber()
							if res.LineNumber() > expressionLineNumber {
								res.SetLineNumber(expressionLineNumber)
							}
							// Remove synthetic local variable
							v.localVariableMaker.RemoveLocalVariable(lvr1.LocalVariable().(intsrv.ILocalVariable))
						}
					}
				}
			}
		}
	}

	v.AbstractJavaSyntaxVisitor.VisitStatements(state)
}

func (v *RemoveBinaryOpReturnStatementsVisitor) VisitBodyDeclaration(_ intmod.IBodyDeclaration) {
}
