package utils

import (
	intcls "bitbucket.org/coontec/go-jd-core/class/interfaces/classpath"
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
	modexp "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/expression"
	modsts "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/statement"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
	"bitbucket.org/coontec/go-jd-core/class/service/converter/model/cfg"
	srvsts "bitbucket.org/coontec/go-jd-core/class/service/converter/model/javasyntax/statement"
	"bitbucket.org/coontec/go-jd-core/class/service/converter/visitor"
	"bitbucket.org/coontec/go-jd-core/class/util"
	"fmt"
	"strings"
)

var GlobalFinallyExceptionExpression = modexp.NewNullExpression(
	_type.NewObjectType("java/lang/Exception",
		"java.lang.Exception", "Exception"))
var GlobalMergeTryWithResourcesStatementVisitor = visitor.NewMergeTryWithResourcesStatementVisitor()

func NewStatementMaker(typeMaker intsrv.ITypeMaker, localVariableMaker intsrv.ILocalVariableMaker,
	comd intsrv.IClassFileConstructorOrMethodDeclaration) intsrv.IStatementMaker {
	classFile := comd.ClassFile()

	m := &StatementMaker{
		typeMaker:                             typeMaker,
		typeBounds:                            comd.TypeBounds(),
		localVariableMaker:                    localVariableMaker,
		majorVersion:                          classFile.MajorVersion(),
		internalTypeName:                      classFile.InternalTypeName(),
		bodyDeclaration:                       comd.BodyDeclaration(),
		stack:                                 util.NewDefaultStack[intmod.IExpression](),
		byteCodeParser:                        NewByteCodeParser(typeMaker, localVariableMaker, classFile, comd.BodyDeclaration(), comd),
		removeFinallyStatementsVisitor:        visitor.NewRemoveFinallyStatementsVisitor(localVariableMaker),
		removeBinaryOpReturnStatementsVisitor: visitor.NewRemoveBinaryOpReturnStatementsVisitor(localVariableMaker),
		updateIntegerConstantTypeVisitor:      visitor.NewUpdateIntegerConstantTypeVisitor(comd.ReturnedType()),
		searchFirstLineNumberVisitor:          visitor.NewSearchFirstLineNumberVisitor(),
		memberVisitor:                         NewStatementMakerMemberVisitor(),
		removeFinallyStatementsFlag:           false,
		mergeTryWithResourcesStatementFlag:    false,
	}

	return m
}

type StatementMaker struct {
	typeMaker                             intsrv.ITypeMaker
	typeBounds                            map[string]intmod.IType
	localVariableMaker                    intsrv.ILocalVariableMaker
	byteCodeParser                        intsrv.IByteCodeParser
	majorVersion                          int
	internalTypeName                      string
	bodyDeclaration                       intsrv.IClassFileBodyDeclaration
	stack                                 util.IStack[intmod.IExpression]
	removeFinallyStatementsVisitor        *visitor.RemoveFinallyStatementsVisitor
	removeBinaryOpReturnStatementsVisitor *visitor.RemoveBinaryOpReturnStatementsVisitor
	updateIntegerConstantTypeVisitor      *visitor.UpdateIntegerConstantTypeVisitor
	searchFirstLineNumberVisitor          *visitor.SearchFirstLineNumberVisitor
	memberVisitor                         *StatementMakerMemberVisitor
	removeFinallyStatementsFlag           bool
	mergeTryWithResourcesStatementFlag    bool
}

func (m *StatementMaker) Make(cfg intsrv.IControlFlowGraph) intmod.IStatements {
	statements := modsts.NewStatements()
	jumps := modsts.NewStatements()
	watchdog := NewWatchDog()

	m.localVariableMaker.PushFrame(statements)

	// Generate statements
	m.makeStatements(watchdog, cfg.Start(), statements, jumps)

	// Remove 'finally' statements
	if m.removeFinallyStatementsFlag {
		m.removeFinallyStatementsVisitor.Init()
		statements.Accept(m.removeFinallyStatementsVisitor)
	}

	// Merge 'try-with-resources' statements
	if m.mergeTryWithResourcesStatementFlag {
		statements.Accept(GlobalMergeTryWithResourcesStatementVisitor)
	}

	// Replace pattern "synthetic_local_var = ...; return synthetic_local_var;" with "return ...;"
	statements.Accept(m.removeBinaryOpReturnStatementsVisitor)

	// Remove last 'return' statement
	if !statements.IsEmpty() && statements.Last().IsReturnStatement() {
		statements.RemoveLast()
	}

	m.localVariableMaker.PopFrame()

	// Update integer constant type to 'byte', 'char', 'short' or 'int'
	statements.Accept(m.updateIntegerConstantTypeVisitor)

	// Change ++i; with i++;
	m.replacePreOperatorWithPostOperator(statements)

	if !jumps.IsEmpty() {
		m.updateJumpStatements(jumps)
	}

	return statements
}

/**
 * A recursive, next neighbour first, statements builder from basic blocks.
 *
 * @param basicBlock Current basic block
 * @param statements List to populate
 */
func (m *StatementMaker) makeStatements(watchdog IWatchDog, basicBlock intsrv.IBasicBlock,
	statements intmod.IStatements, jumps intmod.IStatements) {
	var subStatements, elseStatements intmod.IStatements
	var condition, exp1, exp2 intmod.IExpression

	switch basicBlock.Type() {
	case intsrv.TypeStart:
		watchdog.Check(basicBlock, basicBlock.Next())
		m.makeStatements(watchdog, basicBlock.Next(), statements, jumps)
		break
	case intsrv.TypeEnd:
		break
	case intsrv.TypeStatements:
		watchdog.Check(basicBlock, basicBlock.Next())
	case intsrv.TypeThrow:
		m.parseByteCode(basicBlock, statements)
		m.makeStatements(watchdog, basicBlock.Next(), statements, jumps)
		break
	case intsrv.TypeReturn:
		statements.Add(modsts.Return)
		break
	case intsrv.TypeReturnValue:
	case intsrv.TypeGotoInTernaryOperator:
		m.parseByteCode(basicBlock, statements)
		break
	case intsrv.TypeSwitch:
		m.parseSwitch(watchdog, basicBlock, statements, jumps)
		break
	case intsrv.TypeSwitchBreak:
		statements.Add(modsts.Break)
		break
	case intsrv.TypeTry:
		m.parseTry(watchdog, basicBlock, statements, jumps, false, false)
		break
	case intsrv.TypeTryJsr:
		m.parseTry(watchdog, basicBlock, statements, jumps, true, false)
		break
	case intsrv.TypeTryEclipse:
		m.parseTry(watchdog, basicBlock, statements, jumps, false, true)
		break
	case intsrv.TypeJsr:
		m.parseJSR(watchdog, basicBlock, statements, jumps)
		break
	case intsrv.TypeRet:
		m.parseByteCode(basicBlock, statements)
		break
	case intsrv.TypeIf:
		m.parseIf(watchdog, basicBlock, statements, jumps)
		break
	case intsrv.TypeIfElse:
		watchdog.Check(basicBlock, basicBlock.Condition())
		m.makeStatements(watchdog, basicBlock.Condition(), statements, jumps)
		condition = m.stack.Pop()
		backup := util.NewDefaultStackFrom(m.stack)
		watchdog.Check(basicBlock, basicBlock.Sub1())
		subStatements = m.makeSubStatements2(watchdog, basicBlock.Sub1(), statements, jumps)
		if !basicBlock.Sub2().MatchType(intsrv.TypeLoopEnd|intsrv.TypeLoopContinue|intsrv.TypeLoopStart) &&
			(m.stack.Size() != backup.Size()) {
			m.stack.Copy(backup)
		}
		watchdog.Check(basicBlock, basicBlock.Sub2())
		elseStatements = m.makeSubStatements2(watchdog, basicBlock.Sub2(), statements, jumps)
		statements.Add(modsts.NewIfElseStatement(condition, subStatements, elseStatements))
		watchdog.Check(basicBlock, basicBlock.Next())
		m.makeStatements(watchdog, basicBlock.Next(), statements, jumps)
		break
	case intsrv.TypeCondition:
		if basicBlock.Sub1() != cfg.End {
			m.stack.Push(m.makeExpression(watchdog, basicBlock.Sub1(), statements, jumps))
		}
		if basicBlock.Sub2() != cfg.End {
			m.stack.Push(m.makeExpression(watchdog, basicBlock.Sub2(), statements, jumps))
		}
		m.parseByteCode(basicBlock, statements)
		break
	case intsrv.TypeConditionOr:
		watchdog.Check(basicBlock, basicBlock.Sub1())
		exp1 = m.makeExpression(watchdog, basicBlock.Sub1(), statements, jumps)
		watchdog.Check(basicBlock, basicBlock.Sub2())
		exp2 = m.makeExpression(watchdog, basicBlock.Sub2(), statements, jumps)
		m.stack.Push(modexp.NewBinaryOperatorExpression(basicBlock.FirstLineNumber(),
			_type.PtTypeBoolean, exp1, "||", exp2, 14))
		break
	case intsrv.TypeConditionAnd:
		watchdog.Check(basicBlock, basicBlock.Sub1())
		exp1 = m.makeExpression(watchdog, basicBlock.Sub1(), statements, jumps)
		watchdog.Check(basicBlock, basicBlock.Sub2())
		exp2 = m.makeExpression(watchdog, basicBlock.Sub2(), statements, jumps)
		m.stack.Push(modexp.NewBinaryOperatorExpression(basicBlock.FirstLineNumber(),
			_type.PtTypeBoolean, exp1, "&&", exp2, 13))
		break
	case intsrv.TypeConditionTernaryOperator:
		watchdog.Check(basicBlock, basicBlock.Condition())
		m.makeStatements(watchdog, basicBlock.Condition(), statements, jumps)
		condition = m.stack.Pop()
		backup := util.NewDefaultStackFrom(m.stack)
		watchdog.Check(basicBlock, basicBlock.Sub1())
		exp1 = m.makeExpression(watchdog, basicBlock.Sub1(), statements, jumps)
		if m.stack.Size() != backup.Size() {
			m.stack.Copy(backup)
		}
		watchdog.Check(basicBlock, basicBlock.Sub2())
		exp2 = m.makeExpression(watchdog, basicBlock.Sub2(), statements, jumps)
		m.stack.Push(m.parseTernaryOperator(basicBlock.FirstLineNumber(), condition, exp1, exp2))
		m.parseByteCode(basicBlock, statements)
		break
	case intsrv.TypeTernaryOperator:
		watchdog.Check(basicBlock, basicBlock.Condition())
		m.makeStatements(watchdog, basicBlock.Condition(), statements, jumps)
		condition = m.stack.Pop()
		backup := util.NewDefaultStackFrom(m.stack)
		watchdog.Check(basicBlock, basicBlock.Sub1())
		exp1 = m.makeExpression(watchdog, basicBlock.Sub1(), statements, jumps)
		if m.stack.Size() != backup.Size() {
			m.stack.Copy(backup)
		}
		watchdog.Check(basicBlock, basicBlock.Sub2())
		exp2 = m.makeExpression(watchdog, basicBlock.Sub2(), statements, jumps)
		m.stack.Push(m.parseTernaryOperator(basicBlock.FirstLineNumber(), condition, exp1, exp2))
		watchdog.Check(basicBlock, basicBlock.Next())
		m.makeStatements(watchdog, basicBlock.Next(), statements, jumps)
		break
	case intsrv.TypeLoop:
		m.parseLoop(watchdog, basicBlock, statements, jumps)
		break
	case intsrv.TypeLoopStart, intsrv.TypeLoopContinue:
		statements.Add(modsts.Continue)
		break
	case intsrv.TypeLoopEnd:
		statements.Add(modsts.Break)
		break
	case intsrv.TypeJump:
		jump := srvsts.NewClassFileBreakContinueStatement(basicBlock.FromOffset(), basicBlock.ToOffset())
		statements.Add(jump)
		jumps.Add(jump)
		break
	case intsrv.TypeInfiniteGoto:
		statements.Add(modsts.NewWhileStatement(modexp.True, nil))
		break
	default:
		//assert false : "Unexpected basic block: " + basicBlock.TypeName() + ':' + basicBlock.Index();
		break
	}
}

func (m *StatementMaker) makeSubStatements(watchdog IWatchDog, basicBlock intsrv.IBasicBlock,
	statements intmod.IStatements, jumps intmod.IStatements, updateBasicBlock intsrv.IBasicBlock) intmod.IStatements {
	subStatements := m.makeSubStatements2(watchdog, basicBlock, statements, jumps)

	if updateBasicBlock != nil {
		subStatements.AddAll(m.makeSubStatements2(watchdog, updateBasicBlock, statements, jumps).ToSlice())
	}

	return subStatements
}

func (m *StatementMaker) makeSubStatements2(watchdog IWatchDog, basicBlock intsrv.IBasicBlock,
	statements intmod.IStatements, jumps intmod.IStatements) intmod.IStatements {
	subStatements := modsts.NewStatements()

	if !statements.IsEmpty() && statements.Last().IsMonitorEnterStatement() {
		subStatements.Add(statements.RemoveLast())
	}

	m.localVariableMaker.PushFrame(subStatements)
	m.makeStatements(watchdog, basicBlock, subStatements, jumps)
	m.localVariableMaker.PopFrame()
	m.replacePreOperatorWithPostOperator(subStatements)

	if !subStatements.IsEmpty() && subStatements.First().IsMonitorEnterStatement() {
		statements.Add(subStatements.RemoveFirst())
	}

	return subStatements
}

func (m *StatementMaker) makeExpression(watchdog IWatchDog, basicBlock intsrv.IBasicBlock,
	statements intmod.IStatements, jumps intmod.IStatements) intmod.IExpression {
	initialStatementCount := statements.Size()

	m.makeStatements(watchdog, basicBlock, statements, jumps)

	if m.stack.IsEmpty() {
		// Interesting... Kotlin pattern.
		// https://github.com/JetBrains/intellij-community/blob/master/platform/built-in-server/src/org/jetbrains/builtInWebServer/SingleConnectionNetService.kt
		// final override fun connectToProcess(...)
		return modexp.NewStringConstantExpression("JD-Core does not support Kotlin")
	} else {
		expression := m.stack.Pop()

		if statements.Size() > initialStatementCount {
			// Is a multi-assignment ?
			boe := statements.Last().Expression()

			if boe != modexp.NeNoExpression {
				if boe.RightExpression() == expression {
					// Pattern matched -> Multi-assignment
					statements.RemoveLast()
					expression = boe
				} else if expression.IsNewArray() {
					expression = MakeNewArrayMaker(statements, expression)
				}
			}
		}

		return expression
	}
}

func (m *StatementMaker) parseSwitch(watchdog IWatchDog, basicBlock intsrv.IBasicBlock, statements intmod.IStatements, jumps intmod.IStatements) {
	m.parseByteCode(basicBlock, statements)

	switchCases := basicBlock.SwitchCases()
	switchStatement := statements.Last().(intmod.ISwitchStatement)
	condition := switchStatement.Condition()
	conditionType := condition.Type()
	blocks := util.NewDefaultListWithSlice(switchStatement.Blocks())
	localStack := util.NewDefaultStackFrom(m.stack)

	switchCases.Sort(func(i, j int) bool {
		sc1 := switchCases.Get(i)
		sc2 := switchCases.Get(i)
		diff := sc1.Offset() - sc2.Offset()
		if diff != 0 {
			return false
		}
		return (sc1.Value() - sc2.Value()) == 0
	})

	length := switchCases.Size()
	for i := 0; i < length; i++ {
		sc := switchCases.Get(i)
		bb := sc.BasicBlock()
		j := i + 1

		for (j < length) && (bb == switchCases.Get(j).BasicBlock()) {
			j++
		}

		subStatements := modsts.NewStatements()

		m.stack.Copy(localStack)
		m.makeStatements(watchdog, bb, subStatements, jumps)
		m.replacePreOperatorWithPostOperator(subStatements)

		if sc.IsDefaultCase() {
			blocks.Add(modsts.NewLabelBlock(modsts.DefaultLabe1.(intmod.ILabel), subStatements))
		} else if j == i+1 {
			label := modsts.NewExpressionLabel(modexp.NewIntegerConstantExpression(conditionType, sc.Value()))
			blocks.Add(modsts.NewLabelBlock(label.(intmod.ILabel), subStatements))
		} else {
			labels := util.NewDefaultListWithCapacity[intmod.ILabel](j - i)

			for ; i < j; i++ {
				labels.Add(modsts.NewExpressionLabel(modexp.NewIntegerConstantExpression(conditionType,
					switchCases.Get(i).Value())).(intmod.ILabel))
			}

			blocks.Add(modsts.NewMultiLabelsBlock(labels.ToSlice(), subStatements))
			i--
		}
	}

	size := statements.Size()

	if (size > 3) && condition.IsLocalVariableReferenceExpression() &&
		statements.Get(size-2).IsSwitchStatement() {
		// Check pattern & make 'switch-string'
		MakeSwitchString(m.localVariableMaker, statements, switchStatement)
	} else if condition.IsArrayExpression() {
		// Check pattern & make 'switch-enum'
		MakeSwitchEnum(m.bodyDeclaration, switchStatement)
	}

	m.makeStatements(watchdog, basicBlock.Next(), statements, jumps)
}

func (m *StatementMaker) parseTry(watchdog IWatchDog, basicBlock intsrv.IBasicBlock,
	statements intmod.IStatements, jumps intmod.IStatements, jsr, eclipse bool) {
	var tryStatements intmod.IStatements
	catchClauses := util.NewDefaultList[intmod.ICatchClause]()
	var finallyStatements intmod.IStatements
	// assertStackSize := m.stack.Size()

	tryStatements = m.makeSubStatements2(watchdog, basicBlock.Sub1(), statements, jumps)

	for _, exceptionHandler := range basicBlock.ExceptionHandlers().ToSlice() {
		//assert stack.size() == assertStackSize : "parseTry : problem with stack";

		if exceptionHandler.InternalThrowableName() == "" {
			// 'finally' handler
			m.stack.Push(GlobalFinallyExceptionExpression)

			finallyStatements = m.makeSubStatements2(watchdog, exceptionHandler.BasicBlock(), statements, jumps)

			if !finallyStatements.First().IsMonitorEnterStatement() {
				m.removeFinallyStatementsFlag = m.removeFinallyStatementsFlag || (jsr == false)

				leftExpression := finallyStatements.First().Expression().LeftExpression()

				if leftExpression.IsLocalVariableReferenceExpression() {
					statement := finallyStatements.Last()

					if statement.IsThrowStatement() {
						// Remove synthetic local variable
						expression := statement.Expression()

						if expression.IsLocalVariableReferenceExpression() {
							vre1 := expression.(intsrv.IClassFileLocalVariableReferenceExpression)
							vre2 := leftExpression.(intsrv.IClassFileLocalVariableReferenceExpression)

							if vre1.LocalVariable() == vre2.LocalVariable() {
								m.localVariableMaker.RemoveLocalVariable(vre2.LocalVariable().(intsrv.ILocalVariable))
								// Remove first statement (storage of finally localVariable)
								finallyStatements.RemoveFirst()
							}
						}
					}
				}

				// Remove last statement (throw finally localVariable)
				finallyStatements.RemoveLast()
			}
		} else {
			m.stack.Push(modexp.NewNullExpression(
				m.typeMaker.MakeFromInternalTypeName(
					exceptionHandler.InternalThrowableName())))

			catchStatements := modsts.NewStatements()
			m.localVariableMaker.PushFrame(catchStatements)

			bb := exceptionHandler.BasicBlock()
			lineNumber := bb.ControlFlowGraph().LineNumber(bb.FromOffset())
			index := GetExceptionLocalVariableIndex(bb)
			ot := m.typeMaker.MakeFromInternalTypeName(exceptionHandler.InternalThrowableName())
			offset := bb.FromOffset()
			code := bb.ControlFlowGraph().Method().Attribute("Code").(intcls.IAttributeCode).Code()

			if code[offset] == 58 {
				offset += 2 // ASTORE
			} else {
				offset++ // POP, ASTORE_1 ... ASTORE_3
			}

			exception := m.localVariableMaker.ExceptionLocalVariable(index, offset, ot)

			m.makeStatements(watchdog, bb, catchStatements, jumps)
			m.localVariableMaker.PopFrame()
			m.removeExceptionReference(catchStatements)

			if lineNumber != intmod.UnknownLineNumber {
				m.searchFirstLineNumberVisitor.Init()
				m.searchFirstLineNumberVisitor.VisitStatements(catchStatements)
				if m.searchFirstLineNumberVisitor.LineNumber() == lineNumber {
					lineNumber = intmod.UnknownLineNumber
				}
			}

			m.replacePreOperatorWithPostOperator(catchStatements)

			cc := srvsts.NewCatchClause(lineNumber, ot, exception, catchStatements)

			if exceptionHandler.OtherInternalThrowableNames() != nil {
				for _, name := range exceptionHandler.OtherInternalThrowableNames().ToSlice() {
					cc.AddType(m.typeMaker.MakeFromInternalTypeName(name))
				}
			}

			catchClauses.Add(cc)
		}
	}

	// 'try', 'try-with-resources' or 'synchronized' ?
	var statement intmod.IStatement = nil

	if (finallyStatements != nil) && (finallyStatements.Size() > 0) &&
		finallyStatements.First().IsMonitorExitStatement() {
		statement = MakeSynchronizedStatementMaker(m.localVariableMaker, statements, tryStatements)
	} else {
		if m.majorVersion >= 51 { // (majorVersion >= Java 7)
			// assert jsr == false;
			statement = MakeTryWithResourcesStatementMaker(m.localVariableMaker,
				statements, tryStatements, catchClauses, finallyStatements)
		}
		if statement == nil {
			statement = srvsts.NewClassFileTryStatement(tryStatements,
				catchClauses.ToSlice(), finallyStatements, jsr, eclipse)
		} else {
			m.mergeTryWithResourcesStatementFlag = true
		}
	}

	statements.Add(statement)
	m.makeStatements(watchdog, basicBlock.Next(), statements, jumps)
}

func (m *StatementMaker) removeExceptionReference(catchStatements intmod.IStatements) {
	if (catchStatements.Size() > 0) && catchStatements.First().IsExpressionStatement() {
		exp := catchStatements.First().Expression()

		if exp.IsBinaryOperatorExpression() {
			if exp.LeftExpression().IsLocalVariableReferenceExpression() &&
				exp.RightExpression().IsNullExpression() {
				catchStatements.RemoveFirst()
			}
		} else if exp.IsNullExpression() {
			catchStatements.RemoveFirst()
		}
	}
}

func (m *StatementMaker) parseJSR(watchdog IWatchDog, basicBlock intsrv.IBasicBlock,
	statements intmod.IStatements, jumps intmod.IStatements) {
	statementCount := statements.Size()

	m.parseByteCode(basicBlock, statements)
	m.makeStatements(watchdog, basicBlock.Branch(), statements, jumps)
	m.makeStatements(watchdog, basicBlock.Next(), statements, jumps)

	// Remove synthetic local variable
	es := statements.Get(statementCount).(intmod.IExpressionStatement)
	boe := es.Expression().(intmod.IBinaryOperatorExpression)
	vre := boe.LeftExpression().(intsrv.IClassFileLocalVariableReferenceExpression)

	m.localVariableMaker.RemoveLocalVariable(vre.LocalVariable().(intsrv.ILocalVariable))
	// Remove first statement (storage of JSR return offset)
	statements.RemoveAt(statementCount)
}

func (m *StatementMaker) parseIf(watchdog IWatchDog, basicBlock intsrv.IBasicBlock, statements intmod.IStatements, jumps intmod.IStatements) {
	condition := basicBlock.Condition()

	if condition.Type() == intsrv.TypeConditionAnd {
		condition = condition.Sub1()
	}

	if IsAssertCondition(m.internalTypeName, condition) {
		var cond intmod.IExpression

		if condition == basicBlock.Condition() {
			cond = modexp.NewBooleanExpressionWithLineNumber(condition.FirstLineNumber(), false)
		} else {
			condition = basicBlock.Condition().Sub2()
			condition.InverseCondition()
			m.makeStatements(watchdog, condition, statements, jumps)
			cond = m.stack.Pop()
		}

		subStatements := m.makeSubStatements2(watchdog, basicBlock.Sub1(), statements, jumps)
		var message intmod.IExpression = nil

		if subStatements.First().IsThrowStatement() {
			e := subStatements.First().Expression()
			if e.IsNewExpression() {
				parameters := e.Parameters()
				if (parameters != nil) && !parameters.IsList() {
					message = parameters.First()
				}
			}
		}

		statements.Add(modsts.NewAssertStatement(cond, message))
		m.makeStatements(watchdog, basicBlock.Next(), statements, jumps)
	} else {
		m.makeStatements(watchdog, basicBlock.Condition(), statements, jumps)
		cond := m.stack.Pop()
		backup := util.NewDefaultStackFrom(m.stack)
		subStatements := m.makeSubStatements2(watchdog, basicBlock.Sub1(), statements, jumps)
		if m.stack.Size() != backup.Size() {
			m.stack.Copy(backup)
		}
		statements.Add(modsts.NewIfStatement(cond, subStatements))
		index := statements.Size()
		m.makeStatements(watchdog, basicBlock.Next(), statements, jumps)

		if (subStatements.Size() == 1) &&
			(index+1 == statements.Size()) &&
			(subStatements.First().IsReturnExpressionStatement()) &&
			(statements.Get(index).IsReturnExpressionStatement()) {
			cfres1 := subStatements.First()

			if cond.LineNumber() >= cfres1.LineNumber() {
				cfres2 := statements.Get(index)

				if cfres1.LineNumber() == cfres2.LineNumber() {
					statements.SubList(index-1, statements.Size()).Clear()
					statements.Add(modsts.NewReturnExpressionStatement(
						m.newTernaryOperatorExpression(cfres1.LineNumber(),
							cond, cfres1.Expression(), cfres2.Expression())))
				}
			}
		}
	}
}

func (m *StatementMaker) parseLoop(watchdog IWatchDog, basicBlock intsrv.IBasicBlock, statements intmod.IStatements, jumps intmod.IStatements) {
	sub1 := basicBlock.Sub1()
	var updateBasicBlock intsrv.IBasicBlock = nil

	if (sub1.Type() == intsrv.TypeIf) && (sub1.Condition() == cfg.End) {
		updateBasicBlock = sub1.Next()
		sub1 = sub1.Sub1()
	}

	if sub1.Type() == intsrv.TypeIf {
		ifBB := sub1

		if ifBB.Next() == cfg.LoopEnd {
			// 'while' or 'for' loop
			m.makeStatements(watchdog, ifBB.Condition(), statements, jumps)
			statements.Add(MakeLoopStatementMaker(m.majorVersion, m.typeBounds, m.localVariableMaker,
				basicBlock, statements, m.stack.Pop(), m.makeSubStatements(watchdog,
					ifBB.Sub1(), statements, jumps, updateBasicBlock), jumps))
			m.makeStatements(watchdog, basicBlock.Next(), statements, jumps)
			return
		}

		if ifBB.Sub1() == cfg.LoopEnd {
			if ifBB.Next() == cfg.LoopStart {
				// 'do-while' pattern
				ifBB.Condition().InverseCondition()

				subStatements := modsts.NewStatements()

				m.makeStatements(watchdog, ifBB.Condition(), subStatements, jumps)
				m.replacePreOperatorWithPostOperator(subStatements)
				statements.Add(MakeDoWhileLoop(basicBlock, m.stack.Pop(), subStatements, jumps))
			} else {
				// 'while' or 'for' loop
				ifBB.Condition().InverseCondition()
				m.makeStatements(watchdog, ifBB.Condition(), statements, jumps)
				statements.Add(MakeLoopStatementMaker(m.majorVersion, m.typeBounds, m.localVariableMaker,
					basicBlock, statements, m.stack.Pop(), m.makeSubStatements(watchdog,
						ifBB.Next(), statements, jumps, updateBasicBlock), jumps))
			}

			m.makeStatements(watchdog, basicBlock.Next(), statements, jumps)
			return
		}
	}

	next := sub1.Next()
	last := sub1

	for next.MatchType(intsrv.GroupSingleSuccessor) && (next.Predecessors().Size() == 1) {
		last = next
		next = next.Next()
	}

	if (next == cfg.LoopStart) && (last.Type() == intsrv.TypeIf) &&
		(last.Sub1() == cfg.LoopEnd) && (m.countStartLoop(sub1) == 1) {
		// 'do-while'
		var subStatements intmod.IStatements

		last.Condition().InverseCondition()
		last.SetType(intsrv.TypeEnd)

		if (sub1.Type() == intsrv.TypeLoop) && (sub1.Next() == last) && (m.countStartLoop(sub1.Sub1()) == 0) {
			changeEndLoopToStartLoop(util.NewBitSet(), sub1.Sub1())
			subStatements = m.makeSubStatements(watchdog, sub1.Sub1(), statements, jumps, updateBasicBlock)

			// assert subStatements.Last() == ContinueStatement.CONTINUE :
			//"StatementMaker.parseLoop(...) : unexpected basic block for create a do-while loop";

			subStatements.RemoveLast()
		} else {
			m.createDoWhileContinue(last)
			subStatements = m.makeSubStatements(watchdog, sub1, statements, jumps, updateBasicBlock)
		}

		m.makeStatements(watchdog, last.Condition(), subStatements, jumps)
		statements.Add(MakeDoWhileLoop(basicBlock, m.stack.Pop(), subStatements, jumps))
	} else {
		// Infinite loop
		statements.Add(MakeLoopStatementMaker2(m.localVariableMaker, basicBlock, statements,
			m.makeSubStatements(watchdog, sub1, statements, jumps, updateBasicBlock), jumps))
	}

	m.makeStatements(watchdog, basicBlock.Next(), statements, jumps)
}

func (m *StatementMaker) countStartLoop(bb intsrv.IBasicBlock) int {
	count := 0

	for bb.MatchType(intsrv.GroupSingleSuccessor) {
		switch bb.Type() {
		case intsrv.TypeSwitch:
			for _, switchCase := range bb.SwitchCases().ToSlice() {
				count += m.countStartLoop(switchCase.BasicBlock())
			}
			break
		case intsrv.TypeTry, intsrv.TypeTryJsr, intsrv.TypeTryEclipse:
			count += m.countStartLoop(bb.Sub1())
			for _, exceptionHandler := range bb.ExceptionHandlers().ToSlice() {
				count += m.countStartLoop(exceptionHandler.BasicBlock())
			}
			break
		case intsrv.TypeIfElse, intsrv.TypeTernaryOperator:
			count += m.countStartLoop(bb.Sub2())
		case intsrv.TypeIf:
			count += m.countStartLoop(bb.Sub1())
			break
		}

		bb = bb.Next()
	}

	if bb.Type() == intsrv.TypeLoopStart {
		count++
	}

	return count
}

func (m *StatementMaker) createDoWhileContinue(last intsrv.IBasicBlock) {
	var change bool

	for {
		change = false

		for _, predecessor := range last.Predecessors().ToSlice() {
			if predecessor.Type() == intsrv.TypeIf {
				l := predecessor.Sub1()

				if l.MatchType(intsrv.GroupSingleSuccessor) {
					n := l.Next()

					for n.MatchType(intsrv.GroupSingleSuccessor) {
						l = n
						n = n.Next()
					}

					if n == cfg.End {
						// Transform 'if' to 'if-continue'
						predecessor.Condition().InverseCondition()
						predecessor.SetNext(predecessor.Sub1())
						last.Predecessors().Remove(predecessor)
						last.Predecessors().Add(l)
						l.SetNext(last)
						predecessor.SetSub1(cfg.LoopStart)
						change = true
					}
				}
			}
		}

		if !change {
			break
		}
	}
}

func changeEndLoopToStartLoop(visited util.IBitSet, basicBlock intsrv.IBasicBlock) {
	if !basicBlock.MatchType(intsrv.GroupEnd|intsrv.TypeLoopEnd) && (visited.Get(basicBlock.Index()) == false) {
		visited.Set(basicBlock.Index())

		switch basicBlock.Type() {
		case intsrv.TypeConditionalBranch, intsrv.TypeJsr, intsrv.TypeCondition:
			if basicBlock.Branch() == cfg.LoopEnd {
				basicBlock.SetBranch(cfg.LoopStart)
			} else {
				changeEndLoopToStartLoop(visited, basicBlock.Branch())
			}
		case intsrv.TypeStart, intsrv.TypeStatements, intsrv.TypeGoto,
			intsrv.TypeGotoInTernaryOperator, intsrv.TypeLoop:
			if basicBlock.Next() == cfg.LoopEnd {
				basicBlock.SetNext(cfg.LoopStart)
			} else {
				changeEndLoopToStartLoop(visited, basicBlock.Next())
			}
			break
		case intsrv.TypeTryDeclaration, intsrv.TypeTry, intsrv.TypeTryJsr, intsrv.TypeTryEclipse:
			for _, exceptionHandler := range basicBlock.ExceptionHandlers().ToSlice() {
				if exceptionHandler.BasicBlock() == cfg.LoopEnd {
					exceptionHandler.SetBasicBlock(cfg.LoopStart)
				} else {
					changeEndLoopToStartLoop(visited, exceptionHandler.BasicBlock())
				}
			}
			break
		case intsrv.TypeIfElse, intsrv.TypeTernaryOperator:
			if basicBlock.Sub2() == cfg.LoopEnd {
				basicBlock.SetSub2(cfg.LoopStart)
			} else {
				changeEndLoopToStartLoop(visited, basicBlock.Sub2())
			}
		case intsrv.TypeIf:
			if basicBlock.Sub1() == cfg.LoopEnd {
				basicBlock.SetSub1(cfg.LoopStart)
			} else {
				changeEndLoopToStartLoop(visited, basicBlock.Sub1())
			}
			if basicBlock.Next() == cfg.LoopEnd {
				basicBlock.SetNext(cfg.LoopStart)
			} else {
				changeEndLoopToStartLoop(visited, basicBlock.Next())
			}
			break
		case intsrv.TypeSwitch, intsrv.TypeSwitchDeclaration:
			for _, switchCase := range basicBlock.SwitchCases().ToSlice() {
				if switchCase.BasicBlock() == cfg.LoopEnd {
					switchCase.SetBasicBlock(cfg.LoopStart)
				} else {
					changeEndLoopToStartLoop(visited, switchCase.BasicBlock())
				}
			}
			break
		}
	}
}

func (m *StatementMaker) parseTernaryOperator(lineNumber int, condition, exp1, exp2 intmod.IExpression) intmod.IExpression {
	if _type.OtTypeClass == exp1.Type() && _type.OtTypeClass == exp2.Type() && condition.IsBinaryOperatorExpression() {

		if condition.LeftExpression().IsFieldReferenceExpression() && condition.RightExpression().IsNullExpression() {
			freCond := condition.LeftExpression().(intmod.IFieldReferenceExpression)

			if freCond.InternalTypeName() == m.internalTypeName {
				fieldName := freCond.Name()

				if strings.HasPrefix(fieldName, "class$") {
					if condition.Operator() == "==" && exp1.IsBinaryOperatorExpression() && m.checkFieldReference(fieldName, exp2) {
						if exp1.RightExpression().IsMethodInvocationExpression() && m.checkFieldReference(fieldName, exp1.LeftExpression()) {
							mie := exp1.RightExpression().(intmod.IMethodInvocationExpression)

							if mie.Parameters().IsStringConstantExpression() &&
								mie.Name() == "class$" && mie.InternalTypeName() == m.internalTypeName {
								// JDK 1.4.2 '.class' found ==> Convert '(class$java$lang$String == nil) ? (class$java$lang$String = TestDotClass.class$("java.lang.String") : class$java$lang$String)' to 'String.class'
								return m.createObjectTypeReferenceDotClassExpression(lineNumber, fieldName, mie)
							}
						}
					} else if condition.Operator() == "!=" && exp2.IsBinaryOperatorExpression() && m.checkFieldReference(fieldName, exp1) {
						if exp2.RightExpression().IsMethodInvocationExpression() && m.checkFieldReference(fieldName, exp2.LeftExpression()) {
							mie := exp2.RightExpression().(intmod.IMethodInvocationExpression)

							if mie.Parameters().IsStringConstantExpression() && mie.Name() == "class$" && mie.InternalTypeName() == m.internalTypeName {
								// JDK 1.1.8 '.class' found ==> Convert '(class$java$lang$String != nil) ? class$java$lang$String : (class$java$lang$String = TestDotClass.class$("java.lang.String"))' to 'String.class'
								return m.createObjectTypeReferenceDotClassExpression(lineNumber, fieldName, mie)
							}
						}
					}
				}
			}
		}
	}

	return m.newTernaryOperatorExpression(lineNumber, condition, exp1, exp2)
}

func (m *StatementMaker) newTernaryOperatorExpression(lineNumber int,
	condition, expressionTrue, expressionFalse intmod.IExpression) intmod.ITernaryOperatorExpression {
	expressionTrueType := expressionTrue.Type()
	expressionFalseType := expressionFalse.Type()
	var typ intmod.IType

	if expressionTrue.IsNullExpression() {
		typ = expressionFalseType
	} else if expressionFalse.IsNullExpression() {
		typ = expressionTrueType
	} else if expressionTrueType == expressionFalseType {
		typ = expressionTrueType
	} else if expressionTrueType.IsPrimitiveType() && expressionFalseType.IsPrimitiveType() {
		flags := expressionTrueType.(intmod.IPrimitiveType).Flags() |
			expressionFalseType.(intmod.IPrimitiveType).Flags()

		if (flags & intmod.FlagDouble) != 0 {
			typ = _type.PtTypeDouble
		} else if (flags & intmod.FlagFloat) != 0 {
			typ = _type.PtTypeFloat
		} else if (flags & intmod.FlagLong) != 0 {
			typ = _type.PtTypeLong
		} else {
			flags = expressionTrueType.(intmod.IPrimitiveType).Flags() &
				expressionFalseType.(intmod.IPrimitiveType).Flags()

			if (flags & intmod.FlagInt) != 0 {
				typ = _type.PtTypeInt
			} else if (flags & intmod.FlagShort) != 0 {
				typ = _type.PtTypeShort
			} else if (flags & intmod.FlagChar) != 0 {
				typ = _type.PtTypeChar
			} else if (flags & intmod.FlagByte) != 0 {
				typ = _type.PtTypeByte
			} else {
				typ = _type.PtMaybeBooleanType
			}
		}
	} else if expressionTrueType.IsObjectType() && expressionFalseType.IsObjectType() {
		ot1 := expressionTrueType.(intmod.IObjectType)
		ot2 := expressionFalseType.(intmod.IObjectType)

		if m.typeMaker.IsAssignable(m.typeBounds, ot1, ot2) {
			typ = m.getTernaryOperatorExpressionType(ot1, ot2)
		} else if m.typeMaker.IsAssignable(m.typeBounds, ot2, ot1) {
			typ = m.getTernaryOperatorExpressionType(ot2, ot1)
		} else {
			typ = _type.OtTypeUndefinedObject
		}
	} else {
		typ = _type.OtTypeUndefinedObject
	}

	return modexp.NewTernaryOperatorExpressionWithAll(lineNumber, typ, condition, expressionTrue, expressionFalse)
}

func (m *StatementMaker) getTernaryOperatorExpressionType(ot1, ot2 intmod.IObjectType) intmod.IType {
	if ot1.TypeArguments() == nil {
		return ot1
	} else if ot2.TypeArguments() == nil {
		return ot1.CreateTypeWithArgs(nil)
	} else if ot1.IsTypeArgumentAssignableFrom(m.typeBounds, ot2) {
		return ot1
	} else if ot2.IsTypeArgumentAssignableFrom(m.typeBounds, ot1) {
		return ot1.CreateTypeWithArgs(ot2.TypeArguments())
	} else {
		return ot1.CreateTypeWithArgs(nil)
	}
}

func (m *StatementMaker) checkFieldReference(fieldName string, expression intmod.IExpression) bool {
	if expression.IsFieldReferenceExpression() {
		fre := expression.(intmod.IFieldReferenceExpression)
		return fre.Name() == fieldName && fre.InternalTypeName() == m.internalTypeName
	}

	return false
}

func (m *StatementMaker) createObjectTypeReferenceDotClassExpression(lineNumber int,
	fieldName string, mie intmod.IMethodInvocationExpression) intmod.IExpression {
	// Add SYNTHETIC flags to field
	m.memberVisitor.Init(fieldName)

	for _, field := range m.bodyDeclaration.FieldDeclarations() {
		field.FieldDeclarators().Accept(m.memberVisitor)
		if m.memberVisitor.Found() {
			field.SetFlags(field.Flags() | intcls.AccSynthetic)
			break
		}
	}

	// Add SYNTHETIC flags to method named 'class$'
	m.memberVisitor.Init("class$")

	for _, member := range m.bodyDeclaration.MethodDeclarations() {
		member.Accept(m.memberVisitor)
		if m.memberVisitor.Found() {
			member.SetFlags(member.Flags() | intcls.AccSynthetic)
			break
		}
	}

	typeName := mie.Parameters().StringValue()
	ot := m.typeMaker.MakeFromInternalTypeName(strings.ReplaceAll(typeName, ".", "/"))

	return modexp.NewTypeReferenceDotClassExpressionWithAll(lineNumber, ot)
}

func (m *StatementMaker) parseByteCode(basicBlock intsrv.IBasicBlock, statements intmod.IStatements) {
	m.byteCodeParser.Parse(basicBlock, statements, m.stack)
}

func (m *StatementMaker) replacePreOperatorWithPostOperator(statements intmod.IStatements) {
	iterator := statements.Iterator()

	for iterator.HasNext() {
		statement := iterator.Next()

		if statement.Expression().IsPreOperatorExpression() {
			poe := statement.Expression()
			operator := poe.Operator()

			if "++" == operator || "--" == operator {
				if statement.IsExpressionStatement() {
					es := statement.(intmod.IExpressionStatement)
					// Replace pre-operator statement with post-operator statement
					es.SetExpression(modexp.NewPostOperatorExpressionWithAll(
						poe.LineNumber(), operator, poe.Expression()))
				} else if statement.IsReturnExpressionStatement() {
					res := statement.(intmod.IReturnExpressionStatement)
					// Replace pre-operator statement with post-operator statement
					res.SetExpression(modexp.NewPostOperatorExpressionWithAll(
						poe.LineNumber(), operator, poe.Expression()))
				}
			}
		}
	}
}

func (m *StatementMaker) updateJumpStatements(jumps intmod.IStatements) {
	// assert false : "StatementMaker.updateJumpStatements(stmt) : 'jumps' list is not empty";

	iterator := jumps.Iterator()

	for iterator.HasNext() {
		statement := iterator.Next().(intsrv.IClassFileBreakContinueStatement)

		statement.SetStatement(
			modsts.NewCommentStatement(
				fmt.Sprintf("// Byte code: goto -> %d", statement.TargetOffset())))
	}
}

func NewStatementMakerMemberVisitor() *StatementMakerMemberVisitor {
	return &StatementMakerMemberVisitor{}
}

type StatementMakerMemberVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	name  string
	found bool
}

func (v *StatementMakerMemberVisitor) Init(name string) {
	v.name = name
	v.found = false
}

func (v *StatementMakerMemberVisitor) Found() bool {
	return v.found
}

func (v *StatementMakerMemberVisitor) VisitFieldDeclarator(declaration intmod.IFieldDeclarator) {
	v.found = v.found || declaration.Name() == v.name
}

func (v *StatementMakerMemberVisitor) VisitMethodDeclaration(declaration intmod.IMethodDeclaration) {
	v.found = v.found || declaration.Name() == v.name
}
