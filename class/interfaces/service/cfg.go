package service

import "bitbucket.org/coontec/go-jd-core/class/util"

type ILoop interface {
	Start() IBasicBlock
	Members() util.ISet[IBasicBlock]
	End() IBasicBlock
}

type IControlFlowGraph interface {
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
	FromOffset() int
	ToOffset() int
	Next() IBasicBlock
	Branch() IBasicBlock
	Condition() IBasicBlock
	IsInverseCondition() bool
	Sub1() IBasicBlock
	Sub2() IBasicBlock
	ExceptionHandlers() util.IList[IExceptionHandler]
	SwitchCases() util.IList[ISwitchCase]
	Predecessors() util.ISet[IBasicBlock]

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
