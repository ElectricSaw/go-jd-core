package cfg

import "bitbucket.org/coontec/go-jd-core/class/model/classfile"

func NewControlFlowGraph(method *classfile.Method) *ControlFlowGraph {
	return &ControlFlowGraph{
		Method:              method,
		List:                make([]IBasicBlock, 0),
		OffsetToLineNumbers: make([]int, 0),
	}
}

type ControlFlowGraph struct {
	Method              *classfile.Method
	List                []IBasicBlock
	OffsetToLineNumbers []int
}

func (g *ControlFlowGraph) NewBasicBlock1(original BasicBlock) *BasicBlock {
	basicBlock := NewBasicBlock(g, len(g.List), original)
	g.List = append(g.List, basicBlock)
	return basicBlock
}

func (g *ControlFlowGraph) NewBasicBlock2(fromOffset, toOffset int) *BasicBlock {
	return g.NewBasicBlock3(0, fromOffset, toOffset)
}

func (g *ControlFlowGraph) NewBasicBlock3(typ, fromOffset, toOffset int) *BasicBlock {
	basicBlock := NewBasicBlockWithRaw(g, len(g.List), typ, fromOffset, toOffset, true)
	g.List = append(g.List, basicBlock)
	return basicBlock
}

func (g *ControlFlowGraph) NewBasicBlock4(typ, fromOffset, toOffset int, inverseCondition bool) *BasicBlock {
	basicBlock := NewBasicBlockWithRaw(g, len(g.List), typ, fromOffset, toOffset, inverseCondition)
	g.List = append(g.List, basicBlock)
	return basicBlock
}

func (g *ControlFlowGraph) NewBasicBlock5(typ, fromOffset, toOffset int, predecessors []IBasicBlock) *BasicBlock {
	basicBlock := NewBasicBlockWithRawBasicBlock(g, len(g.List), typ, fromOffset, toOffset, true, predecessors)
	g.List = append(g.List, basicBlock)
	return basicBlock
}

func (g *ControlFlowGraph) SetOffsetToLineNumbers(offsetToLineNumbers []int) {
	g.OffsetToLineNumbers = offsetToLineNumbers
}

func (g *ControlFlowGraph) LineNumber(offset int) int {
	if g.OffsetToLineNumbers == nil {
		return 0
	}

	return g.OffsetToLineNumbers[offset]
}
