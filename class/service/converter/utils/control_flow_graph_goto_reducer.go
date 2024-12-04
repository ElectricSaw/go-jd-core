package utils

import intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"

func ReduceGraphGotoReducer(cfg intsrv.IControlFlowGraph) {
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
