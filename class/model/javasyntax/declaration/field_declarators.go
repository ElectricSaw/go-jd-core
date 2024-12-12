package declaration

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/class/util"
)

func NewFieldDeclarators() intmod.IFieldDeclarators {
	return NewFieldDeclaratorsWithCapacity(0)
}

func NewFieldDeclaratorsWithCapacity(capacity int) intmod.IFieldDeclarators {
	return &FieldDeclarators{
		DefaultList: *util.NewDefaultListWithCapacity[intmod.IFieldDeclarator](capacity).(*util.DefaultList[intmod.IFieldDeclarator]),
	}
}

type FieldDeclarators struct {
	util.DefaultList[intmod.IFieldDeclarator]
}

func (d *FieldDeclarators) FieldDeclaration() intmod.IFieldDeclaration {
	return nil
}

func (d *FieldDeclarators) Name() string {
	return ""
}

func (d *FieldDeclarators) Dimension() int {
	return 0
}

func (d *FieldDeclarators) VariableInitializer() intmod.IVariableInitializer {
	return nil
}

func (d *FieldDeclarators) SetVariableInitializer(_ intmod.IVariableInitializer) {}

func (d *FieldDeclarators) SetFieldDeclaration(fieldDeclaration intmod.IFieldDeclaration) {
	for _, fieldDeclarator := range d.ToSlice() {
		if meta, ok := fieldDeclarator.(intmod.IFieldDeclarator); ok {
			meta.SetFieldDeclaration(fieldDeclaration)
		}
	}
}

func (d *FieldDeclarators) AcceptDeclaration(visitor intmod.IDeclarationVisitor) {
	visitor.VisitFieldDeclarators(d)
}

func (d *FieldDeclarators) String() string {
	return "FieldDeclarators{}"
}
