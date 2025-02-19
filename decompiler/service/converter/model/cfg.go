package model

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/classfile"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

/////////////////////////////////////////////////////////////////////////
//  Global Variable
/////////////////////////////////////////////////////////////////////////

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

var (
	TypeNames = []string{"DELETED", "START", "END", "STATEMENTS", "THROW", "RETURN", "RETURN_VALUE", "SWITCH_DECLARATION", "SWITCH",
		"SWITCH_BREAK", "TRY_DECLARATION", "TRY", "TRY_JSR", "TYPE_TRY_ECLIPSE", "JSR", "RET", "CONDITIONAL_BRANCH",
		"IF", "IF_ELSE", "CONDITION", "CONDITION_OR", "CONDITION_AND", "CONDITION_TERNARY_OPERATOR", "LOOP",
		"LOOP_START", "LOOP_CONTINUE", "LOOP_END", "GOTO", "INFINITE_GOTO", "GOTO_IN_TERNARY_OP", "TERNARY_OP", "JUMP"}

	EmptyExceptionHandlers = util.NewDefaultList[*ExceptionHandler]()
	EmptySwitchCases       = util.NewDefaultList[*SwitchCase]()

	SwitchBreak  = NewImmutableBasicBlock(TypeSwitchBreak)
	LoopStart    = NewImmutableBasicBlock(TypeLoopStart)
	LoopContinue = NewImmutableBasicBlock(TypeLoopContinue)
	LoopEnd      = NewImmutableBasicBlock(TypeLoopEnd)
	End          = newImmutableBasicBlockEnd(TypeEnd)
	Return       = NewImmutableBasicBlock(TypeReturn)
)

/////////////////////////////////////////////////////////////////////////
//  New Functions
/////////////////////////////////////////////////////////////////////////

func NewBasicBlock(controlFlowGraph *ControlFlowGraph, index int, original IBasicBlock) BasicBlock {
	return NewBasicBlockWithBasicBlocks(controlFlowGraph, index, original, util.NewSet[IBasicBlock]())
}

func NewBasicBlockWithBasicBlocks(controlFlowGraph *ControlFlowGraph, index int,
	original IBasicBlock, predecessors util.ISet[IBasicBlock]) BasicBlock {
	b := BasicBlock{
		ControlFlowGraph:  controlFlowGraph,
		Index:             index,
		Type:              original.GetType(),
		FromOffset:        original.GetFromOffset(),
		ToOffset:          original.GetToOffset(),
		Next:              original.GetNext(),
		Branch:            original.GetBranch(),
		Condition:         original.GetCondition(),
		InverseCondition:  original.GetInverseCondition(),
		Sub1:              original.GetSub1(),
		Sub2:              original.GetSub2(),
		ExceptionHandlers: original.GetExceptionHandlers(),
		SwitchCases:       original.GetSwitchCases(),
		Predecessors:      predecessors,
	}
	if b.ExceptionHandlers == nil {
		b.ExceptionHandlers = EmptyExceptionHandlers
	}
	if b.SwitchCases == nil {
		b.SwitchCases = EmptySwitchCases
	}
	return b
}

func NewBasicBlockWithRaw(controlFlowGraph *ControlFlowGraph, index, typ, fromOffset, toOffset int,
	inverseCondition bool) BasicBlock {
	return NewBasicBlockWithRawBasicBlock(controlFlowGraph, index, typ, fromOffset,
		toOffset, inverseCondition, util.NewSet[IBasicBlock]())
}

func NewBasicBlockWithRawBasicBlock(controlFlowGraph *ControlFlowGraph, index, typ, fromOffset, toOffset int,
	inverseCondition bool, predecessors util.ISet[IBasicBlock]) BasicBlock {
	return BasicBlock{
		ControlFlowGraph:  controlFlowGraph,
		Index:             index,
		Type:              typ,
		FromOffset:        fromOffset,
		ToOffset:          toOffset,
		Next:              &End,
		Branch:            &End,
		Condition:         &End,
		Sub1:              &End,
		Sub2:              &End,
		ExceptionHandlers: EmptyExceptionHandlers,
		SwitchCases:       EmptySwitchCases,
		Predecessors:      predecessors,
		InverseCondition:  inverseCondition,
	}
}

func NewExceptionHandler(internalThrowableName string, basicBlock IBasicBlock) ExceptionHandler {
	return ExceptionHandler{
		InternalThrowableName:       internalThrowableName,
		OtherInternalThrowableNames: util.NewDefaultList[string](),
		BasicBlock:                  basicBlock,
	}
}

func NewSwitchCase(basicBlock IBasicBlock) SwitchCase {
	return NewSwitchCaseWithValue(-1, basicBlock)
}

func NewSwitchCaseWithValue(value int, basicBlock IBasicBlock) SwitchCase {
	return SwitchCase{
		Value:       value,
		Offset:      basicBlock.GetFromOffset(),
		BasicBlock:  basicBlock,
		DefaultCase: true,
	}
}

func NewImmutableBasicBlock(typ int) ImmutableBasicBlock {
	return ImmutableBasicBlock{
		ControlFlowGraph:  nil,
		Index:             -1,
		Type:              typ,
		FromOffset:        0,
		ToOffset:          0,
		Next:              &End,
		Branch:            &End,
		Condition:         &End,
		Sub1:              &End,
		Sub2:              &End,
		ExceptionHandlers: EmptyExceptionHandlers,
		SwitchCases:       EmptySwitchCases,
		Predecessors:      util.NewSet[IBasicBlock](),
		InverseCondition:  true,
	}
}

func newImmutableBasicBlockEnd(typ int) ImmutableBasicBlock {
	end := ImmutableBasicBlock{
		ControlFlowGraph:  nil,
		Index:             -1,
		Type:              typ,
		FromOffset:        0,
		ToOffset:          0,
		Next:              nil,
		Branch:            nil,
		Condition:         nil,
		Sub1:              nil,
		Sub2:              nil,
		ExceptionHandlers: EmptyExceptionHandlers,
		SwitchCases:       EmptySwitchCases,
		Predecessors:      util.NewSet[IBasicBlock](),
		InverseCondition:  true,
	}

	end.Next = &end
	end.Branch = &end
	end.Condition = &end
	end.Sub1 = &end
	end.Sub2 = &end

	return end
}

func NewControlFlowGraph(method classfile.Method) ControlFlowGraph {
	return ControlFlowGraph{
		Method:              method,
		BasicBlocks:         util.NewDefaultList[IBasicBlock](),
		OffsetToLineNumbers: make([]int, 0),
	}
}

func NewLoop(start IBasicBlock,
	members util.ISet[IBasicBlock],
	end IBasicBlock) Loop {
	return Loop{
		Start:   start,
		Members: members,
		End:     end,
	}
}

/////////////////////////////////////////////////////////////////////////
//  Interfaces
/////////////////////////////////////////////////////////////////////////

type IBasicBlock interface {
	GetControlFlowGraph() *ControlFlowGraph
	GetIndex() int
	GetType() int
	GetFromOffset() int
	GetToOffset() int
	GetNext() IBasicBlock
	GetBranch() IBasicBlock
	GetCondition() IBasicBlock
	GetInverseCondition() bool
	GetSub1() IBasicBlock
	GetSub2() IBasicBlock
	GetExceptionHandlers() util.IList[*ExceptionHandler]
	GetSwitchCases() util.IList[*SwitchCase]
	GetPredecessors() util.ISet[IBasicBlock]
	FirstLineNumber() int
	LastLineNumber() int
	Contains(basicBlock IBasicBlock) bool
	Replace(old, nevv IBasicBlock)
	ReplaceWithOlds(olds util.ISet[IBasicBlock], nevv IBasicBlock)
	AddExceptionHandler(internalThrowableName string, basicBlock IBasicBlock)
	GenerateInverseCondition()
	MatchType(types int) bool
	TypeName() string
	numberOfTrailingZeros(n int) int
	String() string
	Equals(other IBasicBlock) bool
}

/////////////////////////////////////////////////////////////////////////
//  Structures
/////////////////////////////////////////////////////////////////////////

type BasicBlock struct {
	ControlFlowGraph  *ControlFlowGraph
	Index             int
	Type              int
	FromOffset        int
	ToOffset          int
	Next              IBasicBlock
	Branch            IBasicBlock
	Condition         IBasicBlock
	InverseCondition  bool
	Sub1              IBasicBlock
	Sub2              IBasicBlock
	ExceptionHandlers util.IList[*ExceptionHandler]
	SwitchCases       util.IList[*SwitchCase]
	Predecessors      util.ISet[IBasicBlock]
}

func (b *BasicBlock) GetControlFlowGraph() *ControlFlowGraph {
	return b.ControlFlowGraph
}

func (b *BasicBlock) GetIndex() int {
	return b.Index
}

func (b *BasicBlock) GetType() int {
	return b.Type
}

func (b *BasicBlock) GetFromOffset() int {
	return b.FromOffset
}

func (b *BasicBlock) GetToOffset() int {
	return b.ToOffset
}

func (b *BasicBlock) GetNext() IBasicBlock {
	return b.Next
}

func (b *BasicBlock) GetBranch() IBasicBlock {
	return b.Branch
}

func (b *BasicBlock) GetCondition() IBasicBlock {
	return b.Condition
}

func (b *BasicBlock) GetInverseCondition() bool {
	return b.InverseCondition
}

func (b *BasicBlock) GetSub1() IBasicBlock {
	return b.Sub1
}

func (b *BasicBlock) GetSub2() IBasicBlock {
	return b.Sub2
}

func (b *BasicBlock) GetExceptionHandlers() util.IList[*ExceptionHandler] {
	return b.ExceptionHandlers
}

func (b *BasicBlock) GetSwitchCases() util.IList[*SwitchCase] {
	return b.SwitchCases
}

func (b *BasicBlock) GetPredecessors() util.ISet[IBasicBlock] {
	return b.Predecessors
}

func (b *BasicBlock) FirstLineNumber() int {
	return b.ControlFlowGraph.LineNumber(b.FromOffset)
}

func (b *BasicBlock) LastLineNumber() int {
	return b.ControlFlowGraph.LineNumber(b.ToOffset - 1)
}

func (b *BasicBlock) Contains(basicBlock IBasicBlock) bool {
	if b.Next == basicBlock {
		return true
	}

	if b.Branch == basicBlock {
		return true
	}

	for _, exceptionHandler := range b.ExceptionHandlers.ToSlice() {
		if exceptionHandler.BasicBlock == basicBlock {
			return true
		}
	}

	for _, switchCase := range b.SwitchCases.ToSlice() {
		if switchCase.BasicBlock == basicBlock {
			return true
		}
	}

	if b.Sub1 == basicBlock {
		return true
	}

	if b.Sub2 == basicBlock {
		return true
	}

	return false
}

func (b *BasicBlock) Replace(old, nevv IBasicBlock) {
	fmt.Println("Old: ", old.String())
	fmt.Println("New: ", nevv.String())
	fmt.Println("Next: ", b.Next.String())
	fmt.Println("Branch", b.Branch.String())

	if b.Next == old {
		b.Next = nevv
	}

	if b.Branch == old {
		b.Branch = nevv
	}

	for _, exceptionHandler := range b.ExceptionHandlers.ToSlice() {
		exceptionHandler.Replace(old, nevv)
	}

	for _, switchCase := range b.SwitchCases.ToSlice() {
		switchCase.Replace(old, nevv)
	}

	if b.Sub1 == old {
		b.Sub1 = nevv
	}
	if b.Sub2 == old {
		b.Sub2 = nevv
	}

	if b.Predecessors.Contains(old) {
		_ = b.Predecessors.Remove(old)
		if nevv != &End {
			_ = b.Predecessors.Add(nevv)
		}
	}
}

func (b *BasicBlock) ReplaceWithOlds(olds util.ISet[IBasicBlock], nevv IBasicBlock) {
	if olds.Contains(b.Next) {
		b.Next = nevv
	}

	if olds.Contains(b.Branch) {
		b.Branch = nevv
	}

	for _, exceptionHandler := range b.ExceptionHandlers.ToSlice() {
		exceptionHandler.ReplaceWithOlds(olds, nevv)
	}

	for _, switchCase := range b.SwitchCases.ToSlice() {
		switchCase.ReplaceWithOlds(olds, nevv)
	}

	if olds.Contains(b.Sub1) {
		b.Sub1 = nevv
	}

	if olds.Contains(b.Sub2) {
		b.Sub2 = nevv
	}

	b.Predecessors.RemoveAll(olds.ToSlice())
	b.Predecessors.Add(nevv)
}

func (b *BasicBlock) AddExceptionHandler(internalThrowableName string, basicBlock IBasicBlock) {
	if b.ExceptionHandlers == EmptyExceptionHandlers {
		// Add a first handler
		b.ExceptionHandlers = util.NewDefaultList[*ExceptionHandler]()
		exceptionHandler := NewExceptionHandler(internalThrowableName, basicBlock)
		b.ExceptionHandlers.Add(&exceptionHandler)
	} else {
		for _, exceptionHandler := range b.ExceptionHandlers.ToSlice() {
			if exceptionHandler.BasicBlock == basicBlock {
				// Found -> Add an other 'internalThrowableName'
				exceptionHandler.AddInternalThrowableName(internalThrowableName)
				return
			}
		}
		// Not found -> Add a new handler
		exceptionHandler := NewExceptionHandler(internalThrowableName, basicBlock)
		b.ExceptionHandlers.Add(&exceptionHandler)
	}
}

func (b *BasicBlock) GenerateInverseCondition() {
	switch b.Type {
	case TypeCondition, TypeConditionTernaryOperator, TypeGotoInTernaryOperator:
		b.InverseCondition = true
	case TypeConditionAnd:
		b.Type = TypeConditionOr
		b.Sub1.GenerateInverseCondition()
		b.Sub2.GenerateInverseCondition()
	case TypeConditionOr:
		b.Type = TypeConditionAnd
		b.Sub1.GenerateInverseCondition()
		b.Sub2.GenerateInverseCondition()
	default:
	}
}

func (b *BasicBlock) MatchType(types int) bool {
	return (b.Type & types) != 0
}

func (b *BasicBlock) TypeName() string {
	if b.Type == 0 {
		return TypeNames[0]
	}
	return TypeNames[b.numberOfTrailingZeros(b.Type)+1]
}

func (b *BasicBlock) numberOfTrailingZeros(n int) int {
	if n == 0 {
		return 32 // 32-bit integer의 모든 비트가 0인 경우
	}
	count := 0
	for (n & 1) == 0 {
		n >>= 1
		count++
	}
	return count
}

func (b *BasicBlock) String() string {
	s := fmt.Sprintf("BasicBlock{index=%d, from=%d, to=%d, type=%s, inverseCondition=", b.Index, b.FromOffset, b.ToOffset, b.TypeName())

	if b.InverseCondition {
		s += "true"
	} else {
		s += "false"
	}

	if !b.Predecessors.IsEmpty() {
		s += ", predecessors=["

		length := b.Predecessors.Size()
		for i := 0; i < length; i++ {
			s += fmt.Sprintf("%d", b.Predecessors.Get(i).GetIndex())
			if i != length-1 {
				s += ", "
			}
		}

		s += "]"
	}

	return s + "}"
}

func (b *BasicBlock) Equals(other IBasicBlock) bool {
	return b.GetIndex() == other.GetIndex()
}

type ControlFlowGraph struct {
	Method              classfile.Method
	BasicBlocks         util.IList[IBasicBlock]
	OffsetToLineNumbers []int
}

func (g *ControlFlowGraph) LineNumber(offset int) int {
	if g.OffsetToLineNumbers == nil {
		return 0
	}

	return g.OffsetToLineNumbers[offset]
}

func (g *ControlFlowGraph) Start() IBasicBlock {
	return g.BasicBlocks.Get(0)
}

func (g *ControlFlowGraph) NewBasicBlock1(original IBasicBlock) BasicBlock {
	basicBlock := NewBasicBlock(g, g.BasicBlocks.Size(), original)
	g.BasicBlocks.Add(&basicBlock)
	return basicBlock
}

func (g *ControlFlowGraph) NewBasicBlock2(fromOffset, toOffset int) BasicBlock {
	return g.NewBasicBlock3(0, fromOffset, toOffset)
}

func (g *ControlFlowGraph) NewBasicBlock3(typ, fromOffset, toOffset int) BasicBlock {
	basicBlock := NewBasicBlockWithRaw(g, g.BasicBlocks.Size(), typ, fromOffset, toOffset, true)
	g.BasicBlocks.Add(&basicBlock)
	return basicBlock
}

func (g *ControlFlowGraph) NewBasicBlock4(typ, fromOffset, toOffset int, inverseCondition bool) BasicBlock {
	basicBlock := NewBasicBlockWithRaw(g, g.BasicBlocks.Size(), typ, fromOffset, toOffset, inverseCondition)
	g.BasicBlocks.Add(&basicBlock)
	return basicBlock
}

func (g *ControlFlowGraph) NewBasicBlock5(typ, fromOffset, toOffset int, predecessors util.ISet[IBasicBlock]) BasicBlock {
	basicBlock := NewBasicBlockWithRawBasicBlock(g, g.BasicBlocks.Size(), typ, fromOffset, toOffset, true, predecessors)
	g.BasicBlocks.Add(&basicBlock)
	return basicBlock
}

type Loop struct {
	Start   IBasicBlock
	Members util.ISet[IBasicBlock]
	End     IBasicBlock
}

func (l *Loop) String() string {
	str := fmt.Sprintf("Loop{start=%d, members=[", l.Start.GetIndex())

	if l.Members != nil && l.Members.Size() > 0 {
		length := l.Members.Size()
		for i := 0; i < length; i++ {
			str += fmt.Sprintf("%d", l.Members.Get(i).GetIndex())
			if i < length-1 {
				str += ", "
			}
		}
	}

	str += "], end="
	if l.End != nil {
		str += fmt.Sprintf("%d", l.End.GetIndex())
	}
	str += "}"

	return str
}

/////////////////////////////////////////////////////////////////////////
//  Additional Structures
/////////////////////////////////////////////////////////////////////////

type ExceptionHandler struct {
	InternalThrowableName       string
	OtherInternalThrowableNames util.IList[string]
	BasicBlock                  IBasicBlock
}

func (h *ExceptionHandler) AddInternalThrowableName(internalThrowableName string) {
	if h.OtherInternalThrowableNames == nil {
		h.OtherInternalThrowableNames = util.NewDefaultList[string]()
	}
	h.OtherInternalThrowableNames.Add(internalThrowableName)
}

func (h *ExceptionHandler) Replace(old, nevv IBasicBlock) {
	if h.BasicBlock == old {
		h.BasicBlock = nevv
	}
}

func (h *ExceptionHandler) ReplaceWithOlds(olds util.ISet[IBasicBlock], nevv IBasicBlock) {
	for _, old := range olds.ToSlice() {
		if h.BasicBlock == old {
			h.BasicBlock = nevv
		}
	}
}

func (h *ExceptionHandler) String() string {
	if h.OtherInternalThrowableNames == nil {
		return fmt.Sprintf("BasicBlock.Handler{%s -> %s}", h.InternalThrowableName, h.BasicBlock)
	}
	return fmt.Sprintf("BasicBlock.Handler{%s, %s -> %s}", h.InternalThrowableName, h.OtherInternalThrowableNames, h.BasicBlock)
}

type SwitchCase struct {
	Value       int
	Offset      int
	BasicBlock  IBasicBlock
	DefaultCase bool
}

func (c *SwitchCase) Replace(old, nevv IBasicBlock) {
	if c.BasicBlock == old {
		c.BasicBlock = nevv
	}
}

func (c *SwitchCase) ReplaceWithOlds(olds util.ISet[IBasicBlock], nevv IBasicBlock) {
	for _, old := range olds.ToSlice() {
		if c.BasicBlock == old {
			c.BasicBlock = nevv
		}
	}
}

func (c *SwitchCase) String() string {
	if c.DefaultCase {
		return fmt.Sprintf("BasicBlock.SwitchCase{default: %s}", c.BasicBlock)
	}
	return fmt.Sprintf("BasicBlock.SwitchCase{'%d': %s}", c.Value, c.BasicBlock)
}

type ImmutableBasicBlock struct {
	ControlFlowGraph  *ControlFlowGraph
	Index             int
	Type              int
	FromOffset        int
	ToOffset          int
	Next              IBasicBlock
	Branch            IBasicBlock
	Condition         IBasicBlock
	InverseCondition  bool
	Sub1              IBasicBlock
	Sub2              IBasicBlock
	ExceptionHandlers util.IList[*ExceptionHandler]
	SwitchCases       util.IList[*SwitchCase]
	Predecessors      util.ISet[IBasicBlock]
}

func (b *ImmutableBasicBlock) GetControlFlowGraph() *ControlFlowGraph {
	return b.ControlFlowGraph
}

func (b *ImmutableBasicBlock) GetIndex() int {
	return b.Index
}

func (b *ImmutableBasicBlock) GetType() int {
	return b.Type
}

func (b *ImmutableBasicBlock) GetFromOffset() int {
	return b.FromOffset
}

func (b *ImmutableBasicBlock) GetToOffset() int {
	return b.ToOffset
}

func (b *ImmutableBasicBlock) GetNext() IBasicBlock {
	return b.Next
}

func (b *ImmutableBasicBlock) GetBranch() IBasicBlock {
	return b.Branch
}

func (b *ImmutableBasicBlock) GetCondition() IBasicBlock {
	return b.Condition
}

func (b *ImmutableBasicBlock) GetInverseCondition() bool {
	return b.InverseCondition
}

func (b *ImmutableBasicBlock) GetSub1() IBasicBlock {
	return b.Sub1
}

func (b *ImmutableBasicBlock) GetSub2() IBasicBlock {
	return b.Sub2
}

func (b *ImmutableBasicBlock) GetExceptionHandlers() util.IList[*ExceptionHandler] {
	return b.ExceptionHandlers
}

func (b *ImmutableBasicBlock) GetSwitchCases() util.IList[*SwitchCase] {
	return b.SwitchCases
}

func (b *ImmutableBasicBlock) GetPredecessors() util.ISet[IBasicBlock] {
	return b.Predecessors
}

func (b *ImmutableBasicBlock) FirstLineNumber() int {
	return 0
}

func (b *ImmutableBasicBlock) LastLineNumber() int {
	return 0
}

func (b *ImmutableBasicBlock) Contains(basicBlock IBasicBlock) bool {
	if b.Next == basicBlock {
		return true
	}

	if b.Branch == basicBlock {
		return true
	}

	for _, exceptionHandler := range b.ExceptionHandlers.ToSlice() {
		if exceptionHandler.BasicBlock == basicBlock {
			return true
		}
	}

	for _, switchCase := range b.SwitchCases.ToSlice() {
		if switchCase.BasicBlock == basicBlock {
			return true
		}
	}

	if b.Sub1 == basicBlock {
		return true
	}

	if b.Sub2 == basicBlock {
		return true
	}

	return false
}

func (b *ImmutableBasicBlock) Replace(old, nevv IBasicBlock) {
	fmt.Println("Old: ", old.String())
	fmt.Println("New: ", nevv.String())
	fmt.Println("Next: ", b.Next.String())
	fmt.Println("Branch", b.Branch.String())

	if b.Next == old {
		b.Next = nevv
	}

	if b.Branch == old {
		b.Branch = nevv
	}

	for _, exceptionHandler := range b.ExceptionHandlers.ToSlice() {
		exceptionHandler.Replace(old, nevv)
	}

	for _, switchCase := range b.SwitchCases.ToSlice() {
		switchCase.Replace(old, nevv)
	}

	if b.Sub1 == old {
		b.Sub1 = nevv
	}
	if b.Sub2 == old {
		b.Sub2 = nevv
	}

	if b.Predecessors.Contains(old) {
		_ = b.Predecessors.Remove(old)
		if nevv != &End {
			_ = b.Predecessors.Add(nevv)
		}
	}
}

func (b *ImmutableBasicBlock) ReplaceWithOlds(olds util.ISet[IBasicBlock], nevv IBasicBlock) {
	if olds.Contains(b.Next) {
		b.Next = nevv
	}

	if olds.Contains(b.Branch) {
		b.Branch = nevv
	}

	for _, exceptionHandler := range b.ExceptionHandlers.ToSlice() {
		exceptionHandler.ReplaceWithOlds(olds, nevv)
	}

	for _, switchCase := range b.SwitchCases.ToSlice() {
		switchCase.ReplaceWithOlds(olds, nevv)
	}

	if olds.Contains(b.Sub1) {
		b.Sub1 = nevv
	}

	if olds.Contains(b.Sub2) {
		b.Sub2 = nevv
	}

	b.Predecessors.RemoveAll(olds.ToSlice())
	b.Predecessors.Add(nevv)
}

func (b *ImmutableBasicBlock) AddExceptionHandler(internalThrowableName string, basicBlock IBasicBlock) {
	if b.ExceptionHandlers == EmptyExceptionHandlers {
		// Add a first handler
		b.ExceptionHandlers = util.NewDefaultList[*ExceptionHandler]()
		exceptionHandler := NewExceptionHandler(internalThrowableName, basicBlock)
		b.ExceptionHandlers.Add(&exceptionHandler)
	} else {
		for _, exceptionHandler := range b.ExceptionHandlers.ToSlice() {
			if exceptionHandler.BasicBlock == basicBlock {
				// Found -> Add an other 'internalThrowableName'
				exceptionHandler.AddInternalThrowableName(internalThrowableName)
				return
			}
		}
		// Not found -> Add a new handler
		exceptionHandler := NewExceptionHandler(internalThrowableName, basicBlock)
		b.ExceptionHandlers.Add(&exceptionHandler)
	}
}

func (b *ImmutableBasicBlock) GenerateInverseCondition() {
	switch b.Type {
	case TypeCondition, TypeConditionTernaryOperator, TypeGotoInTernaryOperator:
		b.InverseCondition = true
	case TypeConditionAnd:
		b.Type = TypeConditionOr
		b.Sub1.GenerateInverseCondition()
		b.Sub2.GenerateInverseCondition()
	case TypeConditionOr:
		b.Type = TypeConditionAnd
		b.Sub1.GenerateInverseCondition()
		b.Sub2.GenerateInverseCondition()
	default:
	}
}

func (b *ImmutableBasicBlock) MatchType(types int) bool {
	return (b.Type & types) != 0
}

func (b *ImmutableBasicBlock) TypeName() string {
	if b.Type == 0 {
		return TypeNames[0]
	}
	return TypeNames[b.numberOfTrailingZeros(b.Type)+1]
}

func (b *ImmutableBasicBlock) numberOfTrailingZeros(n int) int {
	if n == 0 {
		return 32 // 32-bit integer의 모든 비트가 0인 경우
	}
	count := 0
	for (n & 1) == 0 {
		n >>= 1
		count++
	}
	return count
}

func (b *ImmutableBasicBlock) String() string {
	s := fmt.Sprintf("BasicBlock{index=%d, from=%d, to=%d, type=%s, inverseCondition=", b.Index, b.FromOffset, b.ToOffset, b.TypeName())

	if b.InverseCondition {
		s += "true"
	} else {
		s += "false"
	}

	if !b.Predecessors.IsEmpty() {
		s += ", predecessors=["

		length := b.Predecessors.Size()
		for i := 0; i < length; i++ {
			s += fmt.Sprintf("%d", b.Predecessors.Get(i).GetIndex())
			if i != length-1 {
				s += ", "
			}
		}

		s += "]"
	}

	return s + "}"
}

func (b *ImmutableBasicBlock) Equals(other IBasicBlock) bool {
	return b.GetIndex() == other.GetIndex()
}

/////////////////////////////////////////////////////////////////////////
//  Functions
/////////////////////////////////////////////////////////////////////////
