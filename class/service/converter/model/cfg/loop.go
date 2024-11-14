package cfg

import "fmt"

func NewLoop(Start IBasicBlock, Members []IBasicBlock, End IBasicBlock) *Loop {
	return &Loop{
		Start:   Start,
		Members: Members,
		End:     End,
	}
}

type Loop struct {
	Start   IBasicBlock
	Members []IBasicBlock
	End     IBasicBlock
}

func (l *Loop) String() string {
	str := fmt.Sprintf("Loop{start=%d, members=[", l.Start.GetIndex())

	length := len(l.Members)
	if l.Members != nil && length > 0 {
		for i := 0; i < length; i++ {
			str += fmt.Sprintf("%d", l.Members[i].GetIndex())
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
