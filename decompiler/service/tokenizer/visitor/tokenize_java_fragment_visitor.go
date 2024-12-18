package visitor

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/api"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/token"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

var Do = token.NewKeywordToken("do")
var Import = token.NewKeywordToken("import")
var For = token.NewKeywordToken("for")
var True = token.NewKeywordToken("true")
var Try = token.NewKeywordToken("try")
var While = token.NewKeywordToken("while")

var DoTokens = util.NewDefaultListWithSlice[intmod.IToken]([]intmod.IToken{Do})
var EmptyForTokens = util.NewDefaultListWithSlice[intmod.IToken]([]intmod.IToken{For, token.InfiniteFor})
var EmptyWhileTokens = util.NewDefaultListWithSlice[intmod.IToken]([]intmod.IToken{While,
	token.Space, token.LeftRoundBracket, True, token.RightRoundBracket})
var TryTokens = util.NewDefaultListWithSlice[intmod.IToken]([]intmod.IToken{Try})

func NewTokenizeJavaFragmentVisitor(capacity int) *TokenizeJavaFragmentVisitor {
	return &TokenizeJavaFragmentVisitor{
		tokens: util.NewDefaultListWithCapacity[intmod.IToken](capacity),
	}
}

type TokenizeJavaFragmentVisitor struct {
	knownLineNumberTokenVisitor   *KnownLineNumberTokenVisitor
	unknownLineNumberTokenVisitor *UnknownLineNumberTokenVisitor
	tokens                        util.IList[intmod.IToken]
}

func (v *TokenizeJavaFragmentVisitor) Tokens() util.IList[intmod.IToken] {
	return v.tokens
}

func (v *TokenizeJavaFragmentVisitor) VisitEndBlockFragment(fragment intmod.IEndBlockFragment) {
	switch fragment.LineCount() {
	case 0:
		v.tokens.Add(token.Space)
		v.tokens.Add(token.EndBlock)
	case 1:
		v.tokens.Add(token.NewLine1)
		v.tokens.Add(token.EndBlock)
	case 2:
		v.tokens.Add(token.NewLine2)
		v.tokens.Add(token.EndBlock)
	default:
		v.tokens.Add(token.NewNewLineToken(fragment.LineCount()))
		v.tokens.Add(token.EndBlock)
	}
}

func (v *TokenizeJavaFragmentVisitor) VisitEndBlockInParameterFragment(fragment intmod.IEndBlockInParameterFragment) {
	switch fragment.LineCount() {
	case 0:
		v.tokens.Add(token.Space)
		v.tokens.Add(token.EndBlock)
		v.tokens.Add(token.Comma)
	case 1:
		v.tokens.Add(token.NewLine1)
		v.tokens.Add(token.EndBlock)
		v.tokens.Add(token.Comma)
	case 2:
		v.tokens.Add(token.NewLine2)
		v.tokens.Add(token.EndBlock)
		v.tokens.Add(token.Comma)
	default:
		v.tokens.Add(token.NewNewLineToken(fragment.LineCount()))
		v.tokens.Add(token.EndBlock)
		v.tokens.Add(token.Comma)
	}
}

func (v *TokenizeJavaFragmentVisitor) VisitEndBodyFragment(fragment intmod.IEndBodyFragment) {
	switch fragment.LineCount() {
	case 0:
		v.tokens.Add(token.Space)
		v.tokens.Add(token.EndBlock)
	case 1:
		if fragment.Start().LineCount() == 0 {
			v.tokens.Add(token.Space)
			v.tokens.Add(token.EndBlock)
			v.tokens.Add(token.NewLine1)
		} else {
			v.tokens.Add(token.NewLine1)
			v.tokens.Add(token.EndBlock)
		}
	case 2:
		v.tokens.Add(token.NewLine1)
		v.tokens.Add(token.EndBlock)
		v.tokens.Add(token.NewLine1)
	default:
		v.tokens.Add(token.NewNewLineToken(fragment.LineCount() - 1))
		v.tokens.Add(token.EndBlock)
		v.tokens.Add(token.NewLine1)
	}
}

func (v *TokenizeJavaFragmentVisitor) VisitEndBodyInParameterFragment(fragment intmod.IEndBodyInParameterFragment) {
	switch fragment.LineCount() {
	case 0:
		v.tokens.Add(token.Space)
		v.tokens.Add(token.EndBlock)
		v.tokens.Add(token.Comma)
	case 1:
		v.tokens.Add(token.NewLine1)
		v.tokens.Add(token.EndBlock)
		v.tokens.Add(token.Comma)
		v.tokens.Add(token.Space)
	case 2:
		v.tokens.Add(token.NewLine1)
		v.tokens.Add(token.EndBlock)
		v.tokens.Add(token.Comma)
		v.tokens.Add(token.NewLine1)
	default:
		v.tokens.Add(token.NewNewLineToken(fragment.LineCount() - 1))
		v.tokens.Add(token.EndBlock)
		v.tokens.Add(token.Comma)
		v.tokens.Add(token.NewLine1)
	}
}

func (v *TokenizeJavaFragmentVisitor) VisitEndMovableJavaBlockFragment(_ intmod.IEndMovableJavaBlockFragment) {
}

func (v *TokenizeJavaFragmentVisitor) VisitEndSingleStatementBlockFragment(fragment intmod.IEndSingleStatementBlockFragment) {
	switch fragment.LineCount() {
	case 0:
		switch fragment.Start().LineCount() {
		case 0, 1:
			v.tokens.Add(token.Space)
			v.tokens.Add(token.EndDeclarationOrStatementBlock)
		default:
			v.tokens.Add(token.Space)
			v.tokens.Add(token.EndBlock)
			v.tokens.Add(token.Space)
		}
	case 1:
		switch fragment.Start().LineCount() {
		case 0:
			v.tokens.Add(token.EndDeclarationOrStatementBlock)
			v.tokens.Add(token.NewLine1)
		default:
			v.tokens.Add(token.NewLine1)
			v.tokens.Add(token.EndBlock)
		}
	case 2:
		switch fragment.Start().LineCount() {
		case 0:
			v.tokens.Add(token.EndDeclarationOrStatementBlock)
			v.tokens.Add(token.NewLine2)
		default:
			v.tokens.Add(token.NewLine1)
			v.tokens.Add(token.EndBlock)
			v.tokens.Add(token.NewLine1)
		}
	default:
		switch fragment.Start().LineCount() {
		case 0:
			v.tokens.Add(token.EndDeclarationOrStatementBlock)
			v.tokens.Add(token.NewNewLineToken(fragment.LineCount()))
		default:
			v.tokens.Add(token.NewLine1)
			v.tokens.Add(token.EndBlock)
			v.tokens.Add(token.NewNewLineToken(fragment.LineCount() - 1))
		}
	}
}

func (v *TokenizeJavaFragmentVisitor) VisitEndStatementsBlockFragment(fragment intmod.IEndStatementsBlockFragment) {
	minimalLineCount := fragment.Group().MinimalLineCount()

	switch fragment.LineCount() {
	case 0:
		v.tokens.Add(token.Space)
		v.tokens.Add(token.EndBlock)
		v.tokens.Add(token.Space)
	case 1:
		if minimalLineCount == 0 {
			v.tokens.Add(token.Space)
			v.tokens.Add(token.EndBlock)
			v.tokens.Add(token.NewLine1)
		} else {
			v.tokens.Add(token.NewLine1)
			v.tokens.Add(token.EndBlock)
			v.tokens.Add(token.Space)
		}
	case 2:
		switch minimalLineCount {
		case 0:
			v.tokens.Add(token.Space)
			v.tokens.Add(token.EndBlock)
			v.tokens.Add(token.NewLine2)
		case 1:
			v.tokens.Add(token.NewLine1)
			v.tokens.Add(token.EndBlock)
			v.tokens.Add(token.NewLine1)
		default:
			v.tokens.Add(token.NewNewLineToken(fragment.LineCount() - 1))
			v.tokens.Add(token.EndBlock)
			v.tokens.Add(token.NewLine1)
		}
	default:
		switch minimalLineCount {
		case 0:
			v.tokens.Add(token.Space)
			v.tokens.Add(token.EndBlock)
			v.tokens.Add(token.NewNewLineToken(fragment.LineCount()))
		case 1:
			v.tokens.Add(token.NewNewLineToken(fragment.LineCount()))
			v.tokens.Add(token.EndBlock)
		default:
			v.tokens.Add(token.NewNewLineToken(fragment.LineCount() - 1))
			v.tokens.Add(token.EndBlock)
			v.tokens.Add(token.NewLine1)
		}
	}
}

func (v *TokenizeJavaFragmentVisitor) VisitImportsFragment(fragment intmod.IImportsFragment) {
	imports := util.NewDefaultListWithSlice(fragment.Imports())
	imports.Sort(func(i, j int) bool {
		return imports.ToSlice()[i].QualifiedName() < imports.ToSlice()[j].QualifiedName()
	})

	v.tokens.Add(token.StartImportStatementsToken)

	for _, imp := range imports.ToSlice() {
		v.tokens.Add(Import)
		v.tokens.Add(token.Space)
		v.tokens.Add(token.NewReferenceToken(intmod.TypeToken, imp.InternalName(),
			imp.QualifiedName(), "", ""))
		v.tokens.Add(token.Semicolon)
		v.tokens.Add(token.NewLine1)
	}

	v.tokens.Add(token.EndImportStatements)
}

func (v *TokenizeJavaFragmentVisitor) VisitLineNumberTokensFragment(fragment intmod.ILineNumberTokensFragment) {
	v.knownLineNumberTokenVisitor.Reset(fragment.FirstLineNumber())

	for _, tkn := range fragment.Tokens() {
		tkn.Accept(v.knownLineNumberTokenVisitor)
	}
}

func (v *TokenizeJavaFragmentVisitor) VisitSpacerBetweenMembersFragment(fragment intmod.ISpacerBetweenMembersFragment) {
	switch fragment.LineCount() {
	case 0:
		v.tokens.Add(token.Space)
	case 1:
		v.tokens.Add(token.NewLine1)
	case 2:
		v.tokens.Add(token.NewLine2)
	default:
		v.tokens.Add(token.NewNewLineToken(fragment.LineCount()))
	}
}

func (v *TokenizeJavaFragmentVisitor) VisitSpacerFragment(fragment intmod.ISpacerFragment) {
	switch fragment.LineCount() {
	case 0:
	case 1:
		v.tokens.Add(token.NewLine1)
	case 2:
		v.tokens.Add(token.NewLine2)
	default:
		v.tokens.Add(token.NewNewLineToken(fragment.LineCount()))
	}
}

func (v *TokenizeJavaFragmentVisitor) VisitSpaceSpacerFragment(fragment intmod.ISpaceSpacerFragment) {
	switch fragment.LineCount() {
	case 0:
		v.tokens.Add(token.Space)
	case 1:
		v.tokens.Add(token.NewLine1)
	case 2:
		v.tokens.Add(token.NewLine2)
	default:
		v.tokens.Add(token.NewNewLineToken(fragment.LineCount()))
	}
}

func (v *TokenizeJavaFragmentVisitor) VisitStartBlockFragment(fragment intmod.IStartBlockFragment) {
	switch fragment.LineCount() {
	case 0:
		v.tokens.Add(token.StartBlock)
		v.tokens.Add(token.Space)
	case 1:
		v.tokens.Add(token.StartBlock)
		v.tokens.Add(token.NewLine1)
	case 2:
		v.tokens.Add(token.StartBlock)
		v.tokens.Add(token.NewLine2)
	default:
		v.tokens.Add(token.StartBlock)
		v.tokens.Add(token.NewNewLineToken(fragment.LineCount()))
	}
}

func (v *TokenizeJavaFragmentVisitor) VisitStartBodyFragment(fragment intmod.IStartBodyFragment) {
	switch fragment.LineCount() {
	case 0:
		v.tokens.Add(token.Space)
		v.tokens.Add(token.StartBlock)
		v.tokens.Add(token.Space)
	case 1:
		v.tokens.Add(token.Space)
		v.tokens.Add(token.StartBlock)
		v.tokens.Add(token.NewLine1)
	case 2:
		v.tokens.Add(token.NewLine1)
		v.tokens.Add(token.StartBlock)
		v.tokens.Add(token.NewLine1)
	default:
		v.tokens.Add(token.NewLine1)
		v.tokens.Add(token.StartBlock)
		v.tokens.Add(token.NewNewLineToken(fragment.LineCount() - 1))
	}
}

func (v *TokenizeJavaFragmentVisitor) VisitStartMovableJavaBlockFragment(fragment intmod.IStartMovableJavaBlockFragment) {
}

func (v *TokenizeJavaFragmentVisitor) VisitStartSingleStatementBlockFragment(fragment intmod.IStartSingleStatementBlockFragment) {
	switch fragment.LineCount() {
	case 0:
		v.tokens.Add(token.StartDeclarationOrStatementBlock)
		v.tokens.Add(token.Space)
	case 1:
		switch fragment.End().LineCount() {
		case 0:
			v.tokens.Add(token.StartDeclarationOrStatementBlock)
			v.tokens.Add(token.NewLine1)
		default:
			v.tokens.Add(token.Space)
			v.tokens.Add(token.StartBlock)
			v.tokens.Add(token.NewLine1)
		}
	case 2:
		v.tokens.Add(token.NewLine1)
		v.tokens.Add(token.StartBlock)
		v.tokens.Add(token.NewLine1)
	default:
		v.tokens.Add(token.NewLine1)
		v.tokens.Add(token.StartBlock)
		v.tokens.Add(token.NewNewLineToken(fragment.LineCount() - 1))
	}
}

func (v *TokenizeJavaFragmentVisitor) VisitStartStatementsBlockFragment(fragment intmod.IStartStatementsBlockFragment) {
	minimalLineCount := fragment.Group().MinimalLineCount()

	switch fragment.LineCount() {
	case 0:
		v.tokens.Add(token.Space)
		v.tokens.Add(token.StartBlock)
		v.tokens.Add(token.Space)
	case 1:
		if minimalLineCount == 0 {
			v.tokens.Add(token.NewLine1)
			v.tokens.Add(token.StartBlock)
			v.tokens.Add(token.Space)
		} else {
			v.tokens.Add(token.Space)
			v.tokens.Add(token.StartBlock)
			v.tokens.Add(token.NewLine1)
		}
	case 2:
		switch minimalLineCount {
		case 0:
			v.tokens.Add(token.NewLine2)
			v.tokens.Add(token.StartBlock)
			v.tokens.Add(token.Space)
		case 1:
			v.tokens.Add(token.Space)
			v.tokens.Add(token.StartBlock)
			v.tokens.Add(token.NewLine2)
		default:
			v.tokens.Add(token.NewLine1)
			v.tokens.Add(token.StartBlock)
			v.tokens.Add(token.NewLine1)
		}
	default:
		switch minimalLineCount {
		case 0:
			v.tokens.Add(token.NewNewLineToken(fragment.LineCount()))
			v.tokens.Add(token.StartBlock)
			v.tokens.Add(token.Space)
		case 1:
			v.tokens.Add(token.Space)
			v.tokens.Add(token.StartBlock)
			v.tokens.Add(token.NewNewLineToken(fragment.LineCount()))
		default:
			v.tokens.Add(token.NewLine1)
			v.tokens.Add(token.StartBlock)
			v.tokens.Add(token.NewNewLineToken(fragment.LineCount() - 1))
		}
	}
}

func (v *TokenizeJavaFragmentVisitor) VisitStartStatementsDoWhileBlockFragment(fragment intmod.IStartStatementsDoWhileBlockFragment) {
	v.visit(fragment, DoTokens.ToSlice())
}

func (v *TokenizeJavaFragmentVisitor) VisitStartStatementsInfiniteForBlockFragment(fragment intmod.IStartStatementsInfiniteForBlockFragment) {
	v.visit(fragment, EmptyForTokens.ToSlice())
}

func (v *TokenizeJavaFragmentVisitor) VisitStartStatementsInfiniteWhileBlockFragment(fragment intmod.IStartStatementsInfiniteWhileBlockFragment) {
	v.visit(fragment, EmptyWhileTokens.ToSlice())
}

func (v *TokenizeJavaFragmentVisitor) VisitStartStatementsTryBlockFragment(fragment intmod.IStartStatementsTryBlockFragment) {
	v.visit(fragment, TryTokens.ToSlice())
}

func (v *TokenizeJavaFragmentVisitor) VisitTokensFragment(fragment intmod.ITokensFragment) {
	for _, tkn := range fragment.Tokens() {
		tkn.Accept(v.unknownLineNumberTokenVisitor)
	}
}

func (v *TokenizeJavaFragmentVisitor) visit(fragment intmod.IStartStatementsBlockFragment, adds []intmod.IToken) {
	minimalLineCount := fragment.Group().MinimalLineCount()

	switch fragment.LineCount() {
	case 0:
		v.tokens.AddAll(adds)
		v.tokens.Add(token.Space)
		v.tokens.Add(token.StartBlock)
		v.tokens.Add(token.Space)
	case 1:
		if minimalLineCount == 0 {
			v.tokens.Add(token.NewLine1)
			v.tokens.AddAll(adds)
			v.tokens.Add(token.Space)
			v.tokens.Add(token.StartBlock)
			v.tokens.Add(token.Space)
		} else {
			v.tokens.AddAll(adds)
			v.tokens.Add(token.Space)
			v.tokens.Add(token.StartBlock)
			v.tokens.Add(token.NewLine1)
		}
	case 2:
		switch minimalLineCount {
		case 0:
			v.tokens.Add(token.NewLine2)
			v.tokens.AddAll(adds)
			v.tokens.Add(token.Space)
			v.tokens.Add(token.StartBlock)
			v.tokens.Add(token.Space)
		case 1:
			v.tokens.Add(token.NewLine1)
			v.tokens.AddAll(adds)
			v.tokens.Add(token.Space)
			v.tokens.Add(token.StartBlock)
			v.tokens.Add(token.NewLine1)
		default:
			v.tokens.AddAll(adds)
			v.tokens.Add(token.NewLine1)
			v.tokens.Add(token.StartBlock)
			v.tokens.Add(token.NewLine1)
		}
	default:
		switch minimalLineCount {
		case 0:
			v.tokens.Add(token.NewNewLineToken(fragment.LineCount()))
			v.tokens.AddAll(adds)
			v.tokens.Add(token.Space)
			v.tokens.Add(token.StartBlock)
			v.tokens.Add(token.Space)
		case 1:
			v.tokens.AddAll(adds)
			v.tokens.Add(token.Space)
			v.tokens.Add(token.StartBlock)
			v.tokens.Add(token.NewNewLineToken(fragment.LineCount()))
		default:
			v.tokens.AddAll(adds)
			v.tokens.Add(token.NewLine1)
			v.tokens.Add(token.StartBlock)
			v.tokens.Add(token.NewNewLineToken(fragment.LineCount() - 1))
		}
	}
}

func NewKnownLineNumberTokenVisitor(parent *TokenizeJavaFragmentVisitor) *KnownLineNumberTokenVisitor {
	return &KnownLineNumberTokenVisitor{
		parent: parent,
	}
}

type KnownLineNumberTokenVisitor struct {
	token.AbstractNopTokenVisitor

	parent            *TokenizeJavaFragmentVisitor
	currentLineNumber int
}

func (v *KnownLineNumberTokenVisitor) Reset(firstLineNumber int) {
	v.currentLineNumber = firstLineNumber
}

func (v *KnownLineNumberTokenVisitor) VisitEndBlockToken(tkn intmod.IEndBlockToken) {
	//assert tkn != EndBlockToken.END_BLOCK : "Unexpected EndBlockToken.END_BLOCK at this step. Uses 'JavaFragmentFactory.addEnd***(fragments)' instead";
	v.parent.tokens.Add(tkn)
}

func (v *KnownLineNumberTokenVisitor) VisitLineNumberToken(tkn intmod.ILineNumberToken) {
	lineNumber := tkn.LineNumber()

	if lineNumber != api.UnknownLineNumber {
		if v.currentLineNumber != api.UnknownLineNumber {
			switch lineNumber - v.currentLineNumber {
			case 0:
				break
			case 1:
				v.parent.tokens.Add(token.NewLine1)
				break
			case 2:
				v.parent.tokens.Add(token.NewLine2)
				break
			default:
				v.parent.tokens.Add(token.NewNewLineToken(lineNumber - v.currentLineNumber))
				break
			}
		}

		v.currentLineNumber = tkn.LineNumber()
		v.parent.tokens.Add(tkn)
	}
}

func (v *KnownLineNumberTokenVisitor) VisitStartBlockToken(tkn intmod.IStartBlockToken) {
	//assert tkn != StartBlockToken.START_BLOCK : "Unexpected StartBlockToken.START_BLOCK at this step. Uses 'JavaFragmentFactory.addStart***(fragments)' instead";
	v.parent.tokens.Add(tkn)
}

func (v *KnownLineNumberTokenVisitor) VisitBooleanConstantToken(tkn intmod.IBooleanConstantToken) {
	v.parent.tokens.Add(tkn)
}

func (v *KnownLineNumberTokenVisitor) VisitCharacterConstantToken(tkn intmod.ICharacterConstantToken) {
	v.parent.tokens.Add(tkn)
}

func (v *KnownLineNumberTokenVisitor) VisitDeclarationToken(tkn intmod.IDeclarationToken) {
	v.parent.tokens.Add(tkn)
}

func (v *KnownLineNumberTokenVisitor) VisitEndMarkerToken(tkn intmod.IEndMarkerToken) {
	v.parent.tokens.Add(tkn)
}

func (v *KnownLineNumberTokenVisitor) VisitKeywordToken(tkn intmod.IKeywordToken) {
	v.parent.tokens.Add(tkn)
}

func (v *KnownLineNumberTokenVisitor) VisitNumericConstantToken(tkn intmod.INumericConstantToken) {
	v.parent.tokens.Add(tkn)
}

func (v *KnownLineNumberTokenVisitor) VisitReferenceToken(tkn intmod.IReferenceToken) {
	v.parent.tokens.Add(tkn)
}

func (v *KnownLineNumberTokenVisitor) VisitStartMarkerToken(tkn intmod.IStartMarkerToken) {
	v.parent.tokens.Add(tkn)
}

func (v *KnownLineNumberTokenVisitor) VisitStringConstantToken(tkn intmod.IStringConstantToken) {
	v.parent.tokens.Add(tkn)
}

func (v *KnownLineNumberTokenVisitor) VisitTextToken(tkn intmod.ITextToken) {
	v.parent.tokens.Add(tkn)
}

func NewUnknownLineNumberTokenVisitor(parent *TokenizeJavaFragmentVisitor) *UnknownLineNumberTokenVisitor {
	return &UnknownLineNumberTokenVisitor{
		parent: parent,
	}
}

type UnknownLineNumberTokenVisitor struct {
	parent *TokenizeJavaFragmentVisitor
}

func (v *UnknownLineNumberTokenVisitor) VisitEndBlockToken(tkn intmod.IEndBlockToken) {
	// assert token != EndBlockToken.END_BLOCK : "Unexpected EndBlockToken.END_BLOCK at this step. Uses 'JavaFragmentFactory.addEnd***(fragments)' instead";
	v.parent.tokens.Add(tkn)
}

func (v *UnknownLineNumberTokenVisitor) VisitLineNumberToken(tkn intmod.ILineNumberToken) {
	// assert token.getLineNumber() == Printer.UNKNOWN_LINE_NUMBER : "LineNumberToken cannot have a known line number. Uses 'LineNumberTokensFragment' instead";
}

func (v *UnknownLineNumberTokenVisitor) VisitStartBlockToken(tkn intmod.IStartBlockToken) {
	// assert token != StartBlockToken.START_BLOCK : "Unexpected StartBlockToken.START_BLOCK at this step. Uses 'JavaFragmentFactory.addStart***(fragments)' instead";
	v.parent.tokens.Add(tkn)
}

func (v *UnknownLineNumberTokenVisitor) VisitBooleanConstantToken(tkn intmod.IBooleanConstantToken) {
	v.parent.tokens.Add(tkn)
}

func (v *UnknownLineNumberTokenVisitor) VisitCharacterConstantToken(tkn intmod.ICharacterConstantToken) {
	v.parent.tokens.Add(tkn)
}

func (v *UnknownLineNumberTokenVisitor) VisitDeclarationToken(tkn intmod.IDeclarationToken) {
	v.parent.tokens.Add(tkn)
}

func (v *UnknownLineNumberTokenVisitor) VisitEndMarkerToken(tkn intmod.IEndMarkerToken) {
	v.parent.tokens.Add(tkn)
}

func (v *UnknownLineNumberTokenVisitor) VisitKeywordToken(tkn intmod.IKeywordToken) {
	v.parent.tokens.Add(tkn)
}

func (v *UnknownLineNumberTokenVisitor) VisitNewLineToken(tkn intmod.INewLineToken) {
	v.parent.tokens.Add(tkn)
}

func (v *UnknownLineNumberTokenVisitor) VisitNumericConstantToken(tkn intmod.INumericConstantToken) {
	v.parent.tokens.Add(tkn)
}

func (v *UnknownLineNumberTokenVisitor) VisitReferenceToken(tkn intmod.IReferenceToken) {
	v.parent.tokens.Add(tkn)
}

func (v *UnknownLineNumberTokenVisitor) VisitStartMarkerToken(tkn intmod.IStartMarkerToken) {
	v.parent.tokens.Add(tkn)
}

func (v *UnknownLineNumberTokenVisitor) VisitStringConstantToken(tkn intmod.IStringConstantToken) {
	v.parent.tokens.Add(tkn)
}

func (v *UnknownLineNumberTokenVisitor) VisitTextToken(tkn intmod.ITextToken) {
	v.parent.tokens.Add(tkn)
}
