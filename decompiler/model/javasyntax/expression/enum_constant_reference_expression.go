package expression

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewEnumConstantReferenceExpression(typ intmod.IObjectType, name string) intmod.IEnumConstantReferenceExpression {
	return NewEnumConstantReferenceExpressionWithAll(0, typ, name)
}

func NewEnumConstantReferenceExpressionWithAll(lineNumber int, typ intmod.IObjectType, name string) intmod.IEnumConstantReferenceExpression {
	e := &EnumConstantReferenceExpression{
		AbstractLineNumberExpression: AbstractLineNumberExpression{
			lineNumber: lineNumber,
		},
		typ:  typ,
		name: name,
	}
	e.SetValue(e)
	return e
}

type EnumConstantReferenceExpression struct {
	AbstractLineNumberExpression

	typ  intmod.IObjectType
	name string
}

func (e *EnumConstantReferenceExpression) Type() intmod.IType {
	return e.typ.(intmod.IType)
}

func (e *EnumConstantReferenceExpression) SetType(typ intmod.IType) {
	e.typ = typ.(intmod.IObjectType)
}

func (e *EnumConstantReferenceExpression) ObjectType() intmod.IObjectType {
	return e.typ
}

func (e *EnumConstantReferenceExpression) Name() string {
	return e.name
}

func (e *EnumConstantReferenceExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitEnumConstantReferenceExpression(e)
}

func (e *EnumConstantReferenceExpression) String() string {
	return fmt.Sprintf("EnumConstantReferenceExpression{type=%s, name=%s}", e.typ.String(), e.name)
}
