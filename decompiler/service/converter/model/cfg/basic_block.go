package cfg

import (
	"fmt"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

var TypeNames = []string{"DELETED", "START", "END", "STATEMENTS", "THROW", "RETURN", "RETURN_VALUE", "SWITCH_DECLARATION", "SWITCH",
	"SWITCH_BREAK", "TRY_DECLARATION", "TRY", "TRY_JSR", "TYPE_TRY_ECLIPSE", "JSR", "RET", "CONDITIONAL_BRANCH",
	"IF", "IF_ELSE", "CONDITION", "CONDITION_OR", "CONDITION_AND", "CONDITION_TERNARY_OPERATOR", "LOOP",
	"LOOP_START", "LOOP_CONTINUE", "LOOP_END", "GOTO", "INFINITE_GOTO", "GOTO_IN_TERNARY_OP", "TERNARY_OP", "JUMP"}

var EmptyExceptionHandlers = util.NewDefaultList[intsrv.IExceptionHandler]()
var EmptySwitchCases = util.NewDefaultList[intsrv.ISwitchCase]()

var SwitchBreak = NewImmutableBasicBlock(intsrv.TypeSwitchBreak).(intsrv.IBasicBlock)
var LoopStart = NewImmutableBasicBlock(intsrv.TypeLoopStart).(intsrv.IBasicBlock)
var LoopContinue = NewImmutableBasicBlock(intsrv.TypeLoopContinue).(intsrv.IBasicBlock)
var LoopEnd = NewImmutableBasicBlock(intsrv.TypeLoopEnd).(intsrv.IBasicBlock)
var End = newImmutableBasicBlockEnd(intsrv.TypeEnd).(intsrv.IBasicBlock)
var Return = NewImmutableBasicBlock(intsrv.TypeReturn).(intsrv.IBasicBlock)

func NewBasicBlock(controlFlowGraph intsrv.IControlFlowGraph, index int, original intsrv.IBasicBlock) intsrv.IBasicBlock {
	return NewBasicBlockWithBasicBlocks(controlFlowGraph, index, original, util.NewSet[intsrv.IBasicBlock]())
}

func NewBasicBlockWithBasicBlocks(controlFlowGraph intsrv.IControlFlowGraph, index int,
	original intsrv.IBasicBlock, predecessors util.ISet[intsrv.IBasicBlock]) intsrv.IBasicBlock {
	b := &BasicBlock{
		controlFlowGraph:  controlFlowGraph,
		index:             index,
		typ:               original.Type(),
		fromOffset:        original.FromOffset(),
		toOffset:          original.ToOffset(),
		next:              original.Next(),
		branch:            original.Branch(),
		condition:         original.Condition(),
		inverseCondition:  original.IsInverseCondition(),
		sub1:              original.Sub1(),
		sub2:              original.Sub2(),
		exceptionHandlers: original.ExceptionHandlers(),
		switchCases:       original.SwitchCases(),
		predecessors:      predecessors,
	}
	if b.exceptionHandlers == nil {
		b.exceptionHandlers = EmptyExceptionHandlers
	}
	if b.switchCases == nil {
		b.switchCases = EmptySwitchCases
	}
	return b
}

func NewBasicBlockWithRaw(controlFlowGraph *ControlFlowGraph, index, typ, fromOffset, toOffset int,
	inverseCondition bool) *BasicBlock {
	return NewBasicBlockWithRawBasicBlock(controlFlowGraph, index, typ, fromOffset,
		toOffset, inverseCondition, util.NewSet[intsrv.IBasicBlock]())
}

func NewBasicBlockWithRawBasicBlock(controlFlowGraph *ControlFlowGraph, index, typ, fromOffset, toOffset int,
	inverseCondition bool, predecessors util.ISet[intsrv.IBasicBlock]) *BasicBlock {
	return &BasicBlock{
		controlFlowGraph:  controlFlowGraph,
		index:             index,
		typ:               typ,
		fromOffset:        fromOffset,
		toOffset:          toOffset,
		next:              End,
		branch:            End,
		condition:         End,
		sub1:              End,
		sub2:              End,
		exceptionHandlers: EmptyExceptionHandlers,
		switchCases:       EmptySwitchCases,
		predecessors:      predecessors,
		inverseCondition:  inverseCondition,
	}
}

type BasicBlock struct {
	controlFlowGraph  intsrv.IControlFlowGraph
	index             int
	typ               int
	fromOffset        int
	toOffset          int
	next              intsrv.IBasicBlock
	branch            intsrv.IBasicBlock
	condition         intsrv.IBasicBlock
	inverseCondition  bool
	sub1              intsrv.IBasicBlock
	sub2              intsrv.IBasicBlock
	exceptionHandlers util.IList[intsrv.IExceptionHandler]
	switchCases       util.IList[intsrv.ISwitchCase]
	predecessors      util.ISet[intsrv.IBasicBlock]
}

func (b *BasicBlock) ControlFlowGraph() intsrv.IControlFlowGraph {
	return b.controlFlowGraph
}

func (b *BasicBlock) Index() int {
	return b.index
}

func (b *BasicBlock) Type() int {
	return b.typ
}

func (b *BasicBlock) SetType(flags int) {
	b.typ = flags
}

func (b *BasicBlock) FromOffset() int {
	return b.fromOffset
}

func (b *BasicBlock) SetFromOffset(offset int) {
	b.fromOffset = offset
}

func (b *BasicBlock) ToOffset() int {
	return b.toOffset
}

func (b *BasicBlock) SetToOffset(offset int) {
	b.toOffset = offset
}

func (b *BasicBlock) Next() intsrv.IBasicBlock {
	return b.next
}

func (b *BasicBlock) SetNext(next intsrv.IBasicBlock) {
	b.next = next
}

func (b *BasicBlock) Branch() intsrv.IBasicBlock {
	return b.branch
}

func (b *BasicBlock) SetBranch(branch intsrv.IBasicBlock) {
	b.branch = branch
}

func (b *BasicBlock) Condition() intsrv.IBasicBlock {
	return b.condition
}

func (b *BasicBlock) SetCondition(condition intsrv.IBasicBlock) {
	b.condition = condition
}

func (b *BasicBlock) IsInverseCondition() bool {
	return b.inverseCondition
}

func (b *BasicBlock) SetInverseCondition(inverseCondition bool) {
	b.inverseCondition = inverseCondition
}

func (b *BasicBlock) Sub1() intsrv.IBasicBlock {
	return b.sub1
}

func (b *BasicBlock) SetSub1(sub intsrv.IBasicBlock) {
	b.sub1 = sub
}

func (b *BasicBlock) Sub2() intsrv.IBasicBlock {
	return b.sub2
}

func (b *BasicBlock) SetSub2(sub intsrv.IBasicBlock) {
	b.sub2 = sub
}

func (b *BasicBlock) ExceptionHandlers() util.IList[intsrv.IExceptionHandler] {
	return b.exceptionHandlers
}

func (b *BasicBlock) SwitchCases() util.IList[intsrv.ISwitchCase] {
	return b.switchCases
}

func (b *BasicBlock) Predecessors() util.ISet[intsrv.IBasicBlock] {
	return b.predecessors
}

func (b *BasicBlock) SetSwitchCases(switchCases util.IList[intsrv.ISwitchCase]) {
	b.switchCases = switchCases
}

func (b *BasicBlock) SetPredecessors(predecessors util.ISet[intsrv.IBasicBlock]) {
	b.predecessors = predecessors
}

func (b *BasicBlock) FirstLineNumber() int {
	return b.controlFlowGraph.LineNumber(b.fromOffset)
}

func (b *BasicBlock) LastLineNumber() int {
	return b.controlFlowGraph.LineNumber(b.toOffset - 1)
}

func (b *BasicBlock) Contains(basicBlock intsrv.IBasicBlock) bool {
	if b.next == basicBlock {
		return true
	}

	if b.branch == basicBlock {
		return true
	}

	for _, exceptionHandler := range b.exceptionHandlers.ToSlice() {
		if exceptionHandler.BasicBlock() == basicBlock {
			return true
		}
	}

	for _, switchCase := range b.switchCases.ToSlice() {
		if switchCase.BasicBlock() == basicBlock {
			return true
		}
	}

	if b.sub1 == basicBlock {
		return true
	}

	if b.sub2 == basicBlock {
		return true
	}

	return false
}

func (b *BasicBlock) Replace(old, nevv intsrv.IBasicBlock) {
	fmt.Println("Old: ", old.String())
	fmt.Println("New: ", nevv.String())
	fmt.Println("Next: ", b.next.String())
	fmt.Println("Branch", b.branch.String())
	//var err error
	//var bn, o, bb, s1, s2 uint64 = 0, 0, 0, 0, 0
	//
	//if b.next != nil {
	//	bn, err = strconv.ParseUint(fmt.Sprintf("%p", b.next), 0, 64)
	//	if err != nil {
	//		bn = 0
	//	}
	//}
	//if old != nil {
	//	o, err = strconv.ParseUint(fmt.Sprintf("%p", old), 0, 64)
	//	if err != nil {
	//		o = 0
	//	}
	//}
	//if b.branch != nil {
	//	bb, err = strconv.ParseUint(fmt.Sprintf("%p", b.branch), 0, 64)
	//	if err != nil {
	//		bb = 0
	//	}
	//}
	//if b.sub1 != nil {
	//	s1, err = strconv.ParseUint(fmt.Sprintf("%p", b.sub1), 0, 64)
	//	if err != nil {
	//		s1 = 0
	//	}
	//}
	//if b.sub2 != nil {
	//	s2, err = strconv.ParseUint(fmt.Sprintf("%p", b.sub2), 0, 64)
	//	if err != nil {
	//		s2 = 0
	//	}
	//}

	if b.next == old {
		b.next = nevv
	}

	if b.branch == old {
		b.branch = nevv
	}

	//if bn == o {
	//	b.next = nevv
	//}
	//
	//if bb == o {
	//	b.branch = nevv
	//}

	for _, exceptionHandler := range b.exceptionHandlers.ToSlice() {
		exceptionHandler.Replace(old, nevv)
	}

	for _, switchCase := range b.switchCases.ToSlice() {
		switchCase.Replace(old, nevv)
	}

	if b.sub1 == old {
		b.sub1 = nevv
	}
	if b.sub2 == old {
		b.sub2 = nevv
	}

	if b.predecessors.Contains(old) {
		_ = b.predecessors.Remove(old)
		if nevv != End {
			_ = b.predecessors.Add(nevv)
		}
	}
}

func (b *BasicBlock) ReplaceWithOlds(olds util.ISet[intsrv.IBasicBlock], nevv intsrv.IBasicBlock) {
	if olds.Contains(b.next) {
		b.next = nevv
	}

	if olds.Contains(b.branch) {
		b.branch = nevv
	}

	for _, exceptionHandler := range b.exceptionHandlers.ToSlice() {
		exceptionHandler.ReplaceWithOlds(olds, nevv)
	}

	for _, switchCase := range b.switchCases.ToSlice() {
		switchCase.ReplaceWithOlds(olds, nevv)
	}

	if olds.Contains(b.sub1) {
		b.sub1 = nevv
	}

	if olds.Contains(b.sub2) {
		b.sub2 = nevv
	}

	b.predecessors.RemoveAll(olds.ToSlice())
	b.predecessors.Add(nevv)
}

func (b *BasicBlock) AddExceptionHandler(internalThrowableName string, basicBlock intsrv.IBasicBlock) {
	if b.exceptionHandlers == EmptyExceptionHandlers {
		// Add a first handler
		b.exceptionHandlers = util.NewDefaultList[intsrv.IExceptionHandler]()
		b.exceptionHandlers.Add(NewExceptionHandler(internalThrowableName, basicBlock))
	} else {
		for _, exceptionHandler := range b.exceptionHandlers.ToSlice() {
			if exceptionHandler.BasicBlock() == basicBlock {
				// Found -> Add an other 'internalThrowableName'
				exceptionHandler.AddInternalThrowableName(internalThrowableName)
				return
			}
		}
		// Not found -> Add a new handler
		b.exceptionHandlers.Add(NewExceptionHandler(internalThrowableName, basicBlock))
	}
}

func (b *BasicBlock) InverseCondition() {
	switch b.typ {
	case intsrv.TypeCondition, intsrv.TypeConditionTernaryOperator, intsrv.TypeGotoInTernaryOperator:
		b.inverseCondition = true
	case intsrv.TypeConditionAnd:
		b.typ = intsrv.TypeConditionOr
		b.sub1.InverseCondition()
		b.sub2.InverseCondition()
	case intsrv.TypeConditionOr:
		b.typ = intsrv.TypeConditionAnd
		b.sub1.InverseCondition()
		b.sub2.InverseCondition()
	default:
	}
}

func (b *BasicBlock) MatchType(types int) bool {
	return (b.typ & types) != 0
}

func (b *BasicBlock) TypeName() string {
	if b.typ == 0 {
		return TypeNames[0]
	}
	return TypeNames[b.numberOfTrailingZeros(b.typ)+1]
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
	s := fmt.Sprintf("BasicBlock{index=%d, from=%d, to=%d, type=%s, inverseCondition=", b.index, b.fromOffset, b.toOffset, b.TypeName())

	if b.inverseCondition {
		s += "true"
	} else {
		s += "false"
	}

	if !b.predecessors.IsEmpty() {
		s += ", predecessors=["

		length := b.predecessors.Size()
		for i := 0; i < length; i++ {
			s += fmt.Sprintf("%d", b.predecessors.Get(i).Index())
			if i != length-1 {
				s += ", "
			}
		}

		s += "]"
	}

	return s + "}"
}

func (b *BasicBlock) Equals(other intsrv.IBasicBlock) bool {
	return b.Index() == other.Index()
}

func NewExceptionHandler(internalThrowableName string, basicBlock intsrv.IBasicBlock) intsrv.IExceptionHandler {
	return &ExceptionHandler{
		internalThrowableName:       internalThrowableName,
		otherInternalThrowableNames: util.NewDefaultList[string](),
		basicBlock:                  basicBlock,
	}
}

type ExceptionHandler struct {
	internalThrowableName       string
	otherInternalThrowableNames util.IList[string]
	basicBlock                  intsrv.IBasicBlock
}

func (h *ExceptionHandler) InternalThrowableName() string {
	return h.internalThrowableName
}

func (h *ExceptionHandler) OtherInternalThrowableNames() util.IList[string] {
	return h.otherInternalThrowableNames
}

func (h *ExceptionHandler) BasicBlock() intsrv.IBasicBlock {
	return h.basicBlock
}

func (h *ExceptionHandler) SetBasicBlock(basicBlock intsrv.IBasicBlock) {
	h.basicBlock = basicBlock
}

func (h *ExceptionHandler) AddInternalThrowableName(internalThrowableName string) {
	if h.otherInternalThrowableNames == nil {
		h.otherInternalThrowableNames = util.NewDefaultList[string]()
	}
	h.otherInternalThrowableNames.Add(internalThrowableName)
}

func (h *ExceptionHandler) Replace(old, nevv intsrv.IBasicBlock) {
	if h.basicBlock == old {
		h.basicBlock = nevv
	}
}

func (h *ExceptionHandler) ReplaceWithOlds(olds util.ISet[intsrv.IBasicBlock], nevv intsrv.IBasicBlock) {
	for _, old := range olds.ToSlice() {
		if h.basicBlock == old {
			h.basicBlock = nevv
		}
	}
}

func (h *ExceptionHandler) String() string {
	if h.otherInternalThrowableNames == nil {
		return fmt.Sprintf("BasicBlock.Handler{%s -> %s}", h.internalThrowableName, h.basicBlock)
	}
	return fmt.Sprintf("BasicBlock.Handler{%s, %s -> %s}", h.internalThrowableName, h.otherInternalThrowableNames, h.basicBlock)
}

func NewSwitchCase(basicBlock intsrv.IBasicBlock) intsrv.ISwitchCase {
	return NewSwitchCaseWithValue(-1, basicBlock)
}

func NewSwitchCaseWithValue(value int, basicBlock intsrv.IBasicBlock) intsrv.ISwitchCase {
	return &SwitchCase{
		value:       value,
		offset:      basicBlock.FromOffset(),
		basicBlock:  basicBlock,
		defaultCase: true,
	}
}

type SwitchCase struct {
	value       int
	offset      int
	basicBlock  intsrv.IBasicBlock
	defaultCase bool
}

func (c *SwitchCase) Value() int {
	return c.value
}

func (c *SwitchCase) Offset() int {
	return c.offset
}

func (c *SwitchCase) BasicBlock() intsrv.IBasicBlock {
	return c.basicBlock
}

func (c *SwitchCase) SetBasicBlock(basicBlock intsrv.IBasicBlock) {
	c.basicBlock = basicBlock
}

func (c *SwitchCase) IsDefaultCase() bool {
	return c.defaultCase
}

func (c *SwitchCase) Replace(old, nevv intsrv.IBasicBlock) {
	if c.basicBlock == old {
		c.basicBlock = nevv
	}
}

func (c *SwitchCase) ReplaceWithOlds(olds util.ISet[intsrv.IBasicBlock], nevv intsrv.IBasicBlock) {
	for _, old := range olds.ToSlice() {
		if c.basicBlock == old {
			c.basicBlock = nevv
		}
	}
}

func (c *SwitchCase) String() string {
	if c.defaultCase {
		return fmt.Sprintf("BasicBlock.SwitchCase{default: %s}", c.basicBlock)
	}
	return fmt.Sprintf("BasicBlock.SwitchCase{'%d': %s}", c.value, c.basicBlock)
}

func NewImmutableBasicBlock(typ int) intsrv.IImmutableBasicBlock {
	return &ImmutableBasicBlock{
		BasicBlock: *NewBasicBlockWithRawBasicBlock(nil, -1, typ,
			0, 0, true, util.NewSet[intsrv.IBasicBlock]()),
	}
}

func newImmutableBasicBlockEnd(typ int) intsrv.IImmutableBasicBlock {
	end := &ImmutableBasicBlock{
		BasicBlock: BasicBlock{
			controlFlowGraph:  nil,
			index:             -1,
			typ:               typ,
			fromOffset:        0,
			toOffset:          0,
			next:              nil,
			branch:            nil,
			condition:         nil,
			sub1:              nil,
			sub2:              nil,
			exceptionHandlers: EmptyExceptionHandlers,
			switchCases:       EmptySwitchCases,
			predecessors:      util.NewSet[intsrv.IBasicBlock](),
			inverseCondition:  true,
		},
	}

	end.next = end
	end.branch = end
	end.condition = end
	end.sub1 = end
	end.sub2 = end

	return end
}

type ImmutableBasicBlock struct {
	BasicBlock
}

func (b *ImmutableBasicBlock) FirstLineNumber() int {
	return 0
}

func (b *ImmutableBasicBlock) LastLineNumber() int {
	return 0
}
