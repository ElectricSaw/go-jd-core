package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

func NewMemberDeclarations() intmod.IMemberDeclarations {
	return NewMemberDeclarationsWithCapacity(0)
}

func NewMemberDeclarationsWithCapacity(capacity int) intmod.IMemberDeclarations {
	return &MemberDeclarations{
		DefaultList: *util.NewDefaultListWithCapacity[intmod.IMemberDeclaration](capacity).(*util.DefaultList[intmod.IMemberDeclaration]),
	}
}

type MemberDeclarations struct {
	AbstractMemberDeclaration
	util.DefaultList[intmod.IMemberDeclaration]
}

func (d *MemberDeclarations) IsList() bool {
	return d.DefaultList.IsList()
}

func (d *MemberDeclarations) Size() int {
	return d.DefaultList.Size()
}

func (d *MemberDeclarations) ToSlice() []intmod.IMemberDeclaration {
	return d.DefaultList.ToSlice()
}

func (d *MemberDeclarations) ToList() *util.DefaultList[intmod.IMemberDeclaration] {
	return d.DefaultList.ToList()
}

func (d *MemberDeclarations) First() intmod.IMemberDeclaration {
	return d.DefaultList.First()
}

func (d *MemberDeclarations) Last() intmod.IMemberDeclaration {
	return d.DefaultList.Last()
}

func (d *MemberDeclarations) Iterator() util.IIterator[intmod.IMemberDeclaration] {
	return d.DefaultList.Iterator()
}

func (d *MemberDeclarations) Accept(visitor intmod.IDeclarationVisitor) {
	visitor.VisitMemberDeclarations(d)
}
