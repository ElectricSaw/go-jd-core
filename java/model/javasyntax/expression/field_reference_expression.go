package expression

import (
	_type "bitbucket.org/coontec/javaClass/java/model/javasyntax/type"
	"fmt"
)

func NewFieldReferenceExpression(typ _type.IType, expression Expression, internalTypeName string, name string, descriptor string) *FieldReferenceExpression {
	return &FieldReferenceExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(typ),
		expression:                       expression,
		internalTypeName:                 internalTypeName,
		name:                             name,
		descriptor:                       descriptor,
	}
}

func NewFieldReferenceExpressionWithAll(lineNumber int, typ _type.IType, expression Expression, internalTypeName string, name string, descriptor string) *FieldReferenceExpression {
	return &FieldReferenceExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		expression:                       expression,
		internalTypeName:                 internalTypeName,
		name:                             name,
		descriptor:                       descriptor,
	}
}

type FieldReferenceExpression struct {
	AbstractLineNumberTypeExpression

	expression       Expression
	internalTypeName string
	name             string
	descriptor       string
}

func (e *FieldReferenceExpression) GetExpression() Expression {
	return e.expression
}

func (e *FieldReferenceExpression) GetInternalTypeName() string {
	return e.internalTypeName
}

func (e *FieldReferenceExpression) GetName() string {
	return e.name
}

func (e *FieldReferenceExpression) GetDescriptor() string {
	return e.descriptor
}

func (e *FieldReferenceExpression) SetExpression(expression Expression) {
	e.expression = expression
}

func (e *FieldReferenceExpression) SetName(name string) {
	e.name = name
}

func (e *FieldReferenceExpression) IsFieldReferenceExpression() bool {
	return true
}

func (e *FieldReferenceExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitFieldReferenceExpression(e)
}

func (e *FieldReferenceExpression) String() string {
	return fmt.Sprintf("FieldReferenceExpression{type=%s, expression=%s, name=%s, descriptor=%s}", e.typ, e.expression, e.name, e.descriptor)
}
