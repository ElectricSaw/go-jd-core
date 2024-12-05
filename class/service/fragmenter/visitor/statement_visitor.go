package visitor

import (
	"bitbucket.org/coontec/go-jd-core/class/api"
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/model/token"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

var Assert = token.NewKeywordToken("assert")
var Break = token.NewKeywordToken("break")
var Case = token.NewKeywordToken("case")
var Catch = token.NewKeywordToken("catch")
var Continue = token.NewKeywordToken("continue")
var Default = token.NewKeywordToken("default")
var Do = token.NewKeywordToken("do")
var Else = token.NewKeywordToken("else")
var Final = token.NewKeywordToken("final")
var Finally = token.NewKeywordToken("finally")
var For = token.NewKeywordToken("for")
var If = token.NewKeywordToken("if")
var Return = token.NewKeywordToken("return")
var Strict = token.NewKeywordToken("strictfp")
var Synchronized = token.NewKeywordToken("synchronized")
var Switch = token.NewKeywordToken("switch")
var Throw = token.NewKeywordToken("throw")
var Transient = token.NewKeywordToken("transient")
var Try = token.NewKeywordToken("try")
var Volatile = token.NewKeywordToken("volatile")
var While = token.NewKeywordToken("while")

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
	v.tokens = NewTokens(v);
	v.tokens.Add(token.StartDeclarationOrStatementBlock);
	v.tokens.Add(Assert);
	v.tokens.Add(token.Space);
	state.Condition().Accept(v);

	msg := state.Message();

	if (msg != nil) {
		v.tokens.Add(token.SpaceColonSpace);
		msg.Accept(v);
	}

	v.tokens.Add(token.Semicolon)
	v.tokens.Add(token.EndDeclarationOrStatementBlock);
	v.fragments.AddTokensFragment(v.tokens);
}

func (v *StatementVisitor) VisitBreakStatement(state intmod.IBreakStatement) {
	v.tokens = NewTokens(v);
	v.tokens.Add(Break);

	if (state.Text() != nil) {
		v.tokens.Add(token.Space);
		v.tokens.Add(v.newTextToken(state.Text()));
	}

	v.tokens.Add(token.Semicolon);
	v.fragments.AddTokensFragment(v.tokens);
}
func (v *StatementVisitor) VisitByteCodeStatement(state intmod.IByteCodeStatement) {
	v.visitComment(state.Text());
}
func (v *StatementVisitor) VisitCommentStatement(state intmod.ICommentStatement) {
	v.visitComment(state.Text());
}

func (v *StatementVisitor) visitComment(text string) {
	v.tokens = NewTokens(v);
	v.tokens.Add(token.StartComment);

	st := util.NewStringTokenizer2(text, "\n");

	for (st.HasMoreTokens()) {
		value, _ := st.NextToken()
		v.tokens.Add(token.NewTextToken(value));
		v.tokens.Add(token.NewLine1);
	}

	v.tokens.RemoveAt(v.tokens.Size()-1);
	v.tokens.Add(token.EndComment);
	v.fragments.AddTokensFragment(v.tokens);
}

func (v *StatementVisitor) VisitContinueStatement(state intmod.IContinueStatement) {
	v.tokens = NewTokens(v);
	v.tokens.Add(Continue);

	if (state.Text() != "") {
		v.tokens.Add(token.Space);
		v.tokens.Add(v.newTextToken(state.Text()));
	}

	v.tokens.Add(token.Semicolon);
	v.fragments.AddTokensFragment(v.tokens);
}

func (v *StatementVisitor) VisitDoWhileStatement(state intmod.IDoWhileStatement) {
	group := AddStartStatementsDoWhileBlock(v.fragments);

	v.SafeAcceptStatement(state.Statements());

	AddEndStatementsBlock(v.fragments, group);

	v.tokens = NewTokens(v);
	v.tokens.Add(For);
	v.tokens.Add(token.Space);
	v.tokens.Add(token.StartParametersBlock);

	state.Condition().Accept(v);

	v.tokens.Add(token.EndParametersBlock);
	v.tokens.Add(token.Semicolon);
	v.fragments.AddTokensFragment(v.tokens);
}

func (v *StatementVisitor) VisitExpressionStatement(state intmod.IExpressionStatement) {
	v.tokens = NewTokens(v);
	v.tokens.Add(token.StartDeclarationOrStatementBlock);

	state.Expression().Accept(v);

	v.tokens.Add(token.Semicolon);
	v.tokens.Add(token.EndDeclarationOrStatementBlock);
	v.fragments.AddTokensFragment(v.tokens);
}

func (v *StatementVisitor) VisitForStatement(state intmod.IForStatement) {
	v.tokens = NewTokens(v);
	v.tokens.Add(For);
	v.tokens.Add(token.Space);
	v.tokens.Add(token.StartParametersBlock);

	v.SafeAcceptDeclaration(state.Declaration());
	v.SafeAcceptExpression(state.Init());

	if (state.Condition() == nil) {
		v.tokens.Add(token.Semicolon);
	} else {
		v.tokens.Add(token.SemicolonSpace);
		state.Condition().Accept(v);
	}

	if (state.Update() == nil) {
		v.tokens.Add(token.Semicolon);
	} else {
		v.tokens.Add(token.SemicolonSpace);
		state.Update().Accept(v);
	}

	v.visitLoopStatements(state.Statements());
}

func (v *StatementVisitor) VisitForEachStatement(state intmod.IForEachStatement) {
	v.tokens = NewTokens(v);
	v.tokens.Add(For);
	v.tokens.Add(token.Space);
	v.tokens.Add(token.StartParametersBlock);

	typ := state.Type();
	typ.AcceptTypeVisitor(v);

	v.tokens.Add(token.Space);
	v.tokens.Add(v.newTextToken(state.Name()));
	v.tokens.Add(token.SpaceColonSpace);

	state.Expression().Accept(v);

	v.visitLoopStatements(state.Statements());
}

func (v *StatementVisitor) visitLoopStatements(state intmod.IStatement) {
	v.tokens.Add(token.EndParametersBlock);
	v.fragments.AddTokensFragment(v.tokens);

	if (state == nil) {
		v.tokens.Add(token.Semicolon);
	} else {
		tmp := v.fragments;
		v.fragments = NewFragments();

		state.Accept(v);

		switch (v.fragments.Size()) {
		case 0:
			v.tokens.Add(token.Semicolon);
		case 1:
			start := AddStartSingleStatementBlock(tmp);
			tmp.AddAll(v.fragments.ToSlice());
			AddEndSingleStatementBlock(tmp, start);
		default:
			group := AddStartStatementsBlock(tmp);
			tmp.AddAll(v.fragments.ToSlice());
			AddEndStatementsBlock(tmp, group);
		}

		v.fragments = tmp;
	}
}

func (v *StatementVisitor) VisitIfStatement(state intmod.IIfStatement) {
	v.tokens = NewTokens(v);
	v.tokens.Add(If);
	v.tokens.Add(token.Space);
	v.tokens.Add(token.StartParametersBlock);

	state.Condition().Accept(v);

	v.tokens.Add(token.EndParametersBlock);
	v.fragments.AddTokensFragment(v.tokens);
	stmt := state.Statements();

	if (stmt == nil) {
		v.fragments.Add(token.Semicolon);
	} else {
		tmp := v.fragments;
		v.fragments = NewFragments();

		stmt.Accept(v);

		switch (stmt.Size()) {
		case 0:
			tmp.Add(token.Semicolon);
		case 1:
			start := AddStartSingleStatementBlock(tmp);
			tmp.AddAll(v.fragments.ToSlice());
			AddEndSingleStatementBlock(tmp, start);
		default:
			group := AddStartStatementsBlock(tmp);
			tmp.AddAll(v.fragments.ToSlice());
			AddEndStatementsBlock(tmp, group);
		}

		v.fragments = tmp;
	}
}

func (v *StatementVisitor) VisitIfElseStatement(state intmod.IIfElseStatement) {
	v.tokens = NewTokens(v);
	v.tokens.Add(If);
	v.tokens.Add(token.Space);
	v.tokens.Add(token.StartParametersBlock);

	state.Condition().Accept(v);

	v.tokens.Add(token.EndParametersBlock);
	fragments.AddTokensFragment(v.tokens);

	StartStatementsBlockFragment.Group group = AddStartStatementsBlock(fragments);
	state.Statements().Accept(v);
	AddEndStatementsBlock(fragments, group);
	visitElseStatements(state.ElseStatements(), group);
}

protected void visitElseStatements(BaseStatement elseStatements, StartStatementsBlockFragment.Group group) {
BaseStatement statementList = elseStatements;

if (elseStatements.IsList()) {
if (elseStatements.Size() == 1) {
statementList = elseStatements.First();
}
}

v.tokens = NewTokens(v);
v.tokens.Add(ELSE);

if (statementList.IsIfElseStatement()) {
v.tokens.Add(token.Space);
v.tokens.Add(IF);
v.tokens.Add(token.Space);
v.tokens.Add(token.StartParametersBlock);

statementList.Condition().Accept(v);

v.tokens.Add(token.EndParametersBlock);
fragments.AddTokensFragment(v.tokens);

AddStartStatementsBlock(fragments, group);
statementList.Statements().Accept(v);
AddEndStatementsBlock(fragments, group);
visitElseStatements(statementList.ElseStatements(), group);
} else if (statementList.IsIfStatement()) {
v.tokens.Add(token.Space);
v.tokens.Add(IF);
v.tokens.Add(token.Space);
v.tokens.Add(token.StartParametersBlock);

statementList.Condition().Accept(v);

v.tokens.Add(token.EndParametersBlock);
fragments.AddTokensFragment(v.tokens);

AddStartStatementsBlock(fragments, group);

statementList.Statements().Accept(v);

AddEndStatementsBlock(fragments, group);
} else {
fragments.AddTokensFragment(v.tokens);

AddStartStatementsBlock(fragments, group);

elseStatements.Accept(v);

AddEndStatementsBlock(fragments, group);
}
}

func (v *StatementVisitor) VisitLabelStatement(state intmod.ILabelStatement) {
	v.tokens = NewTokens(v);
	v.tokens.Add(v.newTextToken(state.Text()));
	v.tokens.Add(token.COLON);

	if (state.Statement() == nil) {
		fragments.AddTokensFragment(v.tokens);
	} else {
		v.tokens.Add(token.Space);
		fragments.AddTokensFragment(v.tokens);
		state.Statement().Accept(v);
	}
}

func (v *StatementVisitor) VisitLambdaExpressionStatement(state intmod.ILambdaExpressionStatement) {
	state.Expression().Accept(v);
}

func (v *StatementVisitor) VisitLocalVariableDeclarationStatement(state intmod.ILocalVariableDeclarationStatement) {
	v.tokens = NewTokens(v);
	v.tokens.Add(token.StartDeclarationOrStatementBlock);

	if (state.IsFinal()) {
		v.tokens.Add(FINAL);
		v.tokens.Add(token.Space);
	}

	Basetyp typ = state.Type();

	typ.Accept(v);

	v.tokens.Add(token.Space);

	state.LocalVariableDeclarators().Accept(v);

	v.tokens.Add(token.Semicolon);
	v.tokens.Add(token.EndDeclarationOrStatementBlock);
	fragments.AddTokensFragment(v.tokens);
}

func (v *StatementVisitor) VisitReturnExpressionStatement(state intmod.IReturnExpressionStatement) {
	v.tokens = NewTokens(v);
	v.tokens.Add(token.StartDeclarationOrStatementBlock);
	v.tokens.AddLineNumberToken(state.LineNumber());
	v.tokens.Add(Return);
	v.tokens.Add(token.Space);

	state.Expression().Accept(v);

	v.tokens.Add(token.Semicolon);
	v.tokens.Add(token.EndDeclarationOrStatementBlock);
	fragments.AddTokensFragment(v.tokens);
}

func (v *StatementVisitor) VisitReturnStatement(state intmod.IReturnStatement) {
	fragments.Add(v.TokensFragment.Return_Semicolon);
}

func (v *StatementVisitor) VisitStatements(state intmod.IStatements) {
	int size = list.Size();

	if (size > 0) {
		Iterator<Statement> iterator = list.Iterator();
		iterator.Next().Accept(v);

		for (int i = 1; i < size; i++) {
			AddSpacerBetweenStatements(fragments);
			iterator.Next().Accept(v);
		}
	}
}

func (v *StatementVisitor) VisitSwitchStatement(state intmod.ISwitchStatement) {
	v.tokens = NewTokens(v);
	v.tokens.Add(SWITCH);
	v.tokens.Add(token.Space);
	v.tokens.Add(token.StartParametersBlock);

	state.Condition().Accept(v);

	v.tokens.Add(token.EndParametersBlock);
	fragments.AddTokensFragment(v.tokens);

	StartStatementsBlockFragment.Group group = AddStartStatementsBlock(fragments);

	Iterator<Switchstate.Block> iterator = state.Blocks().Iterator();

	if (iterator.hasNext()) {
		iterator.Next().Accept(v);

		for (iterator.hasNext()) {
			AddSpacerBetweenSwitchLabelBlock(fragments);
			iterator.Next().Accept(v);
		}
	}

	AddEndStatementsBlock(fragments, group);
	AddSpacerAfterEndStatementsBlock(fragments);
}

func (v *StatementVisitor) VisitSwitchStatement(state intmod.ISwitchStatement statement) {
	state.Label().Accept(v);
	AddSpacerAfterSwitchLabel(fragments);
	state.Statements().Accept(v);
	AddSpacerAfterSwitchBlock(fragments);
}

func (v *StatementVisitor) VisitSwitchStatement(state intmod.ISwitchStatement statement) {
	v.tokens = NewTokens(v);
	v.tokens.Add(DEFAULT);
	v.tokens.Add(token.COLON);
	fragments.AddTokensFragment(v.tokens);
}

func (v *StatementVisitor) VisitSwitchStatement(state intmod.ISwitchStatement statement) {
	v.tokens = NewTokens(v);
	v.tokens.Add(CASE);
	v.tokens.Add(token.Space);

	state.Expression().Accept(v);

	v.tokens.Add(token.COLON);
	fragments.AddTokensFragment(v.tokens);
}

func (v *StatementVisitor) VisitSwitchStatement(state intmod.ISwitchStatement statement) {
	Iterator<Switchstate.Label> iterator = state.Labels().Iterator();

	if (iterator.hasNext()) {
		iterator.Next().Accept(v);

		for (iterator.hasNext()) {
			AddSpacerBetweenSwitchLabels(fragments);
			iterator.Next().Accept(v);
		}
	}

	AddSpacerAfterSwitchLabel(fragments);
	state.Statements().Accept(v);
	AddSpacerAfterSwitchBlock(fragments);
}

func (v *StatementVisitor) VisitSynchronizedStatement(state intmod.ISynchronizedStatement) {
	v.tokens = NewTokens(v);
	v.tokens.Add(SYNCHRONIZED);
	v.tokens.Add(token.Space);
	v.tokens.Add(token.StartParametersBlock);

	state.Monitor().Accept(v);

	v.tokens.Add(token.EndParametersBlock);

	BaseStatement statements = state.Statements();

	if (statements == nil) {
		v.tokens.Add(token.Space);
		v.tokens.Add(token.LEFTRIGHTCURLYBRACKETS);
		fragments.AddTokensFragment(v.tokens);
	} else {
		fragments.AddTokensFragment(v.tokens);
		StartStatementsBlockFragment.Group group = AddStartStatementsBlock(fragments);
		statements.Accept(v);
		AddEndStatementsBlock(fragments, group);
	}
}

func (v *StatementVisitor) VisitThrowStatement(state intmod.IThrowStatement) {
	v.tokens = NewTokens(v);
	v.tokens.Add(token.StartDeclarationOrStatementBlock);
	v.tokens.Add(THROW);
	v.tokens.Add(token.Space);

	state.Expression().Accept(v);

	v.tokens.Add(token.Semicolon);
	v.tokens.Add(token.EndDeclarationOrStatementBlock);
	fragments.AddTokensFragment(v.tokens);
}

func (v *StatementVisitor) VisitTryStatement(state intmod.ITryStatement) {
	List<Trystate.Resource> resources = state.Resources();
	StartStatementsBlockFragment.Group group;

	if (resources == nil) {
		group = AddStartStatementsTryBlock(fragments);
	} else {
		int size = resources.Size();

		Assert size > 0;

		v.tokens = NewTokens(v);
		v.tokens.Add(TRY);
		if (size == 1) {
			v.tokens.Add(token.Space);
		}
		v.tokens.Add(token.StartRESOURCES_Block);

		resources.(0).Accept(v);

		for (int i=1; i<size; i++) {
			v.tokens.Add(token.Semicolon_Space);
			resources.(i).Accept(v);
		}

		v.tokens.Add(token.EndRESOURCES_Block);
		fragments.AddTokensFragment(v.tokens);
		group = AddStartStatementsBlock(fragments);
	}

	visitTryStatement(statement, group);
}

func (v *StatementVisitor) VisitTryStatement(state intmod.ITryStatement resource) {
	Expression expression = resource.Expression();

	v.tokens.AddLineNumberToken(expression);

	Basetyp typ = resource.Type();

	typ.Accept(v);
	v.tokens.Add(token.Space);
	v.tokens.Add(v.newTextToken(resource.Name()));
	v.tokens.Add(token.Space_EQUAL_Space);
	expression.Accept(v);
}

protected void visitTryStatement(TryStatement statement, StartStatementsBlockFragment.Group group) {
int fragmentCount1 = fragments.Size(), fragmentCount2 = fragmentCount1;

state.TryStatements().Accept(v);

if (state.CatchClauses() != nil) {
for (Trystate.CatchClause cc : state.CatchClauses()) {
AddEndStatementsBlock(fragments, group);

Basetyp typ = cc.Type();

v.tokens = NewTokens(v);
v.tokens.Add(CATCH);
v.tokens.Add(token.Space_LEFTROUNDBRACKET);
typ.Accept(v);

if (cc.Othertyps() != nil) {
for (Basetyp othertyp : cc.Othertyps()) {
v.tokens.Add(token.VERTICALLINE);
othertyp.Accept(v);
}
}

v.tokens.Add(token.Space);
v.tokens.Add(v.newTextToken(cc.Name()));
v.tokens.Add(token.RIGHTROUNDBRACKET);

int lineNumber = cc.LineNumber();

if (lineNumber == Expression.UNKNOWN_LINE_NUMBER) {
fragments.AddTokensFragment(v.tokens);
} else {
v.tokens.AddLineNumberToken(lineNumber);
fragments.AddTokensFragment(v.tokens);
}

fragmentCount1 = fragments.Size();
AddStartStatementsBlock(fragments, group);
fragmentCount2 = fragments.Size();
cc.Statements().Accept(v);
}
}

if (state.FinallyStatements() != nil) {
AddEndStatementsBlock(fragments, group);

v.tokens = NewTokens(v);
v.tokens.Add(FINALLY);
fragments.AddTokensFragment(v.tokens);

fragmentCount1 = fragments.Size();
AddStartStatementsBlock(fragments, group);
fragmentCount2 = fragments.Size();
state.FinallyStatements().Accept(v);
}

if (fragmentCount2 == fragments.Size()) {
fragments.SubList(fragmentCount1, fragmentCount2).clear();
v.tokens.Add(token.Space);
v.tokens.Add(token.LEFTRIGHTCURLYBRACKETS);
} else {
AddEndStatementsBlock(fragments, group);
}
}

func (v *StatementVisitor) VisittypDeclarationStatement(state intmod.ItypDeclarationStatement) {
	state.typDeclaration().Accept(v);
	fragments.Add(v.TokensFragment.Semicolon);
}

func (v *StatementVisitor) VisitforStatement(state intmod.IforStatement) {
	v.tokens = NewTokens(v);
	v.tokens.Add(For);
	v.tokens.Add(token.Space);
	v.tokens.Add(token.StartParametersBlock);

	state.Condition().Accept(v);

	v.visitLoopStatements(state.Statements());
}
