package expression

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewFieldReferenceExpression(typ intmod.IType, expression intmod.IExpression,
	internalTypeName string, name string, descriptor string) intmod.IFieldReferenceExpression {
	return NewFieldReferenceExpressionWithAll(0, typ, expression, internalTypeName, name, descriptor)
}

func NewFieldReferenceExpressionWithAll(lineNumber int, typ intmod.IType, expression intmod.IExpression,
	internalTypeName string, name string, descriptor string) intmod.IFieldReferenceExpression {
	e := &FieldReferenceExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		expression:                       expression,
		internalTypeName:                 internalTypeName,
		name:                             name,
		descriptor:                       descriptor,
	}
	e.SetValue(e)
	return e
}

type FieldReferenceExpression struct {
	AbstractLineNumberTypeExpression

	expression       intmod.IExpression
	internalTypeName string
	name             string
	descriptor       string
}

func (e *FieldReferenceExpression) Expression() intmod.IExpression {
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

func (e *FieldReferenceExpression) SetExpression(expression intmod.IExpression) {
	e.expression = expression
}

func (e *FieldReferenceExpression) SetName(name string) {
	e.name = name
}

func (e *FieldReferenceExpression) IsFieldReferenceExpression() bool {
	return true
}

func (e *FieldReferenceExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitFieldReferenceExpression(e)
}

func (e *FieldReferenceExpression) String() string {
	return fmt.Sprintf("FieldReferenceExpression{type=%s, expression=%s, name=%s, descriptor=%s}", e.typ, e.expression, e.name, e.descriptor)
}
