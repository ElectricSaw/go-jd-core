package declaration

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/expression"
)

type AbstractMemberDeclaration struct {
}

func (d *AbstractMemberDeclaration) IsClassDeclaration() bool {
	return false
}

func (d *AbstractMemberDeclaration) Accept(visitor intsyn.IDeclarationVisitor) {

}

type AbstractTypeDeclaration struct {
	AbstractMemberDeclaration
}

func (d *AbstractTypeDeclaration) IsClassDeclaration() bool {
	return false
}

func (d *AbstractTypeDeclaration) Accept(visitor intsyn.IDeclarationVisitor) {

}

func (d *AbstractTypeDeclaration) AnnotationReferences() intsyn.IAnnotationReference {
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

func (d *AbstractTypeDeclaration) BodyDeclaration() intsyn.IBodyDeclaration {
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

func (d *AbstractVariableInitializer) Expression() intsyn.IExpression {
	return expression.NeNoExpression
}

func (d *AbstractVariableInitializer) Accept(visitor intsyn.IDeclarationVisitor) {

}
