package expression

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	"fmt"
)

func NewObjectTypeReferenceExpression(typ intsyn.IObjectType) intsyn.IObjectTypeReferenceExpression {
	return &ObjectTypeReferenceExpression{
		typ:      typ,
		explicit: true,
	}
}

func NewObjectTypeReferenceExpressionWithLineNumber(lineNumber int, typ intsyn.IObjectType) intsyn.IObjectTypeReferenceExpression {
	return &ObjectTypeReferenceExpression{
		lineNumber: lineNumber,
		typ:        typ,
		explicit:   true,
	}
}

func NewObjectTypeReferenceExpressionWithExplicit(typ intsyn.IObjectType, explicit bool) intsyn.IObjectTypeReferenceExpression {
	return &ObjectTypeReferenceExpression{
		typ:      typ,
		explicit: explicit,
	}
}

func NewObjectTypeReferenceExpressionWithAll(lineNumber int, typ intsyn.IObjectType, explicit bool) intsyn.IObjectTypeReferenceExpression {
	return &ObjectTypeReferenceExpression{
		lineNumber: lineNumber,
		typ:        typ,
		explicit:   explicit,
	}
}

type ObjectTypeReferenceExpression struct {
	AbstractExpression

	lineNumber int
	typ        intsyn.IObjectType
	explicit   bool
}

func (e *ObjectTypeReferenceExpression) LineNumber() int {
	return e.lineNumber
}

func (e *ObjectTypeReferenceExpression) ObjectType() intsyn.IObjectType {
	return e.typ
}

func (e *ObjectTypeReferenceExpression) Type() intsyn.IType {
	return e.typ.(intsyn.IType)
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

func (e *ObjectTypeReferenceExpression) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitObjectTypeReferenceExpression(e)
}

func (e *ObjectTypeReferenceExpression) String() string {
	return fmt.Sprintf("ObjectTypeReferenceExpression{%s}", e.typ)
}
