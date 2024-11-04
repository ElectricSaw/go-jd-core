package declaration

import _type "bitbucket.org/coontec/javaClass/java/model/javasyntax/type"

func NewArrayVariableInitializer(typ _type.IType) *ArrayVariableInitializer {
	return &ArrayVariableInitializer{
		Type: typ,
	}
}

type ArrayVariableInitializer struct {
	Type                 _type.IType
	VariableInitializers []VariableInitializer
}

func (i *ArrayVariableInitializer) GetLineNumber() int {
	if len(i.VariableInitializers) == 0 {
		return 0
	}
	return i.VariableInitializers[0].GetLineNumber()
}

func (i *ArrayVariableInitializer) IsExpressionVariableInitializer() bool {
	return false
}

func (i *ArrayVariableInitializer) GetExpression() Expression {
	return NoExpression
}
