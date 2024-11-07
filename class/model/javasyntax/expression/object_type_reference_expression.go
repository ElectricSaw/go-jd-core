package expression

import (
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewObjectTypeReferenceExpression(typ _type.ObjectType) *ObjectTypeReferenceExpression {
	return &ObjectTypeReferenceExpression{
		typ:      typ,
		explicit: true,
	}
}

func NewObjectTypeReferenceExpressionWithLineNumber(lineNumber int, typ _type.ObjectType) *ObjectTypeReferenceExpression {
	return &ObjectTypeReferenceExpression{
		lineNumber: lineNumber,
		typ:        typ,
		explicit:   true,
	}
}

func NewObjectTypeReferenceExpressionWithExplicit(typ _type.ObjectType, explicit bool) *ObjectTypeReferenceExpression {
	return &ObjectTypeReferenceExpression{
		typ:      typ,
		explicit: explicit,
	}
}

func NewObjectTypeReferenceExpressionWithAll(lineNumber int, typ _type.ObjectType, explicit bool) *ObjectTypeReferenceExpression {
	return &ObjectTypeReferenceExpression{
		lineNumber: lineNumber,
		typ:        typ,
		explicit:   explicit,
	}
}

type ObjectTypeReferenceExpression struct {
	AbstractExpression

	lineNumber int
	typ        _type.ObjectType
	explicit   bool
}

func (e *ObjectTypeReferenceExpression) GetLineNumber() int {
	return e.lineNumber
}

func (e *ObjectTypeReferenceExpression) GetObjectType() *_type.ObjectType {
	return &e.typ
}

func (e *ObjectTypeReferenceExpression) GetType() _type.IType {
	return &e.typ
}

func (e *ObjectTypeReferenceExpression) IsExplicit() bool {
	return e.explicit
}

func (e *ObjectTypeReferenceExpression) SetExplicit(explicit bool) {
	e.explicit = explicit
}

func (e *ObjectTypeReferenceExpression) GetPriority() int {
	return 0
}

func (e *ObjectTypeReferenceExpression) IsObjectTypeReferenceExpression() bool {
	return true
}

func (e *ObjectTypeReferenceExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitObjectTypeReferenceExpression(e)
}

func (e *ObjectTypeReferenceExpression) String() string {
	return fmt.Sprintf("ObjectTypeReferenceExpression{%s}", e.typ.String())
}
