package expression

import (
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewEnumConstantReferenceExpression(typ _type.ObjectType, name string) *EnumConstantReferenceExpression {
	return &EnumConstantReferenceExpression{
		AbstractLineNumberExpression: AbstractLineNumberExpression{
			lineNumber: UnknownLineNumber,
		},
		typ:  typ,
		name: name,
	}
}

func NewEnumConstantReferenceExpressionWithAll(lineNumber int, typ _type.ObjectType, name string) *EnumConstantReferenceExpression {
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

	typ  _type.ObjectType
	name string
}

func (e *EnumConstantReferenceExpression) GetType() _type.IType {
	return &e.typ
}

func (e *EnumConstantReferenceExpression) GetObjectType() *_type.ObjectType {
	return &e.typ
}

func (e *EnumConstantReferenceExpression) GetName() string {
	return e.name
}

func (e *EnumConstantReferenceExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitEnumConstantReferenceExpression(e)
}

func (e *EnumConstantReferenceExpression) String() string {
	return fmt.Sprintf("EnumConstantReferenceExpression{type=%s, name=%s}", e.typ.String(), e.name)
}
