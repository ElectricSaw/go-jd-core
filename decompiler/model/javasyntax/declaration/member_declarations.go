package declaration

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewMemberDeclarations() MemberDeclarations {
	return NewMemberDeclarationsWithCapacity(0)
}

func NewMemberDeclarationsWithCapacity(capacity int) MemberDeclarations {
	return &MemberDeclarations{
		DefaultList: *util.NewDefaultListWithCapacity[MemberDeclaration](capacity).(*util.DefaultList[MemberDeclaration]),
	}
}

type MemberDeclarations struct {
	AbstractMemberDeclaration
	util.DefaultList[MemberDeclaration]
}

func (d *MemberDeclarations) IsList() bool {
	return d.DefaultList.IsList()
}

func (d *MemberDeclarations) Size() int {
	return d.DefaultList.Size()
}

func (d *MemberDeclarations) ToSlice() []MemberDeclaration {
	return d.DefaultList.ToSlice()
}

func (d *MemberDeclarations) ToList() *util.DefaultList[MemberDeclaration] {
	return d.DefaultList.ToList()
}

func (d *MemberDeclarations) First() MemberDeclaration {
	return d.DefaultList.First()
}

func (d *MemberDeclarations) Last() MemberDeclaration {
	return d.DefaultList.Last()
}

func (d *MemberDeclarations) Iterator() util.IIterator[MemberDeclaration] {
	return d.DefaultList.Iterator()
}

func (d *MemberDeclarations) AcceptDeclaration(visitor DeclarationVisitor) {
	visitor.VisitMemberDeclarations(d)
}
