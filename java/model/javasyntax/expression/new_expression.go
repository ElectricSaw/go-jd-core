package expression

import (
	"bitbucket.org/coontec/javaClass/java/model/javasyntax/declaration"
	_type "bitbucket.org/coontec/javaClass/java/model/javasyntax/type"
	"fmt"
)

func NewNewExpression(lineNumber int, typ _type.ObjectType, descriptor string) *NewExpression {
	return &NewExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpression(lineNumber),
		typ:                          typ,
		descriptor:                   descriptor,
	}
}

func NewNewExpressionWithBody(lineNumber int, typ _type.ObjectType, descriptor string, bodyDeclaration declaration.BodyDeclaration) *NewExpression {
	return &NewExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpression(lineNumber),
		typ:                          typ,
		descriptor:                   descriptor,
		bodyDeclaration:              bodyDeclaration,
	}
}

type NewExpression struct {
	AbstractLineNumberExpression

	typ             _type.ObjectType
	descriptor      string
	parameters      Expression
	bodyDeclaration declaration.BodyDeclaration
}

func (e *NewExpression) GetObjectType() *_type.ObjectType {
	return &e.typ
}

func (e *NewExpression) SetObjectType(objectType _type.ObjectType) {
	e.typ = objectType
}

func (e *NewExpression) GetType() _type.IType {
	return &e.typ
}

func (e *NewExpression) SetType(objectType _type.ObjectType) {
	e.typ = objectType
}

func (e *NewExpression) GetPriority() int {
	return 0
}

func (e *NewExpression) GetDescriptor() string {
	return e.descriptor
}

func (e *NewExpression) GetParameters() Expression {
	return e.parameters
}

func (e *NewExpression) SetParameters(params Expression) {
	e.parameters = params
}

func (e *NewExpression) GetBodyDeclaration() declaration.BodyDeclaration {
	return e.bodyDeclaration
}

func (e *NewExpression) IsNewExpression() bool {
	return true
}

func (e *NewExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitNewExpression(e)
}

func (e *NewExpression) String() string {
	return fmt.Sprintf("NewExpression{new %s}", e.typ)
}
