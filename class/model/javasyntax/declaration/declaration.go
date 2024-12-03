package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/expression"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

type AbstractMemberDeclaration struct {
	util.DefaultBase[intmod.IMemberDeclaration]
}

func (d *AbstractMemberDeclaration) IsClassDeclaration() bool {
	return false
}

func (d *AbstractMemberDeclaration) Accept(visitor intmod.IDeclarationVisitor) {

}

type AbstractTypeDeclaration struct {
	AbstractMemberDeclaration
}

func (d *AbstractTypeDeclaration) IsClassDeclaration() bool {
	return false
}

func (d *AbstractTypeDeclaration) Accept(visitor intmod.IDeclarationVisitor) {

}

func (d *AbstractTypeDeclaration) AnnotationReferences() intmod.IAnnotationReference {
	return nil
}

func (d *AbstractTypeDeclaration) Flags() int {
	return 0
}

func (d *AbstractTypeDeclaration) SetFlags(_ int) {

}

func (d *AbstractTypeDeclaration) InternalTypeName() string {
	return ""
}

func (d *AbstractTypeDeclaration) Name() string {
	return ""
}

func (d *AbstractTypeDeclaration) BodyDeclaration() intmod.IBodyDeclaration {
	return nil
}

type AbstractVariableInitializer struct {
}

func (d *AbstractVariableInitializer) LineNumber() int {
	return -1
}

func (d *AbstractVariableInitializer) IsExpressionVariableInitializer() bool {
	return false
}

func (d *AbstractVariableInitializer) Expression() intmod.IExpression {
	return expression.NeNoExpression
}

func (d *AbstractVariableInitializer) Accept(visitor intmod.IDeclarationVisitor) {

}
