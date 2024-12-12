package utils

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	modexp "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/expression"
	modsts "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/statement"
	_type "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/type"
	srvsts "github.com/ElectricSaw/go-jd-core/class/service/converter/model/javasyntax/statement"
	"github.com/ElectricSaw/go-jd-core/class/service/converter/visitor"
	"fmt"
	"strings"
)

var GlobalRemoveLastContinueStatementVisitor = visitor.NewRemoveLastContinueStatementVisitor()

func MakeLoopStatementMaker(
	majorVersion int, typeBounds map[string]intmod.IType, localVariableMaker intsrv.ILocalVariableMaker,
	loopBasicBlock intsrv.IBasicBlock, statements intmod.IStatements, condition intmod.IExpression, subStatements intmod.IStatements,
	jumps intmod.IStatements) intmod.IStatement {
	loop := makeLoopStatementMaker(majorVersion, typeBounds, localVariableMaker, loopBasicBlock, statements, condition, subStatements)
	continueOffset := loopBasicBlock.Sub1().FromOffset()
	breakOffset := loopBasicBlock.Next().FromOffset()

	if breakOffset <= 0 {
		breakOffset = loopBasicBlock.ToOffset()
	}

	return makeLabels(loopBasicBlock.Index(), continueOffset, breakOffset, loop, jumps)
}

func makeLoopStatementMaker(majorVersion int, typeBounds map[string]intmod.IType, localVariableMaker intsrv.ILocalVariableMaker,
	loopBasicBlock intsrv.IBasicBlock, statements intmod.IStatements, condition intmod.IExpression, subStatements intmod.IStatements) intmod.IStatement {
	forEachSupported := (majorVersion >= 49) // (majorVersion >= Java 5)

	subStatements.Accept(GlobalRemoveLastContinueStatementVisitor)

	if forEachSupported {
		statement := makeForEachArray(typeBounds, localVariableMaker, statements, condition, subStatements)

		if statement != nil {
			return statement
		}

		statement = makeForEachList(typeBounds, localVariableMaker, statements, condition, subStatements)

		if statement != nil {
			return statement
		}
	}

	lineNumber := intmod.UnknownLineNumber
	if condition != nil {
		lineNumber = condition.LineNumber()
	}
	subStatementsSize := subStatements.Size()

	switch subStatementsSize {
	case 0:
		if lineNumber > 0 {
			// Known line numbers
			init := extractInit(statements, lineNumber)

			if init != nil {
				return newClassFileForStatement(localVariableMaker, loopBasicBlock.FromOffset(),
					loopBasicBlock.ToOffset(), init, condition, nil, nil)
			}
		}
		break
	case 1:
		if lineNumber > 0 {
			// Known line numbers
			subStatement := subStatements.First()
			init := extractInit(statements, lineNumber)

			if subStatement.IsExpressionStatement() {
				subExpression := subStatement.Expression()

				if subExpression.LineNumber() == lineNumber {
					return newClassFileForStatement(
						localVariableMaker, loopBasicBlock.FromOffset(), loopBasicBlock.ToOffset(), init, condition, subExpression, nil)
				} else if init != nil {
					return newClassFileForStatement(
						localVariableMaker, loopBasicBlock.FromOffset(), loopBasicBlock.ToOffset(), init, condition, nil, subStatement)
				}
			} else if init != nil {
				return newClassFileForStatement(
					localVariableMaker, loopBasicBlock.FromOffset(), loopBasicBlock.ToOffset(), init, condition, nil, subStatement)
			}
		} else {
			// Unknown line numbers => Just try to find 'for (expression;;expression)'
			return createForStatementWithoutLineNumber(localVariableMaker, loopBasicBlock, statements, condition, subStatements)
		}
		break
	default:
		if lineNumber > 0 {
			// Known line numbers
			visit := visitor.NewSearchFirstLineNumberVisitor()

			subStatements.First().AcceptStatement(visit)

			firstLineNumber := visit.LineNumber()

			// Populates 'update'
			update := extractUpdate(subStatements, firstLineNumber)
			init := extractInit(statements, lineNumber)

			if (init != nil) || (update.Size() > 0) {
				return newClassFileForStatement(
					localVariableMaker, loopBasicBlock.FromOffset(), loopBasicBlock.ToOffset(), init, condition, update, subStatements)
			}
		} else {
			// Unknown line numbers => Just try to find 'for (expression;;expression)'
			return createForStatementWithoutLineNumber(localVariableMaker, loopBasicBlock, statements, condition, subStatements)
		}
		break
	}

	return modsts.NewWhileStatement(condition, subStatements)
}

func MakeLoopStatementMaker2(localVariableMaker intsrv.ILocalVariableMaker, loopBasicBlock intsrv.IBasicBlock,
	statements intmod.IStatements, subStatements intmod.IStatements, jumps intmod.IStatements) intmod.IStatement {
	subStatements.Accept(GlobalRemoveLastContinueStatementVisitor)

	loop := makeLoopStatementMaker2(localVariableMaker, loopBasicBlock, statements, subStatements)
	continueOffset := loopBasicBlock.Sub1().FromOffset()
	breakOffset := loopBasicBlock.Next().FromOffset()

	if breakOffset <= 0 {
		breakOffset = loopBasicBlock.ToOffset()
	}

	return makeLabels(loopBasicBlock.Index(), continueOffset, breakOffset, loop, jumps)
}

func makeLoopStatementMaker2(localVariableMaker intsrv.ILocalVariableMaker, loopBasicBlock intsrv.IBasicBlock,
	statements intmod.IStatements, subStatements intmod.IStatements) intmod.IStatement {
	subStatementsSize := subStatements.Size()

	if (subStatementsSize > 0) && (subStatements.Last() == modsts.Continue) {
		subStatements.RemoveLast()
		subStatementsSize--
	}

	switch subStatementsSize {
	case 0:
	case 1:
		subStatement := subStatements.First()
		if subStatement.IsExpressionStatement() {
			subExpression := subStatement.Expression()
			lineNumber := subExpression.LineNumber()

			if lineNumber > 0 {
				// Known line numbers
				init := extractInit(statements, lineNumber)

				if init != nil {
					return newClassFileForStatement(localVariableMaker, loopBasicBlock.FromOffset(),
						loopBasicBlock.ToOffset(), init, nil, subExpression, nil)
				}
			}
		}
		break
	default:
		visit := visitor.NewSearchFirstLineNumberVisitor()
		subStatements.Get(0).AcceptStatement(visit)
		firstLineNumber := visit.LineNumber()
		if firstLineNumber > 0 {
			// Populates 'update'
			update := extractUpdate(subStatements, firstLineNumber)

			if update.Size() > 0 {
				// Populates 'init'
				init := extractInit(statements, update.First().LineNumber())

				return newClassFileForStatement(localVariableMaker, loopBasicBlock.FromOffset(),
					loopBasicBlock.ToOffset(), init, nil, update, subStatements)
			}
		} else {
			// Unknown line numbers => Just try to find 'for (expression;;expression)'
			statement := createForStatementWithoutLineNumber(localVariableMaker,
				loopBasicBlock, statements, modexp.True, subStatements)

			if statement != nil {
				return statement
			}
		}
	}

	return modsts.NewWhileStatement(modexp.True, subStatements)
}

func MakeDoWhileLoop(loopBasicBlock intsrv.IBasicBlock, condition intmod.IExpression,
	subStatements intmod.IStatements, jumps intmod.IStatements) intmod.IStatement {
	subStatements.Accept(GlobalRemoveLastContinueStatementVisitor)

	loop := modsts.NewDoWhileStatement(condition, subStatements)
	continueOffset := loopBasicBlock.Sub1().FromOffset()
	breakOffset := loopBasicBlock.Next().FromOffset()

	if breakOffset <= 0 {
		breakOffset = loopBasicBlock.ToOffset()
	}

	return makeLabels(loopBasicBlock.Index(), continueOffset, breakOffset, loop, jumps)
}

func extractInit(statements intmod.IStatements, lineNumber int) intmod.IExpression {
	if lineNumber > 0 {
		switch statements.Size() {
		case 0:
			break
		case 1:
			statement := statements.First()
			if !statement.IsExpressionStatement() {
				break
			}

			expression := statement.Expression()
			if (expression.LineNumber() != lineNumber) ||
				!expression.IsBinaryOperatorExpression() ||
				(expression.RightExpression().LineNumber() == 0) {
				break
			}

			statements.Clear()
			return expression
		default:
			init := modexp.NewExpressions()
			iterator := statements.ListIterator()
			iterator.SetCursor(statements.Size())

			for iterator.HasPrevious() {
				statement := iterator.Previous()

				if !statement.IsExpressionStatement() {
					break
				}

				expression := statement.Expression()

				if (expression.LineNumber() != lineNumber) ||
					!expression.IsBinaryOperatorExpression() ||
					(expression.RightExpression().LineNumber() == 0) {
					break
				}

				init.Add(expression)
				_ = iterator.Remove()
			}

			if init.Size() > 0 {
				if init.Size() > 1 {
					init.Reverse()
				}
				return init
			}
			break
		}
	}

	return nil
}

func extractUpdate(statements intmod.IStatements, firstLineNumber int) intmod.IExpressions {
	update := modexp.NewExpressions()
	iterator := statements.ListIterator()
	iterator.SetCursor(statements.Size())

	// Populates 'update'
	for iterator.HasPrevious() {
		statement := iterator.Previous()
		if !statement.IsExpressionStatement() {
			break
		}
		expression := statement.Expression()
		if expression.LineNumber() >= firstLineNumber {
			break
		}
		_ = iterator.Remove()
		update.Add(expression)
	}

	if update.Size() > 1 {
		update.Reverse()
	}

	return update
}

func createForStatementWithoutLineNumber(localVariableMaker intsrv.ILocalVariableMaker,
	basicBlock intsrv.IBasicBlock, statements intmod.IStatements, condition intmod.IExpression,
	subStatements intmod.IStatements) intmod.IStatement {

	if !statements.IsEmpty() {
		init := statements.Last().Expression()

		if init.LeftExpression().IsLocalVariableReferenceExpression() {
			localVariable := convertTo(init.LeftExpression())
			update := subStatements.Last().Expression()
			var expression intmod.IExpression

			if update.IsBinaryOperatorExpression() {
				expression = update.LeftExpression()
			} else if update.IsPreOperatorExpression() {
				expression = update.Expression()
				update = modexp.NewPostOperatorExpressionWithAll(update.LineNumber(),
					update.Operator(), expression)
			} else if update.IsPostOperatorExpression() {
				expression = update.Expression()
			} else {
				return modsts.NewWhileStatement(condition, subStatements)
			}

			if expression.IsLocalVariableReferenceExpression() && convertTo(expression) == localVariable {
				statements.RemoveLast()
				subStatements.RemoveLast()

				if condition == modexp.True {
					condition = nil
				}

				return newClassFileForStatement(localVariableMaker, basicBlock.FromOffset(),
					basicBlock.ToOffset(), init, condition, update, subStatements)
			}
		}
	}

	return modsts.NewWhileStatement(condition, subStatements)
}

func makeForEachArray(
	typeBounds map[string]intmod.IType, localVariableMaker intsrv.ILocalVariableMaker, statements intmod.IStatements,
	condition intmod.IExpression, subStatements intmod.IStatements) intmod.IStatement {
	if condition == nil {
		return nil
	}

	statementsSize := statements.Size()

	if (statementsSize < 3) || (subStatements.Size() < 2) {
		return nil
	}

	// len$ = arr$.length;
	statement := statements.Get(statementsSize - 2)

	if !statement.IsExpressionStatement() {
		return nil
	}

	lineNumber := condition.LineNumber()
	expression := statement.Expression()

	if (expression.LineNumber() != lineNumber) || !expression.IsBinaryOperatorExpression() {
		return nil
	}

	if !expression.RightExpression().IsLengthExpression() || !expression.LeftExpression().IsLocalVariableReferenceExpression() {
		return nil
	}

	boe := expression
	expression = boe.RightExpression().Expression()

	if !expression.IsLocalVariableReferenceExpression() {
		return nil
	}

	syntheticArray := convertTo(expression)
	syntheticLength := convertTo(boe.LeftExpression())

	if (syntheticArray.Name() != "") && strings.Index(syntheticArray.Name(), "$") == -1 {
		return nil
	}

	// String s = arr$[i$];
	expression = subStatements.First().Expression()

	if !expression.RightExpression().IsArrayExpression() ||
		!expression.LeftExpression().IsLocalVariableReferenceExpression() ||
		(expression.LineNumber() != condition.LineNumber()) {
		return nil
	}

	arrayExpression := expression.RightExpression().(intmod.IArrayExpression)

	if !arrayExpression.Expression().IsLocalVariableReferenceExpression() ||
		!arrayExpression.Index().IsLocalVariableReferenceExpression() {
		return nil
	}
	if convertTo(arrayExpression.Expression()) != syntheticArray {
		return nil
	}

	syntheticIndex := convertTo(arrayExpression.Index())
	item := convertTo(expression.LeftExpression())

	if syntheticIndex.Name() != "" && strings.Index(syntheticIndex.Name(), "$") == -1 {
		return nil
	}

	// arr$ = array;
	expression = statements.Get(statementsSize - 3).Expression()

	if !expression.LeftExpression().IsLocalVariableReferenceExpression() {
		return nil
	}
	if convertTo(expression.LeftExpression()) != syntheticArray {
		return nil
	}

	arrayType := expression.RightExpression().Type()
	array := expression.RightExpression()

	// i$ = 0;
	expression = statements.Get(statementsSize - 1).Expression()

	if (expression.LineNumber() != lineNumber) ||
		!expression.LeftExpression().IsLocalVariableReferenceExpression() ||
		!expression.RightExpression().IsIntegerConstantExpression() {
		return nil
	}
	if (expression.RightExpression().IntegerValue() != 0) ||
		convertTo(expression.LeftExpression()) != syntheticIndex {
		return nil
	}

	// i$ < len$;
	if !condition.LeftExpression().IsLocalVariableReferenceExpression() ||
		!condition.RightExpression().IsLocalVariableReferenceExpression() {
		return nil
	}

	if convertTo(condition.LeftExpression()) != syntheticIndex ||
		convertTo(condition.RightExpression()) != syntheticLength {
		return nil
	}

	// ++i$;
	expression = subStatements.Last().Expression()
	if expression.LineNumber() != lineNumber || !expression.IsPostOperatorExpression() {
		return nil
	}

	// Found
	statements.RemoveLast()
	statements.RemoveLast()
	statements.RemoveLast()

	subStatements.RemoveFirst()
	subStatements.RemoveLast()

	item.SetDeclared(true)
	typ := arrayType.CreateType(arrayType.Dimension() - 1)

	if _type.OtTypeObject == item.Type() {
		item.(intsrv.IObjectLocalVariable).SetType(typeBounds, typ)
	} else if item.Type().IsGenericType() {
		item.(intsrv.IGenericLocalVariable).SetType(typ.(intmod.IGenericType))
	} else {
		item.TypeOnRight(typeBounds, typ)
	}

	localVariableMaker.RemoveLocalVariable(syntheticArray)
	localVariableMaker.RemoveLocalVariable(syntheticIndex)
	localVariableMaker.RemoveLocalVariable(syntheticLength)

	return srvsts.NewClassFileForEachStatement(item, array, subStatements)
}

func makeForEachList(
	typeBounds map[string]intmod.IType, localVariableMaker intsrv.ILocalVariableMaker, statements intmod.IStatements,
	condition intmod.IExpression, subStatements intmod.IStatements) intmod.IStatement {
	if condition == nil {
		return nil
	}

	if (statements.Size() < 1) || (subStatements.Size() < 1) {
		return nil
	}

	// i$.hasNext();
	if !condition.IsMethodInvocationExpression() {
		return nil
	}

	mie := condition.(intmod.IMethodInvocationExpression)

	if mie.Name() != "hasNext" || mie.InternalTypeName() != "java/util/Iterator" ||
		!mie.Expression().IsLocalVariableReferenceExpression() {
		return nil
	}

	syntheticIterator := convertTo(mie.Expression())

	// Iterator i$ = list.Iterator();
	boe := statements.Last().Expression()

	if (boe == nil) ||
		!boe.LeftExpression().IsLocalVariableReferenceExpression() ||
		!boe.RightExpression().IsMethodInvocationExpression() ||
		(boe.LineNumber() != condition.LineNumber()) {
		return nil
	}

	mie = boe.RightExpression().(intmod.IMethodInvocationExpression)

	if mie.Name() != "iterator" || mie.Descriptor() != "()Ljava/util/Iterator;" {
		return nil
	}
	if convertTo(boe.LeftExpression()) != syntheticIterator {
		return nil
	}

	list := mie.Expression()
	if list.IsCastExpression() {
		list = list.Expression()
	}

	// String s = (String)i$.next();
	boe = subStatements.First().Expression()

	if !boe.LeftExpression().IsLocalVariableReferenceExpression() {
		return nil
	}

	expression := boe.RightExpression()

	if boe.RightExpression().IsCastExpression() {
		expression = expression.Expression()
	}
	if !expression.IsMethodInvocationExpression() {
		return nil
	}

	mie = expression.(intmod.IMethodInvocationExpression)
	if mie.Name() != "next" ||
		mie.InternalTypeName() != "java/util/Iterator" ||
		!mie.Expression().IsLocalVariableReferenceExpression() {
		return nil
	}

	if convertTo(mie.Expression()) != syntheticIterator {
		return nil
	}

	// Check if 'i$' is not used in sub-statements
	visitor1 := visitor.NewSearchLocalVariableReferenceVisitor()
	visitor1.Init(syntheticIterator.Index())

	length := subStatements.Size()
	for i := 1; i < length; i++ {
		subStatements.Get(i).AcceptStatement(visitor1)
	}

	if visitor1.ContainsReference() {
		return nil
	}

	// Found
	item := convertTo(boe.LeftExpression())

	statements.RemoveLast()
	subStatements.RemoveAt(0)
	item.SetDeclared(true)

	if len(syntheticIterator.References()) == 3 {
		// If local variable is used only for this loop
		localVariableMaker.RemoveLocalVariable(syntheticIterator)
	}

	typ := list.Type()

	if typ.IsObjectType() {
		listType := typ.(intmod.IObjectType)

		if listType.TypeArguments() == nil {
			if _type.OtTypeObject != item.Type() {
				if list.IsNewExpression() {
					ne := list.(intsrv.IClassFileNewExpression)
					ne.SetType(listType.CreateTypeWithArgs(_type.Diamond))
				} else {
					list = modexp.NewCastExpression(_type.OtTypeIterable.CreateTypeWithArgs(item.Type()), list)
				}
			}
		} else {
			visitor2 := visitor.NewCreateTypeFromTypeArgumentVisitor()
			listType.TypeArguments().AcceptTypeArgumentVisitor(visitor2)
			typ = visitor2.Type()

			if typ != nil {
				if _type.OtTypeObject == item.Type() {
					item.(intsrv.IObjectLocalVariable).SetType(typeBounds, typ)
				} else if item.Type().IsGenericType() {
					item.(intsrv.IGenericLocalVariable).SetType(typ.(intmod.IGenericType))
				} else {
					item.TypeOnRight(typeBounds, typ)
				}
			}
		}
	}

	return srvsts.NewClassFileForEachStatement(item, list, subStatements)
}

func makeLabels(loopIndex, continueOffset, breakOffset int, loop intmod.IStatement, jumps intmod.IStatements) intmod.IStatement {
	if !jumps.IsEmpty() {
		iterator := jumps.Iterator()
		label := fmt.Sprintf("label%d", loopIndex)
		createLabel := false

		for iterator.HasNext() {
			statement := iterator.Next().(intsrv.IClassFileBreakContinueStatement)
			offset := statement.Offset()
			targetOffset := statement.TargetOffset()

			if targetOffset == continueOffset {
				statement.SetStatement(modsts.NewContinueStatement(label))
				createLabel = true
				_ = iterator.Remove()
			} else if targetOffset == breakOffset {
				statement.SetStatement(modsts.NewBreakStatement(label))
				createLabel = true
				_ = iterator.Remove()
			} else if (continueOffset <= offset) && (offset < breakOffset) {
				if (continueOffset <= targetOffset) && (targetOffset < breakOffset) {
					if statement.IsContinueLabel() {
						statement.SetStatement(modsts.NewContinueStatement(label))
						createLabel = true
					} else {
						statement.SetStatement(modsts.Continue)
					}
					_ = iterator.Remove()
				} else {
					statement.SetContinueLabel(true)
				}
			}
		}

		if createLabel {
			return modsts.NewLabelStatement(label, loop)
		}
	}

	return loop
}

func newClassFileForStatement(localVariableMaker intsrv.ILocalVariableMaker, fromOffset, toOffset int,
	init intmod.IExpression, condition intmod.IExpression, update intmod.IExpression,
	statements intmod.IStatement) intsrv.IClassFileForStatement {
	if init != nil {
		visit := visitor.NewSearchFromOffsetVisitor()
		init.Accept(visit)
		offset := visit.Offset()
		if fromOffset > offset {
			fromOffset = offset
		}
	}

	visit := visitor.NewChangeFrameOfLocalVariablesVisitor(localVariableMaker)

	if condition != nil {
		condition.Accept(visit)
	}
	if update != nil {
		update.Accept(visit)
	}

	return srvsts.NewClassFileForStatement(fromOffset, toOffset, init, condition, update, statements)
}

func convertTo(expr intmod.IExpression) intsrv.ILocalVariable {
	return expr.(intsrv.IClassFileLocalVariableReferenceExpression).LocalVariable().(intsrv.ILocalVariable)
}
