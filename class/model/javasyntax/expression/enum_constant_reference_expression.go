package expression

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	"fmt"
)

func NewEnumConstantReferenceExpression(typ intsyn.IObjectType, name string) intsyn.IEnumConstantReferenceExpression {
	return &EnumConstantReferenceExpression{
		AbstractLineNumberExpression: AbstractLineNumberExpression{
			lineNumber: intsyn.UnknownLineNumber,
		},
		typ:  typ,
		name: name,
	}
}

func NewEnumConstantReferenceExpressionWithAll(lineNumber int, typ intsyn.IObjectType, name string) intsyn.IEnumConstantReferenceExpression {
	return &EnumConstantReferenceExpression{
		AbstractLineNumberExpression: AbstractLineNumberExpression{
			lineNumber: lineNumber,
		},
		typ:  typ,
		name: name,
	}
}

type EnumConstantReferenceExpression struct {
	AbstractLineNumberExpression

	typ  intsyn.IObjectType
	name string
}

func (e *EnumConstantReferenceExpression) Type() intsyn.IType {
	return e.typ.(intsyn.IType)
}

func (e *EnumConstantReferenceExpression) SetType(typ intsyn.IType) {
	e.typ = typ.(intsyn.IObjectType)
}

func (e *EnumConstantReferenceExpression) ObjectType() intsyn.IObjectType {
	return e.typ
}

func (e *EnumConstantReferenceExpression) Name() string {
	return e.name
}

func (e *EnumConstantReferenceExpression) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitEnumConstantReferenceExpression(e)
}

func (e *EnumConstantReferenceExpression) String() string {
	return fmt.Sprintf("EnumConstantReferenceExpression{type=%s, name=%s}", e.typ.String(), e.name)
}
