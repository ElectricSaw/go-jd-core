package expression

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	_type "github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/type"
)

func NewTypeReferenceDotClassExpression(typeDotClass intmod.IType) intmod.ITypeReferenceDotClassExpression {
	return NewTypeReferenceDotClassExpressionWithAll(0, typeDotClass)
}

func NewTypeReferenceDotClassExpressionWithAll(lineNumber int, typeDotClass intmod.IType) intmod.ITypeReferenceDotClassExpression {
	e := &TypeReferenceDotClassExpression{
		lineNumber:   lineNumber,
		typeDotClass: typeDotClass,
		typ:          _type.OtTypeClass.CreateTypeWithArgs(typeDotClass).(intmod.IType),
	}
	e.SetValue(e)
	return e
}

type TypeReferenceDotClassExpression struct {
	AbstractExpression

	lineNumber   int
	typeDotClass intmod.IType
	typ          intmod.IType
}

func (e *TypeReferenceDotClassExpression) LineNumber() int {
	return e.lineNumber
}

func (e *TypeReferenceDotClassExpression) TypeDotClass() intmod.IType {
	return e.typeDotClass
}

func (e *TypeReferenceDotClassExpression) Type() intmod.IType {
	return e.typ
}

func (e *TypeReferenceDotClassExpression) Priority() int {
	return 0
}

func (e *TypeReferenceDotClassExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitTypeReferenceDotClassExpression(e)
}

func (e *TypeReferenceDotClassExpression) String() string {
	return fmt.Sprintf("TypeReferenceDotClassExpression{%s}", e.typeDotClass)
}
