package declaration

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/classfile"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewClassFileBodyDeclaration(classFile *classfile.ClassFile, bindings map[string]model.ITypeArgument,
	typeBounds map[string]model.IType, outerBodyDeclaration *ClassFileBodyDeclaration) ClassFileBodyDeclaration {
	d := ClassFileBodyDeclaration{
		DefaultBase:          *util.NewDefaultBase[model.IMemberDeclaration]().(*util.DefaultBase[model.IMemberDeclaration]),
		InternalTypeName:     classFile.InternalTypeName,
		MemberDeclaration:    nil,
		ClassFile:            classFile,
		Bindings:             bindings,
		TypeBounds:           typeBounds,
		OuterBodyDeclaration: outerBodyDeclaration,
	}
	d.SetValue(&d)
	return d
}

type ClassFileBodyDeclaration struct {
	util.DefaultBase[model.IMemberDeclaration]

	InternalTypeName         string
	MemberDeclaration        model.IMemberDeclaration
	ClassFile                *classfile.ClassFile
	FieldDeclarations        util.IList[*ClassFileFieldDeclaration]
	MethodDeclarations       util.IList[IClassFileConstructorOrMethodDeclaration]
	InnerTypeDeclarations    util.IList[IClassFileTypeDeclaration]
	InnerTypeMap             map[string]IClassFileTypeDeclaration
	FirstLineNumber          int
	OuterTypeFieldName       string
	SyntheticInnerFieldNames []string
	OuterBodyDeclaration     *ClassFileBodyDeclaration
	Bindings                 map[string]model.ITypeArgument
	TypeBounds               map[string]model.IType
}

func (d *ClassFileBodyDeclaration) GetInternalTypeName() string {
	return d.InternalTypeName
}

func (d *ClassFileBodyDeclaration) GetMemberDeclaration() model.IMemberDeclaration {
	return d.MemberDeclaration
}

func (d *ClassFileBodyDeclaration) GetFirstLineNumber() int {
	return d.FirstLineNumber
}

func (d *ClassFileBodyDeclaration) SetFieldDeclarations(fieldDeclarations util.IList[*ClassFileFieldDeclaration]) {
	if fieldDeclarations != nil {
		d.FieldDeclarations = fieldDeclarations
		tmp := make([]IClassFileMemberDeclaration, fieldDeclarations.Size())
		for i := range d.FieldDeclarations.Size() {
			tmp[i] = d.FieldDeclarations.Get(i)
		}
		d.UpdateFirstLineNumber(tmp)
	}
}

func (d *ClassFileBodyDeclaration) SetMethodDeclarations(methodDeclarations util.IList[IClassFileConstructorOrMethodDeclaration]) {
	if methodDeclarations != nil {
		d.MethodDeclarations = methodDeclarations
		tmp := make([]IClassFileMemberDeclaration, methodDeclarations.Size())
		for i := range d.MethodDeclarations.Size() {
			tmp[i] = d.MethodDeclarations.Get(i)
		}
		d.UpdateFirstLineNumber(tmp)
	}
}

func (d *ClassFileBodyDeclaration) SetInnerTypeDeclarations(innerTypeDeclarations util.IList[IClassFileTypeDeclaration]) {
	if innerTypeDeclarations != nil {
		d.InnerTypeDeclarations = innerTypeDeclarations
		tmp := make([]IClassFileMemberDeclaration, innerTypeDeclarations.Size())
		for i := range d.InnerTypeDeclarations.Size() {
			tmp[i] = d.InnerTypeDeclarations.Get(i)
		}
		d.UpdateFirstLineNumber(tmp)
		d.InnerTypeMap = make(map[string]IClassFileTypeDeclaration)
		for _, innerType := range innerTypeDeclarations.ToSlice() {
			d.InnerTypeMap[innerType.GetInternalTypeName()] = innerType
		}
	}
}

func (d *ClassFileBodyDeclaration) InnerTypeDeclaration(internalName string) IClassFileTypeDeclaration {
	decla := d.InnerTypeMap[internalName]

	if decla == nil && d.OuterBodyDeclaration != nil {
		return d.OuterBodyDeclaration.InnerTypeDeclaration(internalName)
	}

	return decla
}

func (d *ClassFileBodyDeclaration) RemoveInnerTypeDeclaration(internalName string) IClassFileTypeDeclaration {
	removed := d.InnerTypeMap[internalName]

	delete(d.InnerTypeMap, internalName)
	d.removeInnerTypeDeclaration(removed)

	return removed
}

func (d *ClassFileBodyDeclaration) removeInnerTypeDeclaration(removed IClassFileTypeDeclaration) {
	var index int
	for i, found := range d.InnerTypeDeclarations.ToSlice() {
		if removed == found {
			index = i
			break
		}
	}

	tmp := append(d.InnerTypeDeclarations.ToSlice()[:index], d.InnerTypeDeclarations.ToSlice()[index+1:]...)
	d.InnerTypeDeclarations = util.NewDefaultListWithSlice[IClassFileTypeDeclaration](tmp)
}

func (d *ClassFileBodyDeclaration) UpdateFirstLineNumber(members []IClassFileMemberDeclaration) {
	for _, member := range members {
		lineNumber := member.GetFirstLineNumber()

		if lineNumber > 0 {
			if d.FirstLineNumber == 0 {
				d.FirstLineNumber = lineNumber
			} else if d.FirstLineNumber > lineNumber {
				d.FirstLineNumber = lineNumber
			}
			break
		}
	}
}

func (d *ClassFileBodyDeclaration) IsClassDeclaration() bool {
	return false
}

func (d *ClassFileBodyDeclaration) AcceptDeclaration(visitor model.IDeclarationVisitor) {
	visitor.VisitBodyDeclaration(d)
}

func (d *ClassFileBodyDeclaration) String() string {
	return fmt.Sprintf("ClassFileBodyDeclaration{firstLineNumber=%d}", d.FirstLineNumber)
}
