package utils

import intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"

func ReduceControlFlowGraphGotoReducer(cfg intsrv.IControlFlowGraph) {
	for _, basicBlock := range cfg.BasicBlocks().ToSlice() {
		if basicBlock.Type() == intsrv.TypeGoto {
			successor := basicBlock.Next()

			if basicBlock == successor {
				basicBlock.Predecessors().Remove(basicBlock)
				basicBlock.SetType(intsrv.TypeInfiniteGoto)
			} else {
				successorPredecessors := successor.Predecessors()
				successorPredecessors.Remove(basicBlock)

				for _, predecessor := range basicBlock.Predecessors().ToSlice() {
					predecessor.Replace(basicBlock, successor)
					successorPredecessors.Add(predecessor)
				}

				basicBlock.SetType(intsrv.TypeDeleted)
			}
		}
	}
}
