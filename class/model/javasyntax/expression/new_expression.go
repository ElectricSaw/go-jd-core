package expression

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
)

func NewNewExpression(lineNumber int, typ intmod.IObjectType, descriptor string) intmod.INewExpression {
	return NewNewExpressionWithAll(lineNumber, typ, descriptor, nil)
}

func NewNewExpressionWithAll(lineNumber int, typ intmod.IObjectType, descriptor string, bodyDeclaration intmod.IBodyDeclaration) intmod.INewExpression {
	e := &NewExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpression(lineNumber),
		typ:                          typ,
		descriptor:                   descriptor,
		bodyDeclaration:              bodyDeclaration,
	}
	e.SetValue(e)
	return e
}

type NewExpression struct {
	AbstractLineNumberExpression

	typ             intmod.IObjectType
	descriptor      string
	parameters      intmod.IExpression
	bodyDeclaration intmod.IBodyDeclaration
}

func (e *NewExpression) ObjectType() intmod.IObjectType {
	return e.typ
}

func (e *NewExpression) SetObjectType(objectType intmod.IObjectType) {
	e.typ = objectType
}

func (e *NewExpression) Type() intmod.IType {
	return e.typ.(intmod.IType)
}

func (e *NewExpression) SetType(objectType intmod.IObjectType) {
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

func (e *NewExpression) Parameters() intmod.IExpression {
	return e.parameters
}

func (e *NewExpression) SetParameters(params intmod.IExpression) {
	e.parameters = params
}

func (e *NewExpression) BodyDeclaration() intmod.IBodyDeclaration {
	return e.bodyDeclaration
}

func (e *NewExpression) IsNewExpression() bool {
	return true
}

func (e *NewExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitNewExpression(e)
}

func (e *NewExpression) String() string {
	return fmt.Sprintf("NewExpression{new %s}", e.typ)
}
