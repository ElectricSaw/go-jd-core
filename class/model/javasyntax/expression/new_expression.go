package expression

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewNewExpression(lineNumber int, typ intsyn.IObjectType, descriptor string) intsyn.INewExpression {
	return &NewExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpression(lineNumber),
		typ:                          typ,
		descriptor:                   descriptor,
	}
}

func NewNewExpressionWithAll(lineNumber int, typ intsyn.IObjectType, descriptor string, bodyDeclaration intsyn.IBodyDeclaration) intsyn.INewExpression {
	return &NewExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpression(lineNumber),
		typ:                          typ,
		descriptor:                   descriptor,
		bodyDeclaration:              bodyDeclaration,
	}
}

type NewExpression struct {
	AbstractLineNumberExpression

	typ             intsyn.IObjectType
	descriptor      string
	parameters      intsyn.IExpression
	bodyDeclaration intsyn.IBodyDeclaration
}

func (e *NewExpression) ObjectType() intsyn.IObjectType {
	return e.typ
}

func (e *NewExpression) SetObjectType(objectType intsyn.IObjectType) {
	e.typ = objectType
}

func (e *NewExpression) Type() intsyn.IType {
	return e.typ.(intsyn.IType)
}

func (e *NewExpression) SetType(objectType intsyn.IObjectType) {
	e.typ = objectType
}

func (e *NewExpression) Priority() int {
	return 0
}

func (e *NewExpression) Descriptor() string {
	return e.descriptor
}

func (e *NewExpression) SetDescriptor(descriptor string) {
	e.descriptor = descriptor
}

func (e *NewExpression) Parameters() intsyn.IExpression {
	return e.parameters
}

func (e *NewExpression) SetParameters(params intsyn.IExpression) {
	e.parameters = params
}

func (e *NewExpression) BodyDeclaration() intsyn.IBodyDeclaration {
	return e.bodyDeclaration
}

func (e *NewExpression) IsNewExpression() bool {
	return true
}

func (e *NewExpression) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitNewExpression(e)
}

func (e *NewExpression) String() string {
	return fmt.Sprintf("NewExpression{new %s}", e.typ)
}
