package utils

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	modexp "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/expression"
	modsts "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/statement"
	"bitbucket.org/coontec/go-jd-core/class/util"
	"strings"
)

var MinusOne = -1

func MakeSwitchString(localVariableMaker intsrv.ILocalVariableMaker, statements intmod.IStatements, switchStatement intmod.ISwitchStatement) {
	size := statements.Size()
	previousSwitchStatement := statements.Get(size - 2).(intmod.ISwitchStatement)

	if previousSwitchStatement.Condition().LineNumber() == switchStatement.Condition().LineNumber() &&
		previousSwitchStatement.Condition().IsMethodInvocationExpression() {
		expression := previousSwitchStatement.Condition()

		if expression.IsMethodInvocationExpression() {
			expression = expression.Expression()

			if expression.IsLocalVariableReferenceExpression() {
				syntheticLV1 := expression.(intsrv.IClassFileLocalVariableReferenceExpression).
					LocalVariable().(intsrv.ILocalVariable)
				expression = statements.Get(size - 4).Expression().LeftExpression()

				if expression.IsLocalVariableReferenceExpression() {
					lv2 := expression.(intsrv.IClassFileLocalVariableReferenceExpression).
						LocalVariable().(intsrv.ILocalVariable)

					if syntheticLV1 == lv2 {
						expression = statements.Get(size - 3).Expression()

						if expression.IsBinaryOperatorExpression() {
							boe2 := expression
							expression = boe2.RightExpression()

							if expression.IsIntegerConstantExpression() && MinusOne == expression.IntegerValue() {
								expression = switchStatement.Condition()

								if expression.IsLocalVariableReferenceExpression() {
									syntheticLV2 := expression.(intsrv.IClassFileLocalVariableReferenceExpression).
										LocalVariable().(intsrv.ILocalVariable)

									if syntheticLV2 == boe2.LeftExpression().(intsrv.IClassFileLocalVariableReferenceExpression).
										LocalVariable().(intsrv.ILocalVariable) {
										mie := previousSwitchStatement.Condition().(intmod.IMethodInvocationExpression)

										if mie.Name() == "hashCode" && mie.Descriptor() == "()I" {
											// Pattern found ==> Parse cases of the synthetic switch statement 'previousSwitchStatement'
											mapped := make(map[int]string)
											// Create map<synthetic index -> string>
											for _, block := range previousSwitchStatement.Blocks() {
												stmts := block.Statements()

												// assert (stmts != null) && stmts.IsStatements() && (stmts.Size() > 0);

												for _, stmt := range stmts.ToSlice() {
													if !stmt.IsIfStatement() {
														break
													}

													expression = stmt.Condition()

													if !expression.IsMethodInvocationExpression() {
														break
													}

													expression = expression.Parameters().First()

													if !expression.IsStringConstantExpression() {
														break
													}

													str := expression.StringValue()

													expression = stmt.Statements().First().Expression().RightExpression()

													if !expression.IsIntegerConstantExpression() {
														break
													}

													index := expression.IntegerValue()
													mapped[index] = str
												}
											}

											// Replace synthetic index by string
											for _, block := range switchStatement.Blocks() {
												if block.IsSwitchStatementLabelBlock() {
													lb := block.(intmod.ILabelBlock)

													if lb.Label() != modsts.DefaultLabe1.(intmod.ILabel) {
														el := lb.Label().(intmod.IExpressionLabel)
														nce := el.Expression().(intmod.IIntegerConstantExpression)
														el.SetExpression(modexp.NewStringConstantExpressionWithAll(
															nce.LineNumber(), mapped[nce.IntegerValue()]))
													}
												} else if block.IsSwitchStatementMultiLabelsBlock() {
													lmb := block.(intmod.IMultiLabelsBlock)

													for _, label := range lmb.Labels() {
														if label != modsts.DefaultLabe1.(intmod.ILabel) {
															el := label.(intmod.IExpressionLabel)
															nce := el.Expression().(intmod.IIntegerConstantExpression)
															el.SetExpression(modexp.NewStringConstantExpressionWithAll(
																nce.LineNumber(), mapped[nce.IntegerValue()]))
														}
													}
												}
											}

											// Replace synthetic key
											expression = statements.Get(size - 4).Expression()

											if expression.IsBinaryOperatorExpression() {
												switchStatement.SetCondition(expression.RightExpression())

												// Remove next synthetic statements
												statements.SubList(size-4, size-1).Clear()

												// Remove synthetic local variables
												localVariableMaker.RemoveLocalVariable(syntheticLV1)
												localVariableMaker.RemoveLocalVariable(syntheticLV2)
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func MakeSwitchEnum(bodyDeclaration intsrv.IClassFileBodyDeclaration, switchStatement intmod.ISwitchStatement) {
	expression := switchStatement.Condition().Expression()

	if expression.IsFieldReferenceExpression() {
		fre := expression.(intmod.IFieldReferenceExpression)

		if (fre.Descriptor() == "[I") && strings.HasPrefix(fre.Name(), "$SwitchMap$") {
			syntheticClassDeclaration := bodyDeclaration.InnerTypeDeclaration(fre.InternalTypeName())

			if syntheticClassDeclaration != nil {
				// Javac switch-enum pattern
				bodyDeclaration = syntheticClassDeclaration.BodyDeclaration().(intsrv.IClassFileBodyDeclaration)
				statements := bodyDeclaration.MethodDeclarations()[0].Statements().ToList()
				updateSwitchStatement(switchStatement, searchSwitchMap(fre, statements.Iterator()))
			}
		}
	} else if expression.IsMethodInvocationExpression() {
		mie := expression.(intmod.IMethodInvocationExpression)
		methodName := mie.Name()

		if (mie.Descriptor() == "()[I") && strings.HasPrefix(methodName, "$SWITCH_TABLE$") {
			// Eclipse compiler switch-enum pattern
			for _, declaration := range bodyDeclaration.MethodDeclarations() {
				if declaration.Method().Name() == methodName {
					statements := declaration.Statements().ToList()
					iterator := statements.ListIterator()
					iterator.SetCursor(3)
					updateSwitchStatement(switchStatement, iterator)
					break
				}
			}
		}
	}
}

func searchSwitchMap(fre intmod.IFieldReferenceExpression, iterator util.IIterator[intmod.IStatement]) util.IIterator[intmod.IStatement] {
	name := fre.Name()

	for iterator.HasNext() {
		expression := iterator.Next().Expression().LeftExpression()

		if expression.IsFieldReferenceExpression() && name == expression.Name() {
			return iterator
		}
	}

	return iterator
}

func updateSwitchStatement(switchStatement intmod.ISwitchStatement, iterator util.IIterator[intmod.IStatement]) {
	// Create map<synthetic index -> enum name>
	mapped := make(map[int]string)

	for iterator.HasNext() {
		statement := iterator.Next()
		if !statement.IsTryStatement() {
			break
		}

		statements := statement.TryStatements()
		if !statements.IsList() {
			break
		}

		expression := statements.First().Expression()
		if !expression.IsBinaryOperatorExpression() {
			break
		}

		boe := expression
		expression = boe.RightExpression()

		if !expression.IsIntegerConstantExpression() {
			break
		}

		index := expression.IntegerValue()
		expression = boe.LeftExpression()
		if !expression.IsArrayExpression() {
			break
		}

		expression = expression.Index()
		if !expression.IsMethodInvocationExpression() {
			break
		}

		expression = expression.Expression()
		if !expression.IsFieldReferenceExpression() {
			break
		}

		name := expression.(intmod.IFieldReferenceExpression).Name()
		mapped[index] = name
	}

	// Replace synthetic index by enum name
	expression := switchStatement.Condition().Index().Expression()
	typ := expression.Type().(intmod.IObjectType)

	for _, block := range switchStatement.Blocks() {
		if block.IsSwitchStatementLabelBlock() {
			lb := block.(intmod.ILabelBlock)
			if lb.Label() != modsts.DefaultLabe1.(intmod.ILabel) {
				el := lb.Label().(intmod.IExpressionLabel)
				nce := el.Expression().(intmod.IIntegerConstantExpression)
				el.SetExpression(modexp.NewEnumConstantReferenceExpressionWithAll(
					nce.LineNumber(), typ, mapped[nce.IntegerValue()]))
			}
		} else if block.IsSwitchStatementMultiLabelsBlock() {
			lmb := block.(intmod.IMultiLabelsBlock)
			for _, label := range lmb.Labels() {
				if label != modsts.DefaultLabe1.(intmod.ILabel) {
					el := label.(intmod.IExpressionLabel)
					nce := el.Expression().(intmod.IIntegerConstantExpression)
					el.SetExpression(modexp.NewEnumConstantReferenceExpressionWithAll(
						nce.LineNumber(), typ, mapped[nce.IntegerValue()]))
				}
			}
		}
	}

	// Replace synthetic key
	switchStatement.SetCondition(expression)
}
