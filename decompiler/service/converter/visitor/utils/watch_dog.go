package utils

import (
	"fmt"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
	"log"
)

func NewWatchDog() intsrv.IWatchDog {
	return &WatchDog{
		links: util.NewSet[intsrv.ILink](),
	}
}

type WatchDog struct {
	links util.ISet[intsrv.ILink]
}

func (d *WatchDog) Clear() {
	d.links.Clear()
}

func (d *WatchDog) Check(parent, child intsrv.IBasicBlock) {
	if !child.MatchType(intsrv.GroupEnd) {
		link := NewLink(parent, child)

		if d.links.Contains(link) {
			log.Fatalln(fmt.Sprintf("CFG watchdog: parent=%s, child=%s", parent, child))
			return
		}

		d.links.Add(link)
	}
}

func NewLink(parent, child intsrv.IBasicBlock) intsrv.ILink {
	return &Link{
		parentIndex: parent.Index(),
		childIndex:  child.Index(),
	}
}

type Link struct {
	parentIndex int
	childIndex  int
}

func (l *Link) ParentIndex() int {
	return l.parentIndex
}

func (l *Link) SetParentIndex(parentIndex int) {
	l.parentIndex = parentIndex
}

func (l *Link) ChildIndex() int {
	return l.childIndex
}

func (l *Link) SetChildIndex(childIndex int) {
	l.childIndex = childIndex
}

func (l *Link) HashCode() int {
	return 4807589 + l.parentIndex + 31*l.childIndex
}

func (l *Link) Equals(o intsrv.ILink) bool {
	return l.parentIndex == o.ParentIndex() && l.childIndex == o.ChildIndex()
}
