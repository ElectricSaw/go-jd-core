package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

func NewFieldDeclarators() intmod.IFieldDeclarators {
	return NewFieldDeclaratorsWithCapacity(0)
}

func NewFieldDeclaratorsWithCapacity(capacity int) intmod.IFieldDeclarators {
	return &FieldDeclarators{
		DefaultList: *util.NewDefaultListWithCapacity[intmod.IFieldDeclarator](capacity),
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
		fieldDeclarator.SetFieldDeclaration(fieldDeclaration)
	}
}

func (d *FieldDeclarators) Accept(visitor intmod.IDeclarationVisitor) {
	visitor.VisitFieldDeclarators(d)
}

func (d *FieldDeclarators) String() string {
	return "FieldDeclarators{}"
}
