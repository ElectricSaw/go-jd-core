package declaration

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewTypeDeclarations() TypeDeclarations {
	return NewTypeDeclarationsWithCapacity(0)
}

func NewTypeDeclarationsWithCapacity(capacity int) TypeDeclarations {
	return &TypeDeclarations{
		DefaultList: *util.NewDefaultListWithCapacity[MemberDeclaration](capacity).(*util.DefaultList[MemberDeclaration]),
	}
}

type TypeDeclarations struct {
	AbstractTypeDeclaration
	util.DefaultList[MemberDeclaration]
}

func (d *TypeDeclarations) IsList() bool {
	return d.DefaultList.IsList()
}

func (d *TypeDeclarations) Size() int {
	return d.DefaultList.Size()
}

func (d *TypeDeclarations) ToSlice() []MemberDeclaration {
	return d.DefaultList.ToSlice()
}

func (d *TypeDeclarations) ToList() *util.DefaultList[MemberDeclaration] {
	return d.DefaultList.ToList()
}

func (d *TypeDeclarations) First() MemberDeclaration {
	return d.DefaultList.First()
}

func (d *TypeDeclarations) Last() MemberDeclaration {
	return d.DefaultList.Last()
}

func (d *TypeDeclarations) Iterator() util.IIterator[MemberDeclaration] {
	return d.DefaultList.Iterator()
}

func (d *TypeDeclarations) AcceptDeclaration(visitor DeclarationVisitor) {
	visitor.VisitTypeDeclarations(d)
}
