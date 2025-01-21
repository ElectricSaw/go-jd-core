package declaration

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/expression"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

type AbstractMemberDeclaration struct {
	util.DefaultBase[MemberDeclaration]
}

func (d *AbstractMemberDeclaration) IsClassDeclaration() bool {
	return false
}

func (d *AbstractMemberDeclaration) AcceptDeclaration(visitor DeclarationVisitor) {

}

type AbstractTypeDeclaration struct {
	AbstractMemberDeclaration
}

func (d *AbstractTypeDeclaration) IsClassDeclaration() bool {
	return false
}

func (d *AbstractTypeDeclaration) AcceptDeclaration(visitor DeclarationVisitor) {

}

func (d *AbstractTypeDeclaration) AnnotationReferences() AnnotationReference {
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

func (d *AbstractTypeDeclaration) BodyDeclaration() BodyDeclaration {
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

func (d *AbstractVariableInitializer) Expression() Expression {
	return expression.NeNoExpression
}

func (d *AbstractVariableInitializer) AcceptDeclaration(visitor DeclarationVisitor) {

}
