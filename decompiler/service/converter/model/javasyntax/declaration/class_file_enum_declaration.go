package declaration

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewClassFileEnumDeclaration(annotationReferences *model.AnnotationReference, flags int,
	internalTypeName, name string, interfaces model.IType,
	bodyDeclaration *ClassFileBodyDeclaration) ClassFileEnumDeclaration {
	d := ClassFileEnumDeclaration{
		AnnotationReferences: annotationReferences,
		Flags:                flags,
		InternalTypeName:     internalTypeName,
		Name:                 name,
		BodyDeclaration:      bodyDeclaration,
		Interfaces:           interfaces,
	}
	if bodyDeclaration != nil {
		d.FirstLineNumber = bodyDeclaration.FirstLineNumber
	}
	d.SetValue(&d)
	return d
}

type ClassFileEnumDeclaration struct {
	util.DefaultBase[model.IMemberDeclaration]

	AnnotationReferences *model.AnnotationReference
	Flags                int
	InternalTypeName     string
	Name                 string
	BodyDeclaration      model.IBodyDeclaration
	Interfaces           model.IType
	Constants            util.IList[model.IConstant]
	FirstLineNumber      int
}

func (d *ClassFileEnumDeclaration) GetAnnotationReferences() *model.AnnotationReference {
	return d.AnnotationReferences
}

func (d *ClassFileEnumDeclaration) GetFlag() int {
	return d.Flags
}

func (d *ClassFileEnumDeclaration) GetName() string {
	return d.Name
}

func (d *ClassFileEnumDeclaration) GetInterfaces() model.IType {
	return d.Interfaces
}

func (d *ClassFileEnumDeclaration) GetConstants() util.IList[model.IConstant] {
	return d.Constants
}

func (d *ClassFileEnumDeclaration) GetFirstLineNumber() int {
	return d.FirstLineNumber
}

func (d *ClassFileEnumDeclaration) GetInternalTypeName() string {
	return d.InternalTypeName
}

func (d *ClassFileEnumDeclaration) GetBodyDeclaration() model.IBodyDeclaration {
	return d.BodyDeclaration
}

func (d *ClassFileEnumDeclaration) IsClassDeclaration() bool {
	return false
}

func (d *ClassFileEnumDeclaration) AcceptDeclaration(visitor model.IDeclarationVisitor) {
	visitor.VisitEnumDeclaration(d)
}

func (d *ClassFileEnumDeclaration) String() string {
	return fmt.Sprintf("ClassFileEnumDeclaration{%s, firstLineNumber:%d}", d.InternalTypeName, d.FirstLineNumber)
}

func NewClassFileConstant(lineNumber int, name string, index int, arguments model.IExpression,
	bodyDeclaration model.IBodyDeclaration) ClassFileConstant {
	return ClassFileConstant{
		LineNumber:      lineNumber,
		Name:            name,
		Arguments:       arguments,
		BodyDeclaration: bodyDeclaration,
		Index:           index,
	}
}

type ClassFileConstant struct {
	LineNumber           int
	AnnotationReferences *model.AnnotationReference
	Name                 string
	Arguments            model.IExpression
	BodyDeclaration      model.IBodyDeclaration
	Index                int
}

func (c *ClassFileConstant) GetLineNumber() int {
	return c.LineNumber
}

func (c *ClassFileConstant) GetAnnotationReferences() *model.AnnotationReference {
	return c.AnnotationReferences
}

func (c *ClassFileConstant) GetName() string {
	return c.Name
}

func (c *ClassFileConstant) GetArguments() model.IExpression {
	return c.Arguments
}

func (c *ClassFileConstant) GetBodyDeclaration() model.IBodyDeclaration {
	return c.BodyDeclaration
}

func (c *ClassFileConstant) AcceptDeclaration(visitor model.IDeclarationVisitor) {
	visitor.VisitEnumDeclarationConstant(c)
}

func (c *ClassFileConstant) String() string {
	return fmt.Sprintf("ClassFileConstant{%s : %d}", c.Name, c.Index)
}
