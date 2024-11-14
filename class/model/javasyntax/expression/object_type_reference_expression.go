package expression

import (
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewObjectTypeReferenceExpression(typ _type.IObjectType) *ObjectTypeReferenceExpression {
	return &ObjectTypeReferenceExpression{
		typ:      typ,
		explicit: true,
	}
}

func NewObjectTypeReferenceExpressionWithLineNumber(lineNumber int, typ _type.IObjectType) *ObjectTypeReferenceExpression {
	return &ObjectTypeReferenceExpression{
		lineNumber: lineNumber,
		typ:        typ,
		explicit:   true,
	}
}

func NewObjectTypeReferenceExpressionWithExplicit(typ _type.IObjectType, explicit bool) *ObjectTypeReferenceExpression {
	return &ObjectTypeReferenceExpression{
		typ:      typ,
		explicit: explicit,
	}
}

func NewObjectTypeReferenceExpressionWithAll(lineNumber int, typ _type.IObjectType, explicit bool) *ObjectTypeReferenceExpression {
	return &ObjectTypeReferenceExpression{
		lineNumber: lineNumber,
		typ:        typ,
		explicit:   explicit,
	}
}

type ObjectTypeReferenceExpression struct {
	AbstractExpression

	lineNumber int
	typ        _type.IObjectType
	explicit   bool
}

func (e *ObjectTypeReferenceExpression) LineNumber() int {
	return e.lineNumber
}

func (e *ObjectTypeReferenceExpression) ObjectType() _type.IObjectType {
	return e.typ
}

func (e *ObjectTypeReferenceExpression) Type() _type.IType {
	return e.typ.(_type.IType)
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

func (e *ObjectTypeReferenceExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitObjectTypeReferenceExpression(e)
}

func (e *ObjectTypeReferenceExpression) String() string {
	return fmt.Sprintf("ObjectTypeReferenceExpression{%s}", e.typ)
}
