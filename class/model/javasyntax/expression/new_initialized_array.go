package expression

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/declaration"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewNewInitializedArray(typ _type.IType, arrayInitializer declaration.ArrayVariableInitializer) *NewInitializedArray {
	return &NewInitializedArray{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(typ),
		arrayInitializer:                 arrayInitializer,
	}
}

func NewNewInitializedArrayWithAll(lineNumber int, typ _type.IType, arrayInitializer declaration.ArrayVariableInitializer) *NewInitializedArray {
	return &NewInitializedArray{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		arrayInitializer:                 arrayInitializer,
	}
}

type NewInitializedArray struct {
	AbstractLineNumberTypeExpression

	arrayInitializer declaration.ArrayVariableInitializer
}

func (e *NewInitializedArray) ArrayInitializer() *declaration.ArrayVariableInitializer {
	return &e.arrayInitializer
}

func (e *NewInitializedArray) Priority() int {
	return 0
}

func (e *NewInitializedArray) IsNewInitializedArray() bool {
	return true
}

func (e *NewInitializedArray) Accept(visitor ExpressionVisitor) {
	visitor.VisitNewInitializedArray(e)
}

func (e *NewInitializedArray) String() string {
	return fmt.Sprintf("NewInitializedArray{new %s [%s]}", e.typ, e.arrayInitializer)
}
