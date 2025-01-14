package visitor

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/api"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/service/fragmenter/visitor/fragutil"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewStatementVisitor(loader api.Loader, mainInternalTypeName string, majorVersion int,
	importsFragment intmod.IImportsFragment) *StatementVisitor {
	return &StatementVisitor{
		ExpressionVisitor: *NewExpressionVisitor(loader, mainInternalTypeName, majorVersion, importsFragment),
	}
}

type StatementVisitor struct {
	ExpressionVisitor
}

func (v *StatementVisitor) VisitAssertStatement(state intmod.IAssertStatement) {
	v.tokens = NewTokens(v)
	v.tokens.Add(model.StartDeclarationOrStatementBlock)
	v.tokens.Add(model.Assert)
	v.tokens.Add(model.Space)
	state.Condition().Accept(v)

	msg := state.Message()

	if msg != nil {
		v.tokens.Add(model.SpaceColonSpace)
		msg.Accept(v)
	}

	v.tokens.Add(model.Semicolon)
	v.tokens.Add(model.EndDeclarationOrStatementBlock)
	v.fragments.AddTokensFragment(v.tokens)
}

func (v *StatementVisitor) VisitBreakStatement(state intmod.IBreakStatement) {
	v.tokens = NewTokens(v)
	v.tokens.Add(model.Break)

	if state.Text() != "" {
		v.tokens.Add(model.Space)
		v.tokens.Add(v.newTextToken(state.Text()))
	}

	v.tokens.Add(model.Semicolon)
	v.fragments.AddTokensFragment(v.tokens)
}
func (v *StatementVisitor) VisitByteCodeStatement(state intmod.IByteCodeStatement) {
	v.visitComment(state.Text())
}
func (v *StatementVisitor) VisitCommentStatement(state intmod.ICommentStatement) {
	v.visitComment(state.Text())
}

func (v *StatementVisitor) visitComment(text string) {
	v.tokens = NewTokens(v)
	v.tokens.Add(model.StartComment)

	st := util.NewStringTokenizer2(text, "\n")

	for st.HasMoreTokens() {
		value, _ := st.NextToken()
		v.tokens.Add(model.NewTextToken(value))
		v.tokens.Add(model.NewLine1)
	}

	v.tokens.RemoveAt(v.tokens.Size() - 1)
	v.tokens.Add(model.EndComment)
	v.fragments.AddTokensFragment(v.tokens)
}

func (v *StatementVisitor) VisitContinueStatement(state intmod.IContinueStatement) {
	v.tokens = NewTokens(v)
	v.tokens.Add(model.Continue)

	if state.Text() != "" {
		v.tokens.Add(model.Space)
		v.tokens.Add(v.newTextToken(state.Text()))
	}

	v.tokens.Add(model.Semicolon)
	v.fragments.AddTokensFragment(v.tokens)
}

func (v *StatementVisitor) VisitDoWhileStatement(state intmod.IDoWhileStatement) {
	group := fragutil.AddStartStatementsDoWhileBlock(v.fragments)

	v.SafeAcceptStatement(state.Statements())

	fragutil.AddEndStatementsBlock(v.fragments, group)

	v.tokens = NewTokens(v)
	v.tokens.Add(model.For)
	v.tokens.Add(model.Space)
	v.tokens.Add(model.StartParametersBlock)

	state.Condition().Accept(v)

	v.tokens.Add(model.EndParametersBlock)
	v.tokens.Add(model.Semicolon)
	v.fragments.AddTokensFragment(v.tokens)
}

func (v *StatementVisitor) VisitExpressionStatement(state intmod.IExpressionStatement) {
	v.tokens = NewTokens(v)
	v.tokens.Add(model.StartDeclarationOrStatementBlock)

	state.Expression().Accept(v)

	v.tokens.Add(model.Semicolon)
	v.tokens.Add(model.EndDeclarationOrStatementBlock)
	v.fragments.AddTokensFragment(v.tokens)
}

func (v *StatementVisitor) VisitForStatement(state intmod.IForStatement) {
	v.tokens = NewTokens(v)
	v.tokens.Add(model.For)
	v.tokens.Add(model.Space)
	v.tokens.Add(model.StartParametersBlock)

	v.SafeAcceptDeclaration(state.Declaration())
	v.SafeAcceptExpression(state.Init())

	if state.Condition() == nil {
		v.tokens.Add(model.Semicolon)
	} else {
		v.tokens.Add(model.SemicolonSpace)
		state.Condition().Accept(v)
	}

	if state.Update() == nil {
		v.tokens.Add(model.Semicolon)
	} else {
		v.tokens.Add(model.SemicolonSpace)
		state.Update().Accept(v)
	}

	v.visitLoopStatements(state.Statements())
}

func (v *StatementVisitor) VisitForEachStatement(state intmod.IForEachStatement) {
	v.tokens = NewTokens(v)
	v.tokens.Add(model.For)
	v.tokens.Add(model.Space)
	v.tokens.Add(model.StartParametersBlock)

	typ := state.Type()
	typ.AcceptTypeVisitor(v)

	v.tokens.Add(model.Space)
	v.tokens.Add(v.newTextToken(state.Name()))
	v.tokens.Add(model.SpaceColonSpace)

	state.Expression().Accept(v)

	v.visitLoopStatements(state.Statements())
}

func (v *StatementVisitor) visitLoopStatements(state intmod.IStatement) {
	v.tokens.Add(model.EndParametersBlock)
	v.fragments.AddTokensFragment(v.tokens)

	if state == nil {
		v.tokens.Add(model.Semicolon)
	} else {
		tmp := v.fragments
		v.fragments = NewFragments()

		state.AcceptStatement(v)

		switch v.fragments.Size() {
		case 0:
			v.tokens.Add(model.Semicolon)
		case 1:
			start := fragutil.AddStartSingleStatementBlock(tmp)
			tmp.AddAll(v.fragments.ToSlice())
			fragutil.AddEndSingleStatementBlock(tmp, start)
		default:
			group := fragutil.AddStartStatementsBlock(tmp)
			tmp.AddAll(v.fragments.ToSlice())
			fragutil.AddEndStatementsBlock(tmp, group)
		}

		v.fragments = tmp
	}
}

func (v *StatementVisitor) VisitIfStatement(state intmod.IIfStatement) {
	v.tokens = NewTokens(v)
	v.tokens.Add(model.If)
	v.tokens.Add(model.Space)
	v.tokens.Add(model.StartParametersBlock)

	state.Condition().Accept(v)

	v.tokens.Add(model.EndParametersBlock)
	v.fragments.AddTokensFragment(v.tokens)
	stmt := state.Statements()

	if stmt == nil {
		v.fragments.Add(model.Semicolon)
	} else {
		tmp := v.fragments
		v.fragments = NewFragments()

		stmt.AcceptStatement(v)

		switch stmt.Size() {
		case 0:
			tmp.Add(model.Semicolon)
		case 1:
			start := fragutil.AddStartSingleStatementBlock(tmp)
			tmp.AddAll(v.fragments.ToSlice())
			fragutil.AddEndSingleStatementBlock(tmp, start)
		default:
			group := fragutil.AddStartStatementsBlock(tmp)
			tmp.AddAll(v.fragments.ToSlice())
			fragutil.AddEndStatementsBlock(tmp, group)
		}

		v.fragments = tmp
	}
}

func (v *StatementVisitor) VisitIfElseStatement(state intmod.IIfElseStatement) {
	v.tokens = NewTokens(v)
	v.tokens.Add(model.If)
	v.tokens.Add(model.Space)
	v.tokens.Add(model.StartParametersBlock)

	state.Condition().Accept(v)

	v.tokens.Add(model.EndParametersBlock)
	v.fragments.AddTokensFragment(v.tokens)

	group := fragutil.AddStartStatementsBlock(v.fragments)
	state.Statements().AcceptStatement(v)
	fragutil.AddEndStatementsBlock(v.fragments, group)
	v.visitElseStatements(state.ElseStatements(), group)
}

func (v *StatementVisitor) visitElseStatements(elseStatements intmod.IStatement, group intmod.IStartStatementsBlockFragmentGroup) {
	statementList := elseStatements

	if elseStatements.IsList() {
		if elseStatements.Size() == 1 {
			statementList = elseStatements.First()
		}
	}

	v.tokens = NewTokens(v)
	v.tokens.Add(model.Else)

	if statementList.IsIfElseStatement() {
		v.tokens.Add(model.Space)
		v.tokens.Add(model.If)
		v.tokens.Add(model.Space)
		v.tokens.Add(model.StartParametersBlock)

		statementList.Condition().Accept(v)

		v.tokens.Add(model.EndParametersBlock)
		v.fragments.AddTokensFragment(v.tokens)

		fragutil.AddStartStatementsBlock2(v.fragments, group)
		statementList.Statements().AcceptStatement(v)
		fragutil.AddEndStatementsBlock(v.fragments, group)
		v.visitElseStatements(statementList.ElseStatements(), group)
	} else if statementList.IsIfStatement() {
		v.tokens.Add(model.Space)
		v.tokens.Add(model.If)
		v.tokens.Add(model.Space)
		v.tokens.Add(model.StartParametersBlock)

		statementList.Condition().Accept(v)

		v.tokens.Add(model.EndParametersBlock)
		v.fragments.AddTokensFragment(v.tokens)

		fragutil.AddStartStatementsBlock2(v.fragments, group)

		statementList.Statements().AcceptStatement(v)

		fragutil.AddEndStatementsBlock(v.fragments, group)
	} else {
		v.fragments.AddTokensFragment(v.tokens)

		fragutil.AddStartStatementsBlock2(v.fragments, group)

		elseStatements.AcceptStatement(v)

		fragutil.AddEndStatementsBlock(v.fragments, group)
	}
}

func (v *StatementVisitor) VisitLabelStatement(state intmod.ILabelStatement) {
	v.tokens = NewTokens(v)
	v.tokens.Add(v.newTextToken(state.Text()))
	v.tokens.Add(model.Colon)

	if state.Statement() == nil {
		v.fragments.AddTokensFragment(v.tokens)
	} else {
		v.tokens.Add(model.Space)
		v.fragments.AddTokensFragment(v.tokens)
		state.Statement().AcceptStatement(v)
	}
}

func (v *StatementVisitor) VisitLambdaExpressionStatement(state intmod.ILambdaExpressionStatement) {
	state.Expression().Accept(v)
}

func (v *StatementVisitor) VisitLocalVariableDeclarationStatement(state intmod.ILocalVariableDeclarationStatement) {
	v.tokens = NewTokens(v)
	v.tokens.Add(model.StartDeclarationOrStatementBlock)

	if state.IsFinal() {
		v.tokens.Add(model.Final)
		v.tokens.Add(model.Space)
	}

	typ := state.Type()
	typ.AcceptTypeVisitor(v)

	v.tokens.Add(model.Space)

	state.LocalVariableDeclarators().AcceptDeclaration(v)

	v.tokens.Add(model.Semicolon)
	v.tokens.Add(model.EndDeclarationOrStatementBlock)
	v.fragments.AddTokensFragment(v.tokens)
}

func (v *StatementVisitor) VisitReturnExpressionStatement(state intmod.IReturnExpressionStatement) {
	v.tokens = NewTokens(v)
	v.tokens.Add(model.StartDeclarationOrStatementBlock)
	v.tokens.AddLineNumberTokenAt(state.LineNumber())
	v.tokens.Add(model.Return)
	v.tokens.Add(model.Space)

	state.Expression().Accept(v)

	v.tokens.Add(model.Semicolon)
	v.tokens.Add(model.EndDeclarationOrStatementBlock)
	v.fragments.AddTokensFragment(v.tokens)
}

func (v *StatementVisitor) VisitReturnStatement(_ intmod.IReturnStatement) {
	v.fragments.Add(model.ReturnSemicolon)
}

func (v *StatementVisitor) VisitStatements(list intmod.IStatements) {
	size := list.Size()

	if size > 0 {
		iterator := list.Iterator()
		iterator.Next().AcceptStatement(v)

		for i := 1; i < size; i++ {
			fragutil.AddSpacerBetweenStatements(v.fragments)
			iterator.Next().AcceptStatement(v)
		}
	}
}

func (v *StatementVisitor) VisitSwitchStatement(state intmod.ISwitchStatement) {
	v.tokens = NewTokens(v)
	v.tokens.Add(model.Switch)
	v.tokens.Add(model.Space)
	v.tokens.Add(model.StartParametersBlock)

	state.Condition().Accept(v)

	v.tokens.Add(model.EndParametersBlock)
	v.fragments.AddTokensFragment(v.tokens)

	group := fragutil.AddStartStatementsBlock(v.fragments)
	iterator := util.NewIteratorWithSlice(state.Blocks())

	if iterator.HasNext() {
		iterator.Next().AcceptStatement(v)

		for iterator.HasNext() {
			fragutil.AddSpacerBetweenSwitchLabelBlock(v.fragments)
			iterator.Next().AcceptStatement(v)
		}
	}

	fragutil.AddEndStatementsBlock(v.fragments, group)
	fragutil.AddSpacerAfterEndStatementsBlock(v.fragments)
}

func (v *StatementVisitor) VisitSwitchStatementLabelBlock(state intmod.ILabelBlock) {
	state.Label().AcceptStatement(v)
	fragutil.AddSpacerAfterSwitchLabel(v.fragments)
	state.Statements().AcceptStatement(v)
	fragutil.AddSpacerAfterSwitchBlock(v.fragments)
}

func (v *StatementVisitor) VisitSwitchStatementDefaultLabel(state intmod.IDefaultLabel) {
	v.tokens = NewTokens(v)
	v.tokens.Add(model.Default)
	v.tokens.Add(model.Colon)
	v.fragments.AddTokensFragment(v.tokens)
}

func (v *StatementVisitor) VisitSwitchStatementExpressionLabel(state intmod.IExpressionLabel) {
	v.tokens = NewTokens(v)
	v.tokens.Add(model.Case)
	v.tokens.Add(model.Space)

	state.Expression().Accept(v)

	v.tokens.Add(model.Colon)
	v.fragments.AddTokensFragment(v.tokens)
}

func (v *StatementVisitor) VisitSwitchStatementMultiLabelsBlock(state intmod.IMultiLabelsBlock) {
	iterator := util.NewIteratorWithSlice(state.Labels())

	if iterator.HasNext() {
		iterator.Next().AcceptStatement(v)

		for iterator.HasNext() {
			fragutil.AddSpacerBetweenSwitchLabels(v.fragments)
			iterator.Next().AcceptStatement(v)
		}
	}

	fragutil.AddSpacerAfterSwitchLabel(v.fragments)
	state.Statements().AcceptStatement(v)
	fragutil.AddSpacerAfterSwitchBlock(v.fragments)
}

func (v *StatementVisitor) VisitSynchronizedStatement(state intmod.ISynchronizedStatement) {
	v.tokens = NewTokens(v)
	v.tokens.Add(model.Synchronized)
	v.tokens.Add(model.Space)
	v.tokens.Add(model.StartParametersBlock)

	state.Monitor().Accept(v)

	v.tokens.Add(model.EndParametersBlock)

	statements := state.Statements()

	if statements == nil {
		v.tokens.Add(model.Space)
		v.tokens.Add(model.LeftRightCurlyBrackets)
		v.fragments.AddTokensFragment(v.tokens)
	} else {
		v.fragments.AddTokensFragment(v.tokens)
		group := fragutil.AddStartStatementsBlock(v.fragments)
		statements.AcceptStatement(v)
		fragutil.AddEndStatementsBlock(v.fragments, group)
	}
}

func (v *StatementVisitor) VisitThrowStatement(state intmod.IThrowStatement) {
	v.tokens = NewTokens(v)
	v.tokens.Add(model.StartDeclarationOrStatementBlock)
	v.tokens.Add(model.Throw)
	v.tokens.Add(model.Space)

	state.Expression().Accept(v)

	v.tokens.Add(model.Semicolon)
	v.tokens.Add(model.EndDeclarationOrStatementBlock)
	v.fragments.AddTokensFragment(v.tokens)
}

func (v *StatementVisitor) VisitTryStatement(state intmod.ITryStatement) {
	resources := util.NewDefaultListWithSlice(state.Resources())
	var group intmod.IStartStatementsBlockFragmentGroup

	if resources == nil {
		group = fragutil.AddStartStatementsTryBlock(v.fragments)
	} else {
		size := resources.Size()

		v.tokens = NewTokens(v)
		v.tokens.Add(model.Try)
		if size == 1 {
			v.tokens.Add(model.Space)
		}
		v.tokens.Add(model.StartResourcesBlock)

		resources.Get(0).AcceptStatement(v)

		for i := 1; i < size; i++ {
			v.tokens.Add(model.SemicolonSpace)
			resources.Get(i).AcceptStatement(v)
		}

		v.tokens.Add(model.EndResourcesBlock)
		v.fragments.AddTokensFragment(v.tokens)
		group = fragutil.AddStartStatementsBlock(v.fragments)
	}

	v.visitTryStatement(state, group)
}

func (v *StatementVisitor) VisitTryStatementResource(resource intmod.IResource) {
	expression := resource.Expression()
	v.tokens.AddLineNumberToken(expression)

	typ := resource.Type()
	typ.AcceptTypeVisitor(v)

	v.tokens.Add(model.Space)
	v.tokens.Add(v.newTextToken(resource.Name()))
	v.tokens.Add(model.SpaceEqualSpace)
	expression.Accept(v)
}

func (v *StatementVisitor) visitTryStatement(state intmod.ITryStatement, group intmod.IStartStatementsBlockFragmentGroup) {
	fragmentCount1 := v.fragments.Size()
	fragmentCount2 := fragmentCount1

	state.TryStatements().AcceptStatement(v)

	if state.CatchClauses() != nil {
		for _, cc := range state.CatchClauses() {
			fragutil.AddEndStatementsBlock(v.fragments, group)

			typ := cc.Type()

			v.tokens = NewTokens(v)
			v.tokens.Add(model.Catch)
			v.tokens.Add(model.SpaceLeftRoundBracket)
			typ.AcceptTypeVisitor(v)

			if cc.OtherType() != nil {
				for _, otherType := range cc.OtherType() {
					v.tokens.Add(model.VerticalLine)
					otherType.AcceptTypeVisitor(v)
				}
			}

			v.tokens.Add(model.Space)
			v.tokens.Add(v.newTextToken(cc.Name()))
			v.tokens.Add(model.RightRoundBracket)

			lineNumber := cc.LineNumber()

			if lineNumber == intmod.UnknownLineNumber {
				v.fragments.AddTokensFragment(v.tokens)
			} else {
				v.tokens.AddLineNumberTokenAt(lineNumber)
				v.fragments.AddTokensFragment(v.tokens)
			}

			fragmentCount1 = v.fragments.Size()
			fragutil.AddStartStatementsBlock2(v.fragments, group)
			fragmentCount2 = v.fragments.Size()
			cc.Statements().AcceptStatement(v)
		}
	}

	if state.FinallyStatements() != nil {
		fragutil.AddEndStatementsBlock(v.fragments, group)

		v.tokens = NewTokens(v)
		v.tokens.Add(model.Finally)
		v.fragments.AddTokensFragment(v.tokens)

		fragmentCount1 = v.fragments.Size()
		fragutil.AddStartStatementsBlock2(v.fragments, group)
		fragmentCount2 = v.fragments.Size()
		state.FinallyStatements().AcceptStatement(v)
	}

	if fragmentCount2 == v.fragments.Size() {
		v.fragments.SubList(fragmentCount1, fragmentCount2).Clear()
		v.tokens.Add(model.Space)
		v.tokens.Add(model.LeftRightCurlyBrackets)
	} else {
		fragutil.AddEndStatementsBlock(v.fragments, group)
	}
}

func (v *StatementVisitor) VisitTypeDeclarationStatement(state intmod.ITypeDeclarationStatement) {
	state.TypeDeclaration().AcceptDeclaration(v)
	v.fragments.Add(model.Semicolon)
}

func (v *StatementVisitor) VisitWhileStatement(state intmod.IWhileStatement) {
	v.tokens = NewTokens(v)
	v.tokens.Add(model.For)
	v.tokens.Add(model.Space)
	v.tokens.Add(model.StartParametersBlock)

	state.Condition().Accept(v)

	v.visitLoopStatements(state.Statements())
}
