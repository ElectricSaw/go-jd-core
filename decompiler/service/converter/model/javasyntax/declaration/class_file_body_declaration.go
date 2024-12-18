package declaration

import (
	"fmt"
	intcls "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/classpath"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/declaration"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewClassFileBodyDeclaration(classFile intcls.IClassFile, bindings map[string]intmod.ITypeArgument,
	typeBounds map[string]intmod.IType, outerBodyDeclaration intsrv.IClassFileBodyDeclaration) intsrv.IClassFileBodyDeclaration {
	d := &ClassFileBodyDeclaration{
		BodyDeclaration:      *declaration.NewBodyDeclaration(classFile.InternalTypeName(), nil).(*declaration.BodyDeclaration),
		classFile:            classFile,
		bindings:             bindings,
		typeBounds:           typeBounds,
		outerBodyDeclaration: outerBodyDeclaration,
	}
	d.SetValue(d)
	return d
}

type ClassFileBodyDeclaration struct {
	declaration.BodyDeclaration
	util.DefaultBase[intmod.IMemberDeclaration]

	classFile                intcls.IClassFile
	fieldDeclarations        []intsrv.IClassFileFieldDeclaration
	methodDeclarations       []intsrv.IClassFileConstructorOrMethodDeclaration
	innerTypeDeclarations    []intsrv.IClassFileTypeDeclaration
	innerTypeMap             map[string]intsrv.IClassFileTypeDeclaration
	firstLineNumber          int
	outerTypeFieldName       string
	syntheticInnerFieldNames []string
	outerBodyDeclaration     intsrv.IClassFileBodyDeclaration
	bindings                 map[string]intmod.ITypeArgument
	typeBounds               map[string]intmod.IType
}

func (d *ClassFileBodyDeclaration) FieldDeclarations() []intsrv.IClassFileFieldDeclaration {
	return d.fieldDeclarations
}

func (d *ClassFileBodyDeclaration) SetFieldDeclarations(fieldDeclarations []intsrv.IClassFileFieldDeclaration) {
	if fieldDeclarations != nil {
		d.fieldDeclarations = fieldDeclarations
		tmp := make([]intsrv.IClassFileMemberDeclaration, 0, len(fieldDeclarations))
		for _, v := range d.fieldDeclarations {
			tmp = append(tmp, v.(intsrv.IClassFileMemberDeclaration))
		}
		d.UpdateFirstLineNumber(tmp)
	}
}

func (d *ClassFileBodyDeclaration) MethodDeclarations() []intsrv.IClassFileConstructorOrMethodDeclaration {
	return d.methodDeclarations
}

func (d *ClassFileBodyDeclaration) SetMethodDeclarations(methodDeclarations []intsrv.IClassFileConstructorOrMethodDeclaration) {
	if methodDeclarations != nil {
		d.methodDeclarations = methodDeclarations
		tmp := make([]intsrv.IClassFileMemberDeclaration, 0, len(methodDeclarations))
		for _, v := range d.methodDeclarations {
			tmp = append(tmp, v)
		}
		d.UpdateFirstLineNumber(tmp)
	}
}

func (d *ClassFileBodyDeclaration) InnerTypeDeclarations() []intsrv.IClassFileTypeDeclaration {
	return d.innerTypeDeclarations
}

func (d *ClassFileBodyDeclaration) SetInnerTypeDeclarations(innerTypeDeclarations []intsrv.IClassFileTypeDeclaration) {
	if innerTypeDeclarations != nil {
		d.innerTypeDeclarations = innerTypeDeclarations
		tmp := make([]intsrv.IClassFileMemberDeclaration, 0, len(innerTypeDeclarations))
		for _, v := range d.innerTypeDeclarations {
			tmp = append(tmp, v)
		}
		d.UpdateFirstLineNumber(tmp)
		d.innerTypeMap = make(map[string]intsrv.IClassFileTypeDeclaration)
		for _, innerType := range innerTypeDeclarations {
			d.innerTypeMap[innerType.InternalTypeName()] = innerType
		}
	}
}

func (d *ClassFileBodyDeclaration) InnerTypeDeclaration(internalName string) intsrv.IClassFileTypeDeclaration {
	decla := d.innerTypeMap[internalName]

	if decla == nil && d.outerBodyDeclaration != nil {
		return d.outerBodyDeclaration.InnerTypeDeclaration(internalName)
	}

	return decla
}

func (d *ClassFileBodyDeclaration) RemoveInnerTypeDeclaration(internalName string) intsrv.IClassFileTypeDeclaration {
	removed := d.innerTypeMap[internalName]

	delete(d.innerTypeMap, internalName)
	d.removeInnerTypeDeclaration(removed)

	return removed
}

func (d *ClassFileBodyDeclaration) removeInnerTypeDeclaration(removed intsrv.IClassFileTypeDeclaration) {
	var index int
	for i, found := range d.innerTypeDeclarations {
		if removed == found {
			index = i
			break
		}
	}

	d.innerTypeDeclarations = append(d.innerTypeDeclarations[:index], d.innerTypeDeclarations[index+1:]...)
}

func (d *ClassFileBodyDeclaration) UpdateFirstLineNumber(members []intsrv.IClassFileMemberDeclaration) {
	for _, member := range members {
		lineNumber := member.FirstLineNumber()

		if lineNumber > 0 {
			if d.firstLineNumber == 0 {
				d.firstLineNumber = lineNumber
			} else if d.firstLineNumber > lineNumber {
				d.firstLineNumber = lineNumber
			}
			break
		}
	}
}

func (d *ClassFileBodyDeclaration) ClassFile() intcls.IClassFile {
	return d.classFile
}

func (d *ClassFileBodyDeclaration) FirstLineNumber() int {
	return d.firstLineNumber
}

func (d *ClassFileBodyDeclaration) OuterTypeFieldName() string {
	return d.outerTypeFieldName
}

func (d *ClassFileBodyDeclaration) SetOuterBodyDeclaration(bodyDeclaration intsrv.IClassFileBodyDeclaration) {
	d.outerBodyDeclaration = bodyDeclaration
}

func (d *ClassFileBodyDeclaration) SetOuterTypeFieldName(outerTypeFieldName string) {
	d.outerTypeFieldName = outerTypeFieldName
}

func (d *ClassFileBodyDeclaration) SyntheticInnerFieldNames() []string {
	return d.syntheticInnerFieldNames
}

func (d *ClassFileBodyDeclaration) SetSyntheticInnerFieldNames(names []string) {
	d.syntheticInnerFieldNames = names
}

func (d *ClassFileBodyDeclaration) OuterBodyDeclaration() intsrv.IClassFileBodyDeclaration {
	return d.outerBodyDeclaration
}

func (d *ClassFileBodyDeclaration) Bindings() map[string]intmod.ITypeArgument {
	return d.bindings
}

func (d *ClassFileBodyDeclaration) TypeBounds() map[string]intmod.IType {
	return d.typeBounds
}

func (d *ClassFileBodyDeclaration) IsClassDeclaration() bool {
	return false
}

func (d *ClassFileBodyDeclaration) AcceptDeclaration(visitor intmod.IDeclarationVisitor) {
	visitor.VisitBodyDeclaration(d)
}

func (d *ClassFileBodyDeclaration) String() string {
	return fmt.Sprintf("ClassFileBodyDeclaration{firstLineNumber=%d}", d.firstLineNumber)
}
