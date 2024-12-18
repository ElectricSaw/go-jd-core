package expression

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewObjectTypeReferenceExpression(typ intmod.IObjectType) intmod.IObjectTypeReferenceExpression {
	return NewObjectTypeReferenceExpressionWithAll(0, typ, true)
}

func NewObjectTypeReferenceExpressionWithLineNumber(lineNumber int, typ intmod.IObjectType) intmod.IObjectTypeReferenceExpression {
	return NewObjectTypeReferenceExpressionWithAll(lineNumber, typ, true)
}

func NewObjectTypeReferenceExpressionWithExplicit(typ intmod.IObjectType, explicit bool) intmod.IObjectTypeReferenceExpression {
	return NewObjectTypeReferenceExpressionWithAll(0, typ, explicit)
}

func NewObjectTypeReferenceExpressionWithAll(lineNumber int, typ intmod.IObjectType, explicit bool) intmod.IObjectTypeReferenceExpression {
	e := &ObjectTypeReferenceExpression{
		lineNumber: lineNumber,
		typ:        typ,
		explicit:   explicit,
	}
	e.SetValue(e)
	return e
}

type ObjectTypeReferenceExpression struct {
	AbstractExpression

	lineNumber int
	typ        intmod.IObjectType
	explicit   bool
}

func (e *ObjectTypeReferenceExpression) LineNumber() int {
	return e.lineNumber
}

func (e *ObjectTypeReferenceExpression) ObjectType() intmod.IObjectType {
	return e.typ
}

func (e *ObjectTypeReferenceExpression) Type() intmod.IType {
	return e.typ.(intmod.IType)
}

func (e *ObjectTypeReferenceExpression) IsExplicit() bool {
	return e.explicit
}

func (e *ObjectTypeReferenceExpression) SetExplicit(explicit bool) {
	e.explicit = explicit
}

func (e *ObjectTypeReferenceExpression) Priority() int {
	return 0
}

func (e *ObjectTypeReferenceExpression) IsObjectTypeReferenceExpression() bool {
	return true
}

func (e *ObjectTypeReferenceExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitObjectTypeReferenceExpression(e)
}

func (e *ObjectTypeReferenceExpression) String() string {
	return fmt.Sprintf("ObjectTypeReferenceExpression{%s}", e.typ)
}
