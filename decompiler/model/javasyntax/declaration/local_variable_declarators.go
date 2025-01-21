package declaration

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewLocalVariableDeclarators() LocalVariableDeclarators {
	return NewLocalVariableDeclaratorsWithCapacity(0)
}

func NewLocalVariableDeclaratorsWithCapacity(capacity int) LocalVariableDeclarators {
	return &LocalVariableDeclarators{
		DefaultList: *util.NewDefaultListWithCapacity[LocalVariableDeclarator](capacity).(*util.DefaultList[LocalVariableDeclarator]),
	}
}

type LocalVariableDeclarators struct {
	util.DefaultList[LocalVariableDeclarator]
}

func (d *LocalVariableDeclarators) Name() string { return "" }

func (d *LocalVariableDeclarators) SetName(_ string) {}

func (d *LocalVariableDeclarators) Dimension() int { return 0 }

func (d *LocalVariableDeclarators) SetDimension(_ int) {}

func (d *LocalVariableDeclarators) LineNumber() int {
	if d.Size() == 0 {
		return 0
	}

	return d.Get(0).LineNumber()
}

func (d *LocalVariableDeclarators) VariableInitializer() VariableInitializer { return nil }

func (d *LocalVariableDeclarators) AcceptDeclaration(visitor DeclarationVisitor) {
	visitor.VisitLocalVariableDeclarators(d)
}

func (d *LocalVariableDeclarators) String() string {
	return "LocalVariableDeclarators{}"
}
