package service

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/classpath"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

const (
	TypeDeleted                  = 0
	TypeStart                    = 1 << 0
	TypeEnd                      = 1 << 1
	TypeStatements               = 1 << 2
	TypeThrow                    = 1 << 3
	TypeReturn                   = 1 << 4
	TypeReturnValue              = 1 << 5
	TypeSwitchDeclaration        = 1 << 6
	TypeSwitch                   = 1 << 7
	TypeSwitchBreak              = 1 << 8
	TypeTryDeclaration           = 1 << 9
	TypeTry                      = 1 << 10
	TypeTryJsr                   = 1 << 11
	TypeTryEclipse               = 1 << 12
	TypeJsr                      = 1 << 13
	TypeRet                      = 1 << 14
	TypeConditionalBranch        = 1 << 15
	TypeIf                       = 1 << 16
	TypeIfElse                   = 1 << 17
	TypeCondition                = 1 << 18
	TypeConditionOr              = 1 << 19
	TypeConditionAnd             = 1 << 20
	TypeConditionTernaryOperator = 1 << 21
	TypeLoop                     = 1 << 22
	TypeLoopStart                = 1 << 23
	TypeLoopContinue             = 1 << 24
	TypeLoopEnd                  = 1 << 25
	TypeGoto                     = 1 << 26
	TypeInfiniteGoto             = 1 << 27
	TypeGotoInTernaryOperator    = 1 << 28
	TypeTernaryOperator          = 1 << 29
	TypeJump                     = 1 << 30

	GroupSingleSuccessor = TypeStart | TypeStatements | TypeTryDeclaration | TypeJsr | TypeLoop | TypeIf | TypeIfElse | TypeSwitch | TypeTry | TypeTryJsr | TypeTryEclipse | TypeGoto | TypeGotoInTernaryOperator | TypeTernaryOperator
	GroupSynthetic       = TypeStart | TypeEnd | TypeConditionalBranch | TypeSwitchDeclaration | TypeTryDeclaration | TypeRet | TypeGoto | TypeJump
	GroupCode            = TypeStatements | TypeThrow | TypeReturn | TypeReturnValue | TypeSwitchDeclaration | TypeConditionalBranch | TypeJsr | TypeRet | TypeSwitch | TypeGoto | TypeInfiniteGoto | TypeGotoInTernaryOperator | TypeCondition | TypeConditionTernaryOperator
	GroupEnd             = TypeEnd | TypeThrow | TypeReturn | TypeReturnValue | TypeRet | TypeSwitchBreak | TypeLoopStart | TypeLoopContinue | TypeLoopEnd | TypeInfiniteGoto | TypeJump
	GroupCondition       = TypeCondition | TypeConditionOr | TypeConditionAnd | TypeConditionTernaryOperator
)

type ILoop interface {
	Start() IBasicBlock
	SetStart(start IBasicBlock)
	Members() util.ISet[IBasicBlock]
	SetMembers(members util.ISet[IBasicBlock])
	End() IBasicBlock
	SetEnd(end IBasicBlock)
}

type IControlFlowGraph interface {
	BasicBlocks() util.IList[IBasicBlock]
	Method() intmod.IMethod
	SetOffsetToLineNumbers(offsetToLineNumbers []int)
	LineNumber(offset int) int
	NewBasicBlock1(original IBasicBlock) IBasicBlock
	NewBasicBlock2(fromOffset, toOffset int) IBasicBlock
	NewBasicBlock3(typ, fromOffset, toOffset int) IBasicBlock
	NewBasicBlock4(typ, fromOffset, toOffset int, inverseCondition bool) IBasicBlock
	NewBasicBlock5(typ, fromOffset, toOffset int, predecessors util.ISet[IBasicBlock]) IBasicBlock
}

type IBasicBlock interface {
	ControlFlowGraph() IControlFlowGraph
	Index() int
	Type() int
	SetType(flags int)
	FromOffset() int
	SetFromOffset(offset int)
	ToOffset() int
	SetToOffset(offset int)
	Next() IBasicBlock
	SetNext(next IBasicBlock)
	Branch() IBasicBlock
	SetBranch(branch IBasicBlock)
	Condition() IBasicBlock
	IsInverseCondition() bool
	Sub1() IBasicBlock
	SetSub1(sub IBasicBlock)
	Sub2() IBasicBlock
	SetSub2(sub IBasicBlock)
	ExceptionHandlers() util.IList[IExceptionHandler]
	SwitchCases() util.IList[ISwitchCase]
	Predecessors() util.ISet[IBasicBlock]
	FirstLineNumber() int
	LastLineNumber() int

	Contains(basicBlock IBasicBlock) bool
	Replace(old, nevv IBasicBlock)
	ReplaceWithOlds(olds util.ISet[IBasicBlock], nevv IBasicBlock)
	AddExceptionHandler(internalThrowableName string, basicBlock IBasicBlock)
	InverseCondition()
	MatchType(types int) bool
	TypeName() string
	String() string
	Equals(other IBasicBlock) bool
}

type IExceptionHandler interface {
	InternalThrowableName() string
	OtherInternalThrowableNames() util.IList[string]
	BasicBlock() IBasicBlock

	AddInternalThrowableName(internalThrowableName string)
	Replace(old, nevv IBasicBlock)
	ReplaceWithOlds(olds util.ISet[IBasicBlock], nevv IBasicBlock)
	String() string
}

type ISwitchCase interface {
	Value() int
	Offset() int
	BasicBlock() IBasicBlock
	SetBasicBlock(basicBlock IBasicBlock)
	IsDefaultCase() bool
	Replace(old, nevv IBasicBlock)
	ReplaceWithOlds(olds util.ISet[IBasicBlock], nevv IBasicBlock)
	String() string
}

type IImmutableBasicBlock interface {
	IBasicBlock

	FirstLineNumber() int
	LastLineNumber() int
}
