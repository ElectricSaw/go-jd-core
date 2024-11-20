package expression

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
	"fmt"
)

func NewTypeReferenceDotClassExpression(typeDotClass intsyn.IType) intsyn.ITypeReferenceDotClassExpression {
	return &TypeReferenceDotClassExpression{
		typeDotClass: typeDotClass,
		typ:          _type.OtTypeClass.CreateTypeWithArgs(typeDotClass).(intsyn.IType),
	}
}

func NewTypeReferenceDotClassExpressionWithAll(lineNumber int, typeDotClass intsyn.IType) intsyn.ITypeReferenceDotClassExpression {
	return &TypeReferenceDotClassExpression{
		lineNumber:   lineNumber,
		typeDotClass: typeDotClass,
		typ:          _type.OtTypeClass.CreateTypeWithArgs(typeDotClass).(intsyn.IType),
	}
}

type TypeReferenceDotClassExpression struct {
	AbstractExpression

	lineNumber   int
	typeDotClass intsyn.IType
	typ          intsyn.IType
}

func (e *TypeReferenceDotClassExpression) LineNumber() int {
	return e.lineNumber
}

func (e *TypeReferenceDotClassExpression) TypeDotClass() intsyn.IType {
	return e.typeDotClass
}

func (e *TypeReferenceDotClassExpression) Type() intsyn.IType {
	return e.typ
}

func (e *TypeReferenceDotClassExpression) Priority() int {
	return 0
}

func (e *TypeReferenceDotClassExpression) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitTypeReferenceDotClassExpression(e)
}

func (e *TypeReferenceDotClassExpression) String() string {
	return fmt.Sprintf("TypeReferenceDotClassExpression{%s}", e.typeDotClass)
}
