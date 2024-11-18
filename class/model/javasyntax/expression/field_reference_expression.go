package expression

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	"fmt"
)

func NewFieldReferenceExpression(typ intsyn.IType, expression intsyn.IExpression,
	internalTypeName string, name string, descriptor string) intsyn.IFieldReferenceExpression {
	return &FieldReferenceExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(typ),
		expression:                       expression,
		internalTypeName:                 internalTypeName,
		name:                             name,
		descriptor:                       descriptor,
	}
}

func NewFieldReferenceExpressionWithAll(lineNumber int, typ intsyn.IType, expression intsyn.IExpression,
	internalTypeName string, name string, descriptor string) intsyn.IFieldReferenceExpression {
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

	expression       intsyn.IExpression
	internalTypeName string
	name             string
	descriptor       string
}

func (e *FieldReferenceExpression) Expression() intsyn.IExpression {
	return e.expression
}

func (e *FieldReferenceExpression) InternalTypeName() string {
	return e.internalTypeName
}

func (e *FieldReferenceExpression) Name() string {
	return e.name
}

func (e *FieldReferenceExpression) Descriptor() string {
	return e.descriptor
}

func (e *FieldReferenceExpression) SetExpression(expression intsyn.IExpression) {
	e.expression = expression
}

func (e *FieldReferenceExpression) SetName(name string) {
	e.name = name
}

func (e *FieldReferenceExpression) IsFieldReferenceExpression() bool {
	return true
}

func (e *FieldReferenceExpression) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitFieldReferenceExpression(e)
}

func (e *FieldReferenceExpression) String() string {
	return fmt.Sprintf("FieldReferenceExpression{type=%s, expression=%s, name=%s, descriptor=%s}", e.typ, e.expression, e.name, e.descriptor)
}
