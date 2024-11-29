package cfg

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

func NewControlFlowGraph(method intmod.IMethod) intsrv.IControlFlowGraph {
	return &ControlFlowGraph{
		method:              method,
		list:                util.NewDefaultList[intsrv.IBasicBlock](),
		offsetToLineNumbers: make([]int, 0),
	}
}

type ControlFlowGraph struct {
	method              intmod.IMethod
	list                util.IList[intsrv.IBasicBlock]
	offsetToLineNumbers []int
}

func (g *ControlFlowGraph) BasicBlocks() util.IList[intsrv.IBasicBlock] {
	return g.list
}

func (g *ControlFlowGraph) Method() intmod.IMethod {
	return g.method
}

func (g *ControlFlowGraph) SetOffsetToLineNumbers(offsetToLineNumbers []int) {
	g.offsetToLineNumbers = offsetToLineNumbers
}

func (g *ControlFlowGraph) LineNumber(offset int) int {
	if g.offsetToLineNumbers == nil {
		return 0
	}

	return g.offsetToLineNumbers[offset]
}

func (g *ControlFlowGraph) NewBasicBlock1(original intsrv.IBasicBlock) intsrv.IBasicBlock {
	basicBlock := NewBasicBlock(g, g.list.Size(), original)
	g.list.Add(basicBlock)
	return basicBlock
}

func (g *ControlFlowGraph) NewBasicBlock2(fromOffset, toOffset int) intsrv.IBasicBlock {
	return g.NewBasicBlock3(0, fromOffset, toOffset)
}

func (g *ControlFlowGraph) NewBasicBlock3(typ, fromOffset, toOffset int) intsrv.IBasicBlock {
	basicBlock := NewBasicBlockWithRaw(g, g.list.Size(), typ, fromOffset, toOffset, true)
	g.list.Add(basicBlock)
	return basicBlock
}

func (g *ControlFlowGraph) NewBasicBlock4(typ, fromOffset, toOffset int, inverseCondition bool) intsrv.IBasicBlock {
	basicBlock := NewBasicBlockWithRaw(g, g.list.Size(), typ, fromOffset, toOffset, inverseCondition)
	g.list.Add(basicBlock)
	return basicBlock
}

func (g *ControlFlowGraph) NewBasicBlock5(typ, fromOffset, toOffset int, predecessors util.ISet[intsrv.IBasicBlock]) intsrv.IBasicBlock {
	basicBlock := NewBasicBlockWithRawBasicBlock(g, g.list.Size(), typ, fromOffset, toOffset, true, predecessors)
	g.list.Add(basicBlock)
	return basicBlock
}
