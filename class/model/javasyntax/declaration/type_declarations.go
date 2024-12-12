package declaration

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/class/util"
)

func NewTypeDeclarations() intmod.ITypeDeclarations {
	return NewTypeDeclarationsWithCapacity(0)
}

func NewTypeDeclarationsWithCapacity(capacity int) intmod.ITypeDeclarations {
	return &TypeDeclarations{
		DefaultList: *util.NewDefaultListWithCapacity[intmod.IMemberDeclaration](capacity).(*util.DefaultList[intmod.IMemberDeclaration]),
	}
}

type TypeDeclarations struct {
	AbstractTypeDeclaration
	util.DefaultList[intmod.IMemberDeclaration]
}

func (d *TypeDeclarations) IsList() bool {
	return d.DefaultList.IsList()
}

func (d *TypeDeclarations) Size() int {
	return d.DefaultList.Size()
}

func (d *TypeDeclarations) ToSlice() []intmod.IMemberDeclaration {
	return d.DefaultList.ToSlice()
}

func (d *TypeDeclarations) ToList() *util.DefaultList[intmod.IMemberDeclaration] {
	return d.DefaultList.ToList()
}

func (d *TypeDeclarations) First() intmod.IMemberDeclaration {
	return d.DefaultList.First()
}

func (d *TypeDeclarations) Last() intmod.IMemberDeclaration {
	return d.DefaultList.Last()
}

func (d *TypeDeclarations) Iterator() util.IIterator[intmod.IMemberDeclaration] {
	return d.DefaultList.Iterator()
}

func (d *TypeDeclarations) AcceptDeclaration(visitor intmod.IDeclarationVisitor) {
	visitor.VisitTypeDeclarations(d)
}
