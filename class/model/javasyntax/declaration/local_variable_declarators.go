package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

func NewLocalVariableDeclarators() intmod.ILocalVariableDeclarators {
	return NewLocalVariableDeclaratorsWithCapacity(0)
}

func NewLocalVariableDeclaratorsWithCapacity(capacity int) intmod.ILocalVariableDeclarators {
	return &LocalVariableDeclarators{
		DefaultList: *util.NewDefaultListWithCapacity[intmod.ILocalVariableDeclarator](capacity),
	}
}

type LocalVariableDeclarators struct {
	util.DefaultList[intmod.ILocalVariableDeclarator]
}

func (d *LocalVariableDeclarators) Name() string { return "" }

func (d *LocalVariableDeclarators) SetName(_ string) {}

func (d *LocalVariableDeclarators) Dimension() int { return 0 }

func (d *LocalVariableDeclarators) SetDimension(dimension int) {}

func (d *LocalVariableDeclarators) LineNumber() int {
	if d.Size() == 0 {
		return 0
	}

	return d.Get(0).LineNumber()
}

func (d *LocalVariableDeclarators) VariableInitializer() intmod.IVariableInitializer { return nil }

func (d *LocalVariableDeclarators) Accept(visitor intmod.IDeclarationVisitor) {
	visitor.VisitLocalVariableDeclarators(d)
}

func (d *LocalVariableDeclarators) String() string {
	return "LocalVariableDeclarators{}"
}
