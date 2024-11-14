package expression

import (
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewTypeReferenceDotClassExpression(typeDotClass _type.IType) *TypeReferenceDotClassExpression {
	return &TypeReferenceDotClassExpression{
		typeDotClass: typeDotClass,
		typ:          _type.OtTypeClass.CreateTypeWithArgs(typeDotClass),
	}
}

func NewTypeReferenceDotClassExpressionWithAll(lineNumber int, typeDotClass _type.IType) *TypeReferenceDotClassExpression {
	return &TypeReferenceDotClassExpression{
		lineNumber:   lineNumber,
		typeDotClass: typeDotClass,
		typ:          _type.OtTypeClass.CreateTypeWithArgs(typeDotClass),
	}
}

type TypeReferenceDotClassExpression struct {
	AbstractExpression

	lineNumber   int
	typeDotClass _type.IType
	typ          _type.IType
}

func (e *TypeReferenceDotClassExpression) LineNumber() int {
	return e.lineNumber
}

func (e *TypeReferenceDotClassExpression) GetTypeDotClass() _type.IType {
	return e.typeDotClass
}

func (e *TypeReferenceDotClassExpression) Type() _type.IType {
	return e.typ
}

func (e *TypeReferenceDotClassExpression) Priority() int {
	return 0
}

func (e *TypeReferenceDotClassExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitTypeReferenceDotClassExpression(e)
}

func (e *TypeReferenceDotClassExpression) String() string {
	return fmt.Sprintf("TypeReferenceDotClassExpression{%s}", e.typeDotClass)
}
