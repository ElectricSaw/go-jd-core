package declaration

import (
	"bitbucket.org/coontec/javaClass/class/model/classfile"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/declaration"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewClassFileBodyDeclaration(classFile *classfile.ClassFile, bindings map[string]_type.ITypeArgument, typeBounds map[string]_type.IType, outerBodyDeclaration *ClassFileBodyDeclaration) *ClassFileBodyDeclaration {
	return &ClassFileBodyDeclaration{
		BodyDeclaration:      *declaration.NewBodyDeclaration(classFile.InternalTypeName(), nil),
		classFile:            classFile,
		bindings:             bindings,
		typeBounds:           typeBounds,
		outerBodyDeclaration: outerBodyDeclaration,
	}
}

type ClassFileBodyDeclaration struct {
	declaration.BodyDeclaration

	classFile                *classfile.ClassFile
	fieldDeclarations        []ClassFileFieldDeclaration
	methodDeclarations       []ClassFileConstructorOrMethodDeclaration
	innerTypeDeclarations    []ClassFileTypeDeclaration
	innerTypeMap             map[string]ClassFileTypeDeclaration
	firstLineNumber          int
	outerTypeFieldName       string
	syntheticInnerFieldNames []string
	outerBodyDeclaration     *ClassFileBodyDeclaration
	bindings                 map[string]_type.ITypeArgument
	typeBounds               map[string]_type.IType
}

func (d *ClassFileBodyDeclaration) SetMemberDeclarations(memberDeclaration declaration.IMemberDeclaration) {
	d.SetMemberDeclaration(memberDeclaration)
}

func (d *ClassFileBodyDeclaration) FieldDeclarations() []ClassFileFieldDeclaration {
	return d.fieldDeclarations
}

func (d *ClassFileBodyDeclaration) SetFieldDeclarations(fieldDeclarations []ClassFileFieldDeclaration) {
	if fieldDeclarations != nil {
		d.fieldDeclarations = fieldDeclarations
		tmp := make([]ClassFileMemberDeclaration, 0, len(fieldDeclarations))
		for _, v := range d.fieldDeclarations {
			tmp = append(tmp, &v)
		}
		d.UpdateFirstLineNumber(tmp)
	}
}

func (d *ClassFileBodyDeclaration) MethodDeclaration() []ClassFileConstructorOrMethodDeclaration {
	return d.methodDeclarations
}

func (d *ClassFileBodyDeclaration) SetMethodDeclaration(methodDeclarations []ClassFileConstructorOrMethodDeclaration) {
	if methodDeclarations != nil {
		d.methodDeclarations = methodDeclarations
		tmp := make([]ClassFileMemberDeclaration, 0, len(methodDeclarations))
		for _, v := range d.methodDeclarations {
			tmp = append(tmp, v)
		}
		d.UpdateFirstLineNumber(tmp)
	}
}

func (d *ClassFileBodyDeclaration) InnerTypeDeclarations() []ClassFileTypeDeclaration {
	return d.innerTypeDeclarations
}

func (d *ClassFileBodyDeclaration) SetInnerTypeDeclarations(innerTypeDeclarations []ClassFileTypeDeclaration) {
	if innerTypeDeclarations != nil {
		d.innerTypeDeclarations = innerTypeDeclarations
		tmp := make([]ClassFileMemberDeclaration, 0, len(innerTypeDeclarations))
		for _, v := range d.innerTypeDeclarations {
			tmp = append(tmp, v)
		}
		d.UpdateFirstLineNumber(tmp)
		d.innerTypeMap = make(map[string]ClassFileTypeDeclaration)
		for _, innerType := range innerTypeDeclarations {
			d.innerTypeMap[innerType.InternalTypeName()] = innerType
		}
	}
}

func (d *ClassFileBodyDeclaration) InnerTypeDeclaration(internalName string) ClassFileTypeDeclaration {
	decla := d.innerTypeMap[internalName]

	if decla == nil && d.outerBodyDeclaration != nil {
		return d.outerBodyDeclaration.InnerTypeDeclaration(internalName)
	}

	return decla
}

func (d *ClassFileBodyDeclaration) RemoveInnerTypeDeclaration(internalName string) ClassFileTypeDeclaration {
	removed := d.innerTypeMap[internalName]

	delete(d.innerTypeMap, internalName)
	d.removeInnerTypeDeclaration(removed)

	return removed
}

func (d *ClassFileBodyDeclaration) removeInnerTypeDeclaration(removed ClassFileTypeDeclaration) {
	var index int
	for i, found := range d.innerTypeDeclarations {
		if removed == found {
			index = i
			break
		}
	}

	d.innerTypeDeclarations = append(d.innerTypeDeclarations[:index], d.innerTypeDeclarations[index+1:]...)
}

func (d *ClassFileBodyDeclaration) UpdateFirstLineNumber(members []ClassFileMemberDeclaration) {
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

func (d *ClassFileBodyDeclaration) ClassFile() *classfile.ClassFile {
	return d.classFile
}

func (d *ClassFileBodyDeclaration) FirstLineNumber() int {
	return d.firstLineNumber
}

func (d *ClassFileBodyDeclaration) OuterTypeFieldName() string {
	return d.outerTypeFieldName
}

func (d *ClassFileBodyDeclaration) SetOuterTypeFieldName(outerTypeFieldName string) {
	d.outerTypeFieldName = outerTypeFieldName
}

func (d *ClassFileBodyDeclaration) SyntheticInnerFieldNames() []string {
	return d.syntheticInnerFieldNames
}

func (d *ClassFileBodyDeclaration) OuterBodyDeclaration() *ClassFileBodyDeclaration {
	return d.outerBodyDeclaration
}

func (d *ClassFileBodyDeclaration) Bindings() map[string]_type.ITypeArgument {
	return d.bindings
}

func (d *ClassFileBodyDeclaration) TypeBounds() map[string]_type.IType {
	return d.typeBounds
}

func (d *ClassFileBodyDeclaration) IsClassDeclaration() bool {
	return false
}

func (d *ClassFileBodyDeclaration) String() string {
	return fmt.Sprintf("ClassFileBodyDeclaration{firstLineNumber=%d}", d.firstLineNumber)
}
