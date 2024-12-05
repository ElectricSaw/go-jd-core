package model

type IJavaFragment interface {
	Accept(visitor IJavaFragmentVisitor)
}

type IJavaFragmentVisitor interface {
	VisitEndBodyFragment(fragment IEndBodyFragment)
	VisitEndBlockInParameterFragment(fragment IEndBlockInParameterFragment)
	VisitEndBlockFragment(fragment IEndBlockFragment)
	VisitEndBodyInParameterFragment(fragment IEndBodyInParameterFragment)
	VisitEndMovableJavaBlockFragment(fragment IEndMovableJavaBlockFragment)
	VisitEndSingleStatementBlockFragment(fragment IEndSingleStatementBlockFragment)
	VisitEndStatementsBlockFragment(fragment IEndStatementsBlockFragment)
	VisitImportsFragment(fragment IImportsFragment)
	VisitLineNumberTokensFragment(fragment ILineNumberTokensFragment)
	VisitSpacerBetweenMembersFragment(fragment ISpacerBetweenMembersFragment)
	VisitSpacerFragment(fragment ISpacerFragment)
	VisitSpaceSpacerFragment(fragment ISpaceSpacerFragment)
	VisitStartBlockFragment(fragment IStartBlockFragment)
	VisitStartBodyFragment(fragment IStartBodyFragment)
	VisitStartMovableJavaBlockFragment(fragment IStartMovableJavaBlockFragment)
	VisitStartSingleStatementBlockFragment(fragment IStartSingleStatementBlockFragment)
	VisitStartStatementsBlockFragment(fragment IStartStatementsBlockFragment)
	VisitStartStatementsDoWhileBlockFragment(fragment IStartStatementsDoWhileBlockFragment)
	VisitStartStatementsInfiniteForBlockFragment(fragment IStartStatementsInfiniteForBlockFragment)
	VisitStartStatementsInfiniteWhileBlockFragment(fragment IStartStatementsInfiniteWhileBlockFragment)
	VisitStartStatementsTryBlockFragment(fragment IStartStatementsTryBlockFragment)
	VisitTokensFragment(fragment ITokensFragment)
}

type IEndBlockFragment interface {
	IEndFlexibleBlockFragment

	Start() IStartBlockFragment
	SetStart(start IStartBlockFragment)

	IncLineCount(force bool) bool
	DecLineCount(force bool) bool

	Accept(visitor IJavaFragmentVisitor)
}

type IEndBlockInParameterFragment interface {
	IEndBlockFragment

	Accept(visitor IJavaFragmentVisitor)
}

type IEndBodyFragment interface {
	IEndFlexibleBlockFragment

	Start() IStartBlockFragment
	SetStart(start IStartBlockFragment)

	IncLineCount(force bool) bool
	DecLineCount(force bool) bool

	Accept(visitor IJavaFragmentVisitor)
}

type IEndBodyInParameterFragment interface {
	IEndBodyFragment

	IncLineCount(force bool) bool
	DecLineCount(force bool) bool

	Accept(visitor IJavaFragmentVisitor)
}

type IEndMovableJavaBlockFragment interface {
	IEndMovableBlockFragment

	Accept(visitor IJavaFragmentVisitor)
}

type IEndSingleStatementBlockFragment interface {
	IEndFlexibleBlockFragment

	Start() IStartSingleStatementBlockFragment
	SetStart(start IStartSingleStatementBlockFragment)

	IncLineCount(force bool) bool
	DecLineCount(force bool) bool

	Accept(visitor IJavaFragmentVisitor)
}

type IEndStatementsBlockFragment interface {
	IEndFlexibleBlockFragment

	Group() IStartStatementsBlockFragmentGroup
	SetGroup(start IStartStatementsBlockFragmentGroup)

	Accept(visitor IJavaFragmentVisitor)
}

type IImportsFragment interface {
	IFlexibleFragment

	AddImport(internalName, qualifiedName string)
	IncCounter(internalName string) bool
	IsEmpty() bool
	InitLineCounts()
	Contains(internalName string) bool
	Import(internalName string) IImport
	Imports() []IImport

	Accept(visitor IJavaFragmentVisitor)
}

type IImport interface {
	InternalName() string
	SetInternalName(internalName string)
	QualifiedName() string
	SetQualifiedName(qualifiedName string)
	Counter() int
	SetCounter(counter int)
	IncCounter()
}

type ILineNumberTokensFragment interface {
	IFixedFragment

	TokenAt(index int) IToken
	Tokens() []IToken
	Accept(visitor IJavaFragmentVisitor)
}

type ISearchLineNumberVisitor interface {
	ITokenVisitor

	LineNumber() int
	SetLineNumber(lineNumber int)
	NewLineNumber() int
	SetNewLineNumber(newLineCounter int)

	Reset()
	VisitLineNumberToken(token ILineNumberToken)
	VisitNewLineToken(token INewLineToken)
}

type ISpaceSpacerFragment interface {
	ISpacerFragment

	Accept(visitor IJavaFragmentVisitor)
}

type ISpacerBetweenMembersFragment interface {
	ISpacerBetweenMovableBlocksFragment

	Accept(visitor IJavaFragmentVisitor)
}

type ISpacerFragment interface {
	IFlexibleFragment

	Accept(visitor IJavaFragmentVisitor)
}

type IStartBlockFragment interface {
	IStartFlexibleBlockFragment

	End() IEndBlockFragment
	SetEnd(end IEndBlockFragment)

	IncLineCount(force bool) bool
	DecLineCount(force bool) bool

	Accept(visitor IJavaFragmentVisitor)
}

type IStartBodyFragment interface {
	IStartFlexibleBlockFragment

	End() IEndBodyFragment
	SetEnd(end IEndBodyFragment)

	IncLineCount(force bool) bool
	DecLineCount(force bool) bool

	Accept(visitor IJavaFragmentVisitor)
}

type IStartMovableJavaBlockFragment interface {
	IStartMovableBlockFragment

	Accept(visitor IJavaFragmentVisitor)
}

type IStartSingleStatementBlockFragment interface {
	IStartFlexibleBlockFragment

	End() IEndSingleStatementBlockFragment
	SetEnd(end IEndSingleStatementBlockFragment)

	IncLineCount(force bool) bool
	DecLineCount(force bool) bool

	Accept(visitor IJavaFragmentVisitor)
}

type IStartStatementsBlockFragment interface {
	IStartFlexibleBlockFragment

	Group() IStartStatementsBlockFragmentGroup
	SetGroup(start IStartStatementsBlockFragmentGroup)

	Accept(visitor IJavaFragmentVisitor)
}

type IStartStatementsBlockFragmentGroup interface {
	MinimalLineCount() int
	SetMinimalLineCount(minimalLineCount int)
	Get(index int) IFlexibleFragment
	Add(frag IFlexibleFragment)
	Remove(index int) IFlexibleFragment
	ToSlice() []IFlexibleFragment
}

type IStartStatementsDoWhileBlockFragment interface {
	IStartStatementsBlockFragment

	Accept(visitor IJavaFragmentVisitor)
}

type IStartStatementsInfiniteForBlockFragment interface {
	IStartStatementsBlockFragment

	Accept(visitor IJavaFragmentVisitor)
}

type IStartStatementsInfiniteWhileBlockFragment interface {
	IStartStatementsBlockFragment

	Accept(visitor IJavaFragmentVisitor)
}

type IStartStatementsTryBlockFragment interface {
	IStartStatementsBlockFragment

	Accept(visitor IJavaFragmentVisitor)
}

type ITokensFragment interface {
	IFlexibleFragment

	TokenAt(index int) IToken
	Tokens() []IToken

	Accept(visitor IJavaFragmentVisitor)
}

type ILineCountVisitor interface {
	ITokenVisitor

	LineCount() int
	SetLineCount(lineCount int)

	VisitLineNumberToken(token ILineNumberToken)
}
