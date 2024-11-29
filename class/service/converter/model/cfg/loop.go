package cfg

import (
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/util"
	"fmt"
)

func NewLoop(start intsrv.IBasicBlock,
	members util.ISet[intsrv.IBasicBlock],
	end intsrv.IBasicBlock) intsrv.ILoop {
	return &Loop{
		start:   start,
		members: members,
		end:     end,
	}
}

type Loop struct {
	start   intsrv.IBasicBlock
	members util.ISet[intsrv.IBasicBlock]
	end     intsrv.IBasicBlock
}

func (l *Loop) Start() intsrv.IBasicBlock {
	return l.start
}

func (l *Loop) SetStart(start intsrv.IBasicBlock) {
	l.start = start
}

func (l *Loop) Members() util.ISet[intsrv.IBasicBlock] {
	return l.members
}

func (l *Loop) SetMembers(members util.ISet[intsrv.IBasicBlock]) {
	l.members = members
}

func (l *Loop) End() intsrv.IBasicBlock {
	return l.end
}

func (l *Loop) SetEnd(end intsrv.IBasicBlock) {
	l.end = end
}

func (l *Loop) String() string {
	str := fmt.Sprintf("Loop{start=%d, members=[", l.start.Index())

	if l.members != nil && l.members.Size() > 0 {
		length := l.members.Size()
		for i := 0; i < length; i++ {
			str += fmt.Sprintf("%d", l.members.Get(i).Index())
			if i < length-1 {
				str += ", "
			}
		}
	}

	str += "], end="
	if l.end != nil {
		str += fmt.Sprintf("%d", l.end.Index())
	}
	str += "}"

	return str
}
