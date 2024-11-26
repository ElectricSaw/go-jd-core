package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
	"fmt"
)

func NewNewInitializedArray(typ intmod.IType, arrayInitializer intmod.IArrayVariableInitializer) intmod.INewInitializedArray {
	return NewNewInitializedArrayWithAll(0, typ, arrayInitializer)
}

func NewNewInitializedArrayWithAll(lineNumber int, typ intmod.IType,
	arrayInitializer intmod.IArrayVariableInitializer) intmod.INewInitializedArray {
	e := &NewInitializedArray{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		arrayInitializer:                 arrayInitializer,
	}
	e.SetValue(e)
	return e
}

type NewInitializedArray struct {
	AbstractLineNumberTypeExpression
	util.DefaultBase[intmod.INewInitializedArray]

	arrayInitializer intmod.IArrayVariableInitializer
}

func (e *NewInitializedArray) ArrayInitializer() intmod.IArrayVariableInitializer {
	return e.arrayInitializer
}

func (e *NewInitializedArray) Priority() int {
	return 0
}

func (e *NewInitializedArray) IsNewInitializedArray() bool {
	return true
}

func (e *NewInitializedArray) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitNewInitializedArray(e)
}

func (e *NewInitializedArray) String() string {
	return fmt.Sprintf("NewInitializedArray{new %s [%s]}", e.typ, e.arrayInitializer)
}
