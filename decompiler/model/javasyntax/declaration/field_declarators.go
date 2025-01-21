package declaration

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewFieldDeclarators() FieldDeclarators {
	return NewFieldDeclaratorsWithCapacity(0)
}

func NewFieldDeclaratorsWithCapacity(capacity int) FieldDeclarators {
	return &FieldDeclarators{
		DefaultList: *util.NewDefaultListWithCapacity[FieldDeclarator](capacity).(*util.DefaultList[FieldDeclarator]),
	}
}

type FieldDeclarators struct {
	util.DefaultList[FieldDeclarator]
}

func (d *FieldDeclarators) FieldDeclaration() FieldDeclaration {
	return nil
}

func (d *FieldDeclarators) Name() string {
	return ""
}

func (d *FieldDeclarators) Dimension() int {
	return 0
}

func (d *FieldDeclarators) VariableInitializer() VariableInitializer {
	return nil
}

func (d *FieldDeclarators) SetVariableInitializer(_ VariableInitializer) {}

func (d *FieldDeclarators) SetFieldDeclaration(fieldDeclaration FieldDeclaration) {
	for _, fieldDeclarator := range d.ToSlice() {
		if meta, ok := fieldDeclarator.(FieldDeclarator); ok {
			meta.SetFieldDeclaration(fieldDeclaration)
		}
	}
}

func (d *FieldDeclarators) AcceptDeclaration(visitor DeclarationVisitor) {
	visitor.VisitFieldDeclarators(d)
}

func (d *FieldDeclarators) String() string {
	return "FieldDeclarators{}"
}
