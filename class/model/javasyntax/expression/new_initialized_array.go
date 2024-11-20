package expression

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewNewInitializedArray(typ intsyn.IType, arrayInitializer intsyn.IArrayVariableInitializer) intsyn.INewInitializedArray {
	return &NewInitializedArray{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(typ),
		arrayInitializer:                 arrayInitializer,
	}
}

func NewNewInitializedArrayWithAll(lineNumber int, typ intsyn.IType,
	arrayInitializer intsyn.IArrayVariableInitializer) intsyn.INewInitializedArray {
	return &NewInitializedArray{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		arrayInitializer:                 arrayInitializer,
	}
}

type NewInitializedArray struct {
	AbstractLineNumberTypeExpression

	arrayInitializer intsyn.IArrayVariableInitializer
}

func (e *NewInitializedArray) ArrayInitializer() intsyn.IArrayVariableInitializer {
	return e.arrayInitializer
}

func (e *NewInitializedArray) Priority() int {
	return 0
}

func (e *NewInitializedArray) IsNewInitializedArray() bool {
	return true
}

func (e *NewInitializedArray) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitNewInitializedArray(e)
}

func (e *NewInitializedArray) String() string {
	return fmt.Sprintf("NewInitializedArray{new %s [%s]}", e.typ, e.arrayInitializer)
}
