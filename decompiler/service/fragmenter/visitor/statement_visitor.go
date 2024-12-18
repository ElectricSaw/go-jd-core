package visitor

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/api"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javafragment"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/token"
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
	v.tokens.Add(token.StartDeclarationOrStatementBlock)
	v.tokens.Add(token.Assert)
	v.tokens.Add(token.Space)
	state.Condition().Accept(v)

	msg := state.Message()

	if msg != nil {
		v.tokens.Add(token.SpaceColonSpace)
		msg.Accept(v)
	}

	v.tokens.Add(token.Semicolon)
	v.tokens.Add(token.EndDeclarationOrStatementBlock)
	v.fragments.AddTokensFragment(v.tokens)
}

func (v *StatementVisitor) VisitBreakStatement(state intmod.IBreakStatement) {
	v.tokens = NewTokens(v)
	v.tokens.Add(token.Break)

	if state.Text() != "" {
		v.tokens.Add(token.Space)
		v.tokens.Add(v.newTextToken(state.Text()))
	}

	v.tokens.Add(token.Semicolon)
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
	v.tokens.Add(token.StartComment)

	st := util.NewStringTokenizer2(text, "\n")

	for st.HasMoreTokens() {
		value, _ := st.NextToken()
		v.tokens.Add(token.NewTextToken(value))
		v.tokens.Add(token.NewLine1)
	}

	v.tokens.RemoveAt(v.tokens.Size() - 1)
	v.tokens.Add(token.EndComment)
	v.fragments.AddTokensFragment(v.tokens)
}

func (v *StatementVisitor) VisitContinueStatement(state intmod.IContinueStatement) {
	v.tokens = NewTokens(v)
	v.tokens.Add(token.Continue)

	if state.Text() != "" {
		v.tokens.Add(token.Space)
		v.tokens.Add(v.newTextToken(state.Text()))
	}

	v.tokens.Add(token.Semicolon)
	v.fragments.AddTokensFragment(v.tokens)
}

func (v *StatementVisitor) VisitDoWhileStatement(state intmod.IDoWhileStatement) {
	group := fragutil.AddStartStatementsDoWhileBlock(v.fragments)

	v.SafeAcceptStatement(state.Statements())

	fragutil.AddEndStatementsBlock(v.fragments, group)

	v.tokens = NewTokens(v)
	v.tokens.Add(token.For)
	v.tokens.Add(token.Space)
	v.tokens.Add(token.StartParametersBlock)

	state.Condition().Accept(v)

	v.tokens.Add(token.EndParametersBlock)
	v.tokens.Add(token.Semicolon)
	v.fragments.AddTokensFragment(v.tokens)
}

func (v *StatementVisitor) VisitExpressionStatement(state intmod.IExpressionStatement) {
	v.tokens = NewTokens(v)
	v.tokens.Add(token.StartDeclarationOrStatementBlock)

	state.Expression().Accept(v)

	v.tokens.Add(token.Semicolon)
	v.tokens.Add(token.EndDeclarationOrStatementBlock)
	v.fragments.AddTokensFragment(v.tokens)
}

func (v *StatementVisitor) VisitForStatement(state intmod.IForStatement) {
	v.tokens = NewTokens(v)
	v.tokens.Add(token.For)
	v.tokens.Add(token.Space)
	v.tokens.Add(token.StartParametersBlock)

	v.SafeAcceptDeclaration(state.Declaration())
	v.SafeAcceptExpression(state.Init())

	if state.Condition() == nil {
		v.tokens.Add(token.Semicolon)
	} else {
		v.tokens.Add(token.SemicolonSpace)
		state.Condition().Accept(v)
	}

	if state.Update() == nil {
		v.tokens.Add(token.Semicolon)
	} else {
		v.tokens.Add(token.SemicolonSpace)
		state.Update().Accept(v)
	}

	v.visitLoopStatements(state.Statements())
}

func (v *StatementVisitor) VisitForEachStatement(state intmod.IForEachStatement) {
	v.tokens = NewTokens(v)
	v.tokens.Add(token.For)
	v.tokens.Add(token.Space)
	v.tokens.Add(token.StartParametersBlock)

	typ := state.Type()
	typ.AcceptTypeVisitor(v)

	v.tokens.Add(token.Space)
	v.tokens.Add(v.newTextToken(state.Name()))
	v.tokens.Add(token.SpaceColonSpace)

	state.Expression().Accept(v)

	v.visitLoopStatements(state.Statements())
}

func (v *StatementVisitor) visitLoopStatements(state intmod.IStatement) {
	v.tokens.Add(token.EndParametersBlock)
	v.fragments.AddTokensFragment(v.tokens)

	if state == nil {
		v.tokens.Add(token.Semicolon)
	} else {
		tmp := v.fragments
		v.fragments = NewFragments()

		state.AcceptStatement(v)

		switch v.fragments.Size() {
		case 0:
			v.tokens.Add(token.Semicolon)
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
	v.tokens.Add(token.If)
	v.tokens.Add(token.Space)
	v.tokens.Add(token.StartParametersBlock)

	state.Condition().Accept(v)

	v.tokens.Add(token.EndParametersBlock)
	v.fragments.AddTokensFragment(v.tokens)
	stmt := state.Statements()

	if stmt == nil {
		v.fragments.Add(javafragment.Semicolon)
	} else {
		tmp := v.fragments
		v.fragments = NewFragments()

		stmt.AcceptStatement(v)

		switch stmt.Size() {
		case 0:
			tmp.Add(javafragment.Semicolon)
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
	v.tokens.Add(token.If)
	v.tokens.Add(token.Space)
	v.tokens.Add(token.StartParametersBlock)

	state.Condition().Accept(v)

	v.tokens.Add(token.EndParametersBlock)
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
	v.tokens.Add(token.Else)

	if statementList.IsIfElseStatement() {
		v.tokens.Add(token.Space)
		v.tokens.Add(token.If)
		v.tokens.Add(token.Space)
		v.tokens.Add(token.StartParametersBlock)

		statementList.Condition().Accept(v)

		v.tokens.Add(token.EndParametersBlock)
		v.fragments.AddTokensFragment(v.tokens)

		fragutil.AddStartStatementsBlock2(v.fragments, group)
		statementList.Statements().AcceptStatement(v)
		fragutil.AddEndStatementsBlock(v.fragments, group)
		v.visitElseStatements(statementList.ElseStatements(), group)
	} else if statementList.IsIfStatement() {
		v.tokens.Add(token.Space)
		v.tokens.Add(token.If)
		v.tokens.Add(token.Space)
		v.tokens.Add(token.StartParametersBlock)

		statementList.Condition().Accept(v)

		v.tokens.Add(token.EndParametersBlock)
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
	v.tokens.Add(token.Colon)

	if state.Statement() == nil {
		v.fragments.AddTokensFragment(v.tokens)
	} else {
		v.tokens.Add(token.Space)
		v.fragments.AddTokensFragment(v.tokens)
		state.Statement().AcceptStatement(v)
	}
}

func (v *StatementVisitor) VisitLambdaExpressionStatement(state intmod.ILambdaExpressionStatement) {
	state.Expression().Accept(v)
}

func (v *StatementVisitor) VisitLocalVariableDeclarationStatement(state intmod.ILocalVariableDeclarationStatement) {
	v.tokens = NewTokens(v)
	v.tokens.Add(token.StartDeclarationOrStatementBlock)

	if state.IsFinal() {
		v.tokens.Add(token.Final)
		v.tokens.Add(token.Space)
	}

	typ := state.Type()
	typ.AcceptTypeVisitor(v)

	v.tokens.Add(token.Space)

	state.LocalVariableDeclarators().AcceptDeclaration(v)

	v.tokens.Add(token.Semicolon)
	v.tokens.Add(token.EndDeclarationOrStatementBlock)
	v.fragments.AddTokensFragment(v.tokens)
}

func (v *StatementVisitor) VisitReturnExpressionStatement(state intmod.IReturnExpressionStatement) {
	v.tokens = NewTokens(v)
	v.tokens.Add(token.StartDeclarationOrStatementBlock)
	v.tokens.AddLineNumberTokenAt(state.LineNumber())
	v.tokens.Add(token.Return)
	v.tokens.Add(token.Space)

	state.Expression().Accept(v)

	v.tokens.Add(token.Semicolon)
	v.tokens.Add(token.EndDeclarationOrStatementBlock)
	v.fragments.AddTokensFragment(v.tokens)
}

func (v *StatementVisitor) VisitReturnStatement(_ intmod.IReturnStatement) {
	v.fragments.Add(javafragment.ReturnSemicolon)
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
	v.tokens.Add(token.Switch)
	v.tokens.Add(token.Space)
	v.tokens.Add(token.StartParametersBlock)

	state.Condition().Accept(v)

	v.tokens.Add(token.EndParametersBlock)
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
	v.tokens.Add(token.Default)
	v.tokens.Add(token.Colon)
	v.fragments.AddTokensFragment(v.tokens)
}

func (v *StatementVisitor) VisitSwitchStatementExpressionLabel(state intmod.IExpressionLabel) {
	v.tokens = NewTokens(v)
	v.tokens.Add(token.Case)
	v.tokens.Add(token.Space)

	state.Expression().Accept(v)

	v.tokens.Add(token.Colon)
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
	v.tokens.Add(token.Synchronized)
	v.tokens.Add(token.Space)
	v.tokens.Add(token.StartParametersBlock)

	state.Monitor().Accept(v)

	v.tokens.Add(token.EndParametersBlock)

	statements := state.Statements()

	if statements == nil {
		v.tokens.Add(token.Space)
		v.tokens.Add(token.LeftRightCurlyBrackets)
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
	v.tokens.Add(token.StartDeclarationOrStatementBlock)
	v.tokens.Add(token.Throw)
	v.tokens.Add(token.Space)

	state.Expression().Accept(v)

	v.tokens.Add(token.Semicolon)
	v.tokens.Add(token.EndDeclarationOrStatementBlock)
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
		v.tokens.Add(token.Try)
		if size == 1 {
			v.tokens.Add(token.Space)
		}
		v.tokens.Add(token.StartResourcesBlock)

		resources.Get(0).AcceptStatement(v)

		for i := 1; i < size; i++ {
			v.tokens.Add(token.SemicolonSpace)
			resources.Get(i).AcceptStatement(v)
		}

		v.tokens.Add(token.EndResourcesBlock)
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

	v.tokens.Add(token.Space)
	v.tokens.Add(v.newTextToken(resource.Name()))
	v.tokens.Add(token.SpaceEqualSpace)
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
			v.tokens.Add(token.Catch)
			v.tokens.Add(token.SpaceLeftRoundBracket)
			typ.AcceptTypeVisitor(v)

			if cc.OtherType() != nil {
				for _, otherType := range cc.OtherType() {
					v.tokens.Add(token.VerticalLine)
					otherType.AcceptTypeVisitor(v)
				}
			}

			v.tokens.Add(token.Space)
			v.tokens.Add(v.newTextToken(cc.Name()))
			v.tokens.Add(token.RightRoundBracket)

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
		v.tokens.Add(token.Finally)
		v.fragments.AddTokensFragment(v.tokens)

		fragmentCount1 = v.fragments.Size()
		fragutil.AddStartStatementsBlock2(v.fragments, group)
		fragmentCount2 = v.fragments.Size()
		state.FinallyStatements().AcceptStatement(v)
	}

	if fragmentCount2 == v.fragments.Size() {
		v.fragments.SubList(fragmentCount1, fragmentCount2).Clear()
		v.tokens.Add(token.Space)
		v.tokens.Add(token.LeftRightCurlyBrackets)
	} else {
		fragutil.AddEndStatementsBlock(v.fragments, group)
	}
}

func (v *StatementVisitor) VisitTypeDeclarationStatement(state intmod.ITypeDeclarationStatement) {
	state.TypeDeclaration().AcceptDeclaration(v)
	v.fragments.Add(javafragment.Semicolon)
}

func (v *StatementVisitor) VisitWhileStatement(state intmod.IWhileStatement) {
	v.tokens = NewTokens(v)
	v.tokens.Add(token.For)
	v.tokens.Add(token.Space)
	v.tokens.Add(token.StartParametersBlock)

	state.Condition().Accept(v)

	v.visitLoopStatements(state.Statements())
}
