package cfg

import "fmt"

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

var TypeNames = []string{"DELETED", "START", "END", "STATEMENTS", "THROW", "RETURN", "RETURN_VALUE", "SWITCH_DECLARATION", "SWITCH",
	"SWITCH_BREAK", "TRY_DECLARATION", "TRY", "TRY_JSR", "TYPE_TRY_ECLIPSE", "JSR", "RET", "CONDITIONAL_BRANCH",
	"IF", "IF_ELSE", "CONDITION", "CONDITION_OR", "CONDITION_AND", "CONDITION_TERNARY_OPERATOR", "LOOP",
	"LOOP_START", "LOOP_CONTINUE", "LOOP_END", "GOTO", "INFINITE_GOTO", "GOTO_IN_TERNARY_OP", "TERNARY_OP", "JUMP"}

var EmptyExceptionHandlers = make([]ExceptionHandler, 0)
var EmptySwitchCases = make([]SwitchCase, 0)

var SwitchBreak = NewImmutableBasicBlock(TypeSwitchBreak)
var LoopStart = NewImmutableBasicBlock(TypeLoopStart)
var LoopContinue = NewImmutableBasicBlock(TypeLoopContinue)
var LoopEnd = NewImmutableBasicBlock(TypeLoopEnd)
var End = NewImmutableBasicBlock(TypeEnd)
var Return = NewImmutableBasicBlock(TypeReturn)

type IBasicBlock interface {
	GetIndex() int
	GetFromOffset() int

	Contains(basicBlock IBasicBlock) bool
	Replace(old, nevv IBasicBlock)
	ReplaceWithOlds(olds []IBasicBlock, nevv IBasicBlock)
	AddExceptionHandler(internalThrowableName string, basicBlock IBasicBlock)
	inverseCondition()
	MatchType(types int) bool
	TypeName() string
	String() string
	Equals(other IBasicBlock) bool
}

func NewBasicBlock(controlFlowGraph *ControlFlowGraph, index int, original BasicBlock) *BasicBlock {
	return NewBasicBlockWithBasicBlocks(controlFlowGraph, index, original, make([]IBasicBlock, 0))
}

func NewBasicBlockWithBasicBlocks(controlFlowGraph *ControlFlowGraph, index int, original BasicBlock, predecessors []IBasicBlock) *BasicBlock {
	return &BasicBlock{
		ControlFlowGraph:  controlFlowGraph,
		Index:             index,
		Type:              original.Type,
		FromOffset:        original.FromOffset,
		ToOffset:          original.ToOffset,
		Next:              original.Next,
		Branch:            original.Branch,
		Condition:         original.Condition,
		InverseCondition:  original.InverseCondition,
		Sub1:              original.Sub1,
		Sub2:              original.Sub2,
		ExceptionHandlers: original.ExceptionHandlers,
		SwitchCases:       original.SwitchCases,
		Predecessors:      predecessors,
	}
}

func NewBasicBlockWithRaw(controlFlowGraph *ControlFlowGraph, index, typ, fromOffset, toOffset int, inverseCondition bool) *BasicBlock {
	return NewBasicBlockWithRawBasicBlock(controlFlowGraph, index, typ, fromOffset, toOffset, inverseCondition, make([]IBasicBlock, 0))
}

func NewBasicBlockWithRawBasicBlock(controlFlowGraph *ControlFlowGraph, index, typ, fromOffset, toOffset int, inverseCondition bool, predecessors []IBasicBlock) *BasicBlock {
	return &BasicBlock{
		ControlFlowGraph: controlFlowGraph,
		Index:            index,
		Type:             typ,
		FromOffset:       fromOffset,
		ToOffset:         toOffset,
		Next:             End,
		Branch:           End,
		Condition:        End,
		Sub1:             End,
		Sub2:             End,
		Predecessors:     predecessors,
		InverseCondition: inverseCondition,
	}
}

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
	ExceptionHandlers []ExceptionHandler
	SwitchCases       []SwitchCase
	Predecessors      []IBasicBlock
}

func (b *BasicBlock) GetIndex() int {
	return b.Index
}

func (b *BasicBlock) GetFromOffset() int {
	return b.FromOffset
}

func (b *BasicBlock) Contains(basicBlock IBasicBlock) bool {
	if b.Next == basicBlock {
		return true
	}

	if b.Branch == basicBlock {
		return true
	}

	for _, exceptionHandler := range b.ExceptionHandlers {
		if exceptionHandler.BasicBlock == basicBlock {
			return true
		}
	}

	for _, switchCase := range b.SwitchCases {
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
	if b.Next == old {
		b.Next = nevv
	}

	if b.Branch == old {
		b.Branch = nevv
	}

	for _, exceptionHandler := range b.ExceptionHandlers {
		exceptionHandler.replace(old, nevv)
	}

	for _, switchCase := range b.SwitchCases {
		switchCase.replace(old, nevv)
	}

	if b.Sub1 == old {
		b.Sub1 = nevv
	}
	if b.Sub2 == old {
		b.Sub2 = nevv
	}

	if contains(b.Predecessors, old) {
		b.Predecessors = remove(b.Predecessors, old)
		if nevv != End {
			b.Predecessors = append(b.Predecessors, nevv)
		}
	}
}

func (b *BasicBlock) ReplaceWithOlds(olds []IBasicBlock, nevv IBasicBlock) {
	if contains(olds, b.Next) {
		b.Next = nevv
	}

	if contains(olds, b.Branch) {
		b.Branch = nevv
	}

	for _, exceptionHandler := range b.ExceptionHandlers {
		exceptionHandler.replaceWithOlds(olds, nevv)
	}

	for _, switchCase := range b.SwitchCases {
		switchCase.replaceWithOlds(olds, nevv)
	}

	if contains(olds, b.Sub1) {
		b.Sub1 = nevv
	}

	if contains(olds, b.Sub2) {
		b.Sub2 = nevv
	}

	b.Predecessors = removeAll(b.Predecessors, olds)
	b.Predecessors = append(b.Predecessors, nevv)
}

func (b *BasicBlock) AddExceptionHandler(internalThrowableName string, basicBlock IBasicBlock) {
	if &b.ExceptionHandlers == &EmptyExceptionHandlers {
		// Add a first handler
		b.ExceptionHandlers = make([]ExceptionHandler, 0)
		b.ExceptionHandlers = append(b.ExceptionHandlers, *NewExceptionHandler(internalThrowableName, basicBlock))
	} else {
		for _, exceptionHandler := range b.ExceptionHandlers {
			if exceptionHandler.BasicBlock == basicBlock {
				// Found -> Add an other 'internalThrowableName'
				exceptionHandler.AddInternalThrowableName(internalThrowableName)
				return
			}
		}
		// Not found -> Add a new handler
		b.ExceptionHandlers = append(b.ExceptionHandlers, *NewExceptionHandler(internalThrowableName, basicBlock))
	}
}

func (b *BasicBlock) inverseCondition() {
	switch b.Type {
	case TypeCondition, TypeConditionTernaryOperator, TypeGotoInTernaryOperator:
		b.InverseCondition = true
	case TypeConditionAnd:
		b.Type = TypeConditionOr
		b.Sub1.inverseCondition()
		b.Sub2.inverseCondition()
	case TypeConditionOr:
		b.Type = TypeConditionAnd
		b.Sub1.inverseCondition()
		b.Sub2.inverseCondition()
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

	if !(len(b.Predecessors) == 0) {
		s += ", predecessors=["

		length := len(b.Predecessors)
		for i := 0; i < length; i++ {
			s += fmt.Sprintf("%d", b.Predecessors[i].GetIndex())
			if i != length-1 {
				s += ", "
			}
		}

		s += "]"
	}

	return s + "}"
}

func (b *BasicBlock) Equals(other IBasicBlock) bool {
	return b.Index == other.GetIndex()
}

func NewExceptionHandler(internalThrowableName string, basicBlock IBasicBlock) *ExceptionHandler {
	return &ExceptionHandler{
		InternalThrowableName: internalThrowableName,
		BasicBlock:            basicBlock,
	}
}

type ExceptionHandler struct {
	InternalThrowableName       string
	OtherInternalThrowableNames []string
	BasicBlock                  IBasicBlock
}

func (h *ExceptionHandler) AddInternalThrowableName(internalThrowableName string) {
	if h.OtherInternalThrowableNames == nil {
		h.OtherInternalThrowableNames = make([]string, 0)
	}
	h.OtherInternalThrowableNames = append(h.OtherInternalThrowableNames, internalThrowableName)
}

func (h *ExceptionHandler) replace(old, nevv IBasicBlock) {
	if h.BasicBlock == old {
		h.BasicBlock = nevv
	}
}

func (h *ExceptionHandler) replaceWithOlds(olds []IBasicBlock, nevv IBasicBlock) {
	for _, old := range olds {
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

func NewSwitchCase(basicBlock IBasicBlock) *SwitchCase {
	return &SwitchCase{
		Offset:      basicBlock.GetFromOffset(),
		BasicBlock:  basicBlock,
		DefaultCase: true,
	}
}

func NewSwitchCaseWithValue(value int, basicBlock IBasicBlock) *SwitchCase {
	return &SwitchCase{
		Value:       value,
		Offset:      basicBlock.GetFromOffset(),
		BasicBlock:  basicBlock,
		DefaultCase: true,
	}
}

type SwitchCase struct {
	Value       int
	Offset      int
	BasicBlock  IBasicBlock
	DefaultCase bool
}

func (c *SwitchCase) isDefaultCase() bool {
	return c.DefaultCase
}

func (c *SwitchCase) replace(old, nevv IBasicBlock) {
	if c.BasicBlock == old {
		c.BasicBlock = nevv
	}
}

func (c *SwitchCase) replaceWithOlds(olds []IBasicBlock, nevv IBasicBlock) {
	for _, old := range olds {
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

func NewImmutableBasicBlock(typ int) *ImmutableBasicBlock {
	return &ImmutableBasicBlock{
		BasicBlock: *NewBasicBlockWithRawBasicBlock(nil, -1, typ, 0, 0, true, make([]IBasicBlock, 0)),
	}
}

type ImmutableBasicBlock struct {
	BasicBlock
}

func (b *ImmutableBasicBlock) getFirstLineNumber() int {
	return 0
}

func (b *ImmutableBasicBlock) getLastLineNumber() int {
	return 0
}

func contains(list []IBasicBlock, value IBasicBlock) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

func remove(list []IBasicBlock, value IBasicBlock) []IBasicBlock {
	i := 0
	length := len(list)
	for i = 0; i < length; i++ {
		if list[i] == value {
			return append(list[:i], list[i+1:]...)
		}
	}
	return list
}

func removeAll(list []IBasicBlock, values []IBasicBlock) []IBasicBlock {
	ret := list

	for _, value := range values {
		ret = remove(ret, value)
	}

	return ret
}
